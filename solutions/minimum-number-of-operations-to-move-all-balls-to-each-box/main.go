package main

import "fmt"

func main() {
	fmt.Println(minOperations("110"))
	fmt.Println(minOperations("001011"))
}

func minOperations(boxes string) []int {
	n := len(boxes)
	sets := make(map[int]struct{})

	for i := 0; i < n; i++ {
		if box := boxes[i]; box == '1' {
			sets[i] = struct{}{}
		}
	}

	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var ops int
		for k := range sets {
			if k == i {
				continue
			}

			var dist int
			if i >= k {
				dist = i - k
			} else {
				dist = k - i
			}

			ops += dist
		}

		res = append(res, ops)
	}

	return res
}
