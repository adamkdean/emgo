// +build f411xe

// Peripheral: GPIO_Periph  General Purpose I/O.
// Instances:
//  GPIOA  mmap.GPIOA_BASE
//  GPIOB  mmap.GPIOB_BASE
//  GPIOC  mmap.GPIOC_BASE
//  GPIOD  mmap.GPIOD_BASE
//  GPIOE  mmap.GPIOE_BASE
//  GPIOF  mmap.GPIOF_BASE
//  GPIOG  mmap.GPIOG_BASE
//  GPIOH  mmap.GPIOH_BASE
//  GPIOI  mmap.GPIOI_BASE
//  GPIOJ  mmap.GPIOJ_BASE
//  GPIOK  mmap.GPIOK_BASE
// Registers:
//  0x00 32  MODER   Port mode register.
//  0x04 32  OTYPER  Port output type register.
//  0x08 32  OSPEEDR Port output speed register.
//  0x0C 32  PUPDR   Port pull-up/pull-down register.
//  0x10 32  IDR     Port input data register.
//  0x14 32  ODR     Port output data register.
//  0x18 16  BSRRL   Port bit set/reset low register.
//  0x1A 16  BSRRH   Port bit set/reset high register.
//  0x1C 32  LCKR    Port configuration lock register.
//  0x20 32  AFR[2]  Alternate function registers.
// Import:
//  stm32/o/f411xe/mmap
package gpio

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MODER0    MODER = 0x03 << 0 //+
	MODER0_0  MODER = 0x01 << 0
	MODER0_1  MODER = 0x02 << 0
	MODER1    MODER = 0x03 << 2 //+
	MODER1_0  MODER = 0x01 << 2
	MODER1_1  MODER = 0x02 << 2
	MODER2    MODER = 0x03 << 4 //+
	MODER2_0  MODER = 0x01 << 4
	MODER2_1  MODER = 0x02 << 4
	MODER3    MODER = 0x03 << 6 //+
	MODER3_0  MODER = 0x01 << 6
	MODER3_1  MODER = 0x02 << 6
	MODER4    MODER = 0x03 << 8 //+
	MODER4_0  MODER = 0x01 << 8
	MODER4_1  MODER = 0x02 << 8
	MODER5    MODER = 0x03 << 10 //+
	MODER5_0  MODER = 0x01 << 10
	MODER5_1  MODER = 0x02 << 10
	MODER6    MODER = 0x03 << 12 //+
	MODER6_0  MODER = 0x01 << 12
	MODER6_1  MODER = 0x02 << 12
	MODER7    MODER = 0x03 << 14 //+
	MODER7_0  MODER = 0x01 << 14
	MODER7_1  MODER = 0x02 << 14
	MODER8    MODER = 0x03 << 16 //+
	MODER8_0  MODER = 0x01 << 16
	MODER8_1  MODER = 0x02 << 16
	MODER9    MODER = 0x03 << 18 //+
	MODER9_0  MODER = 0x01 << 18
	MODER9_1  MODER = 0x02 << 18
	MODER10   MODER = 0x03 << 20 //+
	MODER10_0 MODER = 0x01 << 20
	MODER10_1 MODER = 0x02 << 20
	MODER11   MODER = 0x03 << 22 //+
	MODER11_0 MODER = 0x01 << 22
	MODER11_1 MODER = 0x02 << 22
	MODER12   MODER = 0x03 << 24 //+
	MODER12_0 MODER = 0x01 << 24
	MODER12_1 MODER = 0x02 << 24
	MODER13   MODER = 0x03 << 26 //+
	MODER13_0 MODER = 0x01 << 26
	MODER13_1 MODER = 0x02 << 26
	MODER14   MODER = 0x03 << 28 //+
	MODER14_0 MODER = 0x01 << 28
	MODER14_1 MODER = 0x02 << 28
	MODER15   MODER = 0x03 << 30 //+
	MODER15_0 MODER = 0x01 << 30
	MODER15_1 MODER = 0x02 << 30
)

