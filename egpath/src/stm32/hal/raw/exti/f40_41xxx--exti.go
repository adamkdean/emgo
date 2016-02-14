// +build f40_41xxx

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
//  stm32/o/f40_41xxx/mmap
package exti

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	IL0  IMR_Bits = 0x01 << 0  //+ Interrupt Mask on line 0.
	IL1  IMR_Bits = 0x01 << 1  //+ Interrupt Mask on line 1.
	IL2  IMR_Bits = 0x01 << 2  //+ Interrupt Mask on line 2.
	IL3  IMR_Bits = 0x01 << 3  //+ Interrupt Mask on line 3.
	IL4  IMR_Bits = 0x01 << 4  //+ Interrupt Mask on line 4.
	IL5  IMR_Bits = 0x01 << 5  //+ Interrupt Mask on line 5.
	IL6  IMR_Bits = 0x01 << 6  //+ Interrupt Mask on line 6.
	IL7  IMR_Bits = 0x01 << 7  //+ Interrupt Mask on line 7.
	IL8  IMR_Bits = 0x01 << 8  //+ Interrupt Mask on line 8.
	IL9  IMR_Bits = 0x01 << 9  //+ Interrupt Mask on line 9.
	IL10 IMR_Bits = 0x01 << 10 //+ Interrupt Mask on line 10.
	IL11 IMR_Bits = 0x01 << 11 //+ Interrupt Mask on line 11.
	IL12 IMR_Bits = 0x01 << 12 //+ Interrupt Mask on line 12.
	IL13 IMR_Bits = 0x01 << 13 //+ Interrupt Mask on line 13.
	IL14 IMR_Bits = 0x01 << 14 //+ Interrupt Mask on line 14.
	IL15 IMR_Bits = 0x01 << 15 //+ Interrupt Mask on line 15.
	IL16 IMR_Bits = 0x01 << 16 //+ Interrupt Mask on line 16.
	IL17 IMR_Bits = 0x01 << 17 //+ Interrupt Mask on line 17.
	IL18 IMR_Bits = 0x01 << 18 //+ Interrupt Mask on line 18.
	IL19 IMR_Bits = 0x01 << 19 //+ Interrupt Mask on line 19.
	IL23 IMR_Bits = 0x01 << 23 //+ Interrupt Mask on line 23.
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
	IL16n = 16
	IL17n = 17
	IL18n = 18
	IL19n = 19
	IL23n = 23
)

const (
	EL0  EMR_Bits = 0x01 << 0  //+ Event Mask on line 0.
	EL1  EMR_Bits = 0x01 << 1  //+ Event Mask on line 1.
	EL2  EMR_Bits = 0x01 << 2  //+ Event Mask on line 2.
	EL3  EMR_Bits = 0x01 << 3  //+ Event Mask on line 3.
	EL4  EMR_Bits = 0x01 << 4  //+ Event Mask on line 4.
	EL5  EMR_Bits = 0x01 << 5  //+ Event Mask on line 5.
	EL6  EMR_Bits = 0x01 << 6  //+ Event Mask on line 6.
	EL7  EMR_Bits = 0x01 << 7  //+ Event Mask on line 7.
	EL8  EMR_Bits = 0x01 << 8  //+ Event Mask on line 8.
	EL9  EMR_Bits = 0x01 << 9  //+ Event Mask on line 9.
	EL10 EMR_Bits = 0x01 << 10 //+ Event Mask on line 10.
	EL11 EMR_Bits = 0x01 << 11 //+ Event Mask on line 11.
	EL12 EMR_Bits = 0x01 << 12 //+ Event Mask on line 12.
	EL13 EMR_Bits = 0x01 << 13 //+ Event Mask on line 13.
	EL14 EMR_Bits = 0x01 << 14 //+ Event Mask on line 14.
	EL15 EMR_Bits = 0x01 << 15 //+ Event Mask on line 15.
	EL16 EMR_Bits = 0x01 << 16 //+ Event Mask on line 16.
	EL17 EMR_Bits = 0x01 << 17 //+ Event Mask on line 17.
	EL18 EMR_Bits = 0x01 << 18 //+ Event Mask on line 18.
	EL19 EMR_Bits = 0x01 << 19 //+ Event Mask on line 19.
	EL23 EMR_Bits = 0x01 << 23 //+ Event Mask on line 19.
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
	EL16n = 16
	EL17n = 17
	EL18n = 18
	EL19n = 19
	EL23n = 23
)

