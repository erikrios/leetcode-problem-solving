package main

import "fmt"

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	fmt.Println(sortPeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))
}

func sortPeople(names []string, heights []int) []string {
	QuickSort(names, heights)
	return names
}

func QuickSort(names []string, heights []int) {
	if len(heights) <= 1 {
		return
	}

	sort(names, heights, 0, len(heights)-1)
}

func sort(names []string, heights []int, startIndex, endIndex int) {
	if endIndex-startIndex <= 0 {
		return
	}

	p := partition(names, heights, startIndex, endIndex)

	sort(names, heights, startIndex, p-1)
	sort(names, heights, p+1, endIndex)
}

func partition(names []string, heights []int, startIndex, endIndex int) int {
	j := startIndex - 1

	for i := startIndex; i < endIndex; i++ {
		if heights[i] > heights[endIndex] {
			j++
			if i == j {
				continue
			}
			swap(names, heights, j, i)
		}
	}

	swap(names, heights, j+1, endIndex)

	return j + 1
}

func swap(names []string, heights []int, i, j int) {
	heights[i], heights[j] = heights[j], heights[i]
	names[i], names[j] = names[j], names[i]
}
