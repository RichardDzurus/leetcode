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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	length1 := len(nums1)
	length2 := len(nums2)
	stackSecond := make(stack, 0)
	resultNums2 := make([]int, length2)
	mapValue := make(map[int]int)
	final := make([]int, length1)

	for i := 0; i < length2; i += 1 {
		resultNums2[i] = -1
		if i < length1 {
			final[i] = -1
		}
	}

	for i := 0; i < length2; i += 1 {
		for stackSecond.IsEmpty() == false && (nums2[i] > nums2[stackSecond.Top()]) {
			index, _ := stackSecond.Pop()
			resultNums2[index] = nums2[i]
		}
		mapValue[nums2[i]] = i
		stackSecond.Push(i)
	}

	for i := 0; i < length1; i += 1 {
		indexInResults := mapValue[nums1[i]]
		final[i] = resultNums2[indexInResults]
	}

	return final
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i += 1 {
		print(arr[i])
	}
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	result := nextGreaterElement(nums1, nums2)
	printArray(result)
}
