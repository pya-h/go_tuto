package main

import (
	"fmt"
	"math"
)

var globalvar string = "Global variables must always be type specified; [var GLOBALVAR type = ..."

func Globaltest() {
	fmt.Println(globalvar)
}

func main() {
	fmt.Println(`hellowwwww

		asshole!`)

	//var x int = 50
	var str1 string = "test"
	var str2 string = "test"
	fmt.Println(str1 == str2)
	x := "creating new variable with starting value "
	var y = 22
	fmt.Printf("another shapes of var declaration => %s %d\n", x, y)
	Globaltest()

	const constantone = 10
	const complextest complex128 = 5 + 6i

	zi := complextest * (7 + -2i)
	fmt.Printf(`complextest is %f ,
		and zi is then = %f
	`, complextest, zi)

	labledpi := "math.Pi here is " +
		fmt.Sprintf("%f", math.Pi)
	fmt.Println("convert f to str and concat ==>" + labledpi)

	// multiple var definme
	var (
		x1 = 1
		x2 = 1
		x3 = 2
	)

	fmt.Println(x1, x2, x3)
}
