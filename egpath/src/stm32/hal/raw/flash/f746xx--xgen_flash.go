// +build f746xx

package flash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type FLASH_Periph struct {
	ACR     RACR
	KEYR    RKEYR
	OPTKEYR ROPTKEYR
	SR      RSR
	CR      RCR
	OPTCR   [2]ROPTCR
}

func (p *FLASH_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FLASH = (*FLASH_Periph)(unsafe.Pointer(uintptr(mmap.FLASH_R_BASE)))

type ACR uint32

func (b ACR) Field(mask ACR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ACR) J(v int) ACR {
	return ACR(bits.Make32(v, uint32(mask)))
}

type RACR struct{ mmio.U32 }

func (r *RACR) Bits(mask ACR) ACR     { return ACR(r.U32.Bits(uint32(mask))) }
func (r *RACR) StoreBits(mask, b ACR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RACR) SetBits(mask ACR)      { r.U32.SetBits(uint32(mask)) }
func (r *RACR) ClearBits(mask ACR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RACR) Load() ACR             { return ACR(r.U32.Load()) }
func (r *RACR) Store(b ACR)           { r.U32.Store(uint32(b)) }

func (r *RACR) AtomicStoreBits(mask, b ACR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RACR) AtomicSetBits(mask ACR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RACR) AtomicClearBits(mask ACR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMACR struct{ mmio.UM32 }

func (rm RMACR) Load() ACR   { return ACR(rm.UM32.Load()) }
func (rm RMACR) Store(b ACR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) LATENCY() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(LATENCY)}}
}

func (p *FLASH_Periph) PRFTEN() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(PRFTEN)}}
}

func (p *FLASH_Periph) ARTEN() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(ARTEN)}}
}

func (p *FLASH_Periph) ARTRST() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(ARTRST)}}
}

type KEYR uint32

func (b KEYR) Field(mask KEYR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask KEYR) J(v int) KEYR {
	return KEYR(bits.Make32(v, uint32(mask)))
}

type RKEYR struct{ mmio.U32 }

func (r *RKEYR) Bits(mask KEYR) KEYR    { return KEYR(r.U32.Bits(uint32(mask))) }
func (r *RKEYR) StoreBits(mask, b KEYR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RKEYR) SetBits(mask KEYR)      { r.U32.SetBits(uint32(mask)) }
func (r *RKEYR) ClearBits(mask KEYR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RKEYR) Load() KEYR             { return KEYR(r.U32.Load()) }
func (r *RKEYR) Store(b KEYR)           { r.U32.Store(uint32(b)) }

func (r *RKEYR) AtomicStoreBits(mask, b KEYR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RKEYR) AtomicSetBits(mask KEYR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RKEYR) AtomicClearBits(mask KEYR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMKEYR struct{ mmio.UM32 }

func (rm RMKEYR) Load() KEYR   { return KEYR(rm.UM32.Load()) }
func (rm RMKEYR) Store(b KEYR) { rm.UM32.Store(uint32(b)) }

type OPTKEYR uint32

func (b OPTKEYR) Field(mask OPTKEYR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTKEYR) J(v int) OPTKEYR {
	return OPTKEYR(bits.Make32(v, uint32(mask)))
}

type ROPTKEYR struct{ mmio.U32 }

func (r *ROPTKEYR) Bits(mask OPTKEYR) OPTKEYR { return OPTKEYR(r.U32.Bits(uint32(mask))) }
func (r *ROPTKEYR) StoreBits(mask, b OPTKEYR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROPTKEYR) SetBits(mask OPTKEYR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROPTKEYR) ClearBits(mask OPTKEYR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROPTKEYR) Load() OPTKEYR             { return OPTKEYR(r.U32.Load()) }
func (r *ROPTKEYR) Store(b OPTKEYR)           { r.U32.Store(uint32(b)) }

func (r *ROPTKEYR) AtomicStoreBits(mask, b OPTKEYR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROPTKEYR) AtomicSetBits(mask OPTKEYR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROPTKEYR) AtomicClearBits(mask OPTKEYR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOPTKEYR struct{ mmio.UM32 }

func (rm RMOPTKEYR) Load() OPTKEYR   { return OPTKEYR(rm.UM32.Load()) }
func (rm RMOPTKEYR) Store(b OPTKEYR) { rm.UM32.Store(uint32(b)) }

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

func (p *FLASH_Periph) EOP() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(EOP)}}
}

func (p *FLASH_Periph) OPERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(OPERR)}}
}

func (p *FLASH_Periph) WRPERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(WRPERR)}}
}

func (p *FLASH_Periph) PGAERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PGAERR)}}
}

func (p *FLASH_Periph) PGPERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PGPERR)}}
}

func (p *FLASH_Periph) ERSERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(ERSERR)}}
}

func (p *FLASH_Periph) BSY() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(BSY)}}
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

func (p *FLASH_Periph) PG() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PG)}}
}

func (p *FLASH_Periph) SER() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SER)}}
}

func (p *FLASH_Periph) MER() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MER)}}
}

func (p *FLASH_Periph) SNB() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SNB)}}
}

func (p *FLASH_Periph) PSIZE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PSIZE)}}
}

func (p *FLASH_Periph) STRT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(STRT)}}
}

func (p *FLASH_Periph) EOPIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EOPIE)}}
}

func (p *FLASH_Periph) ERRIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ERRIE)}}
}

func (p *FLASH_Periph) LOCK() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LOCK)}}
}

type OPTCR uint32

func (b OPTCR) Field(mask OPTCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTCR) J(v int) OPTCR {
	return OPTCR(bits.Make32(v, uint32(mask)))
}

type ROPTCR struct{ mmio.U32 }

func (r *ROPTCR) Bits(mask OPTCR) OPTCR   { return OPTCR(r.U32.Bits(uint32(mask))) }
func (r *ROPTCR) StoreBits(mask, b OPTCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROPTCR) SetBits(mask OPTCR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROPTCR) ClearBits(mask OPTCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROPTCR) Load() OPTCR             { return OPTCR(r.U32.Load()) }
func (r *ROPTCR) Store(b OPTCR)           { r.U32.Store(uint32(b)) }

func (r *ROPTCR) AtomicStoreBits(mask, b OPTCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROPTCR) AtomicSetBits(mask OPTCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROPTCR) AtomicClearBits(mask OPTCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOPTCR struct{ mmio.UM32 }

func (rm RMOPTCR) Load() OPTCR   { return OPTCR(rm.UM32.Load()) }
func (rm RMOPTCR) Store(b OPTCR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) OPTLOCK(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(OPTLOCK)}}
}

func (p *FLASH_Periph) OPTSTRT(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(OPTSTRT)}}
}

func (p *FLASH_Periph) BOR_LEV(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(BOR_LEV)}}
}

func (p *FLASH_Periph) WWDG_SW(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(WWDG_SW)}}
}

func (p *FLASH_Periph) IWDG_SW(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(IWDG_SW)}}
}

func (p *FLASH_Periph) nRST_STOP(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(nRST_STOP)}}
}

func (p *FLASH_Periph) nRST_STDBY(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(nRST_STDBY)}}
}

func (p *FLASH_Periph) RDP(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(RDP)}}
}

func (p *FLASH_Periph) nWRP(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(nWRP)}}
}

func (p *FLASH_Periph) IWDG_STDBY(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(IWDG_STDBY)}}
}

func (p *FLASH_Periph) IWDG_STOP(n int) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR[n].U32, uint32(IWDG_STOP)}}
}
