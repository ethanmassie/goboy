package cpu

func put2Bytes(buff []byte, read []byte) {
	buff[0] = read[0]
	buff[1] = read[1]
}

func addBytes(buff []byte, x uint16) {
	v := encoding.Uint16(buff)
	v += x
	encoding.PutUint16(buff, v)
}

type Registers struct {
	af    []byte
	bc    []byte
	de    []byte
	hl    []byte
	flags byte
	sp    uint16
	pc    uint16
}

func (r *Registers) A() byte {
	return r.af[0]
}

func (r *Registers) PutA(a byte) {
	r.af[0] = a
}

func (r *Registers) F() byte {
	return r.af[1]
}

func (r *Registers) PutF(f byte) {
	r.af[1] = f
}

func (r *Registers) AF() uint16 {
	return encoding.Uint16(r.af)
}

func (r *Registers) PutAF(af []byte) {
	put2Bytes(r.af, af)
}

func (r *Registers) B() byte {
	return r.bc[0]
}

func (r *Registers) PutB(b byte) {
	r.bc[0] = b
}

func (r *Registers) C() byte {
	return r.bc[1]
}

func (r *Registers) PutC(c byte) {
	r.bc[1] = c
}

func (r *Registers) BC() uint16 {
	return encoding.Uint16(r.bc)
}

func (r *Registers) PutBC(bc []byte) {
	put2Bytes(r.bc, bc)
}

func (r *Registers) D() byte {
	return r.de[0]
}

func (r *Registers) PutD(d byte) {
	r.de[0] = d
}

func (r *Registers) E() byte {
	return r.de[1]
}

func (r *Registers) PutE(e byte) {
	r.de[1] = e
}

func (r *Registers) DE() uint16 {
	return encoding.Uint16(r.de)
}

func (r *Registers) PutDE(de []byte) {
	put2Bytes(r.de, de)
}

func (r *Registers) H() byte {
	return r.hl[0]
}

func (r *Registers) PutH(h byte) {
	r.hl[0] = h
}

func (r *Registers) L() byte {
	return r.hl[1]
}

func (r *Registers) PutL(l byte) {
	r.hl[1] = l
}

func (r *Registers) HL() uint16 {
	return encoding.Uint16(r.hl)
}

func (r *Registers) PutHL(hl []byte) {
	put2Bytes(r.hl, hl)
}

func (r *Registers) Flags() byte {
	return r.flags
}

func (r *Registers) PutFlags(flags byte) {
	r.flags = flags
}

func (r *Registers) SP() uint16 {
	return r.sp
}

func (r *Registers) PutSP(sp uint16) {
	r.sp = sp
}

func (r *Registers) PC() uint16 {
	return r.pc
}

func (r *Registers) PutPC(pc uint16) {
	r.pc = pc
}
