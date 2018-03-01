package rcc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type RCC_Periph struct {
	CR        RCR
	ICSCR     RICSCR
	CFGR      RCFGR
	CIR       RCIR
	AHBRSTR   RAHBRSTR
	APB2RSTR  RAPB2RSTR
	APB1RSTR  RAPB1RSTR
	AHBENR    RAHBENR
	APB2ENR   RAPB2ENR
	APB1ENR   RAPB1ENR
	AHBLPENR  RAHBLPENR
	APB2LPENR RAPB2LPENR
	APB1LPENR RAPB1LPENR
	CSR       RCSR
}

func (p *RCC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var RCC = (*RCC_Periph)(unsafe.Pointer(uintptr(mmap.RCC_BASE)))

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

func (p *RCC_Periph) HSION() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSION)}}
}

func (p *RCC_Periph) HSIRDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSIRDY)}}
}

func (p *RCC_Periph) MSION() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MSION)}}
}

func (p *RCC_Periph) MSIRDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MSIRDY)}}
}

func (p *RCC_Periph) HSEON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSEON)}}
}

func (p *RCC_Periph) HSERDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSERDY)}}
}

func (p *RCC_Periph) HSEBYP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSEBYP)}}
}

func (p *RCC_Periph) PLLON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLLON)}}
}

func (p *RCC_Periph) PLLRDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLLRDY)}}
}

func (p *RCC_Periph) CSSON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CSSON)}}
}

func (p *RCC_Periph) RTCPRE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(RTCPRE)}}
}

type ICSCR uint32

func (b ICSCR) Field(mask ICSCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICSCR) J(v int) ICSCR {
	return ICSCR(bits.Make32(v, uint32(mask)))
}

type RICSCR struct{ mmio.U32 }

func (r *RICSCR) Bits(mask ICSCR) ICSCR   { return ICSCR(r.U32.Bits(uint32(mask))) }
func (r *RICSCR) StoreBits(mask, b ICSCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICSCR) SetBits(mask ICSCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICSCR) ClearBits(mask ICSCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICSCR) Load() ICSCR             { return ICSCR(r.U32.Load()) }
func (r *RICSCR) Store(b ICSCR)           { r.U32.Store(uint32(b)) }

func (r *RICSCR) AtomicStoreBits(mask, b ICSCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RICSCR) AtomicSetBits(mask ICSCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RICSCR) AtomicClearBits(mask ICSCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMICSCR struct{ mmio.UM32 }

func (rm RMICSCR) Load() ICSCR   { return ICSCR(rm.UM32.Load()) }
func (rm RMICSCR) Store(b ICSCR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) HSICAL() RMICSCR {
	return RMICSCR{mmio.UM32{&p.ICSCR.U32, uint32(HSICAL)}}
}

func (p *RCC_Periph) HSITRIM() RMICSCR {
	return RMICSCR{mmio.UM32{&p.ICSCR.U32, uint32(HSITRIM)}}
}

func (p *RCC_Periph) MSIRANGE() RMICSCR {
	return RMICSCR{mmio.UM32{&p.ICSCR.U32, uint32(MSIRANGE)}}
}

func (p *RCC_Periph) MSICAL() RMICSCR {
	return RMICSCR{mmio.UM32{&p.ICSCR.U32, uint32(MSICAL)}}
}

func (p *RCC_Periph) MSITRIM() RMICSCR {
	return RMICSCR{mmio.UM32{&p.ICSCR.U32, uint32(MSITRIM)}}
}

type CFGR uint32

func (b CFGR) Field(mask CFGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR) J(v int) CFGR {
	return CFGR(bits.Make32(v, uint32(mask)))
}

type RCFGR struct{ mmio.U32 }

func (r *RCFGR) Bits(mask CFGR) CFGR    { return CFGR(r.U32.Bits(uint32(mask))) }
func (r *RCFGR) StoreBits(mask, b CFGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) SetBits(mask CFGR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR) ClearBits(mask CFGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR) Load() CFGR             { return CFGR(r.U32.Load()) }
func (r *RCFGR) Store(b CFGR)           { r.U32.Store(uint32(b)) }

func (r *RCFGR) AtomicStoreBits(mask, b CFGR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) AtomicSetBits(mask CFGR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR) AtomicClearBits(mask CFGR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR struct{ mmio.UM32 }

func (rm RMCFGR) Load() CFGR   { return CFGR(rm.UM32.Load()) }
func (rm RMCFGR) Store(b CFGR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SW() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(SW)}}
}

