package ficr

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"nrf5/hal/internal/mmap"
)

type Periph struct {
	_              [4]uint32
	CODEPAGESIZE   RCODEPAGESIZE
	CODESIZE       RCODESIZE
	_              [4]uint32
	CLENR0         RCLENR0
	PPFC           RPPFC
	_              uint32
	NUMRAMBLOCK    RNUMRAMBLOCK
	SIZERAMBLOCK   [4]RSIZERAMBLOCK
	_              [5]uint32
	CONFIGID       RCONFIGID
	DEVICEID       [2]RDEVICEID
	_              [6]uint32
	ER             [4]RER
	IR             [4]RIR
	DEVICEADDRTYPE RDEVICEADDRTYPE
	DEVICEADDR     [2]RDEVICEADDR
	OVERRIDDEN     ROVERRIDDEN
	NRF_1MBIT      [5]RNRF_1MBIT
	_              [10]uint32
	BLE_1MBIT      [5]RBLE_1MBIT
	INFO_PART      RINFO_PART
	INFO_VARIANT   RINFO_VARIANT
	INFO_PACKAGE   RINFO_PACKAGE
	INFO_RAM       RINFO_RAM
	INFO_FLASH     RINFO_FLASH
	_              [188]uint32
	TEMP_A         [6]RTEMP_A
	TEMP_B         [6]RTEMP_B
	TEMP_T         [5]RTEMP_T
	_              [2]uint32
	NFC_TAGHEADER  [4]RNFC_TAGHEADER
}

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FICR = (*Periph)(unsafe.Pointer(uintptr(mmap.FICR_BASE)))

type CODEPAGESIZE uint32

func (b CODEPAGESIZE) Field(mask CODEPAGESIZE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CODEPAGESIZE) J(v int) CODEPAGESIZE {
	return CODEPAGESIZE(bits.Make32(v, uint32(mask)))
}

type RCODEPAGESIZE struct{ mmio.U32 }

func (r *RCODEPAGESIZE) Bits(mask CODEPAGESIZE) CODEPAGESIZE {
	return CODEPAGESIZE(r.U32.Bits(uint32(mask)))
}
func (r *RCODEPAGESIZE) StoreBits(mask, b CODEPAGESIZE) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCODEPAGESIZE) SetBits(mask CODEPAGESIZE)      { r.U32.SetBits(uint32(mask)) }
func (r *RCODEPAGESIZE) ClearBits(mask CODEPAGESIZE)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCODEPAGESIZE) Load() CODEPAGESIZE             { return CODEPAGESIZE(r.U32.Load()) }
func (r *RCODEPAGESIZE) Store(b CODEPAGESIZE)           { r.U32.Store(uint32(b)) }

func (r *RCODEPAGESIZE) AtomicStoreBits(mask, b CODEPAGESIZE) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RCODEPAGESIZE) AtomicSetBits(mask CODEPAGESIZE)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCODEPAGESIZE) AtomicClearBits(mask CODEPAGESIZE) { r.U32.AtomicClearBits(uint32(mask)) }

type RMCODEPAGESIZE struct{ mmio.UM32 }

func (rm RMCODEPAGESIZE) Load() CODEPAGESIZE   { return CODEPAGESIZE(rm.UM32.Load()) }
func (rm RMCODEPAGESIZE) Store(b CODEPAGESIZE) { rm.UM32.Store(uint32(b)) }

type CODESIZE uint32

func (b CODESIZE) Field(mask CODESIZE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CODESIZE) J(v int) CODESIZE {
	return CODESIZE(bits.Make32(v, uint32(mask)))
}

type RCODESIZE struct{ mmio.U32 }

func (r *RCODESIZE) Bits(mask CODESIZE) CODESIZE { return CODESIZE(r.U32.Bits(uint32(mask))) }
func (r *RCODESIZE) StoreBits(mask, b CODESIZE)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCODESIZE) SetBits(mask CODESIZE)       { r.U32.SetBits(uint32(mask)) }
func (r *RCODESIZE) ClearBits(mask CODESIZE)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCODESIZE) Load() CODESIZE              { return CODESIZE(r.U32.Load()) }
func (r *RCODESIZE) Store(b CODESIZE)            { r.U32.Store(uint32(b)) }

func (r *RCODESIZE) AtomicStoreBits(mask, b CODESIZE) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCODESIZE) AtomicSetBits(mask CODESIZE)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCODESIZE) AtomicClearBits(mask CODESIZE)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCODESIZE struct{ mmio.UM32 }

func (rm RMCODESIZE) Load() CODESIZE   { return CODESIZE(rm.UM32.Load()) }
func (rm RMCODESIZE) Store(b CODESIZE) { rm.UM32.Store(uint32(b)) }

type CLENR0 uint32

func (b CLENR0) Field(mask CLENR0) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLENR0) J(v int) CLENR0 {
	return CLENR0(bits.Make32(v, uint32(mask)))
}

