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

func printArray(arr []int) {
	for i := 0; i < len(arr); i += 1 {
		print(arr[i])
		print(" ")
	}
	println()
}

func largestRectangleArea(heights []int) int {
	stack := make(stack, 0)
	maxArea := 0
	n := len(heights)

	for i := 0; i <= n; i++ {
		// We treat the height of the bar at the end as 0 to ensure all bars are processed
		currHeight := 0
		if i < n {
			currHeight = heights[i]
		}

		// Pop from the stack while the current bar is shorter than the stack top
		for !stack.IsEmpty() && currHeight < heights[stack.Top()] {
			height := heights[stack.Top()]
			stack.Pop()

			// Determine the width of the rectangle with the popped height
			width := i
			if !stack.IsEmpty() {
				width = i - stack.Top() - 1
			}

			// Update the maximum area
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		// Push current index to the stack
		stack.Push(i)
	}

	return maxArea
}

func main() {
	heights := []int{4, 1, 1, 1, 1, 1}
	printArray(heights)
	print(largestRectangleArea(heights))
}
