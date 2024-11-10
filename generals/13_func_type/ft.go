package main

import "fmt"

type Func func(x uint)

func (f Func) Inner() {
	fmt.Println("Inner Called.")
}
func Outer(x uint) {
	fmt.Println("Outer called with", x)
}

func main() {
	Outer(10)
	f := Func(Outer)
	f.Inner()
	Func(Outer).Inner()
}
