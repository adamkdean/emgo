// +build f411xe

package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type ADC_Periph struct {
	SR    RSR
	CR1   RCR1
	CR2   RCR2
	SMPR1 RSMPR1
	SMPR2 RSMPR2
	JOFR1 RJOFR1
	JOFR2 RJOFR2
	JOFR3 RJOFR3
	JOFR4 RJOFR4
	HTR   RHTR
	LTR   RLTR
	SQR1  RSQR1
	SQR2  RSQR2
	SQR3  RSQR3
	JSQR  RJSQR
	JDR   [4]RJDR
	DR    RDR
}

func (p *ADC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var ADC1 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC1_BASE)))

//emgo:const
var ADC2 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC2_BASE)))

//emgo:const
var ADC3 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC3_BASE)))

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

func (p *ADC_Periph) AWD() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(AWD)}}
}

func (p *ADC_Periph) EOC() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(EOC)}}
}

func (p *ADC_Periph) JEOC() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(JEOC)}}
}

func (p *ADC_Periph) JSTRT() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(JSTRT)}}
}

func (p *ADC_Periph) STRT() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(STRT)}}
}

func (p *ADC_Periph) OVR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(OVR)}}
}

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

func (p *ADC_Periph) AWDCH() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(AWDCH)}}
}

func (p *ADC_Periph) EOCIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(EOCIE)}}
}

func (p *ADC_Periph) AWDIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(AWDIE)}}
}

func (p *ADC_Periph) JEOCIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(JEOCIE)}}
}

func (p *ADC_Periph) SCAN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SCAN)}}
}

func (p *ADC_Periph) AWDSGL() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(AWDSGL)}}
}

func (p *ADC_Periph) JAUTO() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(JAUTO)}}
}

func (p *ADC_Periph) DISCEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DISCEN)}}
}

func (p *ADC_Periph) JDISCEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(JDISCEN)}}
}

func (p *ADC_Periph) DISCNUM() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DISCNUM)}}
}

func (p *ADC_Periph) JAWDEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(JAWDEN)}}
}

func (p *ADC_Periph) AWDEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(AWDEN)}}
}

func (p *ADC_Periph) RES() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RES)}}
}

func (p *ADC_Periph) OVRIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(OVRIE)}}
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

func (p *ADC_Periph) ADON() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ADON)}}
}

func (p *ADC_Periph) CONT() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(CONT)}}
}

func (p *ADC_Periph) DMA() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(DMA)}}
}

func (p *ADC_Periph) DDS() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(DDS)}}
}

func (p *ADC_Periph) EOCS() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(EOCS)}}
}

func (p *ADC_Periph) ALIGN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ALIGN)}}
}

func (p *ADC_Periph) JEXTSEL() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(JEXTSEL)}}
}

func (p *ADC_Periph) JEXTEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(JEXTEN)}}
}

func (p *ADC_Periph) JSWSTART() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(JSWSTART)}}
}

func (p *ADC_Periph) EXTSEL() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(EXTSEL)}}
}

func (p *ADC_Periph) EXTEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(EXTEN)}}
}

func (p *ADC_Periph) SWSTART() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(SWSTART)}}
}

type SMPR1 uint32

func (b SMPR1) Field(mask SMPR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR1) J(v int) SMPR1 {
	return SMPR1(bits.Make32(v, uint32(mask)))
}

type RSMPR1 struct{ mmio.U32 }

func (r *RSMPR1) Bits(mask SMPR1) SMPR1   { return SMPR1(r.U32.Bits(uint32(mask))) }
func (r *RSMPR1) StoreBits(mask, b SMPR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSMPR1) SetBits(mask SMPR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSMPR1) ClearBits(mask SMPR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSMPR1) Load() SMPR1             { return SMPR1(r.U32.Load()) }
func (r *RSMPR1) Store(b SMPR1)           { r.U32.Store(uint32(b)) }

func (r *RSMPR1) AtomicStoreBits(mask, b SMPR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSMPR1) AtomicSetBits(mask SMPR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSMPR1) AtomicClearBits(mask SMPR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSMPR1 struct{ mmio.UM32 }

func (rm RMSMPR1) Load() SMPR1   { return SMPR1(rm.UM32.Load()) }
func (rm RMSMPR1) Store(b SMPR1) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP10() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP10)}}
}

