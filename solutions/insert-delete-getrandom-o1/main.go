package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
}

type RandomizedSet struct {
	set map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{set: make(map[int]struct{})}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}

	this.set[val] = struct{}{}
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.set[val]; !ok {
		return false
	}

	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	n := len(this.set)
	randNum := rand.Intn(n)
	var counter int

	for k := range this.set {
		if counter == randNum {
			return k
		}
		counter++
	}

	return -1
}
