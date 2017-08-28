package main

func main() {
	const limit int = 100
	print(SqOfSum(limit) - SumOfSq(limit))
}

func SumOfSq(limit int) int {
	res := 0

	for i := 1; i <= limit; i++ {
		res += i * i
	}

	return res
}

func SqOfSum(limit int) int {
	sum := limit * (limit + 1) / 2

	return sum * sum
}