func (p *RCC_Periph) SWS() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(SWS)}}
}

func (p *RCC_Periph) HPRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(HPRE)}}
}

func (p *RCC_Periph) PPRE1() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PPRE1)}}
}

func (p *RCC_Periph) PPRE2() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PPRE2)}}
}

func (p *RCC_Periph) PLLSRC() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PLLSRC)}}
}

func (p *RCC_Periph) PLLMUL() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PLLMUL)}}
}

func (p *RCC_Periph) PLLDIV() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PLLDIV)}}
}

func (p *RCC_Periph) MCOSEL() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(MCOSEL)}}
}

func (p *RCC_Periph) MCOPRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(MCOPRE)}}
}

type CIR uint32

func (b CIR) Field(mask CIR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CIR) J(v int) CIR {
	return CIR(bits.Make32(v, uint32(mask)))
}

type RCIR struct{ mmio.U32 }

func (r *RCIR) Bits(mask CIR) CIR     { return CIR(r.U32.Bits(uint32(mask))) }
func (r *RCIR) StoreBits(mask, b CIR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCIR) SetBits(mask CIR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCIR) ClearBits(mask CIR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCIR) Load() CIR             { return CIR(r.U32.Load()) }
func (r *RCIR) Store(b CIR)           { r.U32.Store(uint32(b)) }

func (r *RCIR) AtomicStoreBits(mask, b CIR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCIR) AtomicSetBits(mask CIR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCIR) AtomicClearBits(mask CIR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCIR struct{ mmio.UM32 }

func (rm RMCIR) Load() CIR   { return CIR(rm.UM32.Load()) }
func (rm RMCIR) Store(b CIR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSIRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYF)}}
}

func (p *RCC_Periph) LSERDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYF)}}
}

func (p *RCC_Periph) HSIRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYF)}}
}

func (p *RCC_Periph) HSERDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYF)}}
}

func (p *RCC_Periph) PLLRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYF)}}
}

func (p *RCC_Periph) MSIRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(MSIRDYF)}}
}

func (p *RCC_Periph) LSECSS() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSECSS)}}
}

func (p *RCC_Periph) CSSF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(CSSF)}}
}

func (p *RCC_Periph) LSIRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYIE)}}
}

func (p *RCC_Periph) LSERDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYIE)}}
}

func (p *RCC_Periph) HSIRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYIE)}}
}

func (p *RCC_Periph) HSERDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYIE)}}
}

func (p *RCC_Periph) PLLRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYIE)}}
}

func (p *RCC_Periph) MSIRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(MSIRDYIE)}}
}

func (p *RCC_Periph) LSECSSIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSECSSIE)}}
}

func (p *RCC_Periph) LSIRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYC)}}
}

func (p *RCC_Periph) LSERDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYC)}}
}

func (p *RCC_Periph) HSIRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYC)}}
}

func (p *RCC_Periph) HSERDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYC)}}
}

func (p *RCC_Periph) PLLRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYC)}}
}

func (p *RCC_Periph) MSIRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(MSIRDYC)}}
}

func (p *RCC_Periph) LSECSSC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSECSSC)}}
}

func (p *RCC_Periph) CSSC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(CSSC)}}
}

type AHBRSTR uint32

func (b AHBRSTR) Field(mask AHBRSTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBRSTR) J(v int) AHBRSTR {
	return AHBRSTR(bits.Make32(v, uint32(mask)))
}

type RAHBRSTR struct{ mmio.U32 }

func (r *RAHBRSTR) Bits(mask AHBRSTR) AHBRSTR { return AHBRSTR(r.U32.Bits(uint32(mask))) }
func (r *RAHBRSTR) StoreBits(mask, b AHBRSTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHBRSTR) SetBits(mask AHBRSTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAHBRSTR) ClearBits(mask AHBRSTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAHBRSTR) Load() AHBRSTR             { return AHBRSTR(r.U32.Load()) }
func (r *RAHBRSTR) Store(b AHBRSTR)           { r.U32.Store(uint32(b)) }

