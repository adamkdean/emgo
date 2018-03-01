// +build l476xx

package quadspi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type QUADSPI_Periph struct {
	CR    RCR
	DCR   RDCR
	SR    RSR
	FCR   RFCR
	DLR   RDLR
	CCR   RCCR
	AR    RAR
	ABR   RABR
	DR    RDR
	PSMKR RPSMKR
	PSMAR RPSMAR
	PIR   RPIR
	LPTR  RLPTR
}

func (p *QUADSPI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var QUADSPI = (*QUADSPI_Periph)(unsafe.Pointer(uintptr(mmap.QSPI_R_BASE)))

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

func (p *QUADSPI_Periph) EN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EN)}}
}

func (p *QUADSPI_Periph) ABORT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ABORT)}}
}

func (p *QUADSPI_Periph) DMAEN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DMAEN)}}
}

func (p *QUADSPI_Periph) TCEN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TCEN)}}
}

func (p *QUADSPI_Periph) SSHIFT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SSHIFT)}}
}

func (p *QUADSPI_Periph) FTHRES() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FTHRES)}}
}

func (p *QUADSPI_Periph) TEIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TEIE)}}
}

func (p *QUADSPI_Periph) TCIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TCIE)}}
}

func (p *QUADSPI_Periph) FTIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FTIE)}}
}

func (p *QUADSPI_Periph) SMIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SMIE)}}
}

func (p *QUADSPI_Periph) TOIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TOIE)}}
}

func (p *QUADSPI_Periph) APMS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(APMS)}}
}

func (p *QUADSPI_Periph) PMM() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PMM)}}
}

func (p *QUADSPI_Periph) PRESCALER() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PRESCALER)}}
}

type DCR uint32

func (b DCR) Field(mask DCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCR) J(v int) DCR {
	return DCR(bits.Make32(v, uint32(mask)))
}

type RDCR struct{ mmio.U32 }

func (r *RDCR) Bits(mask DCR) DCR     { return DCR(r.U32.Bits(uint32(mask))) }
func (r *RDCR) StoreBits(mask, b DCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDCR) SetBits(mask DCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDCR) ClearBits(mask DCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDCR) Load() DCR             { return DCR(r.U32.Load()) }
func (r *RDCR) Store(b DCR)           { r.U32.Store(uint32(b)) }

func (r *RDCR) AtomicStoreBits(mask, b DCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDCR) AtomicSetBits(mask DCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDCR) AtomicClearBits(mask DCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDCR struct{ mmio.UM32 }

func (rm RMDCR) Load() DCR   { return DCR(rm.UM32.Load()) }
func (rm RMDCR) Store(b DCR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) CKMODE() RMDCR {
	return RMDCR{mmio.UM32{&p.DCR.U32, uint32(CKMODE)}}
}

func (p *QUADSPI_Periph) CSHT() RMDCR {
	return RMDCR{mmio.UM32{&p.DCR.U32, uint32(CSHT)}}
}

func (p *QUADSPI_Periph) FSIZE() RMDCR {
	return RMDCR{mmio.UM32{&p.DCR.U32, uint32(FSIZE)}}
}

type SR uint32

func (b SR) Field(mask SR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR) J(v int) SR {
	return SR(bits.Make32(v, uint32(mask)))
}

type RSR struct{ mmio.U32 }

func (r *RSR) Bits(mask SR) SR      { return SR(r.U32.Bits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

func (r *RSR) AtomicStoreBits(mask, b SR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR) AtomicSetBits(mask SR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR) AtomicClearBits(mask SR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) TEF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TEF)}}
}

func (p *QUADSPI_Periph) TCF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TCF)}}
}

func (p *QUADSPI_Periph) FTF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FTF)}}
}

func (p *QUADSPI_Periph) SMF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(SMF)}}
}

func (p *QUADSPI_Periph) TOF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TOF)}}
}

func (p *QUADSPI_Periph) BUSY() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(BUSY)}}
}

func (p *QUADSPI_Periph) FLEVEL() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FLEVEL)}}
}

type FCR uint32

func (b FCR) Field(mask FCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FCR) J(v int) FCR {
	return FCR(bits.Make32(v, uint32(mask)))
}

type RFCR struct{ mmio.U32 }

func (r *RFCR) Bits(mask FCR) FCR     { return FCR(r.U32.Bits(uint32(mask))) }
func (r *RFCR) StoreBits(mask, b FCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFCR) SetBits(mask FCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFCR) ClearBits(mask FCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFCR) Load() FCR             { return FCR(r.U32.Load()) }
func (r *RFCR) Store(b FCR)           { r.U32.Store(uint32(b)) }

func (r *RFCR) AtomicStoreBits(mask, b FCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFCR) AtomicSetBits(mask FCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFCR) AtomicClearBits(mask FCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFCR struct{ mmio.UM32 }

func (rm RMFCR) Load() FCR   { return FCR(rm.UM32.Load()) }
func (rm RMFCR) Store(b FCR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) CTEF() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(CTEF)}}
}

