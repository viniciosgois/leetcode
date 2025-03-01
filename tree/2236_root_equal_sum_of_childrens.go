package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	sum := root.Left.Val + root.Right.Val
	if root.Val == sum {
		return true
	} else {
		return false
	}
}
