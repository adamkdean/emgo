package gpio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type GPIO_Periph struct {
	MODER   RMODER
	OTYPER  ROTYPER
	OSPEEDR ROSPEEDR
	PUPDR   RPUPDR
	IDR     RIDR
	ODR     RODR
	BSRRL   RBSRRL
	BSRRH   RBSRRH
	LCKR    RLCKR
	AFR     [2]RAFR
}

func (p *GPIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var GPIOA = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOA_BASE)))

//emgo:const
var GPIOB = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOB_BASE)))

//emgo:const
var GPIOC = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOC_BASE)))

//emgo:const
var GPIOD = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOD_BASE)))

//emgo:const
var GPIOE = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOE_BASE)))

//emgo:const
var GPIOF = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOF_BASE)))

//emgo:const
var GPIOG = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOG_BASE)))

//emgo:const
var GPIOH = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOH_BASE)))

//emgo:const
var GPIOI = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOI_BASE)))

//emgo:const
var GPIOJ = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOJ_BASE)))

//emgo:const
var GPIOK = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOK_BASE)))

type MODER uint32

func (b MODER) Field(mask MODER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MODER) J(v int) MODER {
	return MODER(bits.Make32(v, uint32(mask)))
}

type RMODER struct{ mmio.U32 }

func (r *RMODER) Bits(mask MODER) MODER   { return MODER(r.U32.Bits(uint32(mask))) }
func (r *RMODER) StoreBits(mask, b MODER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMODER) SetBits(mask MODER)      { r.U32.SetBits(uint32(mask)) }
func (r *RMODER) ClearBits(mask MODER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMODER) Load() MODER             { return MODER(r.U32.Load()) }
func (r *RMODER) Store(b MODER)           { r.U32.Store(uint32(b)) }

func (r *RMODER) AtomicStoreBits(mask, b MODER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMODER) AtomicSetBits(mask MODER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMODER) AtomicClearBits(mask MODER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMODER struct{ mmio.UM32 }

func (rm RMMODER) Load() MODER   { return MODER(rm.UM32.Load()) }
func (rm RMMODER) Store(b MODER) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) MODER0() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER0)}}
}

func (p *GPIO_Periph) MODER1() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER1)}}
}

func (p *GPIO_Periph) MODER2() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER2)}}
}

func (p *GPIO_Periph) MODER3() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER3)}}
}

func (p *GPIO_Periph) MODER4() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER4)}}
}

func (p *GPIO_Periph) MODER5() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER5)}}
}

func (p *GPIO_Periph) MODER6() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER6)}}
}

func (p *GPIO_Periph) MODER7() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER7)}}
}

func (p *GPIO_Periph) MODER8() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER8)}}
}

func (p *GPIO_Periph) MODER9() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER9)}}
}

func (p *GPIO_Periph) MODER10() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER10)}}
}

func (p *GPIO_Periph) MODER11() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER11)}}
}

func (p *GPIO_Periph) MODER12() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER12)}}
}

func (p *GPIO_Periph) MODER13() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER13)}}
}

func (p *GPIO_Periph) MODER14() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER14)}}
}

func (p *GPIO_Periph) MODER15() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER15)}}
}

type OTYPER uint32

func (b OTYPER) Field(mask OTYPER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OTYPER) J(v int) OTYPER {
	return OTYPER(bits.Make32(v, uint32(mask)))
}

type ROTYPER struct{ mmio.U32 }

func (r *ROTYPER) Bits(mask OTYPER) OTYPER  { return OTYPER(r.U32.Bits(uint32(mask))) }
func (r *ROTYPER) StoreBits(mask, b OTYPER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROTYPER) SetBits(mask OTYPER)      { r.U32.SetBits(uint32(mask)) }
func (r *ROTYPER) ClearBits(mask OTYPER)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROTYPER) Load() OTYPER             { return OTYPER(r.U32.Load()) }
func (r *ROTYPER) Store(b OTYPER)           { r.U32.Store(uint32(b)) }

func (r *ROTYPER) AtomicStoreBits(mask, b OTYPER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROTYPER) AtomicSetBits(mask OTYPER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROTYPER) AtomicClearBits(mask OTYPER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOTYPER struct{ mmio.UM32 }

func (rm RMOTYPER) Load() OTYPER   { return OTYPER(rm.UM32.Load()) }
func (rm RMOTYPER) Store(b OTYPER) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OT_0() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_0)}}
}

