// +build l476xx

package swpmi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type SWPMI_Periph struct {
	CR  RCR
	BRR RBRR
	_   uint32
	ISR RISR
	ICR RICR
	IER RIER
	RFL RRFL
	TDR RTDR
	RDR RRDR
	OR  ROR
}

func (p *SWPMI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SWPMI1 = (*SWPMI_Periph)(unsafe.Pointer(uintptr(mmap.SWPMI1_BASE)))

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

func (p *SWPMI_Periph) RXDMA() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(RXDMA)}}
}

func (p *SWPMI_Periph) TXDMA() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TXDMA)}}
}

func (p *SWPMI_Periph) RXMODE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(RXMODE)}}
}

func (p *SWPMI_Periph) TXMODE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TXMODE)}}
}

func (p *SWPMI_Periph) LPBK() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LPBK)}}
}

func (p *SWPMI_Periph) SWPACT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SWPACT)}}
}

func (p *SWPMI_Periph) DEACT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DEACT)}}
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

func (p *SWPMI_Periph) BR() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR)}}
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

func (p *SWPMI_Periph) RXBFF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXBFF)}}
}

func (p *SWPMI_Periph) TXBEF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXBEF)}}
}

func (p *SWPMI_Periph) RXBERF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXBERF)}}
}

func (p *SWPMI_Periph) RXOVRF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXOVRF)}}
}

func (p *SWPMI_Periph) TXUNRF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXUNRF)}}
}

func (p *SWPMI_Periph) RXNE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXNE)}}
}

func (p *SWPMI_Periph) TXE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXE)}}
}

func (p *SWPMI_Periph) TCF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCF)}}
}

func (p *SWPMI_Periph) SRF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(SRF)}}
}

func (p *SWPMI_Periph) SUSP() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(SUSP)}}
}

func (p *SWPMI_Periph) DEACTF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(DEACTF)}}
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

func (p *SWPMI_Periph) CRXBFF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CRXBFF)}}
}

func (p *SWPMI_Periph) CTXBEF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CTXBEF)}}
}

func (p *SWPMI_Periph) CRXBERF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CRXBERF)}}
}

func (p *SWPMI_Periph) CRXOVRF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CRXOVRF)}}
}

func (p *SWPMI_Periph) CTXUNRF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CTXUNRF)}}
}

func (p *SWPMI_Periph) CTCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CTCF)}}
}

func (p *SWPMI_Periph) CSRF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CSRF)}}
}

type IER uint32

func (b IER) Field(mask IER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IER) J(v int) IER {
	return IER(bits.Make32(v, uint32(mask)))
}

type RIER struct{ mmio.U32 }

func (r *RIER) Bits(mask IER) IER     { return IER(r.U32.Bits(uint32(mask))) }
func (r *RIER) StoreBits(mask, b IER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIER) SetBits(mask IER)      { r.U32.SetBits(uint32(mask)) }
func (r *RIER) ClearBits(mask IER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIER) Load() IER             { return IER(r.U32.Load()) }
func (r *RIER) Store(b IER)           { r.U32.Store(uint32(b)) }

func (r *RIER) AtomicStoreBits(mask, b IER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIER) AtomicSetBits(mask IER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIER) AtomicClearBits(mask IER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIER struct{ mmio.UM32 }

func (rm RMIER) Load() IER   { return IER(rm.UM32.Load()) }
func (rm RMIER) Store(b IER) { rm.UM32.Store(uint32(b)) }

func (p *SWPMI_Periph) SRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(SRIE)}}
}

func (p *SWPMI_Periph) TCIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TCIE)}}
}

func (p *SWPMI_Periph) TIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TIE)}}
}

func (p *SWPMI_Periph) RIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RIE)}}
}

func (p *SWPMI_Periph) TXUNRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXUNRIE)}}
}

func (p *SWPMI_Periph) RXOVRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RXOVRIE)}}
}

func (p *SWPMI_Periph) RXBERIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RXBERIE)}}
}

func (p *SWPMI_Periph) TXBEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXBEIE)}}
}

func (p *SWPMI_Periph) RXBFIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RXBFIE)}}
}

type RFL uint32

func (b RFL) Field(mask RFL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RFL) J(v int) RFL {
	return RFL(bits.Make32(v, uint32(mask)))
}

type RRFL struct{ mmio.U32 }

