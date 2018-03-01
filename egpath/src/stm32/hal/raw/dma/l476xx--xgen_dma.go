// +build l476xx

package dma

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type DMA_Periph struct {
	ISR  RISR
	IFCR RIFCR
}

func (p *DMA_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DMA1 = (*DMA_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_BASE)))

//emgo:const
var DMA2 = (*DMA_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_BASE)))

type ISR uint32

func (b ISR) Field(mask ISR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR) J(v int) ISR {
	return ISR(bits.Make32(v, uint32(mask)))
}

type RISR struct{ mmio.U32 }

func (r *RISR) Bits(mask ISR) ISR     { return ISR(r.U32.Bits(uint32(mask))) }
func (r *RISR) StoreBits(mask, b ISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISR) SetBits(mask ISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RISR) ClearBits(mask ISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RISR) Load() ISR             { return ISR(r.U32.Load()) }
func (r *RISR) Store(b ISR)           { r.U32.Store(uint32(b)) }

func (r *RISR) AtomicStoreBits(mask, b ISR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RISR) AtomicSetBits(mask ISR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RISR) AtomicClearBits(mask ISR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMISR struct{ mmio.UM32 }

func (rm RMISR) Load() ISR   { return ISR(rm.UM32.Load()) }
func (rm RMISR) Store(b ISR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) GIF1() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(GIF1)}}
}

func (p *DMA_Periph) TCIF1() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCIF1)}}
}

func (p *DMA_Periph) HTIF1() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(HTIF1)}}
}

func (p *DMA_Periph) TEIF1() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEIF1)}}
}

func (p *DMA_Periph) GIF2() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(GIF2)}}
}

func (p *DMA_Periph) TCIF2() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCIF2)}}
}

func (p *DMA_Periph) HTIF2() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(HTIF2)}}
}

func (p *DMA_Periph) TEIF2() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEIF2)}}
}

func (p *DMA_Periph) GIF3() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(GIF3)}}
}

func (p *DMA_Periph) TCIF3() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCIF3)}}
}

func (p *DMA_Periph) HTIF3() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(HTIF3)}}
}

func (p *DMA_Periph) TEIF3() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEIF3)}}
}

func (p *DMA_Periph) GIF4() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(GIF4)}}
}

func (p *DMA_Periph) TCIF4() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCIF4)}}
}

func (p *DMA_Periph) HTIF4() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(HTIF4)}}
}

func (p *DMA_Periph) TEIF4() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEIF4)}}
}

func (p *DMA_Periph) GIF5() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(GIF5)}}
}

func (p *DMA_Periph) TCIF5() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCIF5)}}
}

func (p *DMA_Periph) HTIF5() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(HTIF5)}}
}

func (p *DMA_Periph) TEIF5() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEIF5)}}
}

func (p *DMA_Periph) GIF6() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(GIF6)}}
}

func (p *DMA_Periph) TCIF6() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCIF6)}}
}

func (p *DMA_Periph) HTIF6() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(HTIF6)}}
}

func (p *DMA_Periph) TEIF6() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEIF6)}}
}

func (p *DMA_Periph) GIF7() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(GIF7)}}
}

func (p *DMA_Periph) TCIF7() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCIF7)}}
}

func (p *DMA_Periph) HTIF7() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(HTIF7)}}
}

func (p *DMA_Periph) TEIF7() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEIF7)}}
}

type IFCR uint32

func (b IFCR) Field(mask IFCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IFCR) J(v int) IFCR {
	return IFCR(bits.Make32(v, uint32(mask)))
}

type RIFCR struct{ mmio.U32 }

func (r *RIFCR) Bits(mask IFCR) IFCR    { return IFCR(r.U32.Bits(uint32(mask))) }
func (r *RIFCR) StoreBits(mask, b IFCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIFCR) SetBits(mask IFCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIFCR) ClearBits(mask IFCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIFCR) Load() IFCR             { return IFCR(r.U32.Load()) }
func (r *RIFCR) Store(b IFCR)           { r.U32.Store(uint32(b)) }

func (r *RIFCR) AtomicStoreBits(mask, b IFCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIFCR) AtomicSetBits(mask IFCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIFCR) AtomicClearBits(mask IFCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIFCR struct{ mmio.UM32 }

func (rm RMIFCR) Load() IFCR   { return IFCR(rm.UM32.Load()) }
func (rm RMIFCR) Store(b IFCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) CGIF1() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CGIF1)}}
}

func (p *DMA_Periph) CTCIF1() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTCIF1)}}
}

func (p *DMA_Periph) CHTIF1() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CHTIF1)}}
}

func (p *DMA_Periph) CTEIF1() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTEIF1)}}
}

func (p *DMA_Periph) CGIF2() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CGIF2)}}
}

func (p *DMA_Periph) CTCIF2() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTCIF2)}}
}

func (p *DMA_Periph) CHTIF2() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CHTIF2)}}
}

func (p *DMA_Periph) CTEIF2() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTEIF2)}}
}

func (p *DMA_Periph) CGIF3() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CGIF3)}}
}

func (p *DMA_Periph) CTCIF3() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTCIF3)}}
}

func (p *DMA_Periph) CHTIF3() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CHTIF3)}}
}

func (p *DMA_Periph) CTEIF3() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTEIF3)}}
}

func (p *DMA_Periph) CGIF4() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CGIF4)}}
}

func (p *DMA_Periph) CTCIF4() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTCIF4)}}
}

func (p *DMA_Periph) CHTIF4() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CHTIF4)}}
}

func (p *DMA_Periph) CTEIF4() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTEIF4)}}
}

func (p *DMA_Periph) CGIF5() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CGIF5)}}
}

func (p *DMA_Periph) CTCIF5() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTCIF5)}}
}

func (p *DMA_Periph) CHTIF5() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CHTIF5)}}
}

func (p *DMA_Periph) CTEIF5() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTEIF5)}}
}

func (p *DMA_Periph) CGIF6() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CGIF6)}}
}

func (p *DMA_Periph) CTCIF6() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTCIF6)}}
}

func (p *DMA_Periph) CHTIF6() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CHTIF6)}}
}

func (p *DMA_Periph) CTEIF6() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTEIF6)}}
}

func (p *DMA_Periph) CGIF7() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CGIF7)}}
}

func (p *DMA_Periph) CTCIF7() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTCIF7)}}
}

func (p *DMA_Periph) CHTIF7() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CHTIF7)}}
}

func (p *DMA_Periph) CTEIF7() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTEIF7)}}
}
