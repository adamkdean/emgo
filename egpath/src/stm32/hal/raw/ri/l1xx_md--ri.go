// +build l1xx_md

// Peripheral: RI_Periph  Routing Interface.
// Instances:
//  RI  mmap.RI_BASE
// Registers:
//  0x00 32  ICR    Input capture register.
//  0x04 32  ASCR1  Analog switches control register.
//  0x08 32  ASCR2  Analog switch control register 2.
//  0x0C 32  HYSCR1 Hysteresis control register 1.
//  0x10 32  HYSCR2 Hysteresis control register 2.
//  0x14 32  HYSCR3 Hysteresis control register 3.
//  0x18 32  HYSCR4 Hysteresis control register 4.
//  0x1C 32  ASMR1  Analog switch mode register 1.
//  0x20 32  CMR1   Channel mask register 1.
//  0x24 32  CICR1  Channel identification for capture register 1.
//  0x28 32  ASMR2  Analog switch mode register 2.
//  0x2C 32  CMR2   Channel mask register 2.
//  0x30 32  CICR2  Channel identification for capture register 2.
//  0x34 32  ASMR3  Analog switch mode register 3.
//  0x38 32  CMR3   Channel mask register 3.
//  0x3C 32  CICR3  Channel identification for capture register3.
//  0x40 32  ASMR4  Analog switch mode register 4.
//  0x44 32  CMR4   Channel mask register 4.
//  0x48 32  CICR4  Channel identification for capture register 4.
//  0x4C 32  ASMR5  Analog switch mode register 5.
//  0x50 32  CMR5   Channel mask register 5.
//  0x54 32  CICR5  Channel identification for capture register 5.
// Import:
//  stm32/o/l1xx_md/mmap
package ri

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	IC1Z   ICR = 0x0F << 0  //+ IC1Z[3:0] bits (Input Capture 1 select bits).
	IC1Z_0 ICR = 0x01 << 0  //  Bit 0.
	IC1Z_1 ICR = 0x02 << 0  //  Bit 1.
	IC1Z_2 ICR = 0x04 << 0  //  Bit 2.
	IC1Z_3 ICR = 0x08 << 0  //  Bit 3.
	IC2Z   ICR = 0x0F << 4  //+ IC2Z[3:0] bits (Input Capture 2 select bits).
	IC2Z_0 ICR = 0x01 << 4  //  Bit 0.
	IC2Z_1 ICR = 0x02 << 4  //  Bit 1.
	IC2Z_2 ICR = 0x04 << 4  //  Bit 2.
	IC2Z_3 ICR = 0x08 << 4  //  Bit 3.
	IC3Z   ICR = 0x0F << 8  //+ IC3Z[3:0] bits (Input Capture 3 select bits).
	IC3Z_0 ICR = 0x01 << 8  //  Bit 0.
	IC3Z_1 ICR = 0x02 << 8  //  Bit 1.
	IC3Z_2 ICR = 0x04 << 8  //  Bit 2.
	IC3Z_3 ICR = 0x08 << 8  //  Bit 3.
	IC4Z   ICR = 0x0F << 12 //+ IC4Z[3:0] bits (Input Capture 4 select bits).
	IC4Z_0 ICR = 0x01 << 12 //  Bit 0.
	IC4Z_1 ICR = 0x02 << 12 //  Bit 1.
	IC4Z_2 ICR = 0x04 << 12 //  Bit 2.
	IC4Z_3 ICR = 0x08 << 12 //  Bit 3.
	TIM    ICR = 0x03 << 16 //+ TIM[3:0] bits (Timers select bits).
	TIM_0  ICR = 0x01 << 16 //  Bit 0.
	TIM_1  ICR = 0x02 << 16 //  Bit 1.
	IC1    ICR = 0x01 << 18 //+ Input capture 1.
	IC2    ICR = 0x01 << 19 //+ Input capture 2.
	IC3    ICR = 0x01 << 20 //+ Input capture 3.
	IC4    ICR = 0x01 << 21 //+ Input capture 4.
)

const (
	IC1Zn = 0
	IC2Zn = 4
	IC3Zn = 8
	IC4Zn = 12
	TIMn  = 16
	IC1n  = 18
	IC2n  = 19
	IC3n  = 20
	IC4n  = 21
)

