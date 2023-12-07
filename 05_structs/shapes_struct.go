package main

import ("fmt";
	"math"
)

type Circle struct {
	x, y, r float64
}

func circle_area(c Circle) float64 {
	return math.Pi * r * r
}

func main() {
	var c Circle
	c2 := new(Circle)
	c3 := Circle {x: 0, r: 3, y: 2}
	c4 := Circle {0, 2, 3}
	fmt.Println(c, *c2, c3)
}