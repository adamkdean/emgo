// Peripheral: FSMC_Bank4_Periph  Flexible Static Memory Controller Bank4.
// Instances:
//  FSMC_Bank4  mmap.FSMC_Bank4_R_BASE
// Registers:
//  0x00 32  PCR4
//  0x04 32  SR4
//  0x08 32  PMEM4
//  0x0C 32  PATT4
//  0x10 32  PIO4
// Import:
//  stm32/o/f10x_hd/mmap
package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PWAITEN PCR4 = 0x01 << 1  //+ Wait feature enable bit.
	PBKEN   PCR4 = 0x01 << 2  //+ PC Card/NAND Flash memory bank enable bit.
	PTYP    PCR4 = 0x01 << 3  //+ Memory type.
	PWID    PCR4 = 0x03 << 4  //+ PWID[1:0] bits (NAND Flash databus width).
	PWID_0  PCR4 = 0x01 << 4  //  Bit 0.
	PWID_1  PCR4 = 0x02 << 4  //  Bit 1.
	ECCEN   PCR4 = 0x01 << 6  //+ ECC computation logic enable bit.
	TCLR    PCR4 = 0x0F << 9  //+ TCLR[3:0] bits (CLE to RE delay).
	TCLR_0  PCR4 = 0x01 << 9  //  Bit 0.
	TCLR_1  PCR4 = 0x02 << 9  //  Bit 1.
	TCLR_2  PCR4 = 0x04 << 9  //  Bit 2.
	TCLR_3  PCR4 = 0x08 << 9  //  Bit 3.
	TAR     PCR4 = 0x0F << 13 //+ TAR[3:0] bits (ALE to RE delay).
	TAR_0   PCR4 = 0x01 << 13 //  Bit 0.
	TAR_1   PCR4 = 0x02 << 13 //  Bit 1.
	TAR_2   PCR4 = 0x04 << 13 //  Bit 2.
	TAR_3   PCR4 = 0x08 << 13 //  Bit 3.
	ECCPS   PCR4 = 0x07 << 17 //+ ECCPS[2:0] bits (ECC page size).
	ECCPS_0 PCR4 = 0x01 << 17 //  Bit 0.
	ECCPS_1 PCR4 = 0x02 << 17 //  Bit 1.
	ECCPS_2 PCR4 = 0x04 << 17 //  Bit 2.
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
	IRS   SR4 = 0x01 << 0 //+ Interrupt Rising Edge status.
	ILS   SR4 = 0x01 << 1 //+ Interrupt Level status.
	IFS   SR4 = 0x01 << 2 //+ Interrupt Falling Edge status.
	IREN  SR4 = 0x01 << 3 //+ Interrupt Rising Edge detection Enable bit.
	ILEN  SR4 = 0x01 << 4 //+ Interrupt Level detection Enable bit.
	IFEN  SR4 = 0x01 << 5 //+ Interrupt Falling Edge detection Enable bit.
	FEMPT SR4 = 0x01 << 6 //+ FIFO empty.
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
	MEMSET4    PMEM4 = 0xFF << 0  //+ MEMSET4[7:0] bits (Common memory 4 setup time).
	MEMSET4_0  PMEM4 = 0x01 << 0  //  Bit 0.
	MEMSET4_1  PMEM4 = 0x02 << 0  //  Bit 1.
	MEMSET4_2  PMEM4 = 0x04 << 0  //  Bit 2.
	MEMSET4_3  PMEM4 = 0x08 << 0  //  Bit 3.
	MEMSET4_4  PMEM4 = 0x10 << 0  //  Bit 4.
	MEMSET4_5  PMEM4 = 0x20 << 0  //  Bit 5.
	MEMSET4_6  PMEM4 = 0x40 << 0  //  Bit 6.
	MEMSET4_7  PMEM4 = 0x80 << 0  //  Bit 7.
	MEMWAIT4   PMEM4 = 0xFF << 8  //+ MEMWAIT4[7:0] bits (Common memory 4 wait time).
	MEMWAIT4_0 PMEM4 = 0x01 << 8  //  Bit 0.
	MEMWAIT4_1 PMEM4 = 0x02 << 8  //  Bit 1.
	MEMWAIT4_2 PMEM4 = 0x04 << 8  //  Bit 2.
	MEMWAIT4_3 PMEM4 = 0x08 << 8  //  Bit 3.
	MEMWAIT4_4 PMEM4 = 0x10 << 8  //  Bit 4.
	MEMWAIT4_5 PMEM4 = 0x20 << 8  //  Bit 5.
	MEMWAIT4_6 PMEM4 = 0x40 << 8  //  Bit 6.
	MEMWAIT4_7 PMEM4 = 0x80 << 8  //  Bit 7.
	MEMHOLD4   PMEM4 = 0xFF << 16 //+ MEMHOLD4[7:0] bits (Common memory 4 hold time).
	MEMHOLD4_0 PMEM4 = 0x01 << 16 //  Bit 0.
	MEMHOLD4_1 PMEM4 = 0x02 << 16 //  Bit 1.
	MEMHOLD4_2 PMEM4 = 0x04 << 16 //  Bit 2.
	MEMHOLD4_3 PMEM4 = 0x08 << 16 //  Bit 3.
	MEMHOLD4_4 PMEM4 = 0x10 << 16 //  Bit 4.
	MEMHOLD4_5 PMEM4 = 0x20 << 16 //  Bit 5.
	MEMHOLD4_6 PMEM4 = 0x40 << 16 //  Bit 6.
	MEMHOLD4_7 PMEM4 = 0x80 << 16 //  Bit 7.
	MEMHIZ4    PMEM4 = 0xFF << 24 //+ MEMHIZ4[7:0] bits (Common memory 4 databus HiZ time).
	MEMHIZ4_0  PMEM4 = 0x01 << 24 //  Bit 0.
	MEMHIZ4_1  PMEM4 = 0x02 << 24 //  Bit 1.
	MEMHIZ4_2  PMEM4 = 0x04 << 24 //  Bit 2.
	MEMHIZ4_3  PMEM4 = 0x08 << 24 //  Bit 3.
	MEMHIZ4_4  PMEM4 = 0x10 << 24 //  Bit 4.
	MEMHIZ4_5  PMEM4 = 0x20 << 24 //  Bit 5.
	MEMHIZ4_6  PMEM4 = 0x40 << 24 //  Bit 6.
	MEMHIZ4_7  PMEM4 = 0x80 << 24 //  Bit 7.
)

