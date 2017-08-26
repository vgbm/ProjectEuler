package main

func main() {

	print(SumMultiplesBelow(1000))

}

func SumMultiplesBelow(n int) int {
	mults := []int{ 3, 5 }
	var sum int = 0

	for i := 1; i < n; i++ {
		if IsMultipleOf(i, mults) {
			sum += i
		}
	}

	return sum
}

func IsMultipleOf(n int, multiples []int) bool {
	for _, mult := range multiples {
		if n % mult == 0 {
			return true
		}
	}

	return false
}