const (
	TR0  RTSR_Bits = 0x01 << 0  //+ Rising trigger event configuration bit of line 0.
	TR1  RTSR_Bits = 0x01 << 1  //+ Rising trigger event configuration bit of line 1.
	TR2  RTSR_Bits = 0x01 << 2  //+ Rising trigger event configuration bit of line 2.
	TR3  RTSR_Bits = 0x01 << 3  //+ Rising trigger event configuration bit of line 3.
	TR4  RTSR_Bits = 0x01 << 4  //+ Rising trigger event configuration bit of line 4.
	TR5  RTSR_Bits = 0x01 << 5  //+ Rising trigger event configuration bit of line 5.
	TR6  RTSR_Bits = 0x01 << 6  //+ Rising trigger event configuration bit of line 6.
	TR7  RTSR_Bits = 0x01 << 7  //+ Rising trigger event configuration bit of line 7.
	TR8  RTSR_Bits = 0x01 << 8  //+ Rising trigger event configuration bit of line 8.
	TR9  RTSR_Bits = 0x01 << 9  //+ Rising trigger event configuration bit of line 9.
	TR10 RTSR_Bits = 0x01 << 10 //+ Rising trigger event configuration bit of line 10.
	TR11 RTSR_Bits = 0x01 << 11 //+ Rising trigger event configuration bit of line 11.
	TR12 RTSR_Bits = 0x01 << 12 //+ Rising trigger event configuration bit of line 12.
	TR13 RTSR_Bits = 0x01 << 13 //+ Rising trigger event configuration bit of line 13.
	TR14 RTSR_Bits = 0x01 << 14 //+ Rising trigger event configuration bit of line 14.
	TR15 RTSR_Bits = 0x01 << 15 //+ Rising trigger event configuration bit of line 15.
	TR16 RTSR_Bits = 0x01 << 16 //+ Rising trigger event configuration bit of line 16.
	TR17 RTSR_Bits = 0x01 << 17 //+ Rising trigger event configuration bit of line 17.
	TR18 RTSR_Bits = 0x01 << 18 //+ Rising trigger event configuration bit of line 18.
	TR19 RTSR_Bits = 0x01 << 19 //+ Rising trigger event configuration bit of line 19.
	TR23 RTSR_Bits = 0x01 << 23 //+ Rising trigger event configuration bit of line 23.
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
	TR18n = 18
	TR19n = 19
	TR23n = 23
)

const (
	TF0  FTSR_Bits = 0x01 << 0  //+ Falling trigger event configuration bit of line 0.
	TF1  FTSR_Bits = 0x01 << 1  //+ Falling trigger event configuration bit of line 1.
	TF2  FTSR_Bits = 0x01 << 2  //+ Falling trigger event configuration bit of line 2.
	TF3  FTSR_Bits = 0x01 << 3  //+ Falling trigger event configuration bit of line 3.
	TF4  FTSR_Bits = 0x01 << 4  //+ Falling trigger event configuration bit of line 4.
	TF5  FTSR_Bits = 0x01 << 5  //+ Falling trigger event configuration bit of line 5.
	TF6  FTSR_Bits = 0x01 << 6  //+ Falling trigger event configuration bit of line 6.
	TF7  FTSR_Bits = 0x01 << 7  //+ Falling trigger event configuration bit of line 7.
	TF8  FTSR_Bits = 0x01 << 8  //+ Falling trigger event configuration bit of line 8.
	TF9  FTSR_Bits = 0x01 << 9  //+ Falling trigger event configuration bit of line 9.
	TF10 FTSR_Bits = 0x01 << 10 //+ Falling trigger event configuration bit of line 10.
	TF11 FTSR_Bits = 0x01 << 11 //+ Falling trigger event configuration bit of line 11.
	TF12 FTSR_Bits = 0x01 << 12 //+ Falling trigger event configuration bit of line 12.
	TF13 FTSR_Bits = 0x01 << 13 //+ Falling trigger event configuration bit of line 13.
	TF14 FTSR_Bits = 0x01 << 14 //+ Falling trigger event configuration bit of line 14.
	TF15 FTSR_Bits = 0x01 << 15 //+ Falling trigger event configuration bit of line 15.
	TF16 FTSR_Bits = 0x01 << 16 //+ Falling trigger event configuration bit of line 16.
	TF17 FTSR_Bits = 0x01 << 17 //+ Falling trigger event configuration bit of line 17.
	TF18 FTSR_Bits = 0x01 << 18 //+ Falling trigger event configuration bit of line 18.
	TF19 FTSR_Bits = 0x01 << 19 //+ Falling trigger event configuration bit of line 19.
	TF23 FTSR_Bits = 0x01 << 23 //+ Falling trigger event configuration bit of line 23.
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
	TF18n = 18
	TF19n = 19
	TF23n = 23
)

