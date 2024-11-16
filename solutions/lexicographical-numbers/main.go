package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(lexicalOrder(13))
	fmt.Println(lexicalOrder(2))
	fmt.Println(lexicalOrder(33))
}

func lexicalOrder(n int) []int {
  var root *BSTNode
  for i := 1; i <= n; i++ {
    root = insert(root, i)
  }

  return bstToList(root)
}

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

func bstToList(root *BSTNode) []int {
  if root == nil {
    return []int{}
  }

  results := bstToList(root.Left)
  results = append(results, root.Val)
  results = append(results, bstToList(root.Right)...)

  return results
}

func insert(root *BSTNode, val int) *BSTNode {
  if root == nil {
    return &BSTNode{Val: val}
  }

  if isXLexicalOrderThanY(val, root.Val) {
    root.Left = insert(root.Left, val)
  } else {
    root.Right = insert(root.Right, val)
  }

  return root
}

func isXLexicalOrderThanY(x, y int) bool {
  xStr := strconv.Itoa(x)
  yStr := strconv.Itoa(y)

  return xStr < yStr
}
