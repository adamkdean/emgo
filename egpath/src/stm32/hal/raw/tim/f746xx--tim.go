// +build f746xx

// Peripheral: TIM_Periph  TIM.
// Instances:
//  TIM2   mmap.TIM2_BASE
//  TIM3   mmap.TIM3_BASE
//  TIM4   mmap.TIM4_BASE
//  TIM5   mmap.TIM5_BASE
//  TIM6   mmap.TIM6_BASE
//  TIM7   mmap.TIM7_BASE
//  TIM12  mmap.TIM12_BASE
//  TIM13  mmap.TIM13_BASE
//  TIM14  mmap.TIM14_BASE
//  TIM1   mmap.TIM1_BASE
//  TIM8   mmap.TIM8_BASE
//  TIM9   mmap.TIM9_BASE
//  TIM10  mmap.TIM10_BASE
//  TIM11  mmap.TIM11_BASE
// Registers:
//  0x00 32  CR1   Control register 1.
//  0x04 32  CR2   Control register 2.
//  0x08 32  SMCR  Slave mode control register.
//  0x0C 32  DIER  DMA/interrupt enable register.
//  0x10 32  SR    Status register.
//  0x14 32  EGR   Event generation register.
//  0x18 32  CCMR1 Capture/compare mode register 1.
//  0x1C 32  CCMR2 Capture/compare mode register 2.
//  0x20 32  CCER  Capture/compare enable register.
//  0x24 32  CNT   Counter register.
//  0x28 32  PSC   Prescaler.
//  0x2C 32  ARR   Auto-reload register.
//  0x30 32  RCR   Repetition counter register.
//  0x34 32  CCR1  Capture/compare register 1.
//  0x38 32  CCR2  Capture/compare register 2.
//  0x3C 32  CCR3  Capture/compare register 3.
//  0x40 32  CCR4  Capture/compare register 4.
//  0x44 32  BDTR  Break and dead-time register.
//  0x48 32  DCR   DMA control register.
//  0x4C 32  DMAR  DMA address for full transfer.
//  0x50 32  OR    Option register.
//  0x54 32  CCMR3 Capture/compare mode register 3.
//  0x58 32  CCR5  Capture/compare mode register5.
//  0x5C 32  CCR6  Capture/compare mode register6.
// Import:
//  stm32/o/f746xx/mmap
package tim

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CEN      CR1 = 0x01 << 0  //+ Counter enable.
	UDIS     CR1 = 0x01 << 1  //+ Update disable.
	URS      CR1 = 0x01 << 2  //+ Update request source.
	OPM      CR1 = 0x01 << 3  //+ One pulse mode.
	DIR      CR1 = 0x01 << 4  //+ Direction.
	CMS      CR1 = 0x03 << 5  //+ CMS[1:0] bits (Center-aligned mode selection).
	CMS_0    CR1 = 0x01 << 5  //  Bit 0.
	CMS_1    CR1 = 0x02 << 5  //  Bit 1.
	ARPE     CR1 = 0x01 << 7  //+ Auto-reload preload enable.
	CKD      CR1 = 0x03 << 8  //+ CKD[1:0] bits (clock division).
	CKD_0    CR1 = 0x01 << 8  //  Bit 0.
	CKD_1    CR1 = 0x02 << 8  //  Bit 1.
	UIFREMAP CR1 = 0x01 << 11 //+ UIF status bit.
)

const (
	CENn      = 0
	UDISn     = 1
	URSn      = 2
	OPMn      = 3
	DIRn      = 4
	CMSn      = 5
	ARPEn     = 7
	CKDn      = 8
	UIFREMAPn = 11
)