const (
	CH    ASCR1 = 0x3FCFFFF << 0 //+ AS_CH[25:18] & AS_CH[15:0] bits ( Analog switches selection bits).
	CH_0  ASCR1 = 0x01 << 0      //  Bit 0.
	CH_1  ASCR1 = 0x02 << 0      //  Bit 1.
	CH_2  ASCR1 = 0x04 << 0      //  Bit 2.
	CH_3  ASCR1 = 0x08 << 0      //  Bit 3.
	CH_4  ASCR1 = 0x10 << 0      //  Bit 4.
	CH_5  ASCR1 = 0x20 << 0      //  Bit 5.
	CH_6  ASCR1 = 0x40 << 0      //  Bit 6.
	CH_7  ASCR1 = 0x80 << 0      //  Bit 7.
	CH_8  ASCR1 = 0x100 << 0     //  Bit 8.
	CH_9  ASCR1 = 0x200 << 0     //  Bit 9.
	CH_10 ASCR1 = 0x400 << 0     //  Bit 10.
	CH_11 ASCR1 = 0x800 << 0     //  Bit 11.
	CH_12 ASCR1 = 0x1000 << 0    //  Bit 12.
	CH_13 ASCR1 = 0x2000 << 0    //  Bit 13.
	CH_14 ASCR1 = 0x4000 << 0    //  Bit 14.
	CH_15 ASCR1 = 0x8000 << 0    //  Bit 15.
	CH_31 ASCR1 = 0x01 << 16     //+ Bit 16.
	CH_18 ASCR1 = 0x40000 << 0   //  Bit 18.
	CH_19 ASCR1 = 0x80000 << 0   //  Bit 19.
	CH_20 ASCR1 = 0x100000 << 0  //  Bit 20.
	CH_21 ASCR1 = 0x200000 << 0  //  Bit 21.
	CH_22 ASCR1 = 0x400000 << 0  //  Bit 22.
	CH_23 ASCR1 = 0x800000 << 0  //  Bit 23.
	CH_24 ASCR1 = 0x1000000 << 0 //  Bit 24.
	CH_25 ASCR1 = 0x2000000 << 0 //  Bit 25.
	VCOMP ASCR1 = 0x01 << 26     //+ ADC analog switch selection for internal node to COMP1.
	CH_27 ASCR1 = 0x400000 << 0  //  Bit 27.
	CH_28 ASCR1 = 0x800000 << 0  //  Bit 28.
	CH_29 ASCR1 = 0x1000000 << 0 //  Bit 29.
	CH_30 ASCR1 = 0x2000000 << 0 //  Bit 30.
	SCM   ASCR1 = 0x01 << 31     //+ I/O Switch control mode.
)

const (
	CHn    = 0
	CH_31n = 16
	VCOMPn = 26
	SCMn   = 31
)

const (
	GR10_1 ASCR2 = 0x01 << 0  //+ GR10-1 selection bit.
	GR10_2 ASCR2 = 0x01 << 1  //+ GR10-2 selection bit.
	GR10_3 ASCR2 = 0x01 << 2  //+ GR10-3 selection bit.
	GR10_4 ASCR2 = 0x01 << 3  //+ GR10-4 selection bit.
	GR6_1  ASCR2 = 0x01 << 4  //+ GR6-1 selection bit.
	GR6_2  ASCR2 = 0x01 << 5  //+ GR6-2 selection bit.
	GR5_1  ASCR2 = 0x01 << 6  //+ GR5-1 selection bit.
	GR5_2  ASCR2 = 0x01 << 7  //+ GR5-2 selection bit.
	GR5_3  ASCR2 = 0x01 << 8  //+ GR5-3 selection bit.
	GR4_1  ASCR2 = 0x01 << 9  //+ GR4-1 selection bit.
	GR4_2  ASCR2 = 0x01 << 10 //+ GR4-2 selection bit.
	GR4_3  ASCR2 = 0x01 << 11 //+ GR4-3 selection bit.
	GR4_4  ASCR2 = 0x01 << 15 //+ GR4-4 selection bit.
	CH0b   ASCR2 = 0x01 << 16 //+ CH0b selection bit.
	CH1b   ASCR2 = 0x01 << 17 //+ CH1b selection bit.
	CH2b   ASCR2 = 0x01 << 18 //+ CH2b selection bit.
	CH3b   ASCR2 = 0x01 << 19 //+ CH3b selection bit.
	CH6b   ASCR2 = 0x01 << 20 //+ CH6b selection bit.
	CH7b   ASCR2 = 0x01 << 21 //+ CH7b selection bit.
	CH8b   ASCR2 = 0x01 << 22 //+ CH8b selection bit.
	CH9b   ASCR2 = 0x01 << 23 //+ CH9b selection bit.
	CH10b  ASCR2 = 0x01 << 24 //+ CH10b selection bit.
	CH11b  ASCR2 = 0x01 << 25 //+ CH11b selection bit.
	CH12b  ASCR2 = 0x01 << 26 //+ CH12b selection bit.
	GR6_3  ASCR2 = 0x01 << 27 //+ GR6-3 selection bit.
	GR6_4  ASCR2 = 0x01 << 28 //+ GR6-4 selection bit.
	GR5_4  ASCR2 = 0x01 << 29 //+ GR5-4 selection bit.
)

const (
	GR10_1n = 0
	GR10_2n = 1
	GR10_3n = 2
	GR10_4n = 3
	GR6_1n  = 4
	GR6_2n  = 5
	GR5_1n  = 6
	GR5_2n  = 7
	GR5_3n  = 8
	GR4_1n  = 9
	GR4_2n  = 10
	GR4_3n  = 11
	GR4_4n  = 15
	CH0bn   = 16
	CH1bn   = 17
	CH2bn   = 18
	CH3bn   = 19
	CH6bn   = 20
	CH7bn   = 21
	CH8bn   = 22
	CH9bn   = 23
	CH10bn  = 24
	CH11bn  = 25
	CH12bn  = 26
	GR6_3n  = 27
	GR6_4n  = 28
	GR5_4n  = 29
)

