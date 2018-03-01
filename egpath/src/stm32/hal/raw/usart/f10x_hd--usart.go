// +build f10x_hd

// Peripheral: USART_Periph  Universal Synchronous Asynchronous Receiver Transmitter.
// Instances:
//  USART2  mmap.USART2_BASE
//  USART3  mmap.USART3_BASE
//  UART4   mmap.UART4_BASE
//  UART5   mmap.UART5_BASE
//  USART1  mmap.USART1_BASE
// Registers:
//  0x00 16  SR
//  0x04 16  DR
//  0x08 16  BRR
//  0x0C 16  CR1
//  0x10 16  CR2
//  0x14 16  CR3
//  0x18 16  GTPR
// Import:
//  stm32/o/f10x_hd/mmap
package usart

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PE   SR = 0x01 << 0 //+ Parity Error.
	FE   SR = 0x01 << 1 //+ Framing Error.
	NE   SR = 0x01 << 2 //+ Noise Error Flag.
	ORE  SR = 0x01 << 3 //+ OverRun Error.
	IDLE SR = 0x01 << 4 //+ IDLE line detected.
	RXNE SR = 0x01 << 5 //+ Read Data Register Not Empty.
	TC   SR = 0x01 << 6 //+ Transmission Complete.
	TXE  SR = 0x01 << 7 //+ Transmit Data Register Empty.
	LBD  SR = 0x01 << 8 //+ LIN Break Detection Flag.
	CTS  SR = 0x01 << 9 //+ CTS Flag.
)

const (
	PEn   = 0
	FEn   = 1
	NEn   = 2
	OREn  = 3
	IDLEn = 4
	RXNEn = 5
	TCn   = 6
	TXEn  = 7
	LBDn  = 8
	CTSn  = 9
)

const (
	DIV_Fraction BRR = 0x0F << 0  //+ Fraction of USARTDIV.
	DIV_Mantissa BRR = 0xFFF << 4 //+ Mantissa of USARTDIV.
)

const (
	DIV_Fractionn = 0
	DIV_Mantissan = 4
)

const (
	SBK    CR1 = 0x01 << 0  //+ Send Break.
	RWU    CR1 = 0x01 << 1  //+ Receiver wakeup.
	RE     CR1 = 0x01 << 2  //+ Receiver Enable.
	TE     CR1 = 0x01 << 3  //+ Transmitter Enable.
	IDLEIE CR1 = 0x01 << 4  //+ IDLE Interrupt Enable.
	RXNEIE CR1 = 0x01 << 5  //+ RXNE Interrupt Enable.
	TCIE   CR1 = 0x01 << 6  //+ Transmission Complete Interrupt Enable.
	TXEIE  CR1 = 0x01 << 7  //+ PE Interrupt Enable.
	PEIE   CR1 = 0x01 << 8  //+ PE Interrupt Enable.
	PS     CR1 = 0x01 << 9  //+ Parity Selection.
	PCE    CR1 = 0x01 << 10 //+ Parity Control Enable.
	WAKE   CR1 = 0x01 << 11 //+ Wakeup method.
	M      CR1 = 0x01 << 12 //+ Word length.
	UE     CR1 = 0x01 << 13 //+ USART Enable.
	OVER8  CR1 = 0x01 << 15 //+ USART Oversmapling 8-bits.
)

const (
	SBKn    = 0
	RWUn    = 1
	REn     = 2
	TEn     = 3
	IDLEIEn = 4
	RXNEIEn = 5
	TCIEn   = 6
	TXEIEn  = 7
	PEIEn   = 8
	PSn     = 9
	PCEn    = 10
	WAKEn   = 11
	Mn      = 12
	UEn     = 13
	OVER8n  = 15
)

const (
	ADD    CR2 = 0x0F << 0  //+ Address of the USART node.
	LBDL   CR2 = 0x01 << 5  //+ LIN Break Detection Length.
	LBDIE  CR2 = 0x01 << 6  //+ LIN Break Detection Interrupt Enable.
	LBCL   CR2 = 0x01 << 8  //+ Last Bit Clock pulse.
	CPHA   CR2 = 0x01 << 9  //+ Clock Phase.
	CPOL   CR2 = 0x01 << 10 //+ Clock Polarity.
	CLKEN  CR2 = 0x01 << 11 //+ Clock Enable.
	STOP   CR2 = 0x03 << 12 //+ STOP[1:0] bits (STOP bits).
	STOP_0 CR2 = 0x01 << 12 //  Bit 0.
	STOP_1 CR2 = 0x02 << 12 //  Bit 1.
	LINEN  CR2 = 0x01 << 14 //+ LIN mode enable.
)

const (
	ADDn   = 0
	LBDLn  = 5
	LBDIEn = 6
	LBCLn  = 8
	CPHAn  = 9
	CPOLn  = 10
	CLKENn = 11
	STOPn  = 12
	LINENn = 14
)

const (
	EIE    CR3 = 0x01 << 0  //+ Error Interrupt Enable.
	IREN   CR3 = 0x01 << 1  //+ IrDA mode Enable.
	IRLP   CR3 = 0x01 << 2  //+ IrDA Low-Power.
	HDSEL  CR3 = 0x01 << 3  //+ Half-Duplex Selection.
	NACK   CR3 = 0x01 << 4  //+ Smartcard NACK enable.
	SCEN   CR3 = 0x01 << 5  //+ Smartcard mode enable.
	DMAR   CR3 = 0x01 << 6  //+ DMA Enable Receiver.
	DMAT   CR3 = 0x01 << 7  //+ DMA Enable Transmitter.
	RTSE   CR3 = 0x01 << 8  //+ RTS Enable.
	CTSE   CR3 = 0x01 << 9  //+ CTS Enable.
	CTSIE  CR3 = 0x01 << 10 //+ CTS Interrupt Enable.
	ONEBIT CR3 = 0x01 << 11 //+ One Bit method.
)

const (
	EIEn    = 0
	IRENn   = 1
	IRLPn   = 2
	HDSELn  = 3
	NACKn   = 4
	SCENn   = 5
	DMARn   = 6
	DMATn   = 7
	RTSEn   = 8
	CTSEn   = 9
	CTSIEn  = 10
	ONEBITn = 11
)

const (
	PSC   GTPR = 0xFF << 0 //+ PSC[7:0] bits (Prescaler value).
	PSC_0 GTPR = 0x01 << 0 //  Bit 0.
	PSC_1 GTPR = 0x02 << 0 //  Bit 1.
	PSC_2 GTPR = 0x04 << 0 //  Bit 2.
	PSC_3 GTPR = 0x08 << 0 //  Bit 3.
	PSC_4 GTPR = 0x10 << 0 //  Bit 4.
	PSC_5 GTPR = 0x20 << 0 //  Bit 5.
	PSC_6 GTPR = 0x40 << 0 //  Bit 6.
	PSC_7 GTPR = 0x80 << 0 //  Bit 7.
	GT    GTPR = 0xFF << 8 //+ Guard time value.
)

const (
	PSCn = 0
	GTn  = 8
)
