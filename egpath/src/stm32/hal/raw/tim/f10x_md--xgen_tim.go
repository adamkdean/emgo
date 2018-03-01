// +build f10x_md

package tim

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type TIM_Periph struct {
	CR1   RCR1
	_     uint16
	CR2   RCR2
	_     uint16
	SMCR  RSMCR
	_     uint16
	DIER  RDIER
	_     uint16
	SR    RSR
	_     uint16
	EGR   REGR
	_     uint16
	CCMR1 RCCMR1
	_     uint16
	CCMR2 RCCMR2
	_     uint16
	CCER  RCCER
	_     uint16
	CNT   RCNT
	_     uint16
	PSC   RPSC
	_     uint16
	ARR   RARR
	_     uint16
	RCR   RRCR
	_     uint16
	CCR1  RCCR1
	_     uint16
	CCR2  RCCR2
	_     uint16
	CCR3  RCCR3
	_     uint16
	CCR4  RCCR4
	_     uint16
	BDTR  RBDTR
	_     uint16
	DCR   RDCR
	_     uint16
	DMAR  RDMAR
}

func (p *TIM_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var TIM2 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM2_BASE)))

//emgo:const
var TIM3 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM3_BASE)))

//emgo:const
var TIM4 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM4_BASE)))

//emgo:const
var TIM5 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM5_BASE)))

//emgo:const
var TIM6 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM6_BASE)))

//emgo:const
var TIM7 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM7_BASE)))

//emgo:const
var TIM12 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM12_BASE)))

//emgo:const
var TIM13 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM13_BASE)))

//emgo:const
var TIM14 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM14_BASE)))

//emgo:const
var TIM1 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM1_BASE)))

//emgo:const
var TIM8 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM8_BASE)))

//emgo:const
var TIM15 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM15_BASE)))

//emgo:const
var TIM16 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM16_BASE)))

//emgo:const
var TIM17 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM17_BASE)))

//emgo:const
var TIM9 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM9_BASE)))

//emgo:const
var TIM10 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM10_BASE)))

//emgo:const
var TIM11 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM11_BASE)))

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

func (p *TIM_Periph) CEN() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(CEN)}}
}

func (p *TIM_Periph) UDIS() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(UDIS)}}
}

func (p *TIM_Periph) URS() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(URS)}}
}

func (p *TIM_Periph) OPM() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(OPM)}}
}

func (p *TIM_Periph) DIR() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(DIR)}}
}

func (p *TIM_Periph) CMS() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(CMS)}}
}

func (p *TIM_Periph) ARPE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(ARPE)}}
}

func (p *TIM_Periph) CKD() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(CKD)}}
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

func (p *TIM_Periph) CCPC() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(CCPC)}}
}

func (p *TIM_Periph) CCUS() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(CCUS)}}
}

func (p *TIM_Periph) CCDS() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(CCDS)}}
}

func (p *TIM_Periph) MMS() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(MMS)}}
}

func (p *TIM_Periph) TI1S() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(TI1S)}}
}

func (p *TIM_Periph) OIS1() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(OIS1)}}
}

func (p *TIM_Periph) OIS1N() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(OIS1N)}}
}

func (p *TIM_Periph) OIS2() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(OIS2)}}
}

func (p *TIM_Periph) OIS2N() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(OIS2N)}}
}

func (p *TIM_Periph) OIS3() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(OIS3)}}
}

func (p *TIM_Periph) OIS3N() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(OIS3N)}}
}

func (p *TIM_Periph) OIS4() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(OIS4)}}
}

type SMCR uint16

func (b SMCR) Field(mask SMCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMCR) J(v int) SMCR {
	return SMCR(bits.Make32(v, uint32(mask)))
}

type RSMCR struct{ mmio.U16 }

func (r *RSMCR) Bits(mask SMCR) SMCR    { return SMCR(r.U16.Bits(uint16(mask))) }
func (r *RSMCR) StoreBits(mask, b SMCR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RSMCR) SetBits(mask SMCR)      { r.U16.SetBits(uint16(mask)) }
func (r *RSMCR) ClearBits(mask SMCR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RSMCR) Load() SMCR             { return SMCR(r.U16.Load()) }
func (r *RSMCR) Store(b SMCR)           { r.U16.Store(uint16(b)) }

