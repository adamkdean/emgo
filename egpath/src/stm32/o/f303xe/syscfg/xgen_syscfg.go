package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type SYSCFG_Periph struct {
	CFGR1      RCFGR1
	RCR        RRCR
	EXTICR     [4]REXTICR
	CFGR2      RCFGR2
	RESERVED0  RRESERVED0
	RESERVED1  RRESERVED1
	RESERVED2  RRESERVED2
	RESERVED4  RRESERVED4
	RESERVED5  RRESERVED5
	RESERVED6  RRESERVED6
	RESERVED7  RRESERVED7
	RESERVED8  RRESERVED8
	RESERVED9  RRESERVED9
	RESERVED10 RRESERVED10
	RESERVED11 RRESERVED11
	CFGR4      RCFGR4
	RESERVED12 RRESERVED12
	RESERVED13 RRESERVED13
}

func (p *SYSCFG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SYSCFG = (*SYSCFG_Periph)(unsafe.Pointer(uintptr(mmap.SYSCFG_BASE)))

type CFGR1 uint32

func (b CFGR1) Field(mask CFGR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR1) J(v int) CFGR1 {
	return CFGR1(bits.Make32(v, uint32(mask)))
}

type RCFGR1 struct{ mmio.U32 }

func (r *RCFGR1) Bits(mask CFGR1) CFGR1   { return CFGR1(r.U32.Bits(uint32(mask))) }
func (r *RCFGR1) StoreBits(mask, b CFGR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR1) SetBits(mask CFGR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR1) ClearBits(mask CFGR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR1) Load() CFGR1             { return CFGR1(r.U32.Load()) }
func (r *RCFGR1) Store(b CFGR1)           { r.U32.Store(uint32(b)) }

func (r *RCFGR1) AtomicStoreBits(mask, b CFGR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR1) AtomicSetBits(mask CFGR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR1) AtomicClearBits(mask CFGR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR1 struct{ mmio.UM32 }

func (rm RMCFGR1) Load() CFGR1   { return CFGR1(rm.UM32.Load()) }
func (rm RMCFGR1) Store(b CFGR1) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) MEM_MODE() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(MEM_MODE)}}
}

func (p *SYSCFG_Periph) USB_IT_RMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(USB_IT_RMP)}}
}

func (p *SYSCFG_Periph) TIM1_ITR3_RMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(TIM1_ITR3_RMP)}}
}

func (p *SYSCFG_Periph) DAC1_TRIG1_RMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(DAC1_TRIG1_RMP)}}
}

func (p *SYSCFG_Periph) DMA_RMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(DMA_RMP)}}
}

func (p *SYSCFG_Periph) I2C_PB6_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_PB6_FMP)}}
}

func (p *SYSCFG_Periph) I2C_PB7_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_PB7_FMP)}}
}

func (p *SYSCFG_Periph) I2C_PB8_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_PB8_FMP)}}
}

func (p *SYSCFG_Periph) I2C_PB9_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_PB9_FMP)}}
}

func (p *SYSCFG_Periph) I2C1_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C1_FMP)}}
}

func (p *SYSCFG_Periph) I2C2_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C2_FMP)}}
}

func (p *SYSCFG_Periph) ENCODER_MODE() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(ENCODER_MODE)}}
}

func (p *SYSCFG_Periph) I2C3_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C3_FMP)}}
}

func (p *SYSCFG_Periph) FPU_IE() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(FPU_IE)}}
}

type RCR uint32

func (b RCR) Field(mask RCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RCR) J(v int) RCR {
	return RCR(bits.Make32(v, uint32(mask)))
}

type RRCR struct{ mmio.U32 }

func (r *RRCR) Bits(mask RCR) RCR     { return RCR(r.U32.Bits(uint32(mask))) }
func (r *RRCR) StoreBits(mask, b RCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRCR) SetBits(mask RCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRCR) ClearBits(mask RCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRCR) Load() RCR             { return RCR(r.U32.Load()) }
func (r *RRCR) Store(b RCR)           { r.U32.Store(uint32(b)) }

func (r *RRCR) AtomicStoreBits(mask, b RCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRCR) AtomicSetBits(mask RCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRCR) AtomicClearBits(mask RCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRCR struct{ mmio.UM32 }

func (rm RMRCR) Load() RCR   { return RCR(rm.UM32.Load()) }
func (rm RMRCR) Store(b RCR) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) PAGE0() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE0)}}
}