const (
	MODER0n  = 0
	MODER1n  = 2
	MODER2n  = 4
	MODER3n  = 6
	MODER4n  = 8
	MODER5n  = 10
	MODER6n  = 12
	MODER7n  = 14
	MODER8n  = 16
	MODER9n  = 18
	MODER10n = 20
	MODER11n = 22
	MODER12n = 24
	MODER13n = 26
	MODER14n = 28
	MODER15n = 30
)

const (
	OT_0  OTYPER = 0x01 << 0  //+
	OT_1  OTYPER = 0x01 << 1  //+
	OT_2  OTYPER = 0x01 << 2  //+
	OT_3  OTYPER = 0x01 << 3  //+
	OT_4  OTYPER = 0x01 << 4  //+
	OT_5  OTYPER = 0x01 << 5  //+
	OT_6  OTYPER = 0x01 << 6  //+
	OT_7  OTYPER = 0x01 << 7  //+
	OT_8  OTYPER = 0x01 << 8  //+
	OT_9  OTYPER = 0x01 << 9  //+
	OT_10 OTYPER = 0x01 << 10 //+
	OT_11 OTYPER = 0x01 << 11 //+
	OT_12 OTYPER = 0x01 << 12 //+
	OT_13 OTYPER = 0x01 << 13 //+
	OT_14 OTYPER = 0x01 << 14 //+
	OT_15 OTYPER = 0x01 << 15 //+
)

const (
	OT_0n  = 0
	OT_1n  = 1
	OT_2n  = 2
	OT_3n  = 3
	OT_4n  = 4
	OT_5n  = 5
	OT_6n  = 6
	OT_7n  = 7
	OT_8n  = 8
	OT_9n  = 9
	OT_10n = 10
	OT_11n = 11
	OT_12n = 12
	OT_13n = 13
	OT_14n = 14
	OT_15n = 15
)

const (
	PUPDR0    PUPDR = 0x03 << 0 //+
	PUPDR0_0  PUPDR = 0x01 << 0
	PUPDR0_1  PUPDR = 0x02 << 0
	PUPDR1    PUPDR = 0x03 << 2 //+
	PUPDR1_0  PUPDR = 0x01 << 2
	PUPDR1_1  PUPDR = 0x02 << 2
	PUPDR2    PUPDR = 0x03 << 4 //+
	PUPDR2_0  PUPDR = 0x01 << 4
	PUPDR2_1  PUPDR = 0x02 << 4
	PUPDR3    PUPDR = 0x03 << 6 //+
	PUPDR3_0  PUPDR = 0x01 << 6
	PUPDR3_1  PUPDR = 0x02 << 6
	PUPDR4    PUPDR = 0x03 << 8 //+
	PUPDR4_0  PUPDR = 0x01 << 8
	PUPDR4_1  PUPDR = 0x02 << 8
	PUPDR5    PUPDR = 0x03 << 10 //+
	PUPDR5_0  PUPDR = 0x01 << 10
	PUPDR5_1  PUPDR = 0x02 << 10
	PUPDR6    PUPDR = 0x03 << 12 //+
	PUPDR6_0  PUPDR = 0x01 << 12
	PUPDR6_1  PUPDR = 0x02 << 12
	PUPDR7    PUPDR = 0x03 << 14 //+
	PUPDR7_0  PUPDR = 0x01 << 14
	PUPDR7_1  PUPDR = 0x02 << 14
	PUPDR8    PUPDR = 0x03 << 16 //+
	PUPDR8_0  PUPDR = 0x01 << 16
	PUPDR8_1  PUPDR = 0x02 << 16
	PUPDR9    PUPDR = 0x03 << 18 //+
	PUPDR9_0  PUPDR = 0x01 << 18
	PUPDR9_1  PUPDR = 0x02 << 18
	PUPDR10   PUPDR = 0x03 << 20 //+
	PUPDR10_0 PUPDR = 0x01 << 20
	PUPDR10_1 PUPDR = 0x02 << 20
	PUPDR11   PUPDR = 0x03 << 22 //+
	PUPDR11_0 PUPDR = 0x01 << 22
	PUPDR11_1 PUPDR = 0x02 << 22
	PUPDR12   PUPDR = 0x03 << 24 //+
	PUPDR12_0 PUPDR = 0x01 << 24
	PUPDR12_1 PUPDR = 0x02 << 24
	PUPDR13   PUPDR = 0x03 << 26 //+
	PUPDR13_0 PUPDR = 0x01 << 26
	PUPDR13_1 PUPDR = 0x02 << 26
	PUPDR14   PUPDR = 0x03 << 28 //+
	PUPDR14_0 PUPDR = 0x01 << 28
	PUPDR14_1 PUPDR = 0x02 << 28
	PUPDR15   PUPDR = 0x03 << 30 //+
	PUPDR15_0 PUPDR = 0x01 << 30
	PUPDR15_1 PUPDR = 0x02 << 30
)

