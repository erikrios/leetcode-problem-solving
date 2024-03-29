package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddr("255.100.50.0"))
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
