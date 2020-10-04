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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var listLength func(*ListNode) int
	listLength = func(l *ListNode) int {
		length := 0
		for nil != l{
			length += 1
			l = l.Next
		}
		return length
	}
	var merge func(*ListNode, *ListNode)
	merge = func(long *ListNode, short *ListNode) {
		if nil == long {
			return
		}
		if nil != short {
			long.Val += short.Val
			short = short.Next
		}
		if long.Val >= 10 {
			long.Val %= 10
			if nil != long.Next {
				long.Next.Val += 1
			} else {
				long.Next = &ListNode{Val: 1}
			}
		}
		long = long.Next
		merge(long, short)
	}
	fmt.Println(listLength(l1))
	fmt.Println("######")
	fmt.Println(listLength(l2))
	fmt.Println("######")

	if listLength(l1) > listLength(l2) {
		merge(l1, l2)
		return l1
	} else {
		merge(l2, l1)
		return l2
	}
	return nil
}
func main() {
	var l1, n1, l2 ListNode
	l1.Val = 1
	l1.Next = &n1
	n1.Val = 8
	// n1.Next = &n2
	// n2.Val = 4

	l2.Val = 0
	l2.Next = nil
	// l2.Next = &n3
	// n3.Val = 3
	// n3.Next = &n4
	// n4.Val = 4

	p := addTwoNumbers(&l1, &l2)
	for nil != p {
		fmt.Println(p.Val)
		p = p.Next
	}
	return
}
