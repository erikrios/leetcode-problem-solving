package main

import (
	"fmt"
	"sort"
)

func main() {
	node := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(kthLargestLevelSum(node, 2))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(kthLargestLevelSum(node, 1))
	fmt.Println(kthLargestLevelSum(node, 2))
	fmt.Println(kthLargestLevelSum(node, 3))
	fmt.Println(kthLargestLevelSum(node, 4))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	results := make([]int64, 0)
	dfs(root, 1, &results)

	if k-1 >= len(results) {
		return -1
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})

	return results[k-1]
}

func dfs(root *TreeNode, level int, results *[]int64) {
	if root == nil {
		return
	}

	val := int64(root.Val)
	if level > len(*results) {
		*results = append(*results, val)
	} else {
		(*results)[level-1] += val
	}

	dfs(root.Left, level+1, results)
	dfs(root.Right, level+1, results)
}
