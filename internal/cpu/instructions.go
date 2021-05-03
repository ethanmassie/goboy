package cpu

type Instruction struct {
	Disassembly   string
	OperandLength byte
	Execute       func(ops ...byte)
}

var CpuInstructions = [256]Instruction{
	0x00: {
		Disassembly:   "NOP",
		OperandLength: 0,
		Execute: func(ops ...byte) {
			// nooperation
		},
	},
	0x01: {
		Disassembly:   "LD BC NN",
		OperandLength: 2,
		Execute: func(ops ...byte) {
			// load operands into reg bc
			CpuRegisters.PutBC(ops)
		},
	},
	0x02: {
		Disassembly:   "LD BC A",
		OperandLength: 0,
		Execute: func(ops ...byte) {
			bc := CpuRegisters.BC()
			a := CpuRegisters.A()
			CpuRam.WriteByteToAddress(bc, a)
		},
	},
	0x03: {
		Disassembly:   "INC BC",
		OperandLength: 0,
		Execute: func(ops ...byte) {
			// increment bc by 1
			addBytes(CpuRegisters.bc, 1)
		},
	},
}
