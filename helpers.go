
package main

import (
	"fmt"
)


// check if a bit is set.
func checkBit(n uint8, pos uint) uint8 {
	return n & (1 << (pos - 1)) 
}
 
// clearBit clears a single bit from a value at the
// specified position.
func clearBit(n uint8, pos uint) uint8 {
	mask := ^(1 << pos)
	n &= uint8(mask)
	return n
}

// printBit dumps the value passed in in binary format.
func printBit(value uint8) {
	var binVal string

	for value > 0 {
		if value % 2 == 0 {
			binVal += "0"
		} else {
			binVal += "1"
		}
		value /= 2
	}
	fmt.Println(binVal)
}

// setBit sets a single bit in the value at the specified position.
func setBit(n uint8, pos uint) uint8 {
	n |= (1 << pos)
	return n
}
