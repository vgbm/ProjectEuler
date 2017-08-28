package main

func main() {
	print(SumOfEvenFibsUnder(4000000))
}

func SumOfEvenFibsUnder(upperLimit int) int {
	prev := 0
	curr := 1
	sum := 0

	for curr <= upperLimit {
		if curr%2 == 0 {
			sum += curr
		}

		//next iter of fib nums
		prev, curr = curr, curr+prev
	}

	return sum
}
