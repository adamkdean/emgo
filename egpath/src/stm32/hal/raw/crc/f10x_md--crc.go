// +build f10x_md

// Peripheral: CRC_Periph  CRC calculation unit.
// Instances:
//  CRC  mmap.CRC_BASE
// Registers:
//  0x00 32  DR
//  0x04  8  IDR
//  0x08 32  CR
// Import:
//  stm32/o/f10x_md/mmap
package crc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	RESET CR_Bits = 0x01 << 0 //+ RESET bit.
)

const (
	RESETn = 0
)
