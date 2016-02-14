package can

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type CAN_TxMailBox_Periph struct {
	TIR  TIR
	TDTR TDTR
	TDLR TDLR
	TDHR TDHR
}

func (p *CAN_TxMailBox_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type TIR_Bits uint32

type TIR struct{ mmio.U32 }

func (r *TIR) Bits(mask TIR_Bits) TIR_Bits { return TIR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TIR) StoreBits(mask, b TIR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TIR) SetBits(mask TIR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *TIR) ClearBits(mask TIR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *TIR) Load() TIR_Bits              { return TIR_Bits(r.U32.Load()) }
func (r *TIR) Store(b TIR_Bits)            { r.U32.Store(uint32(b)) }

type TIR_Mask struct{ mmio.UM32 }

func (rm TIR_Mask) Load() TIR_Bits   { return TIR_Bits(rm.UM32.Load()) }
func (rm TIR_Mask) Store(b TIR_Bits) { rm.UM32.Store(uint32(b)) }

type TDTR_Bits uint32

type TDTR struct{ mmio.U32 }

func (r *TDTR) Bits(mask TDTR_Bits) TDTR_Bits { return TDTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TDTR) StoreBits(mask, b TDTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TDTR) SetBits(mask TDTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TDTR) ClearBits(mask TDTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TDTR) Load() TDTR_Bits               { return TDTR_Bits(r.U32.Load()) }
func (r *TDTR) Store(b TDTR_Bits)             { r.U32.Store(uint32(b)) }

type TDTR_Mask struct{ mmio.UM32 }

func (rm TDTR_Mask) Load() TDTR_Bits   { return TDTR_Bits(rm.UM32.Load()) }
func (rm TDTR_Mask) Store(b TDTR_Bits) { rm.UM32.Store(uint32(b)) }

type TDLR_Bits uint32

type TDLR struct{ mmio.U32 }

func (r *TDLR) Bits(mask TDLR_Bits) TDLR_Bits { return TDLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TDLR) StoreBits(mask, b TDLR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TDLR) SetBits(mask TDLR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TDLR) ClearBits(mask TDLR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TDLR) Load() TDLR_Bits               { return TDLR_Bits(r.U32.Load()) }
func (r *TDLR) Store(b TDLR_Bits)             { r.U32.Store(uint32(b)) }

type TDLR_Mask struct{ mmio.UM32 }

func (rm TDLR_Mask) Load() TDLR_Bits   { return TDLR_Bits(rm.UM32.Load()) }
func (rm TDLR_Mask) Store(b TDLR_Bits) { rm.UM32.Store(uint32(b)) }

type TDHR_Bits uint32

type TDHR struct{ mmio.U32 }

func (r *TDHR) Bits(mask TDHR_Bits) TDHR_Bits { return TDHR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TDHR) StoreBits(mask, b TDHR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TDHR) SetBits(mask TDHR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TDHR) ClearBits(mask TDHR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TDHR) Load() TDHR_Bits               { return TDHR_Bits(r.U32.Load()) }
func (r *TDHR) Store(b TDHR_Bits)             { r.U32.Store(uint32(b)) }

type TDHR_Mask struct{ mmio.UM32 }

func (rm TDHR_Mask) Load() TDHR_Bits   { return TDHR_Bits(rm.UM32.Load()) }
func (rm TDHR_Mask) Store(b TDHR_Bits) { rm.UM32.Store(uint32(b)) }
