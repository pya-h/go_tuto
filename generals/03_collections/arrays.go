package main

import "fmt"

func average_old(scores []float64) float64 {
	sum := 0.0
	length := len(scores)
	for i := 0; i < length; i++ {
		sum += scores[i]
	}
	return sum / float64(length)

}

func average(scores []float64) float64 {
	sum := 0.0
	for _, value := range scores {
		sum += value
	}
	return sum / float64(len(scores))

}

func testaverage() {
	var x [5]float64
	for i := 0; i < 5; i++ {
		fmt.Scanf("%f", &x[i])
	}
	fmt.Println("average of ", x, " = ", average(x[:]))
}

func main() {
	testaverage()

	another := [...]float64{1, 2, 3, 4}
	fmt.Println("average of", another, " = ", average(another[:]))
}
