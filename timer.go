
package main

import (
	"fmt"
	"os"
)

// startTimer sets the Time Stop flag in the TAC register.
func (cpu * cpuContext)startTimer(value uint8, timerBit int) uint8 {
	if timerBit == timeStart {
		value = setBit(value, 2)
	} else if timerBit == timeStop {
		value = clearBit(value, 2)
	} else {
		fmt.Println("Error : setFrequency start/stop bit value incorrect")
		os.Exit(1)
	}
	return value
}

// setFrequency sets the frequency/timer bits in the TAC register.
func (cpu * cpuContext )setFrequency(frequencyBits int, timerBit int) {
	var value uint8 = 0

	switch (frequencyBits) {
	case inputClockSelect11:
		value = setBit(value, 0)
		value = setBit(value, 1)
	case inputClockSelect10:
		value = setBit(value, 1)
	case inputClockSelect01:
		value = setBit(value, 0)
	}
	value = cpu.startTimer(value, timerBit) 
	cpu.memory[cpu.portModeReg.TAC] = value
}