type RCLENR0 struct{ mmio.U32 }

func (r *RCLENR0) Bits(mask CLENR0) CLENR0  { return CLENR0(r.U32.Bits(uint32(mask))) }
func (r *RCLENR0) StoreBits(mask, b CLENR0) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCLENR0) SetBits(mask CLENR0)      { r.U32.SetBits(uint32(mask)) }
func (r *RCLENR0) ClearBits(mask CLENR0)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCLENR0) Load() CLENR0             { return CLENR0(r.U32.Load()) }
func (r *RCLENR0) Store(b CLENR0)           { r.U32.Store(uint32(b)) }

func (r *RCLENR0) AtomicStoreBits(mask, b CLENR0) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCLENR0) AtomicSetBits(mask CLENR0)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCLENR0) AtomicClearBits(mask CLENR0)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCLENR0 struct{ mmio.UM32 }

func (rm RMCLENR0) Load() CLENR0   { return CLENR0(rm.UM32.Load()) }
func (rm RMCLENR0) Store(b CLENR0) { rm.UM32.Store(uint32(b)) }

type PPFC uint32

func (b PPFC) Field(mask PPFC) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PPFC) J(v int) PPFC {
	return PPFC(bits.Make32(v, uint32(mask)))
}

type RPPFC struct{ mmio.U32 }

func (r *RPPFC) Bits(mask PPFC) PPFC    { return PPFC(r.U32.Bits(uint32(mask))) }
func (r *RPPFC) StoreBits(mask, b PPFC) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPPFC) SetBits(mask PPFC)      { r.U32.SetBits(uint32(mask)) }
func (r *RPPFC) ClearBits(mask PPFC)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPPFC) Load() PPFC             { return PPFC(r.U32.Load()) }
func (r *RPPFC) Store(b PPFC)           { r.U32.Store(uint32(b)) }

func (r *RPPFC) AtomicStoreBits(mask, b PPFC) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPPFC) AtomicSetBits(mask PPFC)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPPFC) AtomicClearBits(mask PPFC)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPPFC struct{ mmio.UM32 }

func (rm RMPPFC) Load() PPFC   { return PPFC(rm.UM32.Load()) }
func (rm RMPPFC) Store(b PPFC) { rm.UM32.Store(uint32(b)) }

type NUMRAMBLOCK uint32

func (b NUMRAMBLOCK) Field(mask NUMRAMBLOCK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NUMRAMBLOCK) J(v int) NUMRAMBLOCK {
	return NUMRAMBLOCK(bits.Make32(v, uint32(mask)))
}

type RNUMRAMBLOCK struct{ mmio.U32 }

func (r *RNUMRAMBLOCK) Bits(mask NUMRAMBLOCK) NUMRAMBLOCK {
	return NUMRAMBLOCK(r.U32.Bits(uint32(mask)))
}
func (r *RNUMRAMBLOCK) StoreBits(mask, b NUMRAMBLOCK) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RNUMRAMBLOCK) SetBits(mask NUMRAMBLOCK)      { r.U32.SetBits(uint32(mask)) }
func (r *RNUMRAMBLOCK) ClearBits(mask NUMRAMBLOCK)    { r.U32.ClearBits(uint32(mask)) }
func (r *RNUMRAMBLOCK) Load() NUMRAMBLOCK             { return NUMRAMBLOCK(r.U32.Load()) }
func (r *RNUMRAMBLOCK) Store(b NUMRAMBLOCK)           { r.U32.Store(uint32(b)) }

func (r *RNUMRAMBLOCK) AtomicStoreBits(mask, b NUMRAMBLOCK) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RNUMRAMBLOCK) AtomicSetBits(mask NUMRAMBLOCK)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RNUMRAMBLOCK) AtomicClearBits(mask NUMRAMBLOCK) { r.U32.AtomicClearBits(uint32(mask)) }

type RMNUMRAMBLOCK struct{ mmio.UM32 }

func (rm RMNUMRAMBLOCK) Load() NUMRAMBLOCK   { return NUMRAMBLOCK(rm.UM32.Load()) }
func (rm RMNUMRAMBLOCK) Store(b NUMRAMBLOCK) { rm.UM32.Store(uint32(b)) }

type SIZERAMBLOCK uint32

func (b SIZERAMBLOCK) Field(mask SIZERAMBLOCK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SIZERAMBLOCK) J(v int) SIZERAMBLOCK {
	return SIZERAMBLOCK(bits.Make32(v, uint32(mask)))
}

type RSIZERAMBLOCK struct{ mmio.U32 }

