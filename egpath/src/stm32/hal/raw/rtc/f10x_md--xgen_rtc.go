// +build f10x_md

package rtc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type RTC_Periph struct {
	CRH  CRH
	_    uint16
	CRL  CRL
	_    uint16
	PRLH PRLH
	_    uint16
	PRLL PRLL
	_    uint16
	DIVH DIVH
	_    uint16
	DIVL DIVL
	_    uint16
	CNTH CNTH
	_    uint16
	CNTL CNTL
	_    uint16
	ALRH ALRH
	_    uint16
	ALRL ALRL
}

func (p *RTC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var RTC = (*RTC_Periph)(unsafe.Pointer(uintptr(mmap.RTC_BASE)))

type CRH_Bits uint16

type CRH struct{ mmio.U16 }

func (r *CRH) Bits(mask CRH_Bits) CRH_Bits { return CRH_Bits(r.U16.Bits(uint16(mask))) }
func (r *CRH) StoreBits(mask, b CRH_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CRH) SetBits(mask CRH_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CRH) ClearBits(mask CRH_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CRH) Load() CRH_Bits              { return CRH_Bits(r.U16.Load()) }
func (r *CRH) Store(b CRH_Bits)            { r.U16.Store(uint16(b)) }

type CRH_Mask struct{ mmio.UM16 }

func (rm CRH_Mask) Load() CRH_Bits   { return CRH_Bits(rm.UM16.Load()) }
func (rm CRH_Mask) Store(b CRH_Bits) { rm.UM16.Store(uint16(b)) }

func (p *RTC_Periph) SECIE() CRH_Mask {
	return CRH_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(SECIE)}}
}

func (p *RTC_Periph) ALRIE() CRH_Mask {
	return CRH_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(ALRIE)}}
}

func (p *RTC_Periph) OWIE() CRH_Mask {
	return CRH_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(OWIE)}}
}

type CRL_Bits uint16

type CRL struct{ mmio.U16 }

func (r *CRL) Bits(mask CRL_Bits) CRL_Bits { return CRL_Bits(r.U16.Bits(uint16(mask))) }
func (r *CRL) StoreBits(mask, b CRL_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CRL) SetBits(mask CRL_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CRL) ClearBits(mask CRL_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CRL) Load() CRL_Bits              { return CRL_Bits(r.U16.Load()) }
func (r *CRL) Store(b CRL_Bits)            { r.U16.Store(uint16(b)) }

type CRL_Mask struct{ mmio.UM16 }

func (rm CRL_Mask) Load() CRL_Bits   { return CRL_Bits(rm.UM16.Load()) }
func (rm CRL_Mask) Store(b CRL_Bits) { rm.UM16.Store(uint16(b)) }

func (p *RTC_Periph) SECF() CRL_Mask {
	return CRL_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(SECF)}}
}

func (p *RTC_Periph) ALRF() CRL_Mask {
	return CRL_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(ALRF)}}
}

func (p *RTC_Periph) OWF() CRL_Mask {
	return CRL_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(OWF)}}
}

func (p *RTC_Periph) RSF() CRL_Mask {
	return CRL_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(RSF)}}
}

func (p *RTC_Periph) CNF() CRL_Mask {
	return CRL_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(CNF)}}
}

func (p *RTC_Periph) RTOFF() CRL_Mask {
	return CRL_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(RTOFF)}}
}

type PRLH_Bits uint16

type PRLH struct{ mmio.U16 }

func (r *PRLH) Bits(mask PRLH_Bits) PRLH_Bits { return PRLH_Bits(r.U16.Bits(uint16(mask))) }
func (r *PRLH) StoreBits(mask, b PRLH_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *PRLH) SetBits(mask PRLH_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *PRLH) ClearBits(mask PRLH_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *PRLH) Load() PRLH_Bits               { return PRLH_Bits(r.U16.Load()) }
func (r *PRLH) Store(b PRLH_Bits)             { r.U16.Store(uint16(b)) }

type PRLH_Mask struct{ mmio.UM16 }

func (rm PRLH_Mask) Load() PRLH_Bits   { return PRLH_Bits(rm.UM16.Load()) }
func (rm PRLH_Mask) Store(b PRLH_Bits) { rm.UM16.Store(uint16(b)) }

type PRLL_Bits uint16

type PRLL struct{ mmio.U16 }

func (r *PRLL) Bits(mask PRLL_Bits) PRLL_Bits { return PRLL_Bits(r.U16.Bits(uint16(mask))) }
func (r *PRLL) StoreBits(mask, b PRLL_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *PRLL) SetBits(mask PRLL_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *PRLL) ClearBits(mask PRLL_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *PRLL) Load() PRLL_Bits               { return PRLL_Bits(r.U16.Load()) }
func (r *PRLL) Store(b PRLL_Bits)             { r.U16.Store(uint16(b)) }

type PRLL_Mask struct{ mmio.UM16 }

func (rm PRLL_Mask) Load() PRLL_Bits   { return PRLL_Bits(rm.UM16.Load()) }
func (rm PRLL_Mask) Store(b PRLL_Bits) { rm.UM16.Store(uint16(b)) }

type DIVH_Bits uint16

type DIVH struct{ mmio.U16 }

