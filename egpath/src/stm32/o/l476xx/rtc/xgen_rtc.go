package rtc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type RTC_Periph struct {
	TR      RTR
	DR      RDR
	CR      RCR
	ISR     RISR
	PRER    RPRER
	WUTR    RWUTR
	_       uint32
	ALRMR   [2]RALRMR
	WPR     RWPR
	SSR     RSSR
	SHIFTR  RSHIFTR
	TSTR    RTSTR
	TSDR    RTSDR
	TSSSR   RTSSSR
	CALR    RCALR
	TAMPCR  RTAMPCR
	ALRMSSR [2]RALRMSSR
	OR      ROR
	BKPR    [32]RBKPR
}

func (p *RTC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var RTC = (*RTC_Periph)(unsafe.Pointer(uintptr(mmap.RTC_BASE)))

type TR uint32

func (b TR) Field(mask TR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR) J(v int) TR {
	return TR(bits.Make32(v, uint32(mask)))
}

type RTR struct{ mmio.U32 }

func (r *RTR) Bits(mask TR) TR      { return TR(r.U32.Bits(uint32(mask))) }
func (r *RTR) StoreBits(mask, b TR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTR) SetBits(mask TR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTR) ClearBits(mask TR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTR) Load() TR             { return TR(r.U32.Load()) }
func (r *RTR) Store(b TR)           { r.U32.Store(uint32(b)) }

func (r *RTR) AtomicStoreBits(mask, b TR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTR) AtomicSetBits(mask TR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTR) AtomicClearBits(mask TR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTR struct{ mmio.UM32 }

func (rm RMTR) Load() TR   { return TR(rm.UM32.Load()) }
func (rm RMTR) Store(b TR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) PM() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(PM)}}
}

func (p *RTC_Periph) HT() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(HT)}}
}

func (p *RTC_Periph) HU() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(HU)}}
}

func (p *RTC_Periph) MNT() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(MNT)}}
}

func (p *RTC_Periph) MNU() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(MNU)}}
}

func (p *RTC_Periph) ST() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(ST)}}
}

func (p *RTC_Periph) SU() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(SU)}}
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

func (p *RTC_Periph) YT() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(YT)}}
}

func (p *RTC_Periph) YU() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(YU)}}
}

func (p *RTC_Periph) WDU() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(WDU)}}
}

func (p *RTC_Periph) MT() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(MT)}}
}

func (p *RTC_Periph) MU() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(MU)}}
}

func (p *RTC_Periph) DT() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(DT)}}
}

func (p *RTC_Periph) DU() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(DU)}}
}

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

func (p *RTC_Periph) ITSE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ITSE)}}
}

func (p *RTC_Periph) COE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(COE)}}
}

func (p *RTC_Periph) OSEL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OSEL)}}
}

func (p *RTC_Periph) POL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(POL)}}
}

func (p *RTC_Periph) COSEL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(COSEL)}}
}

func (p *RTC_Periph) BCK() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BCK)}}
}

func (p *RTC_Periph) SUB1H() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SUB1H)}}
}

func (p *RTC_Periph) ADD1H() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADD1H)}}
}

func (p *RTC_Periph) TSIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSIE)}}
}

func (p *RTC_Periph) WUTIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WUTIE)}}
}

func (p *RTC_Periph) ALRBIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRBIE)}}
}

func (p *RTC_Periph) ALRAIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRAIE)}}
}

func (p *RTC_Periph) TSE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSE)}}
}

func (p *RTC_Periph) WUTE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WUTE)}}
}

func (p *RTC_Periph) ALRBE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRBE)}}
}

func (p *RTC_Periph) ALRAE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRAE)}}
}

func (p *RTC_Periph) FMT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FMT)}}
}

func (p *RTC_Periph) BYPSHAD() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BYPSHAD)}}
}

func (p *RTC_Periph) REFCKON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(REFCKON)}}
}

func (p *RTC_Periph) TSEDGE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSEDGE)}}
}

func (p *RTC_Periph) WUCKSEL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WUCKSEL)}}
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

func (p *RTC_Periph) ITSF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ITSF)}}
}

func (p *RTC_Periph) RECALPF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RECALPF)}}
}

func (p *RTC_Periph) TAMP3F() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TAMP3F)}}
}

func (p *RTC_Periph) TAMP2F() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TAMP2F)}}
}