func (r *RAHBRSTR) AtomicStoreBits(mask, b AHBRSTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAHBRSTR) AtomicSetBits(mask AHBRSTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAHBRSTR) AtomicClearBits(mask AHBRSTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAHBRSTR struct{ mmio.UM32 }

func (rm RMAHBRSTR) Load() AHBRSTR   { return AHBRSTR(rm.UM32.Load()) }
func (rm RMAHBRSTR) Store(b AHBRSTR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) GPIOARST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOARST)}}
}

func (p *RCC_Periph) GPIOBRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOBRST)}}
}

func (p *RCC_Periph) GPIOCRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOCRST)}}
}

func (p *RCC_Periph) GPIODRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIODRST)}}
}

func (p *RCC_Periph) GPIOERST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOERST)}}
}

func (p *RCC_Periph) GPIOHRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOHRST)}}
}

func (p *RCC_Periph) GPIOFRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOFRST)}}
}

func (p *RCC_Periph) GPIOGRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOGRST)}}
}

func (p *RCC_Periph) CRCRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(CRCRST)}}
}

func (p *RCC_Periph) FLITFRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(FLITFRST)}}
}

func (p *RCC_Periph) DMA1RST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(DMA1RST)}}
}

func (p *RCC_Periph) DMA2RST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(DMA2RST)}}
}

func (p *RCC_Periph) AESRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(AESRST)}}
}

func (p *RCC_Periph) FSMCRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(FSMCRST)}}
}

type APB2RSTR uint32

func (b APB2RSTR) Field(mask APB2RSTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2RSTR) J(v int) APB2RSTR {
	return APB2RSTR(bits.Make32(v, uint32(mask)))
}

type RAPB2RSTR struct{ mmio.U32 }

func (r *RAPB2RSTR) Bits(mask APB2RSTR) APB2RSTR { return APB2RSTR(r.U32.Bits(uint32(mask))) }
func (r *RAPB2RSTR) StoreBits(mask, b APB2RSTR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2RSTR) SetBits(mask APB2RSTR)       { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2RSTR) ClearBits(mask APB2RSTR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2RSTR) Load() APB2RSTR              { return APB2RSTR(r.U32.Load()) }
func (r *RAPB2RSTR) Store(b APB2RSTR)            { r.U32.Store(uint32(b)) }

func (r *RAPB2RSTR) AtomicStoreBits(mask, b APB2RSTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2RSTR) AtomicSetBits(mask APB2RSTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB2RSTR) AtomicClearBits(mask APB2RSTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB2RSTR struct{ mmio.UM32 }

func (rm RMAPB2RSTR) Load() APB2RSTR   { return APB2RSTR(rm.UM32.Load()) }
func (rm RMAPB2RSTR) Store(b APB2RSTR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGRST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SYSCFGRST)}}
}

func (p *RCC_Periph) TIM9RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM9RST)}}
}

func (p *RCC_Periph) TIM10RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM10RST)}}
}

func (p *RCC_Periph) TIM11RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM11RST)}}
}

func (p *RCC_Periph) ADC1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(ADC1RST)}}
}

func (p *RCC_Periph) SDIORST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SDIORST)}}
}

func (p *RCC_Periph) SPI1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SPI1RST)}}
}

func (p *RCC_Periph) USART1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(USART1RST)}}
}

type APB1RSTR uint32

func (b APB1RSTR) Field(mask APB1RSTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1RSTR) J(v int) APB1RSTR {
	return APB1RSTR(bits.Make32(v, uint32(mask)))
}

type RAPB1RSTR struct{ mmio.U32 }

func (r *RAPB1RSTR) Bits(mask APB1RSTR) APB1RSTR { return APB1RSTR(r.U32.Bits(uint32(mask))) }
func (r *RAPB1RSTR) StoreBits(mask, b APB1RSTR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1RSTR) SetBits(mask APB1RSTR)       { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1RSTR) ClearBits(mask APB1RSTR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1RSTR) Load() APB1RSTR              { return APB1RSTR(r.U32.Load()) }
func (r *RAPB1RSTR) Store(b APB1RSTR)            { r.U32.Store(uint32(b)) }

func (r *RAPB1RSTR) AtomicStoreBits(mask, b APB1RSTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1RSTR) AtomicSetBits(mask APB1RSTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB1RSTR) AtomicClearBits(mask APB1RSTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB1RSTR struct{ mmio.UM32 }

func (rm RMAPB1RSTR) Load() APB1RSTR   { return APB1RSTR(rm.UM32.Load()) }
func (rm RMAPB1RSTR) Store(b APB1RSTR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM2RST)}}
}

