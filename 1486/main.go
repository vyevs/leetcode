package main

import (
	"fmt"
)

func main() {
	fmt.Println(xorOperation(5, 0))
	fmt.Println(xorOperation(4, 3))
}

func xorOperation(n int, start int) int {
    var out int
    for i := 0; i < n; i++ {
        out ^= start + 2 * i
    }
    return out
}