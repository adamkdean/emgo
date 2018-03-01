// +build f746xx

package can

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type CAN_Periph struct {
	MCR   RMCR
	MSR   RMSR
	TSR   RTSR
	RF0R  RRF0R
	RF1R  RRF1R
	IER   RIER
	ESR   RESR
	BTR   RBTR
	_     [100]uint32
	FMR   RFMR
	FM1R  RFM1R
	_     uint32
	FS1R  RFS1R
	_     uint32
	FFA1R RFFA1R
	_     uint32
	FA1R  RFA1R
}

func (p *CAN_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var CAN1 = (*CAN_Periph)(unsafe.Pointer(uintptr(mmap.CAN1_BASE)))

//emgo:const
var CAN2 = (*CAN_Periph)(unsafe.Pointer(uintptr(mmap.CAN2_BASE)))

type MCR uint32

func (b MCR) Field(mask MCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MCR) J(v int) MCR {
	return MCR(bits.Make32(v, uint32(mask)))
}

type RMCR struct{ mmio.U32 }

func (r *RMCR) Bits(mask MCR) MCR     { return MCR(r.U32.Bits(uint32(mask))) }
func (r *RMCR) StoreBits(mask, b MCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMCR) SetBits(mask MCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RMCR) ClearBits(mask MCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMCR) Load() MCR             { return MCR(r.U32.Load()) }
func (r *RMCR) Store(b MCR)           { r.U32.Store(uint32(b)) }

func (r *RMCR) AtomicStoreBits(mask, b MCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMCR) AtomicSetBits(mask MCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMCR) AtomicClearBits(mask MCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMCR struct{ mmio.UM32 }

func (rm RMMCR) Load() MCR   { return MCR(rm.UM32.Load()) }
func (rm RMMCR) Store(b MCR) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) INRQ() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(INRQ)}}
}

func (p *CAN_Periph) SLEEP() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(SLEEP)}}
}

func (p *CAN_Periph) TXFP() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(TXFP)}}
}

func (p *CAN_Periph) RFLM() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(RFLM)}}
}

func (p *CAN_Periph) NART() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(NART)}}
}

func (p *CAN_Periph) AWUM() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(AWUM)}}
}

func (p *CAN_Periph) ABOM() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(ABOM)}}
}

func (p *CAN_Periph) TTCM() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(TTCM)}}
}

func (p *CAN_Periph) RESET() RMMCR {
	return RMMCR{mmio.UM32{&p.MCR.U32, uint32(RESET)}}
}

type MSR uint32

func (b MSR) Field(mask MSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MSR) J(v int) MSR {
	return MSR(bits.Make32(v, uint32(mask)))
}

type RMSR struct{ mmio.U32 }

func (r *RMSR) Bits(mask MSR) MSR     { return MSR(r.U32.Bits(uint32(mask))) }
func (r *RMSR) StoreBits(mask, b MSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMSR) SetBits(mask MSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RMSR) ClearBits(mask MSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMSR) Load() MSR             { return MSR(r.U32.Load()) }
func (r *RMSR) Store(b MSR)           { r.U32.Store(uint32(b)) }

func (r *RMSR) AtomicStoreBits(mask, b MSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMSR) AtomicSetBits(mask MSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMSR) AtomicClearBits(mask MSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMSR struct{ mmio.UM32 }

func (rm RMMSR) Load() MSR   { return MSR(rm.UM32.Load()) }
func (rm RMMSR) Store(b MSR) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) INAK() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(INAK)}}
}

func (p *CAN_Periph) SLAK() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(SLAK)}}
}

func (p *CAN_Periph) ERRI() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(ERRI)}}
}

func (p *CAN_Periph) WKUI() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(WKUI)}}
}

func (p *CAN_Periph) SLAKI() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(SLAKI)}}
}

func (p *CAN_Periph) TXM() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(TXM)}}
}

func (p *CAN_Periph) RXM() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(RXM)}}
}

func (p *CAN_Periph) SAMP() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(SAMP)}}
}

func (p *CAN_Periph) RX() RMMSR {
	return RMMSR{mmio.UM32{&p.MSR.U32, uint32(RX)}}
}

type TSR uint32

func (b TSR) Field(mask TSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSR) J(v int) TSR {
	return TSR(bits.Make32(v, uint32(mask)))
}