func (p *RCC_Periph) TIM3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM3RST)}}
}

func (p *RCC_Periph) TIM4RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM4RST)}}
}

func (p *RCC_Periph) TIM5RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM5RST)}}
}

func (p *RCC_Periph) TIM6RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM6RST)}}
}

func (p *RCC_Periph) TIM7RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM7RST)}}
}

func (p *RCC_Periph) LCDRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(LCDRST)}}
}

func (p *RCC_Periph) WWDGRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(WWDGRST)}}
}

func (p *RCC_Periph) SPI2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI2RST)}}
}

func (p *RCC_Periph) SPI3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI3RST)}}
}

func (p *RCC_Periph) USART2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(USART2RST)}}
}

func (p *RCC_Periph) USART3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(USART3RST)}}
}

func (p *RCC_Periph) UART4RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(UART4RST)}}
}

func (p *RCC_Periph) UART5RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(UART5RST)}}
}

func (p *RCC_Periph) I2C1RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C1RST)}}
}

func (p *RCC_Periph) I2C2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C2RST)}}
}

func (p *RCC_Periph) USBRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(USBRST)}}
}

func (p *RCC_Periph) PWRRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(PWRRST)}}
}

func (p *RCC_Periph) DACRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(DACRST)}}
}

func (p *RCC_Periph) COMPRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(COMPRST)}}
}

type AHBENR uint32

func (b AHBENR) Field(mask AHBENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBENR) J(v int) AHBENR {
	return AHBENR(bits.Make32(v, uint32(mask)))
}

type RAHBENR struct{ mmio.U32 }

func (r *RAHBENR) Bits(mask AHBENR) AHBENR  { return AHBENR(r.U32.Bits(uint32(mask))) }
func (r *RAHBENR) StoreBits(mask, b AHBENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHBENR) SetBits(mask AHBENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAHBENR) ClearBits(mask AHBENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAHBENR) Load() AHBENR             { return AHBENR(r.U32.Load()) }
func (r *RAHBENR) Store(b AHBENR)           { r.U32.Store(uint32(b)) }

func (r *RAHBENR) AtomicStoreBits(mask, b AHBENR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAHBENR) AtomicSetBits(mask AHBENR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAHBENR) AtomicClearBits(mask AHBENR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAHBENR struct{ mmio.UM32 }

func (rm RMAHBENR) Load() AHBENR   { return AHBENR(rm.UM32.Load()) }
func (rm RMAHBENR) Store(b AHBENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) GPIOAEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOAEN)}}
}

func (p *RCC_Periph) GPIOBEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOBEN)}}
}

func (p *RCC_Periph) GPIOCEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOCEN)}}
}

func (p *RCC_Periph) GPIODEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIODEN)}}
}

func (p *RCC_Periph) GPIOEEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOEEN)}}
}

func (p *RCC_Periph) GPIOHEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOHEN)}}
}

func (p *RCC_Periph) GPIOFEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOFEN)}}
}

func (p *RCC_Periph) GPIOGEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOGEN)}}
}

func (p *RCC_Periph) CRCEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(CRCEN)}}
}

func (p *RCC_Periph) FLITFEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(FLITFEN)}}
}

func (p *RCC_Periph) DMA1EN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(DMA1EN)}}
}

func (p *RCC_Periph) DMA2EN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(DMA2EN)}}
}

func (p *RCC_Periph) AESEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(AESEN)}}
}

func (p *RCC_Periph) FSMCEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(FSMCEN)}}
}

type APB2ENR uint32

func (b APB2ENR) Field(mask APB2ENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2ENR) J(v int) APB2ENR {
	return APB2ENR(bits.Make32(v, uint32(mask)))
}

type RAPB2ENR struct{ mmio.U32 }

func (r *RAPB2ENR) Bits(mask APB2ENR) APB2ENR { return APB2ENR(r.U32.Bits(uint32(mask))) }
func (r *RAPB2ENR) StoreBits(mask, b APB2ENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2ENR) SetBits(mask APB2ENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2ENR) ClearBits(mask APB2ENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2ENR) Load() APB2ENR             { return APB2ENR(r.U32.Load()) }
func (r *RAPB2ENR) Store(b APB2ENR)           { r.U32.Store(uint32(b)) }

