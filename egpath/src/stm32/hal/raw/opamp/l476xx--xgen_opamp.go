// +build l476xx

package opamp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type OPAMP_Periph struct {
	CSR   RCSR
	OTR   ROTR
	LPOTR RLPOTR
}

func (p *OPAMP_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var OPAMP = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP_BASE)))

//emgo:const
var OPAMP1 = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP1_BASE)))

//emgo:const
var OPAMP2 = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP2_BASE)))

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

func (p *OPAMP_Periph) OPAMPxEN() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OPAMPxEN)}}
}

func (p *OPAMP_Periph) OPALPM() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OPALPM)}}
}

func (p *OPAMP_Periph) OPAMODE() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OPAMODE)}}
}

func (p *OPAMP_Periph) PGGAIN() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PGGAIN)}}
}

func (p *OPAMP_Periph) VMSEL() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(VMSEL)}}
}

func (p *OPAMP_Periph) VPSEL() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(VPSEL)}}
}

func (p *OPAMP_Periph) CALON() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(CALON)}}
}

func (p *OPAMP_Periph) CALSEL() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(CALSEL)}}
}

func (p *OPAMP_Periph) USERTRIM() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(USERTRIM)}}
}

func (p *OPAMP_Periph) CALOUT() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(CALOUT)}}
}

type OTR uint32

func (b OTR) Field(mask OTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OTR) J(v int) OTR {
	return OTR(bits.Make32(v, uint32(mask)))
}

type ROTR struct{ mmio.U32 }

func (r *ROTR) Bits(mask OTR) OTR     { return OTR(r.U32.Bits(uint32(mask))) }
func (r *ROTR) StoreBits(mask, b OTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROTR) SetBits(mask OTR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROTR) ClearBits(mask OTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROTR) Load() OTR             { return OTR(r.U32.Load()) }
func (r *ROTR) Store(b OTR)           { r.U32.Store(uint32(b)) }

func (r *ROTR) AtomicStoreBits(mask, b OTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROTR) AtomicSetBits(mask OTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROTR) AtomicClearBits(mask OTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOTR struct{ mmio.UM32 }

func (rm RMOTR) Load() OTR   { return OTR(rm.UM32.Load()) }
func (rm RMOTR) Store(b OTR) { rm.UM32.Store(uint32(b)) }

func (p *OPAMP_Periph) TRIMOFFSETN() RMOTR {
	return RMOTR{mmio.UM32{&p.OTR.U32, uint32(TRIMOFFSETN)}}
}

func (p *OPAMP_Periph) TRIMOFFSETP() RMOTR {
	return RMOTR{mmio.UM32{&p.OTR.U32, uint32(TRIMOFFSETP)}}
}

type LPOTR uint32

func (b LPOTR) Field(mask LPOTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LPOTR) J(v int) LPOTR {
	return LPOTR(bits.Make32(v, uint32(mask)))
}

type RLPOTR struct{ mmio.U32 }

func (r *RLPOTR) Bits(mask LPOTR) LPOTR   { return LPOTR(r.U32.Bits(uint32(mask))) }
func (r *RLPOTR) StoreBits(mask, b LPOTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLPOTR) SetBits(mask LPOTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLPOTR) ClearBits(mask LPOTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLPOTR) Load() LPOTR             { return LPOTR(r.U32.Load()) }
func (r *RLPOTR) Store(b LPOTR)           { r.U32.Store(uint32(b)) }

func (r *RLPOTR) AtomicStoreBits(mask, b LPOTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLPOTR) AtomicSetBits(mask LPOTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLPOTR) AtomicClearBits(mask LPOTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLPOTR struct{ mmio.UM32 }

func (rm RMLPOTR) Load() LPOTR   { return LPOTR(rm.UM32.Load()) }
func (rm RMLPOTR) Store(b LPOTR) { rm.UM32.Store(uint32(b)) }

func (p *OPAMP_Periph) TRIMLPOFFSETN() RMLPOTR {
	return RMLPOTR{mmio.UM32{&p.LPOTR.U32, uint32(TRIMLPOFFSETN)}}
}

func (p *OPAMP_Periph) TRIMLPOFFSETP() RMLPOTR {
	return RMLPOTR{mmio.UM32{&p.LPOTR.U32, uint32(TRIMLPOFFSETP)}}
}
