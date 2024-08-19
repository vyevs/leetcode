package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(s1, s2 string) string {
	var out strings.Builder
	
	for len(s1) > 0 && len(s2) > 0 {
		_ = out.WriteByte(s1[0])
		
		s1 = s1[1:]
		
		s1, s2 = s2, s1
	}
	
	_, _ = out.WriteString(s1)
	_, _ = out.WriteString(s2)
	
	return out.String()
}