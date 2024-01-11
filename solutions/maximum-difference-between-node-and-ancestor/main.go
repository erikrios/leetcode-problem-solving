package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: 14,
				Left: &TreeNode{
					Val: 13,
				},
			},
		},
	}

	fmt.Println(maxAncestorDiff(node))

	node = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	fmt.Println(maxAncestorDiff(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

// First ways
func maxAncestors(root *TreeNode, ancestors []int) int {
	if root == nil {
		return 0
	}

	var max int

	recentVal := root.Val
	for _, ancestor := range ancestors {
		diff := ancestor - recentVal
		if diff < 0 {
			diff *= -1
		}

		if diff > max {
			max = diff
		}
	}

	ancestors = append(ancestors, recentVal)

	// Copy the slice to prevent using the same address
	leftAncestors, rightAncestors := make([]int, len(ancestors)), make([]int, len(ancestors))
	copy(leftAncestors, ancestors)
	copy(rightAncestors, ancestors)

	leftMax := maxAncestors(root.Left, ancestors)
	rightMax := maxAncestors(root.Right, ancestors)

	childMax := leftMax
	if rightMax > childMax {
		childMax = rightMax
	}

	if childMax > max {
		return childMax
	}

	return max
}

// Second ways
func dfs(node *TreeNode, min, max int) int {
	if node == nil {
		return max - min
	}

	val := node.Val

	if val < min {
		min = val
	}

	if val > max {
		max = val
	}

	left := dfs(node.Left, min, max)
	right := dfs(node.Right, min, max)

	if left > right {
		return left
	}

	return right
}
