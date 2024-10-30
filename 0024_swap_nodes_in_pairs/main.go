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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	cur := head
	startList := head.Next
	if startList == nil {
		startList = head
	}

	for cur != nil {
		next := cur.Next
		if next == nil {
			if prev != nil {
				prev.Next = cur
			}
			break
		}
		cur.Next = next.Next
		next.Next = cur

		if prev != nil {
			prev.Next = next
		}

		prev = cur
		cur = cur.Next
	}

	return startList
}

func main() {
	input := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	swapped := swapPairs(&input)
	printList(swapped)
}
