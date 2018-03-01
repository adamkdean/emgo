package sdio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type SDIO_Periph struct {
	POWER  RPOWER
	CLKCR  RCLKCR
	ARG    RARG
	CMD    RCMD
	DTIMER RDTIMER
	DLEN   RDLEN
	DCTRL  RDCTRL
	ICR    RICR
	MASK   RMASK
	_      [15]uint32
	FIFO   RFIFO
}

func (p *SDIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SDIO = (*SDIO_Periph)(unsafe.Pointer(uintptr(mmap.SDIO_BASE)))

type POWER uint32

func (b POWER) Field(mask POWER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask POWER) J(v int) POWER {
	return POWER(bits.Make32(v, uint32(mask)))
}

type RPOWER struct{ mmio.U32 }

func (r *RPOWER) Bits(mask POWER) POWER   { return POWER(r.U32.Bits(uint32(mask))) }
func (r *RPOWER) StoreBits(mask, b POWER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPOWER) SetBits(mask POWER)      { r.U32.SetBits(uint32(mask)) }
func (r *RPOWER) ClearBits(mask POWER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPOWER) Load() POWER             { return POWER(r.U32.Load()) }
func (r *RPOWER) Store(b POWER)           { r.U32.Store(uint32(b)) }

func (r *RPOWER) AtomicStoreBits(mask, b POWER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPOWER) AtomicSetBits(mask POWER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPOWER) AtomicClearBits(mask POWER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPOWER struct{ mmio.UM32 }

func (rm RMPOWER) Load() POWER   { return POWER(rm.UM32.Load()) }
func (rm RMPOWER) Store(b POWER) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) PWRCTRL() RMPOWER {
	return RMPOWER{mmio.UM32{&p.POWER.U32, uint32(PWRCTRL)}}
}

type CLKCR uint32

func (b CLKCR) Field(mask CLKCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLKCR) J(v int) CLKCR {
	return CLKCR(bits.Make32(v, uint32(mask)))
}

type RCLKCR struct{ mmio.U32 }

func (r *RCLKCR) Bits(mask CLKCR) CLKCR   { return CLKCR(r.U32.Bits(uint32(mask))) }
func (r *RCLKCR) StoreBits(mask, b CLKCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCLKCR) SetBits(mask CLKCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCLKCR) ClearBits(mask CLKCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCLKCR) Load() CLKCR             { return CLKCR(r.U32.Load()) }
func (r *RCLKCR) Store(b CLKCR)           { r.U32.Store(uint32(b)) }

func (r *RCLKCR) AtomicStoreBits(mask, b CLKCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCLKCR) AtomicSetBits(mask CLKCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCLKCR) AtomicClearBits(mask CLKCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCLKCR struct{ mmio.UM32 }

func (rm RMCLKCR) Load() CLKCR   { return CLKCR(rm.UM32.Load()) }
func (rm RMCLKCR) Store(b CLKCR) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CLKDIV() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(CLKDIV)}}
}

func (p *SDIO_Periph) CLKEN() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(CLKEN)}}
}

func (p *SDIO_Periph) PWRSAV() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(PWRSAV)}}
}

func (p *SDIO_Periph) BYPASS() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(BYPASS)}}
}

func (p *SDIO_Periph) WIDBUS() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(WIDBUS)}}
}

func (p *SDIO_Periph) NEGEDGE() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(NEGEDGE)}}
}

func (p *SDIO_Periph) HWFC_EN() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(HWFC_EN)}}
}

type ARG uint32

func (b ARG) Field(mask ARG) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ARG) J(v int) ARG {
	return ARG(bits.Make32(v, uint32(mask)))
}

type RARG struct{ mmio.U32 }

func (r *RARG) Bits(mask ARG) ARG     { return ARG(r.U32.Bits(uint32(mask))) }
func (r *RARG) StoreBits(mask, b ARG) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RARG) SetBits(mask ARG)      { r.U32.SetBits(uint32(mask)) }
func (r *RARG) ClearBits(mask ARG)    { r.U32.ClearBits(uint32(mask)) }
func (r *RARG) Load() ARG             { return ARG(r.U32.Load()) }
func (r *RARG) Store(b ARG)           { r.U32.Store(uint32(b)) }

func (r *RARG) AtomicStoreBits(mask, b ARG) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RARG) AtomicSetBits(mask ARG)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RARG) AtomicClearBits(mask ARG)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMARG struct{ mmio.UM32 }

