package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for {
		if n == 1 {
			return true
		}

		_, haveSeen := seen[n]
		if haveSeen {
			return false
		}

		seen[n] = struct{}{}

		var next int
		for ; n != 0; n /= 10 {
			digit := n % 10
			next += digit * digit
		}
		n = next
	}
}
