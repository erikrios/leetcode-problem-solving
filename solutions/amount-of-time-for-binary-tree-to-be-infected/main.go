package main

import (
	"fmt"
)

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 10,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Println(amountOfTime(node, 3))

	node = &TreeNode{
		Val: 1,
	}

	fmt.Println(amountOfTime(node, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	mapper := make(map[int]map[int]struct{})
	convertTreeToGraph(root, 0, mapper)

	visited := make(map[int]struct{})
	queue := make([]int, 0)

	queue = append(queue, start)
	visited[start] = struct{}{}

	var minutes int

	for len(queue) > 0 {

		levelSize := len(queue)
		for levelSize > 0 {
			// Dequeue
			current := queue[0]
			queue = queue[1:]

			adjacents := mapper[current]

			for adjacent := range adjacents {
				if _, ok := visited[adjacent]; !ok {
					queue = append(queue, adjacent)
					visited[adjacent] = struct{}{}
				}
			}

			levelSize--
		}

		minutes++
	}

	return minutes - 1
}

func convertTreeToGraph(node *TreeNode, parent int, mapper map[int]map[int]struct{}) {
	if node != nil {
		current := node.Val
		if _, ok := mapper[current]; !ok {
			adjencyList := make(map[int]struct{})

			if parent != 0 {
				adjencyList[parent] = struct{}{}
			}

			if node.Left != nil {
				adjencyList[node.Left.Val] = struct{}{}
			}

			if node.Right != nil {
				adjencyList[node.Right.Val] = struct{}{}
			}

			mapper[current] = adjencyList

			convertTreeToGraph(node.Left, current, mapper)
			convertTreeToGraph(node.Right, current, mapper)
		}
	}
}
