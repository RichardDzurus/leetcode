package main

type stack []int

// Top doesn't modify the stack, so no need for a pointer receiver.
func (s stack) Top() int {
	return s[len(s)-1]
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

// Push needs a pointer receiver to modify the original stack.
func (s *stack) Push(v int) {
	*s = append(*s, v)
}

// Pop also needs a pointer receiver to modify the original stack.
func (s *stack) Pop() (int, bool) {
	l := len(*s)
	if l == 0 {
		return -1, false // Handle the empty stack case.
	}
	value := (*s)[l-1]
	*s = (*s)[:l-1]
	return value, true
}

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	indexStack := make(stack, 0)
	result := make([]int, length)

	for i := 0; i < length; i += 1 {
		result[i] = 0

		for indexStack.IsEmpty() == false && temperatures[i] > temperatures[indexStack.Top()] {
			index, _ := indexStack.Pop()
			result[index] = i - index
		}

		indexStack.Push(i)
	}

	return result
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i += 1 {
		print(arr[i])
		print(" ")
	}
	println()
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	printArray(temperatures)
	printArray(dailyTemperatures(temperatures))
}
