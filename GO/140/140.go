package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordsList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}
func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, wordDict))
	s = "pineapplepenapple"
	wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}
	fmt.Println(wordBreak(s, wordDict))
	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}
