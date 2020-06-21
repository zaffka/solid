package dependencyinversion

import "database/sql"

/*
	INCORRECT VARIANT WITH STRICTLY LINKED HIGH-LEVEL MODEL(BadStats)
	AND LOW-LEVEL MODEL(DataBase with Information)
*/

type User struct {
	name string
}

type CardNumber int

type Information struct {
	user  *User
	cards []*CardNumber
}

type DataBase struct {
	data []Information
}

type BadStats struct {
	db DataBase
}

func (s *BadStats) CardsInfo(userNames []string) []*CardNumber {
	var cardsOfUsers []*CardNumber
	for _, info := range s.db.data { //high-level search is depending on low-level data
		for _, uName := range userNames {
			if uName == info.user.name {
				cardsOfUsers = append(cardsOfUsers, info.cards...)
			}
		}
	}
	return cardsOfUsers
}

/*
	CORRECT VARIANT WITH HIGH-LEVEL MODEL(GoodStats)
	LINKED WITH LOW-LEVEL VIA INTERFACE AND REPLACEABLE REALIZATION
*/

type Searcher interface {
	FindAllCards(userName string) []*CardNumber
}

/*
	Adding a method to make our DataBase struct's behavior match the Searcher interface
*/

func (db *DataBase) FindAllCards(userName string) []*CardNumber {
	for _, info := range db.data {
		if info.user.name == userName {
			return info.cards
		}
	}

	return nil
}

type GoodStats struct {
	Searcher //replaceable by any realization like PossibleDB below
}

func (s *GoodStats) CardsInfo(userNames []string) []*CardNumber {
	var cardsOfUsers []*CardNumber
	for _, uName := range userNames {
		cards := s.FindAllCards(uName)
		if cards != nil {
			cardsOfUsers = append(cardsOfUsers, cards...)
		}
	}
	return cardsOfUsers
}

type PossibleDB struct {
	DB *sql.DB
}

func (pdb *PossibleDB) FindAllCards(userName string) []*CardNumber {
	//pdb.DB.Query("SELECT * FROM information")
	//...or anything else...

	return nil
}
