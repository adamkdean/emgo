// +build f10x_hd

package bkp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type BKP_Periph struct {
	_     uint32
	DR1   DR1
	_     uint16
	DR2   DR2
	_     uint16
	DR3   DR3
	_     uint16
	DR4   DR4
	_     uint16
	DR5   DR5
	_     uint16
	DR6   DR6
	_     uint16
	DR7   DR7
	_     uint16
	DR8   DR8
	_     uint16
	DR9   DR9
	_     uint16
	DR10  DR10
	_     uint16
	RTCCR RTCCR
	_     uint16
	CR    CR
	_     uint16
	CSR   CSR
	_     uint16
	_     [2]uint32
	DR11  DR11
	_     uint16
	DR12  DR12
	_     uint16
	DR13  DR13
	_     uint16
	DR14  DR14
	_     uint16
	DR15  DR15
	_     uint16
	DR16  DR16
	_     uint16
	DR17  DR17
	_     uint16
	DR18  DR18
	_     uint16
	DR19  DR19
	_     uint16
	DR20  DR20
	_     uint16
	DR21  DR21
	_     uint16
	DR22  DR22
	_     uint16
	DR23  DR23
	_     uint16
	DR24  DR24
	_     uint16
	DR25  DR25
	_     uint16
	DR26  DR26
	_     uint16
	DR27  DR27
	_     uint16
	DR28  DR28
	_     uint16
	DR29  DR29
	_     uint16
	DR30  DR30
	_     uint16
	DR31  DR31
	_     uint16
	DR32  DR32
	_     uint16
	DR33  DR33
	_     uint16
	DR34  DR34
	_     uint16
	DR35  DR35
	_     uint16
	DR36  DR36
	_     uint16
	DR37  DR37
	_     uint16
	DR38  DR38
	_     uint16
	DR39  DR39
	_     uint16
	DR40  DR40
	_     uint16
	DR41  DR41
	_     uint16
	DR42  DR42
}

