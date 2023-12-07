package main

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"sort"
)

type Car struct {
	model string
	zero_to_100 float32
}

type Cars []Car
type ByName []Car

func (this Cars) Len() (int) {
	return len(this)
}

func (this Cars) Swap(first, second int) {
	this[first], this[second] = this[second], this[first]
}

func (this Cars) Less(first, second int) bool {
	return this[first].zero_to_100 <= this[second].zero_to_100
}

func (this ByName) Len() (int) {
	return len(this)
}

func (this ByName) Swap(first, second int) {
	this[first], this[second] = this[second], this[first]
}

func (this ByName) Less(first, second int) bool {
	return this[first].model <= this[second].model
}
func main() {
	test_err := errors.New("Test Error")
	fmt.Println(test_err)
	str_bytes := []byte("test")
	str_bytes = append(str_bytes, 65)
	str := string(str_bytes)

	fmt.Println(str, str_bytes)
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	buf.Write([]byte{65, 66})
	fmt.Println(buf, buf.String(), buf.Bytes())

	fmt.Println("Linklist test:")
	x := new(list.List)
	(*x).PushBack(10)
	x.PushBack("test")
	x.PushBack([]byte("test"))

	for item := x.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}

	cars := Cars {
		{"Ferrari", 5.0},
		{"Ford", 6.0},
		{"Porshe", 3.0},
		{"Bogati", 2.5},
	}

	sort.Sort(cars)
	fmt.Println(cars)
	sort.Sort(ByName(cars))
	fmt.Println(cars)
}