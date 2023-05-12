package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
	}

	fmt.Println(increasingBST(node))

	node = &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 7},
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	vals := make([]int, 0)
	inOrder(root, &vals)

	node := &TreeNode{}
	cur := node

	for i := 0; i < len(vals); i++ {
		cur.Right = &TreeNode{Val: vals[i]}
		cur = cur.Right
	}

	return node.Right
}

func inOrder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inOrder(root.Right, vals)
}
