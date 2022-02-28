package coordinate

import (
	"fmt"
	"testing"
)

func TestWGS84AndWebMercator(t *testing.T) {
	err := testCoordinateConvertFunc(WGS84ToWebMercator, WebMercatorToWGS84)
	if err != nil {
		t.Error(err)
	}
}

func TestWGS84ToWebMercatorTile(t *testing.T) {
	x, y, err := WGS84ToWebMercatorTile(Coordinate{116.363665, 39.913441}, 15)
	if err != nil {
		t.Error(err)
	}
	if x != 26975 || y != 12415 {
		t.Error(fmt.Errorf("x, y should be 26975, 12415"))
	}
}
