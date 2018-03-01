package usart

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x6/mmap"
)

type USART_Periph struct {
	CR1  RCR1
	CR2  RCR2
	CR3  RCR3
	BRR  RBRR
	GTPR RGTPR
	RTOR RRTOR
	RQR  RRQR
	ISR  RISR
	ICR  RICR
	RDR  RRDR
	_    uint16
	TDR  RTDR
}

func (p *USART_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var USART1 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART1_BASE)))

type CR1 uint32

func (b CR1) Field(mask CR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1) J(v int) CR1 {
	return CR1(bits.Make32(v, uint32(mask)))
}

type RCR1 struct{ mmio.U32 }

func (r *RCR1) Bits(mask CR1) CR1     { return CR1(r.U32.Bits(uint32(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U32.Load()) }
func (r *RCR1) Store(b CR1)           { r.U32.Store(uint32(b)) }

func (r *RCR1) AtomicStoreBits(mask, b CR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) AtomicSetBits(mask CR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR1) AtomicClearBits(mask CR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR1 struct{ mmio.UM32 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM32.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) UE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(UE)}}
}

func (p *USART_Periph) RE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RE)}}
}

func (p *USART_Periph) TE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(TE)}}
}

func (p *USART_Periph) IDLEIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(IDLEIE)}}
}

func (p *USART_Periph) RXNEIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RXNEIE)}}
}

func (p *USART_Periph) TCIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(TCIE)}}
}

func (p *USART_Periph) TXEIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(TXEIE)}}
}

func (p *USART_Periph) PEIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(PEIE)}}
}

func (p *USART_Periph) PS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(PS)}}
}

func (p *USART_Periph) PCE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(PCE)}}
}

func (p *USART_Periph) WAKE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(WAKE)}}
}

func (p *USART_Periph) M() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(M)}}
}

func (p *USART_Periph) MME() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(MME)}}
}

func (p *USART_Periph) CMIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CMIE)}}
}

func (p *USART_Periph) OVER8() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(OVER8)}}
}

func (p *USART_Periph) DEDT() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DEDT)}}
}

func (p *USART_Periph) DEAT() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DEAT)}}
}

func (p *USART_Periph) RTOIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RTOIE)}}
}

func (p *USART_Periph) EOBIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(EOBIE)}}
}

type CR2 uint32

func (b CR2) Field(mask CR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2) J(v int) CR2 {
	return CR2(bits.Make32(v, uint32(mask)))
}

type RCR2 struct{ mmio.U32 }

func (r *RCR2) Bits(mask CR2) CR2     { return CR2(r.U32.Bits(uint32(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U32.Load()) }
func (r *RCR2) Store(b CR2)           { r.U32.Store(uint32(b)) }

func (r *RCR2) AtomicStoreBits(mask, b CR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) AtomicSetBits(mask CR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR2) AtomicClearBits(mask CR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR2 struct{ mmio.UM32 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM32.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) ADDM7() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ADDM7)}}
}

func (p *USART_Periph) LBCL() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(LBCL)}}
}

func (p *USART_Periph) CPHA() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(CPHA)}}
}

func (p *USART_Periph) CPOL() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(CPOL)}}
}

func (p *USART_Periph) CLKEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(CLKEN)}}
}

func (p *USART_Periph) STOP() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(STOP)}}
}

func (p *USART_Periph) SWAP() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(SWAP)}}
}

func (p *USART_Periph) RXINV() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(RXINV)}}
}

func (p *USART_Periph) TXINV() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TXINV)}}
}

func (p *USART_Periph) DATAINV() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(DATAINV)}}
}

func (p *USART_Periph) MSBFIRST() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(MSBFIRST)}}
}

func (p *USART_Periph) ABREN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ABREN)}}
}

func (p *USART_Periph) ABRMODE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ABRMODE)}}
}

func (p *USART_Periph) RTOEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(RTOEN)}}
}

func (p *USART_Periph) ADD() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ADD)}}
}

type CR3 uint32

