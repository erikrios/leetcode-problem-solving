package main

import "fmt"

func main() {
	fmt.Println(minimumDeletions("aababbab"))
	fmt.Println(minimumDeletions("bbaaaaabb"))
	fmt.Println(minimumDeletions("a"))
}

func minimumDeletions(s string) int {
	n := len(s)
	aRightTotals := make([]int, n)

	for i := n - 1 - 1; i >= 0; i-- {
		aRightTotals[i] = aRightTotals[i+1]
		next := s[i+1]
		if next == 'a' {
			aRightTotals[i] += 1
		}
	}

	res := n
	var bRightTotal int
	for i := 0; i < n; i++ {
		total := aRightTotals[i] + bRightTotal
		if total < res {
			res = total
		}
		cur := s[i]
		if cur == 'b' {
			bRightTotal++
		}
	}

	return res
}
