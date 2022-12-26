package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

type MinStack struct {
	items []int
	min   map[int]int
}

func Constructor() MinStack {
	return MinStack{min: make(map[int]int)}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)

	if len(this.items) == 1 {
		this.min[len(this.items)] = val
	} else {
		prevMin := this.min[len(this.items)-1]
		var min int

		if val < prevMin {
			min = val
		} else {
			min = prevMin
		}

		this.min[len(this.items)] = min
	}
}

func (this *MinStack) Pop() {
	delete(this.min, len(this.items))
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.items)]
}
