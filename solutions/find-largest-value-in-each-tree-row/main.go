package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println(largestValues(node))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(largestValues(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)

	dfs(root, 0, &res)

	return res
}

func dfs(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}

	if level == len(*res) {
		*res = append(*res, root.Val)
	} else if (*res)[level] < root.Val {
		(*res)[level] = root.Val
	}

	nextLevel := level + 1
	dfs(root.Left, nextLevel, res)
	dfs(root.Right, nextLevel, res)
}