func (p *SYSCFG_Periph) PAGE1() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE1)}}
}

func (p *SYSCFG_Periph) PAGE2() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE2)}}
}

func (p *SYSCFG_Periph) PAGE3() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE3)}}
}

func (p *SYSCFG_Periph) PAGE4() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE4)}}
}

func (p *SYSCFG_Periph) PAGE5() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE5)}}
}

func (p *SYSCFG_Periph) PAGE6() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE6)}}
}

func (p *SYSCFG_Periph) PAGE7() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE7)}}
}

func (p *SYSCFG_Periph) PAGE8() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE8)}}
}

func (p *SYSCFG_Periph) PAGE9() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE9)}}
}

func (p *SYSCFG_Periph) PAGE10() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE10)}}
}

func (p *SYSCFG_Periph) PAGE11() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE11)}}
}

func (p *SYSCFG_Periph) PAGE12() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE12)}}
}

func (p *SYSCFG_Periph) PAGE13() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE13)}}
}

func (p *SYSCFG_Periph) PAGE14() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE14)}}
}

func (p *SYSCFG_Periph) PAGE15() RMRCR {
	return RMRCR{mmio.UM32{&p.RCR.U32, uint32(PAGE15)}}
}

type EXTICR uint32

func (b EXTICR) Field(mask EXTICR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EXTICR) J(v int) EXTICR {
	return EXTICR(bits.Make32(v, uint32(mask)))
}

type REXTICR struct{ mmio.U32 }

func (r *REXTICR) Bits(mask EXTICR) EXTICR  { return EXTICR(r.U32.Bits(uint32(mask))) }
func (r *REXTICR) StoreBits(mask, b EXTICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REXTICR) SetBits(mask EXTICR)      { r.U32.SetBits(uint32(mask)) }
func (r *REXTICR) ClearBits(mask EXTICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *REXTICR) Load() EXTICR             { return EXTICR(r.U32.Load()) }
func (r *REXTICR) Store(b EXTICR)           { r.U32.Store(uint32(b)) }

func (r *REXTICR) AtomicStoreBits(mask, b EXTICR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *REXTICR) AtomicSetBits(mask EXTICR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *REXTICR) AtomicClearBits(mask EXTICR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMEXTICR struct{ mmio.UM32 }

func (rm RMEXTICR) Load() EXTICR   { return EXTICR(rm.UM32.Load()) }
func (rm RMEXTICR) Store(b EXTICR) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) EXTI0(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI0)}}
}

func (p *SYSCFG_Periph) EXTI1(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI1)}}
}

func (p *SYSCFG_Periph) EXTI2(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI2)}}
}

func (p *SYSCFG_Periph) EXTI3(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI3)}}
}

type CFGR2 uint32

func (b CFGR2) Field(mask CFGR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR2) J(v int) CFGR2 {
	return CFGR2(bits.Make32(v, uint32(mask)))
}

type RCFGR2 struct{ mmio.U32 }

func (r *RCFGR2) Bits(mask CFGR2) CFGR2   { return CFGR2(r.U32.Bits(uint32(mask))) }
func (r *RCFGR2) StoreBits(mask, b CFGR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR2) SetBits(mask CFGR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR2) ClearBits(mask CFGR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR2) Load() CFGR2             { return CFGR2(r.U32.Load()) }
func (r *RCFGR2) Store(b CFGR2)           { r.U32.Store(uint32(b)) }

func (r *RCFGR2) AtomicStoreBits(mask, b CFGR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR2) AtomicSetBits(mask CFGR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR2) AtomicClearBits(mask CFGR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR2 struct{ mmio.UM32 }

func (rm RMCFGR2) Load() CFGR2   { return CFGR2(rm.UM32.Load()) }
func (rm RMCFGR2) Store(b CFGR2) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) LOCKUP_LOCK() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(LOCKUP_LOCK)}}
}

func (p *SYSCFG_Periph) SRAM_PARITY_LOCK() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(SRAM_PARITY_LOCK)}}
}

func (p *SYSCFG_Periph) PVD_LOCK() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(PVD_LOCK)}}
}

func (p *SYSCFG_Periph) BYP_ADDR_PAR() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(BYP_ADDR_PAR)}}
}

func (p *SYSCFG_Periph) SRAM_PE() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(SRAM_PE)}}
}

type RESERVED0 uint32

