package main

import (
	"fmt"
)

func main() {
	fmt.Println(len(allPossibleFBT(19)))
	fmt.Println(len(allPossibleFBT(9)))
	fmt.Println(len(allPossibleFBT(7)))
	fmt.Println(len(allPossibleFBT(5)))
	fmt.Println(len(allPossibleFBT(3)))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	return possibleFBT(n, map[int][]*TreeNode{})
}

func possibleFBT(n int, memo map[int][]*TreeNode) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}

	if n == 1 {
		return []*TreeNode{{}}
	}

	if vals, ok := memo[n]; ok {
		return vals
	}

	n -= 1

	left := n
	right := 0

	results := make([]*TreeNode, 0)
	for left != 0 {
		leftResults := possibleFBT(left, memo)
		rightResults := possibleFBT(right, memo)

		left--
		right++

		if len(leftResults) == 0 || len(rightResults) == 0 {
			continue
		}

		for _, leftRes := range leftResults {
			for _, rightRes := range rightResults {
				root := &TreeNode{}
				root.Left = leftRes
				root.Right = rightRes
				results = append(results, root)
			}
		}
	}

	memo[n] = results

	return results
}
