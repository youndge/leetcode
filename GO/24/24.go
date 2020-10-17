package main

import (
	"fmt"
)

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

func swapPairs(head *ListNode) *ListNode {
	headnode := &ListNode{0, head}
	temp := headnode
	for nil != temp.Next && nil != temp.Next.Next {
		n1 := temp.Next
		n2 := temp.Next.Next
		temp.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		temp = n1
	}
	return headnode.Next
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	for nil != n1 {
		fmt.Println(n1.Val)
		n1 = n1.Next
	}
	nn := swapPairs(n1)
	for nil != nn {
		fmt.Println(nn.Val)
		nn = nn.Next
	}
	return
}
