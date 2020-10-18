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
type TreeNode struct{
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	}
func isLeafNode(node *TreeNode) bool{
    return node.Left == nil && node.Right == nil
}
func dft(node *TreeNode) (sum int){
    if nil != node.Left{
        if isLeafNode(node.Left) {
            sum+=node.Left.Val
        }else {
        sum += dft(node.Left)
        }
    }
    if !isLeafNode(node.Right) && nil != node.Right{
        sum += dft(node.Right)
    }
    return 
}
func sumOfLeftLeaves(root *TreeNode) int {
    if nil!=root{
        return dft(root)
    }
    return 0
}
func main()  {
	n1 := TreeNode{
		15,nil,nil,
	}
	n2 := TreeNode{
		7,nil,nil,
	}
	n3 := TreeNode{
		20,&n1,&n2,
	}
	n4 := TreeNode{
		9,nil,nil,
	}
	n5 := TreeNode{
		3,&n4,&n3,
	}
	fmt.Println(sumOfLeftLeaves(&n5))
}