func (b CR3) Field(mask CR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR3) J(v int) CR3 {
	return CR3(bits.Make32(v, uint32(mask)))
}

type RCR3 struct{ mmio.U32 }

func (r *RCR3) Bits(mask CR3) CR3     { return CR3(r.U32.Bits(uint32(mask))) }
func (r *RCR3) StoreBits(mask, b CR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR3) SetBits(mask CR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR3) ClearBits(mask CR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR3) Load() CR3             { return CR3(r.U32.Load()) }
func (r *RCR3) Store(b CR3)           { r.U32.Store(uint32(b)) }

func (r *RCR3) AtomicStoreBits(mask, b CR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR3) AtomicSetBits(mask CR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR3) AtomicClearBits(mask CR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR3 struct{ mmio.UM32 }

func (rm RMCR3) Load() CR3   { return CR3(rm.UM32.Load()) }
func (rm RMCR3) Store(b CR3) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) EIE() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EIE)}}
}

func (p *USART_Periph) HDSEL() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(HDSEL)}}
}

func (p *USART_Periph) DMAR() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(DMAR)}}
}

func (p *USART_Periph) DMAT() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(DMAT)}}
}

func (p *USART_Periph) RTSE() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(RTSE)}}
}

func (p *USART_Periph) CTSE() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(CTSE)}}
}

func (p *USART_Periph) CTSIE() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(CTSIE)}}
}

func (p *USART_Periph) ONEBIT() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(ONEBIT)}}
}

func (p *USART_Periph) OVRDIS() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(OVRDIS)}}
}

func (p *USART_Periph) DDRE() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(DDRE)}}
}

func (p *USART_Periph) DEM() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(DEM)}}
}

func (p *USART_Periph) DEP() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(DEP)}}
}

type BRR uint32

func (b BRR) Field(mask BRR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BRR) J(v int) BRR {
	return BRR(bits.Make32(v, uint32(mask)))
}

type RBRR struct{ mmio.U32 }

func (r *RBRR) Bits(mask BRR) BRR     { return BRR(r.U32.Bits(uint32(mask))) }
func (r *RBRR) StoreBits(mask, b BRR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBRR) SetBits(mask BRR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBRR) ClearBits(mask BRR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBRR) Load() BRR             { return BRR(r.U32.Load()) }
func (r *RBRR) Store(b BRR)           { r.U32.Store(uint32(b)) }

func (r *RBRR) AtomicStoreBits(mask, b BRR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBRR) AtomicSetBits(mask BRR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBRR) AtomicClearBits(mask BRR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBRR struct{ mmio.UM32 }

func (rm RMBRR) Load() BRR   { return BRR(rm.UM32.Load()) }
func (rm RMBRR) Store(b BRR) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) DIV_FRACTION() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(DIV_FRACTION)}}
}

func (p *USART_Periph) DIV_MANTISSA() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(DIV_MANTISSA)}}
}

type GTPR uint32

func (b GTPR) Field(mask GTPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask GTPR) J(v int) GTPR {
	return GTPR(bits.Make32(v, uint32(mask)))
}

type RGTPR struct{ mmio.U32 }

func (r *RGTPR) Bits(mask GTPR) GTPR    { return GTPR(r.U32.Bits(uint32(mask))) }
func (r *RGTPR) StoreBits(mask, b GTPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RGTPR) SetBits(mask GTPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RGTPR) ClearBits(mask GTPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RGTPR) Load() GTPR             { return GTPR(r.U32.Load()) }
func (r *RGTPR) Store(b GTPR)           { r.U32.Store(uint32(b)) }

func (r *RGTPR) AtomicStoreBits(mask, b GTPR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RGTPR) AtomicSetBits(mask GTPR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RGTPR) AtomicClearBits(mask GTPR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMGTPR struct{ mmio.UM32 }

func (rm RMGTPR) Load() GTPR   { return GTPR(rm.UM32.Load()) }
func (rm RMGTPR) Store(b GTPR) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) PSC() RMGTPR {
	return RMGTPR{mmio.UM32{&p.GTPR.U32, uint32(PSC)}}
}

