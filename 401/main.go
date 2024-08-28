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
	
	var i uint16
	// We can start with an initial value that has turnedOn bits on right away.
	// This way we don't have to check all the smaller values that have fewer bits on.
	// This is useful for higher bit counts.
	for j := 0; j < turnedOn; j++ {
		i <<= 1
		i |= 1
	}
	for ; i < 1024; i++ {
		ct := bits.OnesCount16(i)
		
		if ct != turnedOn {
			continue
		}
	
		lower6 := i & 0x3f
		upper4 := (i & 0x3c0) >> 6
		if lower6 > 59 || upper4 > 11 {
			continue
		}
		
		s := fmt.Sprintf("%d:", upper4)
		if lower6 < 10 {
			s += "0"
		}
		s += fmt.Sprintf("%d", lower6)
		out = append(out, s)
		
	}
	
	return out
}