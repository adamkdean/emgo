// Peripheral: EXTI_Periph  External Interrupt/Event Controller.
// Instances:
//  EXTI  mmap.EXTI_BASE
// Registers:
//  0x00 32  IMR   Interrupt mask register.
//  0x04 32  EMR   Event mask register.
//  0x08 32  RTSR  Rising trigger selection register.
//  0x0C 32  FTSR  Falling trigger selection register.
//  0x10 32  SWIER Software interrupt event register.
//  0x14 32  PR    Pending register.
// Import:
//  stm32/o/f030x6/mmap
package exti

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	IL0    IMR = 0x01 << 0     //+ Interrupt Mask on line 0.
	IL1    IMR = 0x01 << 1     //+ Interrupt Mask on line 1.
	IL2    IMR = 0x01 << 2     //+ Interrupt Mask on line 2.
	IL3    IMR = 0x01 << 3     //+ Interrupt Mask on line 3.
	IL4    IMR = 0x01 << 4     //+ Interrupt Mask on line 4.
	IL5    IMR = 0x01 << 5     //+ Interrupt Mask on line 5.
	IL6    IMR = 0x01 << 6     //+ Interrupt Mask on line 6.
	IL7    IMR = 0x01 << 7     //+ Interrupt Mask on line 7.
	IL8    IMR = 0x01 << 8     //+ Interrupt Mask on line 8.
	IL9    IMR = 0x01 << 9     //+ Interrupt Mask on line 9.
	IL10   IMR = 0x01 << 10    //+ Interrupt Mask on line 10.
	IL11   IMR = 0x01 << 11    //+ Interrupt Mask on line 11.
	IL12   IMR = 0x01 << 12    //+ Interrupt Mask on line 12.
	IL13   IMR = 0x01 << 13    //+ Interrupt Mask on line 13.
	IL14   IMR = 0x01 << 14    //+ Interrupt Mask on line 14.
	IL15   IMR = 0x01 << 15    //+ Interrupt Mask on line 15.
	IL17   IMR = 0x01 << 17    //+ Interrupt Mask on line 17.
	IL18   IMR = 0x01 << 18    //+ Interrupt Mask on line 18.
	IL19   IMR = 0x01 << 19    //+ Interrupt Mask on line 19.
	IL23   IMR = 0x01 << 23    //+ Interrupt Mask on line 23.
	IMRALL IMR = 0x8EFFFF << 0 //  Interrupt Mask All.
)

const (
	IL0n  = 0
	IL1n  = 1
	IL2n  = 2
	IL3n  = 3
	IL4n  = 4
	IL5n  = 5
	IL6n  = 6
	IL7n  = 7
	IL8n  = 8
	IL9n  = 9
	IL10n = 10
	IL11n = 11
	IL12n = 12
	IL13n = 13
	IL14n = 14
	IL15n = 15
	IL17n = 17
	IL18n = 18
	IL19n = 19
	IL23n = 23
)

const (
	EL0  EMR = 0x01 << 0  //+ Event Mask on line 0.
	EL1  EMR = 0x01 << 1  //+ Event Mask on line 1.
	EL2  EMR = 0x01 << 2  //+ Event Mask on line 2.
	EL3  EMR = 0x01 << 3  //+ Event Mask on line 3.
	EL4  EMR = 0x01 << 4  //+ Event Mask on line 4.
	EL5  EMR = 0x01 << 5  //+ Event Mask on line 5.
	EL6  EMR = 0x01 << 6  //+ Event Mask on line 6.
	EL7  EMR = 0x01 << 7  //+ Event Mask on line 7.
	EL8  EMR = 0x01 << 8  //+ Event Mask on line 8.
	EL9  EMR = 0x01 << 9  //+ Event Mask on line 9.
	EL10 EMR = 0x01 << 10 //+ Event Mask on line 10.
	EL11 EMR = 0x01 << 11 //+ Event Mask on line 11.
	EL12 EMR = 0x01 << 12 //+ Event Mask on line 12.
	EL13 EMR = 0x01 << 13 //+ Event Mask on line 13.
	EL14 EMR = 0x01 << 14 //+ Event Mask on line 14.
	EL15 EMR = 0x01 << 15 //+ Event Mask on line 15.
	EL17 EMR = 0x01 << 17 //+ Event Mask on line 17.
	EL18 EMR = 0x01 << 18 //+ Event Mask on line 18.
	EL19 EMR = 0x01 << 19 //+ Event Mask on line 19.
	EL23 EMR = 0x01 << 23 //+ Event Mask on line 23.
)