const (
	MEMSET4n  = 0
	MEMWAIT4n = 8
	MEMHOLD4n = 16
	MEMHIZ4n  = 24
)

const (
	ATTSET4    PATT4 = 0xFF << 0  //+ ATTSET4[7:0] bits (Attribute memory 4 setup time).
	ATTSET4_0  PATT4 = 0x01 << 0  //  Bit 0.
	ATTSET4_1  PATT4 = 0x02 << 0  //  Bit 1.
	ATTSET4_2  PATT4 = 0x04 << 0  //  Bit 2.
	ATTSET4_3  PATT4 = 0x08 << 0  //  Bit 3.
	ATTSET4_4  PATT4 = 0x10 << 0  //  Bit 4.
	ATTSET4_5  PATT4 = 0x20 << 0  //  Bit 5.
	ATTSET4_6  PATT4 = 0x40 << 0  //  Bit 6.
	ATTSET4_7  PATT4 = 0x80 << 0  //  Bit 7.
	ATTWAIT4   PATT4 = 0xFF << 8  //+ ATTWAIT4[7:0] bits (Attribute memory 4 wait time).
	ATTWAIT4_0 PATT4 = 0x01 << 8  //  Bit 0.
	ATTWAIT4_1 PATT4 = 0x02 << 8  //  Bit 1.
	ATTWAIT4_2 PATT4 = 0x04 << 8  //  Bit 2.
	ATTWAIT4_3 PATT4 = 0x08 << 8  //  Bit 3.
	ATTWAIT4_4 PATT4 = 0x10 << 8  //  Bit 4.
	ATTWAIT4_5 PATT4 = 0x20 << 8  //  Bit 5.
	ATTWAIT4_6 PATT4 = 0x40 << 8  //  Bit 6.
	ATTWAIT4_7 PATT4 = 0x80 << 8  //  Bit 7.
	ATTHOLD4   PATT4 = 0xFF << 16 //+ ATTHOLD4[7:0] bits (Attribute memory 4 hold time).
	ATTHOLD4_0 PATT4 = 0x01 << 16 //  Bit 0.
	ATTHOLD4_1 PATT4 = 0x02 << 16 //  Bit 1.
	ATTHOLD4_2 PATT4 = 0x04 << 16 //  Bit 2.
	ATTHOLD4_3 PATT4 = 0x08 << 16 //  Bit 3.
	ATTHOLD4_4 PATT4 = 0x10 << 16 //  Bit 4.
	ATTHOLD4_5 PATT4 = 0x20 << 16 //  Bit 5.
	ATTHOLD4_6 PATT4 = 0x40 << 16 //  Bit 6.
	ATTHOLD4_7 PATT4 = 0x80 << 16 //  Bit 7.
	ATTHIZ4    PATT4 = 0xFF << 24 //+ ATTHIZ4[7:0] bits (Attribute memory 4 databus HiZ time).
	ATTHIZ4_0  PATT4 = 0x01 << 24 //  Bit 0.
	ATTHIZ4_1  PATT4 = 0x02 << 24 //  Bit 1.
	ATTHIZ4_2  PATT4 = 0x04 << 24 //  Bit 2.
	ATTHIZ4_3  PATT4 = 0x08 << 24 //  Bit 3.
	ATTHIZ4_4  PATT4 = 0x10 << 24 //  Bit 4.
	ATTHIZ4_5  PATT4 = 0x20 << 24 //  Bit 5.
	ATTHIZ4_6  PATT4 = 0x40 << 24 //  Bit 6.
	ATTHIZ4_7  PATT4 = 0x80 << 24 //  Bit 7.
)

