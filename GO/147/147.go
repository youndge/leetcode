package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
func traverse(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ->", head.Val)
		head = head.Next
	}
	fmt.Println()
}
func main() {
	n1 := &ListNode{Val: 4}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n1.Next, n2.Next, n3.Next = n2, n3, n4
	traverse(n1)
	traverse(insertionSortList(n1))

}
