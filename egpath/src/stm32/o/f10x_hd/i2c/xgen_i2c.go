package i2c

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type I2C_Periph struct {
	CR1   RCR1
	_     uint16
	CR2   RCR2
	_     uint16
	OAR1  ROAR1
	_     uint16
	OAR2  ROAR2
	_     uint16
	DR    RDR
	_     uint16
	SR1   RSR1
	_     uint16
	SR2   RSR2
	_     uint16
	CCR   RCCR
	_     uint16
	TRISE RTRISE
}

func (p *I2C_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var I2C1 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C1_BASE)))

//emgo:const
var I2C2 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C2_BASE)))

type CR1 uint16

func (b CR1) Field(mask CR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1) J(v int) CR1 {
	return CR1(bits.Make32(v, uint32(mask)))
}

type RCR1 struct{ mmio.U16 }

func (r *RCR1) Bits(mask CR1) CR1     { return CR1(r.U16.Bits(uint16(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U16.SetBits(uint16(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U16.Load()) }
func (r *RCR1) Store(b CR1)           { r.U16.Store(uint16(b)) }

type RMCR1 struct{ mmio.UM16 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM16.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) PE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(PE)}}
}

func (p *I2C_Periph) SMBUS() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(SMBUS)}}
}

func (p *I2C_Periph) SMBTYPE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(SMBTYPE)}}
}

func (p *I2C_Periph) ENARP() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(ENARP)}}
}

func (p *I2C_Periph) ENPEC() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(ENPEC)}}
}

func (p *I2C_Periph) ENGC() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(ENGC)}}
}

func (p *I2C_Periph) NOSTRETCH() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(NOSTRETCH)}}
}

func (p *I2C_Periph) START() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(START)}}
}

func (p *I2C_Periph) STOP() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(STOP)}}
}

func (p *I2C_Periph) ACK() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(ACK)}}
}

func (p *I2C_Periph) POS() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(POS)}}
}

func (p *I2C_Periph) PEC() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(PEC)}}
}

func (p *I2C_Periph) ALERT() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(ALERT)}}
}

func (p *I2C_Periph) SWRST() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(SWRST)}}
}

type CR2 uint16

func (b CR2) Field(mask CR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2) J(v int) CR2 {
	return CR2(bits.Make32(v, uint32(mask)))
}

type RCR2 struct{ mmio.U16 }

func (r *RCR2) Bits(mask CR2) CR2     { return CR2(r.U16.Bits(uint16(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U16.SetBits(uint16(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U16.Load()) }
func (r *RCR2) Store(b CR2)           { r.U16.Store(uint16(b)) }

type RMCR2 struct{ mmio.UM16 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM16.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) FREQ() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(FREQ)}}
}

func (p *I2C_Periph) ITERREN() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(ITERREN)}}
}

func (p *I2C_Periph) ITEVTEN() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(ITEVTEN)}}
}

func (p *I2C_Periph) ITBUFEN() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(ITBUFEN)}}
}

func (p *I2C_Periph) DMAEN() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(DMAEN)}}
}

func (p *I2C_Periph) LAST() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(LAST)}}
}

type OAR1 uint16

func (b OAR1) Field(mask OAR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OAR1) J(v int) OAR1 {
	return OAR1(bits.Make32(v, uint32(mask)))
}

type ROAR1 struct{ mmio.U16 }

func (r *ROAR1) Bits(mask OAR1) OAR1    { return OAR1(r.U16.Bits(uint16(mask))) }
func (r *ROAR1) StoreBits(mask, b OAR1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *ROAR1) SetBits(mask OAR1)      { r.U16.SetBits(uint16(mask)) }
func (r *ROAR1) ClearBits(mask OAR1)    { r.U16.ClearBits(uint16(mask)) }
func (r *ROAR1) Load() OAR1             { return OAR1(r.U16.Load()) }
func (r *ROAR1) Store(b OAR1)           { r.U16.Store(uint16(b)) }

type RMOAR1 struct{ mmio.UM16 }

func (rm RMOAR1) Load() OAR1   { return OAR1(rm.UM16.Load()) }
func (rm RMOAR1) Store(b OAR1) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) ADD1_7() RMOAR1 {
	return RMOAR1{mmio.UM16{&p.OAR1.U16, uint16(ADD1_7)}}
}

func (p *I2C_Periph) ADD8_9() RMOAR1 {
	return RMOAR1{mmio.UM16{&p.OAR1.U16, uint16(ADD8_9)}}
}