const (
	ATTSET4n  = 0
	ATTWAIT4n = 8
	ATTHOLD4n = 16
	ATTHIZ4n  = 24
)

const (
	IOSET4    PIO4 = 0xFF << 0  //+ IOSET4[7:0] bits (I/O 4 setup time).
	IOSET4_0  PIO4 = 0x01 << 0  //  Bit 0.
	IOSET4_1  PIO4 = 0x02 << 0  //  Bit 1.
	IOSET4_2  PIO4 = 0x04 << 0  //  Bit 2.
	IOSET4_3  PIO4 = 0x08 << 0  //  Bit 3.
	IOSET4_4  PIO4 = 0x10 << 0  //  Bit 4.
	IOSET4_5  PIO4 = 0x20 << 0  //  Bit 5.
	IOSET4_6  PIO4 = 0x40 << 0  //  Bit 6.
	IOSET4_7  PIO4 = 0x80 << 0  //  Bit 7.
	IOWAIT4   PIO4 = 0xFF << 8  //+ IOWAIT4[7:0] bits (I/O 4 wait time).
	IOWAIT4_0 PIO4 = 0x01 << 8  //  Bit 0.
	IOWAIT4_1 PIO4 = 0x02 << 8  //  Bit 1.
	IOWAIT4_2 PIO4 = 0x04 << 8  //  Bit 2.
	IOWAIT4_3 PIO4 = 0x08 << 8  //  Bit 3.
	IOWAIT4_4 PIO4 = 0x10 << 8  //  Bit 4.
	IOWAIT4_5 PIO4 = 0x20 << 8  //  Bit 5.
	IOWAIT4_6 PIO4 = 0x40 << 8  //  Bit 6.
	IOWAIT4_7 PIO4 = 0x80 << 8  //  Bit 7.
	IOHOLD4   PIO4 = 0xFF << 16 //+ IOHOLD4[7:0] bits (I/O 4 hold time).
	IOHOLD4_0 PIO4 = 0x01 << 16 //  Bit 0.
	IOHOLD4_1 PIO4 = 0x02 << 16 //  Bit 1.
	IOHOLD4_2 PIO4 = 0x04 << 16 //  Bit 2.
	IOHOLD4_3 PIO4 = 0x08 << 16 //  Bit 3.
	IOHOLD4_4 PIO4 = 0x10 << 16 //  Bit 4.
	IOHOLD4_5 PIO4 = 0x20 << 16 //  Bit 5.
	IOHOLD4_6 PIO4 = 0x40 << 16 //  Bit 6.
	IOHOLD4_7 PIO4 = 0x80 << 16 //  Bit 7.
	IOHIZ4    PIO4 = 0xFF << 24 //+ IOHIZ4[7:0] bits (I/O 4 databus HiZ time).
	IOHIZ4_0  PIO4 = 0x01 << 24 //  Bit 0.
	IOHIZ4_1  PIO4 = 0x02 << 24 //  Bit 1.
	IOHIZ4_2  PIO4 = 0x04 << 24 //  Bit 2.
	IOHIZ4_3  PIO4 = 0x08 << 24 //  Bit 3.
	IOHIZ4_4  PIO4 = 0x10 << 24 //  Bit 4.
	IOHIZ4_5  PIO4 = 0x20 << 24 //  Bit 5.
	IOHIZ4_6  PIO4 = 0x40 << 24 //  Bit 6.
	IOHIZ4_7  PIO4 = 0x80 << 24 //  Bit 7.
)

const (
	IOSET4n  = 0
	IOWAIT4n = 8
	IOHOLD4n = 16
	IOHIZ4n  = 24
)
