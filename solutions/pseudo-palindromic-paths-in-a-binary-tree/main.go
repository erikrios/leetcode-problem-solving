package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	fmt.Println(pseudoPalindromicPaths(node))

	node = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	fmt.Println(pseudoPalindromicPaths(node))

	node = &TreeNode{Val: 9}

	fmt.Println(pseudoPalindromicPaths(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	return palindromicPaths(root, make(map[byte]bool))
}

func palindromicPaths(root *TreeNode, set map[byte]bool) int {
	if root == nil {
		return 0
	}

	if isLeaf(root) {
		val := byte(root.Val)
		_, ok := set[val]
		negateOrDeleteSetKey(set, val, false)
		var counter int
		for _, v := range set {
			if v {
				counter++
			}
		}
		negateOrDeleteSetKey(set, val, !ok)
		if counter > 1 {
			return 0
		}

		return 1
	}

	val := byte(root.Val)
	_, ok := set[val]
	negateOrDeleteSetKey(set, val, false)

	leftCount := palindromicPaths(root.Left, set)
	rightCount := palindromicPaths(root.Right, set)

	negateOrDeleteSetKey(set, val, !ok)

	return leftCount + rightCount
}

func negateOrDeleteSetKey(set map[byte]bool, key byte, shouldDelete bool) {
	if shouldDelete {
		delete(set, key)
	} else {
		set[key] = !set[key]
	}
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
