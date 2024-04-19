package main

import (
	"fmt"
	"math"
	"os"
)

func updateIt(previous *uint, prime uint) {
	*previous = prime
}

func FindPrime() func() uint {
	previousPrime := uint(1)

	return func() (prime uint) {
		defer func() { // update previous
			previousPrime = prime
		}()
		for prime = previousPrime + 1; ; prime++ {
			end := uint(math.Pow(float64(prime), 0.5))
			var j uint
			for j = 2; j <= end; j++ {
				if prime%j == 0 {
					break
				}
			}
			if j > end {
				// by ref parameter func call
				//updateIt(&previousPrime, prime)
				return
			}
		}
	}
}

func testDefer() {
	fmt.Println("I'm scheduled to run when the func returns!")
}
func main() {
	defer testDefer() // schedule the function to run even after panic or return
	defer func() {
		// error handler
		strErr := recover() // recover panic data
		fmt.Println(strErr)
	}()

	file, err := os.OpenFile("primes.txt", 1, os.FileMode(777))
	if err == nil {
		defer file.Close()
		NextPrime := FindPrime()
		howMany := 1
		fmt.Print("How many prime numbers you want to generate? ")
		ok, _ := fmt.Scanf("%d", &howMany)
		if ok == 0 {
			panic("Please fuck off")
			//return
		}

		for i := 1; i <= howMany; i++ {
			prime := NextPrime()
			fmt.Println(prime)
			maybeEndLine := ""
			if i%10 == 0 {
				maybeEndLine = "\n"
			}
			_, err := file.WriteString(fmt.Sprintf("%d\t%s", prime, maybeEndLine))
			if err != nil {
				panic("Sth went wrong while saving in the file!")
			}
		}

	} else {
		panic("Error opening primes.txt; maybe recreate it?")
	}

}
