package main

import "fmt"

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%v", traverse(t))
}

func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if isLeaf(root) {
		return []int{root.Val}
	}

	nums := make([]int, 0)

	if root.Left != nil {
		childs := traverse(root.Left)
		nums = append(nums, childs...)
	}

	if root.Right != nil {
		childs := traverse(root.Right)
		nums = append(nums, childs...)
	}

	nums = append(nums, root.Val)

	return nums
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	maxIdxPos := maxIndexPos(nums)
	leftNums := nums[:maxIdxPos]
	rightNums := nums[maxIdxPos+1:]

	node := &TreeNode{Val: nums[maxIdxPos]}
	node.Left = constructMaximumBinaryTree(leftNums)
	node.Right = constructMaximumBinaryTree(rightNums)

	return node
}

func maxIndexPos(nums []int) int {
	var maxIndex int

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}
