package main

func main() {}

func findChampion(grid [][]int) int {
	n := len(grid)

	var winner int
	var total int
	for i := 0; i < n; i++ {
		var counter int
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				counter++
			}
		}
		if counter > total {
			total = counter
			winner = i
		}

	}

	return winner
}