func (p *USART_Periph) GT() RMGTPR {
	return RMGTPR{mmio.UM32{&p.GTPR.U32, uint32(GT)}}
}

type RTOR uint32

func (b RTOR) Field(mask RTOR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RTOR) J(v int) RTOR {
	return RTOR(bits.Make32(v, uint32(mask)))
}

type RRTOR struct{ mmio.U32 }

func (r *RRTOR) Bits(mask RTOR) RTOR    { return RTOR(r.U32.Bits(uint32(mask))) }
func (r *RRTOR) StoreBits(mask, b RTOR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRTOR) SetBits(mask RTOR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRTOR) ClearBits(mask RTOR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRTOR) Load() RTOR             { return RTOR(r.U32.Load()) }
func (r *RRTOR) Store(b RTOR)           { r.U32.Store(uint32(b)) }

func (r *RRTOR) AtomicStoreBits(mask, b RTOR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRTOR) AtomicSetBits(mask RTOR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRTOR) AtomicClearBits(mask RTOR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRTOR struct{ mmio.UM32 }

func (rm RMRTOR) Load() RTOR   { return RTOR(rm.UM32.Load()) }
func (rm RMRTOR) Store(b RTOR) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) RTO() RMRTOR {
	return RMRTOR{mmio.UM32{&p.RTOR.U32, uint32(RTO)}}
}

func (p *USART_Periph) BLEN() RMRTOR {
	return RMRTOR{mmio.UM32{&p.RTOR.U32, uint32(BLEN)}}
}

type RQR uint32

func (b RQR) Field(mask RQR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RQR) J(v int) RQR {
	return RQR(bits.Make32(v, uint32(mask)))
}

type RRQR struct{ mmio.U32 }

func (r *RRQR) Bits(mask RQR) RQR     { return RQR(r.U32.Bits(uint32(mask))) }
func (r *RRQR) StoreBits(mask, b RQR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRQR) SetBits(mask RQR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRQR) ClearBits(mask RQR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRQR) Load() RQR             { return RQR(r.U32.Load()) }
func (r *RRQR) Store(b RQR)           { r.U32.Store(uint32(b)) }

func (r *RRQR) AtomicStoreBits(mask, b RQR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRQR) AtomicSetBits(mask RQR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRQR) AtomicClearBits(mask RQR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRQR struct{ mmio.UM32 }

func (rm RMRQR) Load() RQR   { return RQR(rm.UM32.Load()) }
func (rm RMRQR) Store(b RQR) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) ABRRQ() RMRQR {
	return RMRQR{mmio.UM32{&p.RQR.U32, uint32(ABRRQ)}}
}

func (p *USART_Periph) SBKRQ() RMRQR {
	return RMRQR{mmio.UM32{&p.RQR.U32, uint32(SBKRQ)}}
}

func (p *USART_Periph) MMRQ() RMRQR {
	return RMRQR{mmio.UM32{&p.RQR.U32, uint32(MMRQ)}}
}

func (p *USART_Periph) RXFRQ() RMRQR {
	return RMRQR{mmio.UM32{&p.RQR.U32, uint32(RXFRQ)}}
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

func (p *USART_Periph) PE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(PE)}}
}

func (p *USART_Periph) FE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(FE)}}
}

func (p *USART_Periph) NE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(NE)}}
}

func (p *USART_Periph) ORE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ORE)}}
}

func (p *USART_Periph) IDLE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(IDLE)}}
}

func (p *USART_Periph) RXNE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXNE)}}
}

func (p *USART_Periph) TC() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TC)}}
}

func (p *USART_Periph) TXE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXE)}}
}

func (p *USART_Periph) CTSIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(CTSIF)}}
}

func (p *USART_Periph) CTS() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(CTS)}}
}

func (p *USART_Periph) RTOF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RTOF)}}
}

