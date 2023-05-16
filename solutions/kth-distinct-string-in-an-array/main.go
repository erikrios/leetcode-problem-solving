package main

import "fmt"

func main() {
	fmt.Println(kthDistinct([]string{"d", "b", "c", "b", "c", "a"}, 2))
	fmt.Println(kthDistinct([]string{"aaa", "aa", "a"}, 1))
	fmt.Println(kthDistinct([]string{"a", "b", "a"}, 3))
}

func kthDistinct(arr []string, k int) string {
	n := len(arr)
	counts := make(map[string]uint16, n)

	for i := 0; i < n; i++ {
		counts[arr[i]]++
	}

	var res string
	var counter int
	for i := 0; i < n; i++ {
		str := arr[i]
		if counts[str] == 1 {
			counter++
		}

		if counter == k {
			res = str
			break
		}
	}

	return res
}
