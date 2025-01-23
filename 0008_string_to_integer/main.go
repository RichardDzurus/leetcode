package main

import "fmt"

const (
	Start       int = 0
	LeadingZero     = 1
	ValidNum        = 2
)

const MIN = -2_147_483_648
const MIN_STRING = "2147483648"
const MAX = 2_147_483_647
const MAX_STRING = "2147483647"

func shouldRound(s string, isPositive bool) bool {
	if len(s) > 10 {
		return true
	}
	if len(s) < 9 {
		return false
	}
	limitString := MAX_STRING
	if isPositive == false {
		limitString = MIN_STRING
	}

	for i := 0; i < len(s); i += 1 {
		if s[i] > limitString[i] {
			return true
		}
		if s[i] < limitString[i] {
			return false
		}
	}

	return false
}

func charToDigit(c rune) int {
	return int(c) - 48
}

func myAtoi(s string) int {
	isPositive := true
	result := 0
	validNumsString := ""
	state := Start

stringloop:
	for i := 0; i < len(s); i += 1 {
		char := s[i]
		switch state {
		case Start:
			if char == ' ' {
				continue
			}
			if char == '+' {
				isPositive = true
				state = LeadingZero
				continue
			}
			if char == '-' {
				isPositive = false
				state = LeadingZero
				continue
			}
			if char == '0' {
				state = LeadingZero
				continue
			}
			if char >= '1' && char <= '9' {
				isPositive = true
				state = ValidNum
				validNumsString += string(char)
				continue
			}
			break stringloop
		case LeadingZero:
			if char == '0' {
				continue
			}
			if char >= '1' && char <= '9' {
				state = ValidNum
				validNumsString += string(char)
				continue
			}
			break stringloop
		case ValidNum:
			if char >= '0' && char <= '9' {
				state = ValidNum
				validNumsString += string(char)
				continue
			}
			break stringloop
		}
	}

	if isPositive && shouldRound(validNumsString, isPositive) {
		return MAX
	}

	if isPositive == false && shouldRound(validNumsString, isPositive) {
		return MIN
	}

	level := 1
	for j := len(validNumsString) - 1; j >= 0; j -= 1 {
		result += charToDigit(rune(validNumsString[j])) * level
		level *= 10
	}
	if isPositive == false {
		result *= -1
	}
	return result
}

func main() {
	fmt.Println(myAtoi("-123"))
	fmt.Println(myAtoi("123"))
	fmt.Println(myAtoi("    -323a"))
	fmt.Println(myAtoi("   223aaaa"))
	fmt.Println(myAtoi("a   223aaaa"))
	fmt.Println(myAtoi("   -12a   223aaaa"))
	fmt.Println(myAtoi(" 00002022"))
	fmt.Println(myAtoi("  -00220"))
	fmt.Println(myAtoi("-2147483699"))
	fmt.Println(myAtoi("2147483646"))
	fmt.Println(myAtoi("0-1"))
}
