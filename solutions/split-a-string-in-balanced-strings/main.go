package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))
	fmt.Println(balancedStringSplit("LLLLRRRR"))
}

func balancedStringSplit(s string) int {
	var counter int

	var lTotal, rTotal int
	for i := 0; i < len(s); i += 2 {
		switch s[i] {
		case 'L':
			lTotal++
		default:
			rTotal++
		}

		switch s[i+1] {
		case 'L':
			lTotal++
		default:
			rTotal++
		}

		if lTotal == rTotal {
			counter++
		}
	}

	return counter
}
