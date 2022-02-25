package coordinate

import (
	"fmt"
	"testing"
)

var wgs84TestCaseArray = []Coordinate{
	{X: 110, Y: 20},
	{X: 0, Y: 0},
	{X: -10, Y: 10},
	{X: 10, Y: -10},
	{X: -10, Y: -10},
}

var gcj02CaseArray = []Coordinate{
	{X: 116.359425, Y: 39.936404},
	{103.849891, 36.064278},
	{103.849891, 36.064278},
	{114.032219, 22.625665},
	{114.248388, 24.684674},
}

var webMercatorCaseArray = []Coordinate{
	{WebMercatorWidth - 10000, WebMercatorWidth - 10000},
	{10000, 10000},
	{10000, WebMercatorWidth - 10000},
	{WebMercatorWidth - 10000, 10000},
}

var mercatorCaseArray = []Coordinate{
	{WebMercatorHalfWidth - 10000, WebMercatorHalfWidth - 10000},
	{-WebMercatorHalfWidth + 10000, -WebMercatorHalfWidth + 10000},
	{-WebMercatorHalfWidth + 10000, WebMercatorHalfWidth - 10000},
	{WebMercatorHalfWidth - 10000, -WebMercatorHalfWidth + 10000},
}

// testCoordinateConvertFunc test a pair of function, convert the coordinate, and reverse it
func testCoordinateConvertFunc(convertFunc, reverseFunc func(Coordinate) (Coordinate, error)) error {
	for _, testCase := range wgs84TestCaseArray {
		covertCoordinate, err := convertFunc(testCase)
		if err != nil {
			return err
		}

		result, err := reverseFunc(covertCoordinate)
		if err != nil {
			return err
		}
		if result.Equal(testCase) {
			return fmt.Errorf("not equal, convert result: %+v, test case%+v", result, testCase)
		}
	}
	return nil
}

func TestHelper(t *testing.T) {
	// each pair items in transposed matrix is reverse type of source matrix
	typeTestCaseMatrix := [][][]int{
		{{WGS84, WGS84}, {WGS84, WebMercator}, {WGS84, Mercator}, {WGS84, GCJ02}},
		{{WebMercator, WGS84}, {WebMercator, WebMercator}, {WebMercator, Mercator}, {WebMercator, GCJ02}},
		{{Mercator, WGS84}, {Mercator, WebMercator}, {Mercator, Mercator}, {Mercator, GCJ02}},
		{{GCJ02, WGS84}, {GCJ02, WebMercator}, {GCJ02, Mercator}, {GCJ02, GCJ02}},
	}

	typeCaseMap := map[int][]Coordinate{
		WGS84:       wgs84TestCaseArray,
		WebMercator: webMercatorCaseArray,
		Mercator:    mercatorCaseArray,
		GCJ02:       gcj02CaseArray,
	}

	for i := range typeTestCaseMatrix {
		for j := range typeTestCaseMatrix[i] {
			for _, testCase := range typeCaseMap[typeTestCaseMatrix[i][j][0]] {
				covertCoordinate, err := Convert(typeTestCaseMatrix[i][j][0], typeTestCaseMatrix[i][j][1], testCase)
				if err != nil {
					t.Error(err)
				}

				result, err := Convert(typeTestCaseMatrix[j][i][0], typeTestCaseMatrix[j][i][1], covertCoordinate)
				if err != nil {
					t.Error(err)
				}

				if result.Equal(testCase) {
					t.Error(fmt.Errorf("not equal, convert result: %+v, test case%+v", result, testCase))
				}
			}
		}
	}
}
