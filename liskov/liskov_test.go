package liskov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	var currency Currency
	assert.Equal(t, 15, averageMarkup(20, 10, &currency))

	var stock Stock
	assert.Equal(t, 15, averageMarkup(20, 10, &stock)) //this fails
}
