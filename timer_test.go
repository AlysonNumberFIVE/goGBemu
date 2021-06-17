
package main

import (
	"testing"
)

// TestStartTimer the startTimer functionality in the TAC register.
func TestStartTimer(t *testing.T) {
	testCPU := createContext()

	testCPU.setFrequency(11, 1)
	if testCPU.memory[testCPU.portModeReg.TAC] != 7 {
		t.Errorf("memory position doesn't equal 7 (0000 0111)")
	}

	testCPU.setFrequency(10, 1)
	if testCPU.memory[testCPU.portModeReg.TAC] != 6 {
		t.Errorf("memory position doesn't equal 6 (0000 0110")
	}	

	testCPU.setFrequency(00, 1)
	if testCPU.memory[testCPU.portModeReg.TAC] != 4 {
		t.Errorf("memory position doesn't equal (0000 0100)")
	}

	testCPU.setFrequency(10, 0)
	if testCPU.memory[testCPU.portModeReg.TAC] != 2 {
		t.Errorf("memory position doesn't equal (0000 0010)")
	}
	printBit(testCPU.memory[testCPU.portModeReg.TAC])
}
