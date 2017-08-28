package main

func main() {
	print(LargestPrimeFactor(600851475143))
}

func LargestPrimeFactor(num int) int {
	maxPrime := 1

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			num /= i
			maxPrime = i
		}
	}

	return maxPrime
}
