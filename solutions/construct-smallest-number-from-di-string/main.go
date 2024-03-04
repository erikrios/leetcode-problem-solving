package main

import "fmt"

func main() {
	fmt.Println(smallestNumber("IIIDIDDD"))
	fmt.Println(smallestNumber("DDD"))
}

func smallestNumber(pattern string) string {
	n := len(pattern)
	results := make([]byte, n+1)
	results[0] = '1'
	var start int
	var shouldReverse bool
	for i := 0; i < n; i++ {
		results[i+1] = byte(i) + '0' + 2
		switch char := pattern[i]; char {
		case 'I':
			if shouldReverse {
				reverse(results[start : i+1])
			}
			shouldReverse = false
			start = i + 1
		default:
			shouldReverse = true
		}
	}

	if shouldReverse {
		reverse(results[start : n+1])
	}

	return string(results)
}

func reverse(nums []byte) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