func (rm RMARG) Load() ARG   { return ARG(rm.UM32.Load()) }
func (rm RMARG) Store(b ARG) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CMDARG() RMARG {
	return RMARG{mmio.UM32{&p.ARG.U32, uint32(CMDARG)}}
}

type CMD uint32

func (b CMD) Field(mask CMD) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMD) J(v int) CMD {
	return CMD(bits.Make32(v, uint32(mask)))
}

type RCMD struct{ mmio.U32 }

func (r *RCMD) Bits(mask CMD) CMD     { return CMD(r.U32.Bits(uint32(mask))) }
func (r *RCMD) StoreBits(mask, b CMD) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCMD) SetBits(mask CMD)      { r.U32.SetBits(uint32(mask)) }
func (r *RCMD) ClearBits(mask CMD)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCMD) Load() CMD             { return CMD(r.U32.Load()) }
func (r *RCMD) Store(b CMD)           { r.U32.Store(uint32(b)) }

func (r *RCMD) AtomicStoreBits(mask, b CMD) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCMD) AtomicSetBits(mask CMD)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCMD) AtomicClearBits(mask CMD)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCMD struct{ mmio.UM32 }

func (rm RMCMD) Load() CMD   { return CMD(rm.UM32.Load()) }
func (rm RMCMD) Store(b CMD) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CMDINDEX() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(CMDINDEX)}}
}

func (p *SDIO_Periph) WAITRESP() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(WAITRESP)}}
}

func (p *SDIO_Periph) WAITINT() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(WAITINT)}}
}

func (p *SDIO_Periph) WAITPEND() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(WAITPEND)}}
}

func (p *SDIO_Periph) CPSMEN() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(CPSMEN)}}
}

func (p *SDIO_Periph) SDIOSUSPEND() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(SDIOSUSPEND)}}
}

func (p *SDIO_Periph) ENCMDCOMPL() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(ENCMDCOMPL)}}
}

func (p *SDIO_Periph) NIEN() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(NIEN)}}
}

func (p *SDIO_Periph) CEATACMD() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(CEATACMD)}}
}

type DTIMER uint32

func (b DTIMER) Field(mask DTIMER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DTIMER) J(v int) DTIMER {
	return DTIMER(bits.Make32(v, uint32(mask)))
}

type RDTIMER struct{ mmio.U32 }

func (r *RDTIMER) Bits(mask DTIMER) DTIMER  { return DTIMER(r.U32.Bits(uint32(mask))) }
func (r *RDTIMER) StoreBits(mask, b DTIMER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDTIMER) SetBits(mask DTIMER)      { r.U32.SetBits(uint32(mask)) }
func (r *RDTIMER) ClearBits(mask DTIMER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDTIMER) Load() DTIMER             { return DTIMER(r.U32.Load()) }
func (r *RDTIMER) Store(b DTIMER)           { r.U32.Store(uint32(b)) }

func (r *RDTIMER) AtomicStoreBits(mask, b DTIMER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDTIMER) AtomicSetBits(mask DTIMER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDTIMER) AtomicClearBits(mask DTIMER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDTIMER struct{ mmio.UM32 }

func (rm RMDTIMER) Load() DTIMER   { return DTIMER(rm.UM32.Load()) }
func (rm RMDTIMER) Store(b DTIMER) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) DATATIME() RMDTIMER {
	return RMDTIMER{mmio.UM32{&p.DTIMER.U32, uint32(DATATIME)}}
}

type DLEN uint32

func (b DLEN) Field(mask DLEN) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DLEN) J(v int) DLEN {
	return DLEN(bits.Make32(v, uint32(mask)))
}

type RDLEN struct{ mmio.U32 }

func (r *RDLEN) Bits(mask DLEN) DLEN    { return DLEN(r.U32.Bits(uint32(mask))) }
func (r *RDLEN) StoreBits(mask, b DLEN) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDLEN) SetBits(mask DLEN)      { r.U32.SetBits(uint32(mask)) }
func (r *RDLEN) ClearBits(mask DLEN)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDLEN) Load() DLEN             { return DLEN(r.U32.Load()) }
func (r *RDLEN) Store(b DLEN)           { r.U32.Store(uint32(b)) }

func (r *RDLEN) AtomicStoreBits(mask, b DLEN) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDLEN) AtomicSetBits(mask DLEN)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDLEN) AtomicClearBits(mask DLEN)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDLEN struct{ mmio.UM32 }

