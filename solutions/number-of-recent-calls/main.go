package main

import "fmt"

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{requests: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	min := t - 3000
	i := binarySearch(this.requests, min)
	return len(this.requests) - i
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if num := nums[mid]; num == target {
			return mid
		} else if num < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
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

func traverse(root *node, target int) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)

	if root.val >= target {
		lRes := traverse(root.left, target)
		res = append(res, lRes...)
	}
	if root.val >= target {
		res = append(res, root.val)
	}
	rRes := traverse(root.right, target)
	res = append(res, rRes...)

	return res
}

func findMin(root *node, target int) *node {
	if root == nil {
		return nil
	}

	if root.val == target {
		return root
	} else if target < root.val {
		l := findMin(root.left, target)
		if l == nil && root.val > target {
			return root
		} else {
			return l
		}
	} else {
		r := findMin(root.right, target)
		if r == nil && root.val > target {
			return root
		} else {
			return r
		}
	}
}

//               8
//        4            10
//   2              9
//      3

func main() {
	node := insert(nil, 8)
	node = insert(node, 4)
	node = insert(node, 10)
	node = insert(node, 2)
	node = insert(node, 3)
	node = insert(node, 9)

	res := traverse(node, 5)
	fmt.Println(res)

	res = traverse(node, 1)
	fmt.Println(res)

	res = traverse(node, 9)
	fmt.Println(res)

	fmt.Println(findMin(node, 2))
	fmt.Println(findMin(node, 5))
	fmt.Println(findMin(node, -1000))
}
