
package main

import (

	"testing"
)

// TestClearBit tests the functionality than clearBit.
func TestClearBit(t *testing.T) {
	var number uint8 = 8
	var arg uint = 3
	
	if clearBit(number, arg) != 0 {
		t.Errorf("clearBit inaccurate")
	}
	number = 3
	arg = 0
	if clearBit(number, arg) != 2 {
		t.Errorf("clearBit inaccurate")
	}
	number = 10
	arg = 1
	if clearBit(number, arg) != 8 {
		t.Errorf("clearBit inaccurate")
	}
	number = 11
	arg = 1
	if clearBit(number, arg) != 9 {
		t.Errorf("clearBit inaccurate")
	}
}

// TestSetBit tests the bit of
func TestSetBit(t *testing.T) {
	var number uint8 
	var arg uint

	number = 0
	arg = 3
	if setBit(number, arg) != 8 {
		t.Errorf("setBit inaccurate")
	}
	number = 0
	arg = 1
	if setBit(number, arg) != 2 {
		t.Errorf("setBit inaccurate")
	}
	number = 0
	arg = 2
	if setBit(number, arg) != 4 {
		t.Errorf("setBit inaccurate")
	}
}

// TestCheckBit tests checkBit's functionality.
func TestCheckBit(t *testing.T) {
	var number uint8
	var arg uint

	number = 9
	arg = 0
	if checkBit(number, arg) != 0 {
		t.Errorf("check bit incorrect (1001)")
	}
	number = 8
	arg = 3
	if checkBit(number, arg) != 0 {
		t.Errorf("check bit incorrect (1000)")
	}
	number = 11
	arg = 2
	if checkBit(number, arg) == 0 {
		t.Errorf("check bit incorrect (1011)")
	}

}









