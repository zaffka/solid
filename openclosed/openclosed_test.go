package openclosed

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

var seedData = []President{
	{
		person:   "Donald",
		currency: USD,
		status:   Elected,
	},
	{
		person:   "Angela",
		currency: EUR,
		status:   InPower,
	},
	{
		person:   "Vladimir",
		currency: RUB,
		status:   InPower,
	},
	{
		person:   "Dmitry",
		currency: RUB,
		status:   Retired,
	},
	{
		person:   "Emmanuel",
		currency: EUR,
		status:   Elected,
	},
}

func TestGoodFiltering1_Filter(t *testing.T) {
	gf1 := GoodFiltering1{}
	inPower := []President{
		{
			person:   "Angela",
			currency: EUR,
			status:   InPower,
		},
		{
			person:   "Vladimir",
			currency: RUB,
			status:   InPower,
		},
	}

	donald := []President{
		{
			person:   "Donald",
			currency: USD,
			status:   Elected,
		},
	}

	got := gf1.Filter(seedData, StatusParam(InPower))
	assert.Equal(t, got, inPower)

	got = gf1.Filter(seedData, PresNameParam("Donald"))
	assert.Equal(t, got, donald)
}

func TestGoodFiltering2_Filter(t *testing.T) {
	gf2 := GoodFiltering2{}
	rubs := []President{
		{
			person:   "Vladimir",
			currency: RUB,
			status:   InPower,
		},
		{
			person:   "Dmitry",
			currency: RUB,
			status:   Retired,
		},
	}

	paramCurrencyFn := ParamFn(func(p *President) bool {
		return p.currency == RUB
	})

	got := gf2.Filter(seedData, paramCurrencyFn)
	assert.Equal(t, got, rubs)
}
