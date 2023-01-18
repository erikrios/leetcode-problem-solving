package main

import "fmt"

func main() {
	fmt.Println(countMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver"))

	fmt.Println(countMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone"))
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var searchIndex int

	switch ruleKey {
	case "type":
		searchIndex = 0
	case "color":
		searchIndex = 1
	default:
		searchIndex = 2
	}

	var total int

	for _, item := range items {
		if item[searchIndex] == ruleValue {
			total++
		}
	}

	return total
}
