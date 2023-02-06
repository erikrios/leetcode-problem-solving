package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	fmt.Println(deepestLeavesSum(node))

	node = &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}

	fmt.Println(deepestLeavesSum(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	_, sum := sumLeaves(root)
	return sum
}

func sumLeaves(root *TreeNode) (depth, sum int) {
	if root.Left == nil && root.Right == nil {
		depth++
		sum = root.Val
		return
	}

	if root.Left != nil && root.Right != nil {
		dl, sl := sumLeaves(root.Left)
		dr, sr := sumLeaves(root.Right)
		if dl == dr {
			depth = dl + 1
			sum = sl + sr
			return
		} else if dl > dr {
			depth = dl + 1
			sum = sl
			return
		} else {
			depth = dr + 1
			sum = sr
			return
		}
	} else if root.Left != nil {
		dl, sl := sumLeaves(root.Left)
		depth = dl + 1
		sum = sl
		return
	} else {
		dr, sr := sumLeaves(root.Right)
		depth = dr + 1
		sum = sr
		return
	}
}
