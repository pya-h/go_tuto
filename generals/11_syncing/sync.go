package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mtx := new(sync.Mutex)
	fmt.Println("Test without syncing")
	for i := 0; i < 10; i++ {
                go func(i int) {
                        time.Sleep(time.Second)
                        fmt.Println("~nosync: { start: ", i, " }")
                        time.Sleep(time.Second)
                        fmt.Println("~nosync: { end: ", i, " }")
                }(i)
        }

	fmt.Println("Test with syncing through mutex lock:")
	for i := 0; i < 10; i++ {
		go func(i int) {
			mtx.Lock()
			time.Sleep(time.Second)
			fmt.Println("#synced: [ START: ", i, " ]")
			time.Sleep(time.Second)
			fmt.Println("#synced: [ END: ", i, " ]")
			mtx.Unlock()
		}(i)
	}
	var input string
	fmt.Scanln(&input)
}
