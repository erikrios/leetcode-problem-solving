package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(15))
	fmt.Println(reverse(-26))
	fmt.Println(reverse(100))
	fmt.Println(reverse(-50000))
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	strRev := reverseString(str)
	strRev = trimZero(strRev)

	if x < 0 {
		v, _ := strconv.Atoi(string(strRev[:len(strRev)-1]))
		v *= -1
		if v < math.MinInt32 {
			return 0
		}
		return v
	}

	v, _ := strconv.Atoi(strRev)
	if v > math.MaxInt32 {
		return 0
	}
	return v
}

func reverseString(s string) string {
	res := make([]byte, len(s))

	for i := 0; i < (len(s)+1)/2; i++ {
		res[i], res[len(res)-1-i] = s[len(s)-1-i], s[i]
	}

	return string(res)
}

func trimZero(s string) string {
	var counter int

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			counter++
		} else {
			break
		}
	}

	return string(s[counter:])
}
