package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(jewels string, stones string) int {
	jewelsMapper := make(map[rune]bool)

	for _, jewel := range jewels {
		jewelsMapper[jewel] = true
	}

	var total int

	for _, stone := range stones {
		if _, ok := jewelsMapper[stone]; ok {
			total++
		}
	}

	return total
}
