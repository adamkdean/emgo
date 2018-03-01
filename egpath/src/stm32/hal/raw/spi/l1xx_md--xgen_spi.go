// +build l1xx_md

package spi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type SPI_Periph struct {
	CR1     RCR1
	_       uint16
	CR2     RCR2
	_       uint16
	SR      RSR
	_       uint16
	DR      RDR
	_       uint16
	CRCPR   RCRCPR
	_       uint16
	RXCRCR  RRXCRCR
	_       uint16
	TXCRCR  RTXCRCR
	_       uint16
	I2SCFGR RI2SCFGR
	_       uint16
	I2SPR   RI2SPR
}

func (p *SPI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SPI2 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI2_BASE)))

//emgo:const
var SPI3 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI3_BASE)))

//emgo:const
var SPI1 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE)))

type CR1 uint16

func (b CR1) Field(mask CR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1) J(v int) CR1 {
	return CR1(bits.Make32(v, uint32(mask)))
}

type RCR1 struct{ mmio.U16 }

func (r *RCR1) Bits(mask CR1) CR1     { return CR1(r.U16.Bits(uint16(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U16.SetBits(uint16(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U16.Load()) }
func (r *RCR1) Store(b CR1)           { r.U16.Store(uint16(b)) }

type RMCR1 struct{ mmio.UM16 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM16.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) CPHA() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(CPHA)}}
}

func (p *SPI_Periph) CPOL() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(CPOL)}}
}

func (p *SPI_Periph) MSTR() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(MSTR)}}
}

func (p *SPI_Periph) BR() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(BR)}}
}

func (p *SPI_Periph) SPE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(SPE)}}
}

func (p *SPI_Periph) LSBFIRST() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(LSBFIRST)}}
}

func (p *SPI_Periph) SSI() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(SSI)}}
}

func (p *SPI_Periph) SSM() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(SSM)}}
}

func (p *SPI_Periph) RXONLY() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(RXONLY)}}
}

func (p *SPI_Periph) DFF() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(DFF)}}
}

func (p *SPI_Periph) CRCNEXT() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(CRCNEXT)}}
}

func (p *SPI_Periph) CRCEN() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(CRCEN)}}
}

func (p *SPI_Periph) BIDIOE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(BIDIOE)}}
}

func (p *SPI_Periph) BIDIMODE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(BIDIMODE)}}
}

type CR2 uint16

func (b CR2) Field(mask CR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2) J(v int) CR2 {
	return CR2(bits.Make32(v, uint32(mask)))
}

type RCR2 struct{ mmio.U16 }

func (r *RCR2) Bits(mask CR2) CR2     { return CR2(r.U16.Bits(uint16(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U16.SetBits(uint16(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U16.Load()) }
func (r *RCR2) Store(b CR2)           { r.U16.Store(uint16(b)) }

type RMCR2 struct{ mmio.UM16 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM16.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) RXDMAEN() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(RXDMAEN)}}
}

func (p *SPI_Periph) TXDMAEN() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(TXDMAEN)}}
}

func (p *SPI_Periph) SSOE() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(SSOE)}}
}

func (p *SPI_Periph) FRF() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(FRF)}}
}

func (p *SPI_Periph) ERRIE() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(ERRIE)}}
}

func (p *SPI_Periph) RXNEIE() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(RXNEIE)}}
}

func (p *SPI_Periph) TXEIE() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(TXEIE)}}
}

type SR uint16

func (b SR) Field(mask SR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR) J(v int) SR {
	return SR(bits.Make32(v, uint32(mask)))
}

type RSR struct{ mmio.U16 }

func (r *RSR) Bits(mask SR) SR      { return SR(r.U16.Bits(uint16(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RSR) SetBits(mask SR)      { r.U16.SetBits(uint16(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RSR) Load() SR             { return SR(r.U16.Load()) }
func (r *RSR) Store(b SR)           { r.U16.Store(uint16(b)) }

type RMSR struct{ mmio.UM16 }

func (rm RMSR) Load() SR   { return SR(rm.UM16.Load()) }
func (rm RMSR) Store(b SR) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) RXNE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(RXNE)}}
}

func (p *SPI_Periph) TXE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(TXE)}}
}

func (p *SPI_Periph) CHSIDE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CHSIDE)}}
}

