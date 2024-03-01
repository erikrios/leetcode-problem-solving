package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	return FindElements{root: root}
}

func (this *FindElements) Find(target int) bool {
	return find(this.root, 0, target)
}

func find(root *TreeNode, val int, target int) bool {
	if root == nil {
		return false
	}

	if val == target {
		return true
	}

	return find(root.Left, 2*val+1, target) || find(root.Right, 2*val+2, target)
}
