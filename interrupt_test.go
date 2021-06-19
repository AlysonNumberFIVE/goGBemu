
package main

import (
	"testing"
	"fmt"
)

// TestIFSwitchFunction tests the functonality of
// _IFSwitchFunction.
func TestIFSwitchFunction(t *testing.T) {
	testCPU := createContext()

	testCPU._IFSwitchFunction(1, verticalBanking)
	if testCPU.memory[testCPU.iFlags.IF] != 1 {
		t.Errorf("memory value doesn't equal 1 (0000 0001)")
	}

	testCPU._IFSwitchFunction(1, timerOverflow)
	if testCPU.memory[testCPU.iFlags.IF] != 5 {
		t.Errorf("memory value doesn't equal 5 (0000 0101)")
	}

	testCPU._IFSwitchFunction(1, serialIOTransferCompletion)
	if testCPU.memory[testCPU.iFlags.IF] != 13 {
		t.Errorf("memory value doesn't equal 13 (0000 1101")
	}

	testCPU._IFSwitchFunction(1, serialIOTransferCompletion)
	if testCPU.memory[testCPU.iFlags.IF] != 13 {
		t.Errorf("memory value doesn't equal 13 (0000 1101)")
	}

	testCPU._IFSwitchFunction(0, serialIOTransferCompletion)
	if testCPU.memory[testCPU.iFlags.IF] != 5 {
		t.Errorf("memory value doesn't equal 5 (0000 0101")
	}
	fmt.Println(testCPU.memory[testCPU.iFlags.IF])
}

// TestIESwitchFunction tests the functionality of 
// _IESwitchFunction
func TestIESwitchFunction(t *testing.T) {
	testCPU := createContext()

	testCPU._IESwitchFunction(1, verticalBanking)
	if testCPU.memory[testCPU.iFlags.IE] != 1 {
		t.Errorf("memory value doesn't equal 1 (0000 0001)")
	}

	testCPU._IESwitchFunction(1, timerOverflow)
	if testCPU.memory[testCPU.iFlags.IE] != 5 {
		t.Errorf("memory value doesn't equal 5 (0000 0101)")
	}

	testCPU._IESwitchFunction(1, serialIOTransferCompletion)
	if testCPU.memory[testCPU.iFlags.IE] != 13 {
		t.Errorf("memory value doesn't equal 13 (0000 1101")
	}

	testCPU._IESwitchFunction(1, serialIOTransferCompletion)
	if testCPU.memory[testCPU.iFlags.IE] != 13 {
		t.Errorf("memory value doesn't equal 13 (0000 1101)")
	}

	testCPU._IESwitchFunction(0, serialIOTransferCompletion)
	if testCPU.memory[testCPU.iFlags.IE] != 5 {
		t.Errorf("memory value doesn't equal 5 (0000 0101")
	}
	fmt.Println(testCPU.memory[testCPU.iFlags.IF])	
}

// TestInterrupt tests the functionality of interrupt.
func TestInterrupt(t *testing.T) {
	testCPU := createContext()

	testCPU.iFlags.IME = 1
	 
}
















