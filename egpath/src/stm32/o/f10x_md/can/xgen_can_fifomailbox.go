package can

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type CAN_FIFOMailBox_Periph struct {
	RIR  RIR
	RDTR RDTR
	RDLR RDLR
	RDHR RDHR
}

func (p *CAN_FIFOMailBox_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type RIR_Bits uint32

type RIR struct{ mmio.U32 }

func (r *RIR) Bits(mask RIR_Bits) RIR_Bits { return RIR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RIR) StoreBits(mask, b RIR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIR) SetBits(mask RIR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *RIR) ClearBits(mask RIR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *RIR) Load() RIR_Bits              { return RIR_Bits(r.U32.Load()) }
func (r *RIR) Store(b RIR_Bits)            { r.U32.Store(uint32(b)) }

type RIR_Mask struct{ mmio.UM32 }

func (rm RIR_Mask) Load() RIR_Bits   { return RIR_Bits(rm.UM32.Load()) }
func (rm RIR_Mask) Store(b RIR_Bits) { rm.UM32.Store(uint32(b)) }

type RDTR_Bits uint32

type RDTR struct{ mmio.U32 }

func (r *RDTR) Bits(mask RDTR_Bits) RDTR_Bits { return RDTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RDTR) StoreBits(mask, b RDTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDTR) SetBits(mask RDTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RDTR) ClearBits(mask RDTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RDTR) Load() RDTR_Bits               { return RDTR_Bits(r.U32.Load()) }
func (r *RDTR) Store(b RDTR_Bits)             { r.U32.Store(uint32(b)) }

type RDTR_Mask struct{ mmio.UM32 }

func (rm RDTR_Mask) Load() RDTR_Bits   { return RDTR_Bits(rm.UM32.Load()) }
func (rm RDTR_Mask) Store(b RDTR_Bits) { rm.UM32.Store(uint32(b)) }

type RDLR_Bits uint32

type RDLR struct{ mmio.U32 }

func (r *RDLR) Bits(mask RDLR_Bits) RDLR_Bits { return RDLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RDLR) StoreBits(mask, b RDLR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDLR) SetBits(mask RDLR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RDLR) ClearBits(mask RDLR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RDLR) Load() RDLR_Bits               { return RDLR_Bits(r.U32.Load()) }
func (r *RDLR) Store(b RDLR_Bits)             { r.U32.Store(uint32(b)) }

type RDLR_Mask struct{ mmio.UM32 }

func (rm RDLR_Mask) Load() RDLR_Bits   { return RDLR_Bits(rm.UM32.Load()) }
func (rm RDLR_Mask) Store(b RDLR_Bits) { rm.UM32.Store(uint32(b)) }

type RDHR_Bits uint32

type RDHR struct{ mmio.U32 }

func (r *RDHR) Bits(mask RDHR_Bits) RDHR_Bits { return RDHR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RDHR) StoreBits(mask, b RDHR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR) SetBits(mask RDHR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RDHR) ClearBits(mask RDHR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR) Load() RDHR_Bits               { return RDHR_Bits(r.U32.Load()) }
func (r *RDHR) Store(b RDHR_Bits)             { r.U32.Store(uint32(b)) }

type RDHR_Mask struct{ mmio.UM32 }

func (rm RDHR_Mask) Load() RDHR_Bits   { return RDHR_Bits(rm.UM32.Load()) }
func (rm RDHR_Mask) Store(b RDHR_Bits) { rm.UM32.Store(uint32(b)) }
