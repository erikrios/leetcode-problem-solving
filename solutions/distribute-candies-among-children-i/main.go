package main

import (
	"fmt"
)

func main() {
	fmt.Println(distributeCandies(5, 2)) // 3
	fmt.Println(distributeCandies(3, 3)) // 10
	fmt.Println(distributeCandies(1, 2)) // 3
	fmt.Println(distributeCandies(2, 3)) // 6
	fmt.Println(distributeCandies(4, 1)) // 0
	fmt.Println(distributeCandies(2, 1)) // 3
	fmt.Println(distributeCandies(1, 1)) // 3
	fmt.Println(distributeCandies(2, 2)) // 6
	fmt.Println(distributeCandies(3, 1)) // 1
	fmt.Println(distributeCandies(3, 2)) // 7
	fmt.Println(distributeCandies(4, 2)) // 6
}

func distributeCandies(n, limit int) int {
	var res int

	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			for k := 0; k <= limit; k++ {
				if i+j+k == n {
					res++
				}
			}
		}
	}

	return res
}
