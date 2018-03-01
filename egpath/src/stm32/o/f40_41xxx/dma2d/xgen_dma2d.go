package dma2d

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type DMA2D_Periph struct {
	CR      RCR
	ISR     RISR
	IFCR    RIFCR
	FGMAR   RFGMAR
	FGOR    RFGOR
	BGMAR   RBGMAR
	BGOR    RBGOR
	FGPFCCR RFGPFCCR
	FGCOLR  RFGCOLR
	BGPFCCR RBGPFCCR
	BGCOLR  RBGCOLR
	FGCMAR  RFGCMAR
	BGCMAR  RBGCMAR
	OPFCCR  ROPFCCR
	OCOLR   ROCOLR
	OMAR    ROMAR
	OOR     ROOR
	NLR     RNLR
	LWR     RLWR
	AMTCR   RAMTCR
	_       [236]uint32
	FGCLUT  [256]RFGCLUT
	BGCLUT  [256]RBGCLUT
}

func (p *DMA2D_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DMA2D = (*DMA2D_Periph)(unsafe.Pointer(uintptr(mmap.DMA2D_BASE)))

type CR uint32

func (b CR) Field(mask CR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR) J(v int) CR {
	return CR(bits.Make32(v, uint32(mask)))
}

type RCR struct{ mmio.U32 }

func (r *RCR) Bits(mask CR) CR      { return CR(r.U32.Bits(uint32(mask))) }
func (r *RCR) StoreBits(mask, b CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR) SetBits(mask CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR) ClearBits(mask CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR) Load() CR             { return CR(r.U32.Load()) }
func (r *RCR) Store(b CR)           { r.U32.Store(uint32(b)) }

func (r *RCR) AtomicStoreBits(mask, b CR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR) AtomicSetBits(mask CR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR) AtomicClearBits(mask CR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR struct{ mmio.UM32 }

func (rm RMCR) Load() CR   { return CR(rm.UM32.Load()) }
func (rm RMCR) Store(b CR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) START() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(START)}}
}

func (p *DMA2D_Periph) SUSP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SUSP)}}
}

func (p *DMA2D_Periph) ABORT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ABORT)}}
}

func (p *DMA2D_Periph) TEIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TEIE)}}
}

func (p *DMA2D_Periph) TCIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TCIE)}}
}

func (p *DMA2D_Periph) TWIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TWIE)}}
}

func (p *DMA2D_Periph) CAEIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CAEIE)}}
}

func (p *DMA2D_Periph) CTCIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CTCIE)}}
}

func (p *DMA2D_Periph) CEIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CEIE)}}
}

func (p *DMA2D_Periph) MODE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MODE)}}
}

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

func (p *DMA2D_Periph) TEIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEIF)}}
}

func (p *DMA2D_Periph) TCIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCIF)}}
}

func (p *DMA2D_Periph) TWIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TWIF)}}
}

func (p *DMA2D_Periph) CAEIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(CAEIF)}}
}

func (p *DMA2D_Periph) CTCIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(CTCIF)}}
}

func (p *DMA2D_Periph) CEIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(CEIF)}}
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

func (p *DMA2D_Periph) CTEIF() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTEIF)}}
}

func (p *DMA2D_Periph) CTCIF() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTCIF)}}
}

func (p *DMA2D_Periph) CTWIF() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CTWIF)}}
}

func (p *DMA2D_Periph) CAECIF() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CAECIF)}}
}

func (p *DMA2D_Periph) CCTCIF() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CCTCIF)}}
}

func (p *DMA2D_Periph) CCEIF() RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CCEIF)}}
}

type FGMAR uint32

func (b FGMAR) Field(mask FGMAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FGMAR) J(v int) FGMAR {
	return FGMAR(bits.Make32(v, uint32(mask)))
}

type RFGMAR struct{ mmio.U32 }

