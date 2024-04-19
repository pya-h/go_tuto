package main

import (
	"fmt"
	"time"
)

func pinger(test_channel chan string) {
	for {
		test_channel <- "fuck"
	}
}

func ponger(test_channel chan string) {
	for {
		test_channel <- "U"
	}
}

func printer(test_channel chan string) {
	for {
		recv := <-test_channel
		fmt.Println(recv)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var tst chan string = make(chan string)
	go pinger(tst)
	go ponger(tst)
	go printer(tst)

	var inp string
	fmt.Scanln(&inp)

}
