
package main


func (cpu *cpuContext)verifyIFSwitch(position int) {
}

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

func (cpu *cpuContext)interrupt(sigNum int) {
	cpu._IFSwitchFunction(1, sigNum)

	if cpu.iFlags.IME == 1 && checkBit(cpu.iFlags.IE, sigNum) {
		cpu.stackPush()
	}
}



