func (p *QUADSPI_Periph) CTCF() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(CTCF)}}
}

func (p *QUADSPI_Periph) CSMF() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(CSMF)}}
}

func (p *QUADSPI_Periph) CTOF() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(CTOF)}}
}

type DLR uint32

func (b DLR) Field(mask DLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DLR) J(v int) DLR {
	return DLR(bits.Make32(v, uint32(mask)))
}

type RDLR struct{ mmio.U32 }

func (r *RDLR) Bits(mask DLR) DLR     { return DLR(r.U32.Bits(uint32(mask))) }
func (r *RDLR) StoreBits(mask, b DLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDLR) SetBits(mask DLR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDLR) ClearBits(mask DLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDLR) Load() DLR             { return DLR(r.U32.Load()) }
func (r *RDLR) Store(b DLR)           { r.U32.Store(uint32(b)) }

func (r *RDLR) AtomicStoreBits(mask, b DLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDLR) AtomicSetBits(mask DLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDLR) AtomicClearBits(mask DLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDLR struct{ mmio.UM32 }

func (rm RMDLR) Load() DLR   { return DLR(rm.UM32.Load()) }
func (rm RMDLR) Store(b DLR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) DL() RMDLR {
	return RMDLR{mmio.UM32{&p.DLR.U32, uint32(DL)}}
}

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

func (p *QUADSPI_Periph) INSTRUCTION() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(INSTRUCTION)}}
}

func (p *QUADSPI_Periph) IMODE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(IMODE)}}
}

func (p *QUADSPI_Periph) ADMODE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(ADMODE)}}
}

func (p *QUADSPI_Periph) ADSIZE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(ADSIZE)}}
}

func (p *QUADSPI_Periph) ABMODE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(ABMODE)}}
}

func (p *QUADSPI_Periph) ABSIZE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(ABSIZE)}}
}

func (p *QUADSPI_Periph) DCYC() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DCYC)}}
}

func (p *QUADSPI_Periph) DMODE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DMODE)}}
}

func (p *QUADSPI_Periph) FMODE() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(FMODE)}}
}

func (p *QUADSPI_Periph) SIOO() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(SIOO)}}
}

func (p *QUADSPI_Periph) DDRM() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DDRM)}}
}

type AR uint32

func (b AR) Field(mask AR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AR) J(v int) AR {
	return AR(bits.Make32(v, uint32(mask)))
}

type RAR struct{ mmio.U32 }

func (r *RAR) Bits(mask AR) AR      { return AR(r.U32.Bits(uint32(mask))) }
func (r *RAR) StoreBits(mask, b AR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAR) SetBits(mask AR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAR) ClearBits(mask AR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAR) Load() AR             { return AR(r.U32.Load()) }
func (r *RAR) Store(b AR)           { r.U32.Store(uint32(b)) }

