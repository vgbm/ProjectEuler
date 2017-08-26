package main

import "testing"

func TestSumMultiplesBelow_10(t *testing.T) {
	val := SumMultiplesBelow(10)

	if val != 23 {
		t.Errorf("Expected 23, was %d", val)
	}
}

func TestIsMultipleOf_whereTrue(t *testing.T) {

	testCases := []struct {
		desc string
		n int
		multiples []int
	} {
		{desc: "4 is a mult of 2", n: 4, multiples: []int { 2 }},
		{desc: "7 is a mult of 7", n: 7, multiples: []int { 2, 7 }},
		{desc: "15 is a mult of 3 and 5", n: 15, multiples: []int { 3, 5 }},
	}

	for _, testCase := range testCases {
		isMultiple := IsMultipleOf(testCase.n, testCase.multiples)

		if isMultiple == false {
			t.Error(testCase.desc)
		}
	}
}

func TestIsMultipleOf_whereFalse(t *testing.T) {

	testCases := []struct {
		desc string
		n int
		multiples []int
	} {
		{desc: "8 not a mult of 3", n: 8, multiples: []int { 3 }},
		{desc: "3 not a mult of 6", n: 3, multiples: []int { 6 }},
		{desc: "15 not a mult of 4 or 6", n: 15, multiples: []int { 4, 6 }},
	}

	for _, testCase := range testCases {
		isMultiple := IsMultipleOf(testCase.n, testCase.multiples)

		if isMultiple {
			t.Error(testCase.desc)
		}
	}
}