package liskov

type FeeHandler interface {
	SetAskMarkup(int)
	AskMarkup() int

	SetBidMarkup(int)
	BidMarkup() int
}

type Currency struct {
	Ask int
	Bid int
}

func (c *Currency) SetAskMarkup(ask int) { c.Ask = ask }
func (c *Currency) AskMarkup() int       { return c.Ask }
func (c *Currency) SetBidMarkup(bid int) { c.Bid = bid }
func (c *Currency) BidMarkup() int       { return c.Bid }

//Let's imagine Stock type's markups for Ask and Bid prices have to be equal...
type Stock struct {
	Ask int
	Bid int
}

func (s *Stock) SetAskMarkup(ask int) {
	s.Ask = ask
	s.Bid = ask // ...so, we decided to also set Bid here
}

func (s *Stock) AskMarkup() int { return s.Ask }

func (s *Stock) SetBidMarkup(bid int) {
	s.Bid = bid
	s.Ask = bid // ...and Ask here
}

func (s *Stock) BidMarkup() int { return s.Bid }

func averageMarkup(ask, bid int, h FeeHandler) int {

	h.SetAskMarkup(ask)
	h.SetBidMarkup(bid)

	return (h.AskMarkup() + h.BidMarkup()) / 2
}

/*
	var currency Currency
	average(20, 10, &currency)) gives us 15

	var stock Stock
	average(20, 10, &stock) gives us 10 (!)

	In this way Stock's realization of the FeeHandler interface brokes LSP
*/
