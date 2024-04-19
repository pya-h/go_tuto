package main

import ("fmt"; "math")
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)

}

func rectArea(x1, y1, x2, y2 float64) float64 {
	length := distance(x1, y1, x1, y2)
	width := distance(x1, y1, x2, y1)
	return length * width
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10,10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Printf("(%.2f, %.2f) -> (%.2f, %.2f) :=> S = %.4f\n", rx1, ry1, rx2, ry2, rectArea(rx1, ry1, rx2, ry2))
	fmt.Printf("C(%.2f, %.2f), R = %.2f :=> S = %.4f\n", cx, cy, cr, circleArea(cx, cy, cr))
}