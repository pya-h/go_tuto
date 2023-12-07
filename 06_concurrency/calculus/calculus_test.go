package calculus

import "testing"

func TestAverage(t *testing.T) {

	type testPair struct {
		values []float64
		average float64
	}

	tests := [] testPair {
		{[] float64 {10, 20}, 15},
		{[] float64 {1, 2, 3}, 2},
		{[] float64 {10, 15, 20, 30, 5}, 16},

	}

	for _, pair := range tests {
		ave := Average(pair.values)
		if ave != pair.average {
			t.Error(
				"Test on: ", pair.values,
				"Resulted in: ", ave,
				"But", pair.average,
				"Was expected!",
			)
		}
	}
}