func (p *BKP_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var BKP = (*BKP_Periph)(unsafe.Pointer(uintptr(mmap.BKP_BASE)))

type DR1_Bits uint16

type DR1 struct{ mmio.U16 }

func (r *DR1) Bits(mask DR1_Bits) DR1_Bits { return DR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR1) StoreBits(mask, b DR1_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR1) SetBits(mask DR1_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR1) ClearBits(mask DR1_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR1) Load() DR1_Bits              { return DR1_Bits(r.U16.Load()) }
func (r *DR1) Store(b DR1_Bits)            { r.U16.Store(uint16(b)) }

type DR1_Mask struct{ mmio.UM16 }

func (rm DR1_Mask) Load() DR1_Bits   { return DR1_Bits(rm.UM16.Load()) }
func (rm DR1_Mask) Store(b DR1_Bits) { rm.UM16.Store(uint16(b)) }

type DR2_Bits uint16

type DR2 struct{ mmio.U16 }

func (r *DR2) Bits(mask DR2_Bits) DR2_Bits { return DR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR2) StoreBits(mask, b DR2_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR2) SetBits(mask DR2_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR2) ClearBits(mask DR2_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR2) Load() DR2_Bits              { return DR2_Bits(r.U16.Load()) }
func (r *DR2) Store(b DR2_Bits)            { r.U16.Store(uint16(b)) }

type DR2_Mask struct{ mmio.UM16 }

func (rm DR2_Mask) Load() DR2_Bits   { return DR2_Bits(rm.UM16.Load()) }
func (rm DR2_Mask) Store(b DR2_Bits) { rm.UM16.Store(uint16(b)) }

type DR3_Bits uint16

type DR3 struct{ mmio.U16 }

func (r *DR3) Bits(mask DR3_Bits) DR3_Bits { return DR3_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR3) StoreBits(mask, b DR3_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR3) SetBits(mask DR3_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR3) ClearBits(mask DR3_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR3) Load() DR3_Bits              { return DR3_Bits(r.U16.Load()) }
func (r *DR3) Store(b DR3_Bits)            { r.U16.Store(uint16(b)) }

type DR3_Mask struct{ mmio.UM16 }

func (rm DR3_Mask) Load() DR3_Bits   { return DR3_Bits(rm.UM16.Load()) }
func (rm DR3_Mask) Store(b DR3_Bits) { rm.UM16.Store(uint16(b)) }

type DR4_Bits uint16

type DR4 struct{ mmio.U16 }

func (r *DR4) Bits(mask DR4_Bits) DR4_Bits { return DR4_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR4) StoreBits(mask, b DR4_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR4) SetBits(mask DR4_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR4) ClearBits(mask DR4_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR4) Load() DR4_Bits              { return DR4_Bits(r.U16.Load()) }
func (r *DR4) Store(b DR4_Bits)            { r.U16.Store(uint16(b)) }

type DR4_Mask struct{ mmio.UM16 }

func (rm DR4_Mask) Load() DR4_Bits   { return DR4_Bits(rm.UM16.Load()) }
func (rm DR4_Mask) Store(b DR4_Bits) { rm.UM16.Store(uint16(b)) }

type DR5_Bits uint16

type DR5 struct{ mmio.U16 }

func (r *DR5) Bits(mask DR5_Bits) DR5_Bits { return DR5_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR5) StoreBits(mask, b DR5_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR5) SetBits(mask DR5_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR5) ClearBits(mask DR5_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR5) Load() DR5_Bits              { return DR5_Bits(r.U16.Load()) }
func (r *DR5) Store(b DR5_Bits)            { r.U16.Store(uint16(b)) }

type DR5_Mask struct{ mmio.UM16 }

func (rm DR5_Mask) Load() DR5_Bits   { return DR5_Bits(rm.UM16.Load()) }
func (rm DR5_Mask) Store(b DR5_Bits) { rm.UM16.Store(uint16(b)) }

type DR6_Bits uint16

type DR6 struct{ mmio.U16 }

func (r *DR6) Bits(mask DR6_Bits) DR6_Bits { return DR6_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR6) StoreBits(mask, b DR6_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR6) SetBits(mask DR6_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR6) ClearBits(mask DR6_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR6) Load() DR6_Bits              { return DR6_Bits(r.U16.Load()) }
func (r *DR6) Store(b DR6_Bits)            { r.U16.Store(uint16(b)) }

type DR6_Mask struct{ mmio.UM16 }

func (rm DR6_Mask) Load() DR6_Bits   { return DR6_Bits(rm.UM16.Load()) }
func (rm DR6_Mask) Store(b DR6_Bits) { rm.UM16.Store(uint16(b)) }

type DR7_Bits uint16

type DR7 struct{ mmio.U16 }

func (r *DR7) Bits(mask DR7_Bits) DR7_Bits { return DR7_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR7) StoreBits(mask, b DR7_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR7) SetBits(mask DR7_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR7) ClearBits(mask DR7_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR7) Load() DR7_Bits              { return DR7_Bits(r.U16.Load()) }
func (r *DR7) Store(b DR7_Bits)            { r.U16.Store(uint16(b)) }

type DR7_Mask struct{ mmio.UM16 }

func (rm DR7_Mask) Load() DR7_Bits   { return DR7_Bits(rm.UM16.Load()) }
func (rm DR7_Mask) Store(b DR7_Bits) { rm.UM16.Store(uint16(b)) }

type DR8_Bits uint16

type DR8 struct{ mmio.U16 }

func (r *DR8) Bits(mask DR8_Bits) DR8_Bits { return DR8_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR8) StoreBits(mask, b DR8_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR8) SetBits(mask DR8_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR8) ClearBits(mask DR8_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR8) Load() DR8_Bits              { return DR8_Bits(r.U16.Load()) }
func (r *DR8) Store(b DR8_Bits)            { r.U16.Store(uint16(b)) }

type DR8_Mask struct{ mmio.UM16 }

func (rm DR8_Mask) Load() DR8_Bits   { return DR8_Bits(rm.UM16.Load()) }
func (rm DR8_Mask) Store(b DR8_Bits) { rm.UM16.Store(uint16(b)) }

type DR9_Bits uint16

type DR9 struct{ mmio.U16 }

func (r *DR9) Bits(mask DR9_Bits) DR9_Bits { return DR9_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR9) StoreBits(mask, b DR9_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR9) SetBits(mask DR9_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DR9) ClearBits(mask DR9_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DR9) Load() DR9_Bits              { return DR9_Bits(r.U16.Load()) }
func (r *DR9) Store(b DR9_Bits)            { r.U16.Store(uint16(b)) }

type DR9_Mask struct{ mmio.UM16 }

func (rm DR9_Mask) Load() DR9_Bits   { return DR9_Bits(rm.UM16.Load()) }
func (rm DR9_Mask) Store(b DR9_Bits) { rm.UM16.Store(uint16(b)) }

type DR10_Bits uint16

type DR10 struct{ mmio.U16 }

func (r *DR10) Bits(mask DR10_Bits) DR10_Bits { return DR10_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR10) StoreBits(mask, b DR10_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR10) SetBits(mask DR10_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR10) ClearBits(mask DR10_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR10) Load() DR10_Bits               { return DR10_Bits(r.U16.Load()) }
func (r *DR10) Store(b DR10_Bits)             { r.U16.Store(uint16(b)) }

type DR10_Mask struct{ mmio.UM16 }

func (rm DR10_Mask) Load() DR10_Bits   { return DR10_Bits(rm.UM16.Load()) }
func (rm DR10_Mask) Store(b DR10_Bits) { rm.UM16.Store(uint16(b)) }

type RTCCR_Bits uint16

type RTCCR struct{ mmio.U16 }

func (r *RTCCR) Bits(mask RTCCR_Bits) RTCCR_Bits { return RTCCR_Bits(r.U16.Bits(uint16(mask))) }
func (r *RTCCR) StoreBits(mask, b RTCCR_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RTCCR) SetBits(mask RTCCR_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *RTCCR) ClearBits(mask RTCCR_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *RTCCR) Load() RTCCR_Bits                { return RTCCR_Bits(r.U16.Load()) }
func (r *RTCCR) Store(b RTCCR_Bits)              { r.U16.Store(uint16(b)) }

type RTCCR_Mask struct{ mmio.UM16 }

func (rm RTCCR_Mask) Load() RTCCR_Bits   { return RTCCR_Bits(rm.UM16.Load()) }
func (rm RTCCR_Mask) Store(b RTCCR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *BKP_Periph) CAL() RTCCR_Mask {
	return RTCCR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint16(CAL)}}
}

