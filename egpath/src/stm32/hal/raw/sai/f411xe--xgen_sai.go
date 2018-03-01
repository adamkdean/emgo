// +build f411xe

package sai

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type SAI_Periph struct {
	GCR RGCR
}

func (p *SAI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SAI1 = (*SAI_Periph)(unsafe.Pointer(uintptr(mmap.SAI1_BASE)))

type GCR uint32

func (b GCR) Field(mask GCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask GCR) J(v int) GCR {
	return GCR(bits.Make32(v, uint32(mask)))
}

type RGCR struct{ mmio.U32 }

func (r *RGCR) Bits(mask GCR) GCR     { return GCR(r.U32.Bits(uint32(mask))) }
func (r *RGCR) StoreBits(mask, b GCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RGCR) SetBits(mask GCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RGCR) ClearBits(mask GCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RGCR) Load() GCR             { return GCR(r.U32.Load()) }
func (r *RGCR) Store(b GCR)           { r.U32.Store(uint32(b)) }

func (r *RGCR) AtomicStoreBits(mask, b GCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RGCR) AtomicSetBits(mask GCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RGCR) AtomicClearBits(mask GCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMGCR struct{ mmio.UM32 }

func (rm RMGCR) Load() GCR   { return GCR(rm.UM32.Load()) }
func (rm RMGCR) Store(b GCR) { rm.UM32.Store(uint32(b)) }

func (p *SAI_Periph) SYNCIN() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(SYNCIN)}}
}

func (p *SAI_Periph) SYNCOUT() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(SYNCOUT)}}
}