func (r *RSIZERAMBLOCK) Bits(mask SIZERAMBLOCK) SIZERAMBLOCK {
	return SIZERAMBLOCK(r.U32.Bits(uint32(mask)))
}
func (r *RSIZERAMBLOCK) StoreBits(mask, b SIZERAMBLOCK) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSIZERAMBLOCK) SetBits(mask SIZERAMBLOCK)      { r.U32.SetBits(uint32(mask)) }
func (r *RSIZERAMBLOCK) ClearBits(mask SIZERAMBLOCK)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSIZERAMBLOCK) Load() SIZERAMBLOCK             { return SIZERAMBLOCK(r.U32.Load()) }
func (r *RSIZERAMBLOCK) Store(b SIZERAMBLOCK)           { r.U32.Store(uint32(b)) }

func (r *RSIZERAMBLOCK) AtomicStoreBits(mask, b SIZERAMBLOCK) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RSIZERAMBLOCK) AtomicSetBits(mask SIZERAMBLOCK)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSIZERAMBLOCK) AtomicClearBits(mask SIZERAMBLOCK) { r.U32.AtomicClearBits(uint32(mask)) }

type RMSIZERAMBLOCK struct{ mmio.UM32 }

func (rm RMSIZERAMBLOCK) Load() SIZERAMBLOCK   { return SIZERAMBLOCK(rm.UM32.Load()) }
func (rm RMSIZERAMBLOCK) Store(b SIZERAMBLOCK) { rm.UM32.Store(uint32(b)) }

type CONFIGID uint32

func (b CONFIGID) Field(mask CONFIGID) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CONFIGID) J(v int) CONFIGID {
	return CONFIGID(bits.Make32(v, uint32(mask)))
}

type RCONFIGID struct{ mmio.U32 }

func (r *RCONFIGID) Bits(mask CONFIGID) CONFIGID { return CONFIGID(r.U32.Bits(uint32(mask))) }
func (r *RCONFIGID) StoreBits(mask, b CONFIGID)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCONFIGID) SetBits(mask CONFIGID)       { r.U32.SetBits(uint32(mask)) }
func (r *RCONFIGID) ClearBits(mask CONFIGID)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCONFIGID) Load() CONFIGID              { return CONFIGID(r.U32.Load()) }
func (r *RCONFIGID) Store(b CONFIGID)            { r.U32.Store(uint32(b)) }

func (r *RCONFIGID) AtomicStoreBits(mask, b CONFIGID) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCONFIGID) AtomicSetBits(mask CONFIGID)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCONFIGID) AtomicClearBits(mask CONFIGID)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCONFIGID struct{ mmio.UM32 }

func (rm RMCONFIGID) Load() CONFIGID   { return CONFIGID(rm.UM32.Load()) }
func (rm RMCONFIGID) Store(b CONFIGID) { rm.UM32.Store(uint32(b)) }

type DEVICEID uint32

func (b DEVICEID) Field(mask DEVICEID) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEVICEID) J(v int) DEVICEID {
	return DEVICEID(bits.Make32(v, uint32(mask)))
}

type RDEVICEID struct{ mmio.U32 }

func (r *RDEVICEID) Bits(mask DEVICEID) DEVICEID { return DEVICEID(r.U32.Bits(uint32(mask))) }
func (r *RDEVICEID) StoreBits(mask, b DEVICEID)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDEVICEID) SetBits(mask DEVICEID)       { r.U32.SetBits(uint32(mask)) }
func (r *RDEVICEID) ClearBits(mask DEVICEID)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDEVICEID) Load() DEVICEID              { return DEVICEID(r.U32.Load()) }
func (r *RDEVICEID) Store(b DEVICEID)            { r.U32.Store(uint32(b)) }

func (r *RDEVICEID) AtomicStoreBits(mask, b DEVICEID) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDEVICEID) AtomicSetBits(mask DEVICEID)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDEVICEID) AtomicClearBits(mask DEVICEID)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDEVICEID struct{ mmio.UM32 }

func (rm RMDEVICEID) Load() DEVICEID   { return DEVICEID(rm.UM32.Load()) }
func (rm RMDEVICEID) Store(b DEVICEID) { rm.UM32.Store(uint32(b)) }

type ER uint32

func (b ER) Field(mask ER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ER) J(v int) ER {
	return ER(bits.Make32(v, uint32(mask)))
}

type RER struct{ mmio.U32 }

func (r *RER) Bits(mask ER) ER      { return ER(r.U32.Bits(uint32(mask))) }
func (r *RER) StoreBits(mask, b ER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RER) SetBits(mask ER)      { r.U32.SetBits(uint32(mask)) }
func (r *RER) ClearBits(mask ER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RER) Load() ER             { return ER(r.U32.Load()) }
func (r *RER) Store(b ER)           { r.U32.Store(uint32(b)) }