func (p *GPIO_Periph) OT_1() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_1)}}
}

func (p *GPIO_Periph) OT_2() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_2)}}
}

func (p *GPIO_Periph) OT_3() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_3)}}
}

func (p *GPIO_Periph) OT_4() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_4)}}
}

func (p *GPIO_Periph) OT_5() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_5)}}
}

func (p *GPIO_Periph) OT_6() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_6)}}
}

func (p *GPIO_Periph) OT_7() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_7)}}
}

func (p *GPIO_Periph) OT_8() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_8)}}
}

func (p *GPIO_Periph) OT_9() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_9)}}
}

func (p *GPIO_Periph) OT_10() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_10)}}
}

func (p *GPIO_Periph) OT_11() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_11)}}
}

func (p *GPIO_Periph) OT_12() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_12)}}
}

func (p *GPIO_Periph) OT_13() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_13)}}
}

func (p *GPIO_Periph) OT_14() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_14)}}
}

func (p *GPIO_Periph) OT_15() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_15)}}
}

type OSPEEDR uint32

func (b OSPEEDR) Field(mask OSPEEDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OSPEEDR) J(v int) OSPEEDR {
	return OSPEEDR(bits.Make32(v, uint32(mask)))
}

type ROSPEEDR struct{ mmio.U32 }

func (r *ROSPEEDR) Bits(mask OSPEEDR) OSPEEDR { return OSPEEDR(r.U32.Bits(uint32(mask))) }
func (r *ROSPEEDR) StoreBits(mask, b OSPEEDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROSPEEDR) SetBits(mask OSPEEDR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROSPEEDR) ClearBits(mask OSPEEDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROSPEEDR) Load() OSPEEDR             { return OSPEEDR(r.U32.Load()) }
func (r *ROSPEEDR) Store(b OSPEEDR)           { r.U32.Store(uint32(b)) }

func (r *ROSPEEDR) AtomicStoreBits(mask, b OSPEEDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROSPEEDR) AtomicSetBits(mask OSPEEDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROSPEEDR) AtomicClearBits(mask OSPEEDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOSPEEDR struct{ mmio.UM32 }

func (rm RMOSPEEDR) Load() OSPEEDR   { return OSPEEDR(rm.UM32.Load()) }
func (rm RMOSPEEDR) Store(b OSPEEDR) { rm.UM32.Store(uint32(b)) }

type PUPDR uint32

func (b PUPDR) Field(mask PUPDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PUPDR) J(v int) PUPDR {
	return PUPDR(bits.Make32(v, uint32(mask)))
}

type RPUPDR struct{ mmio.U32 }

func (r *RPUPDR) Bits(mask PUPDR) PUPDR   { return PUPDR(r.U32.Bits(uint32(mask))) }
func (r *RPUPDR) StoreBits(mask, b PUPDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPUPDR) SetBits(mask PUPDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPUPDR) ClearBits(mask PUPDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPUPDR) Load() PUPDR             { return PUPDR(r.U32.Load()) }
func (r *RPUPDR) Store(b PUPDR)           { r.U32.Store(uint32(b)) }

func (r *RPUPDR) AtomicStoreBits(mask, b PUPDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPUPDR) AtomicSetBits(mask PUPDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPUPDR) AtomicClearBits(mask PUPDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPUPDR struct{ mmio.UM32 }

func (rm RMPUPDR) Load() PUPDR   { return PUPDR(rm.UM32.Load()) }
func (rm RMPUPDR) Store(b PUPDR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) PUPDR0() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR0)}}
}

func (p *GPIO_Periph) PUPDR1() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR1)}}
}

func (p *GPIO_Periph) PUPDR2() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR2)}}
}

func (p *GPIO_Periph) PUPDR3() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR3)}}
}

func (p *GPIO_Periph) PUPDR4() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR4)}}
}

func (p *GPIO_Periph) PUPDR5() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR5)}}
}

func (p *GPIO_Periph) PUPDR6() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR6)}}
}

func (p *GPIO_Periph) PUPDR7() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR7)}}
}

func (p *GPIO_Periph) PUPDR8() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR8)}}
}

func (p *GPIO_Periph) PUPDR9() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR9)}}
}

func (p *GPIO_Periph) PUPDR10() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR10)}}
}

func (p *GPIO_Periph) PUPDR11() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR11)}}
}

func (p *GPIO_Periph) PUPDR12() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR12)}}
}

