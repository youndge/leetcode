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

func isPalindrome(head *ListNode) bool {
	ln := []int{}
	for nil != head {
		ln = append(ln, head.Val)
		head = head.Next
	}
	h, t := 0, len(ln)-1
	for h <= t {
		if ln[h] != ln[t] {
			return false
		}
		h++
		t--
	}
	return true
}
func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l1.Next = l2
	fmt.Println(isPalindrome(l1))
	l3 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 1}
	l2.Next = l3
	l3.Next = l4
	fmt.Println(isPalindrome(l1))

}
