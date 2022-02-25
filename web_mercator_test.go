package coordinate

import (
	"testing"
)

func TestWGS84AndWebMercator(t *testing.T) {
	err := testCoordinateConvertFunc(WGS84ToWebMercator, WebMercatorToWGS84)
	if err != nil {
		t.Error(err)
	}
}
