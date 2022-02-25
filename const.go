package coordinate

import "math"

const (
	// WGS84 coordinate system WGS84
	WGS84 = iota
	// WebMercator coordinate system web mercator
	WebMercator
	// Mercator coordinate system mercator
	Mercator
	// GCJ02 coordinate system GCJ02 (a special coordinate system in China)
	GCJ02
)

const (
	// WebMercatorWidth web mercator projection width or height
	WebMercatorWidth = WebMercatorHalfWidth * 2

	// WebMercatorHalfWidth half of web mercator projection width or height
	WebMercatorHalfWidth = 20037508.3427892

	// RoundDegree a round degree num
	RoundDegree = 360.0

	// HalfRoundDegree half round degree num
	HalfRoundDegree = RoundDegree / 2

	// QuarterRoundDegree quarter round degree num
	QuarterRoundDegree = HalfRoundDegree / 2

	// DegreeToRadian degree to radian constant
	DegreeToRadian = math.Pi / HalfRoundDegree

	// GCJ02EE a weird param "ee" in GCJ02
	GCJ02EE = 0.006693421622965943

	// GCJ02A a weird param "a" in GCJ02
	GCJ02A = 6378245.0
)

const (
	// ChinaMinLongitude min longitude of china
	ChinaMinLongitude = 72.004000000000005

	// ChinaMaxLongitude max longitude of china
	ChinaMaxLongitude = 137.8347

	// ChinaMinLatitude min latitude of china
	ChinaMinLatitude = 0.8293

	// ChinaMaxLatitude max latitude of china
	ChinaMaxLatitude = 55.827100000000002
)

const (
	// InvalidParam has invalid param here
	InvalidParam = "invalid param"

	// InvalidCoordinateSystemType coordinate system invalid here
	InvalidCoordinateSystemType = "invalid coordinate system"
)
