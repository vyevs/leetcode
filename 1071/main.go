package main

import (
	"fmt"
	
	"github.com/vyevs/gcd"
)

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
}

func gcdOfStrings(s1, s2 string) string {
	if s1 + s2 != s2 + s1 {
		return ""
	}
	
	return s1[:gcd.GCD(len(s1), len(s2))]
}