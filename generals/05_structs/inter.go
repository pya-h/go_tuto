package main

import "fmt"

type Base interface {
	Whatever()
}

type Derived1 struct {
	x int
}

func (self Derived1) Whatever() {
	fmt.Println("This is from derived 1")
}

type Derived2 struct {
	y int
}

func (self Derived2) Whatever() {
	fmt.Println("this is from derived 2")
}

func Gather(x Base) {
	x.Whatever()
}

func main() {

	x := new(Derived1)
	Gather(x)
	y := new(Derived2)
	Gather(y)
}
