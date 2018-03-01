// Peripheral: FSMC_Bank3_Periph  Flexible Static Memory Controller Bank3.
// Instances:
//  FSMC_Bank3  mmap.FSMC_Bank3_R_BASE
// Registers:
//  0x00 32  PCR3  NAND Flash control register 3.
//  0x04 32  SR3   NAND Flash FIFO status and interrupt register 3.
//  0x08 32  PMEM3 NAND Flash Common memory space timing register 3.
//  0x0C 32  PATT3 NAND Flash Attribute memory space timing register 3.
//  0x14 32  ECCR3 NAND Flash ECC result registers 3.
// Import:
//  stm32/o/f40_41xxx/mmap
package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PWAITEN PCR3 = 0x01 << 1  //+ Wait feature enable bit.
	PBKEN   PCR3 = 0x01 << 2  //+ PC Card/NAND Flash memory bank enable bit.
	PTYP    PCR3 = 0x01 << 3  //+ Memory type.
	PWID    PCR3 = 0x03 << 4  //+ PWID[1:0] bits (NAND Flash databus width).
	PWID_0  PCR3 = 0x01 << 4  //  Bit 0.
	PWID_1  PCR3 = 0x02 << 4  //  Bit 1.
	ECCEN   PCR3 = 0x01 << 6  //+ ECC computation logic enable bit.
	TCLR    PCR3 = 0x0F << 9  //+ TCLR[3:0] bits (CLE to RE delay).
	TCLR_0  PCR3 = 0x01 << 9  //  Bit 0.
	TCLR_1  PCR3 = 0x02 << 9  //  Bit 1.
	TCLR_2  PCR3 = 0x04 << 9  //  Bit 2.
	TCLR_3  PCR3 = 0x08 << 9  //  Bit 3.
	TAR     PCR3 = 0x0F << 13 //+ TAR[3:0] bits (ALE to RE delay).
	TAR_0   PCR3 = 0x01 << 13 //  Bit 0.
	TAR_1   PCR3 = 0x02 << 13 //  Bit 1.
	TAR_2   PCR3 = 0x04 << 13 //  Bit 2.
	TAR_3   PCR3 = 0x08 << 13 //  Bit 3.
	ECCPS   PCR3 = 0x07 << 17 //+ ECCPS[2:0] bits (ECC page size).
	ECCPS_0 PCR3 = 0x01 << 17 //  Bit 0.
	ECCPS_1 PCR3 = 0x02 << 17 //  Bit 1.
	ECCPS_2 PCR3 = 0x04 << 17 //  Bit 2.
)

const (
	PWAITENn = 1
	PBKENn   = 2
	PTYPn    = 3
	PWIDn    = 4
	ECCENn   = 6
	TCLRn    = 9
	TARn     = 13
	ECCPSn   = 17
)

const (
	IRS   SR3 = 0x01 << 0 //+ Interrupt Rising Edge status.
	ILS   SR3 = 0x01 << 1 //+ Interrupt Level status.
	IFS   SR3 = 0x01 << 2 //+ Interrupt Falling Edge status.
	IREN  SR3 = 0x01 << 3 //+ Interrupt Rising Edge detection Enable bit.
	ILEN  SR3 = 0x01 << 4 //+ Interrupt Level detection Enable bit.
	IFEN  SR3 = 0x01 << 5 //+ Interrupt Falling Edge detection Enable bit.
	FEMPT SR3 = 0x01 << 6 //+ FIFO empty.
)

const (
	IRSn   = 0
	ILSn   = 1
	IFSn   = 2
	IRENn  = 3
	ILENn  = 4
	IFENn  = 5
	FEMPTn = 6
)

const (
	MEMSET3    PMEM3 = 0xFF << 0  //+ MEMSET3[7:0] bits (Common memory 3 setup time).
	MEMSET3_0  PMEM3 = 0x01 << 0  //  Bit 0.
	MEMSET3_1  PMEM3 = 0x02 << 0  //  Bit 1.
	MEMSET3_2  PMEM3 = 0x04 << 0  //  Bit 2.
	MEMSET3_3  PMEM3 = 0x08 << 0  //  Bit 3.
	MEMSET3_4  PMEM3 = 0x10 << 0  //  Bit 4.
	MEMSET3_5  PMEM3 = 0x20 << 0  //  Bit 5.
	MEMSET3_6  PMEM3 = 0x40 << 0  //  Bit 6.
	MEMSET3_7  PMEM3 = 0x80 << 0  //  Bit 7.
	MEMWAIT3   PMEM3 = 0xFF << 8  //+ MEMWAIT3[7:0] bits (Common memory 3 wait time).
	MEMWAIT3_0 PMEM3 = 0x01 << 8  //  Bit 0.
	MEMWAIT3_1 PMEM3 = 0x02 << 8  //  Bit 1.
	MEMWAIT3_2 PMEM3 = 0x04 << 8  //  Bit 2.
	MEMWAIT3_3 PMEM3 = 0x08 << 8  //  Bit 3.
	MEMWAIT3_4 PMEM3 = 0x10 << 8  //  Bit 4.
	MEMWAIT3_5 PMEM3 = 0x20 << 8  //  Bit 5.
	MEMWAIT3_6 PMEM3 = 0x40 << 8  //  Bit 6.
	MEMWAIT3_7 PMEM3 = 0x80 << 8  //  Bit 7.
	MEMHOLD3   PMEM3 = 0xFF << 16 //+ MEMHOLD3[7:0] bits (Common memory 3 hold time).
	MEMHOLD3_0 PMEM3 = 0x01 << 16 //  Bit 0.
	MEMHOLD3_1 PMEM3 = 0x02 << 16 //  Bit 1.
	MEMHOLD3_2 PMEM3 = 0x04 << 16 //  Bit 2.
	MEMHOLD3_3 PMEM3 = 0x08 << 16 //  Bit 3.
	MEMHOLD3_4 PMEM3 = 0x10 << 16 //  Bit 4.
	MEMHOLD3_5 PMEM3 = 0x20 << 16 //  Bit 5.
	MEMHOLD3_6 PMEM3 = 0x40 << 16 //  Bit 6.
	MEMHOLD3_7 PMEM3 = 0x80 << 16 //  Bit 7.
	MEMHIZ3    PMEM3 = 0xFF << 24 //+ MEMHIZ3[7:0] bits (Common memory 3 databus HiZ time).
	MEMHIZ3_0  PMEM3 = 0x01 << 24 //  Bit 0.
	MEMHIZ3_1  PMEM3 = 0x02 << 24 //  Bit 1.
	MEMHIZ3_2  PMEM3 = 0x04 << 24 //  Bit 2.
	MEMHIZ3_3  PMEM3 = 0x08 << 24 //  Bit 3.
	MEMHIZ3_4  PMEM3 = 0x10 << 24 //  Bit 4.
	MEMHIZ3_5  PMEM3 = 0x20 << 24 //  Bit 5.
	MEMHIZ3_6  PMEM3 = 0x40 << 24 //  Bit 6.
	MEMHIZ3_7  PMEM3 = 0x80 << 24 //  Bit 7.
)

