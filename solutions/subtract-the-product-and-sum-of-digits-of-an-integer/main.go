package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subtractProductAndSum(234))
	fmt.Println(subtractProductAndSum(4421))
}

func subtractProductAndSum(n int) int {
	nStr := strconv.Itoa(n)

	var multiply int = 1
	var sum int

	for _, v := range nStr {
		nVal, _ := strconv.Atoi(string(v))
		multiply *= nVal
		sum += nVal
	}

	return multiply - sum
}
