package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type MyQueue struct {
	inbox  []int
	outbox []int
}

func Constructor() MyQueue {
	return MyQueue{inbox: make([]int, 0), outbox: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.inbox = append(this.inbox, x)
}

func (this *MyQueue) Pop() int {
	val := this.Peek()
	this.outbox = this.outbox[1:]
	return val
}

func (this *MyQueue) Peek() int {
	m := len(this.inbox)
	n := len(this.outbox)

	if n == 0 {
		this.outbox = make([]int, m)
		copy(this.outbox, this.inbox)
		n = len(this.outbox)
		this.inbox = make([]int, 0)
	}

	val := this.outbox[0]
	return val
}

func (this *MyQueue) Empty() bool {
	m := len(this.inbox)
	n := len(this.outbox)
	return m == 0 && n == 0
}
