package cpu

type Memory struct {
	ram []byte
}

func (m *Memory) WriteByteToAddress(addr uint16, v byte) {
	m.ram[addr] = v
}

func (m *Memory) ReadByteFromAddress(addr uint16) byte {
	return m.ram[addr]
}
