package main

import (
	"fmt"
)

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
	fmt.Println(lemonadeChange([]int{5,5,10,20,5,5,5,5,5,5,5,5,5,10,5,5,20,5,20,5}))
}


func lemonadeChange(bills []int) bool {
    var fives, tens int
	
	for _, v := range bills {
		if v == 5 {
			fives++
		} else if v == 10 {
			tens++
		}
		
		chg := v - 5
		
		if chg >= 10 && tens > 0 {
			chg -= 10
			tens--
		}
		
		for chg != 0 && fives != 0 {
			chg -= 5
			fives--
		}
	
		if chg > 0 {
			return false
		}
	}
	
	
	return true
}