func (p *BKP_Periph) CCO() RTCCR_Mask {
	return RTCCR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint16(CCO)}}
}

func (p *BKP_Periph) ASOE() RTCCR_Mask {
	return RTCCR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint16(ASOE)}}
}

func (p *BKP_Periph) ASOS() RTCCR_Mask {
	return RTCCR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint16(ASOS)}}
}

type CR_Bits uint16

type CR struct{ mmio.U16 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U16.SetBits(uint16(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U16.ClearBits(uint16(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U16.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U16.Store(uint16(b)) }

type CR_Mask struct{ mmio.UM16 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM16.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *BKP_Periph) TPE() CR_Mask {
	return CR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint16(TPE)}}
}

func (p *BKP_Periph) TPAL() CR_Mask {
	return CR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint16(TPAL)}}
}

type CSR_Bits uint16

type CSR struct{ mmio.U16 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U16.Bits(uint16(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U16.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U16.Store(uint16(b)) }

type CSR_Mask struct{ mmio.UM16 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM16.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *BKP_Periph) CTE() CSR_Mask {
	return CSR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint16(CTE)}}
}

func (p *BKP_Periph) CTI() CSR_Mask {
	return CSR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint16(CTI)}}
}

func (p *BKP_Periph) TPIE() CSR_Mask {
	return CSR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint16(TPIE)}}
}

func (p *BKP_Periph) TEF() CSR_Mask {
	return CSR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint16(TEF)}}
}

func (p *BKP_Periph) TIF() CSR_Mask {
	return CSR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint16(TIF)}}
}

type DR11_Bits uint16

type DR11 struct{ mmio.U16 }

func (r *DR11) Bits(mask DR11_Bits) DR11_Bits { return DR11_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR11) StoreBits(mask, b DR11_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR11) SetBits(mask DR11_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR11) ClearBits(mask DR11_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR11) Load() DR11_Bits               { return DR11_Bits(r.U16.Load()) }
func (r *DR11) Store(b DR11_Bits)             { r.U16.Store(uint16(b)) }

type DR11_Mask struct{ mmio.UM16 }

func (rm DR11_Mask) Load() DR11_Bits   { return DR11_Bits(rm.UM16.Load()) }
func (rm DR11_Mask) Store(b DR11_Bits) { rm.UM16.Store(uint16(b)) }

type DR12_Bits uint16

type DR12 struct{ mmio.U16 }

func (r *DR12) Bits(mask DR12_Bits) DR12_Bits { return DR12_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR12) StoreBits(mask, b DR12_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR12) SetBits(mask DR12_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR12) ClearBits(mask DR12_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR12) Load() DR12_Bits               { return DR12_Bits(r.U16.Load()) }
func (r *DR12) Store(b DR12_Bits)             { r.U16.Store(uint16(b)) }

type DR12_Mask struct{ mmio.UM16 }

func (rm DR12_Mask) Load() DR12_Bits   { return DR12_Bits(rm.UM16.Load()) }
func (rm DR12_Mask) Store(b DR12_Bits) { rm.UM16.Store(uint16(b)) }

