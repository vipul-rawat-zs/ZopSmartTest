package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	lat  float64
	long float64
}

func CalDistEuc(c1, c2 Coordinate) float64 {
	return math.Sqrt((c1.lat-c2.lat)*(c1.lat-c2.lat) + (c1.long-c2.long)*(c1.long-c2.long))
}

func CalDistHaversine(c1, c2 Coordinate) float64 {
	dLat := (c2.lat - c1.lat) * math.Pi / 180
	dLong := (c2.long - c1.long) * math.Pi / 180

	c2.lat = c2.lat * math.Pi / 180
	c1.lat = c1.lat * math.Pi / 180

	x := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLong/2)*math.Sin(dLong/2)*math.Cos(c1.lat)*math.Cos(c2.lat)

	return 2 * 6371 * math.Asin(math.Sqrt(x))
}

func main() {
	c1 := Coordinate{lat: 51.5007, long: 0.1246}
	c2 := Coordinate{lat: 40.6892, long: 74.0445}
	fmt.Println(CalDistEuc(c1, c2))
	fmt.Println(CalDistHaversine(c1, c2))
}