func (b RESERVED0) Field(mask RESERVED0) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED0) J(v int) RESERVED0 {
	return RESERVED0(bits.Make32(v, uint32(mask)))
}

type RRESERVED0 struct{ mmio.U32 }

func (r *RRESERVED0) Bits(mask RESERVED0) RESERVED0 { return RESERVED0(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED0) StoreBits(mask, b RESERVED0)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED0) SetBits(mask RESERVED0)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED0) ClearBits(mask RESERVED0)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED0) Load() RESERVED0               { return RESERVED0(r.U32.Load()) }
func (r *RRESERVED0) Store(b RESERVED0)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED0) AtomicStoreBits(mask, b RESERVED0) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED0) AtomicSetBits(mask RESERVED0)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED0) AtomicClearBits(mask RESERVED0) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED0 struct{ mmio.UM32 }

func (rm RMRESERVED0) Load() RESERVED0   { return RESERVED0(rm.UM32.Load()) }
func (rm RMRESERVED0) Store(b RESERVED0) { rm.UM32.Store(uint32(b)) }

type RESERVED1 uint32

func (b RESERVED1) Field(mask RESERVED1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED1) J(v int) RESERVED1 {
	return RESERVED1(bits.Make32(v, uint32(mask)))
}

type RRESERVED1 struct{ mmio.U32 }

func (r *RRESERVED1) Bits(mask RESERVED1) RESERVED1 { return RESERVED1(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED1) StoreBits(mask, b RESERVED1)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED1) SetBits(mask RESERVED1)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED1) ClearBits(mask RESERVED1)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED1) Load() RESERVED1               { return RESERVED1(r.U32.Load()) }
func (r *RRESERVED1) Store(b RESERVED1)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED1) AtomicStoreBits(mask, b RESERVED1) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED1) AtomicSetBits(mask RESERVED1)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED1) AtomicClearBits(mask RESERVED1) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED1 struct{ mmio.UM32 }

func (rm RMRESERVED1) Load() RESERVED1   { return RESERVED1(rm.UM32.Load()) }
func (rm RMRESERVED1) Store(b RESERVED1) { rm.UM32.Store(uint32(b)) }

type RESERVED2 uint32

func (b RESERVED2) Field(mask RESERVED2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED2) J(v int) RESERVED2 {
	return RESERVED2(bits.Make32(v, uint32(mask)))
}

type RRESERVED2 struct{ mmio.U32 }

func (r *RRESERVED2) Bits(mask RESERVED2) RESERVED2 { return RESERVED2(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED2) StoreBits(mask, b RESERVED2)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED2) SetBits(mask RESERVED2)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED2) ClearBits(mask RESERVED2)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED2) Load() RESERVED2               { return RESERVED2(r.U32.Load()) }
func (r *RRESERVED2) Store(b RESERVED2)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED2) AtomicStoreBits(mask, b RESERVED2) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED2) AtomicSetBits(mask RESERVED2)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED2) AtomicClearBits(mask RESERVED2) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED2 struct{ mmio.UM32 }

func (rm RMRESERVED2) Load() RESERVED2   { return RESERVED2(rm.UM32.Load()) }
func (rm RMRESERVED2) Store(b RESERVED2) { rm.UM32.Store(uint32(b)) }

type RESERVED4 uint32

func (b RESERVED4) Field(mask RESERVED4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED4) J(v int) RESERVED4 {
	return RESERVED4(bits.Make32(v, uint32(mask)))
}

type RRESERVED4 struct{ mmio.U32 }

func (r *RRESERVED4) Bits(mask RESERVED4) RESERVED4 { return RESERVED4(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED4) StoreBits(mask, b RESERVED4)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED4) SetBits(mask RESERVED4)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED4) ClearBits(mask RESERVED4)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED4) Load() RESERVED4               { return RESERVED4(r.U32.Load()) }
func (r *RRESERVED4) Store(b RESERVED4)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED4) AtomicStoreBits(mask, b RESERVED4) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED4) AtomicSetBits(mask RESERVED4)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED4) AtomicClearBits(mask RESERVED4) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED4 struct{ mmio.UM32 }

func (rm RMRESERVED4) Load() RESERVED4   { return RESERVED4(rm.UM32.Load()) }
func (rm RMRESERVED4) Store(b RESERVED4) { rm.UM32.Store(uint32(b)) }