const (
	PA    HYSCR1 = 0xFFFF << 0  //+ PA[15:0] Port A Hysteresis selection.
	PA_0  HYSCR1 = 0x01 << 0    //  Bit 0.
	PA_1  HYSCR1 = 0x02 << 0    //  Bit 1.
	PA_2  HYSCR1 = 0x04 << 0    //  Bit 2.
	PA_3  HYSCR1 = 0x08 << 0    //  Bit 3.
	PA_4  HYSCR1 = 0x10 << 0    //  Bit 4.
	PA_5  HYSCR1 = 0x20 << 0    //  Bit 5.
	PA_6  HYSCR1 = 0x40 << 0    //  Bit 6.
	PA_7  HYSCR1 = 0x80 << 0    //  Bit 7.
	PA_8  HYSCR1 = 0x100 << 0   //  Bit 8.
	PA_9  HYSCR1 = 0x200 << 0   //  Bit 9.
	PA_10 HYSCR1 = 0x400 << 0   //  Bit 10.
	PA_11 HYSCR1 = 0x800 << 0   //  Bit 11.
	PA_12 HYSCR1 = 0x1000 << 0  //  Bit 12.
	PA_13 HYSCR1 = 0x2000 << 0  //  Bit 13.
	PA_14 HYSCR1 = 0x4000 << 0  //  Bit 14.
	PA_15 HYSCR1 = 0x8000 << 0  //  Bit 15.
	PB    HYSCR1 = 0xFFFF << 16 //+ PB[15:0] Port B Hysteresis selection.
	PB_0  HYSCR1 = 0x01 << 16   //  Bit 0.
	PB_1  HYSCR1 = 0x02 << 16   //  Bit 1.
	PB_2  HYSCR1 = 0x04 << 16   //  Bit 2.
	PB_3  HYSCR1 = 0x08 << 16   //  Bit 3.
	PB_4  HYSCR1 = 0x10 << 16   //  Bit 4.
	PB_5  HYSCR1 = 0x20 << 16   //  Bit 5.
	PB_6  HYSCR1 = 0x40 << 16   //  Bit 6.
	PB_7  HYSCR1 = 0x80 << 16   //  Bit 7.
	PB_8  HYSCR1 = 0x100 << 16  //  Bit 8.
	PB_9  HYSCR1 = 0x200 << 16  //  Bit 9.
	PB_10 HYSCR1 = 0x400 << 16  //  Bit 10.
	PB_11 HYSCR1 = 0x800 << 16  //  Bit 11.
	PB_12 HYSCR1 = 0x1000 << 16 //  Bit 12.
	PB_13 HYSCR1 = 0x2000 << 16 //  Bit 13.
	PB_14 HYSCR1 = 0x4000 << 16 //  Bit 14.
	PB_15 HYSCR1 = 0x8000 << 16 //  Bit 15.
)

const (
	PAn = 0
	PBn = 16
)

const (
	PC    HYSCR2 = 0xFFFF << 0  //+ PC[15:0] Port C Hysteresis selection.
	PC_0  HYSCR2 = 0x01 << 0    //  Bit 0.
	PC_1  HYSCR2 = 0x02 << 0    //  Bit 1.
	PC_2  HYSCR2 = 0x04 << 0    //  Bit 2.
	PC_3  HYSCR2 = 0x08 << 0    //  Bit 3.
	PC_4  HYSCR2 = 0x10 << 0    //  Bit 4.
	PC_5  HYSCR2 = 0x20 << 0    //  Bit 5.
	PC_6  HYSCR2 = 0x40 << 0    //  Bit 6.
	PC_7  HYSCR2 = 0x80 << 0    //  Bit 7.
	PC_8  HYSCR2 = 0x100 << 0   //  Bit 8.
	PC_9  HYSCR2 = 0x200 << 0   //  Bit 9.
	PC_10 HYSCR2 = 0x400 << 0   //  Bit 10.
	PC_11 HYSCR2 = 0x800 << 0   //  Bit 11.
	PC_12 HYSCR2 = 0x1000 << 0  //  Bit 12.
	PC_13 HYSCR2 = 0x2000 << 0  //  Bit 13.
	PC_14 HYSCR2 = 0x4000 << 0  //  Bit 14.
	PC_15 HYSCR2 = 0x8000 << 0  //  Bit 15.
	PD    HYSCR2 = 0xFFFF << 16 //+ PD[15:0] Port D Hysteresis selection.
	PD_0  HYSCR2 = 0x01 << 16   //  Bit 0.
	PD_1  HYSCR2 = 0x02 << 16   //  Bit 1.
	PD_2  HYSCR2 = 0x04 << 16   //  Bit 2.
	PD_3  HYSCR2 = 0x08 << 16   //  Bit 3.
	PD_4  HYSCR2 = 0x10 << 16   //  Bit 4.
	PD_5  HYSCR2 = 0x20 << 16   //  Bit 5.
	PD_6  HYSCR2 = 0x40 << 16   //  Bit 6.
	PD_7  HYSCR2 = 0x80 << 16   //  Bit 7.
	PD_8  HYSCR2 = 0x100 << 16  //  Bit 8.
	PD_9  HYSCR2 = 0x200 << 16  //  Bit 9.
	PD_10 HYSCR2 = 0x400 << 16  //  Bit 10.
	PD_11 HYSCR2 = 0x800 << 16  //  Bit 11.
	PD_12 HYSCR2 = 0x1000 << 16 //  Bit 12.
	PD_13 HYSCR2 = 0x2000 << 16 //  Bit 13.
	PD_14 HYSCR2 = 0x4000 << 16 //  Bit 14.
	PD_15 HYSCR2 = 0x8000 << 16 //  Bit 15.
	PE    HYSCR2 = 0xFFFF << 0  //  PE[15:0] Port E Hysteresis selection.
	PE_0  HYSCR2 = 0x01 << 0    //  Bit 0.
	PE_1  HYSCR2 = 0x02 << 0    //  Bit 1.
	PE_2  HYSCR2 = 0x04 << 0    //  Bit 2.
	PE_3  HYSCR2 = 0x08 << 0    //  Bit 3.
	PE_4  HYSCR2 = 0x10 << 0    //  Bit 4.
	PE_5  HYSCR2 = 0x20 << 0    //  Bit 5.
	PE_6  HYSCR2 = 0x40 << 0    //  Bit 6.
	PE_7  HYSCR2 = 0x80 << 0    //  Bit 7.
	PE_8  HYSCR2 = 0x100 << 0   //  Bit 8.
	PE_9  HYSCR2 = 0x200 << 0   //  Bit 9.
	PE_10 HYSCR2 = 0x400 << 0   //  Bit 10.
	PE_11 HYSCR2 = 0x800 << 0   //  Bit 11.
	PE_12 HYSCR2 = 0x1000 << 0  //  Bit 12.
	PE_13 HYSCR2 = 0x2000 << 0  //  Bit 13.
	PE_14 HYSCR2 = 0x4000 << 0  //  Bit 14.
	PE_15 HYSCR2 = 0x8000 << 0  //  Bit 15.
)

