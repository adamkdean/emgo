// +build f10x_md

package rtc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type RTC_Periph struct {
	CRH  RCRH
	_    uint16
	CRL  RCRL
	_    uint16
	PRLH RPRLH
	_    uint16
	PRLL RPRLL
	_    uint16
	DIVH RDIVH
	_    uint16
	DIVL RDIVL
	_    uint16
	CNTH RCNTH
	_    uint16
	CNTL RCNTL
	_    uint16
	ALRH RALRH
	_    uint16
	ALRL RALRL
}

func (p *RTC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var RTC = (*RTC_Periph)(unsafe.Pointer(uintptr(mmap.RTC_BASE)))

type CRH uint16

func (b CRH) Field(mask CRH) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CRH) J(v int) CRH {
	return CRH(bits.Make32(v, uint32(mask)))
}

type RCRH struct{ mmio.U16 }

func (r *RCRH) Bits(mask CRH) CRH     { return CRH(r.U16.Bits(uint16(mask))) }
func (r *RCRH) StoreBits(mask, b CRH) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCRH) SetBits(mask CRH)      { r.U16.SetBits(uint16(mask)) }
func (r *RCRH) ClearBits(mask CRH)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCRH) Load() CRH             { return CRH(r.U16.Load()) }
func (r *RCRH) Store(b CRH)           { r.U16.Store(uint16(b)) }

type RMCRH struct{ mmio.UM16 }

func (rm RMCRH) Load() CRH   { return CRH(rm.UM16.Load()) }
func (rm RMCRH) Store(b CRH) { rm.UM16.Store(uint16(b)) }

func (p *RTC_Periph) SECIE() RMCRH {
	return RMCRH{mmio.UM16{&p.CRH.U16, uint16(SECIE)}}
}

func (p *RTC_Periph) ALRIE() RMCRH {
	return RMCRH{mmio.UM16{&p.CRH.U16, uint16(ALRIE)}}
}

func (p *RTC_Periph) OWIE() RMCRH {
	return RMCRH{mmio.UM16{&p.CRH.U16, uint16(OWIE)}}
}

type CRL uint16

func (b CRL) Field(mask CRL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CRL) J(v int) CRL {
	return CRL(bits.Make32(v, uint32(mask)))
}

type RCRL struct{ mmio.U16 }

func (r *RCRL) Bits(mask CRL) CRL     { return CRL(r.U16.Bits(uint16(mask))) }
func (r *RCRL) StoreBits(mask, b CRL) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCRL) SetBits(mask CRL)      { r.U16.SetBits(uint16(mask)) }
func (r *RCRL) ClearBits(mask CRL)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCRL) Load() CRL             { return CRL(r.U16.Load()) }
func (r *RCRL) Store(b CRL)           { r.U16.Store(uint16(b)) }

type RMCRL struct{ mmio.UM16 }

func (rm RMCRL) Load() CRL   { return CRL(rm.UM16.Load()) }
func (rm RMCRL) Store(b CRL) { rm.UM16.Store(uint16(b)) }

func (p *RTC_Periph) SECF() RMCRL {
	return RMCRL{mmio.UM16{&p.CRL.U16, uint16(SECF)}}
}

func (p *RTC_Periph) ALRF() RMCRL {
	return RMCRL{mmio.UM16{&p.CRL.U16, uint16(ALRF)}}
}

func (p *RTC_Periph) OWF() RMCRL {
	return RMCRL{mmio.UM16{&p.CRL.U16, uint16(OWF)}}
}

func (p *RTC_Periph) RSF() RMCRL {
	return RMCRL{mmio.UM16{&p.CRL.U16, uint16(RSF)}}
}

func (p *RTC_Periph) CNF() RMCRL {
	return RMCRL{mmio.UM16{&p.CRL.U16, uint16(CNF)}}
}

func (p *RTC_Periph) RTOFF() RMCRL {
	return RMCRL{mmio.UM16{&p.CRL.U16, uint16(RTOFF)}}
}

type PRLH uint16

func (b PRLH) Field(mask PRLH) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PRLH) J(v int) PRLH {
	return PRLH(bits.Make32(v, uint32(mask)))
}

