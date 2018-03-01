// +build f411xe

// Peripheral: CRYP_Periph  Crypto Processor.
// Instances:
//  CRYP  mmap.CRYP_BASE
// Registers:
//  0x00 32  CR         Control register.
//  0x04 32  SR         Status register.
//  0x08 32  DR         Data input register.
//  0x0C 32  DOUT       Data output register.
//  0x10 32  DMACR      DMA control register.
//  0x14 32  IMSCR      Interrupt mask set/clear register.
//  0x18 32  RISR       Raw interrupt status register.
//  0x1C 32  MISR       Masked interrupt status register.
//  0x20 32  K0LR       Key left  register 0.
//  0x24 32  K0RR       Key right register 0.
//  0x28 32  K1LR       Key left  register 1.
//  0x2C 32  K1RR       Key right register 1.
//  0x30 32  K2LR       Key left  register 2.
//  0x34 32  K2RR       Key right register 2.
//  0x38 32  K3LR       Key left  register 3.
//  0x3C 32  K3RR       Key right register 3.
//  0x40 32  IV0LR      Initialization vector left-word  register 0.
//  0x44 32  IV0RR      Initialization vector right-word register 0.
//  0x48 32  IV1LR      Initialization vector left-word  register 1.
//  0x4C 32  IV1RR      Initialization vector right-word register 1.
//  0x50 32  CSGCMCCM0R GCM/GMAC or CCM/CMAC context swap register 0.
//  0x54 32  CSGCMCCM1R GCM/GMAC or CCM/CMAC context swap register 1.
//  0x58 32  CSGCMCCM2R GCM/GMAC or CCM/CMAC context swap register 2.
//  0x5C 32  CSGCMCCM3R GCM/GMAC or CCM/CMAC context swap register 3.
//  0x60 32  CSGCMCCM4R GCM/GMAC or CCM/CMAC context swap register 4.
//  0x64 32  CSGCMCCM5R GCM/GMAC or CCM/CMAC context swap register 5.
//  0x68 32  CSGCMCCM6R GCM/GMAC or CCM/CMAC context swap register 6.
//  0x6C 32  CSGCMCCM7R GCM/GMAC or CCM/CMAC context swap register 7.
//  0x70 32  CSGCM0R    GCM/GMAC context swap register 0.
//  0x74 32  CSGCM1R    GCM/GMAC context swap register 1.
//  0x78 32  CSGCM2R    GCM/GMAC context swap register 2.
//  0x7C 32  CSGCM3R    GCM/GMAC context swap register 3.
//  0x80 32  CSGCM4R    GCM/GMAC context swap register 4.
//  0x84 32  CSGCM5R    GCM/GMAC context swap register 5.
//  0x88 32  CSGCM6R    GCM/GMAC context swap register 6.
//  0x8C 32  CSGCM7R    GCM/GMAC context swap register 7.
// Import:
//  stm32/o/f411xe/mmap
package cryp

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	ALGODIR           CR = 0x01 << 2    //+
	ALGOMODE          CR = 0x10007 << 3 //+
	ALGOMODE_0        CR = 0x01 << 3
	ALGOMODE_1        CR = 0x02 << 3
	ALGOMODE_2        CR = 0x04 << 3
	ALGOMODE_TDES_ECB CR = 0x00 << 3
	ALGOMODE_TDES_CBC CR = 0x01 << 3
	ALGOMODE_DES_ECB  CR = 0x02 << 3
	ALGOMODE_DES_CBC  CR = 0x03 << 3
	ALGOMODE_AES_ECB  CR = 0x04 << 3
	ALGOMODE_AES_CBC  CR = 0x05 << 3
	ALGOMODE_AES_CTR  CR = 0x06 << 3
	ALGOMODE_AES_KEY  CR = 0x07 << 3
	DATATYPE          CR = 0x03 << 6 //+
	DATATYPE_0        CR = 0x01 << 6
	DATATYPE_1        CR = 0x02 << 6
	KEYSIZE           CR = 0x03 << 8 //+
	KEYSIZE_0         CR = 0x01 << 8
	KEYSIZE_1         CR = 0x02 << 8
	FFLUSH            CR = 0x01 << 14 //+
	CRYPEN            CR = 0x01 << 15 //+
	GCM_CCMPH         CR = 0x03 << 16 //+
	GCM_CCMPH_0       CR = 0x01 << 16
	GCM_CCMPH_1       CR = 0x02 << 16
	ALGOMODE_3        CR = 0x10000 << 3
)

const (
	ALGODIRn   = 2
	ALGOMODEn  = 3
	DATATYPEn  = 6
	KEYSIZEn   = 8
	FFLUSHn    = 14
	CRYPENn    = 15
	GCM_CCMPHn = 16
)

const (
	IFEM SR = 0x01 << 0 //+
	IFNF SR = 0x01 << 1 //+
	OFNE SR = 0x01 << 2 //+
	OFFU SR = 0x01 << 3 //+
	BUSY SR = 0x01 << 4 //+
)

const (
	IFEMn = 0
	IFNFn = 1
	OFNEn = 2
	OFFUn = 3
	BUSYn = 4
)

const (
	DIEN DMACR = 0x01 << 0 //+
	DOEN DMACR = 0x01 << 1 //+
)

const (
	DIENn = 0
	DOENn = 1
)

const (
	INIM  IMSCR = 0x01 << 0 //+
	OUTIM IMSCR = 0x01 << 1 //+
)

const (
	INIMn  = 0
	OUTIMn = 1
)

const (
	OUTRIS RISR = 0x01 << 0 //+
	INRIS  RISR = 0x01 << 1 //+
)

const (
	OUTRISn = 0
	INRISn  = 1
)

const (
	INMIS  MISR = 0x01 << 0 //+
	OUTMIS MISR = 0x01 << 1 //+
)

const (
	INMISn  = 0
	OUTMISn = 1
)
