package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if nil == root {
		return 0
	}
	trees := []int{}
	var midorder func(*TreeNode)
	midorder = func(r *TreeNode) {
		if nil == r {
			return
		}
		midorder(r.Left)
		trees = append(trees, r.Val)
		midorder(r.Right)
	}
	midorder(root)
	if len(trees) == 1 {
		return trees[0]
	} else if len(trees) == 2 {
		return trees[1] - trees[0]
	} else if len(trees) > 2 {
		min := trees[1] - trees[0]
		for i := 1; i < len(trees)-1; i++ {
			if trees[i+1]-trees[i] < min {
				min = trees[i+1] - trees[i]
			}
		}
		return min
	}
	return 0
}

func main() {
	ln1 := &TreeNode{Val: 1}
	ln2 := &TreeNode{Val: 2}
	ln3 := &TreeNode{Val: 3}

	ln1.Right = ln3
	ln3.Left = ln2

	fmt.Println(getMinimumDifference(ln1))
	ln4 := &TreeNode{Val: 543}
	ln5 := &TreeNode{Val: 384}
	ln6 := &TreeNode{Val: 652}
	ln7 := &TreeNode{Val: 445}
	ln8 := &TreeNode{Val: 699}
	ln4.Left = ln5
	ln4.Right = ln6
	ln5.Right = ln7
	ln6.Right = ln8
	fmt.Println(getMinimumDifference(ln4))

	return
}
