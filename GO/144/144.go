package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*1.递归*/
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var pt func(*TreeNode)
	pt = func(r *TreeNode) {
		if nil == r {
			return
		}
		ans = append(ans, r.Val)
		pt(r.Left)
		pt(r.Right)
	}
	pt(root)
	return ans
}
func main() {
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t1.Right = t2
	t2.Left = t3
	fmt.Println(preorderTraversal(t1))
}
