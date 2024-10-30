package main

func divide2(dividend int, divisor int) int {
	result := 0
	counter := 0
	sign := 1
	positiveDividend := dividend
	if dividend < -2147483648 {
		positiveDividend = -2147483648
	}
	if dividend > 2147483647 {
		positiveDividend = 2147483647
	}
	if dividend < 0 {
		positiveDividend = -dividend
		sign = -sign
	}
	positiveDivisor := divisor
	if divisor < -2147483648 {
		positiveDivisor = -2147483648
	}
	if divisor > 2147483647 {
		positiveDivisor = 2147483647
	}
	if divisor < 0 {
		positiveDivisor = -divisor
		sign = -sign
	}
	for counter <= positiveDividend {
		counter += positiveDivisor
		result += 1
	}
	if counter > positiveDividend {
		result -= 1
	}
	if sign == -1 {
		result = -result
	}
	restrictedResult := result
	if result < -2147483648 {
		restrictedResult = -2147483648
	}
	if result > 2147483647 {
		restrictedResult = 2147483647
	}
	return restrictedResult
}

// func divide(dividend int, divisor int) {
// 	sign := (dividend < 0) ^ (divisor < 0)
// }

func add(a int, b int) int {
	for b > 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func multiply(a int, b int) int {
	result := 0
	positiveA := a
	isNegativeResult := false
	if a < 0 {
		positiveA = ^a + 1
		isNegativeResult = !isNegativeResult
	}
	positiveB := b
	if b < 0 {
		positiveB = ^b + 1
		isNegativeResult = !isNegativeResult
	}
	for positiveB > 0 {
		if positiveB&1 == 1 {
			result = add(result, positiveA)
		}
		positiveA = positiveA << 1
		positiveB = positiveB >> 1
	}
	if isNegativeResult {
		return ^result + 1
	}
	return result
}

func subtract(a int, b int) int {
	for b > 0 {
		borrow := ^a & b
		a = a ^ b
		b = borrow << 1
	}
	return a
}

const INT_MAX = 1<<31 - 1
const INT_MIN = -1 << 31

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}

	isNegativeResult := false
	positiveDividend := dividend
	if dividend < 0 {
		positiveDividend = ^dividend + 1
		isNegativeResult = !isNegativeResult
	}
	positiveDivisor := divisor
	if divisor < 0 {
		positiveDivisor = ^divisor + 1
		isNegativeResult = !isNegativeResult
	}

	quotient := 0

	bits := 32

	for i := bits - 1; i >= 0; i -= 1 {
		if (positiveDividend >> i) >= positiveDivisor {
			positiveDividend = subtract(positiveDividend, positiveDivisor<<i)

			quotient |= 1 << i
		}
	}

	if isNegativeResult {
		return ^quotient + 1
	}

	return quotient
}

func main() {
	// println(divide(10, 3))
	// println(divide(7, -3))
	// println(divide(8, 2))
	// println(divide(-8, -2))
	// println(divide(1, 3))
	// println(multiply(-8, 7))
	// println(multiply(8, 7))
	println(divide(21, -7))
	println(divide(5, 5))
	println(divide(21, 7))
	println(divide(19, 18))
	println(divide(-2147483648, -1))
}
