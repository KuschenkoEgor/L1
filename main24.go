package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func Distance (one,two Point) float64{
	ln := math.Sqrt(math.Pow((two.x - one.x), 2) + math.Pow((two.y - one.y), 2))
	return ln
}

func main() {
	p1 := Point{x: 1,y: 2}
	p2 := Point{x: 10,y: 2}
	ln := Distance(p1,p2)

	fmt.Println("Расстояние между точками:",ln)
}