type RESERVED5 uint32

func (b RESERVED5) Field(mask RESERVED5) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED5) J(v int) RESERVED5 {
	return RESERVED5(bits.Make32(v, uint32(mask)))
}

type RRESERVED5 struct{ mmio.U32 }

func (r *RRESERVED5) Bits(mask RESERVED5) RESERVED5 { return RESERVED5(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED5) StoreBits(mask, b RESERVED5)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED5) SetBits(mask RESERVED5)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED5) ClearBits(mask RESERVED5)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED5) Load() RESERVED5               { return RESERVED5(r.U32.Load()) }
func (r *RRESERVED5) Store(b RESERVED5)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED5) AtomicStoreBits(mask, b RESERVED5) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED5) AtomicSetBits(mask RESERVED5)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED5) AtomicClearBits(mask RESERVED5) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED5 struct{ mmio.UM32 }

func (rm RMRESERVED5) Load() RESERVED5   { return RESERVED5(rm.UM32.Load()) }
func (rm RMRESERVED5) Store(b RESERVED5) { rm.UM32.Store(uint32(b)) }

type RESERVED6 uint32

func (b RESERVED6) Field(mask RESERVED6) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED6) J(v int) RESERVED6 {
	return RESERVED6(bits.Make32(v, uint32(mask)))
}

type RRESERVED6 struct{ mmio.U32 }

func (r *RRESERVED6) Bits(mask RESERVED6) RESERVED6 { return RESERVED6(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED6) StoreBits(mask, b RESERVED6)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED6) SetBits(mask RESERVED6)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED6) ClearBits(mask RESERVED6)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED6) Load() RESERVED6               { return RESERVED6(r.U32.Load()) }
func (r *RRESERVED6) Store(b RESERVED6)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED6) AtomicStoreBits(mask, b RESERVED6) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED6) AtomicSetBits(mask RESERVED6)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED6) AtomicClearBits(mask RESERVED6) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED6 struct{ mmio.UM32 }

func (rm RMRESERVED6) Load() RESERVED6   { return RESERVED6(rm.UM32.Load()) }
func (rm RMRESERVED6) Store(b RESERVED6) { rm.UM32.Store(uint32(b)) }

type RESERVED7 uint32

func (b RESERVED7) Field(mask RESERVED7) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED7) J(v int) RESERVED7 {
	return RESERVED7(bits.Make32(v, uint32(mask)))
}

type RRESERVED7 struct{ mmio.U32 }

func (r *RRESERVED7) Bits(mask RESERVED7) RESERVED7 { return RESERVED7(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED7) StoreBits(mask, b RESERVED7)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED7) SetBits(mask RESERVED7)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED7) ClearBits(mask RESERVED7)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED7) Load() RESERVED7               { return RESERVED7(r.U32.Load()) }
func (r *RRESERVED7) Store(b RESERVED7)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED7) AtomicStoreBits(mask, b RESERVED7) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED7) AtomicSetBits(mask RESERVED7)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED7) AtomicClearBits(mask RESERVED7) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED7 struct{ mmio.UM32 }

func (rm RMRESERVED7) Load() RESERVED7   { return RESERVED7(rm.UM32.Load()) }
func (rm RMRESERVED7) Store(b RESERVED7) { rm.UM32.Store(uint32(b)) }

type RESERVED8 uint32

func (b RESERVED8) Field(mask RESERVED8) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED8) J(v int) RESERVED8 {
	return RESERVED8(bits.Make32(v, uint32(mask)))
}

type RRESERVED8 struct{ mmio.U32 }

func (r *RRESERVED8) Bits(mask RESERVED8) RESERVED8 { return RESERVED8(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED8) StoreBits(mask, b RESERVED8)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED8) SetBits(mask RESERVED8)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED8) ClearBits(mask RESERVED8)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED8) Load() RESERVED8               { return RESERVED8(r.U32.Load()) }
func (r *RRESERVED8) Store(b RESERVED8)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED8) AtomicStoreBits(mask, b RESERVED8) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED8) AtomicSetBits(mask RESERVED8)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED8) AtomicClearBits(mask RESERVED8) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED8 struct{ mmio.UM32 }