func (r *RAPB2ENR) AtomicStoreBits(mask, b APB2ENR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2ENR) AtomicSetBits(mask APB2ENR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB2ENR) AtomicClearBits(mask APB2ENR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB2ENR struct{ mmio.UM32 }

func (rm RMAPB2ENR) Load() APB2ENR   { return APB2ENR(rm.UM32.Load()) }
func (rm RMAPB2ENR) Store(b APB2ENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGEN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SYSCFGEN)}}
}

func (p *RCC_Periph) TIM9EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM9EN)}}
}

func (p *RCC_Periph) TIM10EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM10EN)}}
}

func (p *RCC_Periph) TIM11EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM11EN)}}
}

func (p *RCC_Periph) ADC1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(ADC1EN)}}
}

func (p *RCC_Periph) SDIOEN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SDIOEN)}}
}

func (p *RCC_Periph) SPI1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SPI1EN)}}
}

func (p *RCC_Periph) USART1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(USART1EN)}}
}

type APB1ENR uint32

func (b APB1ENR) Field(mask APB1ENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1ENR) J(v int) APB1ENR {
	return APB1ENR(bits.Make32(v, uint32(mask)))
}

type RAPB1ENR struct{ mmio.U32 }

func (r *RAPB1ENR) Bits(mask APB1ENR) APB1ENR { return APB1ENR(r.U32.Bits(uint32(mask))) }
func (r *RAPB1ENR) StoreBits(mask, b APB1ENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1ENR) SetBits(mask APB1ENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1ENR) ClearBits(mask APB1ENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1ENR) Load() APB1ENR             { return APB1ENR(r.U32.Load()) }
func (r *RAPB1ENR) Store(b APB1ENR)           { r.U32.Store(uint32(b)) }

func (r *RAPB1ENR) AtomicStoreBits(mask, b APB1ENR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1ENR) AtomicSetBits(mask APB1ENR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB1ENR) AtomicClearBits(mask APB1ENR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB1ENR struct{ mmio.UM32 }

func (rm RMAPB1ENR) Load() APB1ENR   { return APB1ENR(rm.UM32.Load()) }
func (rm RMAPB1ENR) Store(b APB1ENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM2EN)}}
}

func (p *RCC_Periph) TIM3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM3EN)}}
}

func (p *RCC_Periph) TIM4EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM4EN)}}
}

func (p *RCC_Periph) TIM5EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM5EN)}}
}

func (p *RCC_Periph) TIM6EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM6EN)}}
}

func (p *RCC_Periph) TIM7EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM7EN)}}
}

func (p *RCC_Periph) LCDEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(LCDEN)}}
}

func (p *RCC_Periph) WWDGEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(WWDGEN)}}
}

func (p *RCC_Periph) SPI2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(SPI2EN)}}
}

func (p *RCC_Periph) SPI3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(SPI3EN)}}
}

func (p *RCC_Periph) USART2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(USART2EN)}}
}

func (p *RCC_Periph) USART3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(USART3EN)}}
}

func (p *RCC_Periph) UART4EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(UART4EN)}}
}

func (p *RCC_Periph) UART5EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(UART5EN)}}
}

func (p *RCC_Periph) I2C1EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C1EN)}}
}

func (p *RCC_Periph) I2C2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C2EN)}}
}

func (p *RCC_Periph) USBEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(USBEN)}}
}

func (p *RCC_Periph) PWREN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(PWREN)}}
}

func (p *RCC_Periph) DACEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(DACEN)}}
}

func (p *RCC_Periph) COMPEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(COMPEN)}}
}

type AHBLPENR uint32

func (b AHBLPENR) Field(mask AHBLPENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBLPENR) J(v int) AHBLPENR {
	return AHBLPENR(bits.Make32(v, uint32(mask)))
}

type RAHBLPENR struct{ mmio.U32 }

func (r *RAHBLPENR) Bits(mask AHBLPENR) AHBLPENR { return AHBLPENR(r.U32.Bits(uint32(mask))) }
func (r *RAHBLPENR) StoreBits(mask, b AHBLPENR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHBLPENR) SetBits(mask AHBLPENR)       { r.U32.SetBits(uint32(mask)) }
func (r *RAHBLPENR) ClearBits(mask AHBLPENR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RAHBLPENR) Load() AHBLPENR              { return AHBLPENR(r.U32.Load()) }
func (r *RAHBLPENR) Store(b AHBLPENR)            { r.U32.Store(uint32(b)) }

