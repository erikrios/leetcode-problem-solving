package main

import "fmt"

func main() {
	browserHistory := Constructor("leetcode.com")
	browserHistory.Visit("google.com")
	browserHistory.Visit("facebook.com")
	browserHistory.Visit("youtube.com")
	fmt.Println(browserHistory.Back(1))
	fmt.Println(browserHistory.Back(1))
	fmt.Println(browserHistory.Forward(1))
	browserHistory.Visit("linkedin.com")
	fmt.Println(browserHistory.Forward(2))
	fmt.Println(browserHistory.Back(2))
	fmt.Println(browserHistory.Back(7))
}

type BrowserHistory struct {
	node *DoublyListNode
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		node: &DoublyListNode{Val: homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {
	newNode := &DoublyListNode{
		Val: url,
	}

	newNode.Prev = this.node
	this.node.Next = newNode
	this.node = newNode
}

func (this *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps; i++ {
		if this.node.Prev == nil {
			break
		}
		this.node = this.node.Prev
	}

	return this.node.Val
}

func (this *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps; i++ {
		if this.node.Next == nil {
			break
		}
		this.node = this.node.Next
	}

	return this.node.Val
}

type DoublyListNode struct {
	Val  string
	Next *DoublyListNode
	Prev *DoublyListNode
}