type DR13_Bits uint16

type DR13 struct{ mmio.U16 }

func (r *DR13) Bits(mask DR13_Bits) DR13_Bits { return DR13_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR13) StoreBits(mask, b DR13_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR13) SetBits(mask DR13_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR13) ClearBits(mask DR13_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR13) Load() DR13_Bits               { return DR13_Bits(r.U16.Load()) }
func (r *DR13) Store(b DR13_Bits)             { r.U16.Store(uint16(b)) }

type DR13_Mask struct{ mmio.UM16 }

func (rm DR13_Mask) Load() DR13_Bits   { return DR13_Bits(rm.UM16.Load()) }
func (rm DR13_Mask) Store(b DR13_Bits) { rm.UM16.Store(uint16(b)) }

type DR14_Bits uint16

type DR14 struct{ mmio.U16 }

func (r *DR14) Bits(mask DR14_Bits) DR14_Bits { return DR14_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR14) StoreBits(mask, b DR14_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR14) SetBits(mask DR14_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR14) ClearBits(mask DR14_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR14) Load() DR14_Bits               { return DR14_Bits(r.U16.Load()) }
func (r *DR14) Store(b DR14_Bits)             { r.U16.Store(uint16(b)) }

type DR14_Mask struct{ mmio.UM16 }

func (rm DR14_Mask) Load() DR14_Bits   { return DR14_Bits(rm.UM16.Load()) }
func (rm DR14_Mask) Store(b DR14_Bits) { rm.UM16.Store(uint16(b)) }

type DR15_Bits uint16

type DR15 struct{ mmio.U16 }

func (r *DR15) Bits(mask DR15_Bits) DR15_Bits { return DR15_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR15) StoreBits(mask, b DR15_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR15) SetBits(mask DR15_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR15) ClearBits(mask DR15_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR15) Load() DR15_Bits               { return DR15_Bits(r.U16.Load()) }
func (r *DR15) Store(b DR15_Bits)             { r.U16.Store(uint16(b)) }

type DR15_Mask struct{ mmio.UM16 }

func (rm DR15_Mask) Load() DR15_Bits   { return DR15_Bits(rm.UM16.Load()) }
func (rm DR15_Mask) Store(b DR15_Bits) { rm.UM16.Store(uint16(b)) }

type DR16_Bits uint16

type DR16 struct{ mmio.U16 }

func (r *DR16) Bits(mask DR16_Bits) DR16_Bits { return DR16_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR16) StoreBits(mask, b DR16_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR16) SetBits(mask DR16_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR16) ClearBits(mask DR16_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR16) Load() DR16_Bits               { return DR16_Bits(r.U16.Load()) }
func (r *DR16) Store(b DR16_Bits)             { r.U16.Store(uint16(b)) }

type DR16_Mask struct{ mmio.UM16 }

func (rm DR16_Mask) Load() DR16_Bits   { return DR16_Bits(rm.UM16.Load()) }
func (rm DR16_Mask) Store(b DR16_Bits) { rm.UM16.Store(uint16(b)) }

type DR17_Bits uint16

type DR17 struct{ mmio.U16 }

func (r *DR17) Bits(mask DR17_Bits) DR17_Bits { return DR17_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR17) StoreBits(mask, b DR17_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR17) SetBits(mask DR17_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR17) ClearBits(mask DR17_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR17) Load() DR17_Bits               { return DR17_Bits(r.U16.Load()) }
func (r *DR17) Store(b DR17_Bits)             { r.U16.Store(uint16(b)) }

type DR17_Mask struct{ mmio.UM16 }

func (rm DR17_Mask) Load() DR17_Bits   { return DR17_Bits(rm.UM16.Load()) }
func (rm DR17_Mask) Store(b DR17_Bits) { rm.UM16.Store(uint16(b)) }

type DR18_Bits uint16

type DR18 struct{ mmio.U16 }

func (r *DR18) Bits(mask DR18_Bits) DR18_Bits { return DR18_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR18) StoreBits(mask, b DR18_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR18) SetBits(mask DR18_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR18) ClearBits(mask DR18_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR18) Load() DR18_Bits               { return DR18_Bits(r.U16.Load()) }
func (r *DR18) Store(b DR18_Bits)             { r.U16.Store(uint16(b)) }

type DR18_Mask struct{ mmio.UM16 }

func (rm DR18_Mask) Load() DR18_Bits   { return DR18_Bits(rm.UM16.Load()) }
func (rm DR18_Mask) Store(b DR18_Bits) { rm.UM16.Store(uint16(b)) }

