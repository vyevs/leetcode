package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSubsequence("abc","ahbgdc"))
	fmt.Println(isSubsequence("axc","ahbgdc"))
}

func isSubsequence(s, t string) bool {
	if s == "" {
		return true
	}
	var i int
    for j := range t {
		if t[j] == s[i] {
			if i == len(s) - 1 {
				return true
			}
			i++
		}
	}
	
	return false
}