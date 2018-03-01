// +build l476xx

// Peripheral: DBGMCU_Periph  Debug MCU.
// Instances:
//  DBGMCU  mmap.DBGMCU_BASE
// Registers:
//  0x00 32  IDCODE   MCU device ID code.
//  0x04 32  CR       Debug MCU configuration register.
//  0x08 32  APB1FZR1 Debug MCU APB1 freeze register 1.
//  0x0C 32  APB1FZR2 Debug MCU APB1 freeze register 2.
//  0x10 32  APB2FZ   Debug MCU APB2 freeze register.
// Import:
//  stm32/o/l476xx/mmap
package dbgmcu

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	DEV_ID IDCODE = 0xFFF << 0   //+
	REV_ID IDCODE = 0xFFFF << 16 //+
)

const (
	DEV_IDn = 0
	REV_IDn = 16
)

const (
	DBG_SLEEP   CR = 0x01 << 0 //+
	DBG_STOP    CR = 0x01 << 1 //+
	DBG_STANDBY CR = 0x01 << 2 //+
	TRACE_IOEN  CR = 0x01 << 5 //+
	TRACE_MODE  CR = 0x03 << 6 //+
)

const (
	DBG_SLEEPn   = 0
	DBG_STOPn    = 1
	DBG_STANDBYn = 2
	TRACE_IOENn  = 5
	TRACE_MODEn  = 6
)

const (
	DBG_TIM2_STOP   APB1FZR1 = 0x01 << 0  //+
	DBG_TIM3_STOP   APB1FZR1 = 0x01 << 1  //+
	DBG_TIM4_STOP   APB1FZR1 = 0x01 << 2  //+
	DBG_TIM5_STOP   APB1FZR1 = 0x01 << 3  //+
	DBG_TIM6_STOP   APB1FZR1 = 0x01 << 4  //+
	DBG_TIM7_STOP   APB1FZR1 = 0x01 << 5  //+
	DBG_RTC_STOP    APB1FZR1 = 0x01 << 10 //+
	DBG_WWDG_STOP   APB1FZR1 = 0x01 << 11 //+
	DBG_IWDG_STOP   APB1FZR1 = 0x01 << 12 //+
	DBG_I2C1_STOP   APB1FZR1 = 0x01 << 21 //+
	DBG_I2C2_STOP   APB1FZR1 = 0x01 << 22 //+
	DBG_I2C3_STOP   APB1FZR1 = 0x01 << 23 //+
	DBG_CAN_STOP    APB1FZR1 = 0x01 << 25 //+
	DBG_LPTIM1_STOP APB1FZR1 = 0x01 << 31 //+
)

const (
	DBG_TIM2_STOPn   = 0
	DBG_TIM3_STOPn   = 1
	DBG_TIM4_STOPn   = 2
	DBG_TIM5_STOPn   = 3
	DBG_TIM6_STOPn   = 4
	DBG_TIM7_STOPn   = 5
	DBG_RTC_STOPn    = 10
	DBG_WWDG_STOPn   = 11
	DBG_IWDG_STOPn   = 12
	DBG_I2C1_STOPn   = 21
	DBG_I2C2_STOPn   = 22
	DBG_I2C3_STOPn   = 23
	DBG_CAN_STOPn    = 25
	DBG_LPTIM1_STOPn = 31
)

const (
	DBG_LPTIM2_STOP APB1FZR2 = 0x01 << 5 //+
)

const (
	DBG_LPTIM2_STOPn = 5
)

const (
	DBG_TIM1_STOP  APB2FZ = 0x01 << 11 //+
	DBG_TIM8_STOP  APB2FZ = 0x01 << 13 //+
	DBG_TIM15_STOP APB2FZ = 0x01 << 16 //+
	DBG_TIM16_STOP APB2FZ = 0x01 << 17 //+
	DBG_TIM17_STOP APB2FZ = 0x01 << 18 //+
)

const (
	DBG_TIM1_STOPn  = 11
	DBG_TIM8_STOPn  = 13
	DBG_TIM15_STOPn = 16
	DBG_TIM16_STOPn = 17
	DBG_TIM17_STOPn = 18
)