type DR19_Bits uint16

type DR19 struct{ mmio.U16 }

func (r *DR19) Bits(mask DR19_Bits) DR19_Bits { return DR19_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR19) StoreBits(mask, b DR19_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR19) SetBits(mask DR19_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR19) ClearBits(mask DR19_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR19) Load() DR19_Bits               { return DR19_Bits(r.U16.Load()) }
func (r *DR19) Store(b DR19_Bits)             { r.U16.Store(uint16(b)) }

type DR19_Mask struct{ mmio.UM16 }

func (rm DR19_Mask) Load() DR19_Bits   { return DR19_Bits(rm.UM16.Load()) }
func (rm DR19_Mask) Store(b DR19_Bits) { rm.UM16.Store(uint16(b)) }

type DR20_Bits uint16

type DR20 struct{ mmio.U16 }

func (r *DR20) Bits(mask DR20_Bits) DR20_Bits { return DR20_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR20) StoreBits(mask, b DR20_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR20) SetBits(mask DR20_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR20) ClearBits(mask DR20_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR20) Load() DR20_Bits               { return DR20_Bits(r.U16.Load()) }
func (r *DR20) Store(b DR20_Bits)             { r.U16.Store(uint16(b)) }

type DR20_Mask struct{ mmio.UM16 }

func (rm DR20_Mask) Load() DR20_Bits   { return DR20_Bits(rm.UM16.Load()) }
func (rm DR20_Mask) Store(b DR20_Bits) { rm.UM16.Store(uint16(b)) }

type DR21_Bits uint16

type DR21 struct{ mmio.U16 }

func (r *DR21) Bits(mask DR21_Bits) DR21_Bits { return DR21_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR21) StoreBits(mask, b DR21_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR21) SetBits(mask DR21_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR21) ClearBits(mask DR21_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR21) Load() DR21_Bits               { return DR21_Bits(r.U16.Load()) }
func (r *DR21) Store(b DR21_Bits)             { r.U16.Store(uint16(b)) }

type DR21_Mask struct{ mmio.UM16 }

func (rm DR21_Mask) Load() DR21_Bits   { return DR21_Bits(rm.UM16.Load()) }
func (rm DR21_Mask) Store(b DR21_Bits) { rm.UM16.Store(uint16(b)) }

type DR22_Bits uint16

type DR22 struct{ mmio.U16 }

func (r *DR22) Bits(mask DR22_Bits) DR22_Bits { return DR22_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR22) StoreBits(mask, b DR22_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR22) SetBits(mask DR22_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR22) ClearBits(mask DR22_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR22) Load() DR22_Bits               { return DR22_Bits(r.U16.Load()) }
func (r *DR22) Store(b DR22_Bits)             { r.U16.Store(uint16(b)) }

type DR22_Mask struct{ mmio.UM16 }

func (rm DR22_Mask) Load() DR22_Bits   { return DR22_Bits(rm.UM16.Load()) }
func (rm DR22_Mask) Store(b DR22_Bits) { rm.UM16.Store(uint16(b)) }

type DR23_Bits uint16

type DR23 struct{ mmio.U16 }

func (r *DR23) Bits(mask DR23_Bits) DR23_Bits { return DR23_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR23) StoreBits(mask, b DR23_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR23) SetBits(mask DR23_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR23) ClearBits(mask DR23_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR23) Load() DR23_Bits               { return DR23_Bits(r.U16.Load()) }
func (r *DR23) Store(b DR23_Bits)             { r.U16.Store(uint16(b)) }

type DR23_Mask struct{ mmio.UM16 }

func (rm DR23_Mask) Load() DR23_Bits   { return DR23_Bits(rm.UM16.Load()) }
func (rm DR23_Mask) Store(b DR23_Bits) { rm.UM16.Store(uint16(b)) }

type DR24_Bits uint16

type DR24 struct{ mmio.U16 }

func (r *DR24) Bits(mask DR24_Bits) DR24_Bits { return DR24_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR24) StoreBits(mask, b DR24_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR24) SetBits(mask DR24_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR24) ClearBits(mask DR24_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR24) Load() DR24_Bits               { return DR24_Bits(r.U16.Load()) }
func (r *DR24) Store(b DR24_Bits)             { r.U16.Store(uint16(b)) }

type DR24_Mask struct{ mmio.UM16 }

func (rm DR24_Mask) Load() DR24_Bits   { return DR24_Bits(rm.UM16.Load()) }
func (rm DR24_Mask) Store(b DR24_Bits) { rm.UM16.Store(uint16(b)) }

