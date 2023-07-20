package main

import "fmt"

func main() {
	fmt.Println(sortTheStudents([][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}, 2))
	fmt.Println(sortTheStudents([][]int{{3, 4}, {5, 6}}, 0))
}

func sortTheStudents(score [][]int, k int) [][]int {
	m := len(score)

	for i := 0; i < m-1; i++ {
		studentScore := score[i][k]

		for j := i + 1; j < m; j++ {
			nextStudentScore := score[j][k]
			if nextStudentScore > studentScore {
				score[i], score[j] = score[j], score[i]
				studentScore = score[i][k]
			}
		}
	}

	return score
}