const (
	PCn = 0
	PDn = 16
)

const (
	PF    HYSCR3 = 0xFFFF << 16 //+ PF[15:0] Port F Hysteresis selection.
	PF_0  HYSCR3 = 0x01 << 16   //  Bit 0.
	PF_1  HYSCR3 = 0x02 << 16   //  Bit 1.
	PF_2  HYSCR3 = 0x04 << 16   //  Bit 2.
	PF_3  HYSCR3 = 0x08 << 16   //  Bit 3.
	PF_4  HYSCR3 = 0x10 << 16   //  Bit 4.
	PF_5  HYSCR3 = 0x20 << 16   //  Bit 5.
	PF_6  HYSCR3 = 0x40 << 16   //  Bit 6.
	PF_7  HYSCR3 = 0x80 << 16   //  Bit 7.
	PF_8  HYSCR3 = 0x100 << 16  //  Bit 8.
	PF_9  HYSCR3 = 0x200 << 16  //  Bit 9.
	PF_10 HYSCR3 = 0x400 << 16  //  Bit 10.
	PF_11 HYSCR3 = 0x800 << 16  //  Bit 11.
	PF_12 HYSCR3 = 0x1000 << 16 //  Bit 12.
	PF_13 HYSCR3 = 0x2000 << 16 //  Bit 13.
	PF_14 HYSCR3 = 0x4000 << 16 //  Bit 14.
	PF_15 HYSCR3 = 0x8000 << 16 //  Bit 15.
)

const (
	PFn = 16
)

const (
	PG    HYSCR4 = 0xFFFF << 0 //+ PG[15:0] Port G Hysteresis selection.
	PG_0  HYSCR4 = 0x01 << 0   //  Bit 0.
	PG_1  HYSCR4 = 0x02 << 0   //  Bit 1.
	PG_2  HYSCR4 = 0x04 << 0   //  Bit 2.
	PG_3  HYSCR4 = 0x08 << 0   //  Bit 3.
	PG_4  HYSCR4 = 0x10 << 0   //  Bit 4.
	PG_5  HYSCR4 = 0x20 << 0   //  Bit 5.
	PG_6  HYSCR4 = 0x40 << 0   //  Bit 6.
	PG_7  HYSCR4 = 0x80 << 0   //  Bit 7.
	PG_8  HYSCR4 = 0x100 << 0  //  Bit 8.
	PG_9  HYSCR4 = 0x200 << 0  //  Bit 9.
	PG_10 HYSCR4 = 0x400 << 0  //  Bit 10.
	PG_11 HYSCR4 = 0x800 << 0  //  Bit 11.
	PG_12 HYSCR4 = 0x1000 << 0 //  Bit 12.
	PG_13 HYSCR4 = 0x2000 << 0 //  Bit 13.
	PG_14 HYSCR4 = 0x4000 << 0 //  Bit 14.
	PG_15 HYSCR4 = 0x8000 << 0 //  Bit 15.
)

const (
	PGn = 0
)

