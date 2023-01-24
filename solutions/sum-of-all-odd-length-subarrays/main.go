package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Println(sumOddLengthSubarrays([]int{1, 2}))
	fmt.Println(sumOddLengthSubarrays([]int{10, 11, 12}))
}

func sumOddLengthSubarrays(arr []int) int {
	var sum int

	for i := 1; i <= len(arr); i += 2 {
		for j := 0; j+i <= len(arr); j++ {
			subs := arr[j : j+i]
			for k := 0; k < len(subs); k++ {
				sum += subs[k]
			}
		}
	}

	return sum
}
