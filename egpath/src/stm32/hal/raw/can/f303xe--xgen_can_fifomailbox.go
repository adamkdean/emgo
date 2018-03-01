// +build f303xe

package can

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type CAN_FIFOMailBox_Periph struct {
	RIR  RRIR
	RDTR RRDTR
	RDLR RRDLR
	RDHR RRDHR
}

func (p *CAN_FIFOMailBox_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type RIR uint32

func (b RIR) Field(mask RIR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RIR) J(v int) RIR {
	return RIR(bits.Make32(v, uint32(mask)))
}

type RRIR struct{ mmio.U32 }

func (r *RRIR) Bits(mask RIR) RIR     { return RIR(r.U32.Bits(uint32(mask))) }
func (r *RRIR) StoreBits(mask, b RIR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRIR) SetBits(mask RIR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRIR) ClearBits(mask RIR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRIR) Load() RIR             { return RIR(r.U32.Load()) }
func (r *RRIR) Store(b RIR)           { r.U32.Store(uint32(b)) }

func (r *RRIR) AtomicStoreBits(mask, b RIR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRIR) AtomicSetBits(mask RIR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRIR) AtomicClearBits(mask RIR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRIR struct{ mmio.UM32 }

func (rm RMRIR) Load() RIR   { return RIR(rm.UM32.Load()) }
func (rm RMRIR) Store(b RIR) { rm.UM32.Store(uint32(b)) }

type RDTR uint32

func (b RDTR) Field(mask RDTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDTR) J(v int) RDTR {
	return RDTR(bits.Make32(v, uint32(mask)))
}

type RRDTR struct{ mmio.U32 }

func (r *RRDTR) Bits(mask RDTR) RDTR    { return RDTR(r.U32.Bits(uint32(mask))) }
func (r *RRDTR) StoreBits(mask, b RDTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRDTR) SetBits(mask RDTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRDTR) ClearBits(mask RDTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRDTR) Load() RDTR             { return RDTR(r.U32.Load()) }
func (r *RRDTR) Store(b RDTR)           { r.U32.Store(uint32(b)) }

func (r *RRDTR) AtomicStoreBits(mask, b RDTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRDTR) AtomicSetBits(mask RDTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRDTR) AtomicClearBits(mask RDTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRDTR struct{ mmio.UM32 }

func (rm RMRDTR) Load() RDTR   { return RDTR(rm.UM32.Load()) }
func (rm RMRDTR) Store(b RDTR) { rm.UM32.Store(uint32(b)) }

type RDLR uint32

func (b RDLR) Field(mask RDLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDLR) J(v int) RDLR {
	return RDLR(bits.Make32(v, uint32(mask)))
}

type RRDLR struct{ mmio.U32 }

func (r *RRDLR) Bits(mask RDLR) RDLR    { return RDLR(r.U32.Bits(uint32(mask))) }
func (r *RRDLR) StoreBits(mask, b RDLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRDLR) SetBits(mask RDLR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRDLR) ClearBits(mask RDLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRDLR) Load() RDLR             { return RDLR(r.U32.Load()) }
func (r *RRDLR) Store(b RDLR)           { r.U32.Store(uint32(b)) }

func (r *RRDLR) AtomicStoreBits(mask, b RDLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRDLR) AtomicSetBits(mask RDLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRDLR) AtomicClearBits(mask RDLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRDLR struct{ mmio.UM32 }

func (rm RMRDLR) Load() RDLR   { return RDLR(rm.UM32.Load()) }
func (rm RMRDLR) Store(b RDLR) { rm.UM32.Store(uint32(b)) }

type RDHR uint32

func (b RDHR) Field(mask RDHR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDHR) J(v int) RDHR {
	return RDHR(bits.Make32(v, uint32(mask)))
}

type RRDHR struct{ mmio.U32 }

func (r *RRDHR) Bits(mask RDHR) RDHR    { return RDHR(r.U32.Bits(uint32(mask))) }
func (r *RRDHR) StoreBits(mask, b RDHR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRDHR) SetBits(mask RDHR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRDHR) ClearBits(mask RDHR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRDHR) Load() RDHR             { return RDHR(r.U32.Load()) }
func (r *RRDHR) Store(b RDHR)           { r.U32.Store(uint32(b)) }

func (r *RRDHR) AtomicStoreBits(mask, b RDHR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRDHR) AtomicSetBits(mask RDHR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRDHR) AtomicClearBits(mask RDHR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRDHR struct{ mmio.UM32 }

func (rm RMRDHR) Load() RDHR   { return RDHR(rm.UM32.Load()) }
func (rm RMRDHR) Store(b RDHR) { rm.UM32.Store(uint32(b)) }