func (r *RAHBLPENR) AtomicStoreBits(mask, b AHBLPENR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAHBLPENR) AtomicSetBits(mask AHBLPENR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAHBLPENR) AtomicClearBits(mask AHBLPENR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAHBLPENR struct{ mmio.UM32 }

func (rm RMAHBLPENR) Load() AHBLPENR   { return AHBLPENR(rm.UM32.Load()) }
func (rm RMAHBLPENR) Store(b AHBLPENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) GPIOALPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOALPEN)}}
}

func (p *RCC_Periph) GPIOBLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOBLPEN)}}
}

func (p *RCC_Periph) GPIOCLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOCLPEN)}}
}

func (p *RCC_Periph) GPIODLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIODLPEN)}}
}

func (p *RCC_Periph) GPIOELPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOELPEN)}}
}

func (p *RCC_Periph) GPIOHLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOHLPEN)}}
}

func (p *RCC_Periph) GPIOFLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOFLPEN)}}
}

func (p *RCC_Periph) GPIOGLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOGLPEN)}}
}

func (p *RCC_Periph) CRCLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(CRCLPEN)}}
}

func (p *RCC_Periph) FLITFLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(FLITFLPEN)}}
}

func (p *RCC_Periph) SRAMLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(SRAMLPEN)}}
}

func (p *RCC_Periph) DMA1LPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(DMA1LPEN)}}
}

func (p *RCC_Periph) DMA2LPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(DMA2LPEN)}}
}

func (p *RCC_Periph) AESLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(AESLPEN)}}
}

func (p *RCC_Periph) FSMCLPEN() RMAHBLPENR {
	return RMAHBLPENR{mmio.UM32{&p.AHBLPENR.U32, uint32(FSMCLPEN)}}
}

type APB2LPENR uint32

func (b APB2LPENR) Field(mask APB2LPENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2LPENR) J(v int) APB2LPENR {
	return APB2LPENR(bits.Make32(v, uint32(mask)))
}

type RAPB2LPENR struct{ mmio.U32 }

func (r *RAPB2LPENR) Bits(mask APB2LPENR) APB2LPENR { return APB2LPENR(r.U32.Bits(uint32(mask))) }
func (r *RAPB2LPENR) StoreBits(mask, b APB2LPENR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2LPENR) SetBits(mask APB2LPENR)        { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2LPENR) ClearBits(mask APB2LPENR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2LPENR) Load() APB2LPENR               { return APB2LPENR(r.U32.Load()) }
func (r *RAPB2LPENR) Store(b APB2LPENR)             { r.U32.Store(uint32(b)) }

func (r *RAPB2LPENR) AtomicStoreBits(mask, b APB2LPENR) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RAPB2LPENR) AtomicSetBits(mask APB2LPENR)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB2LPENR) AtomicClearBits(mask APB2LPENR) { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB2LPENR struct{ mmio.UM32 }

func (rm RMAPB2LPENR) Load() APB2LPENR   { return APB2LPENR(rm.UM32.Load()) }
func (rm RMAPB2LPENR) Store(b APB2LPENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGLPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(SYSCFGLPEN)}}
}

func (p *RCC_Periph) TIM9LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM9LPEN)}}
}

func (p *RCC_Periph) TIM10LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM10LPEN)}}
}

func (p *RCC_Periph) TIM11LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM11LPEN)}}
}

func (p *RCC_Periph) ADC1LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(ADC1LPEN)}}
}

func (p *RCC_Periph) SDIOLPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(SDIOLPEN)}}
}

func (p *RCC_Periph) SPI1LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(SPI1LPEN)}}
}

func (p *RCC_Periph) USART1LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(USART1LPEN)}}
}

type APB1LPENR uint32

func (b APB1LPENR) Field(mask APB1LPENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1LPENR) J(v int) APB1LPENR {
	return APB1LPENR(bits.Make32(v, uint32(mask)))
}

type RAPB1LPENR struct{ mmio.U32 }

func (r *RAPB1LPENR) Bits(mask APB1LPENR) APB1LPENR { return APB1LPENR(r.U32.Bits(uint32(mask))) }
func (r *RAPB1LPENR) StoreBits(mask, b APB1LPENR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1LPENR) SetBits(mask APB1LPENR)        { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1LPENR) ClearBits(mask APB1LPENR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1LPENR) Load() APB1LPENR               { return APB1LPENR(r.U32.Load()) }
func (r *RAPB1LPENR) Store(b APB1LPENR)             { r.U32.Store(uint32(b)) }

