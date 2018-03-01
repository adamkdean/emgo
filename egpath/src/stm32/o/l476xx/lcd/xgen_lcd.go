package lcd

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type LCD_Periph struct {
	CR  RCR
	FCR RFCR
	SR  RSR
	CLR RCLR
	_   uint32
	RAM [16]RRAM
}

func (p *LCD_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var LCD = (*LCD_Periph)(unsafe.Pointer(uintptr(mmap.LCD_BASE)))

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

func (p *LCD_Periph) LCDEN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LCDEN)}}
}

func (p *LCD_Periph) VSEL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(VSEL)}}
}

func (p *LCD_Periph) DUTY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DUTY)}}
}

func (p *LCD_Periph) BIAS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BIAS)}}
}

func (p *LCD_Periph) MUX_SEG() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MUX_SEG)}}
}

func (p *LCD_Periph) BUFEN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BUFEN)}}
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

func (p *LCD_Periph) HD() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(HD)}}
}

func (p *LCD_Periph) SOFIE() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(SOFIE)}}
}

func (p *LCD_Periph) UDDIE() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(UDDIE)}}
}

func (p *LCD_Periph) PON() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(PON)}}
}

func (p *LCD_Periph) DEAD() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(DEAD)}}
}

func (p *LCD_Periph) CC() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(CC)}}
}

func (p *LCD_Periph) BLINKF() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(BLINKF)}}
}

func (p *LCD_Periph) BLINK() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(BLINK)}}
}

func (p *LCD_Periph) DIV() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(DIV)}}
}

func (p *LCD_Periph) PS() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(PS)}}
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

func (p *LCD_Periph) ENS() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(ENS)}}
}

func (p *LCD_Periph) SOF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(SOF)}}
}

func (p *LCD_Periph) UDR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(UDR)}}
}

func (p *LCD_Periph) UDD() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(UDD)}}
}

func (p *LCD_Periph) RDY() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(RDY)}}
}

func (p *LCD_Periph) FCRSR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FCRSR)}}
}

type CLR uint32

func (b CLR) Field(mask CLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLR) J(v int) CLR {
	return CLR(bits.Make32(v, uint32(mask)))
}

type RCLR struct{ mmio.U32 }

func (r *RCLR) Bits(mask CLR) CLR     { return CLR(r.U32.Bits(uint32(mask))) }
func (r *RCLR) StoreBits(mask, b CLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCLR) SetBits(mask CLR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCLR) ClearBits(mask CLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCLR) Load() CLR             { return CLR(r.U32.Load()) }
func (r *RCLR) Store(b CLR)           { r.U32.Store(uint32(b)) }

func (r *RCLR) AtomicStoreBits(mask, b CLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCLR) AtomicSetBits(mask CLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCLR) AtomicClearBits(mask CLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCLR struct{ mmio.UM32 }

func (rm RMCLR) Load() CLR   { return CLR(rm.UM32.Load()) }
func (rm RMCLR) Store(b CLR) { rm.UM32.Store(uint32(b)) }

func (p *LCD_Periph) SOFC() RMCLR {
	return RMCLR{mmio.UM32{&p.CLR.U32, uint32(SOFC)}}
}

func (p *LCD_Periph) UDDC() RMCLR {
	return RMCLR{mmio.UM32{&p.CLR.U32, uint32(UDDC)}}
}

type RAM uint32

func (b RAM) Field(mask RAM) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RAM) J(v int) RAM {
	return RAM(bits.Make32(v, uint32(mask)))
}

type RRAM struct{ mmio.U32 }

func (r *RRAM) Bits(mask RAM) RAM     { return RAM(r.U32.Bits(uint32(mask))) }
func (r *RRAM) StoreBits(mask, b RAM) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRAM) SetBits(mask RAM)      { r.U32.SetBits(uint32(mask)) }
func (r *RRAM) ClearBits(mask RAM)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRAM) Load() RAM             { return RAM(r.U32.Load()) }
func (r *RRAM) Store(b RAM)           { r.U32.Store(uint32(b)) }

func (r *RRAM) AtomicStoreBits(mask, b RAM) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRAM) AtomicSetBits(mask RAM)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRAM) AtomicClearBits(mask RAM)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRAM struct{ mmio.UM32 }

func (rm RMRAM) Load() RAM   { return RAM(rm.UM32.Load()) }
func (rm RMRAM) Store(b RAM) { rm.UM32.Store(uint32(b)) }

func (p *LCD_Periph) SEGMENT_DATA(n int) RMRAM {
	return RMRAM{mmio.UM32{&p.RAM[n].U32, uint32(SEGMENT_DATA)}}
}
