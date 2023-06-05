package main

import "fmt"

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
	fmt.Println(canBeEqual([]int{7}, []int{7}))
	fmt.Println(canBeEqual([]int{3, 7, 9}, []int{3, 7, 11}))
	fmt.Println(canBeEqual([]int{1, 4, 8, 5, 8}, []int{8, 4, 1, 8, 5}))
	fmt.Println(canBeEqual([]int{1, 5, 8, 5, 8}, []int{8, 1, 1, 8, 5}))
}

func canBeEqual(target []int, arr []int) bool {
	mapper := make(map[int]int)

	for i := 0; i < len(target); i++ {
		num := target[i]
		mapper[num]++
	}

	for i := 0; i < len(arr); i++ {
		num := arr[i]
		if v, ok := mapper[num]; ok {
			v--
			if v == 0 {
				delete(mapper, num)
			} else {
				mapper[num] = v
			}
		} else {
			return false
		}
	}

	return true
}
