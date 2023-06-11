package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countTriples(5))
	fmt.Println(countTriples(10))
}

func countTriples(n int) int {
	var counter int

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			k := i*i + j*j
			d := math.Sqrt(float64(k))
			temp := float64(int(d))
			if d-temp > 0.0 {
				continue
			}
			v := int(d)
			if v <= n {
				if i == j {
					counter++
				} else {
					counter += 2
				}
			} else {
				break
			}
		}
	}

	return counter
}
