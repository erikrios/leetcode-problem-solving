package main

import "fmt"

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
}

func maxLength(arr []string) int {
	return backtrack(arr, make(map[byte]struct{}), 0)
}

func backtrack(arr []string, charSet map[byte]struct{}, i int) int {
	if i == len(arr) {
		return len(charSet)
	}

	var res int
	if !overlap(charSet, arr[i]) {
		for j := 0; j < len(arr[i]); j++ {
			charSet[arr[i][j]] = struct{}{}
		}
		res = backtrack(arr, charSet, i+1)
		for j := 0; j < len(arr[i]); j++ {
			delete(charSet, arr[i][j])
		}
	}
	val := backtrack(arr, charSet, i+1)

	if val > res {
		return val
	}

	return res
}

func overlap(charSet map[byte]struct{}, str string) bool {
	prev := make(map[byte]struct{})

	for i := 0; i < len(str); i++ {
		char := str[i]
		_, charSetOk := charSet[char]
		_, prevOk := prev[char]

		if charSetOk || prevOk {
			return true
		}
		prev[char] = struct{}{}
	}

	return false
}