type RTSR struct{ mmio.U32 }

func (r *RTSR) Bits(mask TSR) TSR     { return TSR(r.U32.Bits(uint32(mask))) }
func (r *RTSR) StoreBits(mask, b TSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTSR) SetBits(mask TSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTSR) ClearBits(mask TSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTSR) Load() TSR             { return TSR(r.U32.Load()) }
func (r *RTSR) Store(b TSR)           { r.U32.Store(uint32(b)) }

func (r *RTSR) AtomicStoreBits(mask, b TSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTSR) AtomicSetBits(mask TSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTSR) AtomicClearBits(mask TSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTSR struct{ mmio.UM32 }

func (rm RMTSR) Load() TSR   { return TSR(rm.UM32.Load()) }
func (rm RMTSR) Store(b TSR) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) RQCP0() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(RQCP0)}}
}

func (p *CAN_Periph) TXOK0() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(TXOK0)}}
}

func (p *CAN_Periph) ALST0() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(ALST0)}}
}

func (p *CAN_Periph) TERR0() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(TERR0)}}
}

func (p *CAN_Periph) ABRQ0() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(ABRQ0)}}
}

func (p *CAN_Periph) RQCP1() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(RQCP1)}}
}

func (p *CAN_Periph) TXOK1() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(TXOK1)}}
}

func (p *CAN_Periph) ALST1() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(ALST1)}}
}

func (p *CAN_Periph) TERR1() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(TERR1)}}
}

func (p *CAN_Periph) ABRQ1() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(ABRQ1)}}
}

func (p *CAN_Periph) RQCP2() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(RQCP2)}}
}

func (p *CAN_Periph) TXOK2() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(TXOK2)}}
}

func (p *CAN_Periph) ALST2() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(ALST2)}}
}

func (p *CAN_Periph) TERR2() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(TERR2)}}
}

func (p *CAN_Periph) ABRQ2() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(ABRQ2)}}
}

func (p *CAN_Periph) CODE() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(CODE)}}
}

func (p *CAN_Periph) TME() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(TME)}}
}

func (p *CAN_Periph) LOW() RMTSR {
	return RMTSR{mmio.UM32{&p.TSR.U32, uint32(LOW)}}
}

type RF0R uint32

func (b RF0R) Field(mask RF0R) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RF0R) J(v int) RF0R {
	return RF0R(bits.Make32(v, uint32(mask)))
}

type RRF0R struct{ mmio.U32 }

func (r *RRF0R) Bits(mask RF0R) RF0R    { return RF0R(r.U32.Bits(uint32(mask))) }
func (r *RRF0R) StoreBits(mask, b RF0R) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRF0R) SetBits(mask RF0R)      { r.U32.SetBits(uint32(mask)) }
func (r *RRF0R) ClearBits(mask RF0R)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRF0R) Load() RF0R             { return RF0R(r.U32.Load()) }
func (r *RRF0R) Store(b RF0R)           { r.U32.Store(uint32(b)) }

func (r *RRF0R) AtomicStoreBits(mask, b RF0R) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRF0R) AtomicSetBits(mask RF0R)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRF0R) AtomicClearBits(mask RF0R)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRF0R struct{ mmio.UM32 }

func (rm RMRF0R) Load() RF0R   { return RF0R(rm.UM32.Load()) }
func (rm RMRF0R) Store(b RF0R) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FMP0() RMRF0R {
	return RMRF0R{mmio.UM32{&p.RF0R.U32, uint32(FMP0)}}
}

func (p *CAN_Periph) FULL0() RMRF0R {
	return RMRF0R{mmio.UM32{&p.RF0R.U32, uint32(FULL0)}}
}

func (p *CAN_Periph) FOVR0() RMRF0R {
	return RMRF0R{mmio.UM32{&p.RF0R.U32, uint32(FOVR0)}}
}

func (p *CAN_Periph) RFOM0() RMRF0R {
	return RMRF0R{mmio.UM32{&p.RF0R.U32, uint32(RFOM0)}}
}

type RF1R uint32

func (b RF1R) Field(mask RF1R) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RF1R) J(v int) RF1R {
	return RF1R(bits.Make32(v, uint32(mask)))
}

type RRF1R struct{ mmio.U32 }

