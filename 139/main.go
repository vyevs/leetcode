package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet","code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple","pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats","dog","sand","and","cat"}))
	fmt.Println(wordBreak("applepenapple", []string{"app", "apple", "pen"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}))
}

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	
	for _, w := range wordDict {
		if strings.HasPrefix(s, w) {
			newS := s[len(w):]
			
			pathWorks := wordBreak(newS, wordDict)
			if pathWorks {
				return true
			}
		}
	}
	
	return false
}