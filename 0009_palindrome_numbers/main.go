package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var digit int
	digits := [15]int{}
	i := 0
	for x > 0 {
		digit = x % 10
		digits[i] = digit
		x = x / 10
		i += 1
	}
	for j := 0; j < i/2; j++ {
		if digits[j] != digits[i-j-1] {
			return false
		}
	}
	return true
}

func main() {
	println(isPalindrome(11))
}