func (r *DIVH) Bits(mask DIVH_Bits) DIVH_Bits { return DIVH_Bits(r.U16.Bits(uint16(mask))) }
func (r *DIVH) StoreBits(mask, b DIVH_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DIVH) SetBits(mask DIVH_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DIVH) ClearBits(mask DIVH_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DIVH) Load() DIVH_Bits               { return DIVH_Bits(r.U16.Load()) }
func (r *DIVH) Store(b DIVH_Bits)             { r.U16.Store(uint16(b)) }

type DIVH_Mask struct{ mmio.UM16 }

func (rm DIVH_Mask) Load() DIVH_Bits   { return DIVH_Bits(rm.UM16.Load()) }
func (rm DIVH_Mask) Store(b DIVH_Bits) { rm.UM16.Store(uint16(b)) }

type DIVL_Bits uint16

type DIVL struct{ mmio.U16 }

func (r *DIVL) Bits(mask DIVL_Bits) DIVL_Bits { return DIVL_Bits(r.U16.Bits(uint16(mask))) }
func (r *DIVL) StoreBits(mask, b DIVL_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DIVL) SetBits(mask DIVL_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DIVL) ClearBits(mask DIVL_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DIVL) Load() DIVL_Bits               { return DIVL_Bits(r.U16.Load()) }
func (r *DIVL) Store(b DIVL_Bits)             { r.U16.Store(uint16(b)) }

type DIVL_Mask struct{ mmio.UM16 }

func (rm DIVL_Mask) Load() DIVL_Bits   { return DIVL_Bits(rm.UM16.Load()) }
func (rm DIVL_Mask) Store(b DIVL_Bits) { rm.UM16.Store(uint16(b)) }

type CNTH_Bits uint16

type CNTH struct{ mmio.U16 }

func (r *CNTH) Bits(mask CNTH_Bits) CNTH_Bits { return CNTH_Bits(r.U16.Bits(uint16(mask))) }
func (r *CNTH) StoreBits(mask, b CNTH_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CNTH) SetBits(mask CNTH_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *CNTH) ClearBits(mask CNTH_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *CNTH) Load() CNTH_Bits               { return CNTH_Bits(r.U16.Load()) }
func (r *CNTH) Store(b CNTH_Bits)             { r.U16.Store(uint16(b)) }

type CNTH_Mask struct{ mmio.UM16 }

func (rm CNTH_Mask) Load() CNTH_Bits   { return CNTH_Bits(rm.UM16.Load()) }
func (rm CNTH_Mask) Store(b CNTH_Bits) { rm.UM16.Store(uint16(b)) }

type CNTL_Bits uint16

type CNTL struct{ mmio.U16 }

func (r *CNTL) Bits(mask CNTL_Bits) CNTL_Bits { return CNTL_Bits(r.U16.Bits(uint16(mask))) }
func (r *CNTL) StoreBits(mask, b CNTL_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CNTL) SetBits(mask CNTL_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *CNTL) ClearBits(mask CNTL_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *CNTL) Load() CNTL_Bits               { return CNTL_Bits(r.U16.Load()) }
func (r *CNTL) Store(b CNTL_Bits)             { r.U16.Store(uint16(b)) }

type CNTL_Mask struct{ mmio.UM16 }

func (rm CNTL_Mask) Load() CNTL_Bits   { return CNTL_Bits(rm.UM16.Load()) }
func (rm CNTL_Mask) Store(b CNTL_Bits) { rm.UM16.Store(uint16(b)) }

type ALRH_Bits uint16

type ALRH struct{ mmio.U16 }

func (r *ALRH) Bits(mask ALRH_Bits) ALRH_Bits { return ALRH_Bits(r.U16.Bits(uint16(mask))) }
func (r *ALRH) StoreBits(mask, b ALRH_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *ALRH) SetBits(mask ALRH_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *ALRH) ClearBits(mask ALRH_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *ALRH) Load() ALRH_Bits               { return ALRH_Bits(r.U16.Load()) }
func (r *ALRH) Store(b ALRH_Bits)             { r.U16.Store(uint16(b)) }

type ALRH_Mask struct{ mmio.UM16 }

func (rm ALRH_Mask) Load() ALRH_Bits   { return ALRH_Bits(rm.UM16.Load()) }
func (rm ALRH_Mask) Store(b ALRH_Bits) { rm.UM16.Store(uint16(b)) }

type ALRL_Bits uint16

type ALRL struct{ mmio.U16 }

func (r *ALRL) Bits(mask ALRL_Bits) ALRL_Bits { return ALRL_Bits(r.U16.Bits(uint16(mask))) }
func (r *ALRL) StoreBits(mask, b ALRL_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *ALRL) SetBits(mask ALRL_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *ALRL) ClearBits(mask ALRL_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *ALRL) Load() ALRL_Bits               { return ALRL_Bits(r.U16.Load()) }
func (r *ALRL) Store(b ALRL_Bits)             { r.U16.Store(uint16(b)) }

type ALRL_Mask struct{ mmio.UM16 }

func (rm ALRL_Mask) Load() ALRL_Bits   { return ALRL_Bits(rm.UM16.Load()) }
func (rm ALRL_Mask) Store(b ALRL_Bits) { rm.UM16.Store(uint16(b)) }