func (p *ADC_Periph) SMP11() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP11)}}
}

func (p *ADC_Periph) SMP12() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP12)}}
}

func (p *ADC_Periph) SMP13() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP13)}}
}

func (p *ADC_Periph) SMP14() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP14)}}
}

func (p *ADC_Periph) SMP15() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP15)}}
}

func (p *ADC_Periph) SMP16() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP16)}}
}

func (p *ADC_Periph) SMP17() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP17)}}
}

func (p *ADC_Periph) SMP18() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP18)}}
}

type SMPR2 uint32

func (b SMPR2) Field(mask SMPR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR2) J(v int) SMPR2 {
	return SMPR2(bits.Make32(v, uint32(mask)))
}

type RSMPR2 struct{ mmio.U32 }

func (r *RSMPR2) Bits(mask SMPR2) SMPR2   { return SMPR2(r.U32.Bits(uint32(mask))) }
func (r *RSMPR2) StoreBits(mask, b SMPR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSMPR2) SetBits(mask SMPR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSMPR2) ClearBits(mask SMPR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSMPR2) Load() SMPR2             { return SMPR2(r.U32.Load()) }
func (r *RSMPR2) Store(b SMPR2)           { r.U32.Store(uint32(b)) }

func (r *RSMPR2) AtomicStoreBits(mask, b SMPR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSMPR2) AtomicSetBits(mask SMPR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSMPR2) AtomicClearBits(mask SMPR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSMPR2 struct{ mmio.UM32 }

func (rm RMSMPR2) Load() SMPR2   { return SMPR2(rm.UM32.Load()) }
func (rm RMSMPR2) Store(b SMPR2) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP0() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP0)}}
}

func (p *ADC_Periph) SMP1() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP1)}}
}

func (p *ADC_Periph) SMP2() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP2)}}
}

func (p *ADC_Periph) SMP3() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP3)}}
}

func (p *ADC_Periph) SMP4() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP4)}}
}

func (p *ADC_Periph) SMP5() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP5)}}
}

func (p *ADC_Periph) SMP6() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP6)}}
}

func (p *ADC_Periph) SMP7() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP7)}}
}

func (p *ADC_Periph) SMP8() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP8)}}
}

func (p *ADC_Periph) SMP9() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP9)}}
}

type JOFR1 uint32

func (b JOFR1) Field(mask JOFR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR1) J(v int) JOFR1 {
	return JOFR1(bits.Make32(v, uint32(mask)))
}

type RJOFR1 struct{ mmio.U32 }

func (r *RJOFR1) Bits(mask JOFR1) JOFR1   { return JOFR1(r.U32.Bits(uint32(mask))) }
func (r *RJOFR1) StoreBits(mask, b JOFR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RJOFR1) SetBits(mask JOFR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RJOFR1) ClearBits(mask JOFR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RJOFR1) Load() JOFR1             { return JOFR1(r.U32.Load()) }
func (r *RJOFR1) Store(b JOFR1)           { r.U32.Store(uint32(b)) }

