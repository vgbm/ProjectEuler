package main

import "testing"

func Test_SumOfEvenFibsUnder(t *testing.T) {
	tt := []struct {
		limit    int
		expected int
	}{
		{3, 2},
		{8, 10},
		{33, 10},
		{34, 44},
	}

	for _, test := range tt {
		result := SumOfEvenFibsUnder(test.limit)

		if result != test.expected {
			t.Errorf("Expected %d, but was %d", test.expected, result)
		}
	}

}
