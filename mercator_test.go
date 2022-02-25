package coordinate

import (
	"testing"
)

func TestWGS84AndMercator(t *testing.T) {
	err := testCoordinateConvertFunc(WGS84ToMercator, MercatorToWGS84)
	if err != nil {
		t.Error(err)
	}
}
