package main

import "testing"

func Test_SumOfSq(t *testing.T) {
	tt := []struct {
		limit    int
		expected int
	}{
		{3, 14},
		{4, 30},
		{10, 385},
	}

	for _, test := range tt {
		result := SumOfSq(test.limit)

		if result != test.expected {
			t.Errorf("Expected %d, but was %d", test.expected, result)
		}
	}
}

func Test_SqOfSum(t *testing.T) {
	tt := []struct {
		limit    int
		expected int
	}{
		{3, 36},
		{4, 100},
		{10, 3025},
	}

	for _, test := range tt {
		result := SqOfSum(test.limit)

		if result != test.expected {
			t.Errorf("Expected %d, but was %d", test.expected, result)
		}
	}
}