func (p *RTC_Periph) TAMP1F() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TAMP1F)}}
}

func (p *RTC_Periph) TSOVF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TSOVF)}}
}

func (p *RTC_Periph) TSF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TSF)}}
}

func (p *RTC_Periph) WUTF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(WUTF)}}
}

func (p *RTC_Periph) ALRBF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALRBF)}}
}

func (p *RTC_Periph) ALRAF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALRAF)}}
}

func (p *RTC_Periph) INIT() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(INIT)}}
}

func (p *RTC_Periph) INITF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(INITF)}}
}

func (p *RTC_Periph) RSF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RSF)}}
}

func (p *RTC_Periph) INITS() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(INITS)}}
}

func (p *RTC_Periph) SHPF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(SHPF)}}
}

func (p *RTC_Periph) WUTWF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(WUTWF)}}
}

func (p *RTC_Periph) ALRBWF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALRBWF)}}
}

func (p *RTC_Periph) ALRAWF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALRAWF)}}
}

type PRER uint32

func (b PRER) Field(mask PRER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PRER) J(v int) PRER {
	return PRER(bits.Make32(v, uint32(mask)))
}

type RPRER struct{ mmio.U32 }

func (r *RPRER) Bits(mask PRER) PRER    { return PRER(r.U32.Bits(uint32(mask))) }
func (r *RPRER) StoreBits(mask, b PRER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPRER) SetBits(mask PRER)      { r.U32.SetBits(uint32(mask)) }
func (r *RPRER) ClearBits(mask PRER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPRER) Load() PRER             { return PRER(r.U32.Load()) }
func (r *RPRER) Store(b PRER)           { r.U32.Store(uint32(b)) }

func (r *RPRER) AtomicStoreBits(mask, b PRER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPRER) AtomicSetBits(mask PRER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPRER) AtomicClearBits(mask PRER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPRER struct{ mmio.UM32 }

func (rm RMPRER) Load() PRER   { return PRER(rm.UM32.Load()) }
func (rm RMPRER) Store(b PRER) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) PREDIV_A() RMPRER {
	return RMPRER{mmio.UM32{&p.PRER.U32, uint32(PREDIV_A)}}
}

func (p *RTC_Periph) PREDIV_S() RMPRER {
	return RMPRER{mmio.UM32{&p.PRER.U32, uint32(PREDIV_S)}}
}

type WUTR uint32

func (b WUTR) Field(mask WUTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WUTR) J(v int) WUTR {
	return WUTR(bits.Make32(v, uint32(mask)))
}

type RWUTR struct{ mmio.U32 }

func (r *RWUTR) Bits(mask WUTR) WUTR    { return WUTR(r.U32.Bits(uint32(mask))) }
func (r *RWUTR) StoreBits(mask, b WUTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWUTR) SetBits(mask WUTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWUTR) ClearBits(mask WUTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWUTR) Load() WUTR             { return WUTR(r.U32.Load()) }
func (r *RWUTR) Store(b WUTR)           { r.U32.Store(uint32(b)) }

func (r *RWUTR) AtomicStoreBits(mask, b WUTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RWUTR) AtomicSetBits(mask WUTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RWUTR) AtomicClearBits(mask WUTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMWUTR struct{ mmio.UM32 }

func (rm RMWUTR) Load() WUTR   { return WUTR(rm.UM32.Load()) }
func (rm RMWUTR) Store(b WUTR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) WUT() RMWUTR {
	return RMWUTR{mmio.UM32{&p.WUTR.U32, uint32(WUT)}}
}

type ALRMR uint32

func (b ALRMR) Field(mask ALRMR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ALRMR) J(v int) ALRMR {
	return ALRMR(bits.Make32(v, uint32(mask)))
}

type RALRMR struct{ mmio.U32 }

func (r *RALRMR) Bits(mask ALRMR) ALRMR   { return ALRMR(r.U32.Bits(uint32(mask))) }
func (r *RALRMR) StoreBits(mask, b ALRMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RALRMR) SetBits(mask ALRMR)      { r.U32.SetBits(uint32(mask)) }
func (r *RALRMR) ClearBits(mask ALRMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RALRMR) Load() ALRMR             { return ALRMR(r.U32.Load()) }
func (r *RALRMR) Store(b ALRMR)           { r.U32.Store(uint32(b)) }

func (r *RALRMR) AtomicStoreBits(mask, b ALRMR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RALRMR) AtomicSetBits(mask ALRMR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RALRMR) AtomicClearBits(mask ALRMR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMALRMR struct{ mmio.UM32 }

func (rm RMALRMR) Load() ALRMR   { return ALRMR(rm.UM32.Load()) }
func (rm RMALRMR) Store(b ALRMR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) AMSK4(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AMSK4)}}
}

