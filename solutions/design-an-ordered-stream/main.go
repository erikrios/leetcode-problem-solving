package main

import "fmt"

func main() {
	os := Constructor(5)
	fmt.Println(os.Insert(3, "ccccc"))
	fmt.Println(os.Insert(1, "aaaaa"))
	fmt.Println(os.Insert(2, "bbbbb"))
	fmt.Println(os.Insert(5, "eeeee"))
	fmt.Println(os.Insert(4, "ddddd"))
}

type OrderedStream struct {
	n      int
	ptr    int
	values []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{n: n, values: make([]string, n, n)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	results := make([]string, 0, this.n)

	this.values[idKey-1] = value

	for this.ptr < this.n && this.values[this.ptr] != "" {
		results = append(results, this.values[this.ptr])
		this.ptr++
	}

	return results
}
