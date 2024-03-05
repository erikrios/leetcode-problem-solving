package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type CustomStack struct {
	nums []int
	ptr  int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{nums: make([]int, maxSize), ptr: -1}
}

func (this *CustomStack) Push(x int) {
	if this.ptr < len(this.nums)-1 {
		this.ptr++
		this.nums[this.ptr] = x
	}
}

func (this *CustomStack) Pop() int {
	if this.ptr == -1 {
		return -1
	}
	num := this.nums[this.ptr]
	this.ptr--
	return num
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k; i++ {
		if i > this.ptr || i > len(this.nums)-1 {
			break
		}
		this.nums[i] += val
	}
}
