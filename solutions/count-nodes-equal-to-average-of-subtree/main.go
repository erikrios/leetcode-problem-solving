package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Println(averageOfSubtree(node))

	node = &TreeNode{Val: 1}

	fmt.Println(averageOfSubtree(node))

	node = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	fmt.Println(averageOfSubtree(node))

	node = &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	fmt.Println(averageOfSubtree(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	_, _, count := nodesAvgs(root)
	return count

}

func nodesAvgs(root *TreeNode) ([]int, []int, int) {
	if isLeaf(root) {
		return []int{root.Val}, []int{root.Val / 1}, 1
	}

	vals := make([]int, 0)
	avgs := make([]int, 0)
	var count int

	if root.Left != nil {
		childVals, childAvgs, childCount := nodesAvgs(root.Left)
		vals = append(vals, childVals...)
		avgs = append(avgs, childAvgs...)
		count += childCount
	}

	if root.Right != nil {
		childVals, childAvgs, childCount := nodesAvgs(root.Right)
		vals = append(vals, childVals...)
		avgs = append(avgs, childAvgs...)
		count += childCount
	}

	vals = append(vals, root.Val)
	avg := sum(vals) / len(vals)
	avgs = append(avgs, avg)

	if root.Val == avg {
		count++
	}

	return vals, avgs, count
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func sum(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
