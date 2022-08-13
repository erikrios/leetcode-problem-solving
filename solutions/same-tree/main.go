package main

import "fmt"

func main() {
	p1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 10,
		},
		Right: &TreeNode{
			Val: 15,
		},
	}

	p2 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 15,
			},
		},
	}

	fmt.Println(isSameTree(p1, p2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	isSame := isSameTree(p.Left, q.Left)
	if !isSame {
		return false
	}

	isSame = isSameTree(p.Right, q.Right)
	if !isSame {
		return false
	}

	return true
}
