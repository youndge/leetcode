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

func oddEvenList(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	evenHead := head.Next
	odd, even := head, evenHead
	for nil != even && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
func traverse(l *ListNode) {
	for nil != l {
		fmt.Printf("%v	", l.Val)
		l = l.Next
	}
	fmt.Println()
}
func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	traverse(n1)
	oddEvenList(n1)
	traverse(n1)
}
