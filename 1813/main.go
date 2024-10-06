package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	fmt.Println(areSentencesSimilar("Hello Jane", "Hello my name is Jane"))
	fmt.Println(areSentencesSimilar("of", "A lot of words"))
	fmt.Println(areSentencesSimilar("Eating right now", "Eating"))
	fmt.Println(areSentencesSimilar("Frog cool", "Frogs are cool"))

	fmt.Println(areSentencesSimilar("A A AAa", "A AAa"))
}

func areSentencesSimilar(s1, s2 string) bool {
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}

	s1Words := splitWords(s1)
	s2Words := splitWords(s2)

	var prefixMatchLen int
	for prefixMatchLen < len(s2Words) && s1Words[prefixMatchLen] == s2Words[prefixMatchLen] {
		prefixMatchLen++
	}

	// If s2 is a prefix of s1, then we can append to s2 to make s1, so the sentences are similar.
	if prefixMatchLen == len(s2Words) {
		return true
	}

	var suffixMatchLen int
	for suffixMatchLen < len(s2Words) &&
		prefixMatchLen+suffixMatchLen < len(s2Words) &&
		s1Words[len(s1Words)-1-suffixMatchLen] == s2Words[len(s2Words)-1-suffixMatchLen] {
		suffixMatchLen++
	}

	// If the sum of the lengths of the longest common suffix and prefix are equal to the length of s2, then we can add to s2 to make it s1.
	return prefixMatchLen+suffixMatchLen == len(s2Words)
}

func splitWords(s string) []string {
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanWords)

	out := make([]string, 0, 5)
	for sc.Scan() {
		t := sc.Text()
		out = append(out, t)
	}

	return out
}
