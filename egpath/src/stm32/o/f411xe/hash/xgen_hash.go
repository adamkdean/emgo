package hash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type HASH_Periph struct {
	CR  RCR
	DIN RDIN
	STR RSTR
	HR  [5]RHR
	IMR RIMR
	SR  RSR
	_   [52]uint32
	CSR [54]RCSR
}

func (p *HASH_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var HASH = (*HASH_Periph)(unsafe.Pointer(uintptr(mmap.HASH_BASE)))

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

func (p *HASH_Periph) INIT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(INIT)}}
}

func (p *HASH_Periph) DMAE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DMAE)}}
}

func (p *HASH_Periph) DATATYPE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DATATYPE)}}
}

func (p *HASH_Periph) MODE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MODE)}}
}

func (p *HASH_Periph) ALGO() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALGO)}}
}

func (p *HASH_Periph) NBW() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(NBW)}}
}

func (p *HASH_Periph) DINNE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DINNE)}}
}

func (p *HASH_Periph) MDMAT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MDMAT)}}
}

func (p *HASH_Periph) LKEY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LKEY)}}
}

type DIN uint32

func (b DIN) Field(mask DIN) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIN) J(v int) DIN {
	return DIN(bits.Make32(v, uint32(mask)))
}

type RDIN struct{ mmio.U32 }

func (r *RDIN) Bits(mask DIN) DIN     { return DIN(r.U32.Bits(uint32(mask))) }
func (r *RDIN) StoreBits(mask, b DIN) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIN) SetBits(mask DIN)      { r.U32.SetBits(uint32(mask)) }
func (r *RDIN) ClearBits(mask DIN)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDIN) Load() DIN             { return DIN(r.U32.Load()) }
func (r *RDIN) Store(b DIN)           { r.U32.Store(uint32(b)) }

func (r *RDIN) AtomicStoreBits(mask, b DIN) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDIN) AtomicSetBits(mask DIN)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDIN) AtomicClearBits(mask DIN)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDIN struct{ mmio.UM32 }

func (rm RMDIN) Load() DIN   { return DIN(rm.UM32.Load()) }
func (rm RMDIN) Store(b DIN) { rm.UM32.Store(uint32(b)) }

type STR uint32

func (b STR) Field(mask STR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask STR) J(v int) STR {
	return STR(bits.Make32(v, uint32(mask)))
}

type RSTR struct{ mmio.U32 }

func (r *RSTR) Bits(mask STR) STR     { return STR(r.U32.Bits(uint32(mask))) }
func (r *RSTR) StoreBits(mask, b STR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSTR) SetBits(mask STR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSTR) ClearBits(mask STR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSTR) Load() STR             { return STR(r.U32.Load()) }
func (r *RSTR) Store(b STR)           { r.U32.Store(uint32(b)) }

func (r *RSTR) AtomicStoreBits(mask, b STR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSTR) AtomicSetBits(mask STR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSTR) AtomicClearBits(mask STR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSTR struct{ mmio.UM32 }

func (rm RMSTR) Load() STR   { return STR(rm.UM32.Load()) }
func (rm RMSTR) Store(b STR) { rm.UM32.Store(uint32(b)) }

func (p *HASH_Periph) NBW() RMSTR {
	return RMSTR{mmio.UM32{&p.STR.U32, uint32(NBW)}}
}

func (p *HASH_Periph) DCAL() RMSTR {
	return RMSTR{mmio.UM32{&p.STR.U32, uint32(DCAL)}}
}

type HR uint32

func (b HR) Field(mask HR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HR) J(v int) HR {
	return HR(bits.Make32(v, uint32(mask)))
}

type RHR struct{ mmio.U32 }

func (r *RHR) Bits(mask HR) HR      { return HR(r.U32.Bits(uint32(mask))) }
func (r *RHR) StoreBits(mask, b HR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHR) SetBits(mask HR)      { r.U32.SetBits(uint32(mask)) }
func (r *RHR) ClearBits(mask HR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHR) Load() HR             { return HR(r.U32.Load()) }
func (r *RHR) Store(b HR)           { r.U32.Store(uint32(b)) }

func (r *RHR) AtomicStoreBits(mask, b HR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHR) AtomicSetBits(mask HR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHR) AtomicClearBits(mask HR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHR struct{ mmio.UM32 }

func (rm RMHR) Load() HR   { return HR(rm.UM32.Load()) }
func (rm RMHR) Store(b HR) { rm.UM32.Store(uint32(b)) }

type IMR uint32

func (b IMR) Field(mask IMR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IMR) J(v int) IMR {
	return IMR(bits.Make32(v, uint32(mask)))
}

type RIMR struct{ mmio.U32 }

func (r *RIMR) Bits(mask IMR) IMR     { return IMR(r.U32.Bits(uint32(mask))) }
func (r *RIMR) StoreBits(mask, b IMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIMR) SetBits(mask IMR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIMR) ClearBits(mask IMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIMR) Load() IMR             { return IMR(r.U32.Load()) }
func (r *RIMR) Store(b IMR)           { r.U32.Store(uint32(b)) }

func (r *RIMR) AtomicStoreBits(mask, b IMR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIMR) AtomicSetBits(mask IMR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIMR) AtomicClearBits(mask IMR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIMR struct{ mmio.UM32 }

func (rm RMIMR) Load() IMR   { return IMR(rm.UM32.Load()) }
func (rm RMIMR) Store(b IMR) { rm.UM32.Store(uint32(b)) }

func (p *HASH_Periph) DINIM() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(DINIM)}}
}

func (p *HASH_Periph) DCIM() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(DCIM)}}
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

func (p *HASH_Periph) DINIS() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(DINIS)}}
}

func (p *HASH_Periph) DCIS() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(DCIS)}}
}

func (p *HASH_Periph) DMAS() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(DMAS)}}
}

func (p *HASH_Periph) BUSY() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(BUSY)}}
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
