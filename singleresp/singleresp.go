package singleresp

import "database/sql"

/*
	INCORRECT VARIANT WITH "GOD OBJECT" (holds external dependencies and mixed logic)
*/

type BadEnvelope struct {
	message string
	db      *sql.DB
}

func (e *BadEnvelope) Pack(msg string) {
	e.message = msg
}

func (e *BadEnvelope) Unpack() string {
	return e.message
}

func (e *BadEnvelope) SaveToDB() error {

	//SAVING ENVELOPE's MESSAGE TO DATABASE

	return nil
}

/*
	CORRECT VARIANT:
	- Envelope's data being manipulated using struct's methods
	- Database logic separated
*/

//Envelope type and it's methods
type Envelope struct {
	message string
}

func (e *Envelope) Pack(msg string) {
	e.message = msg
}

func (e *Envelope) Unpack() string {
	return e.message
}

//DB type, responsible for saving to the persistent storage
type DB struct {
	*sql.DB
}

func (d *DB) Save(env Envelope) error {

	//SAVING ENVELOPE's MESSAGE TO DATABASE

	return nil
}