const (
	CCPC   CR2 = 0x01 << 0  //+ Capture/Compare Preloaded Control.
	CCUS   CR2 = 0x01 << 2  //+ Capture/Compare Control Update Selection.
	CCDS   CR2 = 0x01 << 3  //+ Capture/Compare DMA Selection.
	OIS5   CR2 = 0x01 << 16 //+ Output Idle state 4 (OC4 output).
	OIS6   CR2 = 0x01 << 18 //+ Output Idle state 4 (OC4 output).
	MMS    CR2 = 0x07 << 4  //+ MMS[2:0] bits (Master Mode Selection).
	MMS_0  CR2 = 0x01 << 4  //  Bit 0.
	MMS_1  CR2 = 0x02 << 4  //  Bit 1.
	MMS_2  CR2 = 0x04 << 4  //  Bit 2.
	MMS2   CR2 = 0x0F << 20 //+ MMS[2:0] bits (Master Mode Selection).
	MMS2_0 CR2 = 0x01 << 20 //  Bit 0.
	MMS2_1 CR2 = 0x02 << 20 //  Bit 1.
	MMS2_2 CR2 = 0x04 << 20 //  Bit 2.
	MMS2_3 CR2 = 0x08 << 20 //  Bit 2.
	TI1S   CR2 = 0x01 << 7  //+ TI1 Selection.
	OIS1   CR2 = 0x01 << 8  //+ Output Idle state 1 (OC1 output).
	OIS1N  CR2 = 0x01 << 9  //+ Output Idle state 1 (OC1N output).
	OIS2   CR2 = 0x01 << 10 //+ Output Idle state 2 (OC2 output).
	OIS2N  CR2 = 0x01 << 11 //+ Output Idle state 2 (OC2N output).
	OIS3   CR2 = 0x01 << 12 //+ Output Idle state 3 (OC3 output).
	OIS3N  CR2 = 0x01 << 13 //+ Output Idle state 3 (OC3N output).
	OIS4   CR2 = 0x01 << 14 //+ Output Idle state 4 (OC4 output).
)

const (
	CCPCn  = 0
	CCUSn  = 2
	CCDSn  = 3
	OIS5n  = 16
	OIS6n  = 18
	MMSn   = 4
	MMS2n  = 20
	TI1Sn  = 7
	OIS1n  = 8
	OIS1Nn = 9
	OIS2n  = 10
	OIS2Nn = 11
	OIS3n  = 12
	OIS3Nn = 13
	OIS4n  = 14
)

const (
	SMS    SMCR = 0x10007 << 0 //+ SMS[2:0] bits (Slave mode selection).
	SMS_0  SMCR = 0x01 << 0    //  Bit 0.
	SMS_1  SMCR = 0x02 << 0    //  Bit 1.
	SMS_2  SMCR = 0x04 << 0    //  Bit 2.
	SMS_3  SMCR = 0x10000 << 0 //  Bit 3.
	OCCS   SMCR = 0x01 << 3    //+ OCREF clear selection.
	TS     SMCR = 0x07 << 4    //+ TS[2:0] bits (Trigger selection).
	TS_0   SMCR = 0x01 << 4    //  Bit 0.
	TS_1   SMCR = 0x02 << 4    //  Bit 1.
	TS_2   SMCR = 0x04 << 4    //  Bit 2.
	MSM    SMCR = 0x01 << 7    //+ Master/slave mode.
	ETF    SMCR = 0x0F << 8    //+ ETF[3:0] bits (External trigger filter).
	ETF_0  SMCR = 0x01 << 8    //  Bit 0.
	ETF_1  SMCR = 0x02 << 8    //  Bit 1.
	ETF_2  SMCR = 0x04 << 8    //  Bit 2.
	ETF_3  SMCR = 0x08 << 8    //  Bit 3.
	ETPS   SMCR = 0x03 << 12   //+ ETPS[1:0] bits (External trigger prescaler).
	ETPS_0 SMCR = 0x01 << 12   //  Bit 0.
	ETPS_1 SMCR = 0x02 << 12   //  Bit 1.
	ECE    SMCR = 0x01 << 14   //+ External clock enable.
	ETP    SMCR = 0x01 << 15   //+ External trigger polarity.
)

const (
	SMSn  = 0
	OCCSn = 3
	TSn   = 4
	MSMn  = 7
	ETFn  = 8
	ETPSn = 12
	ECEn  = 14
	ETPn  = 15
)