type RMSMCR struct{ mmio.UM16 }

func (rm RMSMCR) Load() SMCR   { return SMCR(rm.UM16.Load()) }
func (rm RMSMCR) Store(b SMCR) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) SMS() RMSMCR {
	return RMSMCR{mmio.UM16{&p.SMCR.U16, uint16(SMS)}}
}

func (p *TIM_Periph) TS() RMSMCR {
	return RMSMCR{mmio.UM16{&p.SMCR.U16, uint16(TS)}}
}

func (p *TIM_Periph) MSM() RMSMCR {
	return RMSMCR{mmio.UM16{&p.SMCR.U16, uint16(MSM)}}
}

func (p *TIM_Periph) ETF() RMSMCR {
	return RMSMCR{mmio.UM16{&p.SMCR.U16, uint16(ETF)}}
}

func (p *TIM_Periph) ETPS() RMSMCR {
	return RMSMCR{mmio.UM16{&p.SMCR.U16, uint16(ETPS)}}
}

func (p *TIM_Periph) ECE() RMSMCR {
	return RMSMCR{mmio.UM16{&p.SMCR.U16, uint16(ECE)}}
}

func (p *TIM_Periph) ETP() RMSMCR {
	return RMSMCR{mmio.UM16{&p.SMCR.U16, uint16(ETP)}}
}

type DIER uint16

func (b DIER) Field(mask DIER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIER) J(v int) DIER {
	return DIER(bits.Make32(v, uint32(mask)))
}

type RDIER struct{ mmio.U16 }

func (r *RDIER) Bits(mask DIER) DIER    { return DIER(r.U16.Bits(uint16(mask))) }
func (r *RDIER) StoreBits(mask, b DIER) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDIER) SetBits(mask DIER)      { r.U16.SetBits(uint16(mask)) }
func (r *RDIER) ClearBits(mask DIER)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDIER) Load() DIER             { return DIER(r.U16.Load()) }
func (r *RDIER) Store(b DIER)           { r.U16.Store(uint16(b)) }

type RMDIER struct{ mmio.UM16 }

func (rm RMDIER) Load() DIER   { return DIER(rm.UM16.Load()) }
func (rm RMDIER) Store(b DIER) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) UIE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(UIE)}}
}

func (p *TIM_Periph) CC1IE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(CC1IE)}}
}

func (p *TIM_Periph) CC2IE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(CC2IE)}}
}

func (p *TIM_Periph) CC3IE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(CC3IE)}}
}

func (p *TIM_Periph) CC4IE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(CC4IE)}}
}

func (p *TIM_Periph) COMIE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(COMIE)}}
}

func (p *TIM_Periph) TIE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(TIE)}}
}

func (p *TIM_Periph) BIE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(BIE)}}
}

func (p *TIM_Periph) UDE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(UDE)}}
}

func (p *TIM_Periph) CC1DE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(CC1DE)}}
}

func (p *TIM_Periph) CC2DE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(CC2DE)}}
}

func (p *TIM_Periph) CC3DE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(CC3DE)}}
}

func (p *TIM_Periph) CC4DE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(CC4DE)}}
}

func (p *TIM_Periph) COMDE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(COMDE)}}
}

func (p *TIM_Periph) TDE() RMDIER {
	return RMDIER{mmio.UM16{&p.DIER.U16, uint16(TDE)}}
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

func (p *TIM_Periph) UIF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(UIF)}}
}

func (p *TIM_Periph) CC1IF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CC1IF)}}
}

func (p *TIM_Periph) CC2IF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CC2IF)}}
}

func (p *TIM_Periph) CC3IF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CC3IF)}}
}

func (p *TIM_Periph) CC4IF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CC4IF)}}
}

func (p *TIM_Periph) COMIF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(COMIF)}}
}

func (p *TIM_Periph) TIF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(TIF)}}
}

func (p *TIM_Periph) BIF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(BIF)}}
}

func (p *TIM_Periph) CC1OF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CC1OF)}}
}

func (p *TIM_Periph) CC2OF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CC2OF)}}
}

func (p *TIM_Periph) CC3OF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CC3OF)}}
}

