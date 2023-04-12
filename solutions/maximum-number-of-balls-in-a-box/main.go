package main

import "fmt"

func main() {
	fmt.Println(countBalls(1, 10))
	fmt.Println(countBalls(5, 15))
	fmt.Println(countBalls(19, 28))
}

func countBalls(lowLimit int, highLimit int) int {
	boxes := make(map[int]int)

	for i := lowLimit; i <= highLimit; i++ {
		num := i
		var boxNum int

		for num > 0 {
			boxNum += num % 10
			num /= 10
		}

		boxes[boxNum]++
	}

	var maxBox int
	for _, v := range boxes {
		if v > maxBox {
			maxBox = v
		}
	}

	return maxBox
}