func (r *RRF1R) Bits(mask RF1R) RF1R    { return RF1R(r.U32.Bits(uint32(mask))) }
func (r *RRF1R) StoreBits(mask, b RF1R) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRF1R) SetBits(mask RF1R)      { r.U32.SetBits(uint32(mask)) }
func (r *RRF1R) ClearBits(mask RF1R)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRF1R) Load() RF1R             { return RF1R(r.U32.Load()) }
func (r *RRF1R) Store(b RF1R)           { r.U32.Store(uint32(b)) }

func (r *RRF1R) AtomicStoreBits(mask, b RF1R) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRF1R) AtomicSetBits(mask RF1R)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRF1R) AtomicClearBits(mask RF1R)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRF1R struct{ mmio.UM32 }

func (rm RMRF1R) Load() RF1R   { return RF1R(rm.UM32.Load()) }
func (rm RMRF1R) Store(b RF1R) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FMP1() RMRF1R {
	return RMRF1R{mmio.UM32{&p.RF1R.U32, uint32(FMP1)}}
}

func (p *CAN_Periph) FULL1() RMRF1R {
	return RMRF1R{mmio.UM32{&p.RF1R.U32, uint32(FULL1)}}
}

func (p *CAN_Periph) FOVR1() RMRF1R {
	return RMRF1R{mmio.UM32{&p.RF1R.U32, uint32(FOVR1)}}
}

func (p *CAN_Periph) RFOM1() RMRF1R {
	return RMRF1R{mmio.UM32{&p.RF1R.U32, uint32(RFOM1)}}
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

func (p *CAN_Periph) TMEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TMEIE)}}
}

func (p *CAN_Periph) FMPIE0() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(FMPIE0)}}
}

func (p *CAN_Periph) FFIE0() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(FFIE0)}}
}

func (p *CAN_Periph) FOVIE0() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(FOVIE0)}}
}

func (p *CAN_Periph) FMPIE1() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(FMPIE1)}}
}

func (p *CAN_Periph) FFIE1() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(FFIE1)}}
}

func (p *CAN_Periph) FOVIE1() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(FOVIE1)}}
}

func (p *CAN_Periph) EWGIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(EWGIE)}}
}

func (p *CAN_Periph) EPVIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(EPVIE)}}
}

func (p *CAN_Periph) BOFIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(BOFIE)}}
}

func (p *CAN_Periph) LECIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(LECIE)}}
}

func (p *CAN_Periph) ERRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(ERRIE)}}
}

func (p *CAN_Periph) WKUIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(WKUIE)}}
}

func (p *CAN_Periph) SLKIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(SLKIE)}}
}

type ESR uint32

func (b ESR) Field(mask ESR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ESR) J(v int) ESR {
	return ESR(bits.Make32(v, uint32(mask)))
}

type RESR struct{ mmio.U32 }

func (r *RESR) Bits(mask ESR) ESR     { return ESR(r.U32.Bits(uint32(mask))) }
func (r *RESR) StoreBits(mask, b ESR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RESR) SetBits(mask ESR)      { r.U32.SetBits(uint32(mask)) }
func (r *RESR) ClearBits(mask ESR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RESR) Load() ESR             { return ESR(r.U32.Load()) }
func (r *RESR) Store(b ESR)           { r.U32.Store(uint32(b)) }

func (r *RESR) AtomicStoreBits(mask, b ESR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RESR) AtomicSetBits(mask ESR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RESR) AtomicClearBits(mask ESR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMESR struct{ mmio.UM32 }

func (rm RMESR) Load() ESR   { return ESR(rm.UM32.Load()) }
func (rm RMESR) Store(b ESR) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) EWGF() RMESR {
	return RMESR{mmio.UM32{&p.ESR.U32, uint32(EWGF)}}
}

func (p *CAN_Periph) EPVF() RMESR {
	return RMESR{mmio.UM32{&p.ESR.U32, uint32(EPVF)}}
}

func (p *CAN_Periph) BOFF() RMESR {
	return RMESR{mmio.UM32{&p.ESR.U32, uint32(BOFF)}}
}

func (p *CAN_Periph) LEC() RMESR {
	return RMESR{mmio.UM32{&p.ESR.U32, uint32(LEC)}}
}

