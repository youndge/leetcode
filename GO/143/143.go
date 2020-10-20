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

/*1.转为线性表再根据下标重构链表*/
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

//自己写的未通过
// func reorderList(head *ListNode) {
// 	if nil == head || nil == head.Next || nil == head.Next.Next {
// 		return
// 	}
// 	ln := []*ListNode{}
// 	for nil != head {
// 		ln = append(ln, head)
// 		head = head.Next
// 	}
// 	tail := len(ln) - 1
// 	pt, ph := ln[tail], head
// 	for (len(ln)%2 == 0 && tail == len(ln)/2+1) || (len(ln)%2 != 0 && tail == len(ln)/2) {
// 		pt.Next = ph.Next
// 		ph.Next = pt
// 		ph = pt.Next
// 		tail--
// 		pt = ln[tail]
// 	}
// 	return
// }

/*2.寻找链表中点+逆序+合并*/
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList2(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
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

	reorderList(ln1)
	for nil != ln1 {
		fmt.Println(ln1.Val)
		ln1 = ln1.Next
	}
	reorderList2(ln1)
	for nil != ln1 {
		fmt.Println(ln1.Val)
		ln1 = ln1.Next
	}
	return
}
