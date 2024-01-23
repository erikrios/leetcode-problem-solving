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
	return backtrack(arr, make(map[byte]struct{}), 0)
}

// Backtracking algorithm
func backtrack(arr []string, charSet map[byte]struct{}, i int) int {
	n := len(arr)
	if i == n {
		return len(charSet)
	}

	str := arr[i]

	var max int
	if isUnique(charSet, str) {
		// Take
		for i := 0; i < len(str); i++ {
			char := str[i]
			charSet[char] = struct{}{}
		}

		max = backtrack(arr, charSet, i+1)

		for i := 0; i < len(str); i++ {
			char := str[i]
			delete(charSet, char)
		}
	}

	// Not take
	val := backtrack(arr, charSet, i+1)
	if val > max {
		return val
	}

	return max
}

func isUnique(charSet map[byte]struct{}, str string) bool {
	curSet := make(map[byte]struct{})

	for i := 0; i < len(str); i++ {
		char := str[i]

		_, charSetOk := charSet[char]
		_, curSetOk := curSet[char]

		if charSetOk || curSetOk {
			return false
		}

		curSet[char] = struct{}{}
	}

	return true
}
