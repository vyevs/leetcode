package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
    var candidates int
	
	for i, v := range flowerbed {
		if v == 0 {
			prev := i-1
			next := i+1
			if (prev < 0 || flowerbed[prev] == 0) && (next >= len(flowerbed) || flowerbed[next] == 0) {
				candidates++
				flowerbed[i] = 1
			}
		}
	}
	
	return candidates >= n
}