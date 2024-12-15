package main

import (
	"fmt"
	"strings"
)

func main() {
	construct([][]int{
		{0, 1}, {1, 0},
	}).printQuadTree(0)

	fmt.Println("================")

	construct([][]int{
		{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0},
	}).printQuadTree(0)
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// Print visualization of the quadtree
func (n *Node) printQuadTree(indent int) {
	if n == nil {
		return
	}

	// Create indentation
	padding := strings.Repeat("  ", indent)

	if n.IsLeaf {
		fmt.Printf("%sLeaf Node - Value: %v\n", padding, n.Val)
	} else {
		fmt.Printf("%sInternal Node\n", padding)
		fmt.Printf("%sTopLeft:\n", padding)
		n.TopLeft.printQuadTree(indent + 1)
		fmt.Printf("%sTopRight:\n", padding)
		n.TopRight.printQuadTree(indent + 1)
		fmt.Printf("%sBottomLeft:\n", padding)
		n.BottomLeft.printQuadTree(indent + 1)
		fmt.Printf("%sBottomRight:\n", padding)
		n.BottomRight.printQuadTree(indent + 1)
	}
}

func construct(grid [][]int) *Node {
	n := len(grid)

	root := &Node{
		Val:    grid[0][0] == 1,
		IsLeaf: true,
	}

	if n == 1 {
		return root
	}

	topLeftGrid := make([][]int, len(grid[:n/2]))
	copy(topLeftGrid, grid[:n/2])
	for i := 0; i < len(topLeftGrid); i++ {
		topLeftGrid[i] = topLeftGrid[i][:n/2]
	}
	topRightGrid := make([][]int, len(grid[:n/2]))
	copy(topRightGrid, grid[:n/2])
	for i := 0; i < len(topRightGrid); i++ {
		topRightGrid[i] = topRightGrid[i][n/2:]
	}
	bottomLeftGrid := make([][]int, len(grid[n/2:]))
	copy(bottomLeftGrid, grid[n/2:])
	for i := 0; i < len(bottomLeftGrid); i++ {
		bottomLeftGrid[i] = bottomLeftGrid[i][:n/2]
	}
	bottomRightGrid := make([][]int, len(grid[n/2:]))
	copy(bottomRightGrid, grid[n/2:])
	for i := 0; i < len(bottomRightGrid); i++ {
		bottomRightGrid[i] = bottomRightGrid[i][n/2:]
	}

	topLeftNode := construct(topLeftGrid)
	topRightNode := construct(topRightGrid)
	bottomLeftNode := construct(bottomLeftGrid)
	bottomRightNode := construct(bottomRightGrid)

	if topLeftNode.IsLeaf && topRightNode.IsLeaf && bottomLeftNode.IsLeaf && bottomRightNode.IsLeaf && topLeftNode.Val == topRightNode.Val && topRightNode.Val == bottomLeftNode.Val && bottomLeftNode.Val == bottomRightNode.Val {
		return root
	}

	root.IsLeaf = false
	root.TopLeft = topLeftNode
	root.TopRight = topRightNode
	root.BottomLeft = bottomLeftNode
	root.BottomRight = bottomRightNode

	return root
}
