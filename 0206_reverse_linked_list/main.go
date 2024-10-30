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

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	current := head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	input := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reversed := reverseList(&input)
	println("reversed")
	printList(reversed)
}
