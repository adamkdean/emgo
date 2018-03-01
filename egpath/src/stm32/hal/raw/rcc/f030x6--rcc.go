// +build f030x6

// Peripheral: RCC_Periph  Reset and Clock Control.
// Instances:
//  RCC  mmap.RCC_BASE
// Registers:
//  0x00 32  CR       Clock control register.
//  0x04 32  CFGR     Clock configuration register.
//  0x08 32  CIR      Clock interrupt register.
//  0x0C 32  APB2RSTR APB2 peripheral reset register.
//  0x10 32  APB1RSTR APB1 peripheral reset register.
//  0x14 32  AHBENR   AHB peripheral clock register.
//  0x18 32  APB2ENR  APB2 peripheral clock enable register.
//  0x1C 32  APB1ENR  APB1 peripheral clock enable register.
//  0x20 32  BDCR     Backup domain control register.
//  0x24 32  CSR      Clock control & status register.
//  0x28 32  AHBRSTR  AHB peripheral reset register.
//  0x2C 32  CFGR2    Clock configuration register 2.
//  0x30 32  CFGR3    Clock configuration register 3.
//  0x34 32  CR2      Clock control register 2.
// Import:
//  stm32/o/f030x6/mmap
package rcc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	HSION   CR = 0x01 << 0  //+ Internal High Speed clock enable.
	HSIRDY  CR = 0x01 << 1  //+ Internal High Speed clock ready flag.
	HSITRIM CR = 0x1F << 3  //+ Internal High Speed clock trimming.
	HSICAL  CR = 0xFF << 8  //+ Internal High Speed clock Calibration.
	HSEON   CR = 0x01 << 16 //+ External High Speed clock enable.
	HSERDY  CR = 0x01 << 17 //+ External High Speed clock ready flag.
	HSEBYP  CR = 0x01 << 18 //+ External High Speed clock Bypass.
	CSSON   CR = 0x01 << 19 //+ Clock Security System enable.
	PLLON   CR = 0x01 << 24 //+ PLL enable.
	PLLRDY  CR = 0x01 << 25 //+ PLL clock ready flag.
)

const (
	HSIONn   = 0
	HSIRDYn  = 1
	HSITRIMn = 3
	HSICALn  = 8
	HSEONn   = 16
	HSERDYn  = 17
	HSEBYPn  = 18
	CSSONn   = 19
	PLLONn   = 24
	PLLRDYn  = 25
)

