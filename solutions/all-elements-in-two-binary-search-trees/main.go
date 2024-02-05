package main

import "fmt"

func main() {
	node1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	node2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(getAllElements(node1, node2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1, root2 *TreeNode) []int {
	res1, res2 := getElements(root1), getElements(root2)
	m, n := len(res1), len(res2)
	results := make([]int, 0, m+n)

	var i, j int

	for i < m || j < n {
		var min int

		if i == m {
			min = res2[j]
			j++
		} else if j == n {
			min = res1[i]
			i++
		} else {
			val1 := res1[i]
			val2 := res2[j]
			if val1 < val2 {
				min = val1
				i++
			} else {
				min = val2
				j++
			}
		}

		results = append(results, min)
	}

	return results
}

func getElements(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := getElements(root.Left)
	res = append(res, root.Val)
	res = append(res, getElements(root.Right)...)

	return res
}
