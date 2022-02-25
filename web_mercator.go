package coordinate

import (
	"fmt"
	"math"
)

// WGS84ToWebMercator convert coordinate from WGS84 to web mercator
func WGS84ToWebMercator(src Coordinate) (Coordinate, error) {
	var result Coordinate

	if !src.Valid(WGS84) {
		return result, fmt.Errorf("%s: [WGS84] %+v", InvalidParam, src)
	}

	result.X = src.X*WebMercatorHalfWidth/HalfRoundDegree + WebMercatorHalfWidth
	result.Y = math.Log(math.Tan((QuarterRoundDegree+src.Y)*math.Pi/RoundDegree)) / DegreeToRadian
	result.Y = result.Y * WebMercatorHalfWidth / HalfRoundDegree
	result.Y = WebMercatorHalfWidth - result.Y

	if !result.Valid(WebMercator) {
		return result, fmt.Errorf("%s: [WebMercator] %+v", InvalidParam, src)
	}

	return result, nil
}

// WebMercatorToWGS84 convert from web mercator to WGS84
func WebMercatorToWGS84(src Coordinate) (Coordinate, error) {
	var result Coordinate

	if !src.Valid(WebMercator) {
		return result, fmt.Errorf("%s: [WebMercator] %+v", InvalidParam, src)
	}

	result.Y = WebMercatorHalfWidth - src.Y
	result.Y = result.Y * HalfRoundDegree / WebMercatorHalfWidth * DegreeToRadian
	result.Y = math.Atan(math.Exp(result.Y))*RoundDegree/math.Pi - QuarterRoundDegree

	result.X = (src.X - WebMercatorHalfWidth) * HalfRoundDegree / WebMercatorHalfWidth

	if !result.Valid(WGS84) {
		return result, fmt.Errorf("%s: [WGS84] %+v", InvalidParam, src)
	}

	return result, nil
}
