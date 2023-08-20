package main

import "fmt"

func main() {
	myMap := Constructor()
	myMap.Put(1, 5)
	fmt.Println(myMap.Get(1))
	myMap.Put(1, 3)
	fmt.Println(myMap.Get(1))
	fmt.Println(myMap.Get(3))
	myMap.Put(2, 4)
	fmt.Println(myMap.Get(2))
	myMap.Remove(2)
	fmt.Println(myMap.Get(2))
}

const size = 1e6 + 1

type MyHashMap struct {
	list [size]int
}

func Constructor() MyHashMap {
	list := [size]int{}
	for i := 0; i < size; i++ {
		list[i] = -1
	}
	return MyHashMap{list}
}

func (this *MyHashMap) Put(key int, value int) {
	this.list[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.list[key]
}

func (this *MyHashMap) Remove(key int) {
	this.list[key] = -1
}

//type MyHashMap struct {
//	linkeds [9]*LinkedList
//}
//
//func Constructor() MyHashMap {
//	linkeds := [9]*LinkedList{}
//	for i := 0; i < len(linkeds); i++ {
//		linkeds[i] = NewLinkedList()
//	}
//	return MyHashMap{linkeds}
//}
//
//func (this *MyHashMap) Put(key int, value int) {
//	if ok, val := this.linkeds[key%9].Get(key); ok {
//		val.Val = value
//		return
//	}
//
//	keyVal := &KeyVal{Key: key, Val: value}
//	this.linkeds[key%9].Insert(keyVal)
//}
//
//func (this *MyHashMap) Get(key int) int {
//	if ok, val := this.linkeds[key%9].Get(key); ok {
//		return val.Val
//	}
//
//	return -1
//}
//
//func (this *MyHashMap) Remove(key int) {
//	this.linkeds[key%9].Remove(key)
//}
//
//type KeyVal struct {
//	Key int
//	Val int
//}
//
//type Node struct {
//	Val  *KeyVal
//	Next *Node
//}
//
//type LinkedList struct {
//	Head  *Node
//	Tail  *Node
//	Count int
//}
//
//func NewLinkedList() *LinkedList {
//	return &LinkedList{}
//}
//
//func (l *LinkedList) Insert(val *KeyVal) {
//	node := &Node{Val: val}
//	l.Count++
//
//	if l.Head == nil {
//		l.Head = node
//		l.Tail = node
//		return
//	}
//
//	l.Tail.Next = node
//	l.Tail = node
//}
//
//func (l *LinkedList) Remove(key int) {
//	if l.Size() == 0 {
//		return
//	}
//
//	var prev *Node
//	for cur := l.Head; cur != nil; prev, cur = cur, cur.Next {
//		if key == cur.Val.Key {
//			if l.Size() == 1 {
//				l.Head = nil
//				l.Tail = nil
//			} else if l.Size() == 2 && cur == l.Head {
//				l.Head = l.Tail
//			} else if cur == l.Head {
//				l.Head = l.Head.Next
//			} else {
//				prev.Next = cur.Next
//				cur.Next = nil
//			}
//
//			l.Count--
//			return
//		}
//	}
//}
//
//func (l *LinkedList) Size() int {
//	return l.Count
//}
//
//func (l *LinkedList) Get(key int) (ok bool, val *KeyVal) {
//	for cur := l.Head; cur != nil; cur = cur.Next {
//		if key == cur.Val.Key {
//			ok = true
//			val = cur.Val
//			break
//		}
//	}
//
//	return
//}
