package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
}

func maximum69Number(num int) int {
	numStr := strconv.Itoa(num)
	var isChanged bool

	nums := make([]byte, len(numStr), len(numStr))
	for i := 0; i < len(numStr); i++ {
		numByte := numStr[i]
		if !isChanged && numByte == '6' {
			numByte = '9'
			isChanged = true
		}

		nums[i] = numByte
	}

	res, _ := strconv.Atoi(string(nums))
	return res
}
