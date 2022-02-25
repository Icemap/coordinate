# coordinate

![build_badge](https://github.com/Icemap/coordinate/workflows/Go/badge.svg)

[English](README.md) | 中文

一个Golang的地理坐标转换库，支持坐标系:

- [WGS84](https://en.wikipedia.org/wiki/World_Geodetic_System) (平常说的经纬度) 
- [Web墨卡托投影坐标系](https://en.wikipedia.org/wiki/Web_Mercator_projection)
- [墨卡托投影坐标系](https://en.wikipedia.org/wiki/Mercator_projection)
- [GCJ02](https://en.wikipedia.org/wiki/Restrictions_on_geographic_data_in_China#GCJ-02) (国测局加密坐标系/火星坐标系)

## 功能

- 在四种坐标系中互相转换

## 安装

```bash
$ go get -u github.com/Icemap/coordinate
```

## 示例

### `GCJ02`坐标到`WGS84`坐标

```go
import "github.com/Icemap/coordinate"

gcj02Point := Coordinate{X: 110.0, Y: 20.0}
wgs84Point, err := coordinate.Convert(coordinate.GCJ02, coordinate.WGS84, gcj02Point)
```

### `GCJ02`坐标到`Web墨卡托`坐标

```go
import "github.com/Icemap/coordinate"

gcj02Point := Coordinate{X: 110.0, Y: 20.0}
webMercatorPoint, err := coordinate.Convert(coordinate.GCJ02, coordinate.WebMercator, gcj02Point)
```
