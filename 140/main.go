package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Printf("%q\n", wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmt.Printf("%q\n", wordBreak("catsandog", []string{"cats","dog","sand","and","cat"}))
}

func wordBreak(s string, words []string) []string {
	
	
	out := make([]string, 0, 16)
	
	return wordBreakRec(s, words, "", out)
}

func wordBreakRec(s string, words []string, cur string, out []string) []string {
	if len(s) == 0 {
		out = append(out, cur)
		return out
	}
		
	for _, w := range words {
		if strings.HasPrefix(s, w) {
			newCur := cur + w
			newS := s[len(w):]
			
			if newS != "" {
				newCur += " "
			}
			
			out = wordBreakRec(newS, words, newCur, out)
		}
	}
	
	return out
} 