const (
	PA    ASMR1 = 0xFFFF << 0 //+ PA[15:0] Port A analog switch mode selection.
	PA_0  ASMR1 = 0x01 << 0   //  Bit 0.
	PA_1  ASMR1 = 0x02 << 0   //  Bit 1.
	PA_2  ASMR1 = 0x04 << 0   //  Bit 2.
	PA_3  ASMR1 = 0x08 << 0   //  Bit 3.
	PA_4  ASMR1 = 0x10 << 0   //  Bit 4.
	PA_5  ASMR1 = 0x20 << 0   //  Bit 5.
	PA_6  ASMR1 = 0x40 << 0   //  Bit 6.
	PA_7  ASMR1 = 0x80 << 0   //  Bit 7.
	PA_8  ASMR1 = 0x100 << 0  //  Bit 8.
	PA_9  ASMR1 = 0x200 << 0  //  Bit 9.
	PA_10 ASMR1 = 0x400 << 0  //  Bit 10.
	PA_11 ASMR1 = 0x800 << 0  //  Bit 11.
	PA_12 ASMR1 = 0x1000 << 0 //  Bit 12.
	PA_13 ASMR1 = 0x2000 << 0 //  Bit 13.
	PA_14 ASMR1 = 0x4000 << 0 //  Bit 14.
	PA_15 ASMR1 = 0x8000 << 0 //  Bit 15.
)

const (
	PAn = 0
)

const (
	PA    CMR1 = 0xFFFF << 0 //+ PA[15:0] Port A channel masking.
	PA_0  CMR1 = 0x01 << 0   //  Bit 0.
	PA_1  CMR1 = 0x02 << 0   //  Bit 1.
	PA_2  CMR1 = 0x04 << 0   //  Bit 2.
	PA_3  CMR1 = 0x08 << 0   //  Bit 3.
	PA_4  CMR1 = 0x10 << 0   //  Bit 4.
	PA_5  CMR1 = 0x20 << 0   //  Bit 5.
	PA_6  CMR1 = 0x40 << 0   //  Bit 6.
	PA_7  CMR1 = 0x80 << 0   //  Bit 7.
	PA_8  CMR1 = 0x100 << 0  //  Bit 8.
	PA_9  CMR1 = 0x200 << 0  //  Bit 9.
	PA_10 CMR1 = 0x400 << 0  //  Bit 10.
	PA_11 CMR1 = 0x800 << 0  //  Bit 11.
	PA_12 CMR1 = 0x1000 << 0 //  Bit 12.
	PA_13 CMR1 = 0x2000 << 0 //  Bit 13.
	PA_14 CMR1 = 0x4000 << 0 //  Bit 14.
	PA_15 CMR1 = 0x8000 << 0 //  Bit 15.
)

const (
	PAn = 0
)

const (
	PA    CICR1 = 0xFFFF << 0 //+ PA[15:0] Port A channel identification for capture.
	PA_0  CICR1 = 0x01 << 0   //  Bit 0.
	PA_1  CICR1 = 0x02 << 0   //  Bit 1.
	PA_2  CICR1 = 0x04 << 0   //  Bit 2.
	PA_3  CICR1 = 0x08 << 0   //  Bit 3.
	PA_4  CICR1 = 0x10 << 0   //  Bit 4.
	PA_5  CICR1 = 0x20 << 0   //  Bit 5.
	PA_6  CICR1 = 0x40 << 0   //  Bit 6.
	PA_7  CICR1 = 0x80 << 0   //  Bit 7.
	PA_8  CICR1 = 0x100 << 0  //  Bit 8.
	PA_9  CICR1 = 0x200 << 0  //  Bit 9.
	PA_10 CICR1 = 0x400 << 0  //  Bit 10.
	PA_11 CICR1 = 0x800 << 0  //  Bit 11.
	PA_12 CICR1 = 0x1000 << 0 //  Bit 12.
	PA_13 CICR1 = 0x2000 << 0 //  Bit 13.
	PA_14 CICR1 = 0x4000 << 0 //  Bit 14.
	PA_15 CICR1 = 0x8000 << 0 //  Bit 15.
)

const (
	PAn = 0
)

const (
	PB    ASMR2 = 0xFFFF << 0 //+ PB[15:0] Port B analog switch mode selection.
	PB_0  ASMR2 = 0x01 << 0   //  Bit 0.
	PB_1  ASMR2 = 0x02 << 0   //  Bit 1.
	PB_2  ASMR2 = 0x04 << 0   //  Bit 2.
	PB_3  ASMR2 = 0x08 << 0   //  Bit 3.
	PB_4  ASMR2 = 0x10 << 0   //  Bit 4.
	PB_5  ASMR2 = 0x20 << 0   //  Bit 5.
	PB_6  ASMR2 = 0x40 << 0   //  Bit 6.
	PB_7  ASMR2 = 0x80 << 0   //  Bit 7.
	PB_8  ASMR2 = 0x100 << 0  //  Bit 8.
	PB_9  ASMR2 = 0x200 << 0  //  Bit 9.
	PB_10 ASMR2 = 0x400 << 0  //  Bit 10.
	PB_11 ASMR2 = 0x800 << 0  //  Bit 11.
	PB_12 ASMR2 = 0x1000 << 0 //  Bit 12.
	PB_13 ASMR2 = 0x2000 << 0 //  Bit 13.
	PB_14 ASMR2 = 0x4000 << 0 //  Bit 14.
	PB_15 ASMR2 = 0x8000 << 0 //  Bit 15.
)

