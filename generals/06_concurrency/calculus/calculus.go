package calculus

func Average(numbers []float64) (result float64) {
	result = 0

	for _, v := range numbers {
		result += v
	}
	result /= float64(len(numbers))
	return
}