func (p *RTC_Periph) AWDSEL(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AWDSEL)}}
}

func (p *RTC_Periph) ADT(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(ADT)}}
}

func (p *RTC_Periph) ADU(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(ADU)}}
}

func (p *RTC_Periph) AMSK3(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AMSK3)}}
}

func (p *RTC_Periph) APM(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(APM)}}
}

func (p *RTC_Periph) AHT(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AHT)}}
}

func (p *RTC_Periph) AHU(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AHU)}}
}

func (p *RTC_Periph) AMSK2(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AMSK2)}}
}

func (p *RTC_Periph) AMNT(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AMNT)}}
}

func (p *RTC_Periph) AMNU(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AMNU)}}
}

func (p *RTC_Periph) AMSK1(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AMSK1)}}
}

func (p *RTC_Periph) AST(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(AST)}}
}

func (p *RTC_Periph) ASU(n int) RMALRMR {
	return RMALRMR{mmio.UM32{&p.ALRMR[n].U32, uint32(ASU)}}
}

type WPR uint32

func (b WPR) Field(mask WPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WPR) J(v int) WPR {
	return WPR(bits.Make32(v, uint32(mask)))
}

type RWPR struct{ mmio.U32 }

func (r *RWPR) Bits(mask WPR) WPR     { return WPR(r.U32.Bits(uint32(mask))) }
func (r *RWPR) StoreBits(mask, b WPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWPR) SetBits(mask WPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWPR) ClearBits(mask WPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWPR) Load() WPR             { return WPR(r.U32.Load()) }
func (r *RWPR) Store(b WPR)           { r.U32.Store(uint32(b)) }

