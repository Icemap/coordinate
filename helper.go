package coordinate

import "fmt"

type ConvertFunc func(Coordinate) (Coordinate, error)

var fromMap = map[int]ConvertFunc{
	WGS84:       doNothing,
	WebMercator: WebMercatorToWGS84,
	Mercator:    MercatorToWGS84,
	GCJ02:       GCJ02ToWGS84,
}

var toMap = map[int]ConvertFunc{
	WGS84:       doNothing,
	WebMercator: WGS84ToWebMercator,
	Mercator:    WGS84ToMercator,
	GCJ02:       WGS84ToGCJ02,
}

// Convert high level convert function , convert coordinate by convert type
func Convert(from, to int, coordinate Coordinate) (Coordinate, error) {
	if fromFunc, fromExist := fromMap[from]; fromExist {
		if toFunc, toExist := toMap[to]; toExist {
			wgs84Coordinate, err := fromFunc(coordinate)
			if err != nil {
				return Coordinate{}, err
			}
			targetCoordinate, err := toFunc(wgs84Coordinate)
			if err != nil {
				return Coordinate{}, err
			}

			return targetCoordinate, nil
		}
	}

	return Coordinate{}, fmt.Errorf("%s: from: %d, to: %d", InvalidCoordinateSystemType, from, to)
}

// doNothing do nothing function, an adaptation to wgs84 to wgs84
func doNothing(src Coordinate) (Coordinate, error) {
	return src, nil
}
