package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
	fmt.Println(deckRevealedIncreasing([]int{1, 1000}))
}

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	queue := NewQueue()

	for i := 0; i < n; i++ {
		queue.Enqueue(i)
	}

	sort.Ints(deck)

	ans := make([]int, n)
	for _, card := range deck {
		ans[queue.Dequeue()] = card

		if !queue.IsEmpty() {
			queue.Enqueue(queue.Dequeue())
		}
	}

	return ans
}

type Queue struct {
	front *ListNode
	rear  *ListNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val int) {
	newNode := &ListNode{Val: val}
	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.Next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	val := q.front.Val

	if q.front == q.rear {
		q.front = nil
		q.rear = nil
		return val
	}

	q.front = q.front.Next
	return val
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