type DR25_Bits uint16

type DR25 struct{ mmio.U16 }

func (r *DR25) Bits(mask DR25_Bits) DR25_Bits { return DR25_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR25) StoreBits(mask, b DR25_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR25) SetBits(mask DR25_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR25) ClearBits(mask DR25_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR25) Load() DR25_Bits               { return DR25_Bits(r.U16.Load()) }
func (r *DR25) Store(b DR25_Bits)             { r.U16.Store(uint16(b)) }

type DR25_Mask struct{ mmio.UM16 }

func (rm DR25_Mask) Load() DR25_Bits   { return DR25_Bits(rm.UM16.Load()) }
func (rm DR25_Mask) Store(b DR25_Bits) { rm.UM16.Store(uint16(b)) }

type DR26_Bits uint16

type DR26 struct{ mmio.U16 }

func (r *DR26) Bits(mask DR26_Bits) DR26_Bits { return DR26_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR26) StoreBits(mask, b DR26_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR26) SetBits(mask DR26_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR26) ClearBits(mask DR26_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR26) Load() DR26_Bits               { return DR26_Bits(r.U16.Load()) }
func (r *DR26) Store(b DR26_Bits)             { r.U16.Store(uint16(b)) }

type DR26_Mask struct{ mmio.UM16 }

func (rm DR26_Mask) Load() DR26_Bits   { return DR26_Bits(rm.UM16.Load()) }
func (rm DR26_Mask) Store(b DR26_Bits) { rm.UM16.Store(uint16(b)) }

type DR27_Bits uint16

type DR27 struct{ mmio.U16 }

func (r *DR27) Bits(mask DR27_Bits) DR27_Bits { return DR27_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR27) StoreBits(mask, b DR27_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR27) SetBits(mask DR27_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR27) ClearBits(mask DR27_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR27) Load() DR27_Bits               { return DR27_Bits(r.U16.Load()) }
func (r *DR27) Store(b DR27_Bits)             { r.U16.Store(uint16(b)) }

type DR27_Mask struct{ mmio.UM16 }

func (rm DR27_Mask) Load() DR27_Bits   { return DR27_Bits(rm.UM16.Load()) }
func (rm DR27_Mask) Store(b DR27_Bits) { rm.UM16.Store(uint16(b)) }

type DR28_Bits uint16

type DR28 struct{ mmio.U16 }

func (r *DR28) Bits(mask DR28_Bits) DR28_Bits { return DR28_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR28) StoreBits(mask, b DR28_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR28) SetBits(mask DR28_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR28) ClearBits(mask DR28_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR28) Load() DR28_Bits               { return DR28_Bits(r.U16.Load()) }
func (r *DR28) Store(b DR28_Bits)             { r.U16.Store(uint16(b)) }

type DR28_Mask struct{ mmio.UM16 }

func (rm DR28_Mask) Load() DR28_Bits   { return DR28_Bits(rm.UM16.Load()) }
func (rm DR28_Mask) Store(b DR28_Bits) { rm.UM16.Store(uint16(b)) }

type DR29_Bits uint16

type DR29 struct{ mmio.U16 }

func (r *DR29) Bits(mask DR29_Bits) DR29_Bits { return DR29_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR29) StoreBits(mask, b DR29_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR29) SetBits(mask DR29_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR29) ClearBits(mask DR29_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR29) Load() DR29_Bits               { return DR29_Bits(r.U16.Load()) }
func (r *DR29) Store(b DR29_Bits)             { r.U16.Store(uint16(b)) }

type DR29_Mask struct{ mmio.UM16 }

func (rm DR29_Mask) Load() DR29_Bits   { return DR29_Bits(rm.UM16.Load()) }
func (rm DR29_Mask) Store(b DR29_Bits) { rm.UM16.Store(uint16(b)) }

type DR30_Bits uint16

type DR30 struct{ mmio.U16 }

func (r *DR30) Bits(mask DR30_Bits) DR30_Bits { return DR30_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR30) StoreBits(mask, b DR30_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR30) SetBits(mask DR30_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR30) ClearBits(mask DR30_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR30) Load() DR30_Bits               { return DR30_Bits(r.U16.Load()) }
func (r *DR30) Store(b DR30_Bits)             { r.U16.Store(uint16(b)) }

type DR30_Mask struct{ mmio.UM16 }

func (rm DR30_Mask) Load() DR30_Bits   { return DR30_Bits(rm.UM16.Load()) }
func (rm DR30_Mask) Store(b DR30_Bits) { rm.UM16.Store(uint16(b)) }