func (r *RJOFR1) AtomicStoreBits(mask, b JOFR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RJOFR1) AtomicSetBits(mask JOFR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RJOFR1) AtomicClearBits(mask JOFR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMJOFR1 struct{ mmio.UM32 }

func (rm RMJOFR1) Load() JOFR1   { return JOFR1(rm.UM32.Load()) }
func (rm RMJOFR1) Store(b JOFR1) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET1() RMJOFR1 {
	return RMJOFR1{mmio.UM32{&p.JOFR1.U32, uint32(JOFFSET1)}}
}

type JOFR2 uint32

func (b JOFR2) Field(mask JOFR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR2) J(v int) JOFR2 {
	return JOFR2(bits.Make32(v, uint32(mask)))
}

type RJOFR2 struct{ mmio.U32 }

func (r *RJOFR2) Bits(mask JOFR2) JOFR2   { return JOFR2(r.U32.Bits(uint32(mask))) }
func (r *RJOFR2) StoreBits(mask, b JOFR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RJOFR2) SetBits(mask JOFR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RJOFR2) ClearBits(mask JOFR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RJOFR2) Load() JOFR2             { return JOFR2(r.U32.Load()) }
func (r *RJOFR2) Store(b JOFR2)           { r.U32.Store(uint32(b)) }

func (r *RJOFR2) AtomicStoreBits(mask, b JOFR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RJOFR2) AtomicSetBits(mask JOFR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RJOFR2) AtomicClearBits(mask JOFR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMJOFR2 struct{ mmio.UM32 }

func (rm RMJOFR2) Load() JOFR2   { return JOFR2(rm.UM32.Load()) }
func (rm RMJOFR2) Store(b JOFR2) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET2() RMJOFR2 {
	return RMJOFR2{mmio.UM32{&p.JOFR2.U32, uint32(JOFFSET2)}}
}

type JOFR3 uint32

func (b JOFR3) Field(mask JOFR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR3) J(v int) JOFR3 {
	return JOFR3(bits.Make32(v, uint32(mask)))
}

type RJOFR3 struct{ mmio.U32 }

func (r *RJOFR3) Bits(mask JOFR3) JOFR3   { return JOFR3(r.U32.Bits(uint32(mask))) }
func (r *RJOFR3) StoreBits(mask, b JOFR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RJOFR3) SetBits(mask JOFR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RJOFR3) ClearBits(mask JOFR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RJOFR3) Load() JOFR3             { return JOFR3(r.U32.Load()) }
func (r *RJOFR3) Store(b JOFR3)           { r.U32.Store(uint32(b)) }

func (r *RJOFR3) AtomicStoreBits(mask, b JOFR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RJOFR3) AtomicSetBits(mask JOFR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RJOFR3) AtomicClearBits(mask JOFR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMJOFR3 struct{ mmio.UM32 }

func (rm RMJOFR3) Load() JOFR3   { return JOFR3(rm.UM32.Load()) }
func (rm RMJOFR3) Store(b JOFR3) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET3() RMJOFR3 {
	return RMJOFR3{mmio.UM32{&p.JOFR3.U32, uint32(JOFFSET3)}}
}

type JOFR4 uint32

func (b JOFR4) Field(mask JOFR4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR4) J(v int) JOFR4 {
	return JOFR4(bits.Make32(v, uint32(mask)))
}

type RJOFR4 struct{ mmio.U32 }

func (r *RJOFR4) Bits(mask JOFR4) JOFR4   { return JOFR4(r.U32.Bits(uint32(mask))) }
func (r *RJOFR4) StoreBits(mask, b JOFR4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RJOFR4) SetBits(mask JOFR4)      { r.U32.SetBits(uint32(mask)) }
func (r *RJOFR4) ClearBits(mask JOFR4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RJOFR4) Load() JOFR4             { return JOFR4(r.U32.Load()) }
func (r *RJOFR4) Store(b JOFR4)           { r.U32.Store(uint32(b)) }

func (r *RJOFR4) AtomicStoreBits(mask, b JOFR4) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RJOFR4) AtomicSetBits(mask JOFR4)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RJOFR4) AtomicClearBits(mask JOFR4)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMJOFR4 struct{ mmio.UM32 }

func (rm RMJOFR4) Load() JOFR4   { return JOFR4(rm.UM32.Load()) }
func (rm RMJOFR4) Store(b JOFR4) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET4() RMJOFR4 {
	return RMJOFR4{mmio.UM32{&p.JOFR4.U32, uint32(JOFFSET4)}}
}

type HTR uint32

func (b HTR) Field(mask HTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HTR) J(v int) HTR {
	return HTR(bits.Make32(v, uint32(mask)))
}

type RHTR struct{ mmio.U32 }

func (r *RHTR) Bits(mask HTR) HTR     { return HTR(r.U32.Bits(uint32(mask))) }
func (r *RHTR) StoreBits(mask, b HTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHTR) SetBits(mask HTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RHTR) ClearBits(mask HTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHTR) Load() HTR             { return HTR(r.U32.Load()) }
func (r *RHTR) Store(b HTR)           { r.U32.Store(uint32(b)) }

func (r *RHTR) AtomicStoreBits(mask, b HTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHTR) AtomicSetBits(mask HTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHTR) AtomicClearBits(mask HTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHTR struct{ mmio.UM32 }

func (rm RMHTR) Load() HTR   { return HTR(rm.UM32.Load()) }
func (rm RMHTR) Store(b HTR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) HT() RMHTR {
	return RMHTR{mmio.UM32{&p.HTR.U32, uint32(HT)}}
}

type LTR uint32

func (b LTR) Field(mask LTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LTR) J(v int) LTR {
	return LTR(bits.Make32(v, uint32(mask)))
}

type RLTR struct{ mmio.U32 }

func (r *RLTR) Bits(mask LTR) LTR     { return LTR(r.U32.Bits(uint32(mask))) }
func (r *RLTR) StoreBits(mask, b LTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLTR) SetBits(mask LTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLTR) ClearBits(mask LTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLTR) Load() LTR             { return LTR(r.U32.Load()) }
func (r *RLTR) Store(b LTR)           { r.U32.Store(uint32(b)) }

func (r *RLTR) AtomicStoreBits(mask, b LTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLTR) AtomicSetBits(mask LTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLTR) AtomicClearBits(mask LTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLTR struct{ mmio.UM32 }

func (rm RMLTR) Load() LTR   { return LTR(rm.UM32.Load()) }
func (rm RMLTR) Store(b LTR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT() RMLTR {
	return RMLTR{mmio.UM32{&p.LTR.U32, uint32(LT)}}
}

type SQR1 uint32

func (b SQR1) Field(mask SQR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR1) J(v int) SQR1 {
	return SQR1(bits.Make32(v, uint32(mask)))
}

type RSQR1 struct{ mmio.U32 }

func (r *RSQR1) Bits(mask SQR1) SQR1    { return SQR1(r.U32.Bits(uint32(mask))) }
func (r *RSQR1) StoreBits(mask, b SQR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSQR1) SetBits(mask SQR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSQR1) ClearBits(mask SQR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSQR1) Load() SQR1             { return SQR1(r.U32.Load()) }
func (r *RSQR1) Store(b SQR1)           { r.U32.Store(uint32(b)) }

func (r *RSQR1) AtomicStoreBits(mask, b SQR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSQR1) AtomicSetBits(mask SQR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSQR1) AtomicClearBits(mask SQR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSQR1 struct{ mmio.UM32 }

func (rm RMSQR1) Load() SQR1   { return SQR1(rm.UM32.Load()) }
func (rm RMSQR1) Store(b SQR1) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ13() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(SQ13)}}
}

func (p *ADC_Periph) SQ14() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(SQ14)}}
}

