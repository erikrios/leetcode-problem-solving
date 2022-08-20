package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(2))
	fmt.Println(convertToTitle(3))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
}

func convertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}

	columnNumber = columnNumber - 1
	res := 'A' + (columnNumber % 26)
	val := byte(res)
	vals := []byte{val}
	return convertToTitle(columnNumber/26) + string(vals)
}
