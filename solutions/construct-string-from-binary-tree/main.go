package main

import (
	"fmt"
	"strings"
)

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(tree2str(node))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(tree2str(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	val := preorder(root)
	return string(val[1 : len(val)-1])
}

func preorder(root *TreeNode) string {
	if root == nil {
		return "()"
	}

	var builder strings.Builder
	str := fmt.Sprintf("%s%d", "(", root.Val)
	builder.WriteString(str)

	leftStr := preorder(root.Left)

	rightStr := preorder(root.Right)

	if leftStr != "()" && rightStr == "()" {
		builder.WriteString(leftStr)
	} else if leftStr == "()" && rightStr != "()" {
		builder.WriteByte('(')
		builder.WriteByte(')')
		builder.WriteString(rightStr)
	} else if leftStr != "()" && rightStr != "()" {
		builder.WriteString(leftStr)
		builder.WriteString(rightStr)
	}

	builder.WriteByte(')')

	return builder.String()
}
