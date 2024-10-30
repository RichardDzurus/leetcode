package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		print(current.Val)
		current = current.Next
	}
}

func reverseBetweenValues(head *ListNode, left int, right int) *ListNode {
	var startList, curr, prev, lastBeforeStart, lastBeforeEnd, startOfReverse, endOfReverse *ListNode
	curr = head
	startList = head

	for curr != nil {
		next := curr.Next
		if curr.Val < left {
			lastBeforeStart = curr
			curr = next
			continue
		}
		if curr.Val > right {
			if endOfReverse == nil {
				endOfReverse = curr
			}
			curr = next
			continue
		}
		if curr.Val >= left && startOfReverse == nil {
			startOfReverse = curr
		}
		lastBeforeEnd = curr
		curr.Next = prev
		prev = curr
		curr = next
	}

	if lastBeforeStart != nil && lastBeforeEnd != nil {
		lastBeforeStart.Next = lastBeforeEnd

	}
	if startOfReverse != nil && endOfReverse != nil {
		startOfReverse.Next = endOfReverse
	}

	if startList.Val >= left && lastBeforeEnd != nil {
		startList = lastBeforeEnd
	}

	return startList
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var startList, curr, prev, lastBeforeStart, lastBeforeEnd, startOfReverse, endOfReverse *ListNode
	curr = head
	startList = head
	index := 1

	for curr != nil {
		next := curr.Next
		if index < left {
			lastBeforeStart = curr
			curr = next
			index += 1
			continue
		}
		if index > right {
			if endOfReverse == nil {
				endOfReverse = curr
			}
			curr = next
			index += 1
			continue
		}
		if index >= left && startOfReverse == nil {
			startOfReverse = curr
		}
		lastBeforeEnd = curr
		curr.Next = prev
		prev = curr
		curr = next
		index += 1
	}

	if lastBeforeStart != nil && lastBeforeEnd != nil {
		lastBeforeStart.Next = lastBeforeEnd

	}
	if startOfReverse != nil && endOfReverse != nil {
		startOfReverse.Next = endOfReverse
	}

	if left == 1 && lastBeforeEnd != nil {
		startList = lastBeforeEnd
	}

	return startList
}

func main() {
	// input := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	input := ListNode{3, &ListNode{5, nil}}
	left := 2
	right := 2
	reversed := reverseBetween(&input, left, right)
	printList(reversed)
}
