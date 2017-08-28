package main

func main() {

	var palindromes []int
	for i := 1000; i >= 100; i-- {
		palindromes = append(palindromes, LargestPalidrome(i))
	}
	print(Max(palindromes))
}

func LargestPalidrome(base int) int {

	for i := 1000; i >= 100; i-- {
		if IsPalindrome(base * i) {
			return base * i
		}
	}
	return -1
}

func IsPalindrome(n int) bool {
	num := n
	rev := 0

	for num > 0 {
		rev = rev*10 + num%10
		num /= 10
	}

	return n == rev
}

func Max(palindromes []int) int {
	max := palindromes[0]

	for _, num := range palindromes {
		if num > max {
			max = num
		}
	}

	return max
}