func (r *RWPR) AtomicStoreBits(mask, b WPR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RWPR) AtomicSetBits(mask WPR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RWPR) AtomicClearBits(mask WPR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMWPR struct{ mmio.UM32 }

func (rm RMWPR) Load() WPR   { return WPR(rm.UM32.Load()) }
func (rm RMWPR) Store(b WPR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) KEY() RMWPR {
	return RMWPR{mmio.UM32{&p.WPR.U32, uint32(KEY)}}
}

type SSR uint32

func (b SSR) Field(mask SSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SSR) J(v int) SSR {
	return SSR(bits.Make32(v, uint32(mask)))
}

type RSSR struct{ mmio.U32 }

func (r *RSSR) Bits(mask SSR) SSR     { return SSR(r.U32.Bits(uint32(mask))) }
func (r *RSSR) StoreBits(mask, b SSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSSR) SetBits(mask SSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSSR) ClearBits(mask SSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSSR) Load() SSR             { return SSR(r.U32.Load()) }
func (r *RSSR) Store(b SSR)           { r.U32.Store(uint32(b)) }

func (r *RSSR) AtomicStoreBits(mask, b SSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSSR) AtomicSetBits(mask SSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSSR) AtomicClearBits(mask SSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSSR struct{ mmio.UM32 }

func (rm RMSSR) Load() SSR   { return SSR(rm.UM32.Load()) }
func (rm RMSSR) Store(b SSR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) SS() RMSSR {
	return RMSSR{mmio.UM32{&p.SSR.U32, uint32(SS)}}
}

type SHIFTR uint32

func (b SHIFTR) Field(mask SHIFTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SHIFTR) J(v int) SHIFTR {
	return SHIFTR(bits.Make32(v, uint32(mask)))
}

type RSHIFTR struct{ mmio.U32 }

func (r *RSHIFTR) Bits(mask SHIFTR) SHIFTR  { return SHIFTR(r.U32.Bits(uint32(mask))) }
func (r *RSHIFTR) StoreBits(mask, b SHIFTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSHIFTR) SetBits(mask SHIFTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSHIFTR) ClearBits(mask SHIFTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSHIFTR) Load() SHIFTR             { return SHIFTR(r.U32.Load()) }
func (r *RSHIFTR) Store(b SHIFTR)           { r.U32.Store(uint32(b)) }

func (r *RSHIFTR) AtomicStoreBits(mask, b SHIFTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSHIFTR) AtomicSetBits(mask SHIFTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSHIFTR) AtomicClearBits(mask SHIFTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSHIFTR struct{ mmio.UM32 }

func (rm RMSHIFTR) Load() SHIFTR   { return SHIFTR(rm.UM32.Load()) }
func (rm RMSHIFTR) Store(b SHIFTR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) SUBFS() RMSHIFTR {
	return RMSHIFTR{mmio.UM32{&p.SHIFTR.U32, uint32(SUBFS)}}
}

func (p *RTC_Periph) ADD1S() RMSHIFTR {
	return RMSHIFTR{mmio.UM32{&p.SHIFTR.U32, uint32(ADD1S)}}
}

type TSTR uint32

func (b TSTR) Field(mask TSTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSTR) J(v int) TSTR {
	return TSTR(bits.Make32(v, uint32(mask)))
}

type RTSTR struct{ mmio.U32 }

func (r *RTSTR) Bits(mask TSTR) TSTR    { return TSTR(r.U32.Bits(uint32(mask))) }
func (r *RTSTR) StoreBits(mask, b TSTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTSTR) SetBits(mask TSTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTSTR) ClearBits(mask TSTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTSTR) Load() TSTR             { return TSTR(r.U32.Load()) }
func (r *RTSTR) Store(b TSTR)           { r.U32.Store(uint32(b)) }

func (r *RTSTR) AtomicStoreBits(mask, b TSTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTSTR) AtomicSetBits(mask TSTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTSTR) AtomicClearBits(mask TSTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTSTR struct{ mmio.UM32 }

func (rm RMTSTR) Load() TSTR   { return TSTR(rm.UM32.Load()) }
func (rm RMTSTR) Store(b TSTR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) TPM() RMTSTR {
	return RMTSTR{mmio.UM32{&p.TSTR.U32, uint32(TPM)}}
}

func (p *RTC_Periph) THT() RMTSTR {
	return RMTSTR{mmio.UM32{&p.TSTR.U32, uint32(THT)}}
}

func (p *RTC_Periph) THU() RMTSTR {
	return RMTSTR{mmio.UM32{&p.TSTR.U32, uint32(THU)}}
}

func (p *RTC_Periph) TMNT() RMTSTR {
	return RMTSTR{mmio.UM32{&p.TSTR.U32, uint32(TMNT)}}
}

func (p *RTC_Periph) TMNU() RMTSTR {
	return RMTSTR{mmio.UM32{&p.TSTR.U32, uint32(TMNU)}}
}

func (p *RTC_Periph) TST() RMTSTR {
	return RMTSTR{mmio.UM32{&p.TSTR.U32, uint32(TST)}}
}

func (p *RTC_Periph) TSU() RMTSTR {
	return RMTSTR{mmio.UM32{&p.TSTR.U32, uint32(TSU)}}
}

type TSDR uint32

func (b TSDR) Field(mask TSDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSDR) J(v int) TSDR {
	return TSDR(bits.Make32(v, uint32(mask)))
}

type RTSDR struct{ mmio.U32 }

func (r *RTSDR) Bits(mask TSDR) TSDR    { return TSDR(r.U32.Bits(uint32(mask))) }
func (r *RTSDR) StoreBits(mask, b TSDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTSDR) SetBits(mask TSDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTSDR) ClearBits(mask TSDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTSDR) Load() TSDR             { return TSDR(r.U32.Load()) }
func (r *RTSDR) Store(b TSDR)           { r.U32.Store(uint32(b)) }

func (r *RTSDR) AtomicStoreBits(mask, b TSDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTSDR) AtomicSetBits(mask TSDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTSDR) AtomicClearBits(mask TSDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTSDR struct{ mmio.UM32 }

func (rm RMTSDR) Load() TSDR   { return TSDR(rm.UM32.Load()) }
func (rm RMTSDR) Store(b TSDR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) TWDU() RMTSDR {
	return RMTSDR{mmio.UM32{&p.TSDR.U32, uint32(TWDU)}}
}

func (p *RTC_Periph) TMT() RMTSDR {
	return RMTSDR{mmio.UM32{&p.TSDR.U32, uint32(TMT)}}
}

func (p *RTC_Periph) TMU() RMTSDR {
	return RMTSDR{mmio.UM32{&p.TSDR.U32, uint32(TMU)}}
}

func (p *RTC_Periph) TDT() RMTSDR {
	return RMTSDR{mmio.UM32{&p.TSDR.U32, uint32(TDT)}}
}

func (p *RTC_Periph) TDU() RMTSDR {
	return RMTSDR{mmio.UM32{&p.TSDR.U32, uint32(TDU)}}
}

type TSSSR uint32

func (b TSSSR) Field(mask TSSSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSSSR) J(v int) TSSSR {
	return TSSSR(bits.Make32(v, uint32(mask)))
}

type RTSSSR struct{ mmio.U32 }

func (r *RTSSSR) Bits(mask TSSSR) TSSSR   { return TSSSR(r.U32.Bits(uint32(mask))) }
func (r *RTSSSR) StoreBits(mask, b TSSSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTSSSR) SetBits(mask TSSSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTSSSR) ClearBits(mask TSSSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTSSSR) Load() TSSSR             { return TSSSR(r.U32.Load()) }
func (r *RTSSSR) Store(b TSSSR)           { r.U32.Store(uint32(b)) }

func (r *RTSSSR) AtomicStoreBits(mask, b TSSSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTSSSR) AtomicSetBits(mask TSSSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTSSSR) AtomicClearBits(mask TSSSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTSSSR struct{ mmio.UM32 }

func (rm RMTSSSR) Load() TSSSR   { return TSSSR(rm.UM32.Load()) }
func (rm RMTSSSR) Store(b TSSSR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) TSS() RMTSSSR {
	return RMTSSSR{mmio.UM32{&p.TSSSR.U32, uint32(TSS)}}
}

type CALR uint32

func (b CALR) Field(mask CALR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CALR) J(v int) CALR {
	return CALR(bits.Make32(v, uint32(mask)))
}

type RCALR struct{ mmio.U32 }

func (r *RCALR) Bits(mask CALR) CALR    { return CALR(r.U32.Bits(uint32(mask))) }
func (r *RCALR) StoreBits(mask, b CALR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCALR) SetBits(mask CALR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCALR) ClearBits(mask CALR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCALR) Load() CALR             { return CALR(r.U32.Load()) }
func (r *RCALR) Store(b CALR)           { r.U32.Store(uint32(b)) }

func (r *RCALR) AtomicStoreBits(mask, b CALR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCALR) AtomicSetBits(mask CALR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCALR) AtomicClearBits(mask CALR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCALR struct{ mmio.UM32 }

func (rm RMCALR) Load() CALR   { return CALR(rm.UM32.Load()) }
func (rm RMCALR) Store(b CALR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) CALP() RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALP)}}
}

func (p *RTC_Periph) CALW8() RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALW8)}}
}

