package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	fmt.Println(kthSmallest(node, 1))

	node = &TreeNode{
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
		},
	}
	fmt.Println(kthSmallest(node, 3))
	fmt.Println(kthSmallest(node, 4))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var counter int
	return dfs(root, k, &counter)
}

func dfs(root *TreeNode, k int, counter *int) int {
	if root == nil {
		return -1
	}

	res := dfs(root.Left, k, counter)
	*counter += 1
	if *counter == k {
		return root.Val
	}

	if res == -1 {
		return dfs(root.Right, k, counter)
	}

	return res
}
