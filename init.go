
package main

func initMemory() *memory {
	return &memory{
		memSpace:		make([]uint16, 0xffff),
		rst: 		 	[]uint16{0x0000, 0x00ff},
		romData: 	 	[]uint16{0x0100, 0x014f},
		startAddress:	[]uint16{0x0150, 0x7fff},
		LCDRAM:			[]uint16{0x8000, 0x9fff},
		externExp:		[]uint16{0xa000, 0xbfff},
		workRAM:		[]uint16{0xc000, 0Xdfff},
		forbidden:		[]uint16{0xe000, 0xfdff},
		cpuRAM:			[]uint16{0xfe00, 0xffff},
		displayData:	[]uint16{0xfe00, 0xfe9f},
		regFlagSpace:	[]uint16{0xff00, 0xff7f},
		workStackRAM:	[]uint16{0xff80, 0xfffe},
	}
}

func createContext() *cpuContext {
	var portModeReg	portModeRegisters
	var lcdReg		_LCDDisplayRegs
	var iFlags 		interruptFlags

	portModeReg.P1 = 0xff00
	portModeReg.SC = 0xff02
	portModeReg.DIV = 0xff04
	portModeReg.TIMA = 0xff05 
	portModeReg.TMA = 0xff06
	portModeReg.TAC = 0xff07

	lcdReg.SCY = 0xff42
	lcdReg.SCX = 0xff43
	lcdReg.LYC = 0xff45
	lcdReg.WY = 0xff4a
	lcdReg.WX = 0xff4b

	iFlags.IME = 0
	iFlags.IF = 0xff0f 
	iFlags.IE = 0xffff

	iFlags.iAddresses = make(map[uint]uint16)
	iFlags.iAddresses[1] = 0x0040
	iFlags.iAddresses[2] = 0x0048
	iFlags.iAddresses[3] = 0x0050
	iFlags.iAddresses[4] = 0x0058
	iFlags.iAddresses[5] = 0x0060

	return &cpuContext{
		a: 0,
		f: 0,
		b: 0,
		c: 0,
		d: 0,
		e: 0,
		h: 0,
		l: 0,
		sp: 0xff80,
		pc: 0,
		memory: make([]uint8, 0xffff + 1),
		portModeReg: portModeReg,
		lcdReg: lcdReg,
		iFlags: iFlags,
	}
}













