package main

import (
	"fmt"
)

func main() {
	fmt.Println(kthFactor(12, 3))
}

func kthFactor(n int, k int) int {
    var divs int
    for i := 1; i <= n; i++ {
        if n % i == 0 {
            divs++
        }
        if divs == k {
            return i
        }
    }


    return -1
}