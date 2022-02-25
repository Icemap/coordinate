package coordinate

import (
	"fmt"
	"math"
)

// WGS84ToMercator convert coordinate from WGS84 to mercator
func WGS84ToMercator(src Coordinate) (Coordinate, error) {
	var result Coordinate

	if !src.Valid(WGS84) {
		return result, fmt.Errorf("%s: [WGS84] %+v", InvalidParam, src)
	}

	result.X = src.X * WebMercatorHalfWidth / HalfRoundDegree
	result.Y = math.Log(math.Tan((QuarterRoundDegree+src.Y)*math.Pi/RoundDegree)) / DegreeToRadian
	result.Y = result.Y * WebMercatorHalfWidth / HalfRoundDegree

	if !result.Valid(Mercator) {
		return result, fmt.Errorf("%s: [Mercator] %+v", InvalidParam, src)
	}

	return result, nil
}

// MercatorToWGS84 convert from mercator to WGS84
func MercatorToWGS84(src Coordinate) (Coordinate, error) {
	var result Coordinate

	if !src.Valid(Mercator) {
		return result, fmt.Errorf("%s: [Mercator] %+v", InvalidParam, src)
	}

	result.Y = result.Y * HalfRoundDegree / WebMercatorHalfWidth * DegreeToRadian
	result.Y = math.Atan(math.Exp(result.Y))*RoundDegree/math.Pi - QuarterRoundDegree

	result.X = (src.X) * HalfRoundDegree / WebMercatorHalfWidth

	if !result.Valid(WGS84) {
		return result, fmt.Errorf("%s: [WGS84] %+v", InvalidParam, src)
	}

	return result, nil
}
