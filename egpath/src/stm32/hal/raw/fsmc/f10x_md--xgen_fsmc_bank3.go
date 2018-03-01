// +build f10x_md

package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type FSMC_Bank3_Periph struct {
	PCR3  RPCR3
	SR3   RSR3
	PMEM3 RPMEM3
	PATT3 RPATT3
	_     uint32
	ECCR3 RECCR3
}

func (p *FSMC_Bank3_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank3 = (*FSMC_Bank3_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank3_R_BASE)))

type PCR3 uint32

func (b PCR3) Field(mask PCR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCR3) J(v int) PCR3 {
	return PCR3(bits.Make32(v, uint32(mask)))
}

type RPCR3 struct{ mmio.U32 }

func (r *RPCR3) Bits(mask PCR3) PCR3    { return PCR3(r.U32.Bits(uint32(mask))) }
func (r *RPCR3) StoreBits(mask, b PCR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPCR3) SetBits(mask PCR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RPCR3) ClearBits(mask PCR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPCR3) Load() PCR3             { return PCR3(r.U32.Load()) }
func (r *RPCR3) Store(b PCR3)           { r.U32.Store(uint32(b)) }

func (r *RPCR3) AtomicStoreBits(mask, b PCR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPCR3) AtomicSetBits(mask PCR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPCR3) AtomicClearBits(mask PCR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPCR3 struct{ mmio.UM32 }

func (rm RMPCR3) Load() PCR3   { return PCR3(rm.UM32.Load()) }
func (rm RMPCR3) Store(b PCR3) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) PWAITEN() RMPCR3 {
	return RMPCR3{mmio.UM32{&p.PCR3.U32, uint32(PWAITEN)}}
}

func (p *FSMC_Bank3_Periph) PBKEN() RMPCR3 {
	return RMPCR3{mmio.UM32{&p.PCR3.U32, uint32(PBKEN)}}
}

func (p *FSMC_Bank3_Periph) PTYP() RMPCR3 {
	return RMPCR3{mmio.UM32{&p.PCR3.U32, uint32(PTYP)}}
}

func (p *FSMC_Bank3_Periph) PWID() RMPCR3 {
	return RMPCR3{mmio.UM32{&p.PCR3.U32, uint32(PWID)}}
}

func (p *FSMC_Bank3_Periph) ECCEN() RMPCR3 {
	return RMPCR3{mmio.UM32{&p.PCR3.U32, uint32(ECCEN)}}
}

func (p *FSMC_Bank3_Periph) TCLR() RMPCR3 {
	return RMPCR3{mmio.UM32{&p.PCR3.U32, uint32(TCLR)}}
}

func (p *FSMC_Bank3_Periph) TAR() RMPCR3 {
	return RMPCR3{mmio.UM32{&p.PCR3.U32, uint32(TAR)}}
}

func (p *FSMC_Bank3_Periph) ECCPS() RMPCR3 {
	return RMPCR3{mmio.UM32{&p.PCR3.U32, uint32(ECCPS)}}
}

type SR3 uint32

func (b SR3) Field(mask SR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR3) J(v int) SR3 {
	return SR3(bits.Make32(v, uint32(mask)))
}

type RSR3 struct{ mmio.U32 }

func (r *RSR3) Bits(mask SR3) SR3     { return SR3(r.U32.Bits(uint32(mask))) }
func (r *RSR3) StoreBits(mask, b SR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR3) SetBits(mask SR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR3) ClearBits(mask SR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR3) Load() SR3             { return SR3(r.U32.Load()) }
func (r *RSR3) Store(b SR3)           { r.U32.Store(uint32(b)) }

func (r *RSR3) AtomicStoreBits(mask, b SR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR3) AtomicSetBits(mask SR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR3) AtomicClearBits(mask SR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR3 struct{ mmio.UM32 }

func (rm RMSR3) Load() SR3   { return SR3(rm.UM32.Load()) }
func (rm RMSR3) Store(b SR3) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) IRS() RMSR3 {
	return RMSR3{mmio.UM32{&p.SR3.U32, uint32(IRS)}}
}

func (p *FSMC_Bank3_Periph) ILS() RMSR3 {
	return RMSR3{mmio.UM32{&p.SR3.U32, uint32(ILS)}}
}

func (p *FSMC_Bank3_Periph) IFS() RMSR3 {
	return RMSR3{mmio.UM32{&p.SR3.U32, uint32(IFS)}}
}

func (p *FSMC_Bank3_Periph) IREN() RMSR3 {
	return RMSR3{mmio.UM32{&p.SR3.U32, uint32(IREN)}}
}

func (p *FSMC_Bank3_Periph) ILEN() RMSR3 {
	return RMSR3{mmio.UM32{&p.SR3.U32, uint32(ILEN)}}
}

func (p *FSMC_Bank3_Periph) IFEN() RMSR3 {
	return RMSR3{mmio.UM32{&p.SR3.U32, uint32(IFEN)}}
}

func (p *FSMC_Bank3_Periph) FEMPT() RMSR3 {
	return RMSR3{mmio.UM32{&p.SR3.U32, uint32(FEMPT)}}
}

type PMEM3 uint32

