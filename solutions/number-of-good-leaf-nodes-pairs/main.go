package main

import "fmt"

func main() {
	node := &TreeNode{
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
	fmt.Println(countPairs(node, 3))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(countPairs(node, 3))

	node = &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
	}
	fmt.Println(countPairs(node, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	_, good := count(root, distance)
	return good
}

func count(root *TreeNode, distance int) ([]int, int) {
	if root == nil {
		return []int{}, 0
	}

	if isLeaf(root) {
		return []int{1}, 0
	}

	leftCount, leftGood := count(root.Left, distance)
	rightCount, rightGood := count(root.Right, distance)

	var good int
	for _, lc := range leftCount {
		for _, rc := range rightCount {
			if lc+rc <= distance {
				good++
			}
		}
	}

	leftCount = append(leftCount, rightCount...)
	for i := 0; i < len(leftCount); i++ {
		leftCount[i]++
	}

	return leftCount, good + leftGood + rightGood
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
