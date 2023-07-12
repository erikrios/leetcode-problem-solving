package main

import "fmt"

func main() {
	node1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	node2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 2},
	}

	fmt.Println(leafSimilar(node1, node2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1, root2 *TreeNode) bool {
	root1Leafs := leafs(root1)
	root2Leafs := leafs(root2)

	m := len(root1Leafs)
	n := len(root2Leafs)

	if m != n {
		return false
	}

	for i := 0; i < m; i++ {
		if root1Leafs[i] != root2Leafs[i] {
			return false
		}
	}

	return true
}

func leafs(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	res := make([]int, 0)

	if node.Left != nil {
		vals := leafs(node.Left)
		res = append(res, vals...)
	}

	if node.Right != nil {
		vals := leafs(node.Right)
		res = append(res, vals...)
	}

	return res
}