func (p *I2C_Periph) ADD0() RMOAR1 {
	return RMOAR1{mmio.UM16{&p.OAR1.U16, uint16(ADD0)}}
}

func (p *I2C_Periph) ADDMODE() RMOAR1 {
	return RMOAR1{mmio.UM16{&p.OAR1.U16, uint16(ADDMODE)}}
}

type OAR2 uint16

func (b OAR2) Field(mask OAR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OAR2) J(v int) OAR2 {
	return OAR2(bits.Make32(v, uint32(mask)))
}

type ROAR2 struct{ mmio.U16 }

func (r *ROAR2) Bits(mask OAR2) OAR2    { return OAR2(r.U16.Bits(uint16(mask))) }
func (r *ROAR2) StoreBits(mask, b OAR2) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *ROAR2) SetBits(mask OAR2)      { r.U16.SetBits(uint16(mask)) }
func (r *ROAR2) ClearBits(mask OAR2)    { r.U16.ClearBits(uint16(mask)) }
func (r *ROAR2) Load() OAR2             { return OAR2(r.U16.Load()) }
func (r *ROAR2) Store(b OAR2)           { r.U16.Store(uint16(b)) }

type RMOAR2 struct{ mmio.UM16 }

func (rm RMOAR2) Load() OAR2   { return OAR2(rm.UM16.Load()) }
func (rm RMOAR2) Store(b OAR2) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) ENDUAL() RMOAR2 {
	return RMOAR2{mmio.UM16{&p.OAR2.U16, uint16(ENDUAL)}}
}

func (p *I2C_Periph) SECADD1_7() RMOAR2 {
	return RMOAR2{mmio.UM16{&p.OAR2.U16, uint16(SECADD1_7)}}
}

type DR uint16

func (b DR) Field(mask DR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR) J(v int) DR {
	return DR(bits.Make32(v, uint32(mask)))
}

type RDR struct{ mmio.U16 }

func (r *RDR) Bits(mask DR) DR      { return DR(r.U16.Bits(uint16(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDR) SetBits(mask DR)      { r.U16.SetBits(uint16(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDR) Load() DR             { return DR(r.U16.Load()) }
func (r *RDR) Store(b DR)           { r.U16.Store(uint16(b)) }

type RMDR struct{ mmio.UM16 }

func (rm RMDR) Load() DR   { return DR(rm.UM16.Load()) }
func (rm RMDR) Store(b DR) { rm.UM16.Store(uint16(b)) }

type SR1 uint16

func (b SR1) Field(mask SR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR1) J(v int) SR1 {
	return SR1(bits.Make32(v, uint32(mask)))
}

type RSR1 struct{ mmio.U16 }

func (r *RSR1) Bits(mask SR1) SR1     { return SR1(r.U16.Bits(uint16(mask))) }
func (r *RSR1) StoreBits(mask, b SR1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RSR1) SetBits(mask SR1)      { r.U16.SetBits(uint16(mask)) }
func (r *RSR1) ClearBits(mask SR1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RSR1) Load() SR1             { return SR1(r.U16.Load()) }
func (r *RSR1) Store(b SR1)           { r.U16.Store(uint16(b)) }

type RMSR1 struct{ mmio.UM16 }

func (rm RMSR1) Load() SR1   { return SR1(rm.UM16.Load()) }
func (rm RMSR1) Store(b SR1) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) SB() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(SB)}}
}

func (p *I2C_Periph) ADDR() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(ADDR)}}
}

func (p *I2C_Periph) BTF() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(BTF)}}
}

func (p *I2C_Periph) ADD10() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(ADD10)}}
}

func (p *I2C_Periph) STOPF() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(STOPF)}}
}

func (p *I2C_Periph) RXNE() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(RXNE)}}
}

func (p *I2C_Periph) TXE() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(TXE)}}
}

func (p *I2C_Periph) BERR() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(BERR)}}
}

func (p *I2C_Periph) ARLO() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(ARLO)}}
}

func (p *I2C_Periph) AF() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(AF)}}
}

func (p *I2C_Periph) OVR() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(OVR)}}
}

func (p *I2C_Periph) PECERR() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(PECERR)}}
}

func (p *I2C_Periph) TIMEOUT() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(TIMEOUT)}}
}

func (p *I2C_Periph) SMBALERT() RMSR1 {
	return RMSR1{mmio.UM16{&p.SR1.U16, uint16(SMBALERT)}}
}

type SR2 uint16

func (b SR2) Field(mask SR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR2) J(v int) SR2 {
	return SR2(bits.Make32(v, uint32(mask)))
}