const (
	MEMSET3n  = 0
	MEMWAIT3n = 8
	MEMHOLD3n = 16
	MEMHIZ3n  = 24
)

const (
	ATTSET3    PATT3 = 0xFF << 0  //+ ATTSET3[7:0] bits (Attribute memory 3 setup time).
	ATTSET3_0  PATT3 = 0x01 << 0  //  Bit 0.
	ATTSET3_1  PATT3 = 0x02 << 0  //  Bit 1.
	ATTSET3_2  PATT3 = 0x04 << 0  //  Bit 2.
	ATTSET3_3  PATT3 = 0x08 << 0  //  Bit 3.
	ATTSET3_4  PATT3 = 0x10 << 0  //  Bit 4.
	ATTSET3_5  PATT3 = 0x20 << 0  //  Bit 5.
	ATTSET3_6  PATT3 = 0x40 << 0  //  Bit 6.
	ATTSET3_7  PATT3 = 0x80 << 0  //  Bit 7.
	ATTWAIT3   PATT3 = 0xFF << 8  //+ ATTWAIT3[7:0] bits (Attribute memory 3 wait time).
	ATTWAIT3_0 PATT3 = 0x01 << 8  //  Bit 0.
	ATTWAIT3_1 PATT3 = 0x02 << 8  //  Bit 1.
	ATTWAIT3_2 PATT3 = 0x04 << 8  //  Bit 2.
	ATTWAIT3_3 PATT3 = 0x08 << 8  //  Bit 3.
	ATTWAIT3_4 PATT3 = 0x10 << 8  //  Bit 4.
	ATTWAIT3_5 PATT3 = 0x20 << 8  //  Bit 5.
	ATTWAIT3_6 PATT3 = 0x40 << 8  //  Bit 6.
	ATTWAIT3_7 PATT3 = 0x80 << 8  //  Bit 7.
	ATTHOLD3   PATT3 = 0xFF << 16 //+ ATTHOLD3[7:0] bits (Attribute memory 3 hold time).
	ATTHOLD3_0 PATT3 = 0x01 << 16 //  Bit 0.
	ATTHOLD3_1 PATT3 = 0x02 << 16 //  Bit 1.
	ATTHOLD3_2 PATT3 = 0x04 << 16 //  Bit 2.
	ATTHOLD3_3 PATT3 = 0x08 << 16 //  Bit 3.
	ATTHOLD3_4 PATT3 = 0x10 << 16 //  Bit 4.
	ATTHOLD3_5 PATT3 = 0x20 << 16 //  Bit 5.
	ATTHOLD3_6 PATT3 = 0x40 << 16 //  Bit 6.
	ATTHOLD3_7 PATT3 = 0x80 << 16 //  Bit 7.
	ATTHIZ3    PATT3 = 0xFF << 24 //+ ATTHIZ3[7:0] bits (Attribute memory 3 databus HiZ time).
	ATTHIZ3_0  PATT3 = 0x01 << 24 //  Bit 0.
	ATTHIZ3_1  PATT3 = 0x02 << 24 //  Bit 1.
	ATTHIZ3_2  PATT3 = 0x04 << 24 //  Bit 2.
	ATTHIZ3_3  PATT3 = 0x08 << 24 //  Bit 3.
	ATTHIZ3_4  PATT3 = 0x10 << 24 //  Bit 4.
	ATTHIZ3_5  PATT3 = 0x20 << 24 //  Bit 5.
	ATTHIZ3_6  PATT3 = 0x40 << 24 //  Bit 6.
	ATTHIZ3_7  PATT3 = 0x80 << 24 //  Bit 7.
)

const (
	ATTSET3n  = 0
	ATTWAIT3n = 8
	ATTHOLD3n = 16
	ATTHIZ3n  = 24
)

const (
	ECC3 ECCR3 = 0xFFFFFFFF << 0 //+ ECC result.
)

const (
	ECC3n = 0
)
