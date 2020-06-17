package liskov

type FeeHandler interface {
	SetAsk(int)
	GetAsk() int

	SetBid(int)
	GetBid() int
}

type Currency struct {
	Ask int
	Bid int
}

func (c *Currency) SetAsk(ask int) { c.Ask = ask }
func (c *Currency) GetAsk() int    { return c.Ask }
func (c *Currency) SetBid(bid int) { c.Bid = bid }
func (c *Currency) GetBid() int    { return c.Bid }

//Let's imagine Stock type's Ask has to be equal to Bid...
type Stock struct {
	Ask int
	Bid int
}

func (s *Stock) SetAsk(ask int) {
	s.Ask = ask
	s.Bid = ask // ...so, we decided to also set Bid here
}

func (s *Stock) GetAsk() int { return s.Ask }

func (s *Stock) SetBid(bid int) {
	s.Bid = bid
	s.Ask = bid // ...and Ask here
}

func (s *Stock) GetBid() int { return s.Bid }

func average(ask, bid int, h FeeHandler) int {

	h.SetAsk(ask)
	h.SetBid(bid)

	return (h.GetAsk() + h.GetBid()) / 2
}

/*
	var currency Currency
	average(20, 10, &currency)) gives us 15

	var stock Stock
	average(20, 10, &stock) gives us 10 (!)

	In this way Stock's realization of the FeeHandler interface brokes LSP
*/
