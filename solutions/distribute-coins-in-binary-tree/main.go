package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	fmt.Println(distributeCoins(node))

	node = &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	fmt.Println(distributeCoins(node))

	node = &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	fmt.Println(distributeCoins(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	totalMoves, _ := distribute(root)
	return totalMoves
}

func distribute(root *TreeNode) (totalMoves, amount int) {
	if root == nil {
		totalMoves = 0
		amount = 0
		return
	}

	leftTotalMoves, leftAmount := distribute(root.Left)
	rightTotalMoves, rightAmount := distribute(root.Right)

	root.Val += (leftAmount + rightAmount)

	val := root.Val
	amount = val - 1
	totalMoves = amount
	if totalMoves < 0 {
		totalMoves *= -1
	}

	totalMoves += (leftTotalMoves + rightTotalMoves)
	return
}