const (
	SW                       CFGR = 0x03 << 0  //+ SW[1:0] bits (System clock Switch).
	SW_HSI                   CFGR = 0x00 << 0  //  HSI selected as system clock.
	SW_HSE                   CFGR = 0x01 << 0  //  HSE selected as system clock.
	SW_PLL                   CFGR = 0x02 << 0  //  PLL selected as system clock.
	SWS                      CFGR = 0x03 << 2  //+ SWS[1:0] bits (System Clock Switch Status).
	SWS_HSI                  CFGR = 0x00 << 2  //  HSI oscillator used as system clock.
	SWS_HSE                  CFGR = 0x01 << 2  //  HSE oscillator used as system clock.
	SWS_PLL                  CFGR = 0x02 << 2  //  PLL used as system clock.
	HPRE                     CFGR = 0x0F << 4  //+ HPRE[3:0] bits (AHB prescaler).
	HPRE_DIV1                CFGR = 0x00 << 4  //  SYSCLK not divided.
	HPRE_DIV2                CFGR = 0x08 << 4  //  SYSCLK divided by 2.
	HPRE_DIV4                CFGR = 0x09 << 4  //  SYSCLK divided by 4.
	HPRE_DIV8                CFGR = 0x0A << 4  //  SYSCLK divided by 8.
	HPRE_DIV16               CFGR = 0x0B << 4  //  SYSCLK divided by 16.
	HPRE_DIV64               CFGR = 0x0C << 4  //  SYSCLK divided by 64.
	HPRE_DIV128              CFGR = 0x0D << 4  //  SYSCLK divided by 128.
	HPRE_DIV256              CFGR = 0x0E << 4  //  SYSCLK divided by 256.
	HPRE_DIV512              CFGR = 0x0F << 4  //  SYSCLK divided by 512.
	PPRE                     CFGR = 0x07 << 8  //+ PRE[2:0] bits (APB prescaler).
	PPRE_DIV1                CFGR = 0x00 << 8  //  HCLK not divided.
	PPRE_DIV2                CFGR = 0x04 << 8  //  HCLK divided by 2.
	PPRE_DIV4                CFGR = 0x05 << 8  //  HCLK divided by 4.
	PPRE_DIV8                CFGR = 0x06 << 8  //  HCLK divided by 8.
	PPRE_DIV16               CFGR = 0x07 << 8  //  HCLK divided by 16.
	ADCPRE                   CFGR = 0x01 << 14 //+ ADCPRE bit (ADC prescaler).
	ADCPRE_DIV2              CFGR = 0x00 << 14 //  PCLK divided by 2.
	ADCPRE_DIV4              CFGR = 0x01 << 14 //  PCLK divided by 4.
	PLLSRC                   CFGR = 0x01 << 16 //+ PLL entry clock source.
	PLLSRC_HSI_DIV2          CFGR = 0x00 << 16 //  HSI clock divided by 2 selected as PLL entry clock source.
	PLLSRC_HSE_PREDIV        CFGR = 0x01 << 16 //  HSE/PREDIV clock selected as PLL entry clock source.
	PLLXTPRE                 CFGR = 0x01 << 17 //+ HSE divider for PLL entry.
	PLLXTPRE_HSE_PREDIV_DIV1 CFGR = 0x00 << 17 //  HSE/PREDIV clock not divided for PLL entry.
	PLLXTPRE_HSE_PREDIV_DIV2 CFGR = 0x01 << 17 //  HSE/PREDIV clock divided by 2 for PLL entry.
	PLLMUL                   CFGR = 0x0F << 18 //+ PLLMUL[3:0] bits (PLL multiplication factor).
	PLLMUL2                  CFGR = 0x00 << 18 //  PLL input clock*2.
	PLLMUL3                  CFGR = 0x01 << 18 //  PLL input clock*3.
	PLLMUL4                  CFGR = 0x02 << 18 //  PLL input clock*4.
	PLLMUL5                  CFGR = 0x03 << 18 //  PLL input clock*5.
	PLLMUL6                  CFGR = 0x04 << 18 //  PLL input clock*6.
	PLLMUL7                  CFGR = 0x05 << 18 //  PLL input clock*7.
	PLLMUL8                  CFGR = 0x06 << 18 //  PLL input clock*8.
	PLLMUL9                  CFGR = 0x07 << 18 //  PLL input clock*9.
	PLLMUL10                 CFGR = 0x08 << 18 //  PLL input clock10.
	PLLMUL11                 CFGR = 0x09 << 18 //  PLL input clock*11.
	PLLMUL12                 CFGR = 0x0A << 18 //  PLL input clock*12.
	PLLMUL13                 CFGR = 0x0B << 18 //  PLL input clock*13.
	PLLMUL14                 CFGR = 0x0C << 18 //  PLL input clock*14.
	PLLMUL15                 CFGR = 0x0D << 18 //  PLL input clock*15.
	PLLMUL16                 CFGR = 0x0E << 18 //  PLL input clock*16.
	MCO                      CFGR = 0x0F << 24 //+ MCO[3:0] bits (Microcontroller Clock Output).
	MCO_NOCLOCK              CFGR = 0x00 << 24 //  No clock.
	MCO_HSI14                CFGR = 0x01 << 24 //  HSI14 clock selected as MCO source.
	MCO_LSI                  CFGR = 0x02 << 24 //  LSI clock selected as MCO source.
	MCO_LSE                  CFGR = 0x03 << 24 //  LSE clock selected as MCO source.
	MCO_SYSCLK               CFGR = 0x04 << 24 //  System clock selected as MCO source.
	MCO_HSI                  CFGR = 0x05 << 24 //  HSI clock selected as MCO source.
	MCO_HSE                  CFGR = 0x06 << 24 //  HSE clock selected as MCO source.
	MCO_PLL                  CFGR = 0x07 << 24 //  PLL clock divided by 2 selected as MCO source.
	MCOPRE                   CFGR = 0x07 << 28 //+ MCO prescaler.
	MCOPRE_DIV1              CFGR = 0x00 << 28 //  MCO is divided by 1.
	MCOPRE_DIV2              CFGR = 0x01 << 28 //  MCO is divided by 2.
	MCOPRE_DIV4              CFGR = 0x02 << 28 //  MCO is divided by 4.
	MCOPRE_DIV8              CFGR = 0x03 << 28 //  MCO is divided by 8.
	MCOPRE_DIV16             CFGR = 0x04 << 28 //  MCO is divided by 16.
	MCOPRE_DIV32             CFGR = 0x05 << 28 //  MCO is divided by 32.
	MCOPRE_DIV64             CFGR = 0x06 << 28 //  MCO is divided by 64.
	MCOPRE_DIV128            CFGR = 0x07 << 28 //  MCO is divided by 128.
	PLLNODIV                 CFGR = 0x01 << 31 //+ PLL is not divided to MCO.
)

