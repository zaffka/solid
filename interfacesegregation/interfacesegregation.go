package interfacesegregation

type Parcel struct{}

/*
	INCORRECT VARIANT
	ONE INTERFACE REALIZATION(Postomat) FORCED TO HAVE UNNECESSARY METHOD
*/

type Handler interface {
	Receive() Parcel
	Send(Parcel)
}

type Mailman struct{}

func (m *Mailman) Receive() Parcel {
	return Parcel{}
}

func (m *Mailman) Send(p Parcel) {
	//get parcel and deliver
}

type Postomat struct{}

func (p *Postomat) Receive() Parcel {
	return Parcel{}
}

func (p *Postomat) Send(prc Parcel) {
	//you can't send parcel using a postomat
	panic("it is not possible")
}

/*
	CORRECT VARIANT
	HANDLER INTERFACE HAS BEEN TAKEN APART TO SMALLER ONES
*/

type GoodHandler interface {
	Receiver
	Sender
}

type Receiver interface {
	Receive() Parcel
}

type Sender interface {
	Send(Parcel)
}

//Postbox realizes only the Receiver interface
type Postbox struct{}

func (pb *Postbox) Receive() Parcel {
	return Parcel{}
}

//Courier realizes whole GoodHandler interface
type Courier struct{}

func (c *Courier) Receive() Parcel {
	return Parcel{}
}

func (c *Courier) Send(p Parcel) {
	//get parcel and deliver
}