const (
	EL0n  = 0
	EL1n  = 1
	EL2n  = 2
	EL3n  = 3
	EL4n  = 4
	EL5n  = 5
	EL6n  = 6
	EL7n  = 7
	EL8n  = 8
	EL9n  = 9
	EL10n = 10
	EL11n = 11
	EL12n = 12
	EL13n = 13
	EL14n = 14
	EL15n = 15
	EL17n = 17
	EL18n = 18
	EL19n = 19
	EL23n = 23
)

const (
	TR0  RTSR = 0x01 << 0  //+ Rising trigger event configuration bit of line 0.
	TR1  RTSR = 0x01 << 1  //+ Rising trigger event configuration bit of line 1.
	TR2  RTSR = 0x01 << 2  //+ Rising trigger event configuration bit of line 2.
	TR3  RTSR = 0x01 << 3  //+ Rising trigger event configuration bit of line 3.
	TR4  RTSR = 0x01 << 4  //+ Rising trigger event configuration bit of line 4.
	TR5  RTSR = 0x01 << 5  //+ Rising trigger event configuration bit of line 5.
	TR6  RTSR = 0x01 << 6  //+ Rising trigger event configuration bit of line 6.
	TR7  RTSR = 0x01 << 7  //+ Rising trigger event configuration bit of line 7.
	TR8  RTSR = 0x01 << 8  //+ Rising trigger event configuration bit of line 8.
	TR9  RTSR = 0x01 << 9  //+ Rising trigger event configuration bit of line 9.
	TR10 RTSR = 0x01 << 10 //+ Rising trigger event configuration bit of line 10.
	TR11 RTSR = 0x01 << 11 //+ Rising trigger event configuration bit of line 11.
	TR12 RTSR = 0x01 << 12 //+ Rising trigger event configuration bit of line 12.
	TR13 RTSR = 0x01 << 13 //+ Rising trigger event configuration bit of line 13.
	TR14 RTSR = 0x01 << 14 //+ Rising trigger event configuration bit of line 14.
	TR15 RTSR = 0x01 << 15 //+ Rising trigger event configuration bit of line 15.
	TR16 RTSR = 0x01 << 16 //+ Rising trigger event configuration bit of line 16.
	TR17 RTSR = 0x01 << 17 //+ Rising trigger event configuration bit of line 17.
	TR19 RTSR = 0x01 << 19 //+ Rising trigger event configuration bit of line 19.
)

const (
	TR0n  = 0
	TR1n  = 1
	TR2n  = 2
	TR3n  = 3
	TR4n  = 4
	TR5n  = 5
	TR6n  = 6
	TR7n  = 7
	TR8n  = 8
	TR9n  = 9
	TR10n = 10
	TR11n = 11
	TR12n = 12
	TR13n = 13
	TR14n = 14
	TR15n = 15
	TR16n = 16
	TR17n = 17
	TR19n = 19
)

const (
	TF0  FTSR = 0x01 << 0  //+ Falling trigger event configuration bit of line 0.
	TF1  FTSR = 0x01 << 1  //+ Falling trigger event configuration bit of line 1.
	TF2  FTSR = 0x01 << 2  //+ Falling trigger event configuration bit of line 2.
	TF3  FTSR = 0x01 << 3  //+ Falling trigger event configuration bit of line 3.
	TF4  FTSR = 0x01 << 4  //+ Falling trigger event configuration bit of line 4.
	TF5  FTSR = 0x01 << 5  //+ Falling trigger event configuration bit of line 5.
	TF6  FTSR = 0x01 << 6  //+ Falling trigger event configuration bit of line 6.
	TF7  FTSR = 0x01 << 7  //+ Falling trigger event configuration bit of line 7.
	TF8  FTSR = 0x01 << 8  //+ Falling trigger event configuration bit of line 8.
	TF9  FTSR = 0x01 << 9  //+ Falling trigger event configuration bit of line 9.
	TF10 FTSR = 0x01 << 10 //+ Falling trigger event configuration bit of line 10.
	TF11 FTSR = 0x01 << 11 //+ Falling trigger event configuration bit of line 11.
	TF12 FTSR = 0x01 << 12 //+ Falling trigger event configuration bit of line 12.
	TF13 FTSR = 0x01 << 13 //+ Falling trigger event configuration bit of line 13.
	TF14 FTSR = 0x01 << 14 //+ Falling trigger event configuration bit of line 14.
	TF15 FTSR = 0x01 << 15 //+ Falling trigger event configuration bit of line 15.
	TF16 FTSR = 0x01 << 16 //+ Falling trigger event configuration bit of line 16.
	TF17 FTSR = 0x01 << 17 //+ Falling trigger event configuration bit of line 17.
	TF19 FTSR = 0x01 << 19 //+ Falling trigger event configuration bit of line 19.
)