const (
	PBn = 0
)

const (
	PB    CMR2 = 0xFFFF << 0 //+ PB[15:0] Port B channel masking.
	PB_0  CMR2 = 0x01 << 0   //  Bit 0.
	PB_1  CMR2 = 0x02 << 0   //  Bit 1.
	PB_2  CMR2 = 0x04 << 0   //  Bit 2.
	PB_3  CMR2 = 0x08 << 0   //  Bit 3.
	PB_4  CMR2 = 0x10 << 0   //  Bit 4.
	PB_5  CMR2 = 0x20 << 0   //  Bit 5.
	PB_6  CMR2 = 0x40 << 0   //  Bit 6.
	PB_7  CMR2 = 0x80 << 0   //  Bit 7.
	PB_8  CMR2 = 0x100 << 0  //  Bit 8.
	PB_9  CMR2 = 0x200 << 0  //  Bit 9.
	PB_10 CMR2 = 0x400 << 0  //  Bit 10.
	PB_11 CMR2 = 0x800 << 0  //  Bit 11.
	PB_12 CMR2 = 0x1000 << 0 //  Bit 12.
	PB_13 CMR2 = 0x2000 << 0 //  Bit 13.
	PB_14 CMR2 = 0x4000 << 0 //  Bit 14.
	PB_15 CMR2 = 0x8000 << 0 //  Bit 15.
)

const (
	PBn = 0
)

const (
	PB    CICR2 = 0xFFFF << 0 //+ PB[15:0] Port B channel identification for capture.
	PB_0  CICR2 = 0x01 << 0   //  Bit 0.
	PB_1  CICR2 = 0x02 << 0   //  Bit 1.
	PB_2  CICR2 = 0x04 << 0   //  Bit 2.
	PB_3  CICR2 = 0x08 << 0   //  Bit 3.
	PB_4  CICR2 = 0x10 << 0   //  Bit 4.
	PB_5  CICR2 = 0x20 << 0   //  Bit 5.
	PB_6  CICR2 = 0x40 << 0   //  Bit 6.
	PB_7  CICR2 = 0x80 << 0   //  Bit 7.
	PB_8  CICR2 = 0x100 << 0  //  Bit 8.
	PB_9  CICR2 = 0x200 << 0  //  Bit 9.
	PB_10 CICR2 = 0x400 << 0  //  Bit 10.
	PB_11 CICR2 = 0x800 << 0  //  Bit 11.
	PB_12 CICR2 = 0x1000 << 0 //  Bit 12.
	PB_13 CICR2 = 0x2000 << 0 //  Bit 13.
	PB_14 CICR2 = 0x4000 << 0 //  Bit 14.
	PB_15 CICR2 = 0x8000 << 0 //  Bit 15.
)

const (
	PBn = 0
)

const (
	PC    ASMR3 = 0xFFFF << 0 //+ PC[15:0] Port C analog switch mode selection.
	PC_0  ASMR3 = 0x01 << 0   //  Bit 0.
	PC_1  ASMR3 = 0x02 << 0   //  Bit 1.
	PC_2  ASMR3 = 0x04 << 0   //  Bit 2.
	PC_3  ASMR3 = 0x08 << 0   //  Bit 3.
	PC_4  ASMR3 = 0x10 << 0   //  Bit 4.
	PC_5  ASMR3 = 0x20 << 0   //  Bit 5.
	PC_6  ASMR3 = 0x40 << 0   //  Bit 6.
	PC_7  ASMR3 = 0x80 << 0   //  Bit 7.
	PC_8  ASMR3 = 0x100 << 0  //  Bit 8.
	PC_9  ASMR3 = 0x200 << 0  //  Bit 9.
	PC_10 ASMR3 = 0x400 << 0  //  Bit 10.
	PC_11 ASMR3 = 0x800 << 0  //  Bit 11.
	PC_12 ASMR3 = 0x1000 << 0 //  Bit 12.
	PC_13 ASMR3 = 0x2000 << 0 //  Bit 13.
	PC_14 ASMR3 = 0x4000 << 0 //  Bit 14.
	PC_15 ASMR3 = 0x8000 << 0 //  Bit 15.
)

const (
	PCn = 0
)

const (
	PC    CMR3 = 0xFFFF << 0 //+ PC[15:0] Port C channel masking.
	PC_0  CMR3 = 0x01 << 0   //  Bit 0.
	PC_1  CMR3 = 0x02 << 0   //  Bit 1.
	PC_2  CMR3 = 0x04 << 0   //  Bit 2.
	PC_3  CMR3 = 0x08 << 0   //  Bit 3.
	PC_4  CMR3 = 0x10 << 0   //  Bit 4.
	PC_5  CMR3 = 0x20 << 0   //  Bit 5.
	PC_6  CMR3 = 0x40 << 0   //  Bit 6.
	PC_7  CMR3 = 0x80 << 0   //  Bit 7.
	PC_8  CMR3 = 0x100 << 0  //  Bit 8.
	PC_9  CMR3 = 0x200 << 0  //  Bit 9.
	PC_10 CMR3 = 0x400 << 0  //  Bit 10.
	PC_11 CMR3 = 0x800 << 0  //  Bit 11.
	PC_12 CMR3 = 0x1000 << 0 //  Bit 12.
	PC_13 CMR3 = 0x2000 << 0 //  Bit 13.
	PC_14 CMR3 = 0x4000 << 0 //  Bit 14.
	PC_15 CMR3 = 0x8000 << 0 //  Bit 15.
)