func (r *RRFL) Bits(mask RFL) RFL     { return RFL(r.U32.Bits(uint32(mask))) }
func (r *RRFL) StoreBits(mask, b RFL) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRFL) SetBits(mask RFL)      { r.U32.SetBits(uint32(mask)) }
func (r *RRFL) ClearBits(mask RFL)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRFL) Load() RFL             { return RFL(r.U32.Load()) }
func (r *RRFL) Store(b RFL)           { r.U32.Store(uint32(b)) }

func (r *RRFL) AtomicStoreBits(mask, b RFL) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRFL) AtomicSetBits(mask RFL)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRFL) AtomicClearBits(mask RFL)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRFL struct{ mmio.UM32 }

func (rm RMRFL) Load() RFL   { return RFL(rm.UM32.Load()) }
func (rm RMRFL) Store(b RFL) { rm.UM32.Store(uint32(b)) }

func (p *SWPMI_Periph) RFL() RMRFL {
	return RMRFL{mmio.UM32{&p.RFL.U32, uint32(RFL)}}
}

type TDR uint32

func (b TDR) Field(mask TDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TDR) J(v int) TDR {
	return TDR(bits.Make32(v, uint32(mask)))
}

type RTDR struct{ mmio.U32 }

func (r *RTDR) Bits(mask TDR) TDR     { return TDR(r.U32.Bits(uint32(mask))) }
func (r *RTDR) StoreBits(mask, b TDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTDR) SetBits(mask TDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTDR) ClearBits(mask TDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTDR) Load() TDR             { return TDR(r.U32.Load()) }
func (r *RTDR) Store(b TDR)           { r.U32.Store(uint32(b)) }

func (r *RTDR) AtomicStoreBits(mask, b TDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTDR) AtomicSetBits(mask TDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTDR) AtomicClearBits(mask TDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTDR struct{ mmio.UM32 }

func (rm RMTDR) Load() TDR   { return TDR(rm.UM32.Load()) }
func (rm RMTDR) Store(b TDR) { rm.UM32.Store(uint32(b)) }

func (p *SWPMI_Periph) TD() RMTDR {
	return RMTDR{mmio.UM32{&p.TDR.U32, uint32(TD)}}
}

type RDR uint32

func (b RDR) Field(mask RDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDR) J(v int) RDR {
	return RDR(bits.Make32(v, uint32(mask)))
}

type RRDR struct{ mmio.U32 }

func (r *RRDR) Bits(mask RDR) RDR     { return RDR(r.U32.Bits(uint32(mask))) }
func (r *RRDR) StoreBits(mask, b RDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRDR) SetBits(mask RDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRDR) ClearBits(mask RDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRDR) Load() RDR             { return RDR(r.U32.Load()) }
func (r *RRDR) Store(b RDR)           { r.U32.Store(uint32(b)) }

func (r *RRDR) AtomicStoreBits(mask, b RDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRDR) AtomicSetBits(mask RDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRDR) AtomicClearBits(mask RDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRDR struct{ mmio.UM32 }

func (rm RMRDR) Load() RDR   { return RDR(rm.UM32.Load()) }
func (rm RMRDR) Store(b RDR) { rm.UM32.Store(uint32(b)) }

func (p *SWPMI_Periph) RD() RMRDR {
	return RMRDR{mmio.UM32{&p.RDR.U32, uint32(RD)}}
}

type OR uint32

func (b OR) Field(mask OR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OR) J(v int) OR {
	return OR(bits.Make32(v, uint32(mask)))
}

type ROR struct{ mmio.U32 }

func (r *ROR) Bits(mask OR) OR      { return OR(r.U32.Bits(uint32(mask))) }
func (r *ROR) StoreBits(mask, b OR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROR) SetBits(mask OR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROR) ClearBits(mask OR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROR) Load() OR             { return OR(r.U32.Load()) }
func (r *ROR) Store(b OR)           { r.U32.Store(uint32(b)) }

func (r *ROR) AtomicStoreBits(mask, b OR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROR) AtomicSetBits(mask OR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROR) AtomicClearBits(mask OR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOR struct{ mmio.UM32 }

func (rm RMOR) Load() OR   { return OR(rm.UM32.Load()) }
func (rm RMOR) Store(b OR) { rm.UM32.Store(uint32(b)) }

func (p *SWPMI_Periph) TBYP() RMOR {
	return RMOR{mmio.UM32{&p.OR.U32, uint32(TBYP)}}
}

func (p *SWPMI_Periph) CLASS() RMOR {
	return RMOR{mmio.UM32{&p.OR.U32, uint32(CLASS)}}
}
