package memory

type Ram struct {
	ram []byte
}

func NewRam(size int) Ram {
	return Ram{
		ram: make([]byte, size),
	}
}

func (m *Ram) WriteByteToAddress(addr uint16, v byte) {
	m.ram[addr] = v
}

func (m *Ram) ReadByteFromAddress(addr uint16) byte {
	return m.ram[addr]
}