const (
	SWIER0  SWIER_Bits = 0x01 << 0  //+ Software Interrupt on line 0.
	SWIER1  SWIER_Bits = 0x01 << 1  //+ Software Interrupt on line 1.
	SWIER2  SWIER_Bits = 0x01 << 2  //+ Software Interrupt on line 2.
	SWIER3  SWIER_Bits = 0x01 << 3  //+ Software Interrupt on line 3.
	SWIER4  SWIER_Bits = 0x01 << 4  //+ Software Interrupt on line 4.
	SWIER5  SWIER_Bits = 0x01 << 5  //+ Software Interrupt on line 5.
	SWIER6  SWIER_Bits = 0x01 << 6  //+ Software Interrupt on line 6.
	SWIER7  SWIER_Bits = 0x01 << 7  //+ Software Interrupt on line 7.
	SWIER8  SWIER_Bits = 0x01 << 8  //+ Software Interrupt on line 8.
	SWIER9  SWIER_Bits = 0x01 << 9  //+ Software Interrupt on line 9.
	SWIER10 SWIER_Bits = 0x01 << 10 //+ Software Interrupt on line 10.
	SWIER11 SWIER_Bits = 0x01 << 11 //+ Software Interrupt on line 11.
	SWIER12 SWIER_Bits = 0x01 << 12 //+ Software Interrupt on line 12.
	SWIER13 SWIER_Bits = 0x01 << 13 //+ Software Interrupt on line 13.
	SWIER14 SWIER_Bits = 0x01 << 14 //+ Software Interrupt on line 14.
	SWIER15 SWIER_Bits = 0x01 << 15 //+ Software Interrupt on line 15.
	SWIER16 SWIER_Bits = 0x01 << 16 //+ Software Interrupt on line 16.
	SWIER17 SWIER_Bits = 0x01 << 17 //+ Software Interrupt on line 17.
	SWIER18 SWIER_Bits = 0x01 << 18 //+ Software Interrupt on line 18.
	SWIER19 SWIER_Bits = 0x01 << 19 //+ Software Interrupt on line 19.
	SWIER23 SWIER_Bits = 0x01 << 23 //+ Software Interrupt on line 23.
)

const (
	SWIER0n  = 0
	SWIER1n  = 1
	SWIER2n  = 2
	SWIER3n  = 3
	SWIER4n  = 4
	SWIER5n  = 5
	SWIER6n  = 6
	SWIER7n  = 7
	SWIER8n  = 8
	SWIER9n  = 9
	SWIER10n = 10
	SWIER11n = 11
	SWIER12n = 12
	SWIER13n = 13
	SWIER14n = 14
	SWIER15n = 15
	SWIER16n = 16
	SWIER17n = 17
	SWIER18n = 18
	SWIER19n = 19
	SWIER23n = 23
)

const (
	PR0  PR_Bits = 0x01 << 0  //+ Pending bit for line 0.
	PR1  PR_Bits = 0x01 << 1  //+ Pending bit for line 1.
	PR2  PR_Bits = 0x01 << 2  //+ Pending bit for line 2.
	PR3  PR_Bits = 0x01 << 3  //+ Pending bit for line 3.
	PR4  PR_Bits = 0x01 << 4  //+ Pending bit for line 4.
	PR5  PR_Bits = 0x01 << 5  //+ Pending bit for line 5.
	PR6  PR_Bits = 0x01 << 6  //+ Pending bit for line 6.
	PR7  PR_Bits = 0x01 << 7  //+ Pending bit for line 7.
	PR8  PR_Bits = 0x01 << 8  //+ Pending bit for line 8.
	PR9  PR_Bits = 0x01 << 9  //+ Pending bit for line 9.
	PR10 PR_Bits = 0x01 << 10 //+ Pending bit for line 10.
	PR11 PR_Bits = 0x01 << 11 //+ Pending bit for line 11.
	PR12 PR_Bits = 0x01 << 12 //+ Pending bit for line 12.
	PR13 PR_Bits = 0x01 << 13 //+ Pending bit for line 13.
	PR14 PR_Bits = 0x01 << 14 //+ Pending bit for line 14.
	PR15 PR_Bits = 0x01 << 15 //+ Pending bit for line 15.
	PR16 PR_Bits = 0x01 << 16 //+ Pending bit for line 16.
	PR17 PR_Bits = 0x01 << 17 //+ Pending bit for line 17.
	PR18 PR_Bits = 0x01 << 18 //+ Pending bit for line 18.
	PR19 PR_Bits = 0x01 << 19 //+ Pending bit for line 19.
	PR23 PR_Bits = 0x01 << 23 //+ Pending bit for line 23.
)

const (
	PR0n  = 0
	PR1n  = 1
	PR2n  = 2
	PR3n  = 3
	PR4n  = 4
	PR5n  = 5
	PR6n  = 6
	PR7n  = 7
	PR8n  = 8
	PR9n  = 9
	PR10n = 10
	PR11n = 11
	PR12n = 12
	PR13n = 13
	PR14n = 14
	PR15n = 15
	PR16n = 16
	PR17n = 17
	PR18n = 18
	PR19n = 19
	PR23n = 23
)
