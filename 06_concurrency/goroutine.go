package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeDelay(ms uint) {
	delay_time := time.Duration(ms)
	time.Sleep(time.Millisecond * delay_time)
}
func multiplyTable(n int, end uint) {
	for i := 1; i <= int(end); i++ {
		fmt.Println(n, " * ", i, " = ", n * i)
		makeDelay(uint(rand.Intn(250)));
	}
}

func main() {
	defer func() {
		makeDelay(5000);

		err := recover()
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
	} ()
	for i := 2; i < 10; i++ {
		go multiplyTable(i, 10)
	}
	panic("Yeah whatever")
}