func (p *TIM_Periph) CC4OF() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CC4OF)}}
}

type EGR uint16

func (b EGR) Field(mask EGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EGR) J(v int) EGR {
	return EGR(bits.Make32(v, uint32(mask)))
}

type REGR struct{ mmio.U16 }

func (r *REGR) Bits(mask EGR) EGR     { return EGR(r.U16.Bits(uint16(mask))) }
func (r *REGR) StoreBits(mask, b EGR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *REGR) SetBits(mask EGR)      { r.U16.SetBits(uint16(mask)) }
func (r *REGR) ClearBits(mask EGR)    { r.U16.ClearBits(uint16(mask)) }
func (r *REGR) Load() EGR             { return EGR(r.U16.Load()) }
func (r *REGR) Store(b EGR)           { r.U16.Store(uint16(b)) }

type RMEGR struct{ mmio.UM16 }

func (rm RMEGR) Load() EGR   { return EGR(rm.UM16.Load()) }
func (rm RMEGR) Store(b EGR) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) UG() RMEGR {
	return RMEGR{mmio.UM16{&p.EGR.U16, uint16(UG)}}
}

func (p *TIM_Periph) CC1G() RMEGR {
	return RMEGR{mmio.UM16{&p.EGR.U16, uint16(CC1G)}}
}

func (p *TIM_Periph) CC2G() RMEGR {
	return RMEGR{mmio.UM16{&p.EGR.U16, uint16(CC2G)}}
}

func (p *TIM_Periph) CC3G() RMEGR {
	return RMEGR{mmio.UM16{&p.EGR.U16, uint16(CC3G)}}
}

func (p *TIM_Periph) CC4G() RMEGR {
	return RMEGR{mmio.UM16{&p.EGR.U16, uint16(CC4G)}}
}

func (p *TIM_Periph) COMG() RMEGR {
	return RMEGR{mmio.UM16{&p.EGR.U16, uint16(COMG)}}
}

func (p *TIM_Periph) TG() RMEGR {
	return RMEGR{mmio.UM16{&p.EGR.U16, uint16(TG)}}
}

func (p *TIM_Periph) BG() RMEGR {
	return RMEGR{mmio.UM16{&p.EGR.U16, uint16(BG)}}
}

type CCMR1 uint16

func (b CCMR1) Field(mask CCMR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCMR1) J(v int) CCMR1 {
	return CCMR1(bits.Make32(v, uint32(mask)))
}

type RCCMR1 struct{ mmio.U16 }

func (r *RCCMR1) Bits(mask CCMR1) CCMR1   { return CCMR1(r.U16.Bits(uint16(mask))) }
func (r *RCCMR1) StoreBits(mask, b CCMR1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCCMR1) SetBits(mask CCMR1)      { r.U16.SetBits(uint16(mask)) }
func (r *RCCMR1) ClearBits(mask CCMR1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCCMR1) Load() CCMR1             { return CCMR1(r.U16.Load()) }
func (r *RCCMR1) Store(b CCMR1)           { r.U16.Store(uint16(b)) }

type RMCCMR1 struct{ mmio.UM16 }

func (rm RMCCMR1) Load() CCMR1   { return CCMR1(rm.UM16.Load()) }
func (rm RMCCMR1) Store(b CCMR1) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) CC1S() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(CC1S)}}
}

func (p *TIM_Periph) OC1FE() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(OC1FE)}}
}

func (p *TIM_Periph) OC1PE() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(OC1PE)}}
}

func (p *TIM_Periph) OC1M() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(OC1M)}}
}

func (p *TIM_Periph) OC1CE() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(OC1CE)}}
}

func (p *TIM_Periph) CC2S() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(CC2S)}}
}

func (p *TIM_Periph) OC2FE() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(OC2FE)}}
}

func (p *TIM_Periph) OC2PE() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(OC2PE)}}
}

func (p *TIM_Periph) OC2M() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(OC2M)}}
}

func (p *TIM_Periph) OC2CE() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(OC2CE)}}
}

func (p *TIM_Periph) IC1PSC() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(IC1PSC)}}
}

func (p *TIM_Periph) IC1F() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(IC1F)}}
}

func (p *TIM_Periph) IC2PSC() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(IC2PSC)}}
}

