package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"post", "pos", "postal"}))
}

func longestCommonPrefix(strs []string) (result string) {
	res := make([]byte, 0)

	defer func() {
		_ = recover()
	}()

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return
			}
		}
		res = append(res, strs[0][i])
		result = string(res)
	}

	return
}
