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

/*自己的方法：*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dist := 0
	var p1, p2 *ListNode
	dummy := &ListNode{0, head}
	for p1, p2 = dummy, dummy; nil != p1.Next; p1 = p1.Next {
		if dist == n {
			p2 = p2.Next
		} else {
			dist++
		}
	}
	if 1 == n && 1 == dist {
		p2.Next = nil
	} else if dist < n {
		return nil
	} else {
		p2.Next = p2.Next.Next
	}
	return dummy.Next
}

/*别人的方法：*/
// func getLength(head *ListNode) (length int) {
//     for ; head != nil; head = head.Next {
//         length++
//     }
//     return
// }

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     length := getLength(head)
//     dummy := &ListNode{0, head}
//     cur := dummy
//     for i := 0; i < length-n; i++ {
//         cur = cur.Next
//     }
//     cur.Next = cur.Next.Next
//     return dummy.Next
// }

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
	n1 = removeNthFromEnd(n1, 2)
	for nil != n1 {
		fmt.Println(n1.Val)
		n1 = n1.Next
	}
	return
}
