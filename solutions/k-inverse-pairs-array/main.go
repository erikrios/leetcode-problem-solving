package main

import "fmt"

func main() {
	fmt.Println(kInversePairs(3, 0))
	fmt.Println(kInversePairs(3, 1))
	fmt.Println(kInversePairs(1000, 1))
	fmt.Println(kInversePairs(1000, 1000))
}

func kInversePairs(n, k int) int {
	const mod = 1_000_000_007
	prev := make([]int, k+1)
	prev[0] = 1

	for i := 1; i < n+1; i++ {
		cur := make([]int, k+1)
		total := 0
		for j := 0; j < k+1; j++ {
			if j >= i {
				total = (total - prev[j-i] + mod) % mod
			}
			total = (total + prev[j]) % mod
			cur[j] = total
		}
		prev = cur
	}
	return prev[k]
}
