package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(binaryTreePaths(node))

	node = &TreeNode{Val: 1}
	fmt.Println(binaryTreePaths(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	return dfs(root, "")
}

func dfs(root *TreeNode, res string) []string {
	val := root.Val

	if isLeaf(root) {
		str := fmt.Sprintf("%d", val)
		res += str
		return []string{res}
	}

	str := fmt.Sprintf("%d->", val)
	res += str

	results := make([]string, 0)

	if root.Left != nil {
		results = append(results, dfs(root.Left, res)...)
	}

	if root.Right != nil {
		results = append(results, dfs(root.Right, res)...)
	}

	return results
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
