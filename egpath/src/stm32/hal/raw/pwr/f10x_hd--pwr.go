// +build f10x_hd

// Peripheral: PWR_Periph  Power Control.
// Instances:
//  PWR  mmap.PWR_BASE
// Registers:
//  0x00 32  CR
//  0x04 32  CSR
// Import:
//  stm32/o/f10x_hd/mmap
package pwr

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LPDS    CR = 0x01 << 0 //+ Low-Power Deepsleep.
	PDDS    CR = 0x01 << 1 //+ Power Down Deepsleep.
	CWUF    CR = 0x01 << 2 //+ Clear Wakeup Flag.
	CSBF    CR = 0x01 << 3 //+ Clear Standby Flag.
	PVDE    CR = 0x01 << 4 //+ Power Voltage Detector Enable.
	PLS     CR = 0x07 << 5 //+ PLS[2:0] bits (PVD Level Selection).
	PLS_0   CR = 0x01 << 5 //  Bit 0.
	PLS_1   CR = 0x02 << 5 //  Bit 1.
	PLS_2   CR = 0x04 << 5 //  Bit 2.
	PLS_2V2 CR = 0x00 << 5 //  PVD level 2.2V.
	PLS_2V3 CR = 0x01 << 5 //  PVD level 2.3V.
	PLS_2V4 CR = 0x02 << 5 //  PVD level 2.4V.
	PLS_2V5 CR = 0x03 << 5 //  PVD level 2.5V.
	PLS_2V6 CR = 0x04 << 5 //  PVD level 2.6V.
	PLS_2V7 CR = 0x05 << 5 //  PVD level 2.7V.
	PLS_2V8 CR = 0x06 << 5 //  PVD level 2.8V.
	PLS_2V9 CR = 0x07 << 5 //  PVD level 2.9V.
	DBP     CR = 0x01 << 8 //+ Disable Backup Domain write protection.
)

const (
	LPDSn = 0
	PDDSn = 1
	CWUFn = 2
	CSBFn = 3
	PVDEn = 4
	PLSn  = 5
	DBPn  = 8
)

const (
	WUF  CSR = 0x01 << 0 //+ Wakeup Flag.
	SBF  CSR = 0x01 << 1 //+ Standby Flag.
	PVDO CSR = 0x01 << 2 //+ PVD Output.
	EWUP CSR = 0x01 << 8 //+ Enable WKUP pin.
)

const (
	WUFn  = 0
	SBFn  = 1
	PVDOn = 2
	EWUPn = 8
)
