package main

import (
	"fmt"
)

func wordBreak(s string, wordDic []string) []string {
	wordMap := map[string]bool{}
	for _, expect := range wordDic {
		wordMap[expect] = true
	}

	mem := map[int][]string{}

	return findWords(s, 0, mem, wordMap)
}

func findWords(s string, start int, mem map[int][]string, wordMap map[string]bool) []string {
	if val, ok := mem[start]; ok {
		return val
	}

	if start == len(s) {
		return []string{""}
	}

	slice := []string{}
	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]
		if wordMap[word] {
			ws := findWords(s, end, mem, wordMap)
			for _, w := range ws {
				if w != "" {
					slice = append(slice, word+" "+w)
				} else {
					slice = append(slice, word)
				}
			}
		}
	}

	mem[start] = slice
	return slice
}

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	/*
		Expect:
		[
			"cat sand dog",
			"cats and dog"
		]
	*/

	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	/*
		Expect:
		[
		    "pine apple pen apple",
		    "pineapple pen apple",
		    "pine applepen apple"
		]
	*/
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	/*
		Expect:
		[]
	*/
}