func (r *RAR) AtomicStoreBits(mask, b AR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAR) AtomicSetBits(mask AR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAR) AtomicClearBits(mask AR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAR struct{ mmio.UM32 }

func (rm RMAR) Load() AR   { return AR(rm.UM32.Load()) }
func (rm RMAR) Store(b AR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) ADDRESS() RMAR {
	return RMAR{mmio.UM32{&p.AR.U32, uint32(ADDRESS)}}
}

type ABR uint32

func (b ABR) Field(mask ABR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ABR) J(v int) ABR {
	return ABR(bits.Make32(v, uint32(mask)))
}

type RABR struct{ mmio.U32 }

func (r *RABR) Bits(mask ABR) ABR     { return ABR(r.U32.Bits(uint32(mask))) }
func (r *RABR) StoreBits(mask, b ABR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RABR) SetBits(mask ABR)      { r.U32.SetBits(uint32(mask)) }
func (r *RABR) ClearBits(mask ABR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RABR) Load() ABR             { return ABR(r.U32.Load()) }
func (r *RABR) Store(b ABR)           { r.U32.Store(uint32(b)) }

func (r *RABR) AtomicStoreBits(mask, b ABR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RABR) AtomicSetBits(mask ABR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RABR) AtomicClearBits(mask ABR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMABR struct{ mmio.UM32 }

func (rm RMABR) Load() ABR   { return ABR(rm.UM32.Load()) }
func (rm RMABR) Store(b ABR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) ALTERNATE() RMABR {
	return RMABR{mmio.UM32{&p.ABR.U32, uint32(ALTERNATE)}}
}

type DR uint32

func (b DR) Field(mask DR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR) J(v int) DR {
	return DR(bits.Make32(v, uint32(mask)))
}

type RDR struct{ mmio.U32 }

func (r *RDR) Bits(mask DR) DR      { return DR(r.U32.Bits(uint32(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDR) SetBits(mask DR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDR) Load() DR             { return DR(r.U32.Load()) }
func (r *RDR) Store(b DR)           { r.U32.Store(uint32(b)) }

func (r *RDR) AtomicStoreBits(mask, b DR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDR) AtomicSetBits(mask DR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDR) AtomicClearBits(mask DR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDR struct{ mmio.UM32 }

func (rm RMDR) Load() DR   { return DR(rm.UM32.Load()) }
func (rm RMDR) Store(b DR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) DATA() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(DATA)}}
}

type PSMKR uint32

func (b PSMKR) Field(mask PSMKR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PSMKR) J(v int) PSMKR {
	return PSMKR(bits.Make32(v, uint32(mask)))
}

type RPSMKR struct{ mmio.U32 }

func (r *RPSMKR) Bits(mask PSMKR) PSMKR   { return PSMKR(r.U32.Bits(uint32(mask))) }
func (r *RPSMKR) StoreBits(mask, b PSMKR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPSMKR) SetBits(mask PSMKR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPSMKR) ClearBits(mask PSMKR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPSMKR) Load() PSMKR             { return PSMKR(r.U32.Load()) }
func (r *RPSMKR) Store(b PSMKR)           { r.U32.Store(uint32(b)) }

func (r *RPSMKR) AtomicStoreBits(mask, b PSMKR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPSMKR) AtomicSetBits(mask PSMKR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPSMKR) AtomicClearBits(mask PSMKR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPSMKR struct{ mmio.UM32 }

func (rm RMPSMKR) Load() PSMKR   { return PSMKR(rm.UM32.Load()) }
func (rm RMPSMKR) Store(b PSMKR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) MASK() RMPSMKR {
	return RMPSMKR{mmio.UM32{&p.PSMKR.U32, uint32(MASK)}}
}

type PSMAR uint32

func (b PSMAR) Field(mask PSMAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PSMAR) J(v int) PSMAR {
	return PSMAR(bits.Make32(v, uint32(mask)))
}

type RPSMAR struct{ mmio.U32 }

func (r *RPSMAR) Bits(mask PSMAR) PSMAR   { return PSMAR(r.U32.Bits(uint32(mask))) }
func (r *RPSMAR) StoreBits(mask, b PSMAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPSMAR) SetBits(mask PSMAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPSMAR) ClearBits(mask PSMAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPSMAR) Load() PSMAR             { return PSMAR(r.U32.Load()) }
func (r *RPSMAR) Store(b PSMAR)           { r.U32.Store(uint32(b)) }

func (r *RPSMAR) AtomicStoreBits(mask, b PSMAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPSMAR) AtomicSetBits(mask PSMAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPSMAR) AtomicClearBits(mask PSMAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPSMAR struct{ mmio.UM32 }

func (rm RMPSMAR) Load() PSMAR   { return PSMAR(rm.UM32.Load()) }
func (rm RMPSMAR) Store(b PSMAR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) MATCH() RMPSMAR {
	return RMPSMAR{mmio.UM32{&p.PSMAR.U32, uint32(MATCH)}}
}

type PIR uint32

func (b PIR) Field(mask PIR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PIR) J(v int) PIR {
	return PIR(bits.Make32(v, uint32(mask)))
}

type RPIR struct{ mmio.U32 }

func (r *RPIR) Bits(mask PIR) PIR     { return PIR(r.U32.Bits(uint32(mask))) }
func (r *RPIR) StoreBits(mask, b PIR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPIR) SetBits(mask PIR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPIR) ClearBits(mask PIR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPIR) Load() PIR             { return PIR(r.U32.Load()) }
func (r *RPIR) Store(b PIR)           { r.U32.Store(uint32(b)) }

func (r *RPIR) AtomicStoreBits(mask, b PIR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPIR) AtomicSetBits(mask PIR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPIR) AtomicClearBits(mask PIR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPIR struct{ mmio.UM32 }

func (rm RMPIR) Load() PIR   { return PIR(rm.UM32.Load()) }
func (rm RMPIR) Store(b PIR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) INTERVAL() RMPIR {
	return RMPIR{mmio.UM32{&p.PIR.U32, uint32(INTERVAL)}}
}

type LPTR uint32

func (b LPTR) Field(mask LPTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LPTR) J(v int) LPTR {
	return LPTR(bits.Make32(v, uint32(mask)))
}

type RLPTR struct{ mmio.U32 }

func (r *RLPTR) Bits(mask LPTR) LPTR    { return LPTR(r.U32.Bits(uint32(mask))) }
func (r *RLPTR) StoreBits(mask, b LPTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLPTR) SetBits(mask LPTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLPTR) ClearBits(mask LPTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLPTR) Load() LPTR             { return LPTR(r.U32.Load()) }
func (r *RLPTR) Store(b LPTR)           { r.U32.Store(uint32(b)) }

func (r *RLPTR) AtomicStoreBits(mask, b LPTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLPTR) AtomicSetBits(mask LPTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLPTR) AtomicClearBits(mask LPTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLPTR struct{ mmio.UM32 }

func (rm RMLPTR) Load() LPTR   { return LPTR(rm.UM32.Load()) }
func (rm RMLPTR) Store(b LPTR) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) TIMEOUT() RMLPTR {
	return RMLPTR{mmio.UM32{&p.LPTR.U32, uint32(TIMEOUT)}}
}