func (r *RAPB1LPENR) AtomicStoreBits(mask, b APB1LPENR) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RAPB1LPENR) AtomicSetBits(mask APB1LPENR)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB1LPENR) AtomicClearBits(mask APB1LPENR) { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB1LPENR struct{ mmio.UM32 }

func (rm RMAPB1LPENR) Load() APB1LPENR   { return APB1LPENR(rm.UM32.Load()) }
func (rm RMAPB1LPENR) Store(b APB1LPENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM2LPEN)}}
}

func (p *RCC_Periph) TIM3LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM3LPEN)}}
}

func (p *RCC_Periph) TIM4LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM4LPEN)}}
}

func (p *RCC_Periph) TIM5LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM5LPEN)}}
}

func (p *RCC_Periph) TIM6LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM6LPEN)}}
}

func (p *RCC_Periph) TIM7LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM7LPEN)}}
}

func (p *RCC_Periph) LCDLPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(LCDLPEN)}}
}

func (p *RCC_Periph) WWDGLPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(WWDGLPEN)}}
}

func (p *RCC_Periph) SPI2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(SPI2LPEN)}}
}

func (p *RCC_Periph) SPI3LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(SPI3LPEN)}}
}

func (p *RCC_Periph) USART2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(USART2LPEN)}}
}

func (p *RCC_Periph) USART3LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(USART3LPEN)}}
}

func (p *RCC_Periph) UART4LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(UART4LPEN)}}
}

func (p *RCC_Periph) UART5LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(UART5LPEN)}}
}

func (p *RCC_Periph) I2C1LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(I2C1LPEN)}}
}

func (p *RCC_Periph) I2C2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(I2C2LPEN)}}
}

func (p *RCC_Periph) USBLPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(USBLPEN)}}
}

func (p *RCC_Periph) PWRLPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(PWRLPEN)}}
}

func (p *RCC_Periph) DACLPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(DACLPEN)}}
}

func (p *RCC_Periph) COMPLPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(COMPLPEN)}}
}

type CSR uint32

func (b CSR) Field(mask CSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR) J(v int) CSR {
	return CSR(bits.Make32(v, uint32(mask)))
}

type RCSR struct{ mmio.U32 }

func (r *RCSR) Bits(mask CSR) CSR     { return CSR(r.U32.Bits(uint32(mask))) }
func (r *RCSR) StoreBits(mask, b CSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) SetBits(mask CSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCSR) ClearBits(mask CSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCSR) Load() CSR             { return CSR(r.U32.Load()) }
func (r *RCSR) Store(b CSR)           { r.U32.Store(uint32(b)) }

func (r *RCSR) AtomicStoreBits(mask, b CSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) AtomicSetBits(mask CSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCSR) AtomicClearBits(mask CSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCSR struct{ mmio.UM32 }

func (rm RMCSR) Load() CSR   { return CSR(rm.UM32.Load()) }
func (rm RMCSR) Store(b CSR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSION() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSION)}}
}

func (p *RCC_Periph) LSIRDY() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSIRDY)}}
}

func (p *RCC_Periph) LSEON() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSEON)}}
}

func (p *RCC_Periph) LSERDY() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSERDY)}}
}

func (p *RCC_Periph) LSEBYP() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSEBYP)}}
}

func (p *RCC_Periph) LSECSSON() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSECSSON)}}
}

func (p *RCC_Periph) LSECSSD() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSECSSD)}}
}

func (p *RCC_Periph) RTCSEL() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(RTCSEL)}}
}

func (p *RCC_Periph) RTCEN() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(RTCEN)}}
}

func (p *RCC_Periph) RTCRST() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(RTCRST)}}
}

func (p *RCC_Periph) RMVF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(RMVF)}}
}

func (p *RCC_Periph) OBLRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OBLRSTF)}}
}

func (p *RCC_Periph) PINRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PINRSTF)}}
}

func (p *RCC_Periph) PORRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PORRSTF)}}
}

func (p *RCC_Periph) SFTRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(SFTRSTF)}}
}

func (p *RCC_Periph) IWDGRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(IWDGRSTF)}}
}

func (p *RCC_Periph) WWDGRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(WWDGRSTF)}}
}

func (p *RCC_Periph) LPWRRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LPWRRSTF)}}
}