func (p *USART_Periph) ABRE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ABRE)}}
}

func (p *USART_Periph) ABRF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ABRF)}}
}

func (p *USART_Periph) BUSY() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(BUSY)}}
}

func (p *USART_Periph) CMF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(CMF)}}
}

func (p *USART_Periph) SBKF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(SBKF)}}
}

func (p *USART_Periph) RWU() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RWU)}}
}

func (p *USART_Periph) TEACK() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TEACK)}}
}

func (p *USART_Periph) REACK() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(REACK)}}
}

type ICR uint32

func (b ICR) Field(mask ICR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR) J(v int) ICR {
	return ICR(bits.Make32(v, uint32(mask)))
}

type RICR struct{ mmio.U32 }

func (r *RICR) Bits(mask ICR) ICR     { return ICR(r.U32.Bits(uint32(mask))) }
func (r *RICR) StoreBits(mask, b ICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICR) SetBits(mask ICR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICR) ClearBits(mask ICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICR) Load() ICR             { return ICR(r.U32.Load()) }
func (r *RICR) Store(b ICR)           { r.U32.Store(uint32(b)) }

func (r *RICR) AtomicStoreBits(mask, b ICR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RICR) AtomicSetBits(mask ICR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RICR) AtomicClearBits(mask ICR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMICR struct{ mmio.UM32 }

func (rm RMICR) Load() ICR   { return ICR(rm.UM32.Load()) }
func (rm RMICR) Store(b ICR) { rm.UM32.Store(uint32(b)) }

func (p *USART_Periph) PECF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(PECF)}}
}

func (p *USART_Periph) FECF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(FECF)}}
}

func (p *USART_Periph) NCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(NCF)}}
}

func (p *USART_Periph) ORECF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(ORECF)}}
}

func (p *USART_Periph) IDLECF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(IDLECF)}}
}

func (p *USART_Periph) TCCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(TCCF)}}
}

func (p *USART_Periph) CTSCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CTSCF)}}
}

func (p *USART_Periph) RTOCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(RTOCF)}}
}

func (p *USART_Periph) CMCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CMCF)}}
}

type RDR uint16

func (b RDR) Field(mask RDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDR) J(v int) RDR {
	return RDR(bits.Make32(v, uint32(mask)))
}

type RRDR struct{ mmio.U16 }

func (r *RRDR) Bits(mask RDR) RDR     { return RDR(r.U16.Bits(uint16(mask))) }
func (r *RRDR) StoreBits(mask, b RDR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RRDR) SetBits(mask RDR)      { r.U16.SetBits(uint16(mask)) }
func (r *RRDR) ClearBits(mask RDR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RRDR) Load() RDR             { return RDR(r.U16.Load()) }
func (r *RRDR) Store(b RDR)           { r.U16.Store(uint16(b)) }

type RMRDR struct{ mmio.UM16 }

func (rm RMRDR) Load() RDR   { return RDR(rm.UM16.Load()) }
func (rm RMRDR) Store(b RDR) { rm.UM16.Store(uint16(b)) }

type TDR uint16

func (b TDR) Field(mask TDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TDR) J(v int) TDR {
	return TDR(bits.Make32(v, uint32(mask)))
}

type RTDR struct{ mmio.U16 }

func (r *RTDR) Bits(mask TDR) TDR     { return TDR(r.U16.Bits(uint16(mask))) }
func (r *RTDR) StoreBits(mask, b TDR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RTDR) SetBits(mask TDR)      { r.U16.SetBits(uint16(mask)) }
func (r *RTDR) ClearBits(mask TDR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RTDR) Load() TDR             { return TDR(r.U16.Load()) }
func (r *RTDR) Store(b TDR)           { r.U16.Store(uint16(b)) }

type RMTDR struct{ mmio.UM16 }

func (rm RMTDR) Load() TDR   { return TDR(rm.UM16.Load()) }
func (rm RMTDR) Store(b TDR) { rm.UM16.Store(uint16(b)) }
