package main

import (
	"fmt"
)

func main() {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(averageOfLevels(node))

	node = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 20,
		},
	}

	fmt.Println(averageOfLevels(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LevelTrack struct {
	Level int
	Node  *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	queue := NewQueue()
	queue.Enqueue(LevelTrack{0, root})
	records := make([][]int, 0)

	for !queue.IsEmpty() {
		temp := queue.Dequeue()
		if temp.Level > len(records)-1 {
			records = append(records, []int{temp.Node.Val})
		} else {
			records[temp.Level] = append(records[temp.Level], temp.Node.Val)
		}

		if temp.Node.Left != nil {
			queue.Enqueue(LevelTrack{temp.Level + 1, temp.Node.Left})
		}

		if temp.Node.Right != nil {
			queue.Enqueue(LevelTrack{temp.Level + 1, temp.Node.Right})
		}
	}

	n := len(records)
	results := make([]float64, n, n)
	for i := 0; i < n; i++ {
		record := records[i]

		var sum int
		for j := 0; j < len(record); j++ {
			sum += record[j]
		}

		avg := float64(sum) / float64(len(record))
		results[i] = avg
	}

	return results
}

type node struct {
	val  LevelTrack
	prev *node
	next *node
}

func NewNode(val LevelTrack, next, prev *node) *node {
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

func NewLL() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (d *DoublyLinkedList) Insert(v LevelTrack) {
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

func (d *DoublyLinkedList) InsertAt(index int, v LevelTrack) {
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

func (d *DoublyLinkedList) Delete() (result LevelTrack) {
	if d.head == nil {
		return
	}

	if d.head == d.tail {
		result = d.head.val
		d.head, d.tail = nil, nil
		return
	}

	result = d.tail.val
	oldail := d.tail
	d.tail.prev.next = nil
	d.tail = d.tail.prev
	oldail.prev = nil
	return
}

func (d *DoublyLinkedList) ToList() []LevelTrack {
	results := make([]LevelTrack, 0)

	cur := d.head
	for cur != nil {
		results = append(results, cur.val)
		cur = cur.next
	}

	return results
}

type Queue struct {
	ll *DoublyLinkedList
}

func NewQueue() *Queue {
	ll := NewLL()
	return &Queue{ll: ll}
}

func (q *Queue) Enqueue(v LevelTrack) {
	q.ll.InsertAt(0, v)
}

func (q *Queue) Dequeue() (result LevelTrack) {
	result = q.ll.Delete()
	return
}

func (q *Queue) IsEmpty() bool {
	return len(q.ToList()) == 0
}

func (q *Queue) ToList() []LevelTrack {
	listRev := q.ll.ToList()

	if len(listRev) == 0 {
		return []LevelTrack{}
	}

	results := make([]LevelTrack, len(listRev))

	for i := 0; i <= len(listRev)/2; i++ {
		results[i], results[len(listRev)-i-1] = listRev[len(listRev)-i-1], listRev[i]
	}

	return results
}
