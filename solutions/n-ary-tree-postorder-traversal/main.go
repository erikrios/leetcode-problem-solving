package main

import "fmt"

func main() {
	node := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(postorder(node))
}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	if len(root.Children) == 0 {
		return []int{root.Val}
	}

	results := make([]int, 0)
	for i := 0; i < len(root.Children); i++ {
		child := root.Children[i]
		values := postorder(child)
		results = append(results, values...)
	}
	results = append(results, root.Val)

	return results
}
