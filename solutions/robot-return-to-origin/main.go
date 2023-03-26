package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
}

func judgeCircle(moves string) bool {
	var x, y int

	for i := 0; i < len(moves); i++ {
		move(&x, &y, moves[i])
	}

	return x == 0 && y == 0
}

func move(x, y *int, direction byte) {
	switch direction {
	case 'L':
		*x--
	case 'R':
		*x++
	case 'U':
		*y++
	default:
		*y--
	}
}
