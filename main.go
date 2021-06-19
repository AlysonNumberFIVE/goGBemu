
package main


// interrupt starting addresses 
const VBankingStart uint16 = 0x0040 			// vertical banking      - priority 1
const LCDCStatusIntStart uint16 = 0x0048		// LCDC status interrupt - priority 2
const timerOverflowStart uint16 = 0x0050		// Timer overflow        - priority 3
const serialTransCompStart uint16 = 0x0058   // Serial transfer comp  - priority 4
const inputSignalLowStart uint16 = 0x0060    // P10-P13 input sig low - priority 5


// Bit positions for each of the interrupt bits in the
// IF and IE registers.
const verticalBanking uint = 0
const LCDC uint = 1
const timerOverflow uint = 2
const serialIOTransferCompletion uint = 3
const terminalNegativeEdge uint = 4

// interrupt flag locations.
type interruptFlags struct {
	IF 				uint16 // 0xff0f
	IE 				uint16 // 0xffff
	IME 			uint16 // ??? - seemingsly has no address and is a 1 or 0.
	iAddresses		map[uint]uint16
}

// LCD display registers.
type _LCDDisplayRegs struct {
	LCDC	uint16 // 0xff40
	STAT 	uint16 // 0xff41
	SCY		uint16 // 0xff42
	SCX		uint16 // 0xff43
	LY 		uint16 // 0xff44
	LYC 	uint16 // 0xff45
	DMA  	uint16 // 0xff46
	BGP		uint16 // 0xff47
	OBP0	uint16 // 0xff48
	OBP1	uint16 // 0xff49
	WY 		uint16 // 0xff4a
	WX 		uint16 // 0xff4b
	OAM		[] uint16 // 0xfe00+/-0xfe9f
}

// sound register locations.
type soundRegisters struct {
	NRxx	[] uint16 // 0xff10, 0xff26
	WavRAM	[] uint16 // 0xff30, 0xff3f
}

// bit position of values in P1 register
// to be passed into setBit.
const _P1p10 int = 0
const _P1p11 int = 1
const _P1p12 int = 2
const _P1p13 int = 3
const _P1p14 int = 4
const _P1p15 int = 5

// DIV register bit values. Bit positions to send to
// setBit
const _8192Hz int = 0
const _4096Hz int = 1
const _2048Hz int = 2
const _1024Hz int = 3
const _512Hz int = 4
const _256Hz int = 5
const _126Hz int = 6
const _64Hz int = 7

// bit positions for TAC register. Values to be passed into
// bitSet.
//   values for inputClockSelect/TAC bytes are:
//  	00 = (4.096 KHz)
//		01 = (262.144 kHz)
//		10 = (65.536 KHz)
//		11 = (16.384 KHz)

const inputClockSelect11 int = 11
const inputClockSelect10 int = 10
const inputClockSelect01 int = 1
const timeStop int = 0
const timeStart int = 1
const timeStopStart int = 2

// port mode register locations.
type portModeRegisters struct {
	P1 		uint16 // 0xff00
	SB 		uint16 // 0xff01
	SC 		uint16 // 0xff02
	DIV 	uint16 // 0xff04
	TIMA	uint16 // 0xff05
	TMA		uint16 // 0xff06
	TAC		uint16 // 0xff07
}

// array index for memory start and end values in 
// memory struct.
const memStart int = 0
const memEnd int = 1

// memory
type memory struct {
	memSpace 		[] uint16  // entire memory space 
	rst 			[] uint16  // 0x0000, 0x00ff
	romData 		[] uint16  // 0x0100, 0x014f
	startAddress    [] uint16  // 0x0150, 0x7fff
	LCDRAM 			[] uint16  // 0x8000, 0x9fff
	externExp 		[] uint16  // 0xa000, 0xbfff
	workRAM 		[] uint16  // 0xc000, 0Xdfff
	forbidden 		[] uint16  // 0xe000, 0xfdff
	cpuRAM			[] uint16  // 0xfe00, 0xffff
	displayData 	[] uint16  // 0xfe00, 0xfe9f
	regFlagSpace	[] uint16  // 0xff00, 0xff7f
	workStackRAM	[] uint16  // 0xff80, 0xfffe
}

type cpuContext struct {
	a 				uint8
	f 				uint8
	b 				uint8
	c 				uint8
	d 				uint8
	e 				uint8
	h 				uint8
	l 				uint8
	sp 				uint16
	pc 				uint16
	memory 			[]uint8

	portModeReg 	portModeRegisters
	soundReg 		soundRegisters
	lcdReg 			_LCDDisplayRegs
	iFlags 			interruptFlags

}

func (cpu * cpuContext) readMemory(location uint16) uint8 {
	return cpu.memory[location]
}

func (cpu * cpuContext) writeMemory(location uint16, data uint8) {
	cpu.memory[location] = data
}


/*
func (cpu * cpuContext) loadRom(romFile string) []byte {
	rom, err := ioutil.ReadFile(romFile)
	if err != nil {
		fmt.Println("Error : unable to read rom file")
		os.Exit(1)
	}

}		
*/
func main() {
}