type DR31_Bits uint16

type DR31 struct{ mmio.U16 }

func (r *DR31) Bits(mask DR31_Bits) DR31_Bits { return DR31_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR31) StoreBits(mask, b DR31_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR31) SetBits(mask DR31_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR31) ClearBits(mask DR31_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR31) Load() DR31_Bits               { return DR31_Bits(r.U16.Load()) }
func (r *DR31) Store(b DR31_Bits)             { r.U16.Store(uint16(b)) }

type DR31_Mask struct{ mmio.UM16 }

func (rm DR31_Mask) Load() DR31_Bits   { return DR31_Bits(rm.UM16.Load()) }
func (rm DR31_Mask) Store(b DR31_Bits) { rm.UM16.Store(uint16(b)) }

type DR32_Bits uint16

type DR32 struct{ mmio.U16 }

func (r *DR32) Bits(mask DR32_Bits) DR32_Bits { return DR32_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR32) StoreBits(mask, b DR32_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR32) SetBits(mask DR32_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR32) ClearBits(mask DR32_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR32) Load() DR32_Bits               { return DR32_Bits(r.U16.Load()) }
func (r *DR32) Store(b DR32_Bits)             { r.U16.Store(uint16(b)) }

type DR32_Mask struct{ mmio.UM16 }

func (rm DR32_Mask) Load() DR32_Bits   { return DR32_Bits(rm.UM16.Load()) }
func (rm DR32_Mask) Store(b DR32_Bits) { rm.UM16.Store(uint16(b)) }

type DR33_Bits uint16

type DR33 struct{ mmio.U16 }

func (r *DR33) Bits(mask DR33_Bits) DR33_Bits { return DR33_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR33) StoreBits(mask, b DR33_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR33) SetBits(mask DR33_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR33) ClearBits(mask DR33_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR33) Load() DR33_Bits               { return DR33_Bits(r.U16.Load()) }
func (r *DR33) Store(b DR33_Bits)             { r.U16.Store(uint16(b)) }

type DR33_Mask struct{ mmio.UM16 }

func (rm DR33_Mask) Load() DR33_Bits   { return DR33_Bits(rm.UM16.Load()) }
func (rm DR33_Mask) Store(b DR33_Bits) { rm.UM16.Store(uint16(b)) }

type DR34_Bits uint16

type DR34 struct{ mmio.U16 }

func (r *DR34) Bits(mask DR34_Bits) DR34_Bits { return DR34_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR34) StoreBits(mask, b DR34_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR34) SetBits(mask DR34_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR34) ClearBits(mask DR34_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR34) Load() DR34_Bits               { return DR34_Bits(r.U16.Load()) }
func (r *DR34) Store(b DR34_Bits)             { r.U16.Store(uint16(b)) }

type DR34_Mask struct{ mmio.UM16 }

func (rm DR34_Mask) Load() DR34_Bits   { return DR34_Bits(rm.UM16.Load()) }
func (rm DR34_Mask) Store(b DR34_Bits) { rm.UM16.Store(uint16(b)) }

type DR35_Bits uint16

type DR35 struct{ mmio.U16 }

func (r *DR35) Bits(mask DR35_Bits) DR35_Bits { return DR35_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR35) StoreBits(mask, b DR35_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR35) SetBits(mask DR35_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR35) ClearBits(mask DR35_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR35) Load() DR35_Bits               { return DR35_Bits(r.U16.Load()) }
func (r *DR35) Store(b DR35_Bits)             { r.U16.Store(uint16(b)) }

type DR35_Mask struct{ mmio.UM16 }

func (rm DR35_Mask) Load() DR35_Bits   { return DR35_Bits(rm.UM16.Load()) }
func (rm DR35_Mask) Store(b DR35_Bits) { rm.UM16.Store(uint16(b)) }

type DR36_Bits uint16

type DR36 struct{ mmio.U16 }

func (r *DR36) Bits(mask DR36_Bits) DR36_Bits { return DR36_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR36) StoreBits(mask, b DR36_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR36) SetBits(mask DR36_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR36) ClearBits(mask DR36_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR36) Load() DR36_Bits               { return DR36_Bits(r.U16.Load()) }
func (r *DR36) Store(b DR36_Bits)             { r.U16.Store(uint16(b)) }

type DR36_Mask struct{ mmio.UM16 }

func (rm DR36_Mask) Load() DR36_Bits   { return DR36_Bits(rm.UM16.Load()) }
func (rm DR36_Mask) Store(b DR36_Bits) { rm.UM16.Store(uint16(b)) }

