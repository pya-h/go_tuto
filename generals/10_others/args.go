package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 6, "the max value: ")
	another := flag.Int("another", 10, "thi another value: ")

	flag.Parse()

	fmt.Println(rand.Intn(*maxp))

	fmt.Println(flag.Args(), *another)
}
