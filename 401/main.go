package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%q\n", readBinaryWatch(1))
	fmt.Printf("%q\n", readBinaryWatch(2))
	fmt.Printf("%q\n", readBinaryWatch(3))
	fmt.Printf("%q\n", readBinaryWatch(9))
}


func readBinaryWatch(turnedOn int) []string {
	out := make([]string, 0, 16)
	
	var h, m uint16
	for h = 0; h < 12; h++ {
		for m = 0; m < 60; m++ {
			if bits.OnesCount16(h) + bits.OnesCount16(m) == turnedOn {
				s := fmt.Sprintf("%d:", h)
				if m < 10 {
					s += "0"
				}
				s += fmt.Sprintf("%d", m)
				out = append(out, s)
			}
		}
	}
	
	return out
}