const (
	PCn = 0
)

const (
	PC    CICR3 = 0xFFFF << 0 //+ PC[15:0] Port C channel identification for capture.
	PC_0  CICR3 = 0x01 << 0   //  Bit 0.
	PC_1  CICR3 = 0x02 << 0   //  Bit 1.
	PC_2  CICR3 = 0x04 << 0   //  Bit 2.
	PC_3  CICR3 = 0x08 << 0   //  Bit 3.
	PC_4  CICR3 = 0x10 << 0   //  Bit 4.
	PC_5  CICR3 = 0x20 << 0   //  Bit 5.
	PC_6  CICR3 = 0x40 << 0   //  Bit 6.
	PC_7  CICR3 = 0x80 << 0   //  Bit 7.
	PC_8  CICR3 = 0x100 << 0  //  Bit 8.
	PC_9  CICR3 = 0x200 << 0  //  Bit 9.
	PC_10 CICR3 = 0x400 << 0  //  Bit 10.
	PC_11 CICR3 = 0x800 << 0  //  Bit 11.
	PC_12 CICR3 = 0x1000 << 0 //  Bit 12.
	PC_13 CICR3 = 0x2000 << 0 //  Bit 13.
	PC_14 CICR3 = 0x4000 << 0 //  Bit 14.
	PC_15 CICR3 = 0x8000 << 0 //  Bit 15.
)

const (
	PCn = 0
)

const (
	PF    ASMR4 = 0xFFFF << 0 //+ PF[15:0] Port F analog switch mode selection.
	PF_0  ASMR4 = 0x01 << 0   //  Bit 0.
	PF_1  ASMR4 = 0x02 << 0   //  Bit 1.
	PF_2  ASMR4 = 0x04 << 0   //  Bit 2.
	PF_3  ASMR4 = 0x08 << 0   //  Bit 3.
	PF_4  ASMR4 = 0x10 << 0   //  Bit 4.
	PF_5  ASMR4 = 0x20 << 0   //  Bit 5.
	PF_6  ASMR4 = 0x40 << 0   //  Bit 6.
	PF_7  ASMR4 = 0x80 << 0   //  Bit 7.
	PF_8  ASMR4 = 0x100 << 0  //  Bit 8.
	PF_9  ASMR4 = 0x200 << 0  //  Bit 9.
	PF_10 ASMR4 = 0x400 << 0  //  Bit 10.
	PF_11 ASMR4 = 0x800 << 0  //  Bit 11.
	PF_12 ASMR4 = 0x1000 << 0 //  Bit 12.
	PF_13 ASMR4 = 0x2000 << 0 //  Bit 13.
	PF_14 ASMR4 = 0x4000 << 0 //  Bit 14.
	PF_15 ASMR4 = 0x8000 << 0 //  Bit 15.
)

const (
	PFn = 0
)

const (
	PF    CMR4 = 0xFFFF << 0 //+ PF[15:0] Port F channel masking.
	PF_0  CMR4 = 0x01 << 0   //  Bit 0.
	PF_1  CMR4 = 0x02 << 0   //  Bit 1.
	PF_2  CMR4 = 0x04 << 0   //  Bit 2.
	PF_3  CMR4 = 0x08 << 0   //  Bit 3.
	PF_4  CMR4 = 0x10 << 0   //  Bit 4.
	PF_5  CMR4 = 0x20 << 0   //  Bit 5.
	PF_6  CMR4 = 0x40 << 0   //  Bit 6.
	PF_7  CMR4 = 0x80 << 0   //  Bit 7.
	PF_8  CMR4 = 0x100 << 0  //  Bit 8.
	PF_9  CMR4 = 0x200 << 0  //  Bit 9.
	PF_10 CMR4 = 0x400 << 0  //  Bit 10.
	PF_11 CMR4 = 0x800 << 0  //  Bit 11.
	PF_12 CMR4 = 0x1000 << 0 //  Bit 12.
	PF_13 CMR4 = 0x2000 << 0 //  Bit 13.
	PF_14 CMR4 = 0x4000 << 0 //  Bit 14.
	PF_15 CMR4 = 0x8000 << 0 //  Bit 15.
)

const (
	PFn = 0
)