const (
	TF0n  = 0
	TF1n  = 1
	TF2n  = 2
	TF3n  = 3
	TF4n  = 4
	TF5n  = 5
	TF6n  = 6
	TF7n  = 7
	TF8n  = 8
	TF9n  = 9
	TF10n = 10
	TF11n = 11
	TF12n = 12
	TF13n = 13
	TF14n = 14
	TF15n = 15
	TF16n = 16
	TF17n = 17
	TF19n = 19
)

const (
	SWI0  SWIER = 0x01 << 0  //+ Software Interrupt on line 0.
	SWI1  SWIER = 0x01 << 1  //+ Software Interrupt on line 1.
	SWI2  SWIER = 0x01 << 2  //+ Software Interrupt on line 2.
	SWI3  SWIER = 0x01 << 3  //+ Software Interrupt on line 3.
	SWI4  SWIER = 0x01 << 4  //+ Software Interrupt on line 4.
	SWI5  SWIER = 0x01 << 5  //+ Software Interrupt on line 5.
	SWI6  SWIER = 0x01 << 6  //+ Software Interrupt on line 6.
	SWI7  SWIER = 0x01 << 7  //+ Software Interrupt on line 7.
	SWI8  SWIER = 0x01 << 8  //+ Software Interrupt on line 8.
	SWI9  SWIER = 0x01 << 9  //+ Software Interrupt on line 9.
	SWI10 SWIER = 0x01 << 10 //+ Software Interrupt on line 10.
	SWI11 SWIER = 0x01 << 11 //+ Software Interrupt on line 11.
	SWI12 SWIER = 0x01 << 12 //+ Software Interrupt on line 12.
	SWI13 SWIER = 0x01 << 13 //+ Software Interrupt on line 13.
	SWI14 SWIER = 0x01 << 14 //+ Software Interrupt on line 14.
	SWI15 SWIER = 0x01 << 15 //+ Software Interrupt on line 15.
	SWI16 SWIER = 0x01 << 16 //+ Software Interrupt on line 16.
	SWI17 SWIER = 0x01 << 17 //+ Software Interrupt on line 17.
	SWI19 SWIER = 0x01 << 19 //+ Software Interrupt on line 19.
)

const (
	SWI0n  = 0
	SWI1n  = 1
	SWI2n  = 2
	SWI3n  = 3
	SWI4n  = 4
	SWI5n  = 5
	SWI6n  = 6
	SWI7n  = 7
	SWI8n  = 8
	SWI9n  = 9
	SWI10n = 10
	SWI11n = 11
	SWI12n = 12
	SWI13n = 13
	SWI14n = 14
	SWI15n = 15
	SWI16n = 16
	SWI17n = 17
	SWI19n = 19
)

const (
	PIF0  PR = 0x01 << 0  //+ Pending bit 0.
	PIF1  PR = 0x01 << 1  //+ Pending bit 1.
	PIF2  PR = 0x01 << 2  //+ Pending bit 2.
	PIF3  PR = 0x01 << 3  //+ Pending bit 3.
	PIF4  PR = 0x01 << 4  //+ Pending bit 4.
	PIF5  PR = 0x01 << 5  //+ Pending bit 5.
	PIF6  PR = 0x01 << 6  //+ Pending bit 6.
	PIF7  PR = 0x01 << 7  //+ Pending bit 7.
	PIF8  PR = 0x01 << 8  //+ Pending bit 8.
	PIF9  PR = 0x01 << 9  //+ Pending bit 9.
	PIF10 PR = 0x01 << 10 //+ Pending bit 10.
	PIF11 PR = 0x01 << 11 //+ Pending bit 11.
	PIF12 PR = 0x01 << 12 //+ Pending bit 12.
	PIF13 PR = 0x01 << 13 //+ Pending bit 13.
	PIF14 PR = 0x01 << 14 //+ Pending bit 14.
	PIF15 PR = 0x01 << 15 //+ Pending bit 15.
	PIF16 PR = 0x01 << 16 //+ Pending bit 16.
	PIF17 PR = 0x01 << 17 //+ Pending bit 17.
	PIF19 PR = 0x01 << 19 //+ Pending bit 19.
)

const (
	PIF0n  = 0
	PIF1n  = 1
	PIF2n  = 2
	PIF3n  = 3
	PIF4n  = 4
	PIF5n  = 5
	PIF6n  = 6
	PIF7n  = 7
	PIF8n  = 8
	PIF9n  = 9
	PIF10n = 10
	PIF11n = 11
	PIF12n = 12
	PIF13n = 13
	PIF14n = 14
	PIF15n = 15
	PIF16n = 16
	PIF17n = 17
	PIF19n = 19
)
