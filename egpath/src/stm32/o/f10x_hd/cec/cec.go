// Peripheral: CEC_Periph  Consumer Electronics Control (CEC).
// Instances:
//  CEC  mmap.CEC_BASE
// Registers:
//  0x00 32  CFGR
//  0x04 32  OAR
//  0x08 32  PRES
//  0x0C 32  ESR
//  0x10 32  CSR
//  0x14 32  TXD
//  0x18 32  RXD
// Import:
//  stm32/o/f10x_hd/mmap
package cec

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PE   CFGR = 0x01 << 0 //+ Peripheral Enable.
	IE   CFGR = 0x01 << 1 //+ Interrupt Enable.
	BTEM CFGR = 0x01 << 2 //+ Bit Timing Error Mode.
	BPEM CFGR = 0x01 << 3 //+ Bit Period Error Mode.
)

const (
	PEn   = 0
	IEn   = 1
	BTEMn = 2
	BPEMn = 3
)

const (
	OA   OAR = 0x0F << 0 //+ OA[3:0]: Own Address.
	OA_0 OAR = 0x01 << 0 //  Bit 0.
	OA_1 OAR = 0x02 << 0 //  Bit 1.
	OA_2 OAR = 0x04 << 0 //  Bit 2.
	OA_3 OAR = 0x08 << 0 //  Bit 3.
)

const (
	OAn = 0
)

const (
	BTE   ESR = 0x01 << 0 //+ Bit Timing Error.
	BPE   ESR = 0x01 << 1 //+ Bit Period Error.
	RBTFE ESR = 0x01 << 2 //+ Rx Block Transfer Finished Error.
	SBE   ESR = 0x01 << 3 //+ Start Bit Error.
	ACKE  ESR = 0x01 << 4 //+ Block Acknowledge Error.
	LINE  ESR = 0x01 << 5 //+ Line Error.
	TBTFE ESR = 0x01 << 6 //+ Tx Block Transfer Finished Error.
)

const (
	BTEn   = 0
	BPEn   = 1
	RBTFEn = 2
	SBEn   = 3
	ACKEn  = 4
	LINEn  = 5
	TBTFEn = 6
)

const (
	TSOM  CSR = 0x01 << 0 //+ Tx Start Of Message.
	TEOM  CSR = 0x01 << 1 //+ Tx End Of Message.
	TERR  CSR = 0x01 << 2 //+ Tx Error.
	TBTRF CSR = 0x01 << 3 //+ Tx Byte Transfer Request or Block Transfer Finished.
	RSOM  CSR = 0x01 << 4 //+ Rx Start Of Message.
	REOM  CSR = 0x01 << 5 //+ Rx End Of Message.
	RERR  CSR = 0x01 << 6 //+ Rx Error.
	RBTF  CSR = 0x01 << 7 //+ Rx Block Transfer Finished.
)

const (
	TSOMn  = 0
	TEOMn  = 1
	TERRn  = 2
	TBTRFn = 3
	RSOMn  = 4
	REOMn  = 5
	RERRn  = 6
	RBTFn  = 7
)
