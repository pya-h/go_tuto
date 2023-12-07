package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List // only normal access allowed

	x.PushBack(10)
	// (*x).PushBack(10)   // WRONG: cannot indirect a list in this way

	x.PushBack(2.3)
	x.PushBack("String")

	for y := x.Front(); y != nil; y = y.Next() {
		fmt.Println(y.Value)
	}

	x2 := new(list.List)  // normal & indirect access allowed
	(*x2).PushBack(3)
	x2.PushBack(234)
	for y := x2.Front(); y != nil; y = y.Next() {
		fmt.Println(y.Value.(int))
	}
}
