package main

import "fmt"

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
	fmt.Println(makeGood("s"))
}

func makeGood(s string) string {
	chars := []byte(s)

	for i, j := 0, 1; j < len(chars); {
		cur := chars[i]
		next := chars[j]

		if isBad(cur, next) {
			front := chars[:i]
			back := chars[j+1:]

			front = append(front, back...)
			chars = front

			if i > 0 {
				i--
				j--
			}
		} else {
			i++
			j++
		}
	}

	return string(chars)
}

func isBad(x, y byte) bool {
	isCurLower := isLower(x)
	isNextLower := isLower(y)

	if isCurLower && isNextLower {
		return false
	}

	if !isCurLower && !isNextLower {
		return false
	}

	if toLower(x) != toLower(y) {
		return false
	}

	return true
}

func toLower(char byte) byte {
	if isLower(char) {
		return char
	}

	return 'a' + char - 'A'
}

func isLower(char byte) bool {
	return char >= 'a' && char <= 'z'
}