func (rm RMRESERVED8) Load() RESERVED8   { return RESERVED8(rm.UM32.Load()) }
func (rm RMRESERVED8) Store(b RESERVED8) { rm.UM32.Store(uint32(b)) }

type RESERVED9 uint32

func (b RESERVED9) Field(mask RESERVED9) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED9) J(v int) RESERVED9 {
	return RESERVED9(bits.Make32(v, uint32(mask)))
}

type RRESERVED9 struct{ mmio.U32 }

func (r *RRESERVED9) Bits(mask RESERVED9) RESERVED9 { return RESERVED9(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED9) StoreBits(mask, b RESERVED9)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED9) SetBits(mask RESERVED9)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED9) ClearBits(mask RESERVED9)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED9) Load() RESERVED9               { return RESERVED9(r.U32.Load()) }
func (r *RRESERVED9) Store(b RESERVED9)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED9) AtomicStoreBits(mask, b RESERVED9) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED9) AtomicSetBits(mask RESERVED9)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED9) AtomicClearBits(mask RESERVED9) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED9 struct{ mmio.UM32 }

func (rm RMRESERVED9) Load() RESERVED9   { return RESERVED9(rm.UM32.Load()) }
func (rm RMRESERVED9) Store(b RESERVED9) { rm.UM32.Store(uint32(b)) }

type RESERVED10 uint32

func (b RESERVED10) Field(mask RESERVED10) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED10) J(v int) RESERVED10 {
	return RESERVED10(bits.Make32(v, uint32(mask)))
}

type RRESERVED10 struct{ mmio.U32 }

func (r *RRESERVED10) Bits(mask RESERVED10) RESERVED10 { return RESERVED10(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED10) StoreBits(mask, b RESERVED10)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED10) SetBits(mask RESERVED10)         { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED10) ClearBits(mask RESERVED10)       { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED10) Load() RESERVED10                { return RESERVED10(r.U32.Load()) }
func (r *RRESERVED10) Store(b RESERVED10)              { r.U32.Store(uint32(b)) }

func (r *RRESERVED10) AtomicStoreBits(mask, b RESERVED10) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED10) AtomicSetBits(mask RESERVED10)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED10) AtomicClearBits(mask RESERVED10) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED10 struct{ mmio.UM32 }

func (rm RMRESERVED10) Load() RESERVED10   { return RESERVED10(rm.UM32.Load()) }
func (rm RMRESERVED10) Store(b RESERVED10) { rm.UM32.Store(uint32(b)) }

type RESERVED11 uint32

func (b RESERVED11) Field(mask RESERVED11) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED11) J(v int) RESERVED11 {
	return RESERVED11(bits.Make32(v, uint32(mask)))
}

type RRESERVED11 struct{ mmio.U32 }

func (r *RRESERVED11) Bits(mask RESERVED11) RESERVED11 { return RESERVED11(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED11) StoreBits(mask, b RESERVED11)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED11) SetBits(mask RESERVED11)         { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED11) ClearBits(mask RESERVED11)       { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED11) Load() RESERVED11                { return RESERVED11(r.U32.Load()) }
func (r *RRESERVED11) Store(b RESERVED11)              { r.U32.Store(uint32(b)) }

func (r *RRESERVED11) AtomicStoreBits(mask, b RESERVED11) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED11) AtomicSetBits(mask RESERVED11)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED11) AtomicClearBits(mask RESERVED11) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED11 struct{ mmio.UM32 }

func (rm RMRESERVED11) Load() RESERVED11   { return RESERVED11(rm.UM32.Load()) }
func (rm RMRESERVED11) Store(b RESERVED11) { rm.UM32.Store(uint32(b)) }

type CFGR4 uint32

func (b CFGR4) Field(mask CFGR4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR4) J(v int) CFGR4 {
	return CFGR4(bits.Make32(v, uint32(mask)))
}

type RCFGR4 struct{ mmio.U32 }

func (r *RCFGR4) Bits(mask CFGR4) CFGR4   { return CFGR4(r.U32.Bits(uint32(mask))) }
func (r *RCFGR4) StoreBits(mask, b CFGR4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR4) SetBits(mask CFGR4)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR4) ClearBits(mask CFGR4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR4) Load() CFGR4             { return CFGR4(r.U32.Load()) }
func (r *RCFGR4) Store(b CFGR4)           { r.U32.Store(uint32(b)) }