const (
	UIE   DIER = 0x01 << 0  //+ Update interrupt enable.
	CC1IE DIER = 0x01 << 1  //+ Capture/Compare 1 interrupt enable.
	CC2IE DIER = 0x01 << 2  //+ Capture/Compare 2 interrupt enable.
	CC3IE DIER = 0x01 << 3  //+ Capture/Compare 3 interrupt enable.
	CC4IE DIER = 0x01 << 4  //+ Capture/Compare 4 interrupt enable.
	COMIE DIER = 0x01 << 5  //+ COM interrupt enable.
	TIE   DIER = 0x01 << 6  //+ Trigger interrupt enable.
	BIE   DIER = 0x01 << 7  //+ Break interrupt enable.
	UDE   DIER = 0x01 << 8  //+ Update DMA request enable.
	CC1DE DIER = 0x01 << 9  //+ Capture/Compare 1 DMA request enable.
	CC2DE DIER = 0x01 << 10 //+ Capture/Compare 2 DMA request enable.
	CC3DE DIER = 0x01 << 11 //+ Capture/Compare 3 DMA request enable.
	CC4DE DIER = 0x01 << 12 //+ Capture/Compare 4 DMA request enable.
	COMDE DIER = 0x01 << 13 //+ COM DMA request enable.
	TDE   DIER = 0x01 << 14 //+ Trigger DMA request enable.
)

const (
	UIEn   = 0
	CC1IEn = 1
	CC2IEn = 2
	CC3IEn = 3
	CC4IEn = 4
	COMIEn = 5
	TIEn   = 6
	BIEn   = 7
	UDEn   = 8
	CC1DEn = 9
	CC2DEn = 10
	CC3DEn = 11
	CC4DEn = 12
	COMDEn = 13
	TDEn   = 14
)

const (
	UIF   SR = 0x01 << 0  //+ Update interrupt Flag.
	CC1IF SR = 0x01 << 1  //+ Capture/Compare 1 interrupt Flag.
	CC2IF SR = 0x01 << 2  //+ Capture/Compare 2 interrupt Flag.
	CC3IF SR = 0x01 << 3  //+ Capture/Compare 3 interrupt Flag.
	CC4IF SR = 0x01 << 4  //+ Capture/Compare 4 interrupt Flag.
	COMIF SR = 0x01 << 5  //+ COM interrupt Flag.
	TIF   SR = 0x01 << 6  //+ Trigger interrupt Flag.
	BIF   SR = 0x01 << 7  //+ Break interrupt Flag.
	B2IF  SR = 0x01 << 8  //+ Break2 interrupt Flag.
	CC1OF SR = 0x01 << 9  //+ Capture/Compare 1 Overcapture Flag.
	CC2OF SR = 0x01 << 10 //+ Capture/Compare 2 Overcapture Flag.
	CC3OF SR = 0x01 << 11 //+ Capture/Compare 3 Overcapture Flag.
	CC4OF SR = 0x01 << 12 //+ Capture/Compare 4 Overcapture Flag.
)

const (
	UIFn   = 0
	CC1IFn = 1
	CC2IFn = 2
	CC3IFn = 3
	CC4IFn = 4
	COMIFn = 5
	TIFn   = 6
	BIFn   = 7
	B2IFn  = 8
	CC1OFn = 9
	CC2OFn = 10
	CC3OFn = 11
	CC4OFn = 12
)

const (
	UG   EGR = 0x01 << 0 //+ Update Generation.
	CC1G EGR = 0x01 << 1 //+ Capture/Compare 1 Generation.
	CC2G EGR = 0x01 << 2 //+ Capture/Compare 2 Generation.
	CC3G EGR = 0x01 << 3 //+ Capture/Compare 3 Generation.
	CC4G EGR = 0x01 << 4 //+ Capture/Compare 4 Generation.
	COMG EGR = 0x01 << 5 //+ Capture/Compare Control Update Generation.
	TG   EGR = 0x01 << 6 //+ Trigger Generation.
	BG   EGR = 0x01 << 7 //+ Break Generation.
	B2G  EGR = 0x01 << 8 //+ Break2 Generation.
)

const (
	UGn   = 0
	CC1Gn = 1
	CC2Gn = 2
	CC3Gn = 3
	CC4Gn = 4
	COMGn = 5
	TGn   = 6
	BGn   = 7
	B2Gn  = 8
)

