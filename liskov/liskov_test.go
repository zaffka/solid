package liskov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	var currency Currency
	assert.Equal(t, 15, average(20, 10, &currency))

	var stock Stock
	assert.Equal(t, 10, average(20, 10, &stock)) //has to be 15 as well
}