func (r *RER) AtomicStoreBits(mask, b ER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RER) AtomicSetBits(mask ER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RER) AtomicClearBits(mask ER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMER struct{ mmio.UM32 }

func (rm RMER) Load() ER   { return ER(rm.UM32.Load()) }
func (rm RMER) Store(b ER) { rm.UM32.Store(uint32(b)) }

type IR uint32

func (b IR) Field(mask IR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IR) J(v int) IR {
	return IR(bits.Make32(v, uint32(mask)))
}

type RIR struct{ mmio.U32 }

func (r *RIR) Bits(mask IR) IR      { return IR(r.U32.Bits(uint32(mask))) }
func (r *RIR) StoreBits(mask, b IR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIR) SetBits(mask IR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIR) ClearBits(mask IR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIR) Load() IR             { return IR(r.U32.Load()) }
func (r *RIR) Store(b IR)           { r.U32.Store(uint32(b)) }

func (r *RIR) AtomicStoreBits(mask, b IR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIR) AtomicSetBits(mask IR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIR) AtomicClearBits(mask IR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIR struct{ mmio.UM32 }

func (rm RMIR) Load() IR   { return IR(rm.UM32.Load()) }
func (rm RMIR) Store(b IR) { rm.UM32.Store(uint32(b)) }

type DEVICEADDRTYPE uint32

func (b DEVICEADDRTYPE) Field(mask DEVICEADDRTYPE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEVICEADDRTYPE) J(v int) DEVICEADDRTYPE {
	return DEVICEADDRTYPE(bits.Make32(v, uint32(mask)))
}

type RDEVICEADDRTYPE struct{ mmio.U32 }

func (r *RDEVICEADDRTYPE) Bits(mask DEVICEADDRTYPE) DEVICEADDRTYPE {
	return DEVICEADDRTYPE(r.U32.Bits(uint32(mask)))
}
func (r *RDEVICEADDRTYPE) StoreBits(mask, b DEVICEADDRTYPE) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDEVICEADDRTYPE) SetBits(mask DEVICEADDRTYPE)      { r.U32.SetBits(uint32(mask)) }
func (r *RDEVICEADDRTYPE) ClearBits(mask DEVICEADDRTYPE)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDEVICEADDRTYPE) Load() DEVICEADDRTYPE             { return DEVICEADDRTYPE(r.U32.Load()) }
func (r *RDEVICEADDRTYPE) Store(b DEVICEADDRTYPE)           { r.U32.Store(uint32(b)) }

func (r *RDEVICEADDRTYPE) AtomicStoreBits(mask, b DEVICEADDRTYPE) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RDEVICEADDRTYPE) AtomicSetBits(mask DEVICEADDRTYPE)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDEVICEADDRTYPE) AtomicClearBits(mask DEVICEADDRTYPE) { r.U32.AtomicClearBits(uint32(mask)) }

type RMDEVICEADDRTYPE struct{ mmio.UM32 }

func (rm RMDEVICEADDRTYPE) Load() DEVICEADDRTYPE   { return DEVICEADDRTYPE(rm.UM32.Load()) }
func (rm RMDEVICEADDRTYPE) Store(b DEVICEADDRTYPE) { rm.UM32.Store(uint32(b)) }

type DEVICEADDR uint32

func (b DEVICEADDR) Field(mask DEVICEADDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEVICEADDR) J(v int) DEVICEADDR {
	return DEVICEADDR(bits.Make32(v, uint32(mask)))
}

type RDEVICEADDR struct{ mmio.U32 }

func (r *RDEVICEADDR) Bits(mask DEVICEADDR) DEVICEADDR { return DEVICEADDR(r.U32.Bits(uint32(mask))) }
func (r *RDEVICEADDR) StoreBits(mask, b DEVICEADDR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDEVICEADDR) SetBits(mask DEVICEADDR)         { r.U32.SetBits(uint32(mask)) }
func (r *RDEVICEADDR) ClearBits(mask DEVICEADDR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RDEVICEADDR) Load() DEVICEADDR                { return DEVICEADDR(r.U32.Load()) }
func (r *RDEVICEADDR) Store(b DEVICEADDR)              { r.U32.Store(uint32(b)) }

func (r *RDEVICEADDR) AtomicStoreBits(mask, b DEVICEADDR) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RDEVICEADDR) AtomicSetBits(mask DEVICEADDR)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDEVICEADDR) AtomicClearBits(mask DEVICEADDR) { r.U32.AtomicClearBits(uint32(mask)) }

type RMDEVICEADDR struct{ mmio.UM32 }

func (rm RMDEVICEADDR) Load() DEVICEADDR   { return DEVICEADDR(rm.UM32.Load()) }
func (rm RMDEVICEADDR) Store(b DEVICEADDR) { rm.UM32.Store(uint32(b)) }

type OVERRIDDEN uint32

func (b OVERRIDDEN) Field(mask OVERRIDDEN) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OVERRIDDEN) J(v int) OVERRIDDEN {
	return OVERRIDDEN(bits.Make32(v, uint32(mask)))
}

type ROVERRIDDEN struct{ mmio.U32 }

