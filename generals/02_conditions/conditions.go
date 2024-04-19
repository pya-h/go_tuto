package main

import "fmt"

func fibbonacci(n int64) int64 {
	var (
		x1 int64 = 0
		x2 int64 = 1
	)
	var i int64
	for i = 0; i < n; i++ {
		x3 := x2
		x2 += x1
		x1 = x3
	}

	return x1
}

func trifibbo() {
	const N = 100
	var fibboidx int64 = 1
	var width int64 = fibbonacci(fibboidx)
	var i int64
	for i = 0; i < N; i++ {
		if i >= width {
			fibboidx++
			width = fibbonacci(fibboidx)
			fmt.Println()
		}
		fmt.Printf("%d\t", fibbonacci(i))
	}
}

func primes(end int) string {
	result := ""
	for i := 2; i <= end; i++ {
		prime := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			result += fmt.Sprintf("\t%d", i)

		}
	}
	return result
}

func main() {
	allprimeslessthan1000 := primes(1000)
	fmt.Println(allprimeslessthan1000)
}