func (p *SPI_Periph) UDR() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(UDR)}}
}

func (p *SPI_Periph) CRCERR() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CRCERR)}}
}

func (p *SPI_Periph) MODF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(MODF)}}
}

func (p *SPI_Periph) OVR() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(OVR)}}
}

func (p *SPI_Periph) BSY() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(BSY)}}
}

type DR uint16

func (b DR) Field(mask DR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR) J(v int) DR {
	return DR(bits.Make32(v, uint32(mask)))
}

type RDR struct{ mmio.U16 }

func (r *RDR) Bits(mask DR) DR      { return DR(r.U16.Bits(uint16(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDR) SetBits(mask DR)      { r.U16.SetBits(uint16(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDR) Load() DR             { return DR(r.U16.Load()) }
func (r *RDR) Store(b DR)           { r.U16.Store(uint16(b)) }

type RMDR struct{ mmio.UM16 }

func (rm RMDR) Load() DR   { return DR(rm.UM16.Load()) }
func (rm RMDR) Store(b DR) { rm.UM16.Store(uint16(b)) }

type CRCPR uint16

func (b CRCPR) Field(mask CRCPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CRCPR) J(v int) CRCPR {
	return CRCPR(bits.Make32(v, uint32(mask)))
}

type RCRCPR struct{ mmio.U16 }

func (r *RCRCPR) Bits(mask CRCPR) CRCPR   { return CRCPR(r.U16.Bits(uint16(mask))) }
func (r *RCRCPR) StoreBits(mask, b CRCPR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCRCPR) SetBits(mask CRCPR)      { r.U16.SetBits(uint16(mask)) }
func (r *RCRCPR) ClearBits(mask CRCPR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCRCPR) Load() CRCPR             { return CRCPR(r.U16.Load()) }
func (r *RCRCPR) Store(b CRCPR)           { r.U16.Store(uint16(b)) }

type RMCRCPR struct{ mmio.UM16 }

func (rm RMCRCPR) Load() CRCPR   { return CRCPR(rm.UM16.Load()) }
func (rm RMCRCPR) Store(b CRCPR) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) CRCPOLY() RMCRCPR {
	return RMCRCPR{mmio.UM16{&p.CRCPR.U16, uint16(CRCPOLY)}}
}

type RXCRCR uint16

func (b RXCRCR) Field(mask RXCRCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RXCRCR) J(v int) RXCRCR {
	return RXCRCR(bits.Make32(v, uint32(mask)))
}

type RRXCRCR struct{ mmio.U16 }

func (r *RRXCRCR) Bits(mask RXCRCR) RXCRCR  { return RXCRCR(r.U16.Bits(uint16(mask))) }
func (r *RRXCRCR) StoreBits(mask, b RXCRCR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RRXCRCR) SetBits(mask RXCRCR)      { r.U16.SetBits(uint16(mask)) }
func (r *RRXCRCR) ClearBits(mask RXCRCR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RRXCRCR) Load() RXCRCR             { return RXCRCR(r.U16.Load()) }
func (r *RRXCRCR) Store(b RXCRCR)           { r.U16.Store(uint16(b)) }

type RMRXCRCR struct{ mmio.UM16 }

func (rm RMRXCRCR) Load() RXCRCR   { return RXCRCR(rm.UM16.Load()) }
func (rm RMRXCRCR) Store(b RXCRCR) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) RXCRC() RMRXCRCR {
	return RMRXCRCR{mmio.UM16{&p.RXCRCR.U16, uint16(RXCRC)}}
}

type TXCRCR uint16

func (b TXCRCR) Field(mask TXCRCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TXCRCR) J(v int) TXCRCR {
	return TXCRCR(bits.Make32(v, uint32(mask)))
}

type RTXCRCR struct{ mmio.U16 }

func (r *RTXCRCR) Bits(mask TXCRCR) TXCRCR  { return TXCRCR(r.U16.Bits(uint16(mask))) }
func (r *RTXCRCR) StoreBits(mask, b TXCRCR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RTXCRCR) SetBits(mask TXCRCR)      { r.U16.SetBits(uint16(mask)) }
func (r *RTXCRCR) ClearBits(mask TXCRCR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RTXCRCR) Load() TXCRCR             { return TXCRCR(r.U16.Load()) }
func (r *RTXCRCR) Store(b TXCRCR)           { r.U16.Store(uint16(b)) }

type RMTXCRCR struct{ mmio.UM16 }

func (rm RMTXCRCR) Load() TXCRCR   { return TXCRCR(rm.UM16.Load()) }
func (rm RMTXCRCR) Store(b TXCRCR) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) TXCRC() RMTXCRCR {
	return RMTXCRCR{mmio.UM16{&p.TXCRCR.U16, uint16(TXCRC)}}
}

type I2SCFGR uint16

func (b I2SCFGR) Field(mask I2SCFGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask I2SCFGR) J(v int) I2SCFGR {
	return I2SCFGR(bits.Make32(v, uint32(mask)))
}

type RI2SCFGR struct{ mmio.U16 }

func (r *RI2SCFGR) Bits(mask I2SCFGR) I2SCFGR { return I2SCFGR(r.U16.Bits(uint16(mask))) }
func (r *RI2SCFGR) StoreBits(mask, b I2SCFGR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RI2SCFGR) SetBits(mask I2SCFGR)      { r.U16.SetBits(uint16(mask)) }
func (r *RI2SCFGR) ClearBits(mask I2SCFGR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RI2SCFGR) Load() I2SCFGR             { return I2SCFGR(r.U16.Load()) }
func (r *RI2SCFGR) Store(b I2SCFGR)           { r.U16.Store(uint16(b)) }

type RMI2SCFGR struct{ mmio.UM16 }

func (rm RMI2SCFGR) Load() I2SCFGR   { return I2SCFGR(rm.UM16.Load()) }
func (rm RMI2SCFGR) Store(b I2SCFGR) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) CHLEN() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM16{&p.I2SCFGR.U16, uint16(CHLEN)}}
}

