package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type FSMC_Bank2_Periph struct {
	PCR2  RPCR2
	SR2   RSR2
	PMEM2 RPMEM2
	PATT2 RPATT2
	_     uint32
	ECCR2 RECCR2
}

func (p *FSMC_Bank2_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank2 = (*FSMC_Bank2_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank2_R_BASE)))

type PCR2 uint32

func (b PCR2) Field(mask PCR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCR2) J(v int) PCR2 {
	return PCR2(bits.Make32(v, uint32(mask)))
}

type RPCR2 struct{ mmio.U32 }

func (r *RPCR2) Bits(mask PCR2) PCR2    { return PCR2(r.U32.Bits(uint32(mask))) }
func (r *RPCR2) StoreBits(mask, b PCR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPCR2) SetBits(mask PCR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RPCR2) ClearBits(mask PCR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPCR2) Load() PCR2             { return PCR2(r.U32.Load()) }
func (r *RPCR2) Store(b PCR2)           { r.U32.Store(uint32(b)) }

func (r *RPCR2) AtomicStoreBits(mask, b PCR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPCR2) AtomicSetBits(mask PCR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPCR2) AtomicClearBits(mask PCR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPCR2 struct{ mmio.UM32 }

func (rm RMPCR2) Load() PCR2   { return PCR2(rm.UM32.Load()) }
func (rm RMPCR2) Store(b PCR2) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) PWAITEN() RMPCR2 {
	return RMPCR2{mmio.UM32{&p.PCR2.U32, uint32(PWAITEN)}}
}

func (p *FSMC_Bank2_Periph) PBKEN() RMPCR2 {
	return RMPCR2{mmio.UM32{&p.PCR2.U32, uint32(PBKEN)}}
}

func (p *FSMC_Bank2_Periph) PTYP() RMPCR2 {
	return RMPCR2{mmio.UM32{&p.PCR2.U32, uint32(PTYP)}}
}

func (p *FSMC_Bank2_Periph) PWID() RMPCR2 {
	return RMPCR2{mmio.UM32{&p.PCR2.U32, uint32(PWID)}}
}

func (p *FSMC_Bank2_Periph) ECCEN() RMPCR2 {
	return RMPCR2{mmio.UM32{&p.PCR2.U32, uint32(ECCEN)}}
}

func (p *FSMC_Bank2_Periph) TCLR() RMPCR2 {
	return RMPCR2{mmio.UM32{&p.PCR2.U32, uint32(TCLR)}}
}

func (p *FSMC_Bank2_Periph) TAR() RMPCR2 {
	return RMPCR2{mmio.UM32{&p.PCR2.U32, uint32(TAR)}}
}

func (p *FSMC_Bank2_Periph) ECCPS() RMPCR2 {
	return RMPCR2{mmio.UM32{&p.PCR2.U32, uint32(ECCPS)}}
}

type SR2 uint32

func (b SR2) Field(mask SR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR2) J(v int) SR2 {
	return SR2(bits.Make32(v, uint32(mask)))
}

type RSR2 struct{ mmio.U32 }

func (r *RSR2) Bits(mask SR2) SR2     { return SR2(r.U32.Bits(uint32(mask))) }
func (r *RSR2) StoreBits(mask, b SR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR2) SetBits(mask SR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR2) ClearBits(mask SR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR2) Load() SR2             { return SR2(r.U32.Load()) }
func (r *RSR2) Store(b SR2)           { r.U32.Store(uint32(b)) }

func (r *RSR2) AtomicStoreBits(mask, b SR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR2) AtomicSetBits(mask SR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR2) AtomicClearBits(mask SR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR2 struct{ mmio.UM32 }

func (rm RMSR2) Load() SR2   { return SR2(rm.UM32.Load()) }
func (rm RMSR2) Store(b SR2) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) IRS() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(IRS)}}
}

func (p *FSMC_Bank2_Periph) ILS() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(ILS)}}
}

func (p *FSMC_Bank2_Periph) IFS() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(IFS)}}
}

func (p *FSMC_Bank2_Periph) IREN() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(IREN)}}
}

func (p *FSMC_Bank2_Periph) ILEN() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(ILEN)}}
}

func (p *FSMC_Bank2_Periph) IFEN() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(IFEN)}}
}

func (p *FSMC_Bank2_Periph) FEMPT() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(FEMPT)}}
}

type PMEM2 uint32

func (b PMEM2) Field(mask PMEM2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMEM2) J(v int) PMEM2 {
	return PMEM2(bits.Make32(v, uint32(mask)))
}

type RPMEM2 struct{ mmio.U32 }

