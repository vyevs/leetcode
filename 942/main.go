package main

import (
	"fmt"
)

func main() {
	fmt.Println(diStringMatch("DIIII"))
}


func diStringMatch(s string) []int {
	perm := make([]int, 0, len(s) + 1)
	
	min, max := 0, len(s)
	for _, c := range s {
		if c == 'I' {
			perm = append(perm, min)
			min++
		} else {
			perm = append(perm, max)
			max--
		}
	}
	
	perm = append(perm, max)

	
	return perm
}