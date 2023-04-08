package main

import "fmt"

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}

func sortByBits(arr []int) []int {
	var bitCounters [16]*node

	for i := 0; i < len(arr); i++ {
		num := arr[i]

		var counter int
		for num > 0 {
			digit := num % 2
			if digit == 1 {
				counter++
			}
			num /= 2
		}

		bitCounters[counter] = insert(bitCounters[counter], arr[i])
	}

	results := make([]int, 0)
	for i := 0; i < len(bitCounters); i++ {
		node := bitCounters[i]

		if node == nil {
			continue
		}

		nums := traverse(node)
		results = append(results, nums...)
	}

	return results
}

type node struct {
	val   int
	left  *node
	right *node
}

func insert(root *node, val int) *node {
	if root == nil {
		return &node{val: val}
	}

	if val <= root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}

	return root
}

func traverse(root *node) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)

	lRes := traverse(root.left)
	res = append(res, lRes...)
	res = append(res, root.val)
	rRes := traverse(root.right)
	res = append(res, rRes...)

	return res
}
