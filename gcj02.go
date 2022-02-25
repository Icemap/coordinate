package coordinate

import (
	"fmt"
	"math"
)

// Do not to try to understand this part of code
// It's only a rule to convert coordinate in China by government
// Forget it, and treat it like a magic, nobody knows why

// WGS84ToGCJ02 convert coordinate from WGS84 to GCJ02
func WGS84ToGCJ02(src Coordinate) (Coordinate, error) {
	var result Coordinate

	if !src.Valid(WGS84) {
		return result, fmt.Errorf("%s: [WGS84] %+v", InvalidParam, src)
	}

	// gcj02 effect in China only
	if !src.In(ChinaBoundary) {
		return src, nil
	}

	offset := calculateOffset(src)
	result = Coordinate{
		X: src.X + offset.X,
		Y: src.Y + offset.Y,
	}

	if !result.Valid(GCJ02) {
		return result, fmt.Errorf("%s: [GCJ02] %+v", InvalidParam, src)
	}

	return result, nil
}

// GCJ02ToWGS84 convert from GCJ02 to WGS84
func GCJ02ToWGS84(src Coordinate) (Coordinate, error) {
	var result Coordinate

	if !src.Valid(GCJ02) {
		return result, fmt.Errorf("%s: [GCJ02] %+v", InvalidParam, src)
	}

	// gcj02 effect in China only
	if !src.In(ChinaBoundary) {
		return src, nil
	}

	offset := calculateOffset(src)
	result = Coordinate{
		X: src.X - offset.X,
		Y: src.Y - offset.Y,
	}

	if !result.Valid(WGS84) {
		return result, fmt.Errorf("%s: [WGS84] %+v", InvalidParam, src)
	}

	return result, nil
}

func calculateOffset(src Coordinate) Coordinate {
	offsetX := transformLon(Coordinate{X: src.X - 105, Y: src.Y - 35})
	offsetY := transformLat(Coordinate{X: src.X - 105, Y: src.Y - 35})
	radianY := offsetY / HalfRoundDegree * math.Pi

	magicNumber := math.Sin(radianY)
	magicNumber = 1.0 - GCJ02EE*magicNumber*magicNumber
	sqrtMagic := math.Sqrt(magicNumber)

	offsetX = offsetX * HalfRoundDegree / (GCJ02A / sqrtMagic * math.Cos(radianY) * math.Pi)
	offsetY = offsetY * HalfRoundDegree / (GCJ02A * (1.0 - GCJ02EE) / (magicNumber * sqrtMagic) * math.Pi)

	return Coordinate{offsetX, offsetY}
}

func transformLat(src Coordinate) float64 {
	ret := -100.0 + 2.0*src.X + 3.0*src.Y + 0.2*src.Y*src.Y + 0.1*src.X*src.Y + 0.2*math.Sqrt(math.Abs(src.X))
	ret += (20.0*math.Sin(6.0*src.X*math.Pi) + 20.0*math.Sin(2.0*src.X*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(src.Y*math.Pi) + 40.0*math.Sin(src.Y/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(src.Y/12.0*math.Pi) + 320.0*math.Sin(src.Y*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLon(src Coordinate) float64 {
	ret := 300.0 + src.X + 2.0*src.Y + 0.1*src.X*src.X + 0.1*src.X*src.Y + 0.1*math.Sqrt(math.Abs(src.X))
	ret += (20.0*math.Sin(6.0*src.X*math.Pi) + 20.0*math.Sin(2.0*src.X*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(src.X*math.Pi) + 40.0*math.Sin(src.X/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(src.X/12.0*math.Pi) + 300.0*math.Sin(src.X/30.0*math.Pi)) * 2.0 / 3.0

	return ret
}
