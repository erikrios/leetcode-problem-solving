package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findWinners([][]int{
		{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9},
	}))
	fmt.Println(findWinners([][]int{
		{2, 3}, {1, 3}, {5, 4}, {6, 4},
	}))
}

func findWinners(matches [][]int) [][]int {
	n := len(matches)

	matchMap := make(map[int][2]int)
	for i := 0; i < n; i++ {
		match := matches[i]
		winner := match[0]
		loser := match[1]

		winnerStatsRecords := matchMap[winner]
		winnerStatsRecords[0]++
		matchMap[winner] = winnerStatsRecords

		loserStatsRecords := matchMap[loser]
		loserStatsRecords[1]++
		matchMap[loser] = loserStatsRecords
	}

	alwaysWinPlayers := make([]int, 0)
	loseExactlyOnePlayers := make([]int, 0)
	for k, values := range matchMap {
		winCount := values[0]
		loseCount := values[1]

		if winCount > 0 && loseCount == 0 {
			alwaysWinPlayers = append(alwaysWinPlayers, k)
		} else if loseCount == 1 {
			loseExactlyOnePlayers = append(loseExactlyOnePlayers, k)
		}
	}

	sort.Ints(alwaysWinPlayers)
	sort.Ints(loseExactlyOnePlayers)

	return [][]int{alwaysWinPlayers, loseExactlyOnePlayers}
}
