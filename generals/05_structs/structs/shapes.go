package main

import ("fmt";
	"math"
)
type Shape interface {
	area() float64
}
type Curve struct {
	name string
	group string
}

type Circle struct {
	Curve
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	Curve
	x1, x2, y1, y2 float64
}

func (rect *Rectangle) area () float64 {
	var width, length float64
	if length = rect.x2 - rect.x1; length < 0 {
		length *= -1
	}

	if width = rect.y2 - rect.y1; width < 0 {
		width *= -1
	}
	return width * length

}

func (s *Curve) describe () string {
	return fmt.Sprint("Curve is a ", s.group, " named: ", s.name)
}

func totalArea(shapes ...Shape) (total float64) {
	total = 0
	for _, s := range shapes {
		total += s.area()
	}
	return
}

type ComplexShape struct {
	segments []Shape
}

func (cs *ComplexShape) area() (area float64) {
	area = 0
	for _, shape := range cs.segments {
		area += shape.area()
	}
	return
}
func main() {
	var c Circle
	c2 := new(Circle)
	c3 := Circle {Curve: Curve {"c3", "circle"}, x: 0, r: 3, y: 2}
	c4 := Circle {Curve{"c4", "circle"}, 0, 2, 3}
	fmt.Println(c, *c2, c3, c4)
	fmt.Println( c3,c3.describe(), "with area of ", c3.area())
	r1 := Rectangle{Curve: Curve{name: "R1", group: "rectangle"}, x1: 0, x2: 10, y1: 0, y2: 5}
	fmt.Println(r1, r1.describe(), "with area of ",  r1.area())
	fmt.Println(totalArea(&c4, &c3, c2, &c, &r1))
	cs := ComplexShape { segments: []Shape{&c3, &r1 }}
	fmt.Println("Complex Shape area is: ", cs.area())
}