func (r *RPMEM2) Bits(mask PMEM2) PMEM2   { return PMEM2(r.U32.Bits(uint32(mask))) }
func (r *RPMEM2) StoreBits(mask, b PMEM2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPMEM2) SetBits(mask PMEM2)      { r.U32.SetBits(uint32(mask)) }
func (r *RPMEM2) ClearBits(mask PMEM2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPMEM2) Load() PMEM2             { return PMEM2(r.U32.Load()) }
func (r *RPMEM2) Store(b PMEM2)           { r.U32.Store(uint32(b)) }

func (r *RPMEM2) AtomicStoreBits(mask, b PMEM2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPMEM2) AtomicSetBits(mask PMEM2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPMEM2) AtomicClearBits(mask PMEM2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPMEM2 struct{ mmio.UM32 }

func (rm RMPMEM2) Load() PMEM2   { return PMEM2(rm.UM32.Load()) }
func (rm RMPMEM2) Store(b PMEM2) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) MEMSET2() RMPMEM2 {
	return RMPMEM2{mmio.UM32{&p.PMEM2.U32, uint32(MEMSET2)}}
}

func (p *FSMC_Bank2_Periph) MEMWAIT2() RMPMEM2 {
	return RMPMEM2{mmio.UM32{&p.PMEM2.U32, uint32(MEMWAIT2)}}
}

func (p *FSMC_Bank2_Periph) MEMHOLD2() RMPMEM2 {
	return RMPMEM2{mmio.UM32{&p.PMEM2.U32, uint32(MEMHOLD2)}}
}

func (p *FSMC_Bank2_Periph) MEMHIZ2() RMPMEM2 {
	return RMPMEM2{mmio.UM32{&p.PMEM2.U32, uint32(MEMHIZ2)}}
}

type PATT2 uint32

func (b PATT2) Field(mask PATT2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PATT2) J(v int) PATT2 {
	return PATT2(bits.Make32(v, uint32(mask)))
}

type RPATT2 struct{ mmio.U32 }

func (r *RPATT2) Bits(mask PATT2) PATT2   { return PATT2(r.U32.Bits(uint32(mask))) }
func (r *RPATT2) StoreBits(mask, b PATT2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPATT2) SetBits(mask PATT2)      { r.U32.SetBits(uint32(mask)) }
func (r *RPATT2) ClearBits(mask PATT2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPATT2) Load() PATT2             { return PATT2(r.U32.Load()) }
func (r *RPATT2) Store(b PATT2)           { r.U32.Store(uint32(b)) }

func (r *RPATT2) AtomicStoreBits(mask, b PATT2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPATT2) AtomicSetBits(mask PATT2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPATT2) AtomicClearBits(mask PATT2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPATT2 struct{ mmio.UM32 }

func (rm RMPATT2) Load() PATT2   { return PATT2(rm.UM32.Load()) }
func (rm RMPATT2) Store(b PATT2) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) ATTSET2() RMPATT2 {
	return RMPATT2{mmio.UM32{&p.PATT2.U32, uint32(ATTSET2)}}
}

func (p *FSMC_Bank2_Periph) ATTWAIT2() RMPATT2 {
	return RMPATT2{mmio.UM32{&p.PATT2.U32, uint32(ATTWAIT2)}}
}

func (p *FSMC_Bank2_Periph) ATTHOLD2() RMPATT2 {
	return RMPATT2{mmio.UM32{&p.PATT2.U32, uint32(ATTHOLD2)}}
}

func (p *FSMC_Bank2_Periph) ATTHIZ2() RMPATT2 {
	return RMPATT2{mmio.UM32{&p.PATT2.U32, uint32(ATTHIZ2)}}
}

type ECCR2 uint32

func (b ECCR2) Field(mask ECCR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ECCR2) J(v int) ECCR2 {
	return ECCR2(bits.Make32(v, uint32(mask)))
}

type RECCR2 struct{ mmio.U32 }

func (r *RECCR2) Bits(mask ECCR2) ECCR2   { return ECCR2(r.U32.Bits(uint32(mask))) }
func (r *RECCR2) StoreBits(mask, b ECCR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RECCR2) SetBits(mask ECCR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RECCR2) ClearBits(mask ECCR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RECCR2) Load() ECCR2             { return ECCR2(r.U32.Load()) }
func (r *RECCR2) Store(b ECCR2)           { r.U32.Store(uint32(b)) }

func (r *RECCR2) AtomicStoreBits(mask, b ECCR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RECCR2) AtomicSetBits(mask ECCR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RECCR2) AtomicClearBits(mask ECCR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMECCR2 struct{ mmio.UM32 }

func (rm RMECCR2) Load() ECCR2   { return ECCR2(rm.UM32.Load()) }
func (rm RMECCR2) Store(b ECCR2) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) ECC2() RMECCR2 {
	return RMECCR2{mmio.UM32{&p.ECCR2.U32, uint32(ECC2)}}
}
