package main

import "fmt"

func main() {
	fmt.Println(minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
	fmt.Println(minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
	fmt.Println(minMovesToSeat([]int{2, 2, 6, 6}, []int{1, 3, 2, 6}))
	fmt.Println(minMovesToSeat([]int{12, 14, 19, 19, 12}, []int{19, 2, 17, 20, 7}))
}

func minMovesToSeat(seats []int, students []int) int {
	QuickSort(seats)
	QuickSort(students)

	var moves int

	for i := 0; i < len(seats); i++ {
		seat, student := seats[i], students[i]

		if seat > student {
			moves += seat - student
		} else {
			moves += student - seat
		}
	}

	if moves < 0 {
		moves *= -1
	}

	return moves
}

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, startIndex, endIndex int) {
	if endIndex-startIndex <= 0 {
		return
	}

	p := partition(nums, startIndex, endIndex)

	sort(nums, startIndex, p-1)
	sort(nums, p+1, endIndex)
}

func partition(nums []int, startIndex, endIndex int) int {
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

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