type DR37_Bits uint16

type DR37 struct{ mmio.U16 }

func (r *DR37) Bits(mask DR37_Bits) DR37_Bits { return DR37_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR37) StoreBits(mask, b DR37_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR37) SetBits(mask DR37_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR37) ClearBits(mask DR37_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR37) Load() DR37_Bits               { return DR37_Bits(r.U16.Load()) }
func (r *DR37) Store(b DR37_Bits)             { r.U16.Store(uint16(b)) }

type DR37_Mask struct{ mmio.UM16 }

func (rm DR37_Mask) Load() DR37_Bits   { return DR37_Bits(rm.UM16.Load()) }
func (rm DR37_Mask) Store(b DR37_Bits) { rm.UM16.Store(uint16(b)) }

type DR38_Bits uint16

type DR38 struct{ mmio.U16 }

func (r *DR38) Bits(mask DR38_Bits) DR38_Bits { return DR38_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR38) StoreBits(mask, b DR38_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR38) SetBits(mask DR38_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR38) ClearBits(mask DR38_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR38) Load() DR38_Bits               { return DR38_Bits(r.U16.Load()) }
func (r *DR38) Store(b DR38_Bits)             { r.U16.Store(uint16(b)) }

type DR38_Mask struct{ mmio.UM16 }

func (rm DR38_Mask) Load() DR38_Bits   { return DR38_Bits(rm.UM16.Load()) }
func (rm DR38_Mask) Store(b DR38_Bits) { rm.UM16.Store(uint16(b)) }

type DR39_Bits uint16

type DR39 struct{ mmio.U16 }

func (r *DR39) Bits(mask DR39_Bits) DR39_Bits { return DR39_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR39) StoreBits(mask, b DR39_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR39) SetBits(mask DR39_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR39) ClearBits(mask DR39_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR39) Load() DR39_Bits               { return DR39_Bits(r.U16.Load()) }
func (r *DR39) Store(b DR39_Bits)             { r.U16.Store(uint16(b)) }

type DR39_Mask struct{ mmio.UM16 }

func (rm DR39_Mask) Load() DR39_Bits   { return DR39_Bits(rm.UM16.Load()) }
func (rm DR39_Mask) Store(b DR39_Bits) { rm.UM16.Store(uint16(b)) }

type DR40_Bits uint16

type DR40 struct{ mmio.U16 }

func (r *DR40) Bits(mask DR40_Bits) DR40_Bits { return DR40_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR40) StoreBits(mask, b DR40_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR40) SetBits(mask DR40_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR40) ClearBits(mask DR40_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR40) Load() DR40_Bits               { return DR40_Bits(r.U16.Load()) }
func (r *DR40) Store(b DR40_Bits)             { r.U16.Store(uint16(b)) }

type DR40_Mask struct{ mmio.UM16 }

func (rm DR40_Mask) Load() DR40_Bits   { return DR40_Bits(rm.UM16.Load()) }
func (rm DR40_Mask) Store(b DR40_Bits) { rm.UM16.Store(uint16(b)) }

type DR41_Bits uint16

type DR41 struct{ mmio.U16 }

func (r *DR41) Bits(mask DR41_Bits) DR41_Bits { return DR41_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR41) StoreBits(mask, b DR41_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR41) SetBits(mask DR41_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR41) ClearBits(mask DR41_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR41) Load() DR41_Bits               { return DR41_Bits(r.U16.Load()) }
func (r *DR41) Store(b DR41_Bits)             { r.U16.Store(uint16(b)) }

type DR41_Mask struct{ mmio.UM16 }

func (rm DR41_Mask) Load() DR41_Bits   { return DR41_Bits(rm.UM16.Load()) }
func (rm DR41_Mask) Store(b DR41_Bits) { rm.UM16.Store(uint16(b)) }

type DR42_Bits uint16

type DR42 struct{ mmio.U16 }

func (r *DR42) Bits(mask DR42_Bits) DR42_Bits { return DR42_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR42) StoreBits(mask, b DR42_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR42) SetBits(mask DR42_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DR42) ClearBits(mask DR42_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DR42) Load() DR42_Bits               { return DR42_Bits(r.U16.Load()) }
func (r *DR42) Store(b DR42_Bits)             { r.U16.Store(uint16(b)) }

type DR42_Mask struct{ mmio.UM16 }

func (rm DR42_Mask) Load() DR42_Bits   { return DR42_Bits(rm.UM16.Load()) }
func (rm DR42_Mask) Store(b DR42_Bits) { rm.UM16.Store(uint16(b)) }
