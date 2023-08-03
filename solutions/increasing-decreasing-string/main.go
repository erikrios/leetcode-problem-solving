package main

import "fmt"

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
	fmt.Println(sortString("rat"))
}

func sortString(s string) string {
	sMapper := make(map[byte]uint16)
	for i := 0; i < len(s); i++ {
		sMapper[s[i]]++
	}

	chars := make([]byte, len(sMapper), len(sMapper))
	var index int
	for k := range sMapper {
		chars[index] = k
		index++
	}

	quickSort(chars)

	index = 0
	results := make([]byte, len(s), len(s))

	for len(sMapper) != 0 {
		for i := 0; i < len(chars); i++ {
			char := chars[i]
			v, ok := sMapper[char]
			if ok {
				results[index] = char
				index++
				v--
				if v == 0 {
					delete(sMapper, char)
				} else {
					sMapper[char] = v
				}
			}
		}

		for i := len(chars) - 1; i >= 0; i-- {
			char := chars[i]
			v, ok := sMapper[char]
			if ok {
				results[index] = char
				index++
				v--
				if v == 0 {
					delete(sMapper, char)
				} else {
					sMapper[char] = v
				}
			}
		}
	}

	return string(results)
}

func quickSort(nums []byte) {
	if len(nums) <= 1 {
		return
	}

	sort(nums, 0, len(nums)-1)
}

func sort(nums []byte, startIndex, endIndex int) {
	if endIndex-startIndex <= 0 {
		return
	}

	p := partition(nums, startIndex, endIndex)

	sort(nums, startIndex, p-1)
	sort(nums, p+1, endIndex)
}

func partition(nums []byte, startIndex, endIndex int) int {
	j := startIndex - 1

	for i := startIndex; i < endIndex; i++ {
		if nums[i] < nums[endIndex] {
			j++
			if i == j {
				continue
			}
			swap(nums, j, i)
		}
	}

	swap(nums, j+1, endIndex)

	return j + 1
}

func swap(nums []byte, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
