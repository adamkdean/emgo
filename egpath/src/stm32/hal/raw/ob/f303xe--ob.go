// +build f303xe

// Peripheral: OB_Periph  Option Bytes Registers.
// Instances:
//  OB  mmap.OB_BASE
// Registers:
//  0x00 16  RDP  FLASH option byte Read protection.
//  0x02 16  USER FLASH option byte user options.
//  0x08 16  WRP0 FLASH option byte write protection 0.
//  0x0A 16  WRP1 FLASH option byte write protection 1.
//  0x0C 16  WRP2 FLASH option byte write protection 2.
//  0x0E 16  WRP3 FLASH option byte write protection 3.
// Import:
//  stm32/o/f303xe/mmap
package ob

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	RDP  RDP = 0xFF << 0 //+ Read protection option byte.
	nRDP RDP = 0xFF << 8 //+ Read protection complemented option byte.
)

const (
	RDPn  = 0
	nRDPn = 8
)

const (
	USER  USER = 0xFF << 16 //+ User option byte.
	nUSER USER = 0xFF << 24 //+ User complemented option byte.
)

const (
	USERn  = 16
	nUSERn = 24
)

const (
	WRP0  WRP0 = 0xFF << 0 //+ Flash memory write protection option bytes.
	nWRP0 WRP0 = 0xFF << 8 //+ Flash memory write protection complemented option bytes.
)

const (
	WRP0n  = 0
	nWRP0n = 8
)

const (
	WRP1  WRP1 = 0xFF << 16 //+ Flash memory write protection option bytes.
	nWRP1 WRP1 = 0xFF << 24 //+ Flash memory write protection complemented option bytes.
)

const (
	WRP1n  = 16
	nWRP1n = 24
)

const (
	WRP2  WRP2 = 0xFF << 0 //+ Flash memory write protection option bytes.
	nWRP2 WRP2 = 0xFF << 8 //+ Flash memory write protection complemented option bytes.
)

const (
	WRP2n  = 0
	nWRP2n = 8
)

const (
	WRP3  WRP3 = 0xFF << 16 //+ Flash memory write protection option bytes.
	nWRP3 WRP3 = 0xFF << 24 //+ Flash memory write protection complemented option bytes.
)

const (
	WRP3n  = 16
	nWRP3n = 24
)
