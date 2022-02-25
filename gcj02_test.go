package coordinate

import "testing"

func TestWGS84AndGCJ02(t *testing.T) {
	err := testCoordinateConvertFunc(WGS84ToGCJ02, GCJ02ToWGS84)
	if err != nil {
		t.Error(err)
	}
}