const (
	SWn       = 0
	SWSn      = 2
	HPREn     = 4
	PPREn     = 8
	ADCPREn   = 14
	PLLSRCn   = 16
	PLLXTPREn = 17
	PLLMULn   = 18
	MCOn      = 24
	MCOPREn   = 28
	PLLNODIVn = 31
)

const (
	LSIRDYF    CIR = 0x01 << 0  //+ LSI Ready Interrupt flag.
	LSERDYF    CIR = 0x01 << 1  //+ LSE Ready Interrupt flag.
	HSIRDYF    CIR = 0x01 << 2  //+ HSI Ready Interrupt flag.
	HSERDYF    CIR = 0x01 << 3  //+ HSE Ready Interrupt flag.
	PLLRDYF    CIR = 0x01 << 4  //+ PLL Ready Interrupt flag.
	HSI14RDYF  CIR = 0x01 << 5  //+ HSI14 Ready Interrupt flag.
	CSSF       CIR = 0x01 << 7  //+ Clock Security System Interrupt flag.
	LSIRDYIE   CIR = 0x01 << 8  //+ LSI Ready Interrupt Enable.
	LSERDYIE   CIR = 0x01 << 9  //+ LSE Ready Interrupt Enable.
	HSIRDYIE   CIR = 0x01 << 10 //+ HSI Ready Interrupt Enable.
	HSERDYIE   CIR = 0x01 << 11 //+ HSE Ready Interrupt Enable.
	PLLRDYIE   CIR = 0x01 << 12 //+ PLL Ready Interrupt Enable.
	HSI14RDYIE CIR = 0x01 << 13 //+ HSI14 Ready Interrupt Enable.
	LSIRDYC    CIR = 0x01 << 16 //+ LSI Ready Interrupt Clear.
	LSERDYC    CIR = 0x01 << 17 //+ LSE Ready Interrupt Clear.
	HSIRDYC    CIR = 0x01 << 18 //+ HSI Ready Interrupt Clear.
	HSERDYC    CIR = 0x01 << 19 //+ HSE Ready Interrupt Clear.
	PLLRDYC    CIR = 0x01 << 20 //+ PLL Ready Interrupt Clear.
	HSI14RDYC  CIR = 0x01 << 21 //+ HSI14 Ready Interrupt Clear.
	CSSC       CIR = 0x01 << 23 //+ Clock Security System Interrupt Clear.
)

const (
	LSIRDYFn    = 0
	LSERDYFn    = 1
	HSIRDYFn    = 2
	HSERDYFn    = 3
	PLLRDYFn    = 4
	HSI14RDYFn  = 5
	CSSFn       = 7
	LSIRDYIEn   = 8
	LSERDYIEn   = 9
	HSIRDYIEn   = 10
	HSERDYIEn   = 11
	PLLRDYIEn   = 12
	HSI14RDYIEn = 13
	LSIRDYCn    = 16
	LSERDYCn    = 17
	HSIRDYCn    = 18
	HSERDYCn    = 19
	PLLRDYCn    = 20
	HSI14RDYCn  = 21
	CSSCn       = 23
)