func (r *ROVERRIDDEN) Bits(mask OVERRIDDEN) OVERRIDDEN { return OVERRIDDEN(r.U32.Bits(uint32(mask))) }
func (r *ROVERRIDDEN) StoreBits(mask, b OVERRIDDEN)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROVERRIDDEN) SetBits(mask OVERRIDDEN)         { r.U32.SetBits(uint32(mask)) }
func (r *ROVERRIDDEN) ClearBits(mask OVERRIDDEN)       { r.U32.ClearBits(uint32(mask)) }
func (r *ROVERRIDDEN) Load() OVERRIDDEN                { return OVERRIDDEN(r.U32.Load()) }
func (r *ROVERRIDDEN) Store(b OVERRIDDEN)              { r.U32.Store(uint32(b)) }

func (r *ROVERRIDDEN) AtomicStoreBits(mask, b OVERRIDDEN) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *ROVERRIDDEN) AtomicSetBits(mask OVERRIDDEN)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROVERRIDDEN) AtomicClearBits(mask OVERRIDDEN) { r.U32.AtomicClearBits(uint32(mask)) }

type RMOVERRIDDEN struct{ mmio.UM32 }

func (rm RMOVERRIDDEN) Load() OVERRIDDEN   { return OVERRIDDEN(rm.UM32.Load()) }
func (rm RMOVERRIDDEN) Store(b OVERRIDDEN) { rm.UM32.Store(uint32(b)) }

func (p *Periph) NRF_1MBIT_OK() RMOVERRIDDEN {
	return RMOVERRIDDEN{mmio.UM32{&p.OVERRIDDEN.U32, uint32(NRF_1MBIT_OK)}}
}

func (p *Periph) BLE_1MBIT_OK() RMOVERRIDDEN {
	return RMOVERRIDDEN{mmio.UM32{&p.OVERRIDDEN.U32, uint32(BLE_1MBIT_OK)}}
}

type NRF_1MBIT uint32

func (b NRF_1MBIT) Field(mask NRF_1MBIT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NRF_1MBIT) J(v int) NRF_1MBIT {
	return NRF_1MBIT(bits.Make32(v, uint32(mask)))
}

type RNRF_1MBIT struct{ mmio.U32 }

func (r *RNRF_1MBIT) Bits(mask NRF_1MBIT) NRF_1MBIT { return NRF_1MBIT(r.U32.Bits(uint32(mask))) }
func (r *RNRF_1MBIT) StoreBits(mask, b NRF_1MBIT)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RNRF_1MBIT) SetBits(mask NRF_1MBIT)        { r.U32.SetBits(uint32(mask)) }
func (r *RNRF_1MBIT) ClearBits(mask NRF_1MBIT)      { r.U32.ClearBits(uint32(mask)) }
func (r *RNRF_1MBIT) Load() NRF_1MBIT               { return NRF_1MBIT(r.U32.Load()) }
func (r *RNRF_1MBIT) Store(b NRF_1MBIT)             { r.U32.Store(uint32(b)) }

func (r *RNRF_1MBIT) AtomicStoreBits(mask, b NRF_1MBIT) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RNRF_1MBIT) AtomicSetBits(mask NRF_1MBIT)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RNRF_1MBIT) AtomicClearBits(mask NRF_1MBIT) { r.U32.AtomicClearBits(uint32(mask)) }

type RMNRF_1MBIT struct{ mmio.UM32 }

func (rm RMNRF_1MBIT) Load() NRF_1MBIT   { return NRF_1MBIT(rm.UM32.Load()) }
func (rm RMNRF_1MBIT) Store(b NRF_1MBIT) { rm.UM32.Store(uint32(b)) }

type BLE_1MBIT uint32

func (b BLE_1MBIT) Field(mask BLE_1MBIT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BLE_1MBIT) J(v int) BLE_1MBIT {
	return BLE_1MBIT(bits.Make32(v, uint32(mask)))
}

type RBLE_1MBIT struct{ mmio.U32 }

func (r *RBLE_1MBIT) Bits(mask BLE_1MBIT) BLE_1MBIT { return BLE_1MBIT(r.U32.Bits(uint32(mask))) }
func (r *RBLE_1MBIT) StoreBits(mask, b BLE_1MBIT)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBLE_1MBIT) SetBits(mask BLE_1MBIT)        { r.U32.SetBits(uint32(mask)) }
func (r *RBLE_1MBIT) ClearBits(mask BLE_1MBIT)      { r.U32.ClearBits(uint32(mask)) }
func (r *RBLE_1MBIT) Load() BLE_1MBIT               { return BLE_1MBIT(r.U32.Load()) }
func (r *RBLE_1MBIT) Store(b BLE_1MBIT)             { r.U32.Store(uint32(b)) }

