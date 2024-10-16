package main

import "fmt"

func main() {
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
	fmt.Println(findXSum([]int{3, 8, 7, 8, 7, 5}, 2, 2))
}

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	freqs := make(map[int]int)

	results := make([]int, 0)
	for i, j := 0, 0; j < n; j++ {
		num := nums[j]
		freqs[num]++
		if j < k-1 {
			continue
		}

		takes := make([]int, x)

		for k, v := range freqs {
			for i, take := range takes {
				if take == 0 {
					takes[i] = k
					break
				} else {
					if v > freqs[take] {
						temps := make([]int, len(takes[i:]))
						copy(temps, takes[i:])
						takes[i] = k
						copy(takes[i+1:], temps)
						break
					} else if v == freqs[take] {
						if k > take {
							temps := make([]int, len(takes[i:]))
							copy(temps, takes[i:])
							takes[i] = k
							copy(takes[i+1:], temps)
							break
						}
					}
				}
			}
		}

		var sum int
		for _, take := range takes {
			sum += take * freqs[take]
		}

		results = append(results, sum)

		freqs[nums[i]]--
		if freqs[nums[i]] == 0 {
			delete(freqs, nums[i])
		}
		i++
	}

	return results
}