const (
	SYSCFGRST APB2RSTR = 0x01 << 0  //+ SYSCFG reset.
	ADCRST    APB2RSTR = 0x01 << 9  //+ ADC reset.
	TIM1RST   APB2RSTR = 0x01 << 11 //+ TIM1 reset.
	SPI1RST   APB2RSTR = 0x01 << 12 //+ SPI1 reset.
	USART1RST APB2RSTR = 0x01 << 14 //+ USART1 reset.
	TIM16RST  APB2RSTR = 0x01 << 17 //+ TIM16 reset.
	TIM17RST  APB2RSTR = 0x01 << 18 //+ TIM17 reset.
	DBGMCURST APB2RSTR = 0x01 << 22 //+ DBGMCU reset.
)

const (
	SYSCFGRSTn = 0
	ADCRSTn    = 9
	TIM1RSTn   = 11
	SPI1RSTn   = 12
	USART1RSTn = 14
	TIM16RSTn  = 17
	TIM17RSTn  = 18
	DBGMCURSTn = 22
)

const (
	TIM3RST  APB1RSTR = 0x01 << 1  //+ Timer 3 reset.
	TIM14RST APB1RSTR = 0x01 << 8  //+ Timer 14 reset.
	WWDGRST  APB1RSTR = 0x01 << 11 //+ Window Watchdog reset.
	I2C1RST  APB1RSTR = 0x01 << 21 //+ I2C 1 reset.
	PWRRST   APB1RSTR = 0x01 << 28 //+ PWR reset.
)

const (
	TIM3RSTn  = 1
	TIM14RSTn = 8
	WWDGRSTn  = 11
	I2C1RSTn  = 21
	PWRRSTn   = 28
)

const (
	DMAEN   AHBENR = 0x01 << 0  //+ DMA1 clock enable.
	SRAMEN  AHBENR = 0x01 << 2  //+ SRAM interface clock enable.
	FLITFEN AHBENR = 0x01 << 4  //+ FLITF clock enable.
	CRCEN   AHBENR = 0x01 << 6  //+ CRC clock enable.
	GPIOAEN AHBENR = 0x01 << 17 //+ GPIOA clock enable.
	GPIOBEN AHBENR = 0x01 << 18 //+ GPIOB clock enable.
	GPIOCEN AHBENR = 0x01 << 19 //+ GPIOC clock enable.
	GPIODEN AHBENR = 0x01 << 20 //+ GPIOD clock enable.
	GPIOFEN AHBENR = 0x01 << 22 //+ GPIOF clock enable.
)

const (
	DMAENn   = 0
	SRAMENn  = 2
	FLITFENn = 4
	CRCENn   = 6
	GPIOAENn = 17
	GPIOBENn = 18
	GPIOCENn = 19
	GPIODENn = 20
	GPIOFENn = 22
)

const (
	SYSCFGCOMPEN APB2ENR = 0x01 << 0  //+ SYSCFG and comparator clock enable.
	ADCEN        APB2ENR = 0x01 << 9  //+ ADC1 clock enable.
	TIM1EN       APB2ENR = 0x01 << 11 //+ TIM1 clock enable.
	SPI1EN       APB2ENR = 0x01 << 12 //+ SPI1 clock enable.
	USART1EN     APB2ENR = 0x01 << 14 //+ USART1 clock enable.
	TIM16EN      APB2ENR = 0x01 << 17 //+ TIM16 clock enable.
	TIM17EN      APB2ENR = 0x01 << 18 //+ TIM17 clock enable.
	DBGMCUEN     APB2ENR = 0x01 << 22 //+ DBGMCU clock enable.
)