const (
	CC1S     CCMR1 = 0x03 << 0    //+ CC1S[1:0] bits (Capture/Compare 1 Selection).
	CC1S_0   CCMR1 = 0x01 << 0    //  Bit 0.
	CC1S_1   CCMR1 = 0x02 << 0    //  Bit 1.
	OC1FE    CCMR1 = 0x01 << 2    //+ Output Compare 1 Fast enable.
	OC1PE    CCMR1 = 0x01 << 3    //+ Output Compare 1 Preload enable.
	OC1M     CCMR1 = 0x1007 << 4  //+ OC1M[2:0] bits (Output Compare 1 Mode).
	OC1M_0   CCMR1 = 0x01 << 4    //  Bit 0.
	OC1M_1   CCMR1 = 0x02 << 4    //  Bit 1.
	OC1M_2   CCMR1 = 0x04 << 4    //  Bit 2.
	OC1M_3   CCMR1 = 0x1000 << 4  //  Bit 3.
	OC1CE    CCMR1 = 0x01 << 7    //+ Output Compare 1Clear Enable.
	CC2S     CCMR1 = 0x03 << 8    //+ CC2S[1:0] bits (Capture/Compare 2 Selection).
	CC2S_0   CCMR1 = 0x01 << 8    //  Bit 0.
	CC2S_1   CCMR1 = 0x02 << 8    //  Bit 1.
	OC2FE    CCMR1 = 0x01 << 10   //+ Output Compare 2 Fast enable.
	OC2PE    CCMR1 = 0x01 << 11   //+ Output Compare 2 Preload enable.
	OC2M     CCMR1 = 0x1007 << 12 //+ OC2M[2:0] bits (Output Compare 2 Mode).
	OC2M_0   CCMR1 = 0x01 << 12   //  Bit 0.
	OC2M_1   CCMR1 = 0x02 << 12   //  Bit 1.
	OC2M_2   CCMR1 = 0x04 << 12   //  Bit 2.
	OC2M_3   CCMR1 = 0x1000 << 12 //  Bit 3.
	OC2CE    CCMR1 = 0x01 << 15   //+ Output Compare 2 Clear Enable.
	IC1PSC   CCMR1 = 0x03 << 2    //+ IC1PSC[1:0] bits (Input Capture 1 Prescaler).
	IC1PSC_0 CCMR1 = 0x01 << 2    //  Bit 0.
	IC1PSC_1 CCMR1 = 0x01 << 3    //  Bit 1.
	IC1F     CCMR1 = 0x0F << 4    //+ IC1F[3:0] bits (Input Capture 1 Filter).
	IC1F_0   CCMR1 = 0x01 << 4    //  Bit 0.
	IC1F_1   CCMR1 = 0x02 << 4    //  Bit 1.
	IC1F_2   CCMR1 = 0x04 << 4    //  Bit 2.
	IC1F_3   CCMR1 = 0x01 << 7    //  Bit 3.
	IC2PSC   CCMR1 = 0x03 << 10   //+ IC2PSC[1:0] bits (Input Capture 2 Prescaler).
	IC2PSC_0 CCMR1 = 0x01 << 10   //  Bit 0.
	IC2PSC_1 CCMR1 = 0x01 << 11   //  Bit 1.
	IC2F     CCMR1 = 0x0F << 12   //+ IC2F[3:0] bits (Input Capture 2 Filter).
	IC2F_0   CCMR1 = 0x01 << 12   //  Bit 0.
	IC2F_1   CCMR1 = 0x02 << 12   //  Bit 1.
	IC2F_2   CCMR1 = 0x04 << 12   //  Bit 2.
	IC2F_3   CCMR1 = 0x01 << 15   //  Bit 3.
)

const (
	CC1Sn   = 0
	OC1FEn  = 2
	OC1PEn  = 3
	OC1Mn   = 4
	OC1CEn  = 7
	CC2Sn   = 8
	OC2FEn  = 10
	OC2PEn  = 11
	OC2Mn   = 12
	OC2CEn  = 15
	IC1PSCn = 2
	IC1Fn   = 4
	IC2PSCn = 10
	IC2Fn   = 12
)

