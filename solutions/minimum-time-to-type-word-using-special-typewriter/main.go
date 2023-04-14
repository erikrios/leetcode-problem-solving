package main

import "fmt"

func main() {
	fmt.Println(minTimeToType("abc"))
	fmt.Println(minTimeToType("bza"))
	fmt.Println(minTimeToType("zjpc"))
}

func minTimeToType(word string) int {
	var total int
	var pointer byte = 'a'

	for i := 0; i < len(word); i++ {
		letter := word[i]
		if letter == pointer {
			total++
			continue
		}

		vRight := sumClockWise(pointer, letter)
		vLeft := sumCounterClockWise(pointer, letter)

		if vLeft < vRight {
			total += vLeft
		} else {
			total += vRight
		}

		total++
		pointer = letter
	}

	return total
}

func sumClockWise(origin, target byte) int {
	if origin <= target {
		return int(target - origin)
	} else {
		return int('z'-origin+target-'a') + 1
	}
}

func sumCounterClockWise(origin, target byte) int {
	if origin < target {
		return int(origin-'a'+'z'-target) + 1
	} else {
		return int(origin - target)
	}
}