const (
	SYSCFGCOMPENn = 0
	ADCENn        = 9
	TIM1ENn       = 11
	SPI1ENn       = 12
	USART1ENn     = 14
	TIM16ENn      = 17
	TIM17ENn      = 18
	DBGMCUENn     = 22
)

const (
	TIM3EN  APB1ENR = 0x01 << 1  //+ Timer 3 clock enable.
	TIM14EN APB1ENR = 0x01 << 8  //+ Timer 14 clock enable.
	WWDGEN  APB1ENR = 0x01 << 11 //+ Window Watchdog clock enable.
	I2C1EN  APB1ENR = 0x01 << 21 //+ I2C1 clock enable.
	PWREN   APB1ENR = 0x01 << 28 //+ PWR clock enable.
)

const (
	TIM3ENn  = 1
	TIM14ENn = 8
	WWDGENn  = 11
	I2C1ENn  = 21
	PWRENn   = 28
)

const (
	LSEON          BDCR = 0x01 << 0  //+ External Low Speed oscillator enable.
	LSERDY         BDCR = 0x01 << 1  //+ External Low Speed oscillator Ready.
	LSEBYP         BDCR = 0x01 << 2  //+ External Low Speed oscillator Bypass.
	LSEDRV         BDCR = 0x03 << 3  //+ LSEDRV[1:0] bits (LSE Osc. drive capability).
	RTCSEL         BDCR = 0x03 << 8  //+ RTCSEL[1:0] bits (RTC clock source selection).
	RTCSEL_NOCLOCK BDCR = 0x00 << 8  //  No clock.
	RTCSEL_LSE     BDCR = 0x01 << 8  //  LSE oscillator clock used as RTC clock.
	RTCSEL_LSI     BDCR = 0x02 << 8  //  LSI oscillator clock used as RTC clock.
	RTCSEL_HSE     BDCR = 0x03 << 8  //  HSE oscillator clock divided by 128 used as RTC clock.
	RTCEN          BDCR = 0x01 << 15 //+ RTC clock enable.
	BDRST          BDCR = 0x01 << 16 //+ Backup domain software reset.
)

const (
	LSEONn  = 0
	LSERDYn = 1
	LSEBYPn = 2
	LSEDRVn = 3
	RTCSELn = 8
	RTCENn  = 15
	BDRSTn  = 16
)

const (
	LSION      CSR = 0x01 << 0  //+ Internal Low Speed oscillator enable.
	LSIRDY     CSR = 0x01 << 1  //+ Internal Low Speed oscillator Ready.
	V18PWRRSTF CSR = 0x01 << 23 //+ V1.8 power domain reset flag.
	RMVF       CSR = 0x01 << 24 //+ Remove reset flag.
	OBLRSTF    CSR = 0x01 << 25 //+ OBL reset flag.
	PINRSTF    CSR = 0x01 << 26 //+ PIN reset flag.
	PORRSTF    CSR = 0x01 << 27 //+ POR/PDR reset flag.
	SFTRSTF    CSR = 0x01 << 28 //+ Software Reset flag.
	IWDGRSTF   CSR = 0x01 << 29 //+ Independent Watchdog reset flag.
	WWDGRSTF   CSR = 0x01 << 30 //+ Window watchdog reset flag.
	LPWRRSTF   CSR = 0x01 << 31 //+ Low-Power reset flag.
)

const (
	LSIONn      = 0
	LSIRDYn     = 1
	V18PWRRSTFn = 23
	RMVFn       = 24
	OBLRSTFn    = 25
	PINRSTFn    = 26
	PORRSTFn    = 27
	SFTRSTFn    = 28
	IWDGRSTFn   = 29
	WWDGRSTFn   = 30
	LPWRRSTFn   = 31
)

const (
	GPIOARST AHBRSTR = 0x01 << 17 //+ GPIOA reset.
	GPIOBRST AHBRSTR = 0x01 << 18 //+ GPIOB reset.
	GPIOCRST AHBRSTR = 0x01 << 19 //+ GPIOC reset.
	GPIODRST AHBRSTR = 0x01 << 20 //+ GPIOD reset.
	GPIOFRST AHBRSTR = 0x01 << 22 //+ GPIOF reset.
)

