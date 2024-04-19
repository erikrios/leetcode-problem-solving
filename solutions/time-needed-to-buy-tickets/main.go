package main

import "fmt"

func main() {
	fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))
	fmt.Println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))
}

func timeRequiredToBuy(tickets []int, k int) int {
	n := len(tickets)
	var i int
	var counter int

	for {
		if i == n {
			i = 0
		}

		if tickets[i] == 0 {
			i++
			continue
		}

		tickets[i]--
		counter++

		if i == k && tickets[i] == 0 {
			break
		}

		i++
	}

	return counter
}
