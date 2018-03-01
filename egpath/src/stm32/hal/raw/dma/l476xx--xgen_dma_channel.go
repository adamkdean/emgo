// +build l476xx

package dma

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type DMA_Channel_Periph struct {
	CCR   RCCR
	CNDTR RCNDTR
	CPAR  RCPAR
	CMAR  RCMAR
}

func (p *DMA_Channel_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DMA1_Channel1 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel1_BASE)))

//emgo:const
var DMA1_Channel2 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel2_BASE)))

//emgo:const
var DMA1_Channel3 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel3_BASE)))

//emgo:const
var DMA1_Channel4 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel4_BASE)))

//emgo:const
var DMA1_Channel5 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel5_BASE)))

//emgo:const
var DMA1_Channel6 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel6_BASE)))

//emgo:const
var DMA1_Channel7 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Channel7_BASE)))

//emgo:const
var DMA2_Channel1 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel1_BASE)))

//emgo:const
var DMA2_Channel2 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel2_BASE)))

//emgo:const
var DMA2_Channel3 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel3_BASE)))

//emgo:const
var DMA2_Channel4 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel4_BASE)))

//emgo:const
var DMA2_Channel5 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel5_BASE)))

//emgo:const
var DMA2_Channel6 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel6_BASE)))

//emgo:const
var DMA2_Channel7 = (*DMA_Channel_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Channel7_BASE)))

type CCR uint32

func (b CCR) Field(mask CCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR) J(v int) CCR {
	return CCR(bits.Make32(v, uint32(mask)))
}

type RCCR struct{ mmio.U32 }

func (r *RCCR) Bits(mask CCR) CCR     { return CCR(r.U32.Bits(uint32(mask))) }
func (r *RCCR) StoreBits(mask, b CCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCCR) SetBits(mask CCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCCR) ClearBits(mask CCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCCR) Load() CCR             { return CCR(r.U32.Load()) }
func (r *RCCR) Store(b CCR)           { r.U32.Store(uint32(b)) }

func (r *RCCR) AtomicStoreBits(mask, b CCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCCR) AtomicSetBits(mask CCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCCR) AtomicClearBits(mask CCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCCR struct{ mmio.UM32 }

func (rm RMCCR) Load() CCR   { return CCR(rm.UM32.Load()) }
func (rm RMCCR) Store(b CCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Channel_Periph) EN() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(EN)}}
}

func (p *DMA_Channel_Periph) TCIE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(TCIE)}}
}

func (p *DMA_Channel_Periph) HTIE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(HTIE)}}
}

func (p *DMA_Channel_Periph) TEIE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(TEIE)}}
}

func (p *DMA_Channel_Periph) DIR() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DIR)}}
}

func (p *DMA_Channel_Periph) CIRC() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(CIRC)}}
}

func (p *DMA_Channel_Periph) PINC() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(PINC)}}
}

func (p *DMA_Channel_Periph) MINC() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(MINC)}}
}

func (p *DMA_Channel_Periph) PSIZE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(PSIZE)}}
}

func (p *DMA_Channel_Periph) MSIZE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(MSIZE)}}
}

func (p *DMA_Channel_Periph) PL() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(PL)}}
}

func (p *DMA_Channel_Periph) MEM2MEM() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(MEM2MEM)}}
}

type CNDTR uint32

func (b CNDTR) Field(mask CNDTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CNDTR) J(v int) CNDTR {
	return CNDTR(bits.Make32(v, uint32(mask)))
}

type RCNDTR struct{ mmio.U32 }

func (r *RCNDTR) Bits(mask CNDTR) CNDTR   { return CNDTR(r.U32.Bits(uint32(mask))) }
func (r *RCNDTR) StoreBits(mask, b CNDTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCNDTR) SetBits(mask CNDTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCNDTR) ClearBits(mask CNDTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCNDTR) Load() CNDTR             { return CNDTR(r.U32.Load()) }
func (r *RCNDTR) Store(b CNDTR)           { r.U32.Store(uint32(b)) }

func (r *RCNDTR) AtomicStoreBits(mask, b CNDTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCNDTR) AtomicSetBits(mask CNDTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCNDTR) AtomicClearBits(mask CNDTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCNDTR struct{ mmio.UM32 }

func (rm RMCNDTR) Load() CNDTR   { return CNDTR(rm.UM32.Load()) }
func (rm RMCNDTR) Store(b CNDTR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Channel_Periph) NDT() RMCNDTR {
	return RMCNDTR{mmio.UM32{&p.CNDTR.U32, uint32(NDT)}}
}

type CPAR uint32

func (b CPAR) Field(mask CPAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPAR) J(v int) CPAR {
	return CPAR(bits.Make32(v, uint32(mask)))
}

type RCPAR struct{ mmio.U32 }

func (r *RCPAR) Bits(mask CPAR) CPAR    { return CPAR(r.U32.Bits(uint32(mask))) }
func (r *RCPAR) StoreBits(mask, b CPAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCPAR) SetBits(mask CPAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCPAR) ClearBits(mask CPAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCPAR) Load() CPAR             { return CPAR(r.U32.Load()) }
func (r *RCPAR) Store(b CPAR)           { r.U32.Store(uint32(b)) }

func (r *RCPAR) AtomicStoreBits(mask, b CPAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCPAR) AtomicSetBits(mask CPAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCPAR) AtomicClearBits(mask CPAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCPAR struct{ mmio.UM32 }

func (rm RMCPAR) Load() CPAR   { return CPAR(rm.UM32.Load()) }
func (rm RMCPAR) Store(b CPAR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Channel_Periph) PA() RMCPAR {
	return RMCPAR{mmio.UM32{&p.CPAR.U32, uint32(PA)}}
}

type CMAR uint32

func (b CMAR) Field(mask CMAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMAR) J(v int) CMAR {
	return CMAR(bits.Make32(v, uint32(mask)))
}

type RCMAR struct{ mmio.U32 }

func (r *RCMAR) Bits(mask CMAR) CMAR    { return CMAR(r.U32.Bits(uint32(mask))) }
func (r *RCMAR) StoreBits(mask, b CMAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCMAR) SetBits(mask CMAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCMAR) ClearBits(mask CMAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCMAR) Load() CMAR             { return CMAR(r.U32.Load()) }
func (r *RCMAR) Store(b CMAR)           { r.U32.Store(uint32(b)) }

func (r *RCMAR) AtomicStoreBits(mask, b CMAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCMAR) AtomicSetBits(mask CMAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCMAR) AtomicClearBits(mask CMAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCMAR struct{ mmio.UM32 }

func (rm RMCMAR) Load() CMAR   { return CMAR(rm.UM32.Load()) }
func (rm RMCMAR) Store(b CMAR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Channel_Periph) MA() RMCMAR {
	return RMCMAR{mmio.UM32{&p.CMAR.U32, uint32(MA)}}
}
