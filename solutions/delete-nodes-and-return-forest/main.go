package main

import (
	"fmt"
	"strings"
)

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(delNodes(node, []int{2, 3}))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(delNodes(node, []int{3, 5}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return "nil"
	}

	var result []string
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				result = append(result, "nil")
				continue
			}

			result = append(result, fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}

	// Trim trailing "nil" values for a cleaner visualization
	for len(result) > 0 && result[len(result)-1] == "nil" {
		result = result[:len(result)-1]
	}

	return strings.Join(result, ", ")
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	results := make([]*TreeNode, 0)
	mapper := make(map[int]struct{})

	for _, num := range toDelete {
		mapper[num] = struct{}{}
	}

	if !delete(root, mapper, &results) {
		results = append(results, root)
	}

	return results
}

func delete(root *TreeNode, toDelete map[int]struct{}, results *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	isLeftDeleted := delete(root.Left, toDelete, results)
	isRightDeleted := delete(root.Right, toDelete, results)

	if isLeftDeleted {
		root.Left = nil
	}

	if isRightDeleted {
		root.Right = nil
	}

	_, shouldDelete := toDelete[root.Val]

	if shouldDelete {
		if root.Left != nil {
			*results = append(*results, root.Left)
		}

		if root.Right != nil {
			*results = append(*results, root.Right)
		}
	}

	return shouldDelete
}