const (
	CC3S     CCMR2 = 0x03 << 0    //+ CC3S[1:0] bits (Capture/Compare 3 Selection).
	CC3S_0   CCMR2 = 0x01 << 0    //  Bit 0.
	CC3S_1   CCMR2 = 0x02 << 0    //  Bit 1.
	OC3FE    CCMR2 = 0x01 << 2    //+ Output Compare 3 Fast enable.
	OC3PE    CCMR2 = 0x01 << 3    //+ Output Compare 3 Preload enable.
	OC3M     CCMR2 = 0x1007 << 4  //+ OC3M[2:0] bits (Output Compare 3 Mode).
	OC3M_0   CCMR2 = 0x01 << 4    //  Bit 0.
	OC3M_1   CCMR2 = 0x02 << 4    //  Bit 1.
	OC3M_2   CCMR2 = 0x04 << 4    //  Bit 2.
	OC3M_3   CCMR2 = 0x1000 << 4  //  Bit 3.
	OC3CE    CCMR2 = 0x01 << 7    //+ Output Compare 3 Clear Enable.
	CC4S     CCMR2 = 0x03 << 8    //+ CC4S[1:0] bits (Capture/Compare 4 Selection).
	CC4S_0   CCMR2 = 0x01 << 8    //  Bit 0.
	CC4S_1   CCMR2 = 0x02 << 8    //  Bit 1.
	OC4FE    CCMR2 = 0x01 << 10   //+ Output Compare 4 Fast enable.
	OC4PE    CCMR2 = 0x01 << 11   //+ Output Compare 4 Preload enable.
	OC4M     CCMR2 = 0x1007 << 12 //+ OC4M[2:0] bits (Output Compare 4 Mode).
	OC4M_0   CCMR2 = 0x01 << 12   //  Bit 0.
	OC4M_1   CCMR2 = 0x02 << 12   //  Bit 1.
	OC4M_2   CCMR2 = 0x04 << 12   //  Bit 2.
	OC4M_3   CCMR2 = 0x1000 << 12 //  Bit 3.
	OC4CE    CCMR2 = 0x01 << 15   //+ Output Compare 4 Clear Enable.
	IC3PSC   CCMR2 = 0x03 << 2    //+ IC3PSC[1:0] bits (Input Capture 3 Prescaler).
	IC3PSC_0 CCMR2 = 0x01 << 2    //  Bit 0.
	IC3PSC_1 CCMR2 = 0x01 << 3    //  Bit 1.
	IC3F     CCMR2 = 0x0F << 4    //+ IC3F[3:0] bits (Input Capture 3 Filter).
	IC3F_0   CCMR2 = 0x01 << 4    //  Bit 0.
	IC3F_1   CCMR2 = 0x02 << 4    //  Bit 1.
	IC3F_2   CCMR2 = 0x04 << 4    //  Bit 2.
	IC3F_3   CCMR2 = 0x01 << 7    //  Bit 3.
	IC4PSC   CCMR2 = 0x03 << 10   //+ IC4PSC[1:0] bits (Input Capture 4 Prescaler).
	IC4PSC_0 CCMR2 = 0x01 << 10   //  Bit 0.
	IC4PSC_1 CCMR2 = 0x01 << 11   //  Bit 1.
	IC4F     CCMR2 = 0x0F << 12   //+ IC4F[3:0] bits (Input Capture 4 Filter).
	IC4F_0   CCMR2 = 0x01 << 12   //  Bit 0.
	IC4F_1   CCMR2 = 0x02 << 12   //  Bit 1.
	IC4F_2   CCMR2 = 0x04 << 12   //  Bit 2.
	IC4F_3   CCMR2 = 0x01 << 15   //  Bit 3.
)

const (
	CC3Sn   = 0
	OC3FEn  = 2
	OC3PEn  = 3
	OC3Mn   = 4
	OC3CEn  = 7
	CC4Sn   = 8
	OC4FEn  = 10
	OC4PEn  = 11
	OC4Mn   = 12
	OC4CEn  = 15
	IC3PSCn = 2
	IC3Fn   = 4
	IC4PSCn = 10
	IC4Fn   = 12
)