func (p *RTC_Periph) CALW16() RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALW16)}}
}

func (p *RTC_Periph) CALM() RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALM)}}
}

type TAMPCR uint32

func (b TAMPCR) Field(mask TAMPCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TAMPCR) J(v int) TAMPCR {
	return TAMPCR(bits.Make32(v, uint32(mask)))
}

type RTAMPCR struct{ mmio.U32 }

func (r *RTAMPCR) Bits(mask TAMPCR) TAMPCR  { return TAMPCR(r.U32.Bits(uint32(mask))) }
func (r *RTAMPCR) StoreBits(mask, b TAMPCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTAMPCR) SetBits(mask TAMPCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTAMPCR) ClearBits(mask TAMPCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTAMPCR) Load() TAMPCR             { return TAMPCR(r.U32.Load()) }
func (r *RTAMPCR) Store(b TAMPCR)           { r.U32.Store(uint32(b)) }

func (r *RTAMPCR) AtomicStoreBits(mask, b TAMPCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTAMPCR) AtomicSetBits(mask TAMPCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTAMPCR) AtomicClearBits(mask TAMPCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTAMPCR struct{ mmio.UM32 }

func (rm RMTAMPCR) Load() TAMPCR   { return TAMPCR(rm.UM32.Load()) }
func (rm RMTAMPCR) Store(b TAMPCR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) TAMP3MF() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3MF)}}
}

func (p *RTC_Periph) TAMP3NOERASE() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3NOERASE)}}
}

func (p *RTC_Periph) TAMP3IE() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3IE)}}
}

func (p *RTC_Periph) TAMP2MF() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2MF)}}
}

func (p *RTC_Periph) TAMP2NOERASE() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2NOERASE)}}
}

