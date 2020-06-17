package openclosed

type Currency int
type Status int

const (
	Unknown Currency = iota
	USD
	EUR
	RUB

	InPower Status = iota
	Retired
	Elected
)

type President struct {
	person   string
	currency Currency
	status   Status
}

/*
	INCORRECT VARIANT WITH EXCESSIVE FILTERING METHODS
	IT WILL BE DIFFICULT TO SUPPORT AND EXTEND SUCH CODE
*/

type BadFiltering struct {
	//some dependencies
}

func (*BadFiltering) ListByStatus(presidents []President, status Status) []President {
	var res []President
	for _, p := range presidents {
		if p.status == status {
			res = append(res, p)
		}
	}
	return res
}

func (*BadFiltering) ListByCurrency(presidents []President, currency Currency) []President {
	var res []President
	for _, p := range presidents {
		if p.currency == currency {
			res = append(res, p)
		}
	}
	return res
}

/*
	GOOD VARIANT #1 WITH FILTERING USING INTERFACES
*/

//GoodFiltering1 - main filtering struct
type GoodFiltering1 struct {
	//some dependencies
}

//Parametr interface - lets us use parameterized filtering, without changing any filtering methods
type Parametr interface {
	Match(*President) bool
}

//Filter - it's just one method, which gets any kind of filtering params without modifying any logic
func (*GoodFiltering1) Filter(presidents []President, param Parametr) []President {
	var res []President
	for _, pres := range presidents {
		if param.Match(&pres) {
			res = append(res, pres)
		}

	}
	return res
}

//StatusParam - specification for filtering by Status
type StatusParam Status

func (sp StatusParam) Match(p *President) bool {
	return Status(sp) == p.status
}

//CurrencyParam - specification for filtering by Currency
type CurrencyParam Currency

func (cp CurrencyParam) Match(p *President) bool {
	return Currency(cp) == p.currency
}

//PresNameParam - specification for filtering by president's name
type PresNameParam string

func (pn PresNameParam) Match(p *President) bool {
	return string(pn) == p.person
}

/*
	GOOD VARIANT #2 WITH FILTERING USING PARAM FUNCS
*/

type ParamFn func(*President) bool

//GoodFiltering1 - main filtering struct
type GoodFiltering2 struct {
	//some dependencies
}

//Filter - it's just one method, which gets any kind of filtering params without modifying any logic
func (*GoodFiltering2) Filter(presidents []President, param ParamFn) []President {
	var res []President
	for _, pres := range presidents {
		if param(&pres) {
			res = append(res, pres)
		}
	}
	return res
}

/*
	As an example:

	gf2 := GoodFiltering2{}

	paramCurrencyFn := ParamFn(func(p *President) bool {
		return p.currency == RUB
	})

	result := gf2.Filter(seedData, paramCurrencyFn)
*/
