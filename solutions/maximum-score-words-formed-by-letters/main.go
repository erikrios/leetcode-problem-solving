package main

import "fmt"

func main() {
	fmt.Println(maxScoreWords([]string{"dog", "cat", "dad", "good"}, []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}, []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(maxScoreWords([]string{"xxxz", "ax", "bx", "cx"}, []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}, []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}))
	fmt.Println(maxScoreWords([]string{"leetcode"}, []byte{'l', 'e', 't', 'c', 'o', 'd'}, []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}))
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	mapper := make(map[byte]struct {
		count int
		score int
	})

	for _, letter := range letters {
		if val, ok := mapper[letter]; !ok {
			mapper[letter] = struct {
				count int
				score int
			}{1, score[letter-'a']}
		} else {
			val.count += 1
			mapper[letter] = val
		}
	}

	return backtracking(mapper, 0, words, "")
}

func backtracking(
	mapper map[byte]struct {
		count int
		score int
	},
	i int,
	words []string,
	wordComs string,
) int {
	if i >= len(words) {
		checkers := make(map[byte]int)
		for j := 0; j < len(wordComs); j++ {
			checkers[wordComs[j]]++
		}

		var score int

		for k, v := range checkers {
			if mapperVal, ok := mapper[k]; !ok {
				return 0
			} else if v > mapperVal.count {
				return 0
			} else {
				score += v * mapperVal.score
			}
		}

		return score
	}

	// Take
	wordComs += words[i]
	aScore := backtracking(mapper, i+1, words, wordComs)

	// Not take
	wordComs = wordComs[:len(wordComs)-len(words[i])]
	if bScore := backtracking(mapper, i+1, words, wordComs); bScore > aScore {
		return bScore
	}

	return aScore
}
