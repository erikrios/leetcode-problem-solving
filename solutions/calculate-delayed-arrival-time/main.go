package main

import "fmt"

func main() {
	fmt.Println(findDelayedArrivalTime(15, 5))
	fmt.Println(findDelayedArrivalTime(13, 11))
}

func findDelayedArrivalTime(arrivalTime, delayedTime int) int {
	const timeFormat = 24
	res := arrivalTime + delayedTime
	if res >= timeFormat {
		res %= timeFormat
	}

	return res
}
