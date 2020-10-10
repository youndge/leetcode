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

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for nil != fast {
		if nil == fast.Next {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
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

	fmt.Println(detectCycle(ln1))

	return
}
