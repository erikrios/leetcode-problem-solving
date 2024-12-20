package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(findTilt(node))

	node = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(findTilt(node))

	node = &TreeNode{
		Val: 21,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 14,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Println(findTilt(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	_, tilt := dfs(root)
	return tilt
}

func dfs(root *TreeNode) (sum int, tilt int) {
	if root == nil {
		return
	}

	leftSum, leftTilt := dfs(root.Left)
	rightSum, rightTilt := dfs(root.Right)

	sum = leftSum + rightSum + root.Val
	var absSum int
	if leftSum > rightSum {
		absSum = leftSum - rightSum
	} else {
		absSum = rightSum - leftSum
	}

	tilt = absSum + leftTilt + rightTilt

	return
}