func (p *CAN_Periph) TEC() RMESR {
	return RMESR{mmio.UM32{&p.ESR.U32, uint32(TEC)}}
}

func (p *CAN_Periph) REC() RMESR {
	return RMESR{mmio.UM32{&p.ESR.U32, uint32(REC)}}
}

type BTR uint32

func (b BTR) Field(mask BTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BTR) J(v int) BTR {
	return BTR(bits.Make32(v, uint32(mask)))
}

type RBTR struct{ mmio.U32 }

func (r *RBTR) Bits(mask BTR) BTR     { return BTR(r.U32.Bits(uint32(mask))) }
func (r *RBTR) StoreBits(mask, b BTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBTR) SetBits(mask BTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBTR) ClearBits(mask BTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBTR) Load() BTR             { return BTR(r.U32.Load()) }
func (r *RBTR) Store(b BTR)           { r.U32.Store(uint32(b)) }

func (r *RBTR) AtomicStoreBits(mask, b BTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBTR) AtomicSetBits(mask BTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBTR) AtomicClearBits(mask BTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBTR struct{ mmio.UM32 }

func (rm RMBTR) Load() BTR   { return BTR(rm.UM32.Load()) }
func (rm RMBTR) Store(b BTR) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) BRP() RMBTR {
	return RMBTR{mmio.UM32{&p.BTR.U32, uint32(BRP)}}
}

func (p *CAN_Periph) TS1() RMBTR {
	return RMBTR{mmio.UM32{&p.BTR.U32, uint32(TS1)}}
}

func (p *CAN_Periph) TS2() RMBTR {
	return RMBTR{mmio.UM32{&p.BTR.U32, uint32(TS2)}}
}

func (p *CAN_Periph) SJW() RMBTR {
	return RMBTR{mmio.UM32{&p.BTR.U32, uint32(SJW)}}
}

func (p *CAN_Periph) LBKM() RMBTR {
	return RMBTR{mmio.UM32{&p.BTR.U32, uint32(LBKM)}}
}

func (p *CAN_Periph) SILM() RMBTR {
	return RMBTR{mmio.UM32{&p.BTR.U32, uint32(SILM)}}
}

type FMR uint32

func (b FMR) Field(mask FMR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FMR) J(v int) FMR {
	return FMR(bits.Make32(v, uint32(mask)))
}

type RFMR struct{ mmio.U32 }

func (r *RFMR) Bits(mask FMR) FMR     { return FMR(r.U32.Bits(uint32(mask))) }
func (r *RFMR) StoreBits(mask, b FMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFMR) SetBits(mask FMR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFMR) ClearBits(mask FMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFMR) Load() FMR             { return FMR(r.U32.Load()) }
func (r *RFMR) Store(b FMR)           { r.U32.Store(uint32(b)) }

func (r *RFMR) AtomicStoreBits(mask, b FMR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFMR) AtomicSetBits(mask FMR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFMR) AtomicClearBits(mask FMR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFMR struct{ mmio.UM32 }

func (rm RMFMR) Load() FMR   { return FMR(rm.UM32.Load()) }
func (rm RMFMR) Store(b FMR) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FINIT() RMFMR {
	return RMFMR{mmio.UM32{&p.FMR.U32, uint32(FINIT)}}
}

func (p *CAN_Periph) CAN2SB() RMFMR {
	return RMFMR{mmio.UM32{&p.FMR.U32, uint32(CAN2SB)}}
}

type FM1R uint32

func (b FM1R) Field(mask FM1R) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FM1R) J(v int) FM1R {
	return FM1R(bits.Make32(v, uint32(mask)))
}

type RFM1R struct{ mmio.U32 }

func (r *RFM1R) Bits(mask FM1R) FM1R    { return FM1R(r.U32.Bits(uint32(mask))) }
func (r *RFM1R) StoreBits(mask, b FM1R) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFM1R) SetBits(mask FM1R)      { r.U32.SetBits(uint32(mask)) }
func (r *RFM1R) ClearBits(mask FM1R)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFM1R) Load() FM1R             { return FM1R(r.U32.Load()) }
func (r *RFM1R) Store(b FM1R)           { r.U32.Store(uint32(b)) }

func (r *RFM1R) AtomicStoreBits(mask, b FM1R) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFM1R) AtomicSetBits(mask FM1R)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFM1R) AtomicClearBits(mask FM1R)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFM1R struct{ mmio.UM32 }

