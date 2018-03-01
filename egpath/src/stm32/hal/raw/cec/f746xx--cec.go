// +build f746xx

// Peripheral: CEC_Periph  HDMI-CEC.
// Instances:
//  CEC  mmap.CEC_BASE
// Registers:
//  0x00 32  CR   Control register.
//  0x04 32  CFGR Configuration register.
//  0x08 32  TXDR Tx data register.
//  0x0C 32  RXDR Rx Data Register.
//  0x10 32  ISR  Interrupt and Status Register.
//  0x14 32  IER  Interrupt enable register.
// Import:
//  stm32/o/f746xx/mmap
package cec

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CECEN CR = 0x01 << 0 //+ CEC Enable.
	TXSOM CR = 0x01 << 1 //+ CEC Tx Start Of Message.
	TXEOM CR = 0x01 << 2 //+ CEC Tx End Of Message.
)

const (
	CECENn = 0
	TXSOMn = 1
	TXEOMn = 2
)

const (
	SFT      CFGR = 0x07 << 0    //+ CEC Signal Free Time.
	RXTOL    CFGR = 0x01 << 3    //+ CEC Tolerance.
	BRESTP   CFGR = 0x01 << 4    //+ CEC Rx Stop.
	BREGEN   CFGR = 0x01 << 5    //+ CEC Bit Rising Error generation.
	LBPEGEN  CFGR = 0x01 << 6    //+ CEC Long Period Error generation.
	BRDNOGEN CFGR = 0x01 << 7    //+ CEC Broadcast no Error generation.
	SFTOPT   CFGR = 0x01 << 8    //+ CEC Signal Free Time optional.
	OAR      CFGR = 0x7FFF << 16 //+ CEC Own Address.
	LSTN     CFGR = 0x01 << 31   //+ CEC Listen mode.
)

const (
	SFTn      = 0
	RXTOLn    = 3
	BRESTPn   = 4
	BREGENn   = 5
	LBPEGENn  = 6
	BRDNOGENn = 7
	SFTOPTn   = 8
	OARn      = 16
	LSTNn     = 31
)

const (
	TXD TXDR = 0xFF << 0 //+ CEC Tx Data.
	RXD TXDR = 0xFF << 0 //  CEC Rx Data.
)

const (
	TXDn = 0
)

const (
	RXBR   ISR = 0x01 << 0  //+ CEC Rx-Byte Received.
	RXEND  ISR = 0x01 << 1  //+ CEC End Of Reception.
	RXOVR  ISR = 0x01 << 2  //+ CEC Rx-Overrun.
	BRE    ISR = 0x01 << 3  //+ CEC Rx Bit Rising Error.
	SBPE   ISR = 0x01 << 4  //+ CEC Rx Short Bit period Error.
	LBPE   ISR = 0x01 << 5  //+ CEC Rx Long Bit period Error.
	RXACKE ISR = 0x01 << 6  //+ CEC Rx Missing Acknowledge.
	ARBLST ISR = 0x01 << 7  //+ CEC Arbitration Lost.
	TXBR   ISR = 0x01 << 8  //+ CEC Tx Byte Request.
	TXEND  ISR = 0x01 << 9  //+ CEC End of Transmission.
	TXUDR  ISR = 0x01 << 10 //+ CEC Tx-Buffer Underrun.
	TXERR  ISR = 0x01 << 11 //+ CEC Tx-Error.
	TXACKE ISR = 0x01 << 12 //+ CEC Tx Missing Acknowledge.
)

const (
	RXBRn   = 0
	RXENDn  = 1
	RXOVRn  = 2
	BREn    = 3
	SBPEn   = 4
	LBPEn   = 5
	RXACKEn = 6
	ARBLSTn = 7
	TXBRn   = 8
	TXENDn  = 9
	TXUDRn  = 10
	TXERRn  = 11
	TXACKEn = 12
)

const (
	RXBRIE   IER = 0x01 << 0  //+ CEC Rx-Byte Received IT Enable.
	RXENDIE  IER = 0x01 << 1  //+ CEC End Of Reception IT Enable.
	RXOVRIE  IER = 0x01 << 2  //+ CEC Rx-Overrun IT Enable.
	BREIE    IER = 0x01 << 3  //+ CEC Rx Bit Rising Error IT Enable.
	SBPEIE   IER = 0x01 << 4  //+ CEC Rx Short Bit period Error IT Enable.
	LBPEIE   IER = 0x01 << 5  //+ CEC Rx Long Bit period Error IT Enable.
	RXACKEIE IER = 0x01 << 6  //+ CEC Rx Missing Acknowledge IT Enable.
	ARBLSTIE IER = 0x01 << 7  //+ CEC Arbitration Lost IT Enable.
	TXBRIE   IER = 0x01 << 8  //+ CEC Tx Byte Request  IT Enable.
	TXENDIE  IER = 0x01 << 9  //+ CEC End of Transmission IT Enable.
	TXUDRIE  IER = 0x01 << 10 //+ CEC Tx-Buffer Underrun IT Enable.
	TXERRIE  IER = 0x01 << 11 //+ CEC Tx-Error IT Enable.
	TXACKEIE IER = 0x01 << 12 //+ CEC Tx Missing Acknowledge IT Enable.
)

const (
	RXBRIEn   = 0
	RXENDIEn  = 1
	RXOVRIEn  = 2
	BREIEn    = 3
	SBPEIEn   = 4
	LBPEIEn   = 5
	RXACKEIEn = 6
	ARBLSTIEn = 7
	TXBRIEn   = 8
	TXENDIEn  = 9
	TXUDRIEn  = 10
	TXERRIEn  = 11
	TXACKEIEn = 12
)