func (p *GPIO_Periph) PUPDR13() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR13)}}
}

func (p *GPIO_Periph) PUPDR14() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR14)}}
}

func (p *GPIO_Periph) PUPDR15() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR15)}}
}

type IDR uint32

func (b IDR) Field(mask IDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDR) J(v int) IDR {
	return IDR(bits.Make32(v, uint32(mask)))
}

type RIDR struct{ mmio.U32 }

func (r *RIDR) Bits(mask IDR) IDR     { return IDR(r.U32.Bits(uint32(mask))) }
func (r *RIDR) StoreBits(mask, b IDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIDR) SetBits(mask IDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIDR) ClearBits(mask IDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIDR) Load() IDR             { return IDR(r.U32.Load()) }
func (r *RIDR) Store(b IDR)           { r.U32.Store(uint32(b)) }

func (r *RIDR) AtomicStoreBits(mask, b IDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIDR) AtomicSetBits(mask IDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIDR) AtomicClearBits(mask IDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIDR struct{ mmio.UM32 }

func (rm RMIDR) Load() IDR   { return IDR(rm.UM32.Load()) }
func (rm RMIDR) Store(b IDR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) IDR_0() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_0)}}
}

func (p *GPIO_Periph) IDR_1() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_1)}}
}

func (p *GPIO_Periph) IDR_2() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_2)}}
}

func (p *GPIO_Periph) IDR_3() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_3)}}
}

func (p *GPIO_Periph) IDR_4() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_4)}}
}

func (p *GPIO_Periph) IDR_5() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_5)}}
}

func (p *GPIO_Periph) IDR_6() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_6)}}
}

func (p *GPIO_Periph) IDR_7() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_7)}}
}

func (p *GPIO_Periph) IDR_8() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_8)}}
}

func (p *GPIO_Periph) IDR_9() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_9)}}
}

func (p *GPIO_Periph) IDR_10() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_10)}}
}

func (p *GPIO_Periph) IDR_11() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_11)}}
}

func (p *GPIO_Periph) IDR_12() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_12)}}
}

func (p *GPIO_Periph) IDR_13() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_13)}}
}

func (p *GPIO_Periph) IDR_14() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_14)}}
}

func (p *GPIO_Periph) IDR_15() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR_15)}}
}

type ODR uint32

func (b ODR) Field(mask ODR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ODR) J(v int) ODR {
	return ODR(bits.Make32(v, uint32(mask)))
}

type RODR struct{ mmio.U32 }

func (r *RODR) Bits(mask ODR) ODR     { return ODR(r.U32.Bits(uint32(mask))) }
func (r *RODR) StoreBits(mask, b ODR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RODR) SetBits(mask ODR)      { r.U32.SetBits(uint32(mask)) }
func (r *RODR) ClearBits(mask ODR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RODR) Load() ODR             { return ODR(r.U32.Load()) }
func (r *RODR) Store(b ODR)           { r.U32.Store(uint32(b)) }

func (r *RODR) AtomicStoreBits(mask, b ODR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RODR) AtomicSetBits(mask ODR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RODR) AtomicClearBits(mask ODR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMODR struct{ mmio.UM32 }

func (rm RMODR) Load() ODR   { return ODR(rm.UM32.Load()) }
func (rm RMODR) Store(b ODR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) ODR_0() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_0)}}
}

func (p *GPIO_Periph) ODR_1() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_1)}}
}

func (p *GPIO_Periph) ODR_2() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_2)}}
}

func (p *GPIO_Periph) ODR_3() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_3)}}
}

func (p *GPIO_Periph) ODR_4() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_4)}}
}

func (p *GPIO_Periph) ODR_5() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_5)}}
}

func (p *GPIO_Periph) ODR_6() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_6)}}
}

func (p *GPIO_Periph) ODR_7() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_7)}}
}

func (p *GPIO_Periph) ODR_8() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_8)}}
}

func (p *GPIO_Periph) ODR_9() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_9)}}
}

func (p *GPIO_Periph) ODR_10() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_10)}}
}

func (p *GPIO_Periph) ODR_11() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_11)}}
}

func (p *GPIO_Periph) ODR_12() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_12)}}
}

func (p *GPIO_Periph) ODR_13() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_13)}}
}

func (p *GPIO_Periph) ODR_14() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_14)}}
}

func (p *GPIO_Periph) ODR_15() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR_15)}}
}