func (p *SPI_Periph) DATLEN() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM16{&p.I2SCFGR.U16, uint16(DATLEN)}}
}

func (p *SPI_Periph) CKPOL() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM16{&p.I2SCFGR.U16, uint16(CKPOL)}}
}

func (p *SPI_Periph) I2SSTD() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM16{&p.I2SCFGR.U16, uint16(I2SSTD)}}
}

func (p *SPI_Periph) PCMSYNC() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM16{&p.I2SCFGR.U16, uint16(PCMSYNC)}}
}

func (p *SPI_Periph) I2SCFG() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM16{&p.I2SCFGR.U16, uint16(I2SCFG)}}
}

func (p *SPI_Periph) I2SE() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM16{&p.I2SCFGR.U16, uint16(I2SE)}}
}

func (p *SPI_Periph) I2SMOD() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM16{&p.I2SCFGR.U16, uint16(I2SMOD)}}
}

type I2SPR uint16

func (b I2SPR) Field(mask I2SPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask I2SPR) J(v int) I2SPR {
	return I2SPR(bits.Make32(v, uint32(mask)))
}

type RI2SPR struct{ mmio.U16 }

func (r *RI2SPR) Bits(mask I2SPR) I2SPR   { return I2SPR(r.U16.Bits(uint16(mask))) }
func (r *RI2SPR) StoreBits(mask, b I2SPR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RI2SPR) SetBits(mask I2SPR)      { r.U16.SetBits(uint16(mask)) }
func (r *RI2SPR) ClearBits(mask I2SPR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RI2SPR) Load() I2SPR             { return I2SPR(r.U16.Load()) }
func (r *RI2SPR) Store(b I2SPR)           { r.U16.Store(uint16(b)) }

type RMI2SPR struct{ mmio.UM16 }

func (rm RMI2SPR) Load() I2SPR   { return I2SPR(rm.UM16.Load()) }
func (rm RMI2SPR) Store(b I2SPR) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) I2SDIV() RMI2SPR {
	return RMI2SPR{mmio.UM16{&p.I2SPR.U16, uint16(I2SDIV)}}
}

func (p *SPI_Periph) ODD() RMI2SPR {
	return RMI2SPR{mmio.UM16{&p.I2SPR.U16, uint16(ODD)}}
}

func (p *SPI_Periph) MCKOE() RMI2SPR {
	return RMI2SPR{mmio.UM16{&p.I2SPR.U16, uint16(MCKOE)}}
}