func (r *RBLE_1MBIT) AtomicStoreBits(mask, b BLE_1MBIT) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RBLE_1MBIT) AtomicSetBits(mask BLE_1MBIT)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBLE_1MBIT) AtomicClearBits(mask BLE_1MBIT) { r.U32.AtomicClearBits(uint32(mask)) }

type RMBLE_1MBIT struct{ mmio.UM32 }

func (rm RMBLE_1MBIT) Load() BLE_1MBIT   { return BLE_1MBIT(rm.UM32.Load()) }
func (rm RMBLE_1MBIT) Store(b BLE_1MBIT) { rm.UM32.Store(uint32(b)) }

type INFO_PART uint32

func (b INFO_PART) Field(mask INFO_PART) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_PART) J(v int) INFO_PART {
	return INFO_PART(bits.Make32(v, uint32(mask)))
}

type RINFO_PART struct{ mmio.U32 }

func (r *RINFO_PART) Bits(mask INFO_PART) INFO_PART { return INFO_PART(r.U32.Bits(uint32(mask))) }
func (r *RINFO_PART) StoreBits(mask, b INFO_PART)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RINFO_PART) SetBits(mask INFO_PART)        { r.U32.SetBits(uint32(mask)) }
func (r *RINFO_PART) ClearBits(mask INFO_PART)      { r.U32.ClearBits(uint32(mask)) }
func (r *RINFO_PART) Load() INFO_PART               { return INFO_PART(r.U32.Load()) }
func (r *RINFO_PART) Store(b INFO_PART)             { r.U32.Store(uint32(b)) }

func (r *RINFO_PART) AtomicStoreBits(mask, b INFO_PART) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RINFO_PART) AtomicSetBits(mask INFO_PART)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RINFO_PART) AtomicClearBits(mask INFO_PART) { r.U32.AtomicClearBits(uint32(mask)) }

type RMINFO_PART struct{ mmio.UM32 }

func (rm RMINFO_PART) Load() INFO_PART   { return INFO_PART(rm.UM32.Load()) }
func (rm RMINFO_PART) Store(b INFO_PART) { rm.UM32.Store(uint32(b)) }

type INFO_VARIANT uint32

func (b INFO_VARIANT) Field(mask INFO_VARIANT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_VARIANT) J(v int) INFO_VARIANT {
	return INFO_VARIANT(bits.Make32(v, uint32(mask)))
}

type RINFO_VARIANT struct{ mmio.U32 }

func (r *RINFO_VARIANT) Bits(mask INFO_VARIANT) INFO_VARIANT {
	return INFO_VARIANT(r.U32.Bits(uint32(mask)))
}
func (r *RINFO_VARIANT) StoreBits(mask, b INFO_VARIANT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RINFO_VARIANT) SetBits(mask INFO_VARIANT)      { r.U32.SetBits(uint32(mask)) }
func (r *RINFO_VARIANT) ClearBits(mask INFO_VARIANT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RINFO_VARIANT) Load() INFO_VARIANT             { return INFO_VARIANT(r.U32.Load()) }
func (r *RINFO_VARIANT) Store(b INFO_VARIANT)           { r.U32.Store(uint32(b)) }

func (r *RINFO_VARIANT) AtomicStoreBits(mask, b INFO_VARIANT) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RINFO_VARIANT) AtomicSetBits(mask INFO_VARIANT)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RINFO_VARIANT) AtomicClearBits(mask INFO_VARIANT) { r.U32.AtomicClearBits(uint32(mask)) }

type RMINFO_VARIANT struct{ mmio.UM32 }

func (rm RMINFO_VARIANT) Load() INFO_VARIANT   { return INFO_VARIANT(rm.UM32.Load()) }
func (rm RMINFO_VARIANT) Store(b INFO_VARIANT) { rm.UM32.Store(uint32(b)) }

type INFO_PACKAGE uint32

func (b INFO_PACKAGE) Field(mask INFO_PACKAGE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_PACKAGE) J(v int) INFO_PACKAGE {
	return INFO_PACKAGE(bits.Make32(v, uint32(mask)))
}

type RINFO_PACKAGE struct{ mmio.U32 }

func (r *RINFO_PACKAGE) Bits(mask INFO_PACKAGE) INFO_PACKAGE {
	return INFO_PACKAGE(r.U32.Bits(uint32(mask)))
}
func (r *RINFO_PACKAGE) StoreBits(mask, b INFO_PACKAGE) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RINFO_PACKAGE) SetBits(mask INFO_PACKAGE)      { r.U32.SetBits(uint32(mask)) }
func (r *RINFO_PACKAGE) ClearBits(mask INFO_PACKAGE)    { r.U32.ClearBits(uint32(mask)) }
func (r *RINFO_PACKAGE) Load() INFO_PACKAGE             { return INFO_PACKAGE(r.U32.Load()) }
func (r *RINFO_PACKAGE) Store(b INFO_PACKAGE)           { r.U32.Store(uint32(b)) }