type BSRRL uint16

func (b BSRRL) Field(mask BSRRL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BSRRL) J(v int) BSRRL {
	return BSRRL(bits.Make32(v, uint32(mask)))
}

type RBSRRL struct{ mmio.U16 }

func (r *RBSRRL) Bits(mask BSRRL) BSRRL   { return BSRRL(r.U16.Bits(uint16(mask))) }
func (r *RBSRRL) StoreBits(mask, b BSRRL) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RBSRRL) SetBits(mask BSRRL)      { r.U16.SetBits(uint16(mask)) }
func (r *RBSRRL) ClearBits(mask BSRRL)    { r.U16.ClearBits(uint16(mask)) }
func (r *RBSRRL) Load() BSRRL             { return BSRRL(r.U16.Load()) }
func (r *RBSRRL) Store(b BSRRL)           { r.U16.Store(uint16(b)) }

type RMBSRRL struct{ mmio.UM16 }

func (rm RMBSRRL) Load() BSRRL   { return BSRRL(rm.UM16.Load()) }
func (rm RMBSRRL) Store(b BSRRL) { rm.UM16.Store(uint16(b)) }

type BSRRH uint16

func (b BSRRH) Field(mask BSRRH) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BSRRH) J(v int) BSRRH {
	return BSRRH(bits.Make32(v, uint32(mask)))
}

type RBSRRH struct{ mmio.U16 }

func (r *RBSRRH) Bits(mask BSRRH) BSRRH   { return BSRRH(r.U16.Bits(uint16(mask))) }
func (r *RBSRRH) StoreBits(mask, b BSRRH) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RBSRRH) SetBits(mask BSRRH)      { r.U16.SetBits(uint16(mask)) }
func (r *RBSRRH) ClearBits(mask BSRRH)    { r.U16.ClearBits(uint16(mask)) }
func (r *RBSRRH) Load() BSRRH             { return BSRRH(r.U16.Load()) }
func (r *RBSRRH) Store(b BSRRH)           { r.U16.Store(uint16(b)) }

type RMBSRRH struct{ mmio.UM16 }

func (rm RMBSRRH) Load() BSRRH   { return BSRRH(rm.UM16.Load()) }
func (rm RMBSRRH) Store(b BSRRH) { rm.UM16.Store(uint16(b)) }

type LCKR uint32

func (b LCKR) Field(mask LCKR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LCKR) J(v int) LCKR {
	return LCKR(bits.Make32(v, uint32(mask)))
}

type RLCKR struct{ mmio.U32 }

func (r *RLCKR) Bits(mask LCKR) LCKR    { return LCKR(r.U32.Bits(uint32(mask))) }
func (r *RLCKR) StoreBits(mask, b LCKR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLCKR) SetBits(mask LCKR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLCKR) ClearBits(mask LCKR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLCKR) Load() LCKR             { return LCKR(r.U32.Load()) }
func (r *RLCKR) Store(b LCKR)           { r.U32.Store(uint32(b)) }

func (r *RLCKR) AtomicStoreBits(mask, b LCKR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLCKR) AtomicSetBits(mask LCKR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLCKR) AtomicClearBits(mask LCKR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLCKR struct{ mmio.UM32 }

func (rm RMLCKR) Load() LCKR   { return LCKR(rm.UM32.Load()) }
func (rm RMLCKR) Store(b LCKR) { rm.UM32.Store(uint32(b)) }

type AFR uint32

func (b AFR) Field(mask AFR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AFR) J(v int) AFR {
	return AFR(bits.Make32(v, uint32(mask)))
}

type RAFR struct{ mmio.U32 }

func (r *RAFR) Bits(mask AFR) AFR     { return AFR(r.U32.Bits(uint32(mask))) }
func (r *RAFR) StoreBits(mask, b AFR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAFR) SetBits(mask AFR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAFR) ClearBits(mask AFR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAFR) Load() AFR             { return AFR(r.U32.Load()) }
func (r *RAFR) Store(b AFR)           { r.U32.Store(uint32(b)) }

func (r *RAFR) AtomicStoreBits(mask, b AFR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAFR) AtomicSetBits(mask AFR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAFR) AtomicClearBits(mask AFR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAFR struct{ mmio.UM32 }

func (rm RMAFR) Load() AFR   { return AFR(rm.UM32.Load()) }
func (rm RMAFR) Store(b AFR) { rm.UM32.Store(uint32(b)) }
