
package main


func (cpu *cpuContext)stackPushPC() bool {
	if cpu.sp == 0xfffe {
		return false
	}
	cpu.memory[cpu.sp] = uint8(cpu.pc)
	cpu.sp++
	return true
}

func (cpu *cpuContext)stackPopPC() bool {
	if cpu.sp == 0xff80 {
		return false
	}
	cpu.pc = uint16(cpu.memory[cpu.sp])
	cpu.sp--
	return true
}

func (cpu *cpuContext)pusha() {
	cpu.memory[cpu.sp] = cpu.a 
	cpu.sp++
	cpu.memory[cpu.sp] = cpu.f
	cpu.sp++
	cpu.memory[cpu.sp] = cpu.b 
	cpu.sp++ 
	cpu.memory[cpu.sp] = cpu.c 
	cpu.sp++
	cpu.memory[cpu.sp] = cpu.d 
	cpu.sp++ 
	cpu.memory[cpu.sp] = cpu.e 
	cpu.sp++
	cpu.memory[cpu.sp] = cpu.h 
	cpu.sp++
	cpu.memory[cpu.sp] = cpu.l 
	cpu.sp++
}

func (cpu *cpuContext)popa() {
	cpu.sp--
	cpu.l = cpu.memory[cpu.sp] 
	cpu.sp--
	cpu.h = cpu.memory[cpu.sp]
	cpu.sp--
	cpu.e = cpu.memory[cpu.sp]
	cpu.sp--
	cpu.d = cpu.memory[cpu.sp]
	cpu.sp--
	cpu.c = cpu.memory[cpu.sp]
	cpu.sp--
	cpu.b = cpu.memory[cpu.sp]
	cpu.sp--
	cpu.f = cpu.memory[cpu.sp]
	cpu.sp--
	cpu.a = cpu.memory[cpu.sp]
}