func (r *RINFO_PACKAGE) AtomicStoreBits(mask, b INFO_PACKAGE) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RINFO_PACKAGE) AtomicSetBits(mask INFO_PACKAGE)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RINFO_PACKAGE) AtomicClearBits(mask INFO_PACKAGE) { r.U32.AtomicClearBits(uint32(mask)) }

type RMINFO_PACKAGE struct{ mmio.UM32 }

func (rm RMINFO_PACKAGE) Load() INFO_PACKAGE   { return INFO_PACKAGE(rm.UM32.Load()) }
func (rm RMINFO_PACKAGE) Store(b INFO_PACKAGE) { rm.UM32.Store(uint32(b)) }

type INFO_RAM uint32

func (b INFO_RAM) Field(mask INFO_RAM) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_RAM) J(v int) INFO_RAM {
	return INFO_RAM(bits.Make32(v, uint32(mask)))
}

type RINFO_RAM struct{ mmio.U32 }

func (r *RINFO_RAM) Bits(mask INFO_RAM) INFO_RAM { return INFO_RAM(r.U32.Bits(uint32(mask))) }
func (r *RINFO_RAM) StoreBits(mask, b INFO_RAM)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RINFO_RAM) SetBits(mask INFO_RAM)       { r.U32.SetBits(uint32(mask)) }
func (r *RINFO_RAM) ClearBits(mask INFO_RAM)     { r.U32.ClearBits(uint32(mask)) }
func (r *RINFO_RAM) Load() INFO_RAM              { return INFO_RAM(r.U32.Load()) }
func (r *RINFO_RAM) Store(b INFO_RAM)            { r.U32.Store(uint32(b)) }

func (r *RINFO_RAM) AtomicStoreBits(mask, b INFO_RAM) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RINFO_RAM) AtomicSetBits(mask INFO_RAM)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RINFO_RAM) AtomicClearBits(mask INFO_RAM)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMINFO_RAM struct{ mmio.UM32 }

func (rm RMINFO_RAM) Load() INFO_RAM   { return INFO_RAM(rm.UM32.Load()) }
func (rm RMINFO_RAM) Store(b INFO_RAM) { rm.UM32.Store(uint32(b)) }

type INFO_FLASH uint32

func (b INFO_FLASH) Field(mask INFO_FLASH) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_FLASH) J(v int) INFO_FLASH {
	return INFO_FLASH(bits.Make32(v, uint32(mask)))
}

type RINFO_FLASH struct{ mmio.U32 }

func (r *RINFO_FLASH) Bits(mask INFO_FLASH) INFO_FLASH { return INFO_FLASH(r.U32.Bits(uint32(mask))) }
func (r *RINFO_FLASH) StoreBits(mask, b INFO_FLASH)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RINFO_FLASH) SetBits(mask INFO_FLASH)         { r.U32.SetBits(uint32(mask)) }
func (r *RINFO_FLASH) ClearBits(mask INFO_FLASH)       { r.U32.ClearBits(uint32(mask)) }
func (r *RINFO_FLASH) Load() INFO_FLASH                { return INFO_FLASH(r.U32.Load()) }
func (r *RINFO_FLASH) Store(b INFO_FLASH)              { r.U32.Store(uint32(b)) }

func (r *RINFO_FLASH) AtomicStoreBits(mask, b INFO_FLASH) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RINFO_FLASH) AtomicSetBits(mask INFO_FLASH)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RINFO_FLASH) AtomicClearBits(mask INFO_FLASH) { r.U32.AtomicClearBits(uint32(mask)) }

type RMINFO_FLASH struct{ mmio.UM32 }

func (rm RMINFO_FLASH) Load() INFO_FLASH   { return INFO_FLASH(rm.UM32.Load()) }
func (rm RMINFO_FLASH) Store(b INFO_FLASH) { rm.UM32.Store(uint32(b)) }

type TEMP_A uint32

func (b TEMP_A) Field(mask TEMP_A) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TEMP_A) J(v int) TEMP_A {
	return TEMP_A(bits.Make32(v, uint32(mask)))
}

type RTEMP_A struct{ mmio.U32 }

func (r *RTEMP_A) Bits(mask TEMP_A) TEMP_A  { return TEMP_A(r.U32.Bits(uint32(mask))) }
func (r *RTEMP_A) StoreBits(mask, b TEMP_A) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTEMP_A) SetBits(mask TEMP_A)      { r.U32.SetBits(uint32(mask)) }
func (r *RTEMP_A) ClearBits(mask TEMP_A)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTEMP_A) Load() TEMP_A             { return TEMP_A(r.U32.Load()) }
func (r *RTEMP_A) Store(b TEMP_A)           { r.U32.Store(uint32(b)) }

func (r *RTEMP_A) AtomicStoreBits(mask, b TEMP_A) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTEMP_A) AtomicSetBits(mask TEMP_A)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTEMP_A) AtomicClearBits(mask TEMP_A)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTEMP_A struct{ mmio.UM32 }