func (rm RMFM1R) Load() FM1R   { return FM1R(rm.UM32.Load()) }
func (rm RMFM1R) Store(b FM1R) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FBM() RMFM1R {
	return RMFM1R{mmio.UM32{&p.FM1R.U32, uint32(FBM)}}
}

type FS1R uint32

func (b FS1R) Field(mask FS1R) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FS1R) J(v int) FS1R {
	return FS1R(bits.Make32(v, uint32(mask)))
}

type RFS1R struct{ mmio.U32 }

func (r *RFS1R) Bits(mask FS1R) FS1R    { return FS1R(r.U32.Bits(uint32(mask))) }
func (r *RFS1R) StoreBits(mask, b FS1R) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFS1R) SetBits(mask FS1R)      { r.U32.SetBits(uint32(mask)) }
func (r *RFS1R) ClearBits(mask FS1R)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFS1R) Load() FS1R             { return FS1R(r.U32.Load()) }
func (r *RFS1R) Store(b FS1R)           { r.U32.Store(uint32(b)) }

func (r *RFS1R) AtomicStoreBits(mask, b FS1R) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFS1R) AtomicSetBits(mask FS1R)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFS1R) AtomicClearBits(mask FS1R)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFS1R struct{ mmio.UM32 }

func (rm RMFS1R) Load() FS1R   { return FS1R(rm.UM32.Load()) }
func (rm RMFS1R) Store(b FS1R) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FSC() RMFS1R {
	return RMFS1R{mmio.UM32{&p.FS1R.U32, uint32(FSC)}}
}

type FFA1R uint32

func (b FFA1R) Field(mask FFA1R) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FFA1R) J(v int) FFA1R {
	return FFA1R(bits.Make32(v, uint32(mask)))
}

type RFFA1R struct{ mmio.U32 }

func (r *RFFA1R) Bits(mask FFA1R) FFA1R   { return FFA1R(r.U32.Bits(uint32(mask))) }
func (r *RFFA1R) StoreBits(mask, b FFA1R) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFFA1R) SetBits(mask FFA1R)      { r.U32.SetBits(uint32(mask)) }
func (r *RFFA1R) ClearBits(mask FFA1R)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFFA1R) Load() FFA1R             { return FFA1R(r.U32.Load()) }
func (r *RFFA1R) Store(b FFA1R)           { r.U32.Store(uint32(b)) }

func (r *RFFA1R) AtomicStoreBits(mask, b FFA1R) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFFA1R) AtomicSetBits(mask FFA1R)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFFA1R) AtomicClearBits(mask FFA1R)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFFA1R struct{ mmio.UM32 }

func (rm RMFFA1R) Load() FFA1R   { return FFA1R(rm.UM32.Load()) }
func (rm RMFFA1R) Store(b FFA1R) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FFA() RMFFA1R {
	return RMFFA1R{mmio.UM32{&p.FFA1R.U32, uint32(FFA)}}
}

type FA1R uint32

func (b FA1R) Field(mask FA1R) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FA1R) J(v int) FA1R {
	return FA1R(bits.Make32(v, uint32(mask)))
}

type RFA1R struct{ mmio.U32 }

func (r *RFA1R) Bits(mask FA1R) FA1R    { return FA1R(r.U32.Bits(uint32(mask))) }
func (r *RFA1R) StoreBits(mask, b FA1R) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFA1R) SetBits(mask FA1R)      { r.U32.SetBits(uint32(mask)) }
func (r *RFA1R) ClearBits(mask FA1R)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFA1R) Load() FA1R             { return FA1R(r.U32.Load()) }
func (r *RFA1R) Store(b FA1R)           { r.U32.Store(uint32(b)) }

func (r *RFA1R) AtomicStoreBits(mask, b FA1R) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFA1R) AtomicSetBits(mask FA1R)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFA1R) AtomicClearBits(mask FA1R)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFA1R struct{ mmio.UM32 }

func (rm RMFA1R) Load() FA1R   { return FA1R(rm.UM32.Load()) }
func (rm RMFA1R) Store(b FA1R) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FACT() RMFA1R {
	return RMFA1R{mmio.UM32{&p.FA1R.U32, uint32(FACT)}}
}
