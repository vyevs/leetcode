package main

import "fmt"

func main() {
	fmt.Println(getLucky("iiii", 1))
	fmt.Println(getLucky("leetcode", 2))
	fmt.Println(getLucky("zbax", 2))
}

func getLucky(s string, k int) int {
	var sum int
	for i := range len(s) {
		c := s[i]

		sum += sumDigits(int(c - 'a' + 1))
	}

	for range k - 1 {
		sum = sumDigits(sum)
	}

	return sum
}

func sumDigits(n int) int {
	var sum int
	for ; n != 0; n /= 10 {
		sum += n % 10
	}
	return sum
}
