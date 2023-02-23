package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil || root.Right == nil {
		var res bool
		if root.Val == 1 {
			res = true
		}
		return res
	}

	leftVal := evaluateTree(root.Left)
	rightVal := evaluateTree(root.Right)

	if root.Val == 2 {
		return leftVal || rightVal
	} else {
		return leftVal && rightVal
	}
}
