package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	printTree(balanceBST(node), "", true)

	node = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	printTree(balanceBST(node), "", true)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	if isLeft {
		fmt.Printf("%s└── %d\n", prefix, node.Val)
		prefix += "    "
	} else {
		fmt.Printf("%s┌── %d\n", prefix, node.Val)
		prefix += "│   "
	}

	if node.Left != nil || node.Right != nil {
		printTree(node.Right, prefix, false)
		printTree(node.Left, prefix, true)
	}
}

func balanceBST(root *TreeNode) *TreeNode {
	results := make([]int, 0)
	bstToSortedSlice(root, &results)
	return sortedSliceToBst(0, len(results)-1, results)
}

func bstToSortedSlice(root *TreeNode, results *[]int) {
	if root == nil {
		return
	}

	bstToSortedSlice(root.Left, results)
	*results = append(*results, root.Val)
	bstToSortedSlice(root.Right, results)
}

func sortedSliceToBst(start int, end int, results []int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &TreeNode{Val: results[mid]}

	root.Left = sortedSliceToBst(start, mid-1, results)
	root.Right = sortedSliceToBst(mid+1, end, results)

	return root
}
