package main

import "fmt"

func main() {
	fmt.Println(countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
	fmt.Println(countSeniors([]string{"1313579440F2036", "2921522980M5644"}))
}

func countSeniors(details []string) int {
	var counter int

	for i := 0; i < len(details); i++ {
		detail := details[i]
		firstChar := detail[11] - '0'
		secondChar := detail[12] - '0'
		age := firstChar*10 + secondChar
		if age > 60 {
			counter++
		}
	}

	return counter
}
