package main

import (
	"fmt"
)

func main() {
	fmt.Println(reformat("abc123"))
	fmt.Println(reformat("123abcd"))
	fmt.Println(reformat("1234abc"))
	fmt.Println(reformat("leetcode"))
}

// a1b2c3d
func reformat(s string) string {
    var letters, digits []byte
    for _, c := range s {
        if c >= '0' && c <= '9' {
            digits = append(digits, byte(c))
        } else {
            letters = append(letters, byte(c))
        }
    }

    diff := abs(len(letters)-len(digits))
    if diff > 1 {
        return ""
    }

    out := merge(letters, digits)
    return string(out)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func merge(s1, s2 []byte) []byte {
	n := len(s1) + len(s2)
	
    out := make([]byte, 0, len(s1)+len(s2))

	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	
	for i := 0; i < n; i++ {
		out = append(out, s1[0])
		s1 = s1[1:]
		s1, s2 = s2, s1
	}
	
	return out
}