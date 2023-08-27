package main

import "fmt"

func main() {
	fmt.Println(convertTime("02:30", "04:35"))
	fmt.Println(convertTime("11:00", "11:01"))
}

func convertTime(current, correct string) int {
	timeRange := toMinutes(correct) - toMinutes(current)
	var counter int

	for timeRange > 0 {
		var reduce int

		if timeRange >= 60 {
			reduce = 60
		} else if timeRange >= 15 && timeRange < 60 {
			reduce = 15
		} else if timeRange >= 5 && timeRange < 15 {
			reduce = 5
		} else {
			reduce = 1
		}

		timeRange -= reduce
		counter++
	}

	return counter
}

func toMinutes(time string) int {
	hours := time[:2]
	minutes := time[3:]

	var res int

	res += int(minutes[1] - '0')
	res += (int(minutes[0]-'0') * 10)

	res += (int(hours[1]-'0') * 60)
	res += (int(hours[0]-'0') * 60 * 10)

	return res
}
