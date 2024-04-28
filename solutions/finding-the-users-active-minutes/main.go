package main

import "fmt"

func main() {
	fmt.Println(findingUsersActiveMinutes([][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5))
	fmt.Println(findingUsersActiveMinutes([][]int{{1, 1}, {2, 2}, {2, 3}}, 4))
}

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	n := len(logs)

	usersActiveMinute := make(map[int]map[int]struct{})

	for i := 0; i < n; i++ {
		user, minute := logs[i][0], logs[i][1]

		if _, ok := usersActiveMinute[user]; !ok {
			usersActiveMinute[user] = map[int]struct{}{}
		}
		usersActiveMinute[user][minute] = struct{}{}
	}

	results := make([]int, k)

	for _, v := range usersActiveMinute {
		if len(v)-1 >= k {
			continue
		}

		results[len(v)-1]++
	}

	return results
}
