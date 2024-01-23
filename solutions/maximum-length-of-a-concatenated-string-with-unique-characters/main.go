package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Println(maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
	fmt.Println(maxLength([]string{"cha", "r", "acetslb", "ers"}))
	fmt.Println(maxLength([]string{"fui", "lo", "yr", "i", "hxo", "rou", "q", "spu", "d", "lo", "p", "xjb", "idm", "bwj", "s", "ec"}))
	fmt.Println(maxLength([]string{"dw", "q", "ux", "j", "he", "ev", "ly", "zix", "tth", "x", "t", "r", "ty", "n", "sei", "mb"}))
	fmt.Println(maxLength([]string{"z", "chgtccakarmgp", "ieyfhzxtcczjhs", "i", "kxowcdbynshauqikgg", "aklbjxkczzjiqldciekn", "cvabiynubojuwa", "ctmszammcjwdkyigd", "vswykwxueeo", "ua", "rmwest", "jmjivmbnoexaat", "obbar", "cyek", "vvfxooaacpxdjzsstzbn", "t"}))
	fmt.Println(maxLength([]string{"aa", "bb"}))
	fmt.Println(maxLength([]string{"abc", "def", "bp", "dq", "eg", "fh"}))
	fmt.Println(maxLength([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}))
}

func maxLength(arr []string) int {
	subsets := backtrack(arr, make([]string, 0), 0)
	var max int

loop:
	for i := 0; i < len(subsets); i++ {
		subset := subsets[i]
		charSet := make(map[byte]struct{})

		for j := 0; j < len(subset); j++ {
			sub := subset[j]
			for k := 0; k < len(sub); k++ {
				char := sub[k]
				if _, ok := charSet[char]; ok {
					continue loop
				}
				charSet[char] = struct{}{}
			}
		}

		length := len(charSet)
		if length > max {
			max = length
		}
	}

	return max
}

// Backtracking algorithm
func backtrack(arr, subset []string, i int) [][]string {
	n := len(arr)
	if i == n {
		subsetCopy := make([]string, len(subset))
		copy(subsetCopy, subset)
		return [][]string{subsetCopy}
	}

	str := arr[i]

	results := make([][]string, 0)

	// Take
	subset = append(subset, str)
	res := backtrack(arr, subset, i+1)
	results = append(results, res...)

	// Not take
	subset = subset[:len(subset)-1]
	res = backtrack(arr, subset, i+1)
	results = append(results, res...)

	return results
}