func (p *ADC_Periph) SQ15() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(SQ15)}}
}

func (p *ADC_Periph) SQ16() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(SQ16)}}
}

func (p *ADC_Periph) L() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(L)}}
}

type SQR2 uint32

func (b SQR2) Field(mask SQR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR2) J(v int) SQR2 {
	return SQR2(bits.Make32(v, uint32(mask)))
}

type RSQR2 struct{ mmio.U32 }

func (r *RSQR2) Bits(mask SQR2) SQR2    { return SQR2(r.U32.Bits(uint32(mask))) }
func (r *RSQR2) StoreBits(mask, b SQR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSQR2) SetBits(mask SQR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSQR2) ClearBits(mask SQR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSQR2) Load() SQR2             { return SQR2(r.U32.Load()) }
func (r *RSQR2) Store(b SQR2)           { r.U32.Store(uint32(b)) }

func (r *RSQR2) AtomicStoreBits(mask, b SQR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSQR2) AtomicSetBits(mask SQR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSQR2) AtomicClearBits(mask SQR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSQR2 struct{ mmio.UM32 }

func (rm RMSQR2) Load() SQR2   { return SQR2(rm.UM32.Load()) }
func (rm RMSQR2) Store(b SQR2) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ7() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ7)}}
}

func (p *ADC_Periph) SQ8() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ8)}}
}

func (p *ADC_Periph) SQ9() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ9)}}
}

func (p *ADC_Periph) SQ10() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ10)}}
}

func (p *ADC_Periph) SQ11() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ11)}}
}

