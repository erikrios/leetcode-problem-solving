package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Nums []int
	I    int
}

func Constructor(root *TreeNode) BSTIterator {
	nums := make([]int, 0)
	toList(root, &nums)
	return BSTIterator{Nums: nums, I: 0}
}

func toList(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	toList(root.Left, nums)
	*nums = append(*nums, root.Val)
	toList(root.Right, nums)
}

func (this *BSTIterator) Next() int {
	val := this.Nums[this.I]
	this.I += 1
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.I < len(this.Nums)
}
