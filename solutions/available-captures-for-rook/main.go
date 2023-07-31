package main

func main() {
}

func numRookCaptures(board [][]byte) int {
	var row, col int

loop:
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				row, col = i, j
				break loop
			}
		}
	}

	counter := countTop(row, col, board)
	counter += countDown(row, col, board)
	counter += countLeft(row, col, board)
	counter += countRight(row, col, board)

	return counter
}

func countTop(row, col int, board [][]byte) int {
	var counter int

	for row >= 0 {
		if val := board[row][col]; val == 'B' {
			break
		} else if val == 'p' {
			counter++
			break
		}

		row--
	}

	return counter
}

func countDown(row, col int, board [][]byte) int {
	var counter int

	for row < len(board) {
		if val := board[row][col]; val == 'B' {
			break
		} else if val == 'p' {
			counter++
			break
		}

		row++
	}

	return counter
}

func countLeft(row, col int, board [][]byte) int {
	var counter int

	for col >= 0 {
		if val := board[row][col]; val == 'B' {
			break
		} else if val == 'p' {
			counter++
			break
		}

		col--
	}

	return counter
}

func countRight(row, col int, board [][]byte) int {
	var counter int

	for col < len(board[row]) {
		if val := board[row][col]; val == 'B' {
			break
		} else if val == 'p' {
			counter++
			break
		}

		col++
	}

	return counter
}
