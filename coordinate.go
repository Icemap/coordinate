package coordinate

import "math"

type (
	// Coordinate point struct
	Coordinate struct {
		X float64
		Y float64
	}

	// Boundary coordinate system boundary struct
	Boundary struct {
		Min Coordinate
		Max Coordinate
	}
)

// BoundaryArray coordinate system boundary array
var BoundaryArray = []Boundary{
	{Min: Coordinate{-HalfRoundDegree, -QuarterRoundDegree}, Max: Coordinate{HalfRoundDegree, QuarterRoundDegree}},
	{Min: Coordinate{0, 0}, Max: Coordinate{WebMercatorWidth, WebMercatorWidth}},
	{Min: Coordinate{-WebMercatorHalfWidth, -WebMercatorHalfWidth}, Max: Coordinate{WebMercatorHalfWidth, WebMercatorHalfWidth}},
	{Min: Coordinate{-HalfRoundDegree, -QuarterRoundDegree}, Max: Coordinate{HalfRoundDegree, QuarterRoundDegree}},
}

// ChinaBoundary roughly rectangle boundary of china
var ChinaBoundary = Boundary{Min: Coordinate{ChinaMinLongitude, ChinaMinLatitude}, Max: Coordinate{ChinaMaxLongitude, ChinaMaxLatitude}}

// EqualWithAccuracy coordinate equal judge with parametric accuracy
func (coordinate *Coordinate) EqualWithAccuracy(another Coordinate, accuracy float64) bool {
	return math.Abs(coordinate.X-another.X) < accuracy && math.Abs(coordinate.Y-another.Y) < accuracy
}

// Equal coordinate equal judge with default accuracy (10 ^ -6)
func (coordinate *Coordinate) Equal(another Coordinate) bool {
	return coordinate.EqualWithAccuracy(another, 10^-6)
}

// Valid check coordinate valid in coordinate system
func (coordinate *Coordinate) Valid(coordinateSystem int) bool {
	boundary := BoundaryArray[coordinateSystem]
	return coordinate.In(boundary)
}

// In check if coordinate in the boundary
func (coordinate *Coordinate) In(boundary Boundary) bool {
	return coordinate.X >= boundary.Min.X && coordinate.X <= boundary.Max.X &&
		coordinate.Y >= boundary.Min.Y && coordinate.Y <= boundary.Max.Y
}
