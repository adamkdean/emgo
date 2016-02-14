package cec

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type CEC_Periph struct {
	CFGR CFGR
	OAR  OAR
	PRES PRES
	ESR  ESR
	CSR  CSR
	TXD  TXD
	RXD  RXD
}

func (p *CEC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var CEC = (*CEC_Periph)(unsafe.Pointer(uintptr(mmap.CEC_BASE)))

type CFGR_Bits uint32

type CFGR struct{ mmio.U32 }

func (r *CFGR) Bits(mask CFGR_Bits) CFGR_Bits { return CFGR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFGR) StoreBits(mask, b CFGR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR) SetBits(mask CFGR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CFGR) ClearBits(mask CFGR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CFGR) Load() CFGR_Bits               { return CFGR_Bits(r.U32.Load()) }
func (r *CFGR) Store(b CFGR_Bits)             { r.U32.Store(uint32(b)) }

type CFGR_Mask struct{ mmio.UM32 }

func (rm CFGR_Mask) Load() CFGR_Bits   { return CFGR_Bits(rm.UM32.Load()) }
func (rm CFGR_Mask) Store(b CFGR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CEC_Periph) PE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(PE)}}
}

func (p *CEC_Periph) IE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IE)}}
}

func (p *CEC_Periph) BTEM() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(BTEM)}}
}

func (p *CEC_Periph) BPEM() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(BPEM)}}
}

type OAR_Bits uint32

type OAR struct{ mmio.U32 }

func (r *OAR) Bits(mask OAR_Bits) OAR_Bits { return OAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OAR) StoreBits(mask, b OAR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OAR) SetBits(mask OAR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *OAR) ClearBits(mask OAR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *OAR) Load() OAR_Bits              { return OAR_Bits(r.U32.Load()) }
func (r *OAR) Store(b OAR_Bits)            { r.U32.Store(uint32(b)) }

type OAR_Mask struct{ mmio.UM32 }

func (rm OAR_Mask) Load() OAR_Bits   { return OAR_Bits(rm.UM32.Load()) }
func (rm OAR_Mask) Store(b OAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CEC_Periph) OA() OAR_Mask {
	return OAR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OA)}}
}

type PRES_Bits uint32

type PRES struct{ mmio.U32 }

func (r *PRES) Bits(mask PRES_Bits) PRES_Bits { return PRES_Bits(r.U32.Bits(uint32(mask))) }
func (r *PRES) StoreBits(mask, b PRES_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PRES) SetBits(mask PRES_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PRES) ClearBits(mask PRES_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PRES) Load() PRES_Bits               { return PRES_Bits(r.U32.Load()) }
func (r *PRES) Store(b PRES_Bits)             { r.U32.Store(uint32(b)) }

type PRES_Mask struct{ mmio.UM32 }

func (rm PRES_Mask) Load() PRES_Bits   { return PRES_Bits(rm.UM32.Load()) }
func (rm PRES_Mask) Store(b PRES_Bits) { rm.UM32.Store(uint32(b)) }

type ESR_Bits uint32

type ESR struct{ mmio.U32 }

func (r *ESR) Bits(mask ESR_Bits) ESR_Bits { return ESR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ESR) StoreBits(mask, b ESR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ESR) SetBits(mask ESR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ESR) ClearBits(mask ESR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ESR) Load() ESR_Bits              { return ESR_Bits(r.U32.Load()) }
func (r *ESR) Store(b ESR_Bits)            { r.U32.Store(uint32(b)) }

type ESR_Mask struct{ mmio.UM32 }

func (rm ESR_Mask) Load() ESR_Bits   { return ESR_Bits(rm.UM32.Load()) }
func (rm ESR_Mask) Store(b ESR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CEC_Periph) BTE() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(BTE)}}
}

func (p *CEC_Periph) BPE() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(BPE)}}
}

func (p *CEC_Periph) RBTFE() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(RBTFE)}}
}

func (p *CEC_Periph) SBE() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SBE)}}
}

func (p *CEC_Periph) ACKE() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(ACKE)}}
}

func (p *CEC_Periph) LINE() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(LINE)}}
}

func (p *CEC_Periph) TBTFE() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(TBTFE)}}
}

type CSR_Bits uint32

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CEC_Periph) TSOM() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(TSOM)}}
}

func (p *CEC_Periph) TEOM() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(TEOM)}}
}

func (p *CEC_Periph) TERR() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(TERR)}}
}

func (p *CEC_Periph) TBTRF() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(TBTRF)}}
}

func (p *CEC_Periph) RSOM() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(RSOM)}}
}

func (p *CEC_Periph) REOM() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(REOM)}}
}

func (p *CEC_Periph) RERR() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(RERR)}}
}

func (p *CEC_Periph) RBTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(RBTF)}}
}

type TXD_Bits uint32

type TXD struct{ mmio.U32 }

func (r *TXD) Bits(mask TXD_Bits) TXD_Bits { return TXD_Bits(r.U32.Bits(uint32(mask))) }
func (r *TXD) StoreBits(mask, b TXD_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TXD) SetBits(mask TXD_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *TXD) ClearBits(mask TXD_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *TXD) Load() TXD_Bits              { return TXD_Bits(r.U32.Load()) }
func (r *TXD) Store(b TXD_Bits)            { r.U32.Store(uint32(b)) }

type TXD_Mask struct{ mmio.UM32 }

func (rm TXD_Mask) Load() TXD_Bits   { return TXD_Bits(rm.UM32.Load()) }
func (rm TXD_Mask) Store(b TXD_Bits) { rm.UM32.Store(uint32(b)) }

type RXD_Bits uint32

type RXD struct{ mmio.U32 }

func (r *RXD) Bits(mask RXD_Bits) RXD_Bits { return RXD_Bits(r.U32.Bits(uint32(mask))) }
func (r *RXD) StoreBits(mask, b RXD_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXD) SetBits(mask RXD_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *RXD) ClearBits(mask RXD_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *RXD) Load() RXD_Bits              { return RXD_Bits(r.U32.Load()) }
func (r *RXD) Store(b RXD_Bits)            { r.U32.Store(uint32(b)) }

type RXD_Mask struct{ mmio.UM32 }

func (rm RXD_Mask) Load() RXD_Bits   { return RXD_Bits(rm.UM32.Load()) }
func (rm RXD_Mask) Store(b RXD_Bits) { rm.UM32.Store(uint32(b)) }
