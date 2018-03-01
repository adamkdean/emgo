// +build f40_41xxx

package dcmi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type DCMI_Periph struct {
	CR      RCR
	SR      RSR
	RISR    RRISR
	IER     RIER
	MISR    RMISR
	ICR     RICR
	ESCR    RESCR
	ESUR    RESUR
	CWSTRTR RCWSTRTR
	CWSIZER RCWSIZER
	DR      RDR
}

func (p *DCMI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DCMI = (*DCMI_Periph)(unsafe.Pointer(uintptr(mmap.DCMI_BASE)))

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

func (p *DCMI_Periph) CAPTURE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CAPTURE)}}
}

func (p *DCMI_Periph) CM() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CM)}}
}

func (p *DCMI_Periph) CROP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CROP)}}
}

func (p *DCMI_Periph) JPEG() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(JPEG)}}
}

func (p *DCMI_Periph) ESS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ESS)}}
}

func (p *DCMI_Periph) PCKPOL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PCKPOL)}}
}

func (p *DCMI_Periph) HSPOL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSPOL)}}
}

func (p *DCMI_Periph) VSPOL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(VSPOL)}}
}

func (p *DCMI_Periph) FCRC_0() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FCRC_0)}}
}

func (p *DCMI_Periph) FCRC_1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FCRC_1)}}
}

func (p *DCMI_Periph) EDM_0() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EDM_0)}}
}

func (p *DCMI_Periph) EDM_1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EDM_1)}}
}

func (p *DCMI_Periph) CRE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CRE)}}
}

func (p *DCMI_Periph) ENABLE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ENABLE)}}
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

func (p *DCMI_Periph) HSYNC() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(HSYNC)}}
}

func (p *DCMI_Periph) VSYNC() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(VSYNC)}}
}

func (p *DCMI_Periph) FNE() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FNE)}}
}

type RISR uint32

func (b RISR) Field(mask RISR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RISR) J(v int) RISR {
	return RISR(bits.Make32(v, uint32(mask)))
}

type RRISR struct{ mmio.U32 }

func (r *RRISR) Bits(mask RISR) RISR    { return RISR(r.U32.Bits(uint32(mask))) }
func (r *RRISR) StoreBits(mask, b RISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRISR) SetBits(mask RISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRISR) ClearBits(mask RISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRISR) Load() RISR             { return RISR(r.U32.Load()) }
func (r *RRISR) Store(b RISR)           { r.U32.Store(uint32(b)) }

func (r *RRISR) AtomicStoreBits(mask, b RISR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRISR) AtomicSetBits(mask RISR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRISR) AtomicClearBits(mask RISR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRISR struct{ mmio.UM32 }

func (rm RMRISR) Load() RISR   { return RISR(rm.UM32.Load()) }
func (rm RMRISR) Store(b RISR) { rm.UM32.Store(uint32(b)) }

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

func (p *DCMI_Periph) FRAME_IE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(FRAME_IE)}}
}

func (p *DCMI_Periph) OVR_IE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(OVR_IE)}}
}

func (p *DCMI_Periph) ERR_IE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(ERR_IE)}}
}

func (p *DCMI_Periph) VSYNC_IE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(VSYNC_IE)}}
}

func (p *DCMI_Periph) LINE_IE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(LINE_IE)}}
}

type MISR uint32

func (b MISR) Field(mask MISR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MISR) J(v int) MISR {
	return MISR(bits.Make32(v, uint32(mask)))
}

type RMISR struct{ mmio.U32 }

func (r *RMISR) Bits(mask MISR) MISR    { return MISR(r.U32.Bits(uint32(mask))) }
func (r *RMISR) StoreBits(mask, b MISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMISR) SetBits(mask MISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RMISR) ClearBits(mask MISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMISR) Load() MISR             { return MISR(r.U32.Load()) }
func (r *RMISR) Store(b MISR)           { r.U32.Store(uint32(b)) }

func (r *RMISR) AtomicStoreBits(mask, b MISR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMISR) AtomicSetBits(mask MISR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMISR) AtomicClearBits(mask MISR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMISR struct{ mmio.UM32 }

func (rm RMMISR) Load() MISR   { return MISR(rm.UM32.Load()) }
func (rm RMMISR) Store(b MISR) { rm.UM32.Store(uint32(b)) }

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

func (p *DCMI_Periph) FRAME_ISC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(FRAME_ISC)}}
}

func (p *DCMI_Periph) OVR_ISC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(OVR_ISC)}}
}

func (p *DCMI_Periph) ERR_ISC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(ERR_ISC)}}
}

func (p *DCMI_Periph) VSYNC_ISC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(VSYNC_ISC)}}
}

func (p *DCMI_Periph) LINE_ISC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(LINE_ISC)}}
}

type ESCR uint32

func (b ESCR) Field(mask ESCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ESCR) J(v int) ESCR {
	return ESCR(bits.Make32(v, uint32(mask)))
}

type RESCR struct{ mmio.U32 }

func (r *RESCR) Bits(mask ESCR) ESCR    { return ESCR(r.U32.Bits(uint32(mask))) }
func (r *RESCR) StoreBits(mask, b ESCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RESCR) SetBits(mask ESCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RESCR) ClearBits(mask ESCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RESCR) Load() ESCR             { return ESCR(r.U32.Load()) }
func (r *RESCR) Store(b ESCR)           { r.U32.Store(uint32(b)) }

