package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	shorterner := Constructor()
	encoded := shorterner.encode("https://leetcode.com/erikrios")
	fmt.Println(encoded)
	decoded := shorterner.decode(encoded)
	fmt.Println(decoded)
}

type Codec struct {
	urls    map[string]string
	baseUrl string
}

func Constructor() Codec {
	return Codec{
		urls:    make(map[string]string),
		baseUrl: "https://er.ik",
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	path := randStr(5)
	c.urls[path] = longUrl
	return fmt.Sprintf("%s/%s", c.baseUrl, path)
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	path := strings.TrimPrefix(shortUrl, fmt.Sprintf("%s/", c.baseUrl))
	return c.urls[path]
}

const aplhabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		j := rand.Intn(len(aplhabets))
		char := aplhabets[j]
		res = append(res, char)
	}

	return string(res)
}
