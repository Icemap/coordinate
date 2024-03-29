package coordinate

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// Map URL
const (
	// GoogleSatelliteURL google.com/maps satellite map URL
	GoogleSatelliteURL = "http://mt[0,1,2,3].google.com/vt/lyrs=y&x={x}&y={y}&z={z}&s=Gali"
	// GoogleImageURL google.com/maps image map URL
	GoogleImageURL = "http://mt[0,1,2,3].google.com/vt/lyrs=m&gl=CN&x={x}&y={y}&z={z}&s=Gali"
	// GoogleTerrainURL google.com/maps terrain map URL
	GoogleTerrainURL = "http://mt[0,1,2,3].google.com/vt/lyrs=p&gl=CN&x={x}&y={y}&z={z}&s=Gali"
	// AMapSatelliteURL amap.com satellite map URL
	AMapSatelliteURL = "http://webst0[1,2,3,4].is.autonavi.com/appmaptile?style=6&x={x}&y={y}&z={z}"
	// AMapCoverURL amap.com cover map URL
	AMapCoverURL = "http://webst0[1,2,3,4].is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zhcn&size=1&scale=1&style=8"
	// AMapImageURL amap.com image map URL
	AMapImageURL = "http://webrd0[1,2,3,4].is.autonavi.com/appmaptile?lang=zh_cn&size=1&scale=1&style=8&x={x}&y={y}&z={z}"

	// GoogleWithoutLabelSuffix google map add this suffix to drop the label
	GoogleWithoutLabelSuffix = "&apistyle=s.t%3A0%7Cs.e%3Al%7Cp.v%3Aoff"
)

// Map Type
const (
	// GoogleSatellite google.com/maps satellite map
	GoogleSatellite = "GoogleSatellite"
	// GoogleImage google.com/maps image map
	GoogleImage = "GoogleImage"
	// GoogleTerrain google.com/maps terrain map
	GoogleTerrain = "GoogleTerrain"
	// AMapSatellite amap.com satellite map
	AMapSatellite = "AMapSatellite"
	// AMapCover amap.com cover map
	AMapCover = "AMapCover"
	// AMapImage amap.com image map
	AMapImage = "AMapImage"
)

type TiltStyle struct {
	// Google Specify Config

	// Google map tiles without label
	GoogleWithLabel bool
}

// WebMercatorTileToURLWithTiltStyle
// Convert web mercator tile number to URL, get domain randomly.
// And set the tilt style parameters.
func WebMercatorTileToURLWithTiltStyle(mapType string, x, y, z int,
	tiltStyle TiltStyle) string {
	urlModel := ""
	switch mapType {
	case GoogleSatellite:
		urlModel = GoogleSatelliteURL
	case GoogleImage:
		urlModel = GoogleImageURL
	case GoogleTerrain:
		urlModel = GoogleTerrainURL
	case AMapSatellite:
		urlModel = AMapSatelliteURL
	case AMapCover:
		urlModel = AMapCoverURL
	case AMapImage:
		urlModel = AMapImageURL
	}

	urlModel = strings.Replace(urlModel, "{x}", strconv.Itoa(x), 1)
	urlModel = strings.Replace(urlModel, "{y}", strconv.Itoa(y), 1)
	urlModel = strings.Replace(urlModel, "{z}", strconv.Itoa(z), 1)

	reg := regexp.MustCompile(`\[.*\]`)
	found := string(reg.Find([]byte(urlModel)))
	found = strings.TrimSuffix(strings.TrimPrefix(found, "["), "]")
	subArray := strings.Split(found, ",")
	urlModel = string(reg.ReplaceAll([]byte(urlModel), []byte(subArray[rand.Intn(len(subArray))])))

	// Google type
	if mapType == GoogleSatellite || mapType == GoogleImage || mapType == GoogleTerrain {
		if !tiltStyle.GoogleWithLabel {
			urlModel += GoogleWithoutLabelSuffix
		}
	}

	return urlModel
}

// WebMercatorTileToURL convert web mercator tile number to URL, get domain randomly
func WebMercatorTileToURL(mapType string, x, y, z int) string {
	return WebMercatorTileToURLWithTiltStyle(mapType, x, y, z, TiltStyle{GoogleWithLabel: true})
}
