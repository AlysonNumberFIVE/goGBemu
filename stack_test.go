
package main

import (
	"testing"
	"fmt"
)

// TestStackPushPC tests the functionality of stackPushPC.
func TestStackPushPC(t *testing.T) {
	testCPU := createContext()

	testCPU.pc = 42
	testCPU.stackPushPC()
	if testCPU.memory[testCPU.sp - 1] != 42 && testCPU.sp != 0xff81 {
		t.Errorf("Stack pushing error:")
		fmt.Println("SP is ", testCPU.sp)
		fmt.Println("Memory is ", testCPU.memory[testCPU.sp - 1])
	}
	testCPU.pc = 99
	testCPU.stackPushPC()
	if testCPU.memory[testCPU.sp - 1] != 99 && testCPU.sp != 0xff82 {
		t.Errorf("Stack pushing error:")
		fmt.Println("SP is ", testCPU.sp)
		fmt.Println("Memory is ", testCPU.memory[testCPU.sp - 1])
	}
	testCPU.stackPopPC()
	testCPU.stackPopPC()

	if testCPU.pc != 42 && testCPU.sp != 0xff80 {
		t.Errorf("Stack popping error:")
	}
	if testCPU.stackPopPC() == true {
		t.Errorf("Stack popping error: true returned for empty stack space.")
	}
}

// TestPusha tests the functionality of the function pusha
func TestPusha(t *testing.T) {
	testCPU := createContext()

	testCPU.a = 1
	testCPU.f = 2
	testCPU.b = 3
	testCPU.c = 4
	testCPU.d = 5
	testCPU.e = 6
	testCPU.h = 7
	testCPU.l = 8

	testCPU.pusha()

	testCPU.a = 0
	testCPU.f = 0
	testCPU.b = 0
	testCPU.c = 0
	testCPU.d = 0
	testCPU.e = 0
	testCPU.h = 0
	testCPU.l = 0

	testCPU.popa()

	if testCPU.a != 1 {
		t.Errorf("reg a != 1, a == %d", testCPU.a)
	}
	if testCPU.f != 2 {
		t.Errorf("reg f != 2, f == %d", testCPU.f)
	}
	if testCPU.b != 3 {
		t.Errorf("reg b != 3, b == %d", testCPU.b)
	}
	if testCPU.c != 4 {
		t.Errorf("reg c != 1, c == %d", testCPU.c)
	}
	if testCPU.d != 5 {
		t.Errorf("reg d != 2, d == %d", testCPU.d)
	}
	if testCPU.e != 6 {
		t.Errorf("reg e != 3, e == %d", testCPU.e)
	}
	if testCPU.h != 7 {
		t.Errorf("reg h != 2, h == %d", testCPU.h)
	}
	if testCPU.l != 8 {
		t.Errorf("reg l != 3, l == %d", testCPU.l)
	}
}
