package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minBitFlips(10, 7))
	fmt.Println(minBitFlips(3, 4))
}

func minBitFlips(start int, goal int) int {
	sBit := strconv.FormatInt(int64(start), 2)
	gBit := strconv.FormatInt(int64(goal), 2)

	if len(sBit) > len(gBit) {
		return flipCounter(sBit, gBit)
	} else {
		return flipCounter(gBit, sBit)
	}
}

func flipCounter(fBit, sBit string) (result int) {
	for i := 0; i < len(fBit); i++ {
		if len(sBit)-1-i < 0 {
			if fBit[len(fBit)-1-i] != '0' {
				result++
			}
			continue
		}

		if fBit[len(fBit)-1-i] != sBit[len(sBit)-1-i] {
			result++
		}
	}

	return result
}
