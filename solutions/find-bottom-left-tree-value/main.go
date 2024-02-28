package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(findBottomLeftValue(node))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(findBottomLeftValue(node))

	node = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
		},
	}
	fmt.Println(findBottomLeftValue(node))

	node = &TreeNode{
		Val: 0,
		Right: &TreeNode{
			Val: -1,
		},
	}
	fmt.Println(findBottomLeftValue(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	vals := make([][]int, 0)
	find(root, -1, &vals)
	return vals[len(vals)-1][0]
}

func find(root *TreeNode, level int, vals *[][]int) {
	level++

	if root == nil {
		return
	}

	if level >= len(*vals) {
		for i := len(*vals); i <= level; i++ {
			*vals = append(*vals, []int{})
		}
	}
	(*vals)[level] = append((*vals)[level], root.Val)

	find(root.Left, level, vals)
	find(root.Right, level, vals)
}
