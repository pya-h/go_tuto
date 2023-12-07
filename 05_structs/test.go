package main

import "log"

type X struct {
	Var1 int
}

func (x *X) Test() {
	x.Var1 = 10
	log.Println(x)

}

func TestX(x *X) {
	(*x).Var1 = 11
	x.Var1 = 11
	log.Println(x)
	
}

func main() {
	var x X
	x.Test()
	log.Println(x)
	TestX(&x)
	log.Println(x)

}