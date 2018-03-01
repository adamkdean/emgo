// +build f746xx

// Peripheral: IWDG_Periph  Independent WATCHDOG.
// Instances:
//  IWDG  mmap.IWDG_BASE
// Registers:
//  0x00 32  KR   Key register.
//  0x04 32  PR   Prescaler register.
//  0x08 32  RLR  Reload register.
//  0x0C 32  SR   Status register.
//  0x10 32  WINR Window register.
// Import:
//  stm32/o/f746xx/mmap
package iwdg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	KEY KR = 0xFFFF << 0 //+ Key value (write only, read 0000h).
)

const (
	KEYn = 0
)

const (
	PR   PR = 0x07 << 0 //+ PR[2:0] (Prescaler divider).
	PR_0 PR = 0x01 << 0 //  Bit 0.
	PR_1 PR = 0x02 << 0 //  Bit 1.
	PR_2 PR = 0x04 << 0 //  Bit 2.
)

const (
	PRn = 0
)

const (
	RL RLR = 0xFFF << 0 //+ Watchdog counter reload value.
)

const (
	RLn = 0
)

const (
	PVU SR = 0x01 << 0 //+ Watchdog prescaler value update.
	RVU SR = 0x01 << 1 //+ Watchdog counter reload value update.
	WVU SR = 0x01 << 2 //+ Watchdog counter window value update.
)

const (
	PVUn = 0
	RVUn = 1
	WVUn = 2
)

const (
	WIN WINR = 0xFFF << 0 //+ Watchdog counter window value.
)

const (
	WINn = 0
)
