package main

import (
	"fmt"
	"strconv"
)

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(sumNumbers(node))

	node = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	fmt.Println(sumNumbers(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return calculateSum(root, make([]byte, 0))
}

func calculateSum(root *TreeNode, nums []byte) int {
	if isLeaf(root) {
		nums = append(nums, '0'+byte(root.Val))
		num, _ := strconv.Atoi(string(nums))
		nums = nums[:len(nums)-1]
		return num
	}

	nums = append(nums, '0'+byte(root.Val))

	var sum int

	if root.Left != nil {
		sum += calculateSum(root.Left, nums)
	}

	if root.Right != nil {
		sum += calculateSum(root.Right, nums)
	}

	nums = nums[:len(nums)-1]

	return sum
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
