package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(findTarget(node, 9))
	fmt.Println(findTarget(node, 28))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return findTargetRec(root, k, map[int]struct{}{})
}

func findTargetRec(root *TreeNode, k int, mapper map[int]struct{}) bool {
	if root == nil {
		return false
	}

	val := root.Val

	if _, ok := mapper[val]; ok {
		return true
	}

	mapper[k-val] = struct{}{}

	return findTargetRec(root.Left, k, mapper) || findTargetRec(root.Right, k, mapper)
}
