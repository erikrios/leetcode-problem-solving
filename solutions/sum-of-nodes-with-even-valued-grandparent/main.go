package main

import "fmt"

func main() {
	node := &TreeNode{
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

	fmt.Println(sumEvenGrandparent(node))

	node = &TreeNode{
		Val: 1,
	}

	fmt.Println(sumEvenGrandparent(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := sumEvenGrandparent(root.Left)
	r := sumEvenGrandparent(root.Right)

	var sum int
	if root.Val%2 == 0 {
		sum = sumGrandchild(root)
	}

	return sum + l + r
}

func sumGrandchild(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return sumChild(root.Left) + sumChild(root.Right)
}

func sumChild(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var l, r int

	if root.Left != nil {
		l = root.Left.Val
	}

	if root.Right != nil {
		r = root.Right.Val
	}

	return l + r
}
