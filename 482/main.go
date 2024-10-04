package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 3))
	fmt.Println(licenseKeyFormatting("aaaa", 2))
	fmt.Println(licenseKeyFormatting("--a-a-a-a--", 2))

}

func licenseKeyFormatting(s string, k int) string {

	out := make([]byte, 0, len(s))

	var acc int
	for i := len(s) - 1; i >= 0; i-- {

		c := s[i]

		if c == '-' {
			continue
		}

		if acc == k {
			out = append(out, '-')
			acc = 0
		}

		out = append(out, toUpper(c))
		acc++
	}

	slices.Reverse(out)

	return string(out)
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}