const (
	CC1E  CCER = 0x01 << 0  //+ Capture/Compare 1 output enable.
	CC1P  CCER = 0x01 << 1  //+ Capture/Compare 1 output Polarity.
	CC1NE CCER = 0x01 << 2  //+ Capture/Compare 1 Complementary output enable.
	CC1NP CCER = 0x01 << 3  //+ Capture/Compare 1 Complementary output Polarity.
	CC2E  CCER = 0x01 << 4  //+ Capture/Compare 2 output enable.
	CC2P  CCER = 0x01 << 5  //+ Capture/Compare 2 output Polarity.
	CC2NE CCER = 0x01 << 6  //+ Capture/Compare 2 Complementary output enable.
	CC2NP CCER = 0x01 << 7  //+ Capture/Compare 2 Complementary output Polarity.
	CC3E  CCER = 0x01 << 8  //+ Capture/Compare 3 output enable.
	CC3P  CCER = 0x01 << 9  //+ Capture/Compare 3 output Polarity.
	CC3NE CCER = 0x01 << 10 //+ Capture/Compare 3 Complementary output enable.
	CC3NP CCER = 0x01 << 11 //+ Capture/Compare 3 Complementary output Polarity.
	CC4E  CCER = 0x01 << 12 //+ Capture/Compare 4 output enable.
	CC4P  CCER = 0x01 << 13 //+ Capture/Compare 4 output Polarity.
	CC4NP CCER = 0x01 << 15 //+ Capture/Compare 4 Complementary output Polarity.
	CC5E  CCER = 0x01 << 16 //+ Capture/Compare 5 output enable.
	CC5P  CCER = 0x01 << 17 //+ Capture/Compare 5 output Polarity.
	CC6E  CCER = 0x01 << 20 //+ Capture/Compare 6 output enable.
	CC6P  CCER = 0x01 << 21 //+ Capture/Compare 6 output Polarity.
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NEn = 2
	CC1NPn = 3
	CC2En  = 4
	CC2Pn  = 5
	CC2NEn = 6
	CC2NPn = 7
	CC3En  = 8
	CC3Pn  = 9
	CC3NEn = 10
	CC3NPn = 11
	CC4En  = 12
	CC4Pn  = 13
	CC4NPn = 15
	CC5En  = 16
	CC5Pn  = 17
	CC6En  = 20
	CC6Pn  = 21
)

const (
	REP RCR = 0xFF << 0 //+ Repetition Counter Value.
)

const (
	REPn = 0
)

const (
	DTG    BDTR = 0xFF << 0  //+ DTG[0:7] bits (Dead-Time Generator set-up).
	DTG_0  BDTR = 0x01 << 0  //  Bit 0.
	DTG_1  BDTR = 0x02 << 0  //  Bit 1.
	DTG_2  BDTR = 0x04 << 0  //  Bit 2.
	DTG_3  BDTR = 0x08 << 0  //  Bit 3.
	DTG_4  BDTR = 0x10 << 0  //  Bit 4.
	DTG_5  BDTR = 0x20 << 0  //  Bit 5.
	DTG_6  BDTR = 0x40 << 0  //  Bit 6.
	DTG_7  BDTR = 0x80 << 0  //  Bit 7.
	LOCK   BDTR = 0x03 << 8  //+ LOCK[1:0] bits (Lock Configuration).
	LOCK_0 BDTR = 0x01 << 8  //  Bit 0.
	LOCK_1 BDTR = 0x02 << 8  //  Bit 1.
	OSSI   BDTR = 0x01 << 10 //+ Off-State Selection for Idle mode.
	OSSR   BDTR = 0x01 << 11 //+ Off-State Selection for Run mode.
	BKE    BDTR = 0x01 << 12 //+ Break enable.
	BKP    BDTR = 0x01 << 13 //+ Break Polarity.
	AOE    BDTR = 0x01 << 14 //+ Automatic Output enable.
	MOE    BDTR = 0x01 << 15 //+ Main Output enable.
	BKF    BDTR = 0x0F << 16 //+ Break Filter for Break1.
	BK2F   BDTR = 0x0F << 20 //+ Break Filter for Break2.
	BK2E   BDTR = 0x01 << 24 //+ Break enable for Break2.
	BK2P   BDTR = 0x01 << 25 //+ Break Polarity for Break2.
)

const (
	DTGn  = 0
	LOCKn = 8
	OSSIn = 10
	OSSRn = 11
	BKEn  = 12
	BKPn  = 13
	AOEn  = 14
	MOEn  = 15
	BKFn  = 16
	BK2Fn = 20
	BK2En = 24
	BK2Pn = 25
)

