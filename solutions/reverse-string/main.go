package main

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	reverseString([]byte{'H', 'a', 'n', 'n', 'a', 'h'})
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
