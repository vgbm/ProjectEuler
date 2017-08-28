package main

const LIMIT int = 20

func main() {

	i := LIMIT

	for {
		if isDivByAllUnderLimit(i) {
			break
		}

		i++
	}

	print(i)
}

func isDivByAllUnderLimit(n int) bool {

	for i := 2; i <= LIMIT; i++ {
		if n%i != 0 {
			return false
		}
	}

	return true
}