const (
	PUPDR0n  = 0
	PUPDR1n  = 2
	PUPDR2n  = 4
	PUPDR3n  = 6
	PUPDR4n  = 8
	PUPDR5n  = 10
	PUPDR6n  = 12
	PUPDR7n  = 14
	PUPDR8n  = 16
	PUPDR9n  = 18
	PUPDR10n = 20
	PUPDR11n = 22
	PUPDR12n = 24
	PUPDR13n = 26
	PUPDR14n = 28
	PUPDR15n = 30
)

const (
	IDR_0  IDR = 0x01 << 0  //+
	IDR_1  IDR = 0x01 << 1  //+
	IDR_2  IDR = 0x01 << 2  //+
	IDR_3  IDR = 0x01 << 3  //+
	IDR_4  IDR = 0x01 << 4  //+
	IDR_5  IDR = 0x01 << 5  //+
	IDR_6  IDR = 0x01 << 6  //+
	IDR_7  IDR = 0x01 << 7  //+
	IDR_8  IDR = 0x01 << 8  //+
	IDR_9  IDR = 0x01 << 9  //+
	IDR_10 IDR = 0x01 << 10 //+
	IDR_11 IDR = 0x01 << 11 //+
	IDR_12 IDR = 0x01 << 12 //+
	IDR_13 IDR = 0x01 << 13 //+
	IDR_14 IDR = 0x01 << 14 //+
	IDR_15 IDR = 0x01 << 15 //+
)

const (
	IDR_0n  = 0
	IDR_1n  = 1
	IDR_2n  = 2
	IDR_3n  = 3
	IDR_4n  = 4
	IDR_5n  = 5
	IDR_6n  = 6
	IDR_7n  = 7
	IDR_8n  = 8
	IDR_9n  = 9
	IDR_10n = 10
	IDR_11n = 11
	IDR_12n = 12
	IDR_13n = 13
	IDR_14n = 14
	IDR_15n = 15
)

const (
	ODR_0  ODR = 0x01 << 0  //+
	ODR_1  ODR = 0x01 << 1  //+
	ODR_2  ODR = 0x01 << 2  //+
	ODR_3  ODR = 0x01 << 3  //+
	ODR_4  ODR = 0x01 << 4  //+
	ODR_5  ODR = 0x01 << 5  //+
	ODR_6  ODR = 0x01 << 6  //+
	ODR_7  ODR = 0x01 << 7  //+
	ODR_8  ODR = 0x01 << 8  //+
	ODR_9  ODR = 0x01 << 9  //+
	ODR_10 ODR = 0x01 << 10 //+
	ODR_11 ODR = 0x01 << 11 //+
	ODR_12 ODR = 0x01 << 12 //+
	ODR_13 ODR = 0x01 << 13 //+
	ODR_14 ODR = 0x01 << 14 //+
	ODR_15 ODR = 0x01 << 15 //+
)

const (
	ODR_0n  = 0
	ODR_1n  = 1
	ODR_2n  = 2
	ODR_3n  = 3
	ODR_4n  = 4
	ODR_5n  = 5
	ODR_6n  = 6
	ODR_7n  = 7
	ODR_8n  = 8
	ODR_9n  = 9
	ODR_10n = 10
	ODR_11n = 11
	ODR_12n = 12
	ODR_13n = 13
	ODR_14n = 14
	ODR_15n = 15
)