func (rm RMDLEN) Load() DLEN   { return DLEN(rm.UM32.Load()) }
func (rm RMDLEN) Store(b DLEN) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) DATALENGTH() RMDLEN {
	return RMDLEN{mmio.UM32{&p.DLEN.U32, uint32(DATALENGTH)}}
}

type DCTRL uint32

func (b DCTRL) Field(mask DCTRL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCTRL) J(v int) DCTRL {
	return DCTRL(bits.Make32(v, uint32(mask)))
}

type RDCTRL struct{ mmio.U32 }

func (r *RDCTRL) Bits(mask DCTRL) DCTRL   { return DCTRL(r.U32.Bits(uint32(mask))) }
func (r *RDCTRL) StoreBits(mask, b DCTRL) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDCTRL) SetBits(mask DCTRL)      { r.U32.SetBits(uint32(mask)) }
func (r *RDCTRL) ClearBits(mask DCTRL)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDCTRL) Load() DCTRL             { return DCTRL(r.U32.Load()) }
func (r *RDCTRL) Store(b DCTRL)           { r.U32.Store(uint32(b)) }

func (r *RDCTRL) AtomicStoreBits(mask, b DCTRL) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDCTRL) AtomicSetBits(mask DCTRL)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDCTRL) AtomicClearBits(mask DCTRL)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDCTRL struct{ mmio.UM32 }

func (rm RMDCTRL) Load() DCTRL   { return DCTRL(rm.UM32.Load()) }
func (rm RMDCTRL) Store(b DCTRL) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) DTEN() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DTEN)}}
}

func (p *SDIO_Periph) DTDIR() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DTDIR)}}
}

func (p *SDIO_Periph) DTMODE() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DTMODE)}}
}

func (p *SDIO_Periph) DMAEN() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DMAEN)}}
}

func (p *SDIO_Periph) DBLOCKSIZE() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DBLOCKSIZE)}}
}

func (p *SDIO_Periph) RWSTART() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(RWSTART)}}
}

func (p *SDIO_Periph) RWSTOP() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(RWSTOP)}}
}

func (p *SDIO_Periph) RWMOD() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(RWMOD)}}
}

func (p *SDIO_Periph) SDIOEN() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(SDIOEN)}}
}

type ICR uint32

func (b ICR) Field(mask ICR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR) J(v int) ICR {
	return ICR(bits.Make32(v, uint32(mask)))
}

type RICR struct{ mmio.U32 }

func (r *RICR) Bits(mask ICR) ICR     { return ICR(r.U32.Bits(uint32(mask))) }
func (r *RICR) StoreBits(mask, b ICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICR) SetBits(mask ICR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICR) ClearBits(mask ICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICR) Load() ICR             { return ICR(r.U32.Load()) }
func (r *RICR) Store(b ICR)           { r.U32.Store(uint32(b)) }

func (r *RICR) AtomicStoreBits(mask, b ICR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RICR) AtomicSetBits(mask ICR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RICR) AtomicClearBits(mask ICR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMICR struct{ mmio.UM32 }

func (rm RMICR) Load() ICR   { return ICR(rm.UM32.Load()) }
func (rm RMICR) Store(b ICR) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CCRCFAILC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CCRCFAILC)}}
}

func (p *SDIO_Periph) DCRCFAILC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(DCRCFAILC)}}
}

func (p *SDIO_Periph) CTIMEOUTC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CTIMEOUTC)}}
}

func (p *SDIO_Periph) DTIMEOUTC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(DTIMEOUTC)}}
}

func (p *SDIO_Periph) TXUNDERRC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(TXUNDERRC)}}
}

func (p *SDIO_Periph) RXOVERRC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(RXOVERRC)}}
}

func (p *SDIO_Periph) CMDRENDC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CMDRENDC)}}
}

func (p *SDIO_Periph) CMDSENTC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CMDSENTC)}}
}

func (p *SDIO_Periph) DATAENDC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(DATAENDC)}}
}

func (p *SDIO_Periph) STBITERRC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(STBITERRC)}}
}

func (p *SDIO_Periph) DBCKENDC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(DBCKENDC)}}
}

func (p *SDIO_Periph) SDIOITC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(SDIOITC)}}
}

func (p *SDIO_Periph) CEATAENDC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CEATAENDC)}}
}

type MASK uint32