func (r *RCFGR4) AtomicStoreBits(mask, b CFGR4) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR4) AtomicSetBits(mask CFGR4)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR4) AtomicClearBits(mask CFGR4)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR4 struct{ mmio.UM32 }

func (rm RMCFGR4) Load() CFGR4   { return CFGR4(rm.UM32.Load()) }
func (rm RMCFGR4) Store(b CFGR4) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) ADC12_EXT2_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC12_EXT2_RMP)}}
}

func (p *SYSCFG_Periph) ADC12_EXT3_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC12_EXT3_RMP)}}
}

func (p *SYSCFG_Periph) ADC12_EXT5_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC12_EXT5_RMP)}}
}

func (p *SYSCFG_Periph) ADC12_EXT13_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC12_EXT13_RMP)}}
}

func (p *SYSCFG_Periph) ADC12_EXT15_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC12_EXT15_RMP)}}
}

func (p *SYSCFG_Periph) ADC12_JEXT3_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC12_JEXT3_RMP)}}
}

func (p *SYSCFG_Periph) ADC12_JEXT6_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC12_JEXT6_RMP)}}
}

func (p *SYSCFG_Periph) ADC12_JEXT13_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC12_JEXT13_RMP)}}
}

func (p *SYSCFG_Periph) ADC34_EXT5_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC34_EXT5_RMP)}}
}

func (p *SYSCFG_Periph) ADC34_EXT6_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC34_EXT6_RMP)}}
}

func (p *SYSCFG_Periph) ADC34_EXT15_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC34_EXT15_RMP)}}
}

func (p *SYSCFG_Periph) ADC34_JEXT5_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC34_JEXT5_RMP)}}
}

func (p *SYSCFG_Periph) ADC34_JEXT11_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC34_JEXT11_RMP)}}
}

func (p *SYSCFG_Periph) ADC34_JEXT14_RMP() RMCFGR4 {
	return RMCFGR4{mmio.UM32{&p.CFGR4.U32, uint32(ADC34_JEXT14_RMP)}}
}

type RESERVED12 uint32

func (b RESERVED12) Field(mask RESERVED12) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED12) J(v int) RESERVED12 {
	return RESERVED12(bits.Make32(v, uint32(mask)))
}

type RRESERVED12 struct{ mmio.U32 }

func (r *RRESERVED12) Bits(mask RESERVED12) RESERVED12 { return RESERVED12(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED12) StoreBits(mask, b RESERVED12)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED12) SetBits(mask RESERVED12)         { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED12) ClearBits(mask RESERVED12)       { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED12) Load() RESERVED12                { return RESERVED12(r.U32.Load()) }
func (r *RRESERVED12) Store(b RESERVED12)              { r.U32.Store(uint32(b)) }

func (r *RRESERVED12) AtomicStoreBits(mask, b RESERVED12) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED12) AtomicSetBits(mask RESERVED12)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED12) AtomicClearBits(mask RESERVED12) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED12 struct{ mmio.UM32 }

func (rm RMRESERVED12) Load() RESERVED12   { return RESERVED12(rm.UM32.Load()) }
func (rm RMRESERVED12) Store(b RESERVED12) { rm.UM32.Store(uint32(b)) }

type RESERVED13 uint32

func (b RESERVED13) Field(mask RESERVED13) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED13) J(v int) RESERVED13 {
	return RESERVED13(bits.Make32(v, uint32(mask)))
}

type RRESERVED13 struct{ mmio.U32 }

func (r *RRESERVED13) Bits(mask RESERVED13) RESERVED13 { return RESERVED13(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED13) StoreBits(mask, b RESERVED13)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED13) SetBits(mask RESERVED13)         { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED13) ClearBits(mask RESERVED13)       { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED13) Load() RESERVED13                { return RESERVED13(r.U32.Load()) }
func (r *RRESERVED13) Store(b RESERVED13)              { r.U32.Store(uint32(b)) }

func (r *RRESERVED13) AtomicStoreBits(mask, b RESERVED13) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED13) AtomicSetBits(mask RESERVED13)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED13) AtomicClearBits(mask RESERVED13) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED13 struct{ mmio.UM32 }

func (rm RMRESERVED13) Load() RESERVED13   { return RESERVED13(rm.UM32.Load()) }
func (rm RMRESERVED13) Store(b RESERVED13) { rm.UM32.Store(uint32(b)) }
