package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
	fmt.Println(groupAnagrams([]string{"hhhhu", "tttti", "tttit", "hhhuh", "hhuhh", "tittt"}))
	fmt.Println(groupAnagrams([]string{"", "b", ""}))
	fmt.Println(groupAnagrams([]string{"ac", "c"}))
	fmt.Println(groupAnagrams([]string{"ddddddddddg", "dgggggggggg"}))
}

type Key [26]byte

func strKey(str string) (key Key) {
	for i := 0; i < len(str); i++ {
		key[str[i]-'a']++
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Key][]string)

	for i := 0; i < len(strs); i++ {
		str := strs[i]
		key := strKey(str)
		groups[key] = append(groups[key], str)
	}

	results := make([][]string, 0, len(groups))
	for _, v := range groups {
		results = append(results, v)
	}

	return results
}