func (p *TIM_Periph) IC2F() RMCCMR1 {
	return RMCCMR1{mmio.UM16{&p.CCMR1.U16, uint16(IC2F)}}
}

type CCMR2 uint16

func (b CCMR2) Field(mask CCMR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCMR2) J(v int) CCMR2 {
	return CCMR2(bits.Make32(v, uint32(mask)))
}

type RCCMR2 struct{ mmio.U16 }

func (r *RCCMR2) Bits(mask CCMR2) CCMR2   { return CCMR2(r.U16.Bits(uint16(mask))) }
func (r *RCCMR2) StoreBits(mask, b CCMR2) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCCMR2) SetBits(mask CCMR2)      { r.U16.SetBits(uint16(mask)) }
func (r *RCCMR2) ClearBits(mask CCMR2)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCCMR2) Load() CCMR2             { return CCMR2(r.U16.Load()) }
func (r *RCCMR2) Store(b CCMR2)           { r.U16.Store(uint16(b)) }

type RMCCMR2 struct{ mmio.UM16 }

func (rm RMCCMR2) Load() CCMR2   { return CCMR2(rm.UM16.Load()) }
func (rm RMCCMR2) Store(b CCMR2) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) CC3S() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(CC3S)}}
}

func (p *TIM_Periph) OC3FE() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(OC3FE)}}
}

func (p *TIM_Periph) OC3PE() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(OC3PE)}}
}

func (p *TIM_Periph) OC3M() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(OC3M)}}
}

func (p *TIM_Periph) OC3CE() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(OC3CE)}}
}

func (p *TIM_Periph) CC4S() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(CC4S)}}
}

func (p *TIM_Periph) OC4FE() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(OC4FE)}}
}

func (p *TIM_Periph) OC4PE() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(OC4PE)}}
}

func (p *TIM_Periph) OC4M() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(OC4M)}}
}

func (p *TIM_Periph) OC4CE() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(OC4CE)}}
}

func (p *TIM_Periph) IC3PSC() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(IC3PSC)}}
}

func (p *TIM_Periph) IC3F() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(IC3F)}}
}

func (p *TIM_Periph) IC4PSC() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(IC4PSC)}}
}

func (p *TIM_Periph) IC4F() RMCCMR2 {
	return RMCCMR2{mmio.UM16{&p.CCMR2.U16, uint16(IC4F)}}
}

type CCER uint16

func (b CCER) Field(mask CCER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCER) J(v int) CCER {
	return CCER(bits.Make32(v, uint32(mask)))
}

type RCCER struct{ mmio.U16 }

func (r *RCCER) Bits(mask CCER) CCER    { return CCER(r.U16.Bits(uint16(mask))) }
func (r *RCCER) StoreBits(mask, b CCER) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCCER) SetBits(mask CCER)      { r.U16.SetBits(uint16(mask)) }
func (r *RCCER) ClearBits(mask CCER)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCCER) Load() CCER             { return CCER(r.U16.Load()) }
func (r *RCCER) Store(b CCER)           { r.U16.Store(uint16(b)) }

type RMCCER struct{ mmio.UM16 }

func (rm RMCCER) Load() CCER   { return CCER(rm.UM16.Load()) }
func (rm RMCCER) Store(b CCER) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) CC1E() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC1E)}}
}

func (p *TIM_Periph) CC1P() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC1P)}}
}

func (p *TIM_Periph) CC1NE() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC1NE)}}
}

func (p *TIM_Periph) CC1NP() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC1NP)}}
}

func (p *TIM_Periph) CC2E() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC2E)}}
}

func (p *TIM_Periph) CC2P() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC2P)}}
}

func (p *TIM_Periph) CC2NE() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC2NE)}}
}

func (p *TIM_Periph) CC2NP() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC2NP)}}
}

func (p *TIM_Periph) CC3E() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC3E)}}
}

func (p *TIM_Periph) CC3P() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC3P)}}
}

func (p *TIM_Periph) CC3NE() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC3NE)}}
}

func (p *TIM_Periph) CC3NP() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC3NP)}}
}

func (p *TIM_Periph) CC4E() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC4E)}}
}

