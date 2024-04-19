package main

import (
	"fmt"
	"stack"
)

func main() {
	s := stack.NewStack()
	s.Push("test")
	fmt.Println(s)
}
