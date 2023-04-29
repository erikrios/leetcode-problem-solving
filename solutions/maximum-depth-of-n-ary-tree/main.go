package main

import "fmt"

func main() {
	node := &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{Val: 5},
					&Node{Val: 6},
				},
			},
			&Node{Val: 2},
			&Node{Val: 4},
		},
	}

	fmt.Println(maxDepth(node))
}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	var max int
	for i := 0; i < len(root.Children); i++ {
		child := root.Children[i]
		v := maxDepth(child)
		if v > max {
			max = v
		}
	}

	return max + 1
}
