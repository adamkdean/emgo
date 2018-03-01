package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type ADC_Common_Periph struct {
	CSR RCSR
	CCR RCCR
	CDR RCDR
}

func (p *ADC_Common_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var ADC = (*ADC_Common_Periph)(unsafe.Pointer(uintptr(mmap.ADC_BASE)))

type CSR uint32

func (b CSR) Field(mask CSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR) J(v int) CSR {
	return CSR(bits.Make32(v, uint32(mask)))
}

type RCSR struct{ mmio.U32 }

func (r *RCSR) Bits(mask CSR) CSR     { return CSR(r.U32.Bits(uint32(mask))) }
func (r *RCSR) StoreBits(mask, b CSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) SetBits(mask CSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCSR) ClearBits(mask CSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCSR) Load() CSR             { return CSR(r.U32.Load()) }
func (r *RCSR) Store(b CSR)           { r.U32.Store(uint32(b)) }

func (r *RCSR) AtomicStoreBits(mask, b CSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) AtomicSetBits(mask CSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCSR) AtomicClearBits(mask CSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCSR struct{ mmio.UM32 }

func (rm RMCSR) Load() CSR   { return CSR(rm.UM32.Load()) }
func (rm RMCSR) Store(b CSR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) AWD1() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD1)}}
}

func (p *ADC_Common_Periph) EOC1() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOC1)}}
}

func (p *ADC_Common_Periph) JEOC1() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JEOC1)}}
}

func (p *ADC_Common_Periph) JSTRT1() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JSTRT1)}}
}

func (p *ADC_Common_Periph) STRT1() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(STRT1)}}
}

func (p *ADC_Common_Periph) OVR1() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OVR1)}}
}

func (p *ADC_Common_Periph) AWD2() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD2)}}
}

func (p *ADC_Common_Periph) EOC2() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOC2)}}
}

func (p *ADC_Common_Periph) JEOC2() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JEOC2)}}
}

func (p *ADC_Common_Periph) JSTRT2() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JSTRT2)}}
}

func (p *ADC_Common_Periph) STRT2() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(STRT2)}}
}

func (p *ADC_Common_Periph) OVR2() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OVR2)}}
}

func (p *ADC_Common_Periph) AWD3() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD3)}}
}

func (p *ADC_Common_Periph) EOC3() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOC3)}}
}

func (p *ADC_Common_Periph) JEOC3() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JEOC3)}}
}

func (p *ADC_Common_Periph) JSTRT3() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JSTRT3)}}
}

func (p *ADC_Common_Periph) STRT3() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(STRT3)}}
}

func (p *ADC_Common_Periph) OVR3() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OVR3)}}
}

type CCR uint32

func (b CCR) Field(mask CCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR) J(v int) CCR {
	return CCR(bits.Make32(v, uint32(mask)))
}

type RCCR struct{ mmio.U32 }

func (r *RCCR) Bits(mask CCR) CCR     { return CCR(r.U32.Bits(uint32(mask))) }
func (r *RCCR) StoreBits(mask, b CCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCCR) SetBits(mask CCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCCR) ClearBits(mask CCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCCR) Load() CCR             { return CCR(r.U32.Load()) }
func (r *RCCR) Store(b CCR)           { r.U32.Store(uint32(b)) }

func (r *RCCR) AtomicStoreBits(mask, b CCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCCR) AtomicSetBits(mask CCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCCR) AtomicClearBits(mask CCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCCR struct{ mmio.UM32 }

func (rm RMCCR) Load() CCR   { return CCR(rm.UM32.Load()) }
func (rm RMCCR) Store(b CCR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) MULTI() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(MULTI)}}
}

func (p *ADC_Common_Periph) DELAY() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DELAY)}}
}

func (p *ADC_Common_Periph) DDS() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DDS)}}
}

func (p *ADC_Common_Periph) DMA() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DMA)}}
}

func (p *ADC_Common_Periph) ADCPRE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(ADCPRE)}}
}

func (p *ADC_Common_Periph) VBATE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(VBATE)}}
}

func (p *ADC_Common_Periph) TSVREFE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(TSVREFE)}}
}

type CDR uint32

func (b CDR) Field(mask CDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CDR) J(v int) CDR {
	return CDR(bits.Make32(v, uint32(mask)))
}

type RCDR struct{ mmio.U32 }

func (r *RCDR) Bits(mask CDR) CDR     { return CDR(r.U32.Bits(uint32(mask))) }
func (r *RCDR) StoreBits(mask, b CDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCDR) SetBits(mask CDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCDR) ClearBits(mask CDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCDR) Load() CDR             { return CDR(r.U32.Load()) }
func (r *RCDR) Store(b CDR)           { r.U32.Store(uint32(b)) }

func (r *RCDR) AtomicStoreBits(mask, b CDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCDR) AtomicSetBits(mask CDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCDR) AtomicClearBits(mask CDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCDR struct{ mmio.UM32 }

func (rm RMCDR) Load() CDR   { return CDR(rm.UM32.Load()) }
func (rm RMCDR) Store(b CDR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) DATA1() RMCDR {
	return RMCDR{mmio.UM32{&p.CDR.U32, uint32(DATA1)}}
}

func (p *ADC_Common_Periph) DATA2() RMCDR {
	return RMCDR{mmio.UM32{&p.CDR.U32, uint32(DATA2)}}
}
