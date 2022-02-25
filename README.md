# coordinate

![build_badge](https://github.com/Icemap/coordinate/workflows/Go/badge.svg)

English | [中文](README_ch.md)

A golang library for geographic coordinate conversions, supporting the following coordinate systems:

- [WGS84](https://en.wikipedia.org/wiki/World_Geodetic_System)
- [Web Mercator](https://en.wikipedia.org/wiki/Web_Mercator_projection)
- [Mercator](https://en.wikipedia.org/wiki/Mercator_projection)
- [GCJ02](https://en.wikipedia.org/wiki/Restrictions_on_geographic_data_in_China#GCJ-02)

## Feature

- Convert coordinates between this coordinate systems

## Install

```bash
$ go get -u github.com/Icemap/coordinate
```

## Example

### `GCJ02` coordinate system to `WGS84` coordinate system

```go
import "github.com/Icemap/coordinate"

gcj02Point := Coordinate{X: 110.0, Y: 20.0}
wgs84Point, err := coordinate.Convert(coordinate.GCJ02, coordinate.WGS84, gcj02Point)
```

### `GCJ02` coordinate system to `Web Mercator` coordinate system

```go
import "github.com/Icemap/coordinate"

gcj02Point := Coordinate{X: 110.0, Y: 20.0}
webMercatorPoint, err := coordinate.Convert(coordinate.GCJ02, coordinate.WebMercator, gcj02Point)
```
