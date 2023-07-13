package main

import "fmt"

func main() {
	set := Constructor()
	set.Add(1)
	set.Add(2)
	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(3))
	set.Add(2)
	fmt.Println(set.Contains(2))
	set.Remove(2)
	fmt.Println(set.Contains(2))
}

type MyHashSet struct {
	Vals [1_000_001]bool
}

func Constructor() MyHashSet {
	return MyHashSet{Vals: [1_000_001]bool{}}
}

func (m *MyHashSet) Add(key int) {
	m.Vals[key] = true
}

func (m *MyHashSet) Remove(key int) {
	m.Vals[key] = false
}

func (m *MyHashSet) Contains(key int) bool {
	return m.Vals[key]
}
