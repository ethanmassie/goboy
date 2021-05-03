package cpu

import (
	"encoding/binary"
	"unsafe"

	"emassie.dev/go-boy/internal/memory"
)

// dependent on the endianness of the architecture
var encoding binary.ByteOrder

// detect endianess and set encoding accordingly so we can use the binary package
// see https://stackoverflow.com/questions/51332658/any-better-way-to-check-endianness-in-go
func setEncoding() {
	var i int = 0x0100
	var ptr = unsafe.Pointer(&i)
	if *(*byte)(ptr) == 0x01 {
		encoding = binary.BigEndian
	} else if *(*byte)(ptr) == 0x00 {
		encoding = binary.LittleEndian
	}
}

var CpuRegisters Registers

func initializeRegisters() {
	CpuRegisters = Registers{
		af:    make([]byte, 2),
		bc:    make([]byte, 2),
		de:    make([]byte, 2),
		hl:    make([]byte, 2),
		flags: 0x00,
		sp:    0x00,
		pc:    0x00,
	}
}

var CpuRam memory.Ram

func initializeMemory() {
	CpuRam = memory.NewRam(8192)
}

func Init() {
	setEncoding()
	initializeRegisters()
	initializeMemory()
}