func (p *ADC_Periph) SQ12() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ12)}}
}

type SQR3 uint32

func (b SQR3) Field(mask SQR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR3) J(v int) SQR3 {
	return SQR3(bits.Make32(v, uint32(mask)))
}

type RSQR3 struct{ mmio.U32 }

func (r *RSQR3) Bits(mask SQR3) SQR3    { return SQR3(r.U32.Bits(uint32(mask))) }
func (r *RSQR3) StoreBits(mask, b SQR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSQR3) SetBits(mask SQR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RSQR3) ClearBits(mask SQR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSQR3) Load() SQR3             { return SQR3(r.U32.Load()) }
func (r *RSQR3) Store(b SQR3)           { r.U32.Store(uint32(b)) }

func (r *RSQR3) AtomicStoreBits(mask, b SQR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSQR3) AtomicSetBits(mask SQR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSQR3) AtomicClearBits(mask SQR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSQR3 struct{ mmio.UM32 }

func (rm RMSQR3) Load() SQR3   { return SQR3(rm.UM32.Load()) }
func (rm RMSQR3) Store(b SQR3) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ1() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ1)}}
}

func (p *ADC_Periph) SQ2() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ2)}}
}

func (p *ADC_Periph) SQ3() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ3)}}
}

func (p *ADC_Periph) SQ4() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ4)}}
}

func (p *ADC_Periph) SQ5() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ5)}}
}

func (p *ADC_Periph) SQ6() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ6)}}
}

type JSQR uint32

func (b JSQR) Field(mask JSQR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JSQR) J(v int) JSQR {
	return JSQR(bits.Make32(v, uint32(mask)))
}

type RJSQR struct{ mmio.U32 }

func (r *RJSQR) Bits(mask JSQR) JSQR    { return JSQR(r.U32.Bits(uint32(mask))) }
func (r *RJSQR) StoreBits(mask, b JSQR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RJSQR) SetBits(mask JSQR)      { r.U32.SetBits(uint32(mask)) }
func (r *RJSQR) ClearBits(mask JSQR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RJSQR) Load() JSQR             { return JSQR(r.U32.Load()) }
func (r *RJSQR) Store(b JSQR)           { r.U32.Store(uint32(b)) }

func (r *RJSQR) AtomicStoreBits(mask, b JSQR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RJSQR) AtomicSetBits(mask JSQR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RJSQR) AtomicClearBits(mask JSQR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMJSQR struct{ mmio.UM32 }

func (rm RMJSQR) Load() JSQR   { return JSQR(rm.UM32.Load()) }
func (rm RMJSQR) Store(b JSQR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JSQ1() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JSQ1)}}
}

func (p *ADC_Periph) JSQ2() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JSQ2)}}
}

func (p *ADC_Periph) JSQ3() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JSQ3)}}
}

func (p *ADC_Periph) JSQ4() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JSQ4)}}
}

func (p *ADC_Periph) JL() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JL)}}
}

type JDR uint32

func (b JDR) Field(mask JDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JDR) J(v int) JDR {
	return JDR(bits.Make32(v, uint32(mask)))
}

type RJDR struct{ mmio.U32 }

func (r *RJDR) Bits(mask JDR) JDR     { return JDR(r.U32.Bits(uint32(mask))) }
func (r *RJDR) StoreBits(mask, b JDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RJDR) SetBits(mask JDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RJDR) ClearBits(mask JDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RJDR) Load() JDR             { return JDR(r.U32.Load()) }
func (r *RJDR) Store(b JDR)           { r.U32.Store(uint32(b)) }

func (r *RJDR) AtomicStoreBits(mask, b JDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RJDR) AtomicSetBits(mask JDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RJDR) AtomicClearBits(mask JDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMJDR struct{ mmio.UM32 }

func (rm RMJDR) Load() JDR   { return JDR(rm.UM32.Load()) }
func (rm RMJDR) Store(b JDR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JDATA(n int) RMJDR {
	return RMJDR{mmio.UM32{&p.JDR[n].U32, uint32(JDATA)}}
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

func (p *ADC_Periph) DATA() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(DATA)}}
}

func (p *ADC_Periph) ADC2DATA() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(ADC2DATA)}}
}
