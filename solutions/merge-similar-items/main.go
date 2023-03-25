package main

import "fmt"

func main() {
	fmt.Println(mergeSimilarItems([][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}}))
	fmt.Println(mergeSimilarItems([][]int{{1, 1}, {3, 2}, {2, 3}}, [][]int{{2, 1}, {3, 2}, {1, 3}}))
	fmt.Println(mergeSimilarItems([][]int{{1, 3}, {2, 2}}, [][]int{{7, 1}, {2, 2}, {1, 4}}))
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	records := make(map[int]int)
	bst := &bst{}

	for i := 0; i < len(items1) || i < len(items2); i++ {
		if i < len(items1) {
			value := items1[i][0]
			weight := items1[i][1]

			if _, ok := records[value]; !ok {
				bst.insert(value)
			}

			records[value] += weight
		}

		if i < len(items2) {
			value := items2[i][0]
			weight := items2[i][1]

			if _, ok := records[value]; !ok {
				bst.insert(value)
			}

			records[value] += weight
		}
	}

	results := make([][]int, len(records), len(records))
	for i, uniqueVals := 0, bst.toList(); i < len(uniqueVals); i++ {
		uniqueVal := uniqueVals[i]
		weight := records[uniqueVal]
		results[i] = []int{uniqueVal, weight}
	}

	return results
}

type node struct {
	val   int
	left  *node
	right *node
}

type bst struct {
	root *node
}

func (b *bst) insert(v int) {
	b.root = insert(b.root, v)
}

func (b *bst) toList() []int {
	return inorder(b.root)
}

func insert(root *node, v int) *node {
	if root == nil {
		root = &node{val: v}
		return root
	}

	if v < root.val {
		root.left = insert(root.left, v)
	} else {
		root.right = insert(root.right, v)
	}

	return root
}

func inorder(root *node) []int {
	if root == nil {
		return []int{}
	}

	results := make([]int, 0)

	leftVals := inorder(root.left)
	results = append(results, leftVals...)

	results = append(results, root.val)

	rightVals := inorder(root.right)
	results = append(results, rightVals...)

	return results
}