func (r *RESCR) AtomicStoreBits(mask, b ESCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RESCR) AtomicSetBits(mask ESCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RESCR) AtomicClearBits(mask ESCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMESCR struct{ mmio.UM32 }

func (rm RMESCR) Load() ESCR   { return ESCR(rm.UM32.Load()) }
func (rm RMESCR) Store(b ESCR) { rm.UM32.Store(uint32(b)) }

func (p *DCMI_Periph) FSC() RMESCR {
	return RMESCR{mmio.UM32{&p.ESCR.U32, uint32(FSC)}}
}

func (p *DCMI_Periph) LSC() RMESCR {
	return RMESCR{mmio.UM32{&p.ESCR.U32, uint32(LSC)}}
}

func (p *DCMI_Periph) LEC() RMESCR {
	return RMESCR{mmio.UM32{&p.ESCR.U32, uint32(LEC)}}
}

func (p *DCMI_Periph) FEC() RMESCR {
	return RMESCR{mmio.UM32{&p.ESCR.U32, uint32(FEC)}}
}

type ESUR uint32

func (b ESUR) Field(mask ESUR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ESUR) J(v int) ESUR {
	return ESUR(bits.Make32(v, uint32(mask)))
}

type RESUR struct{ mmio.U32 }

func (r *RESUR) Bits(mask ESUR) ESUR    { return ESUR(r.U32.Bits(uint32(mask))) }
func (r *RESUR) StoreBits(mask, b ESUR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RESUR) SetBits(mask ESUR)      { r.U32.SetBits(uint32(mask)) }
func (r *RESUR) ClearBits(mask ESUR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RESUR) Load() ESUR             { return ESUR(r.U32.Load()) }
func (r *RESUR) Store(b ESUR)           { r.U32.Store(uint32(b)) }

func (r *RESUR) AtomicStoreBits(mask, b ESUR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RESUR) AtomicSetBits(mask ESUR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RESUR) AtomicClearBits(mask ESUR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMESUR struct{ mmio.UM32 }

func (rm RMESUR) Load() ESUR   { return ESUR(rm.UM32.Load()) }
func (rm RMESUR) Store(b ESUR) { rm.UM32.Store(uint32(b)) }

func (p *DCMI_Periph) FSU() RMESUR {
	return RMESUR{mmio.UM32{&p.ESUR.U32, uint32(FSU)}}
}

func (p *DCMI_Periph) LSU() RMESUR {
	return RMESUR{mmio.UM32{&p.ESUR.U32, uint32(LSU)}}
}

func (p *DCMI_Periph) LEU() RMESUR {
	return RMESUR{mmio.UM32{&p.ESUR.U32, uint32(LEU)}}
}

func (p *DCMI_Periph) FEU() RMESUR {
	return RMESUR{mmio.UM32{&p.ESUR.U32, uint32(FEU)}}
}

type CWSTRTR uint32

func (b CWSTRTR) Field(mask CWSTRTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CWSTRTR) J(v int) CWSTRTR {
	return CWSTRTR(bits.Make32(v, uint32(mask)))
}

type RCWSTRTR struct{ mmio.U32 }

func (r *RCWSTRTR) Bits(mask CWSTRTR) CWSTRTR { return CWSTRTR(r.U32.Bits(uint32(mask))) }
func (r *RCWSTRTR) StoreBits(mask, b CWSTRTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCWSTRTR) SetBits(mask CWSTRTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCWSTRTR) ClearBits(mask CWSTRTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCWSTRTR) Load() CWSTRTR             { return CWSTRTR(r.U32.Load()) }
func (r *RCWSTRTR) Store(b CWSTRTR)           { r.U32.Store(uint32(b)) }

func (r *RCWSTRTR) AtomicStoreBits(mask, b CWSTRTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCWSTRTR) AtomicSetBits(mask CWSTRTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCWSTRTR) AtomicClearBits(mask CWSTRTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCWSTRTR struct{ mmio.UM32 }

func (rm RMCWSTRTR) Load() CWSTRTR   { return CWSTRTR(rm.UM32.Load()) }
func (rm RMCWSTRTR) Store(b CWSTRTR) { rm.UM32.Store(uint32(b)) }

type CWSIZER uint32

func (b CWSIZER) Field(mask CWSIZER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CWSIZER) J(v int) CWSIZER {
	return CWSIZER(bits.Make32(v, uint32(mask)))
}

type RCWSIZER struct{ mmio.U32 }

func (r *RCWSIZER) Bits(mask CWSIZER) CWSIZER { return CWSIZER(r.U32.Bits(uint32(mask))) }
func (r *RCWSIZER) StoreBits(mask, b CWSIZER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCWSIZER) SetBits(mask CWSIZER)      { r.U32.SetBits(uint32(mask)) }
func (r *RCWSIZER) ClearBits(mask CWSIZER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCWSIZER) Load() CWSIZER             { return CWSIZER(r.U32.Load()) }
func (r *RCWSIZER) Store(b CWSIZER)           { r.U32.Store(uint32(b)) }

func (r *RCWSIZER) AtomicStoreBits(mask, b CWSIZER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCWSIZER) AtomicSetBits(mask CWSIZER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCWSIZER) AtomicClearBits(mask CWSIZER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCWSIZER struct{ mmio.UM32 }

func (rm RMCWSIZER) Load() CWSIZER   { return CWSIZER(rm.UM32.Load()) }
func (rm RMCWSIZER) Store(b CWSIZER) { rm.UM32.Store(uint32(b)) }

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

func (p *DCMI_Periph) BYTE0() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(BYTE0)}}
}

func (p *DCMI_Periph) BYTE1() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(BYTE1)}}
}

func (p *DCMI_Periph) BYTE2() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(BYTE2)}}
}

func (p *DCMI_Periph) BYTE3() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(BYTE3)}}
}
