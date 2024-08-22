package main

import (
	"fmt"
)

func main() {
	fmt.Println(kidsWithCandies([]int{2,3,5,1,3}, 3))
	fmt.Println(kidsWithCandies([]int{4,2,1,1,2}, 1))
	fmt.Println(kidsWithCandies([]int{12,1,12}, 10))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
    m := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > m {
			 m = candies[i]
		}
	}
	
	out := make([]bool, len(candies))
	for i, v := range candies {
		if v + extraCandies >= m {
			out[i] = true
		}
	}
	
	return out
}