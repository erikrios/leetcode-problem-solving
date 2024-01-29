package main

import "fmt"

func main() {
	fmt.Println(bstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
	fmt.Println(bstFromPreorder([]int{1, 3}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}

	for i := 1; i < len(preorder); i++ {
		insert(root, preorder[i])
	}

	return root
}

func insert(root *TreeNode, val int) {
	currentVal := root.Val

	if val < currentVal {
		if root.Left == nil {
			newNode := &TreeNode{Val: val}
			root.Left = newNode
			return
		}
		insert(root.Left, val)
	} else {
		if root.Right == nil {
			newNode := &TreeNode{Val: val}
			root.Right = newNode
			return
		}
		insert(root.Right, val)
	}
}
