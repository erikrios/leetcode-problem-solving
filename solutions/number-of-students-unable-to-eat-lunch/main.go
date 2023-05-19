package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}

func countStudents(students []int, sandwiches []int) int {
	studentsCounter := make(map[int]int, 2)
	studentsQueue := NewQueue()
	sandwichesStack := NewStack()

	for i := 0; i < len(students); i++ {
		student := students[i]
		studentsQueue.Enqueue(student)
		studentsCounter[student]++

		sandwich := sandwiches[len(sandwiches)-1-i]
		sandwichesStack.Push(sandwich)
	}

	for len(studentsCounter) > 0 {
		sandwich := sandwichesStack.Peek()
		student := studentsQueue.Dequeue()

		if sandwich == student {
			sandwichesStack.Pop()
			studentsCounter[student]--
			if studentsCounter[student] == 0 {
				delete(studentsCounter, student)
			}
		} else {
			studentsQueue.Enqueue(student)
			if _, ok := studentsCounter[sandwich]; !ok {
				return studentsCounter[student]
			}
		}
	}

	return 0
}

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		return 0
	}

	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v
}

func (s *Stack) Peek() int {
	if s.isEmpty() {
		return 0
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

type Queue struct {
	ll *DoublyLinkedList
}

func NewQueue() *Queue {
	ll := NewDoublyLinkedList()
	return &Queue{ll: ll}
}

func (q *Queue) Enqueue(v int) {
	q.ll.InsertAt(0, v)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return 0
	}

	return q.ll.Delete()
}

func (q *Queue) IsEmpty() bool {
	return len(q.ToList()) == 0
}

func (q *Queue) ToList() []int {
	listRev := q.ll.ToList()

	if len(listRev) == 0 {
		return []int{}
	}

	results := make([]int, len(listRev))

	for i := 0; i <= len(listRev)/2; i++ {
		results[i], results[len(listRev)-i-1] = listRev[len(listRev)-i-1], listRev[i]
	}

	return results
}

type node struct {
	val  int
	prev *node
	next *node
}

func NewNode(val int, next, prev *node) *node {
	return &node{
		val:  val,
		next: next,
		prev: prev,
	}
}

type DoublyLinkedList struct {
	head *node
	tail *node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (d *DoublyLinkedList) Insert(v int) {
	newNode := NewNode(v, nil, nil)

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode
}

func (d *DoublyLinkedList) InsertAt(index int, v int) {
	newNode := NewNode(v, nil, nil)

	if index == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
			return
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
			return
		}
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	cur := d.head
	for i := 0; i < index; i++ {
		if cur.next != nil {
			cur = cur.next
		} else {
			cur.next = newNode
			newNode.prev = cur
			d.tail = newNode
			return
		}
	}

	newNode.next = cur
	newNode.prev = cur.prev
	cur.prev.next = newNode
	cur.prev = newNode
}

func (d *DoublyLinkedList) Find(v int) bool {
	var isExists bool

	cur := d.head
	for cur != nil {
		if reflect.DeepEqual(cur.val, v) {
			isExists = true
			break
		}
		cur = cur.next
	}

	return isExists
}

func (d *DoublyLinkedList) Delete() (result int) {
	if d.head == nil {
		return 0
	}

	if d.head == d.tail {
		result := d.head.val
		d.head, d.tail = nil, nil
		return result
	}

	result = d.tail.val
	oldTail := d.tail
	d.tail.prev.next = nil
	d.tail = d.tail.prev
	oldTail.prev = nil
	return
}

func (d *DoublyLinkedList) ToList() []int {
	results := make([]int, 0)

	cur := d.head
	for cur != nil {
		results = append(results, cur.val)
		cur = cur.next
	}

	return results
}
