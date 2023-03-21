package main

import "fmt"

func main() {
	node := &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 5,
					},
					&Node{
						Val: 6,
					},
				},
			},
			&Node{
				Val: 2,
			},
			&Node{
				Val: 4,
			},
		},
	}

	fmt.Println(preorder(node))
}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	if len(root.Children) == 0 {
		return []int{root.Val}
	}

	results := make([]int, 0)

	results = append(results, root.Val)

	for i := 0; i < len(root.Children); i++ {
		child := root.Children[i]
		values := preorder(child)

		for j := 0; j < len(values); j++ {
			value := values[j]
			results = append(results, value)
		}
	}

	return results
}
