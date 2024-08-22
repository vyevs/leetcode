package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("aA"))
}

func reverseVowels(s string) string {
	isVowel := func(b byte) bool {
		switch b {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U': return true
		}
		return false
	}
	bs := []byte(s)
	
	for i, j := 0, len(bs)-1; i < j; {
		if !isVowel(bs[i]) {
			i++
			continue
		}
		if !isVowel(bs[j]) {
			j--
			continue
		}
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	
	return string(bs)
}