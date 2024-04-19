package main

import "fmt"

func main() {
	sl := make([]float64, 4, 4)
	sl[0] = 1
	fmt.Println(cap(sl))
	sl3 := append(sl, 5, 10)
	sl = append(sl3, -4.4)
	s2 := make([]float64, len(sl)+1)
	copy(s2, append(sl, -666.666)) // the length and capacity of slices are IMPORTANT when using copy func
	fmt.Println(sl, s2, sl3, cap(sl))
}
