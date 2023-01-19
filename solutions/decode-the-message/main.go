package main

import "fmt"

func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
	fmt.Println(decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
}

func decodeMessage(key string, message string) string {
	mapper := make(map[byte]byte)

	var alphabet byte = 'a'
	for i := 0; i < len(key); i++ {
		if key[i] == ' ' {
			continue
		}

		if _, ok := mapper[key[i]]; ok {
			continue
		}

		mapper[key[i]] = alphabet
		alphabet++
	}

	decoded := make([]byte, len(message), len(message))

	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			decoded[i] = ' '
			continue
		}

		decoded[i] = mapper[message[i]]
	}

	return string(decoded)
}
