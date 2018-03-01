// +build l476xx

package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type ADC_Common_Periph struct {
	CSR RCSR
	_   uint32
	CCR RCCR
	CDR RCDR
}

func (p *ADC_Common_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var ADC123_COMMON = (*ADC_Common_Periph)(unsafe.Pointer(uintptr(mmap.ADC123_COMMON_BASE)))

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

func (p *ADC_Common_Periph) ADRDY_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(ADRDY_MST)}}
}

func (p *ADC_Common_Periph) EOSMP_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOSMP_MST)}}
}

func (p *ADC_Common_Periph) EOC_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOC_MST)}}
}

func (p *ADC_Common_Periph) EOS_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOS_MST)}}
}

func (p *ADC_Common_Periph) OVR_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OVR_MST)}}
}

func (p *ADC_Common_Periph) JEOC_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JEOC_MST)}}
}

func (p *ADC_Common_Periph) JEOS_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JEOS_MST)}}
}

func (p *ADC_Common_Periph) AWD1_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD1_MST)}}
}

func (p *ADC_Common_Periph) AWD2_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD2_MST)}}
}

func (p *ADC_Common_Periph) AWD3_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD3_MST)}}
}

func (p *ADC_Common_Periph) JQOVF_MST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JQOVF_MST)}}
}

func (p *ADC_Common_Periph) ADRDY_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(ADRDY_SLV)}}
}

func (p *ADC_Common_Periph) EOSMP_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOSMP_SLV)}}
}

func (p *ADC_Common_Periph) EOC_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOC_SLV)}}
}

func (p *ADC_Common_Periph) EOS_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EOS_SLV)}}
}

func (p *ADC_Common_Periph) OVR_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OVR_SLV)}}
}

func (p *ADC_Common_Periph) JEOC_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JEOC_SLV)}}
}

func (p *ADC_Common_Periph) JEOS_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JEOS_SLV)}}
}

func (p *ADC_Common_Periph) AWD1_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD1_SLV)}}
}

func (p *ADC_Common_Periph) AWD2_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD2_SLV)}}
}

func (p *ADC_Common_Periph) AWD3_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(AWD3_SLV)}}
}

func (p *ADC_Common_Periph) JQOVF_SLV() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(JQOVF_SLV)}}
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

func (p *ADC_Common_Periph) DUAL() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DUAL)}}
}

func (p *ADC_Common_Periph) DELAY() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DELAY)}}
}

func (p *ADC_Common_Periph) MDMACFG() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(MDMACFG)}}
}

func (p *ADC_Common_Periph) MDMA() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(MDMA)}}
}

func (p *ADC_Common_Periph) CKMODE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(CKMODE)}}
}

func (p *ADC_Common_Periph) PRESC() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(PRESC)}}
}

func (p *ADC_Common_Periph) VREFEN() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(VREFEN)}}
}

func (p *ADC_Common_Periph) TSEN() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(TSEN)}}
}

func (p *ADC_Common_Periph) VBATEN() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(VBATEN)}}
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

func (p *ADC_Common_Periph) RDATA_MST() RMCDR {
	return RMCDR{mmio.UM32{&p.CDR.U32, uint32(RDATA_MST)}}
}

func (p *ADC_Common_Periph) RDATA_SLV() RMCDR {
	return RMCDR{mmio.UM32{&p.CDR.U32, uint32(RDATA_SLV)}}
}
