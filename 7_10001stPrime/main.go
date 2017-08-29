package main

import "math"

func main() {

	const limit int = 10001
	counter := 2
	prime := 4

	for i := prime; counter < limit; i++ {
		if isPrime(i) {
			counter++
			prime = i
		}
	}

	print(prime)
}

func isPrime(n int) bool {
	//only do up to sqrt(n) to run in sqrt(n) time
	limit := int(math.Ceil(math.Sqrt(float64(n))))

	//fmt.Printf("%d | %d \n", n, limit)

	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
