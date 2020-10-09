package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if nil == head || nil == head.Next {
		return false
	}
	p1, p2 := head, head.Next
	for p1 != p2 {
		if nil == p2 || nil == p2.Next {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return true
}
func main() {
	ln1 := &ListNode{Val: 1}
	ln2 := &ListNode{Val: 2}
	ln3 := &ListNode{Val: 3}
	ln4 := &ListNode{Val: 4}
	ln5 := &ListNode{Val: 5}
	ln1.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5
	ln5.Next = ln3

	fmt.Println(hasCycle(ln1))
	return
}
