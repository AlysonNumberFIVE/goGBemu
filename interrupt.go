
package main

import (
	"fmt"
)

// _IFSwitchFunction controls the interrupt bits of the IF register.
func (cpu *cpuContext)_IFSwitchFunction(bitVal int, position uint) {
	if bitVal == 1 {
		cpu.memory[cpu.iFlags.IF] = setBit(cpu.memory[cpu.iFlags.IF], position)
	} else if bitVal == 0 {
		cpu.memory[cpu.iFlags.IF] = clearBit(cpu.memory[cpu.iFlags.IF], position)
	}
}

// _IESwitchFunction controls the interrupt bits in the IE register.
func (cpu *cpuContext)_IESwitchFunction(bitVal int, position uint) {
	if bitVal == 1 {
		cpu.memory[cpu.iFlags.IE] = setBit(cpu.memory[cpu.iFlags.IE], position)
	} else if bitVal == 0 {
		cpu.memory[cpu.iFlags.IE] = clearBit(cpu.memory[cpu.iFlags.IE], position)
	}
}

// runInterrupt runs the specified interrupt.
func (cpu *cpuContext)runInterrupt(sigNum uint) {
	if cpu.iFlags.IME == 1 {
		cpu.iFlags.IME = 0						// Disable
		cpu._IFSwitchFunction(0, sigNum)	// Acknowledge
		cpu.stackPushPC()						// push
		cpu.pusha()					
		cpu.pc = cpu.iFlags.iAddresses[sigNum]	

		// run interrupt.
		cpu.popa()
		cpu.stackPopPC()
		cpu.iFlags.IME = 1
	}	
}

// inteerrupt processes all interrupts during a HALT.
func (cpu *cpuContext)interrupt() (bool, int) {
	var interruptChecker uint = 1

	for interruptChecker <= 5 {
		if checkBit(cpu.memory[cpu.iFlags.IF], interruptChecker) != 0 && checkBit(cpu.memory[cpu.iFlags.IE], interruptChecker) != 0 {
			cpu.runInterrupt(interruptChecker - 1)
			return true, int(interruptChecker)
		}
		interruptChecker++
	}
	return false, -1
}
