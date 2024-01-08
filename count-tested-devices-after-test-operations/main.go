package main

import "fmt"

func main() {
	fmt.Println(countTestedDevices([]int{1, 1, 2, 1, 3}))
	fmt.Println(countTestedDevices([]int{0, 1, 2}))
}

func countTestedDevices(batteryPercentages []int) int {
	var testedDevices int

	for i := 0; i < len(batteryPercentages); i++ {
		batteryPercentage := batteryPercentages[i]
		if batteryPercentage-testedDevices > 0 {
			testedDevices++
		}
	}

	return testedDevices
}