func (p *TIM_Periph) CC4P() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC4P)}}
}

func (p *TIM_Periph) CC4NP() RMCCER {
	return RMCCER{mmio.UM16{&p.CCER.U16, uint16(CC4NP)}}
}

type CNT uint16

func (b CNT) Field(mask CNT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CNT) J(v int) CNT {
	return CNT(bits.Make32(v, uint32(mask)))
}

type RCNT struct{ mmio.U16 }

func (r *RCNT) Bits(mask CNT) CNT     { return CNT(r.U16.Bits(uint16(mask))) }
func (r *RCNT) StoreBits(mask, b CNT) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCNT) SetBits(mask CNT)      { r.U16.SetBits(uint16(mask)) }
func (r *RCNT) ClearBits(mask CNT)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCNT) Load() CNT             { return CNT(r.U16.Load()) }
func (r *RCNT) Store(b CNT)           { r.U16.Store(uint16(b)) }

type RMCNT struct{ mmio.UM16 }

func (rm RMCNT) Load() CNT   { return CNT(rm.UM16.Load()) }
func (rm RMCNT) Store(b CNT) { rm.UM16.Store(uint16(b)) }

type PSC uint16

func (b PSC) Field(mask PSC) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PSC) J(v int) PSC {
	return PSC(bits.Make32(v, uint32(mask)))
}

type RPSC struct{ mmio.U16 }

func (r *RPSC) Bits(mask PSC) PSC     { return PSC(r.U16.Bits(uint16(mask))) }
func (r *RPSC) StoreBits(mask, b PSC) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RPSC) SetBits(mask PSC)      { r.U16.SetBits(uint16(mask)) }
func (r *RPSC) ClearBits(mask PSC)    { r.U16.ClearBits(uint16(mask)) }
func (r *RPSC) Load() PSC             { return PSC(r.U16.Load()) }
func (r *RPSC) Store(b PSC)           { r.U16.Store(uint16(b)) }

type RMPSC struct{ mmio.UM16 }

func (rm RMPSC) Load() PSC   { return PSC(rm.UM16.Load()) }
func (rm RMPSC) Store(b PSC) { rm.UM16.Store(uint16(b)) }

type ARR uint16

func (b ARR) Field(mask ARR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ARR) J(v int) ARR {
	return ARR(bits.Make32(v, uint32(mask)))
}

type RARR struct{ mmio.U16 }

func (r *RARR) Bits(mask ARR) ARR     { return ARR(r.U16.Bits(uint16(mask))) }
func (r *RARR) StoreBits(mask, b ARR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RARR) SetBits(mask ARR)      { r.U16.SetBits(uint16(mask)) }
func (r *RARR) ClearBits(mask ARR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RARR) Load() ARR             { return ARR(r.U16.Load()) }
func (r *RARR) Store(b ARR)           { r.U16.Store(uint16(b)) }

type RMARR struct{ mmio.UM16 }

func (rm RMARR) Load() ARR   { return ARR(rm.UM16.Load()) }
func (rm RMARR) Store(b ARR) { rm.UM16.Store(uint16(b)) }

type RCR uint16

func (b RCR) Field(mask RCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RCR) J(v int) RCR {
	return RCR(bits.Make32(v, uint32(mask)))
}

type RRCR struct{ mmio.U16 }

func (r *RRCR) Bits(mask RCR) RCR     { return RCR(r.U16.Bits(uint16(mask))) }
func (r *RRCR) StoreBits(mask, b RCR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RRCR) SetBits(mask RCR)      { r.U16.SetBits(uint16(mask)) }
func (r *RRCR) ClearBits(mask RCR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RRCR) Load() RCR             { return RCR(r.U16.Load()) }
func (r *RRCR) Store(b RCR)           { r.U16.Store(uint16(b)) }

type RMRCR struct{ mmio.UM16 }

func (rm RMRCR) Load() RCR   { return RCR(rm.UM16.Load()) }
func (rm RMRCR) Store(b RCR) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) REP() RMRCR {
	return RMRCR{mmio.UM16{&p.RCR.U16, uint16(REP)}}
}

type CCR1 uint16

func (b CCR1) Field(mask CCR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR1) J(v int) CCR1 {
	return CCR1(bits.Make32(v, uint32(mask)))
}

