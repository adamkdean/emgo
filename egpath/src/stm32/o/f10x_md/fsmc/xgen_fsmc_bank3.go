package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type FSMC_Bank3_Periph struct {
	PCR3  PCR3
	SR3   SR3
	PMEM3 PMEM3
	PATT3 PATT3
	_     uint32
	ECCR3 ECCR3
}

func (p *FSMC_Bank3_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var FSMC_Bank3 = (*FSMC_Bank3_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank3_R_BASE)))

type PCR3_Bits uint32

type PCR3 struct{ mmio.U32 }

func (r *PCR3) Bits(mask PCR3_Bits) PCR3_Bits { return PCR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *PCR3) StoreBits(mask, b PCR3_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCR3) SetBits(mask PCR3_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PCR3) ClearBits(mask PCR3_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PCR3) Load() PCR3_Bits               { return PCR3_Bits(r.U32.Load()) }
func (r *PCR3) Store(b PCR3_Bits)             { r.U32.Store(uint32(b)) }

type PCR3_Mask struct{ mmio.UM32 }

func (rm PCR3_Mask) Load() PCR3_Bits   { return PCR3_Bits(rm.UM32.Load()) }
func (rm PCR3_Mask) Store(b PCR3_Bits) { rm.UM32.Store(uint32(b)) }

type SR3_Bits uint32

type SR3 struct{ mmio.U32 }

func (r *SR3) Bits(mask SR3_Bits) SR3_Bits { return SR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR3) StoreBits(mask, b SR3_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR3) SetBits(mask SR3_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *SR3) ClearBits(mask SR3_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *SR3) Load() SR3_Bits              { return SR3_Bits(r.U32.Load()) }
func (r *SR3) Store(b SR3_Bits)            { r.U32.Store(uint32(b)) }

type SR3_Mask struct{ mmio.UM32 }

func (rm SR3_Mask) Load() SR3_Bits   { return SR3_Bits(rm.UM32.Load()) }
func (rm SR3_Mask) Store(b SR3_Bits) { rm.UM32.Store(uint32(b)) }

type PMEM3_Bits uint32

type PMEM3 struct{ mmio.U32 }

func (r *PMEM3) Bits(mask PMEM3_Bits) PMEM3_Bits { return PMEM3_Bits(r.U32.Bits(uint32(mask))) }
func (r *PMEM3) StoreBits(mask, b PMEM3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PMEM3) SetBits(mask PMEM3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PMEM3) ClearBits(mask PMEM3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PMEM3) Load() PMEM3_Bits                { return PMEM3_Bits(r.U32.Load()) }
func (r *PMEM3) Store(b PMEM3_Bits)              { r.U32.Store(uint32(b)) }

type PMEM3_Mask struct{ mmio.UM32 }

func (rm PMEM3_Mask) Load() PMEM3_Bits   { return PMEM3_Bits(rm.UM32.Load()) }
func (rm PMEM3_Mask) Store(b PMEM3_Bits) { rm.UM32.Store(uint32(b)) }

type PATT3_Bits uint32

type PATT3 struct{ mmio.U32 }

func (r *PATT3) Bits(mask PATT3_Bits) PATT3_Bits { return PATT3_Bits(r.U32.Bits(uint32(mask))) }
func (r *PATT3) StoreBits(mask, b PATT3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PATT3) SetBits(mask PATT3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PATT3) ClearBits(mask PATT3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PATT3) Load() PATT3_Bits                { return PATT3_Bits(r.U32.Load()) }
func (r *PATT3) Store(b PATT3_Bits)              { r.U32.Store(uint32(b)) }

type PATT3_Mask struct{ mmio.UM32 }

func (rm PATT3_Mask) Load() PATT3_Bits   { return PATT3_Bits(rm.UM32.Load()) }
func (rm PATT3_Mask) Store(b PATT3_Bits) { rm.UM32.Store(uint32(b)) }

type ECCR3_Bits uint32

type ECCR3 struct{ mmio.U32 }

func (r *ECCR3) Bits(mask ECCR3_Bits) ECCR3_Bits { return ECCR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *ECCR3) StoreBits(mask, b ECCR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ECCR3) SetBits(mask ECCR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ECCR3) ClearBits(mask ECCR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ECCR3) Load() ECCR3_Bits                { return ECCR3_Bits(r.U32.Load()) }
func (r *ECCR3) Store(b ECCR3_Bits)              { r.U32.Store(uint32(b)) }

type ECCR3_Mask struct{ mmio.UM32 }

func (rm ECCR3_Mask) Load() ECCR3_Bits   { return ECCR3_Bits(rm.UM32.Load()) }
func (rm ECCR3_Mask) Store(b ECCR3_Bits) { rm.UM32.Store(uint32(b)) }