type RSR2 struct{ mmio.U16 }

func (r *RSR2) Bits(mask SR2) SR2     { return SR2(r.U16.Bits(uint16(mask))) }
func (r *RSR2) StoreBits(mask, b SR2) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RSR2) SetBits(mask SR2)      { r.U16.SetBits(uint16(mask)) }
func (r *RSR2) ClearBits(mask SR2)    { r.U16.ClearBits(uint16(mask)) }
func (r *RSR2) Load() SR2             { return SR2(r.U16.Load()) }
func (r *RSR2) Store(b SR2)           { r.U16.Store(uint16(b)) }

type RMSR2 struct{ mmio.UM16 }

func (rm RMSR2) Load() SR2   { return SR2(rm.UM16.Load()) }
func (rm RMSR2) Store(b SR2) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) MSL() RMSR2 {
	return RMSR2{mmio.UM16{&p.SR2.U16, uint16(MSL)}}
}

func (p *I2C_Periph) BUSY() RMSR2 {
	return RMSR2{mmio.UM16{&p.SR2.U16, uint16(BUSY)}}
}

func (p *I2C_Periph) TRA() RMSR2 {
	return RMSR2{mmio.UM16{&p.SR2.U16, uint16(TRA)}}
}

func (p *I2C_Periph) GENCALL() RMSR2 {
	return RMSR2{mmio.UM16{&p.SR2.U16, uint16(GENCALL)}}
}

func (p *I2C_Periph) SMBDEFAULT() RMSR2 {
	return RMSR2{mmio.UM16{&p.SR2.U16, uint16(SMBDEFAULT)}}
}

func (p *I2C_Periph) SMBHOST() RMSR2 {
	return RMSR2{mmio.UM16{&p.SR2.U16, uint16(SMBHOST)}}
}

func (p *I2C_Periph) DUALF() RMSR2 {
	return RMSR2{mmio.UM16{&p.SR2.U16, uint16(DUALF)}}
}

func (p *I2C_Periph) PECVAL() RMSR2 {
	return RMSR2{mmio.UM16{&p.SR2.U16, uint16(PECVAL)}}
}

type CCR uint16

func (b CCR) Field(mask CCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR) J(v int) CCR {
	return CCR(bits.Make32(v, uint32(mask)))
}

type RCCR struct{ mmio.U16 }

func (r *RCCR) Bits(mask CCR) CCR     { return CCR(r.U16.Bits(uint16(mask))) }
func (r *RCCR) StoreBits(mask, b CCR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCCR) SetBits(mask CCR)      { r.U16.SetBits(uint16(mask)) }
func (r *RCCR) ClearBits(mask CCR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCCR) Load() CCR             { return CCR(r.U16.Load()) }
func (r *RCCR) Store(b CCR)           { r.U16.Store(uint16(b)) }

type RMCCR struct{ mmio.UM16 }

func (rm RMCCR) Load() CCR   { return CCR(rm.UM16.Load()) }
func (rm RMCCR) Store(b CCR) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) CCRVAL() RMCCR {
	return RMCCR{mmio.UM16{&p.CCR.U16, uint16(CCRVAL)}}
}

func (p *I2C_Periph) DUTY() RMCCR {
	return RMCCR{mmio.UM16{&p.CCR.U16, uint16(DUTY)}}
}

func (p *I2C_Periph) FS() RMCCR {
	return RMCCR{mmio.UM16{&p.CCR.U16, uint16(FS)}}
}

type TRISE uint16

func (b TRISE) Field(mask TRISE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TRISE) J(v int) TRISE {
	return TRISE(bits.Make32(v, uint32(mask)))
}

type RTRISE struct{ mmio.U16 }

func (r *RTRISE) Bits(mask TRISE) TRISE   { return TRISE(r.U16.Bits(uint16(mask))) }
func (r *RTRISE) StoreBits(mask, b TRISE) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RTRISE) SetBits(mask TRISE)      { r.U16.SetBits(uint16(mask)) }
func (r *RTRISE) ClearBits(mask TRISE)    { r.U16.ClearBits(uint16(mask)) }
func (r *RTRISE) Load() TRISE             { return TRISE(r.U16.Load()) }
func (r *RTRISE) Store(b TRISE)           { r.U16.Store(uint16(b)) }

type RMTRISE struct{ mmio.UM16 }

func (rm RMTRISE) Load() TRISE   { return TRISE(rm.UM16.Load()) }
func (rm RMTRISE) Store(b TRISE) { rm.UM16.Store(uint16(b)) }
