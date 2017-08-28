package main

import "testing"

func Test_LargestPalidrome(t *testing.T) {
	tt := []struct {
		base     int
		expected int
	}{
		{1, 999},
		{11, 10901},
	}

	for _, test := range tt {
		result := LargestPalidrome(test.base)

		if result != test.expected {
			t.Errorf("Expected %d, but was %d", test.expected, result)
		}
	}
}

func Test_IsPalindrome(t *testing.T) {
	tt := []struct {
		n        int
		expected bool
	}{
		{3, true},
		{80, false},
		{1001, true},
		{101, true},
		{102, false},
	}

	for _, test := range tt {
		result := IsPalindrome(test.n)

		if result != test.expected {
			t.Errorf("For %d, expected %t, but was %t", test.n, test.expected, result)
		}
	}
}

func Test_Max(t *testing.T) {
	tt := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 4}, 4},
		{[]int{3, 4, 2, 6, 1}, 6},
	}

	for _, test := range tt {
		result := Max(test.nums)

		if result != test.expected {
			t.Errorf("Expected %d, but was %d", test.expected, result)
		}
	}
}