func (p *RTC_Periph) TAMP2IE() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2IE)}}
}

func (p *RTC_Periph) TAMP1MF() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1MF)}}
}

func (p *RTC_Periph) TAMP1NOERASE() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1NOERASE)}}
}

func (p *RTC_Periph) TAMP1IE() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1IE)}}
}

func (p *RTC_Periph) TAMPPUDIS() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPPUDIS)}}
}

func (p *RTC_Periph) TAMPPRCH() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPPRCH)}}
}

func (p *RTC_Periph) TAMPFLT() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPFLT)}}
}

func (p *RTC_Periph) TAMPFREQ() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPFREQ)}}
}

func (p *RTC_Periph) TAMPTS() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPTS)}}
}

func (p *RTC_Periph) TAMP3TRG() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3TRG)}}
}

func (p *RTC_Periph) TAMP3E() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3E)}}
}

func (p *RTC_Periph) TAMP2TRG() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2TRG)}}
}

func (p *RTC_Periph) TAMP2E() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2E)}}
}

func (p *RTC_Periph) TAMPIE() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPIE)}}
}

func (p *RTC_Periph) TAMP1TRG() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1TRG)}}
}

func (p *RTC_Periph) TAMP1E() RMTAMPCR {
	return RMTAMPCR{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1E)}}
}

type ALRMSSR uint32

func (b ALRMSSR) Field(mask ALRMSSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ALRMSSR) J(v int) ALRMSSR {
	return ALRMSSR(bits.Make32(v, uint32(mask)))
}

type RALRMSSR struct{ mmio.U32 }

func (r *RALRMSSR) Bits(mask ALRMSSR) ALRMSSR { return ALRMSSR(r.U32.Bits(uint32(mask))) }
func (r *RALRMSSR) StoreBits(mask, b ALRMSSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RALRMSSR) SetBits(mask ALRMSSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RALRMSSR) ClearBits(mask ALRMSSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RALRMSSR) Load() ALRMSSR             { return ALRMSSR(r.U32.Load()) }
func (r *RALRMSSR) Store(b ALRMSSR)           { r.U32.Store(uint32(b)) }

func (r *RALRMSSR) AtomicStoreBits(mask, b ALRMSSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RALRMSSR) AtomicSetBits(mask ALRMSSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RALRMSSR) AtomicClearBits(mask ALRMSSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMALRMSSR struct{ mmio.UM32 }

func (rm RMALRMSSR) Load() ALRMSSR   { return ALRMSSR(rm.UM32.Load()) }
func (rm RMALRMSSR) Store(b ALRMSSR) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) AMASKSS(n int) RMALRMSSR {
	return RMALRMSSR{mmio.UM32{&p.ALRMSSR[n].U32, uint32(AMASKSS)}}
}

func (p *RTC_Periph) ASS(n int) RMALRMSSR {
	return RMALRMSSR{mmio.UM32{&p.ALRMSSR[n].U32, uint32(ASS)}}
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

func (p *RTC_Periph) OUT_RMP() RMOR {
	return RMOR{mmio.UM32{&p.OR.U32, uint32(OUT_RMP)}}
}

func (p *RTC_Periph) ALARMOUTTYPE() RMOR {
	return RMOR{mmio.UM32{&p.OR.U32, uint32(ALARMOUTTYPE)}}
}

type BKPR uint32

func (b BKPR) Field(mask BKPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BKPR) J(v int) BKPR {
	return BKPR(bits.Make32(v, uint32(mask)))
}

type RBKPR struct{ mmio.U32 }

func (r *RBKPR) Bits(mask BKPR) BKPR    { return BKPR(r.U32.Bits(uint32(mask))) }
func (r *RBKPR) StoreBits(mask, b BKPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBKPR) SetBits(mask BKPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBKPR) ClearBits(mask BKPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBKPR) Load() BKPR             { return BKPR(r.U32.Load()) }
func (r *RBKPR) Store(b BKPR)           { r.U32.Store(uint32(b)) }

func (r *RBKPR) AtomicStoreBits(mask, b BKPR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBKPR) AtomicSetBits(mask BKPR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBKPR) AtomicClearBits(mask BKPR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBKPR struct{ mmio.UM32 }

func (rm RMBKPR) Load() BKPR   { return BKPR(rm.UM32.Load()) }
func (rm RMBKPR) Store(b BKPR) { rm.UM32.Store(uint32(b)) }
