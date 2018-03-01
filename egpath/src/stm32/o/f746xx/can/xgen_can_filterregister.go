package can

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type CAN_FilterRegister_Periph struct {
	FR1 RFR1
	FR2 RFR2
}

func (p *CAN_FilterRegister_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type FR1 uint32

func (b FR1) Field(mask FR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FR1) J(v int) FR1 {
	return FR1(bits.Make32(v, uint32(mask)))
}

type RFR1 struct{ mmio.U32 }

func (r *RFR1) Bits(mask FR1) FR1     { return FR1(r.U32.Bits(uint32(mask))) }
func (r *RFR1) StoreBits(mask, b FR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFR1) SetBits(mask FR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RFR1) ClearBits(mask FR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFR1) Load() FR1             { return FR1(r.U32.Load()) }
func (r *RFR1) Store(b FR1)           { r.U32.Store(uint32(b)) }

func (r *RFR1) AtomicStoreBits(mask, b FR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFR1) AtomicSetBits(mask FR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFR1) AtomicClearBits(mask FR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFR1 struct{ mmio.UM32 }

func (rm RMFR1) Load() FR1   { return FR1(rm.UM32.Load()) }
func (rm RMFR1) Store(b FR1) { rm.UM32.Store(uint32(b)) }

type FR2 uint32

func (b FR2) Field(mask FR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FR2) J(v int) FR2 {
	return FR2(bits.Make32(v, uint32(mask)))
}

type RFR2 struct{ mmio.U32 }

func (r *RFR2) Bits(mask FR2) FR2     { return FR2(r.U32.Bits(uint32(mask))) }
func (r *RFR2) StoreBits(mask, b FR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFR2) SetBits(mask FR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RFR2) ClearBits(mask FR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFR2) Load() FR2             { return FR2(r.U32.Load()) }
func (r *RFR2) Store(b FR2)           { r.U32.Store(uint32(b)) }

func (r *RFR2) AtomicStoreBits(mask, b FR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFR2) AtomicSetBits(mask FR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFR2) AtomicClearBits(mask FR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFR2 struct{ mmio.UM32 }

func (rm RMFR2) Load() FR2   { return FR2(rm.UM32.Load()) }
func (rm RMFR2) Store(b FR2) { rm.UM32.Store(uint32(b)) }