func (b MASK) Field(mask MASK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MASK) J(v int) MASK {
	return MASK(bits.Make32(v, uint32(mask)))
}

type RMASK struct{ mmio.U32 }

func (r *RMASK) Bits(mask MASK) MASK    { return MASK(r.U32.Bits(uint32(mask))) }
func (r *RMASK) StoreBits(mask, b MASK) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMASK) SetBits(mask MASK)      { r.U32.SetBits(uint32(mask)) }
func (r *RMASK) ClearBits(mask MASK)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMASK) Load() MASK             { return MASK(r.U32.Load()) }
func (r *RMASK) Store(b MASK)           { r.U32.Store(uint32(b)) }

func (r *RMASK) AtomicStoreBits(mask, b MASK) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMASK) AtomicSetBits(mask MASK)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMASK) AtomicClearBits(mask MASK)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMASK struct{ mmio.UM32 }

func (rm RMMASK) Load() MASK   { return MASK(rm.UM32.Load()) }
func (rm RMMASK) Store(b MASK) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CCRCFAILIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CCRCFAILIE)}}
}

func (p *SDIO_Periph) DCRCFAILIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(DCRCFAILIE)}}
}

func (p *SDIO_Periph) CTIMEOUTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CTIMEOUTIE)}}
}

func (p *SDIO_Periph) DTIMEOUTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(DTIMEOUTIE)}}
}

func (p *SDIO_Periph) TXUNDERRIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXUNDERRIE)}}
}

func (p *SDIO_Periph) RXOVERRIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXOVERRIE)}}
}

func (p *SDIO_Periph) CMDRENDIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CMDRENDIE)}}
}

func (p *SDIO_Periph) CMDSENTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CMDSENTIE)}}
}

func (p *SDIO_Periph) DATAENDIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(DATAENDIE)}}
}

func (p *SDIO_Periph) STBITERRIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(STBITERRIE)}}
}

func (p *SDIO_Periph) DBCKENDIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(DBCKENDIE)}}
}

func (p *SDIO_Periph) CMDACTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CMDACTIE)}}
}

func (p *SDIO_Periph) TXACTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXACTIE)}}
}

func (p *SDIO_Periph) RXACTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXACTIE)}}
}

func (p *SDIO_Periph) TXFIFOHEIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXFIFOHEIE)}}
}

func (p *SDIO_Periph) RXFIFOHFIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXFIFOHFIE)}}
}

func (p *SDIO_Periph) TXFIFOFIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXFIFOFIE)}}
}

func (p *SDIO_Periph) RXFIFOFIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXFIFOFIE)}}
}

func (p *SDIO_Periph) TXFIFOEIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXFIFOEIE)}}
}

func (p *SDIO_Periph) RXFIFOEIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXFIFOEIE)}}
}

func (p *SDIO_Periph) TXDAVLIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXDAVLIE)}}
}

func (p *SDIO_Periph) RXDAVLIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXDAVLIE)}}
}

func (p *SDIO_Periph) SDIOITIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(SDIOITIE)}}
}

func (p *SDIO_Periph) CEATAENDIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CEATAENDIE)}}
}

type FIFO uint32

func (b FIFO) Field(mask FIFO) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FIFO) J(v int) FIFO {
	return FIFO(bits.Make32(v, uint32(mask)))
}

type RFIFO struct{ mmio.U32 }

func (r *RFIFO) Bits(mask FIFO) FIFO    { return FIFO(r.U32.Bits(uint32(mask))) }
func (r *RFIFO) StoreBits(mask, b FIFO) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFIFO) SetBits(mask FIFO)      { r.U32.SetBits(uint32(mask)) }
func (r *RFIFO) ClearBits(mask FIFO)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFIFO) Load() FIFO             { return FIFO(r.U32.Load()) }
func (r *RFIFO) Store(b FIFO)           { r.U32.Store(uint32(b)) }

func (r *RFIFO) AtomicStoreBits(mask, b FIFO) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFIFO) AtomicSetBits(mask FIFO)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFIFO) AtomicClearBits(mask FIFO)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFIFO struct{ mmio.UM32 }

func (rm RMFIFO) Load() FIFO   { return FIFO(rm.UM32.Load()) }
func (rm RMFIFO) Store(b FIFO) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) FIFODATA() RMFIFO {
	return RMFIFO{mmio.UM32{&p.FIFO.U32, uint32(FIFODATA)}}
}