func (r *RFGMAR) Bits(mask FGMAR) FGMAR   { return FGMAR(r.U32.Bits(uint32(mask))) }
func (r *RFGMAR) StoreBits(mask, b FGMAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFGMAR) SetBits(mask FGMAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFGMAR) ClearBits(mask FGMAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFGMAR) Load() FGMAR             { return FGMAR(r.U32.Load()) }
func (r *RFGMAR) Store(b FGMAR)           { r.U32.Store(uint32(b)) }

func (r *RFGMAR) AtomicStoreBits(mask, b FGMAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFGMAR) AtomicSetBits(mask FGMAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFGMAR) AtomicClearBits(mask FGMAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFGMAR struct{ mmio.UM32 }

func (rm RMFGMAR) Load() FGMAR   { return FGMAR(rm.UM32.Load()) }
func (rm RMFGMAR) Store(b FGMAR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() RMFGMAR {
	return RMFGMAR{mmio.UM32{&p.FGMAR.U32, uint32(MA)}}
}

type FGOR uint32

func (b FGOR) Field(mask FGOR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FGOR) J(v int) FGOR {
	return FGOR(bits.Make32(v, uint32(mask)))
}

type RFGOR struct{ mmio.U32 }

func (r *RFGOR) Bits(mask FGOR) FGOR    { return FGOR(r.U32.Bits(uint32(mask))) }
func (r *RFGOR) StoreBits(mask, b FGOR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFGOR) SetBits(mask FGOR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFGOR) ClearBits(mask FGOR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFGOR) Load() FGOR             { return FGOR(r.U32.Load()) }
func (r *RFGOR) Store(b FGOR)           { r.U32.Store(uint32(b)) }

func (r *RFGOR) AtomicStoreBits(mask, b FGOR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFGOR) AtomicSetBits(mask FGOR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFGOR) AtomicClearBits(mask FGOR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFGOR struct{ mmio.UM32 }

func (rm RMFGOR) Load() FGOR   { return FGOR(rm.UM32.Load()) }
func (rm RMFGOR) Store(b FGOR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) LO() RMFGOR {
	return RMFGOR{mmio.UM32{&p.FGOR.U32, uint32(LO)}}
}

type BGMAR uint32

func (b BGMAR) Field(mask BGMAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BGMAR) J(v int) BGMAR {
	return BGMAR(bits.Make32(v, uint32(mask)))
}

type RBGMAR struct{ mmio.U32 }

func (r *RBGMAR) Bits(mask BGMAR) BGMAR   { return BGMAR(r.U32.Bits(uint32(mask))) }
func (r *RBGMAR) StoreBits(mask, b BGMAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBGMAR) SetBits(mask BGMAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBGMAR) ClearBits(mask BGMAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBGMAR) Load() BGMAR             { return BGMAR(r.U32.Load()) }
func (r *RBGMAR) Store(b BGMAR)           { r.U32.Store(uint32(b)) }

func (r *RBGMAR) AtomicStoreBits(mask, b BGMAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBGMAR) AtomicSetBits(mask BGMAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBGMAR) AtomicClearBits(mask BGMAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBGMAR struct{ mmio.UM32 }

func (rm RMBGMAR) Load() BGMAR   { return BGMAR(rm.UM32.Load()) }
func (rm RMBGMAR) Store(b BGMAR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() RMBGMAR {
	return RMBGMAR{mmio.UM32{&p.BGMAR.U32, uint32(MA)}}
}

type BGOR uint32

func (b BGOR) Field(mask BGOR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BGOR) J(v int) BGOR {
	return BGOR(bits.Make32(v, uint32(mask)))
}

type RBGOR struct{ mmio.U32 }

func (r *RBGOR) Bits(mask BGOR) BGOR    { return BGOR(r.U32.Bits(uint32(mask))) }
func (r *RBGOR) StoreBits(mask, b BGOR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBGOR) SetBits(mask BGOR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBGOR) ClearBits(mask BGOR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBGOR) Load() BGOR             { return BGOR(r.U32.Load()) }
func (r *RBGOR) Store(b BGOR)           { r.U32.Store(uint32(b)) }

func (r *RBGOR) AtomicStoreBits(mask, b BGOR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBGOR) AtomicSetBits(mask BGOR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBGOR) AtomicClearBits(mask BGOR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBGOR struct{ mmio.UM32 }

func (rm RMBGOR) Load() BGOR   { return BGOR(rm.UM32.Load()) }
func (rm RMBGOR) Store(b BGOR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) LO() RMBGOR {
	return RMBGOR{mmio.UM32{&p.BGOR.U32, uint32(LO)}}
}

type FGPFCCR uint32

func (b FGPFCCR) Field(mask FGPFCCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FGPFCCR) J(v int) FGPFCCR {
	return FGPFCCR(bits.Make32(v, uint32(mask)))
}

type RFGPFCCR struct{ mmio.U32 }

func (r *RFGPFCCR) Bits(mask FGPFCCR) FGPFCCR { return FGPFCCR(r.U32.Bits(uint32(mask))) }
func (r *RFGPFCCR) StoreBits(mask, b FGPFCCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFGPFCCR) SetBits(mask FGPFCCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFGPFCCR) ClearBits(mask FGPFCCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFGPFCCR) Load() FGPFCCR             { return FGPFCCR(r.U32.Load()) }
func (r *RFGPFCCR) Store(b FGPFCCR)           { r.U32.Store(uint32(b)) }

func (r *RFGPFCCR) AtomicStoreBits(mask, b FGPFCCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFGPFCCR) AtomicSetBits(mask FGPFCCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFGPFCCR) AtomicClearBits(mask FGPFCCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFGPFCCR struct{ mmio.UM32 }

func (rm RMFGPFCCR) Load() FGPFCCR   { return FGPFCCR(rm.UM32.Load()) }
func (rm RMFGPFCCR) Store(b FGPFCCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) CM() RMFGPFCCR {
	return RMFGPFCCR{mmio.UM32{&p.FGPFCCR.U32, uint32(CM)}}
}

func (p *DMA2D_Periph) CCM() RMFGPFCCR {
	return RMFGPFCCR{mmio.UM32{&p.FGPFCCR.U32, uint32(CCM)}}
}

func (p *DMA2D_Periph) START() RMFGPFCCR {
	return RMFGPFCCR{mmio.UM32{&p.FGPFCCR.U32, uint32(START)}}
}

func (p *DMA2D_Periph) CS() RMFGPFCCR {
	return RMFGPFCCR{mmio.UM32{&p.FGPFCCR.U32, uint32(CS)}}
}

func (p *DMA2D_Periph) AM() RMFGPFCCR {
	return RMFGPFCCR{mmio.UM32{&p.FGPFCCR.U32, uint32(AM)}}
}

func (p *DMA2D_Periph) ALPHA() RMFGPFCCR {
	return RMFGPFCCR{mmio.UM32{&p.FGPFCCR.U32, uint32(ALPHA)}}
}

type FGCOLR uint32

func (b FGCOLR) Field(mask FGCOLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FGCOLR) J(v int) FGCOLR {
	return FGCOLR(bits.Make32(v, uint32(mask)))
}

type RFGCOLR struct{ mmio.U32 }

func (r *RFGCOLR) Bits(mask FGCOLR) FGCOLR  { return FGCOLR(r.U32.Bits(uint32(mask))) }
func (r *RFGCOLR) StoreBits(mask, b FGCOLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFGCOLR) SetBits(mask FGCOLR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFGCOLR) ClearBits(mask FGCOLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFGCOLR) Load() FGCOLR             { return FGCOLR(r.U32.Load()) }
func (r *RFGCOLR) Store(b FGCOLR)           { r.U32.Store(uint32(b)) }

func (r *RFGCOLR) AtomicStoreBits(mask, b FGCOLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFGCOLR) AtomicSetBits(mask FGCOLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFGCOLR) AtomicClearBits(mask FGCOLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFGCOLR struct{ mmio.UM32 }

func (rm RMFGCOLR) Load() FGCOLR   { return FGCOLR(rm.UM32.Load()) }
func (rm RMFGCOLR) Store(b FGCOLR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) BLUE() RMFGCOLR {
	return RMFGCOLR{mmio.UM32{&p.FGCOLR.U32, uint32(BLUE)}}
}

func (p *DMA2D_Periph) GREEN() RMFGCOLR {
	return RMFGCOLR{mmio.UM32{&p.FGCOLR.U32, uint32(GREEN)}}
}

func (p *DMA2D_Periph) RED() RMFGCOLR {
	return RMFGCOLR{mmio.UM32{&p.FGCOLR.U32, uint32(RED)}}
}

type BGPFCCR uint32

func (b BGPFCCR) Field(mask BGPFCCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BGPFCCR) J(v int) BGPFCCR {
	return BGPFCCR(bits.Make32(v, uint32(mask)))
}

type RBGPFCCR struct{ mmio.U32 }

func (r *RBGPFCCR) Bits(mask BGPFCCR) BGPFCCR { return BGPFCCR(r.U32.Bits(uint32(mask))) }
func (r *RBGPFCCR) StoreBits(mask, b BGPFCCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBGPFCCR) SetBits(mask BGPFCCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBGPFCCR) ClearBits(mask BGPFCCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBGPFCCR) Load() BGPFCCR             { return BGPFCCR(r.U32.Load()) }
func (r *RBGPFCCR) Store(b BGPFCCR)           { r.U32.Store(uint32(b)) }

func (r *RBGPFCCR) AtomicStoreBits(mask, b BGPFCCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBGPFCCR) AtomicSetBits(mask BGPFCCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBGPFCCR) AtomicClearBits(mask BGPFCCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBGPFCCR struct{ mmio.UM32 }

func (rm RMBGPFCCR) Load() BGPFCCR   { return BGPFCCR(rm.UM32.Load()) }
func (rm RMBGPFCCR) Store(b BGPFCCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) CM() RMBGPFCCR {
	return RMBGPFCCR{mmio.UM32{&p.BGPFCCR.U32, uint32(CM)}}
}

func (p *DMA2D_Periph) CCM() RMBGPFCCR {
	return RMBGPFCCR{mmio.UM32{&p.BGPFCCR.U32, uint32(CCM)}}
}

func (p *DMA2D_Periph) START() RMBGPFCCR {
	return RMBGPFCCR{mmio.UM32{&p.BGPFCCR.U32, uint32(START)}}
}

func (p *DMA2D_Periph) CS() RMBGPFCCR {
	return RMBGPFCCR{mmio.UM32{&p.BGPFCCR.U32, uint32(CS)}}
}

func (p *DMA2D_Periph) AM() RMBGPFCCR {
	return RMBGPFCCR{mmio.UM32{&p.BGPFCCR.U32, uint32(AM)}}
}

func (p *DMA2D_Periph) ALPHA() RMBGPFCCR {
	return RMBGPFCCR{mmio.UM32{&p.BGPFCCR.U32, uint32(ALPHA)}}
}

type BGCOLR uint32

func (b BGCOLR) Field(mask BGCOLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BGCOLR) J(v int) BGCOLR {
	return BGCOLR(bits.Make32(v, uint32(mask)))
}

type RBGCOLR struct{ mmio.U32 }

func (r *RBGCOLR) Bits(mask BGCOLR) BGCOLR  { return BGCOLR(r.U32.Bits(uint32(mask))) }
func (r *RBGCOLR) StoreBits(mask, b BGCOLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBGCOLR) SetBits(mask BGCOLR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBGCOLR) ClearBits(mask BGCOLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBGCOLR) Load() BGCOLR             { return BGCOLR(r.U32.Load()) }
func (r *RBGCOLR) Store(b BGCOLR)           { r.U32.Store(uint32(b)) }

func (r *RBGCOLR) AtomicStoreBits(mask, b BGCOLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBGCOLR) AtomicSetBits(mask BGCOLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBGCOLR) AtomicClearBits(mask BGCOLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBGCOLR struct{ mmio.UM32 }

func (rm RMBGCOLR) Load() BGCOLR   { return BGCOLR(rm.UM32.Load()) }
func (rm RMBGCOLR) Store(b BGCOLR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) BLUE() RMBGCOLR {
	return RMBGCOLR{mmio.UM32{&p.BGCOLR.U32, uint32(BLUE)}}
}

func (p *DMA2D_Periph) GREEN() RMBGCOLR {
	return RMBGCOLR{mmio.UM32{&p.BGCOLR.U32, uint32(GREEN)}}
}

func (p *DMA2D_Periph) RED() RMBGCOLR {
	return RMBGCOLR{mmio.UM32{&p.BGCOLR.U32, uint32(RED)}}
}

type FGCMAR uint32

func (b FGCMAR) Field(mask FGCMAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FGCMAR) J(v int) FGCMAR {
	return FGCMAR(bits.Make32(v, uint32(mask)))
}

type RFGCMAR struct{ mmio.U32 }

func (r *RFGCMAR) Bits(mask FGCMAR) FGCMAR  { return FGCMAR(r.U32.Bits(uint32(mask))) }
func (r *RFGCMAR) StoreBits(mask, b FGCMAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFGCMAR) SetBits(mask FGCMAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFGCMAR) ClearBits(mask FGCMAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFGCMAR) Load() FGCMAR             { return FGCMAR(r.U32.Load()) }
func (r *RFGCMAR) Store(b FGCMAR)           { r.U32.Store(uint32(b)) }

func (r *RFGCMAR) AtomicStoreBits(mask, b FGCMAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFGCMAR) AtomicSetBits(mask FGCMAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFGCMAR) AtomicClearBits(mask FGCMAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFGCMAR struct{ mmio.UM32 }

func (rm RMFGCMAR) Load() FGCMAR   { return FGCMAR(rm.UM32.Load()) }
func (rm RMFGCMAR) Store(b FGCMAR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() RMFGCMAR {
	return RMFGCMAR{mmio.UM32{&p.FGCMAR.U32, uint32(MA)}}
}

type BGCMAR uint32

func (b BGCMAR) Field(mask BGCMAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BGCMAR) J(v int) BGCMAR {
	return BGCMAR(bits.Make32(v, uint32(mask)))
}

type RBGCMAR struct{ mmio.U32 }

func (r *RBGCMAR) Bits(mask BGCMAR) BGCMAR  { return BGCMAR(r.U32.Bits(uint32(mask))) }
func (r *RBGCMAR) StoreBits(mask, b BGCMAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBGCMAR) SetBits(mask BGCMAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBGCMAR) ClearBits(mask BGCMAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBGCMAR) Load() BGCMAR             { return BGCMAR(r.U32.Load()) }
func (r *RBGCMAR) Store(b BGCMAR)           { r.U32.Store(uint32(b)) }

func (r *RBGCMAR) AtomicStoreBits(mask, b BGCMAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBGCMAR) AtomicSetBits(mask BGCMAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBGCMAR) AtomicClearBits(mask BGCMAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBGCMAR struct{ mmio.UM32 }

func (rm RMBGCMAR) Load() BGCMAR   { return BGCMAR(rm.UM32.Load()) }
func (rm RMBGCMAR) Store(b BGCMAR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() RMBGCMAR {
	return RMBGCMAR{mmio.UM32{&p.BGCMAR.U32, uint32(MA)}}
}

type OPFCCR uint32

func (b OPFCCR) Field(mask OPFCCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPFCCR) J(v int) OPFCCR {
	return OPFCCR(bits.Make32(v, uint32(mask)))
}

type ROPFCCR struct{ mmio.U32 }

func (r *ROPFCCR) Bits(mask OPFCCR) OPFCCR  { return OPFCCR(r.U32.Bits(uint32(mask))) }
func (r *ROPFCCR) StoreBits(mask, b OPFCCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROPFCCR) SetBits(mask OPFCCR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROPFCCR) ClearBits(mask OPFCCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROPFCCR) Load() OPFCCR             { return OPFCCR(r.U32.Load()) }
func (r *ROPFCCR) Store(b OPFCCR)           { r.U32.Store(uint32(b)) }

func (r *ROPFCCR) AtomicStoreBits(mask, b OPFCCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROPFCCR) AtomicSetBits(mask OPFCCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROPFCCR) AtomicClearBits(mask OPFCCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOPFCCR struct{ mmio.UM32 }

func (rm RMOPFCCR) Load() OPFCCR   { return OPFCCR(rm.UM32.Load()) }
func (rm RMOPFCCR) Store(b OPFCCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) CM() RMOPFCCR {
	return RMOPFCCR{mmio.UM32{&p.OPFCCR.U32, uint32(CM)}}
}

type OCOLR uint32

func (b OCOLR) Field(mask OCOLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OCOLR) J(v int) OCOLR {
	return OCOLR(bits.Make32(v, uint32(mask)))
}

type ROCOLR struct{ mmio.U32 }

func (r *ROCOLR) Bits(mask OCOLR) OCOLR   { return OCOLR(r.U32.Bits(uint32(mask))) }
func (r *ROCOLR) StoreBits(mask, b OCOLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROCOLR) SetBits(mask OCOLR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROCOLR) ClearBits(mask OCOLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROCOLR) Load() OCOLR             { return OCOLR(r.U32.Load()) }
func (r *ROCOLR) Store(b OCOLR)           { r.U32.Store(uint32(b)) }

func (r *ROCOLR) AtomicStoreBits(mask, b OCOLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROCOLR) AtomicSetBits(mask OCOLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROCOLR) AtomicClearBits(mask OCOLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOCOLR struct{ mmio.UM32 }

func (rm RMOCOLR) Load() OCOLR   { return OCOLR(rm.UM32.Load()) }
func (rm RMOCOLR) Store(b OCOLR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) BLUE_1() RMOCOLR {
	return RMOCOLR{mmio.UM32{&p.OCOLR.U32, uint32(BLUE_1)}}
}

func (p *DMA2D_Periph) GREEN_1() RMOCOLR {
	return RMOCOLR{mmio.UM32{&p.OCOLR.U32, uint32(GREEN_1)}}
}

func (p *DMA2D_Periph) RED_1() RMOCOLR {
	return RMOCOLR{mmio.UM32{&p.OCOLR.U32, uint32(RED_1)}}
}

func (p *DMA2D_Periph) ALPHA_1() RMOCOLR {
	return RMOCOLR{mmio.UM32{&p.OCOLR.U32, uint32(ALPHA_1)}}
}

type OMAR uint32

func (b OMAR) Field(mask OMAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OMAR) J(v int) OMAR {
	return OMAR(bits.Make32(v, uint32(mask)))
}

type ROMAR struct{ mmio.U32 }

func (r *ROMAR) Bits(mask OMAR) OMAR    { return OMAR(r.U32.Bits(uint32(mask))) }
func (r *ROMAR) StoreBits(mask, b OMAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROMAR) SetBits(mask OMAR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROMAR) ClearBits(mask OMAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROMAR) Load() OMAR             { return OMAR(r.U32.Load()) }
func (r *ROMAR) Store(b OMAR)           { r.U32.Store(uint32(b)) }

func (r *ROMAR) AtomicStoreBits(mask, b OMAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROMAR) AtomicSetBits(mask OMAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROMAR) AtomicClearBits(mask OMAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOMAR struct{ mmio.UM32 }

func (rm RMOMAR) Load() OMAR   { return OMAR(rm.UM32.Load()) }
func (rm RMOMAR) Store(b OMAR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() RMOMAR {
	return RMOMAR{mmio.UM32{&p.OMAR.U32, uint32(MA)}}
}

type OOR uint32

func (b OOR) Field(mask OOR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OOR) J(v int) OOR {
	return OOR(bits.Make32(v, uint32(mask)))
}

type ROOR struct{ mmio.U32 }

func (r *ROOR) Bits(mask OOR) OOR     { return OOR(r.U32.Bits(uint32(mask))) }
func (r *ROOR) StoreBits(mask, b OOR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROOR) SetBits(mask OOR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROOR) ClearBits(mask OOR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROOR) Load() OOR             { return OOR(r.U32.Load()) }
func (r *ROOR) Store(b OOR)           { r.U32.Store(uint32(b)) }

func (r *ROOR) AtomicStoreBits(mask, b OOR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROOR) AtomicSetBits(mask OOR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROOR) AtomicClearBits(mask OOR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOOR struct{ mmio.UM32 }

func (rm RMOOR) Load() OOR   { return OOR(rm.UM32.Load()) }
func (rm RMOOR) Store(b OOR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) LO() RMOOR {
	return RMOOR{mmio.UM32{&p.OOR.U32, uint32(LO)}}
}

type NLR uint32

func (b NLR) Field(mask NLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NLR) J(v int) NLR {
	return NLR(bits.Make32(v, uint32(mask)))
}

type RNLR struct{ mmio.U32 }

func (r *RNLR) Bits(mask NLR) NLR     { return NLR(r.U32.Bits(uint32(mask))) }
func (r *RNLR) StoreBits(mask, b NLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RNLR) SetBits(mask NLR)      { r.U32.SetBits(uint32(mask)) }
func (r *RNLR) ClearBits(mask NLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RNLR) Load() NLR             { return NLR(r.U32.Load()) }
func (r *RNLR) Store(b NLR)           { r.U32.Store(uint32(b)) }

func (r *RNLR) AtomicStoreBits(mask, b NLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RNLR) AtomicSetBits(mask NLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RNLR) AtomicClearBits(mask NLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMNLR struct{ mmio.UM32 }

func (rm RMNLR) Load() NLR   { return NLR(rm.UM32.Load()) }
func (rm RMNLR) Store(b NLR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) NL() RMNLR {
	return RMNLR{mmio.UM32{&p.NLR.U32, uint32(NL)}}
}

func (p *DMA2D_Periph) PL() RMNLR {
	return RMNLR{mmio.UM32{&p.NLR.U32, uint32(PL)}}
}

type LWR uint32

func (b LWR) Field(mask LWR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LWR) J(v int) LWR {
	return LWR(bits.Make32(v, uint32(mask)))
}

type RLWR struct{ mmio.U32 }

func (r *RLWR) Bits(mask LWR) LWR     { return LWR(r.U32.Bits(uint32(mask))) }
func (r *RLWR) StoreBits(mask, b LWR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLWR) SetBits(mask LWR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLWR) ClearBits(mask LWR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLWR) Load() LWR             { return LWR(r.U32.Load()) }
func (r *RLWR) Store(b LWR)           { r.U32.Store(uint32(b)) }

func (r *RLWR) AtomicStoreBits(mask, b LWR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLWR) AtomicSetBits(mask LWR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLWR) AtomicClearBits(mask LWR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLWR struct{ mmio.UM32 }

func (rm RMLWR) Load() LWR   { return LWR(rm.UM32.Load()) }
func (rm RMLWR) Store(b LWR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) LW() RMLWR {
	return RMLWR{mmio.UM32{&p.LWR.U32, uint32(LW)}}
}

type AMTCR uint32

func (b AMTCR) Field(mask AMTCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AMTCR) J(v int) AMTCR {
	return AMTCR(bits.Make32(v, uint32(mask)))
}

type RAMTCR struct{ mmio.U32 }

func (r *RAMTCR) Bits(mask AMTCR) AMTCR   { return AMTCR(r.U32.Bits(uint32(mask))) }
func (r *RAMTCR) StoreBits(mask, b AMTCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAMTCR) SetBits(mask AMTCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAMTCR) ClearBits(mask AMTCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAMTCR) Load() AMTCR             { return AMTCR(r.U32.Load()) }
func (r *RAMTCR) Store(b AMTCR)           { r.U32.Store(uint32(b)) }

func (r *RAMTCR) AtomicStoreBits(mask, b AMTCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAMTCR) AtomicSetBits(mask AMTCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAMTCR) AtomicClearBits(mask AMTCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAMTCR struct{ mmio.UM32 }

func (rm RMAMTCR) Load() AMTCR   { return AMTCR(rm.UM32.Load()) }
func (rm RMAMTCR) Store(b AMTCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) EN() RMAMTCR {
	return RMAMTCR{mmio.UM32{&p.AMTCR.U32, uint32(EN)}}
}

func (p *DMA2D_Periph) DT() RMAMTCR {
	return RMAMTCR{mmio.UM32{&p.AMTCR.U32, uint32(DT)}}
}

type FGCLUT uint32

func (b FGCLUT) Field(mask FGCLUT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FGCLUT) J(v int) FGCLUT {
	return FGCLUT(bits.Make32(v, uint32(mask)))
}

type RFGCLUT struct{ mmio.U32 }

func (r *RFGCLUT) Bits(mask FGCLUT) FGCLUT  { return FGCLUT(r.U32.Bits(uint32(mask))) }
func (r *RFGCLUT) StoreBits(mask, b FGCLUT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFGCLUT) SetBits(mask FGCLUT)      { r.U32.SetBits(uint32(mask)) }
func (r *RFGCLUT) ClearBits(mask FGCLUT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFGCLUT) Load() FGCLUT             { return FGCLUT(r.U32.Load()) }
func (r *RFGCLUT) Store(b FGCLUT)           { r.U32.Store(uint32(b)) }

func (r *RFGCLUT) AtomicStoreBits(mask, b FGCLUT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFGCLUT) AtomicSetBits(mask FGCLUT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFGCLUT) AtomicClearBits(mask FGCLUT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFGCLUT struct{ mmio.UM32 }

func (rm RMFGCLUT) Load() FGCLUT   { return FGCLUT(rm.UM32.Load()) }
func (rm RMFGCLUT) Store(b FGCLUT) { rm.UM32.Store(uint32(b)) }

type BGCLUT uint32

func (b BGCLUT) Field(mask BGCLUT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BGCLUT) J(v int) BGCLUT {
	return BGCLUT(bits.Make32(v, uint32(mask)))
}

type RBGCLUT struct{ mmio.U32 }

func (r *RBGCLUT) Bits(mask BGCLUT) BGCLUT  { return BGCLUT(r.U32.Bits(uint32(mask))) }
func (r *RBGCLUT) StoreBits(mask, b BGCLUT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBGCLUT) SetBits(mask BGCLUT)      { r.U32.SetBits(uint32(mask)) }
func (r *RBGCLUT) ClearBits(mask BGCLUT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBGCLUT) Load() BGCLUT             { return BGCLUT(r.U32.Load()) }
func (r *RBGCLUT) Store(b BGCLUT)           { r.U32.Store(uint32(b)) }

func (r *RBGCLUT) AtomicStoreBits(mask, b BGCLUT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBGCLUT) AtomicSetBits(mask BGCLUT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBGCLUT) AtomicClearBits(mask BGCLUT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBGCLUT struct{ mmio.UM32 }

func (rm RMBGCLUT) Load() BGCLUT   { return BGCLUT(rm.UM32.Load()) }
func (rm RMBGCLUT) Store(b BGCLUT) { rm.UM32.Store(uint32(b)) }
