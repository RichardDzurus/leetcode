package main

func printArr(arr []int) {
	for i := 0; i < len(arr); i += 1 {
		print(arr[i])
	}
}

func plusOne(digits []int) []int {
	size := len(digits)
	plusOne := true
	for i := size - 1; i >= 0; i -= 1 {
		if plusOne == false {
			continue
		}
		if digits[i] == 9 {
			digits[i] = 0
			plusOne = true
		} else {
			digits[i] += 1
			plusOne = false
		}
	}
	if plusOne == false {
		return digits
	}
	result := make([]int, size+1)
	result[0] = 1
	for i := 0; i < size; i += 1 {
		result[i+1] = digits[i]
	}
	return result
}

func main() {
	printArr(plusOne([]int{1, 2, 3}))
	println()
	printArr(plusOne([]int{9, 9, 9}))
	println()
	printArr(plusOne([]int{9, 9, 8}))
	println()
	printArr(plusOne([]int{9, 0, 9}))
	println()
	printArr(plusOne([]int{1, 9, 9}))
}
