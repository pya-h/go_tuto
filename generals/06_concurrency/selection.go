package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	c1 := make(chan string, 1) // buffered or async channel
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "first one"
			time.Sleep(time.Second * 2)
		}
	} ()

	go func() {
		for {
			c2 <- "second one"
			time.Sleep(time.Second * 3)
		}
	} ()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("CASE_1:", msg1)

		case msg2 := <-c2:
			fmt.Println("CASE_2:", msg2)

		case msg3 := <- time.After(time.Second): // what to do when one seconds passes while blocking
//		case msg3 <- time.After(time.Second):
			fmt.Println("timeout", msg3)
/*		default:
			fmt.Println("Nadda")
			time.Sleep(time.Second)*/
		}
	}

	var inp string
	fmt.Scanln(&inp)

}