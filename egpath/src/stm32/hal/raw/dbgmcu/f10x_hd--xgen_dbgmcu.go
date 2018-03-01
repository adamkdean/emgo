// +build f10x_hd

package dbgmcu

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type DBGMCU_Periph struct {
	IDCODE RIDCODE
	CR     RCR
}

func (p *DBGMCU_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DBGMCU = (*DBGMCU_Periph)(unsafe.Pointer(uintptr(mmap.DBGMCU_BASE)))

type IDCODE uint32

func (b IDCODE) Field(mask IDCODE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDCODE) J(v int) IDCODE {
	return IDCODE(bits.Make32(v, uint32(mask)))
}

type RIDCODE struct{ mmio.U32 }

func (r *RIDCODE) Bits(mask IDCODE) IDCODE  { return IDCODE(r.U32.Bits(uint32(mask))) }
func (r *RIDCODE) StoreBits(mask, b IDCODE) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIDCODE) SetBits(mask IDCODE)      { r.U32.SetBits(uint32(mask)) }
func (r *RIDCODE) ClearBits(mask IDCODE)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIDCODE) Load() IDCODE             { return IDCODE(r.U32.Load()) }
func (r *RIDCODE) Store(b IDCODE)           { r.U32.Store(uint32(b)) }

func (r *RIDCODE) AtomicStoreBits(mask, b IDCODE) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIDCODE) AtomicSetBits(mask IDCODE)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIDCODE) AtomicClearBits(mask IDCODE)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIDCODE struct{ mmio.UM32 }

func (rm RMIDCODE) Load() IDCODE   { return IDCODE(rm.UM32.Load()) }
func (rm RMIDCODE) Store(b IDCODE) { rm.UM32.Store(uint32(b)) }

func (p *DBGMCU_Periph) DEV_ID() RMIDCODE {
	return RMIDCODE{mmio.UM32{&p.IDCODE.U32, uint32(DEV_ID)}}
}

func (p *DBGMCU_Periph) REV_ID() RMIDCODE {
	return RMIDCODE{mmio.UM32{&p.IDCODE.U32, uint32(REV_ID)}}
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

func (p *DBGMCU_Periph) DBG_SLEEP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_SLEEP)}}
}

func (p *DBGMCU_Periph) DBG_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_STOP)}}
}

func (p *DBGMCU_Periph) DBG_STANDBY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_STANDBY)}}
}

func (p *DBGMCU_Periph) TRACE_IOEN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TRACE_IOEN)}}
}

func (p *DBGMCU_Periph) TRACE_MODE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TRACE_MODE)}}
}

func (p *DBGMCU_Periph) DBG_IWDG_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_IWDG_STOP)}}
}

func (p *DBGMCU_Periph) DBG_WWDG_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_WWDG_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM1_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM1_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM2_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM2_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM3_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM3_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM4_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM4_STOP)}}
}

func (p *DBGMCU_Periph) DBG_CAN1_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_CAN1_STOP)}}
}

func (p *DBGMCU_Periph) DBG_I2C1_SMBUS_TIMEOUT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_I2C1_SMBUS_TIMEOUT)}}
}

func (p *DBGMCU_Periph) DBG_I2C2_SMBUS_TIMEOUT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_I2C2_SMBUS_TIMEOUT)}}
}

func (p *DBGMCU_Periph) DBG_TIM8_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM8_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM5_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM5_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM6_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM6_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM7_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM7_STOP)}}
}

func (p *DBGMCU_Periph) DBG_CAN2_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_CAN2_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM15_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM15_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM16_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM16_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM17_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM17_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM12_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM12_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM13_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM13_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM14_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM14_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM9_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM9_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM10_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM10_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM11_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_TIM11_STOP)}}
}
