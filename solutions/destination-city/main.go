package main

import (
	"fmt"
)

func main() {
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
	fmt.Println(destCity([][]string{{"A", "Z"}}))
}

func destCity(paths [][]string) string {
	cities := make(map[string]*Node)

	for i := 0; i < len(paths); i++ {
		path := paths[i]
		cityA := path[0]
		cityB := path[1]

		var nodeA, nodeB *Node
		var ok bool

		if nodeA, ok = cities[cityA]; !ok {
			node := &Node{Val: cityA}
			cities[cityA] = node
			nodeA = node
		}

		if nodeB, ok = cities[cityB]; !ok {
			node := &Node{Val: cityB}
			cities[cityB] = node
			nodeB = node
		}

		nodeA.Next = nodeB
	}

	var result string
	for k, v := range cities {
		if v.Next == nil {
			result = k
			break
		}
	}

	return result
}

type Node struct {
	Val  string
	Next *Node
}