type RPRLH struct{ mmio.U16 }

func (r *RPRLH) Bits(mask PRLH) PRLH    { return PRLH(r.U16.Bits(uint16(mask))) }
func (r *RPRLH) StoreBits(mask, b PRLH) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RPRLH) SetBits(mask PRLH)      { r.U16.SetBits(uint16(mask)) }
func (r *RPRLH) ClearBits(mask PRLH)    { r.U16.ClearBits(uint16(mask)) }
func (r *RPRLH) Load() PRLH             { return PRLH(r.U16.Load()) }
func (r *RPRLH) Store(b PRLH)           { r.U16.Store(uint16(b)) }

type RMPRLH struct{ mmio.UM16 }

func (rm RMPRLH) Load() PRLH   { return PRLH(rm.UM16.Load()) }
func (rm RMPRLH) Store(b PRLH) { rm.UM16.Store(uint16(b)) }

type PRLL uint16

func (b PRLL) Field(mask PRLL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PRLL) J(v int) PRLL {
	return PRLL(bits.Make32(v, uint32(mask)))
}

type RPRLL struct{ mmio.U16 }

func (r *RPRLL) Bits(mask PRLL) PRLL    { return PRLL(r.U16.Bits(uint16(mask))) }
func (r *RPRLL) StoreBits(mask, b PRLL) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RPRLL) SetBits(mask PRLL)      { r.U16.SetBits(uint16(mask)) }
func (r *RPRLL) ClearBits(mask PRLL)    { r.U16.ClearBits(uint16(mask)) }
func (r *RPRLL) Load() PRLL             { return PRLL(r.U16.Load()) }
func (r *RPRLL) Store(b PRLL)           { r.U16.Store(uint16(b)) }

type RMPRLL struct{ mmio.UM16 }

func (rm RMPRLL) Load() PRLL   { return PRLL(rm.UM16.Load()) }
func (rm RMPRLL) Store(b PRLL) { rm.UM16.Store(uint16(b)) }

type DIVH uint16

func (b DIVH) Field(mask DIVH) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIVH) J(v int) DIVH {
	return DIVH(bits.Make32(v, uint32(mask)))
}

type RDIVH struct{ mmio.U16 }

func (r *RDIVH) Bits(mask DIVH) DIVH    { return DIVH(r.U16.Bits(uint16(mask))) }
func (r *RDIVH) StoreBits(mask, b DIVH) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDIVH) SetBits(mask DIVH)      { r.U16.SetBits(uint16(mask)) }
func (r *RDIVH) ClearBits(mask DIVH)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDIVH) Load() DIVH             { return DIVH(r.U16.Load()) }
func (r *RDIVH) Store(b DIVH)           { r.U16.Store(uint16(b)) }

type RMDIVH struct{ mmio.UM16 }

func (rm RMDIVH) Load() DIVH   { return DIVH(rm.UM16.Load()) }
func (rm RMDIVH) Store(b DIVH) { rm.UM16.Store(uint16(b)) }

type DIVL uint16

func (b DIVL) Field(mask DIVL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIVL) J(v int) DIVL {
	return DIVL(bits.Make32(v, uint32(mask)))
}

type RDIVL struct{ mmio.U16 }

func (r *RDIVL) Bits(mask DIVL) DIVL    { return DIVL(r.U16.Bits(uint16(mask))) }
func (r *RDIVL) StoreBits(mask, b DIVL) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDIVL) SetBits(mask DIVL)      { r.U16.SetBits(uint16(mask)) }
func (r *RDIVL) ClearBits(mask DIVL)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDIVL) Load() DIVL             { return DIVL(r.U16.Load()) }
func (r *RDIVL) Store(b DIVL)           { r.U16.Store(uint16(b)) }

type RMDIVL struct{ mmio.UM16 }

func (rm RMDIVL) Load() DIVL   { return DIVL(rm.UM16.Load()) }
func (rm RMDIVL) Store(b DIVL) { rm.UM16.Store(uint16(b)) }

type CNTH uint16