const (
	DBA   DCR = 0x1F << 0 //+ DBA[4:0] bits (DMA Base Address).
	DBA_0 DCR = 0x01 << 0 //  Bit 0.
	DBA_1 DCR = 0x02 << 0 //  Bit 1.
	DBA_2 DCR = 0x04 << 0 //  Bit 2.
	DBA_3 DCR = 0x08 << 0 //  Bit 3.
	DBA_4 DCR = 0x10 << 0 //  Bit 4.
	DBL   DCR = 0x1F << 8 //+ DBL[4:0] bits (DMA Burst Length).
	DBL_0 DCR = 0x01 << 8 //  Bit 0.
	DBL_1 DCR = 0x02 << 8 //  Bit 1.
	DBL_2 DCR = 0x04 << 8 //  Bit 2.
	DBL_3 DCR = 0x08 << 8 //  Bit 3.
	DBL_4 DCR = 0x10 << 8 //  Bit 4.
)

const (
	DBAn = 0
	DBLn = 8
)

const (
	DMAB DMAR = 0xFFFF << 0 //+ DMA register for burst accesses.
)

const (
	DMABn = 0
)

const (
	TI4_RMP    OR = 0x03 << 6  //+ TI4_RMP[1:0] bits (TIM5 Input 4 remap).
	TI4_RMP_0  OR = 0x01 << 6  //  Bit 0.
	TI4_RMP_1  OR = 0x02 << 6  //  Bit 1.
	ITR1_RMP   OR = 0x03 << 10 //+ ITR1_RMP[1:0] bits (TIM2 Internal trigger 1 remap).
	ITR1_RMP_0 OR = 0x01 << 10 //  Bit 0.
	ITR1_RMP_1 OR = 0x02 << 10 //  Bit 1.
)

const (
	TI4_RMPn  = 6
	ITR1_RMPn = 10
)

const (
	OC5FE  CCMR3 = 0x01 << 2    //+ Output Compare 5 Fast enable.
	OC5PE  CCMR3 = 0x01 << 3    //+ Output Compare 5 Preload enable.
	OC5M   CCMR3 = 0x1007 << 4  //+ OC5M[2:0] bits (Output Compare 5 Mode).
	OC5M_0 CCMR3 = 0x01 << 4    //  Bit 0.
	OC5M_1 CCMR3 = 0x02 << 4    //  Bit 1.
	OC5M_2 CCMR3 = 0x04 << 4    //  Bit 2.
	OC5M_3 CCMR3 = 0x1000 << 4  //  Bit 3.
	OC5CE  CCMR3 = 0x01 << 7    //+ Output Compare 5 Clear Enable.
	OC6FE  CCMR3 = 0x01 << 10   //+ Output Compare 4 Fast enable.
	OC6PE  CCMR3 = 0x01 << 11   //+ Output Compare 4 Preload enable.
	OC6M   CCMR3 = 0x1007 << 12 //+ OC4M[2:0] bits (Output Compare 4 Mode).
	OC6M_0 CCMR3 = 0x01 << 12   //  Bit 0.
	OC6M_1 CCMR3 = 0x02 << 12   //  Bit 1.
	OC6M_2 CCMR3 = 0x04 << 12   //  Bit 2.
	OC6M_3 CCMR3 = 0x1000 << 12 //  Bit 3.
	OC6CE  CCMR3 = 0x01 << 15   //+ Output Compare 4 Clear Enable.
)

const (
	OC5FEn = 2
	OC5PEn = 3
	OC5Mn  = 4
	OC5CEn = 7
	OC6FEn = 10
	OC6PEn = 11
	OC6Mn  = 12
	OC6CEn = 15
)

const (
	CCR5V CCR5 = 0xFFFFFFFF << 0 //+ Capture/Compare 5 Value.
	GC5C1 CCR5 = 0x20000000 << 0 //  Group Channel 5 and Channel 1.
	GC5C2 CCR5 = 0x40000000 << 0 //  Group Channel 5 and Channel 2.
	GC5C3 CCR5 = 0x80000000 << 0 //  Group Channel 5 and Channel 3.
)

const (
	CCR5Vn = 0
)