func (rm RMTEMP_A) Load() TEMP_A   { return TEMP_A(rm.UM32.Load()) }
func (rm RMTEMP_A) Store(b TEMP_A) { rm.UM32.Store(uint32(b)) }

type TEMP_B uint32

func (b TEMP_B) Field(mask TEMP_B) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TEMP_B) J(v int) TEMP_B {
	return TEMP_B(bits.Make32(v, uint32(mask)))
}

type RTEMP_B struct{ mmio.U32 }

func (r *RTEMP_B) Bits(mask TEMP_B) TEMP_B  { return TEMP_B(r.U32.Bits(uint32(mask))) }
func (r *RTEMP_B) StoreBits(mask, b TEMP_B) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTEMP_B) SetBits(mask TEMP_B)      { r.U32.SetBits(uint32(mask)) }
func (r *RTEMP_B) ClearBits(mask TEMP_B)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTEMP_B) Load() TEMP_B             { return TEMP_B(r.U32.Load()) }
func (r *RTEMP_B) Store(b TEMP_B)           { r.U32.Store(uint32(b)) }

func (r *RTEMP_B) AtomicStoreBits(mask, b TEMP_B) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTEMP_B) AtomicSetBits(mask TEMP_B)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTEMP_B) AtomicClearBits(mask TEMP_B)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTEMP_B struct{ mmio.UM32 }

func (rm RMTEMP_B) Load() TEMP_B   { return TEMP_B(rm.UM32.Load()) }
func (rm RMTEMP_B) Store(b TEMP_B) { rm.UM32.Store(uint32(b)) }

type TEMP_T uint32

func (b TEMP_T) Field(mask TEMP_T) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TEMP_T) J(v int) TEMP_T {
	return TEMP_T(bits.Make32(v, uint32(mask)))
}

type RTEMP_T struct{ mmio.U32 }

func (r *RTEMP_T) Bits(mask TEMP_T) TEMP_T  { return TEMP_T(r.U32.Bits(uint32(mask))) }
func (r *RTEMP_T) StoreBits(mask, b TEMP_T) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTEMP_T) SetBits(mask TEMP_T)      { r.U32.SetBits(uint32(mask)) }
func (r *RTEMP_T) ClearBits(mask TEMP_T)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTEMP_T) Load() TEMP_T             { return TEMP_T(r.U32.Load()) }
func (r *RTEMP_T) Store(b TEMP_T)           { r.U32.Store(uint32(b)) }

func (r *RTEMP_T) AtomicStoreBits(mask, b TEMP_T) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTEMP_T) AtomicSetBits(mask TEMP_T)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTEMP_T) AtomicClearBits(mask TEMP_T)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTEMP_T struct{ mmio.UM32 }

func (rm RMTEMP_T) Load() TEMP_T   { return TEMP_T(rm.UM32.Load()) }
func (rm RMTEMP_T) Store(b TEMP_T) { rm.UM32.Store(uint32(b)) }

type NFC_TAGHEADER uint32

func (b NFC_TAGHEADER) Field(mask NFC_TAGHEADER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NFC_TAGHEADER) J(v int) NFC_TAGHEADER {
	return NFC_TAGHEADER(bits.Make32(v, uint32(mask)))
}

type RNFC_TAGHEADER struct{ mmio.U32 }

func (r *RNFC_TAGHEADER) Bits(mask NFC_TAGHEADER) NFC_TAGHEADER {
	return NFC_TAGHEADER(r.U32.Bits(uint32(mask)))
}
func (r *RNFC_TAGHEADER) StoreBits(mask, b NFC_TAGHEADER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RNFC_TAGHEADER) SetBits(mask NFC_TAGHEADER)      { r.U32.SetBits(uint32(mask)) }
func (r *RNFC_TAGHEADER) ClearBits(mask NFC_TAGHEADER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RNFC_TAGHEADER) Load() NFC_TAGHEADER             { return NFC_TAGHEADER(r.U32.Load()) }
func (r *RNFC_TAGHEADER) Store(b NFC_TAGHEADER)           { r.U32.Store(uint32(b)) }

func (r *RNFC_TAGHEADER) AtomicStoreBits(mask, b NFC_TAGHEADER) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RNFC_TAGHEADER) AtomicSetBits(mask NFC_TAGHEADER)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RNFC_TAGHEADER) AtomicClearBits(mask NFC_TAGHEADER) { r.U32.AtomicClearBits(uint32(mask)) }

type RMNFC_TAGHEADER struct{ mmio.UM32 }

func (rm RMNFC_TAGHEADER) Load() NFC_TAGHEADER   { return NFC_TAGHEADER(rm.UM32.Load()) }
func (rm RMNFC_TAGHEADER) Store(b NFC_TAGHEADER) { rm.UM32.Store(uint32(b)) }