const (
	PF    CICR4 = 0xFFFF << 0 //+ PF[15:0] Port F channel identification for capture.
	PF_0  CICR4 = 0x01 << 0   //  Bit 0.
	PF_1  CICR4 = 0x02 << 0   //  Bit 1.
	PF_2  CICR4 = 0x04 << 0   //  Bit 2.
	PF_3  CICR4 = 0x08 << 0   //  Bit 3.
	PF_4  CICR4 = 0x10 << 0   //  Bit 4.
	PF_5  CICR4 = 0x20 << 0   //  Bit 5.
	PF_6  CICR4 = 0x40 << 0   //  Bit 6.
	PF_7  CICR4 = 0x80 << 0   //  Bit 7.
	PF_8  CICR4 = 0x100 << 0  //  Bit 8.
	PF_9  CICR4 = 0x200 << 0  //  Bit 9.
	PF_10 CICR4 = 0x400 << 0  //  Bit 10.
	PF_11 CICR4 = 0x800 << 0  //  Bit 11.
	PF_12 CICR4 = 0x1000 << 0 //  Bit 12.
	PF_13 CICR4 = 0x2000 << 0 //  Bit 13.
	PF_14 CICR4 = 0x4000 << 0 //  Bit 14.
	PF_15 CICR4 = 0x8000 << 0 //  Bit 15.
)

const (
	PFn = 0
)

const (
	PG    ASMR5 = 0xFFFF << 0 //+ PG[15:0] Port G analog switch mode selection.
	PG_0  ASMR5 = 0x01 << 0   //  Bit 0.
	PG_1  ASMR5 = 0x02 << 0   //  Bit 1.
	PG_2  ASMR5 = 0x04 << 0   //  Bit 2.
	PG_3  ASMR5 = 0x08 << 0   //  Bit 3.
	PG_4  ASMR5 = 0x10 << 0   //  Bit 4.
	PG_5  ASMR5 = 0x20 << 0   //  Bit 5.
	PG_6  ASMR5 = 0x40 << 0   //  Bit 6.
	PG_7  ASMR5 = 0x80 << 0   //  Bit 7.
	PG_8  ASMR5 = 0x100 << 0  //  Bit 8.
	PG_9  ASMR5 = 0x200 << 0  //  Bit 9.
	PG_10 ASMR5 = 0x400 << 0  //  Bit 10.
	PG_11 ASMR5 = 0x800 << 0  //  Bit 11.
	PG_12 ASMR5 = 0x1000 << 0 //  Bit 12.
	PG_13 ASMR5 = 0x2000 << 0 //  Bit 13.
	PG_14 ASMR5 = 0x4000 << 0 //  Bit 14.
	PG_15 ASMR5 = 0x8000 << 0 //  Bit 15.
)

const (
	PGn = 0
)

const (
	PG    CMR5 = 0xFFFF << 0 //+ PG[15:0] Port G channel masking.
	PG_0  CMR5 = 0x01 << 0   //  Bit 0.
	PG_1  CMR5 = 0x02 << 0   //  Bit 1.
	PG_2  CMR5 = 0x04 << 0   //  Bit 2.
	PG_3  CMR5 = 0x08 << 0   //  Bit 3.
	PG_4  CMR5 = 0x10 << 0   //  Bit 4.
	PG_5  CMR5 = 0x20 << 0   //  Bit 5.
	PG_6  CMR5 = 0x40 << 0   //  Bit 6.
	PG_7  CMR5 = 0x80 << 0   //  Bit 7.
	PG_8  CMR5 = 0x100 << 0  //  Bit 8.
	PG_9  CMR5 = 0x200 << 0  //  Bit 9.
	PG_10 CMR5 = 0x400 << 0  //  Bit 10.
	PG_11 CMR5 = 0x800 << 0  //  Bit 11.
	PG_12 CMR5 = 0x1000 << 0 //  Bit 12.
	PG_13 CMR5 = 0x2000 << 0 //  Bit 13.
	PG_14 CMR5 = 0x4000 << 0 //  Bit 14.
	PG_15 CMR5 = 0x8000 << 0 //  Bit 15.
)

const (
	PGn = 0
)

const (
	PG    CICR5 = 0xFFFF << 0 //+ PG[15:0] Port G channel identification for capture.
	PG_0  CICR5 = 0x01 << 0   //  Bit 0.
	PG_1  CICR5 = 0x02 << 0   //  Bit 1.
	PG_2  CICR5 = 0x04 << 0   //  Bit 2.
	PG_3  CICR5 = 0x08 << 0   //  Bit 3.
	PG_4  CICR5 = 0x10 << 0   //  Bit 4.
	PG_5  CICR5 = 0x20 << 0   //  Bit 5.
	PG_6  CICR5 = 0x40 << 0   //  Bit 6.
	PG_7  CICR5 = 0x80 << 0   //  Bit 7.
	PG_8  CICR5 = 0x100 << 0  //  Bit 8.
	PG_9  CICR5 = 0x200 << 0  //  Bit 9.
	PG_10 CICR5 = 0x400 << 0  //  Bit 10.
	PG_11 CICR5 = 0x800 << 0  //  Bit 11.
	PG_12 CICR5 = 0x1000 << 0 //  Bit 12.
	PG_13 CICR5 = 0x2000 << 0 //  Bit 13.
	PG_14 CICR5 = 0x4000 << 0 //  Bit 14.
	PG_15 CICR5 = 0x8000 << 0 //  Bit 15.
)

const (
	PGn = 0
)