const (
	GPIOARSTn = 17
	GPIOBRSTn = 18
	GPIOCRSTn = 19
	GPIODRSTn = 20
	GPIOFRSTn = 22
)

const (
	PREDIV       CFGR2 = 0x0F << 0 //+ PREDIV[3:0] bits.
	PREDIV_DIV1  CFGR2 = 0x00 << 0 //  PREDIV input clock not divided.
	PREDIV_DIV2  CFGR2 = 0x01 << 0 //  PREDIV input clock divided by 2.
	PREDIV_DIV3  CFGR2 = 0x02 << 0 //  PREDIV input clock divided by 3.
	PREDIV_DIV4  CFGR2 = 0x03 << 0 //  PREDIV input clock divided by 4.
	PREDIV_DIV5  CFGR2 = 0x04 << 0 //  PREDIV input clock divided by 5.
	PREDIV_DIV6  CFGR2 = 0x05 << 0 //  PREDIV input clock divided by 6.
	PREDIV_DIV7  CFGR2 = 0x06 << 0 //  PREDIV input clock divided by 7.
	PREDIV_DIV8  CFGR2 = 0x07 << 0 //  PREDIV input clock divided by 8.
	PREDIV_DIV9  CFGR2 = 0x08 << 0 //  PREDIV input clock divided by 9.
	PREDIV_DIV10 CFGR2 = 0x09 << 0 //  PREDIV input clock divided by 10.
	PREDIV_DIV11 CFGR2 = 0x0A << 0 //  PREDIV input clock divided by 11.
	PREDIV_DIV12 CFGR2 = 0x0B << 0 //  PREDIV input clock divided by 12.
	PREDIV_DIV13 CFGR2 = 0x0C << 0 //  PREDIV input clock divided by 13.
	PREDIV_DIV14 CFGR2 = 0x0D << 0 //  PREDIV input clock divided by 14.
	PREDIV_DIV15 CFGR2 = 0x0E << 0 //  PREDIV input clock divided by 15.
	PREDIV_DIV16 CFGR2 = 0x0F << 0 //  PREDIV input clock divided by 16.
)

const (
	PREDIVn = 0
)

const (
	USART1SW        CFGR3 = 0x03 << 0 //+ USART1SW[1:0] bits.
	USART1SW_PCLK   CFGR3 = 0x00 << 0 //  PCLK clock used as USART1 clock source.
	USART1SW_SYSCLK CFGR3 = 0x01 << 0 //  System clock selected as USART1 clock source.
	USART1SW_LSE    CFGR3 = 0x02 << 0 //  LSE oscillator clock used as USART1 clock source.
	USART1SW_HSI    CFGR3 = 0x03 << 0 //  HSI oscillator clock used as USART1 clock source.
	I2C1SW          CFGR3 = 0x01 << 4 //+ I2C1SW bits.
	I2C1SW_HSI      CFGR3 = 0x00 << 4 //  HSI oscillator clock used as I2C1 clock source.
	I2C1SW_SYSCLK   CFGR3 = 0x01 << 4 //  System clock selected as I2C1 clock source.
)

const (
	USART1SWn = 0
	I2C1SWn   = 4
)

const (
	HSI14ON   CR2 = 0x01 << 0 //+ Internal High Speed 14MHz clock enable.
	HSI14RDY  CR2 = 0x01 << 1 //+ Internal High Speed 14MHz clock ready flag.
	HSI14DIS  CR2 = 0x01 << 2 //+ Internal High Speed 14MHz clock disable.
	HSI14TRIM CR2 = 0x1F << 3 //+ Internal High Speed 14MHz clock trimming.
	HSI14CAL  CR2 = 0xFF << 8 //+ Internal High Speed 14MHz clock Calibration.
)

const (
	HSI14ONn   = 0
	HSI14RDYn  = 1
	HSI14DISn  = 2
	HSI14TRIMn = 3
	HSI14CALn  = 8
)
