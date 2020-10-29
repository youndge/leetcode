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

//深度优先搜索
func sumNumbers(root *TreeNode) int {
	var dfs func(r *TreeNode, prevSum int) int
	dfs = func(r *TreeNode, prevSum int) int {
		if nil == r {
			return 0
		}
		sum := prevSum*10 + r.Val
		if nil == r.Left && nil == r.Right {
			return sum
		}
		return dfs(r.Left, sum) + dfs(r.Right, sum)
	}
	return dfs(root, 0)
}

//广度优先搜索 + 队列
type pair struct {
	node *TreeNode
	num  int
}

func sumNumbers2(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	queue := []pair{{root, root.Val}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		left, right, num := p.node.Left, p.node.Right, p.num
		if left == nil && right == nil {
			sum += num
		} else {
			if left != nil {
				queue = append(queue, pair{left, num*10 + left.Val})
			}
			if right != nil {
				queue = append(queue, pair{right, num*10 + right.Val})
			}
		}
	}
	return
}

func traverse(r *TreeNode) {
	if nil == r {
		return
	}
	fmt.Printf("%d  ", r.Val)
	traverse(r.Left)
	traverse(r.Right)

}
func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Left = n2
	n1.Right = n3
	traverse(n1)
	fmt.Println()
	fmt.Println(sumNumbers(n1))
	fmt.Println(sumNumbers2(n1))

	n1 = &TreeNode{Val: 4}
	n2 = &TreeNode{Val: 9}
	n3 = &TreeNode{Val: 0}
	n4 := &TreeNode{Val: 5}
	n5 := &TreeNode{Val: 1}

	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	traverse(n1)
	fmt.Println()

	fmt.Println(sumNumbers(n1))
	fmt.Println(sumNumbers2(n1))
}
