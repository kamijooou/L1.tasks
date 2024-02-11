package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func dist(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1, p2 := new(Point), new(Point)
	p1.x, p1.y = 2, 7
	p2.x, p2.y = 5, 4
	fmt.Println(dist(p1, p2))
}
