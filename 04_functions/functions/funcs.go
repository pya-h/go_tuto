package main

import "fmt"

func imaxArr(nums []float64) [2]float64 {
	max := nums[0]
	idx := 0
	for i, v := range nums {
		if v > max {
			idx = i
			max = v
		}
	}
	return [...]float64{float64(idx), max} // [...] means count the array sizer by counting the init values between { and }
	// [...] differs with [], [...] is fixed size, but [] is undetermined size
}

func imax(nums []float64) (idx int, max float64) {
	max = nums[0]
	idx = 0
	for i, v := range nums {
		if v > max {
			idx = i
			max = v
		}
	}
	return
}

type GenericNumber interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

func GenericSum[Number GenericNumber](nums ...[]Number) Number {
	var sum Number = 0.0

	for _, num := range nums {
		for _, v := range num {
			sum += v
		}
	}

	return sum
}

func product[Number GenericNumber](x ...Number) (product Number) {
	product = 1

	for _, v := range x {
		product *= v
	}
	return
}

func main() {

	x := make([]int, 3, 3)
	x = append(x, 1)
	var y map[string]int = make(map[string]int)
	y["one"] = 1
	y["two"] = 2

	z := map[string]int{
		"z": 1,
	}
	z["x"] = x[0]
	fmt.Println(z)
	fmt.Println(imaxArr([]float64{10, -2, 3, 1, 11, 0, -4, -1}))
	index, max := imax([]float64{10, -2, 3, 1, 11, 0, -4, -1})
	fmt.Printf("func can return multiple values with single types! and also name the return vars! cool %d, %.2f\n", index, max)

	fmt.Println("Test the GenericSum function:")
	test := [...]int{1, 2, 3, 4, 5}
	fmt.Println("sum(test) = ", GenericSum(test[:]))
	testi := [...]int{10, 20, 30, 40, 50}
	fmt.Println("sum(test) = ", GenericSum(test[:], testi[:]))

	test2 := []float32{1.2, 3.2, 22.33, 54.1, -2.3, -22}
	fmt.Println("sum(test2) = ", GenericSum(test2))

	nonlocal := [...]complex128{2, 2.3, 2 + 3i}
	// INNER FUNC
	innerProduct := func() (product complex128) {
		product = 1

		for _, v := range nonlocal {
			product *= v
		}
		return
	}

	fmt.Println(innerProduct())
	fmt.Println(product(2+0i, 2-3i, 2+3i))

}
