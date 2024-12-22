package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9}))
	fmt.Println(largestNumber([]int{432, 43243}))
	fmt.Println(largestNumber([]int{8308, 8308, 830}))
	fmt.Println(largestNumber([]int{111311, 1113}))
	fmt.Println(largestNumber([]int{0, 0}))
	fmt.Println(largestNumber([]int{0, 0, 0}))
}

func largestNumber(nums []int) string {
	var root *BSTNode
	for _, num := range nums {
		root = insert(root, num)
	}

	var builder strings.Builder
	traverse(root, &builder)

	results := strings.TrimLeft(builder.String(), "0")
	if len(results) == 0 {
		return "0"
	}

	return results
}

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

func insert(root *BSTNode, val int) *BSTNode {
	if root == nil {
		return &BSTNode{Val: val}
	}

	if !isXGreaterThanY(val, root.Val) {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}

	return root
}

func traverse(root *BSTNode, builder *strings.Builder) {
	if root == nil {
		return
	}

	traverse(root.Right, builder)
	builder.WriteString(strconv.Itoa(root.Val))
	traverse(root.Left, builder)
}

func isXGreaterThanY(x, y int) bool {
	xStr := strconv.Itoa(x)
	yStr := strconv.Itoa(y)

	xLen := len(xStr)
	yLen := len(yStr)

	for i := 0; i < xLen+yLen; i++ {
		xNum := xStr[i%xLen]
		yNum := yStr[i%yLen]

		if xNum > yNum {
			return true
		} else if yNum > xNum {
			return false
		}
	}

	return true
}
