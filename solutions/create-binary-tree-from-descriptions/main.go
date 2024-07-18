package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	mapper := make(map[int]struct {
		node   *TreeNode
		isRoot bool
	})

	for _, desc := range descriptions {
		parentVal := desc[0]
		childVal := desc[1]
		isLeft := desc[2] == 1

		parent, ok := mapper[parentVal]
		if !ok {
			node := &TreeNode{Val: parentVal}
			parent.node = node
			parent.isRoot = true
			mapper[parentVal] = parent
		}

		child, ok := mapper[childVal]
		if !ok {
			node := &TreeNode{Val: childVal}
			child.node = node
		}
		child.isRoot = false
		mapper[childVal] = child

		if isLeft {
			parent.node.Left = child.node
		} else {
			parent.node.Right = child.node
		}
	}

	for _, v := range mapper {
		if v.isRoot {
			return v.node
		}
	}

	return nil
}
