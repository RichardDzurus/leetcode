package main

func reverse(x int) int {
	positiveValue := x
	if x < 0 {
		positiveValue = -1 * x
	}

	level := 1
	positiveValueCp := positiveValue
	for positiveValueCp > 0 {
		positiveValueCp /= 10
		level *= 10
	}
	level /= 10

	reversedNum := 0
	for positiveValue > 0 {
		reversedNum += level * (positiveValue % 10)
		level /= 10
		positiveValue /= 10
	}

	if reversedNum > 2147483647 {
		return 0
	}

	if reversedNum < -2147483648 {
		return 0
	}

	if x < 0 {
		return -1 * reversedNum
	}
	return reversedNum
}

func main() {
	print(reverse(2147483647))
}
