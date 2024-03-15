package main

import "fmt"

func main() {
	fmt.Println(numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
	fmt.Println(numberOfBeams([]string{"000", "111", "000"}))
}

func numberOfBeams(bank []string) int {
	var totalBeams int
	var prevLaserCount int

	for i := 0; i < len(bank); i++ {
		row := bank[i]
		var totalLaser int
		for j := 0; j < len(row); j++ {
			if row[j] == '1' {
				totalLaser++
			}
		}

		if totalLaser == 0 {
			continue
		}

		totalBeams += prevLaserCount * totalLaser
		prevLaserCount = totalLaser
	}

	return totalBeams
}