func (b PMEM3) Field(mask PMEM3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMEM3) J(v int) PMEM3 {
	return PMEM3(bits.Make32(v, uint32(mask)))
}

type RPMEM3 struct{ mmio.U32 }

func (r *RPMEM3) Bits(mask PMEM3) PMEM3   { return PMEM3(r.U32.Bits(uint32(mask))) }
func (r *RPMEM3) StoreBits(mask, b PMEM3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPMEM3) SetBits(mask PMEM3)      { r.U32.SetBits(uint32(mask)) }
func (r *RPMEM3) ClearBits(mask PMEM3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPMEM3) Load() PMEM3             { return PMEM3(r.U32.Load()) }
func (r *RPMEM3) Store(b PMEM3)           { r.U32.Store(uint32(b)) }

func (r *RPMEM3) AtomicStoreBits(mask, b PMEM3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPMEM3) AtomicSetBits(mask PMEM3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPMEM3) AtomicClearBits(mask PMEM3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPMEM3 struct{ mmio.UM32 }

func (rm RMPMEM3) Load() PMEM3   { return PMEM3(rm.UM32.Load()) }
func (rm RMPMEM3) Store(b PMEM3) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) MEMSET3() RMPMEM3 {
	return RMPMEM3{mmio.UM32{&p.PMEM3.U32, uint32(MEMSET3)}}
}

func (p *FSMC_Bank3_Periph) MEMWAIT3() RMPMEM3 {
	return RMPMEM3{mmio.UM32{&p.PMEM3.U32, uint32(MEMWAIT3)}}
}

func (p *FSMC_Bank3_Periph) MEMHOLD3() RMPMEM3 {
	return RMPMEM3{mmio.UM32{&p.PMEM3.U32, uint32(MEMHOLD3)}}
}

func (p *FSMC_Bank3_Periph) MEMHIZ3() RMPMEM3 {
	return RMPMEM3{mmio.UM32{&p.PMEM3.U32, uint32(MEMHIZ3)}}
}

type PATT3 uint32

func (b PATT3) Field(mask PATT3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PATT3) J(v int) PATT3 {
	return PATT3(bits.Make32(v, uint32(mask)))
}

type RPATT3 struct{ mmio.U32 }

func (r *RPATT3) Bits(mask PATT3) PATT3   { return PATT3(r.U32.Bits(uint32(mask))) }
func (r *RPATT3) StoreBits(mask, b PATT3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPATT3) SetBits(mask PATT3)      { r.U32.SetBits(uint32(mask)) }
func (r *RPATT3) ClearBits(mask PATT3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPATT3) Load() PATT3             { return PATT3(r.U32.Load()) }
func (r *RPATT3) Store(b PATT3)           { r.U32.Store(uint32(b)) }

func (r *RPATT3) AtomicStoreBits(mask, b PATT3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPATT3) AtomicSetBits(mask PATT3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPATT3) AtomicClearBits(mask PATT3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPATT3 struct{ mmio.UM32 }

func (rm RMPATT3) Load() PATT3   { return PATT3(rm.UM32.Load()) }
func (rm RMPATT3) Store(b PATT3) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) ATTSET3() RMPATT3 {
	return RMPATT3{mmio.UM32{&p.PATT3.U32, uint32(ATTSET3)}}
}

func (p *FSMC_Bank3_Periph) ATTWAIT3() RMPATT3 {
	return RMPATT3{mmio.UM32{&p.PATT3.U32, uint32(ATTWAIT3)}}
}

func (p *FSMC_Bank3_Periph) ATTHOLD3() RMPATT3 {
	return RMPATT3{mmio.UM32{&p.PATT3.U32, uint32(ATTHOLD3)}}
}

func (p *FSMC_Bank3_Periph) ATTHIZ3() RMPATT3 {
	return RMPATT3{mmio.UM32{&p.PATT3.U32, uint32(ATTHIZ3)}}
}

type ECCR3 uint32

func (b ECCR3) Field(mask ECCR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ECCR3) J(v int) ECCR3 {
	return ECCR3(bits.Make32(v, uint32(mask)))
}

type RECCR3 struct{ mmio.U32 }

func (r *RECCR3) Bits(mask ECCR3) ECCR3   { return ECCR3(r.U32.Bits(uint32(mask))) }
func (r *RECCR3) StoreBits(mask, b ECCR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RECCR3) SetBits(mask ECCR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RECCR3) ClearBits(mask ECCR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RECCR3) Load() ECCR3             { return ECCR3(r.U32.Load()) }
func (r *RECCR3) Store(b ECCR3)           { r.U32.Store(uint32(b)) }

func (r *RECCR3) AtomicStoreBits(mask, b ECCR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RECCR3) AtomicSetBits(mask ECCR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RECCR3) AtomicClearBits(mask ECCR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMECCR3 struct{ mmio.UM32 }

func (rm RMECCR3) Load() ECCR3   { return ECCR3(rm.UM32.Load()) }
func (rm RMECCR3) Store(b ECCR3) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank3_Periph) ECC3() RMECCR3 {
	return RMECCR3{mmio.UM32{&p.ECCR3.U32, uint32(ECC3)}}
}