type RCCR1 struct{ mmio.U16 }

func (r *RCCR1) Bits(mask CCR1) CCR1    { return CCR1(r.U16.Bits(uint16(mask))) }
func (r *RCCR1) StoreBits(mask, b CCR1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCCR1) SetBits(mask CCR1)      { r.U16.SetBits(uint16(mask)) }
func (r *RCCR1) ClearBits(mask CCR1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCCR1) Load() CCR1             { return CCR1(r.U16.Load()) }
func (r *RCCR1) Store(b CCR1)           { r.U16.Store(uint16(b)) }

type RMCCR1 struct{ mmio.UM16 }

func (rm RMCCR1) Load() CCR1   { return CCR1(rm.UM16.Load()) }
func (rm RMCCR1) Store(b CCR1) { rm.UM16.Store(uint16(b)) }

type CCR2 uint16

func (b CCR2) Field(mask CCR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR2) J(v int) CCR2 {
	return CCR2(bits.Make32(v, uint32(mask)))
}

type RCCR2 struct{ mmio.U16 }

func (r *RCCR2) Bits(mask CCR2) CCR2    { return CCR2(r.U16.Bits(uint16(mask))) }
func (r *RCCR2) StoreBits(mask, b CCR2) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCCR2) SetBits(mask CCR2)      { r.U16.SetBits(uint16(mask)) }
func (r *RCCR2) ClearBits(mask CCR2)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCCR2) Load() CCR2             { return CCR2(r.U16.Load()) }
func (r *RCCR2) Store(b CCR2)           { r.U16.Store(uint16(b)) }

type RMCCR2 struct{ mmio.UM16 }

func (rm RMCCR2) Load() CCR2   { return CCR2(rm.UM16.Load()) }
func (rm RMCCR2) Store(b CCR2) { rm.UM16.Store(uint16(b)) }

type CCR3 uint16

func (b CCR3) Field(mask CCR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR3) J(v int) CCR3 {
	return CCR3(bits.Make32(v, uint32(mask)))
}

type RCCR3 struct{ mmio.U16 }

func (r *RCCR3) Bits(mask CCR3) CCR3    { return CCR3(r.U16.Bits(uint16(mask))) }
func (r *RCCR3) StoreBits(mask, b CCR3) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCCR3) SetBits(mask CCR3)      { r.U16.SetBits(uint16(mask)) }
func (r *RCCR3) ClearBits(mask CCR3)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCCR3) Load() CCR3             { return CCR3(r.U16.Load()) }
func (r *RCCR3) Store(b CCR3)           { r.U16.Store(uint16(b)) }

type RMCCR3 struct{ mmio.UM16 }

func (rm RMCCR3) Load() CCR3   { return CCR3(rm.UM16.Load()) }
func (rm RMCCR3) Store(b CCR3) { rm.UM16.Store(uint16(b)) }

type CCR4 uint16

func (b CCR4) Field(mask CCR4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR4) J(v int) CCR4 {
	return CCR4(bits.Make32(v, uint32(mask)))
}

type RCCR4 struct{ mmio.U16 }

func (r *RCCR4) Bits(mask CCR4) CCR4    { return CCR4(r.U16.Bits(uint16(mask))) }
func (r *RCCR4) StoreBits(mask, b CCR4) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCCR4) SetBits(mask CCR4)      { r.U16.SetBits(uint16(mask)) }
func (r *RCCR4) ClearBits(mask CCR4)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCCR4) Load() CCR4             { return CCR4(r.U16.Load()) }
func (r *RCCR4) Store(b CCR4)           { r.U16.Store(uint16(b)) }

type RMCCR4 struct{ mmio.UM16 }

func (rm RMCCR4) Load() CCR4   { return CCR4(rm.UM16.Load()) }
func (rm RMCCR4) Store(b CCR4) { rm.UM16.Store(uint16(b)) }

type BDTR uint16

func (b BDTR) Field(mask BDTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BDTR) J(v int) BDTR {
	return BDTR(bits.Make32(v, uint32(mask)))
}

type RBDTR struct{ mmio.U16 }