func (b CNTH) Field(mask CNTH) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CNTH) J(v int) CNTH {
	return CNTH(bits.Make32(v, uint32(mask)))
}

type RCNTH struct{ mmio.U16 }

func (r *RCNTH) Bits(mask CNTH) CNTH    { return CNTH(r.U16.Bits(uint16(mask))) }
func (r *RCNTH) StoreBits(mask, b CNTH) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCNTH) SetBits(mask CNTH)      { r.U16.SetBits(uint16(mask)) }
func (r *RCNTH) ClearBits(mask CNTH)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCNTH) Load() CNTH             { return CNTH(r.U16.Load()) }
func (r *RCNTH) Store(b CNTH)           { r.U16.Store(uint16(b)) }

type RMCNTH struct{ mmio.UM16 }

func (rm RMCNTH) Load() CNTH   { return CNTH(rm.UM16.Load()) }
func (rm RMCNTH) Store(b CNTH) { rm.UM16.Store(uint16(b)) }

type CNTL uint16

func (b CNTL) Field(mask CNTL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CNTL) J(v int) CNTL {
	return CNTL(bits.Make32(v, uint32(mask)))
}

type RCNTL struct{ mmio.U16 }

func (r *RCNTL) Bits(mask CNTL) CNTL    { return CNTL(r.U16.Bits(uint16(mask))) }
func (r *RCNTL) StoreBits(mask, b CNTL) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCNTL) SetBits(mask CNTL)      { r.U16.SetBits(uint16(mask)) }
func (r *RCNTL) ClearBits(mask CNTL)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCNTL) Load() CNTL             { return CNTL(r.U16.Load()) }
func (r *RCNTL) Store(b CNTL)           { r.U16.Store(uint16(b)) }

type RMCNTL struct{ mmio.UM16 }

func (rm RMCNTL) Load() CNTL   { return CNTL(rm.UM16.Load()) }
func (rm RMCNTL) Store(b CNTL) { rm.UM16.Store(uint16(b)) }

type ALRH uint16

func (b ALRH) Field(mask ALRH) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ALRH) J(v int) ALRH {
	return ALRH(bits.Make32(v, uint32(mask)))
}

type RALRH struct{ mmio.U16 }

func (r *RALRH) Bits(mask ALRH) ALRH    { return ALRH(r.U16.Bits(uint16(mask))) }
func (r *RALRH) StoreBits(mask, b ALRH) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RALRH) SetBits(mask ALRH)      { r.U16.SetBits(uint16(mask)) }
func (r *RALRH) ClearBits(mask ALRH)    { r.U16.ClearBits(uint16(mask)) }
func (r *RALRH) Load() ALRH             { return ALRH(r.U16.Load()) }
func (r *RALRH) Store(b ALRH)           { r.U16.Store(uint16(b)) }

type RMALRH struct{ mmio.UM16 }

func (rm RMALRH) Load() ALRH   { return ALRH(rm.UM16.Load()) }
func (rm RMALRH) Store(b ALRH) { rm.UM16.Store(uint16(b)) }

type ALRL uint16

func (b ALRL) Field(mask ALRL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ALRL) J(v int) ALRL {
	return ALRL(bits.Make32(v, uint32(mask)))
}

type RALRL struct{ mmio.U16 }

func (r *RALRL) Bits(mask ALRL) ALRL    { return ALRL(r.U16.Bits(uint16(mask))) }
func (r *RALRL) StoreBits(mask, b ALRL) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RALRL) SetBits(mask ALRL)      { r.U16.SetBits(uint16(mask)) }
func (r *RALRL) ClearBits(mask ALRL)    { r.U16.ClearBits(uint16(mask)) }
func (r *RALRL) Load() ALRL             { return ALRL(r.U16.Load()) }
func (r *RALRL) Store(b ALRL)           { r.U16.Store(uint16(b)) }

type RMALRL struct{ mmio.UM16 }

func (rm RMALRL) Load() ALRL   { return ALRL(rm.UM16.Load()) }
func (rm RMALRL) Store(b ALRL) { rm.UM16.Store(uint16(b)) }
