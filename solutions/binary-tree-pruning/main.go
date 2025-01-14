package main

func main() {
	node := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	pruneTree(node)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if !prune(root) {
		return nil
	}

	return root
}

func prune(root *TreeNode) bool {
	if root == nil {
		return false
	}

	res := root.Val == 1

	leftRes := prune(root.Left)
	rightRes := prune(root.Right)

	if !leftRes {
		root.Left = nil
	}

	if !rightRes {
		root.Right = nil
	}

	return res || leftRes || rightRes
}