func (r *RBDTR) Bits(mask BDTR) BDTR    { return BDTR(r.U16.Bits(uint16(mask))) }
func (r *RBDTR) StoreBits(mask, b BDTR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RBDTR) SetBits(mask BDTR)      { r.U16.SetBits(uint16(mask)) }
func (r *RBDTR) ClearBits(mask BDTR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RBDTR) Load() BDTR             { return BDTR(r.U16.Load()) }
func (r *RBDTR) Store(b BDTR)           { r.U16.Store(uint16(b)) }

type RMBDTR struct{ mmio.UM16 }

func (rm RMBDTR) Load() BDTR   { return BDTR(rm.UM16.Load()) }
func (rm RMBDTR) Store(b BDTR) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) DTG() RMBDTR {
	return RMBDTR{mmio.UM16{&p.BDTR.U16, uint16(DTG)}}
}

func (p *TIM_Periph) LOCK() RMBDTR {
	return RMBDTR{mmio.UM16{&p.BDTR.U16, uint16(LOCK)}}
}

func (p *TIM_Periph) OSSI() RMBDTR {
	return RMBDTR{mmio.UM16{&p.BDTR.U16, uint16(OSSI)}}
}

func (p *TIM_Periph) OSSR() RMBDTR {
	return RMBDTR{mmio.UM16{&p.BDTR.U16, uint16(OSSR)}}
}

func (p *TIM_Periph) BKE() RMBDTR {
	return RMBDTR{mmio.UM16{&p.BDTR.U16, uint16(BKE)}}
}

func (p *TIM_Periph) BKP() RMBDTR {
	return RMBDTR{mmio.UM16{&p.BDTR.U16, uint16(BKP)}}
}

func (p *TIM_Periph) AOE() RMBDTR {
	return RMBDTR{mmio.UM16{&p.BDTR.U16, uint16(AOE)}}
}

func (p *TIM_Periph) MOE() RMBDTR {
	return RMBDTR{mmio.UM16{&p.BDTR.U16, uint16(MOE)}}
}

type DCR uint16

func (b DCR) Field(mask DCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCR) J(v int) DCR {
	return DCR(bits.Make32(v, uint32(mask)))
}

type RDCR struct{ mmio.U16 }

func (r *RDCR) Bits(mask DCR) DCR     { return DCR(r.U16.Bits(uint16(mask))) }
func (r *RDCR) StoreBits(mask, b DCR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDCR) SetBits(mask DCR)      { r.U16.SetBits(uint16(mask)) }
func (r *RDCR) ClearBits(mask DCR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDCR) Load() DCR             { return DCR(r.U16.Load()) }
func (r *RDCR) Store(b DCR)           { r.U16.Store(uint16(b)) }

type RMDCR struct{ mmio.UM16 }

func (rm RMDCR) Load() DCR   { return DCR(rm.UM16.Load()) }
func (rm RMDCR) Store(b DCR) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) DBA() RMDCR {
	return RMDCR{mmio.UM16{&p.DCR.U16, uint16(DBA)}}
}

func (p *TIM_Periph) DBL() RMDCR {
	return RMDCR{mmio.UM16{&p.DCR.U16, uint16(DBL)}}
}

type DMAR uint16

func (b DMAR) Field(mask DMAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DMAR) J(v int) DMAR {
	return DMAR(bits.Make32(v, uint32(mask)))
}

type RDMAR struct{ mmio.U16 }

func (r *RDMAR) Bits(mask DMAR) DMAR    { return DMAR(r.U16.Bits(uint16(mask))) }
func (r *RDMAR) StoreBits(mask, b DMAR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDMAR) SetBits(mask DMAR)      { r.U16.SetBits(uint16(mask)) }
func (r *RDMAR) ClearBits(mask DMAR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDMAR) Load() DMAR             { return DMAR(r.U16.Load()) }
func (r *RDMAR) Store(b DMAR)           { r.U16.Store(uint16(b)) }

type RMDMAR struct{ mmio.UM16 }

func (rm RMDMAR) Load() DMAR   { return DMAR(rm.UM16.Load()) }
func (rm RMDMAR) Store(b DMAR) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) DMAB() RMDMAR {
	return RMDMAR{mmio.UM16{&p.DMAR.U16, uint16(DMAB)}}
}
