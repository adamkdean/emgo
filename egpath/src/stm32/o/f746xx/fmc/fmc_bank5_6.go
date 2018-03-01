// Peripheral: FMC_Bank5_6_Periph  Flexible Memory Controller Bank5_6.
// Instances:
//  FMC_Bank5_6  mmap.FMC_Bank5_6_R_BASE
// Registers:
//  0x00 32  SDCR[2] SDRAM Control registers .
//  0x08 32  SDTR[2] SDRAM Timing registers .
//  0x10 32  SDCMR   SDRAM Command Mode register.
//  0x14 32  SDRTR   SDRAM Refresh Timer register.
//  0x18 32  SDSR    SDRAM Status register.
// Import:
//  stm32/o/f746xx/mmap
package fmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	NC       SDCR = 0x03 << 0  //+ NC[1:0] bits (Number of column bits).
	NC_0     SDCR = 0x01 << 0  //  Bit 0.
	NC_1     SDCR = 0x02 << 0  //  Bit 1.
	NR       SDCR = 0x03 << 2  //+ NR[1:0] bits (Number of row bits).
	NR_0     SDCR = 0x01 << 2  //  Bit 0.
	NR_1     SDCR = 0x02 << 2  //  Bit 1.
	SDMWID   SDCR = 0x03 << 4  //+ NR[1:0] bits (Number of row bits).
	SDMWID_0 SDCR = 0x01 << 4  //  Bit 0.
	SDMWID_1 SDCR = 0x02 << 4  //  Bit 1.
	NB       SDCR = 0x01 << 6  //+ Number of internal bank.
	CAS      SDCR = 0x03 << 7  //+ CAS[1:0] bits (CAS latency).
	CAS_0    SDCR = 0x01 << 7  //  Bit 0.
	CAS_1    SDCR = 0x02 << 7  //  Bit 1.
	WP       SDCR = 0x01 << 9  //+ Write protection.
	SDCLK    SDCR = 0x03 << 10 //+ SDRAM clock configuration.
	SDCLK_0  SDCR = 0x01 << 10 //  Bit 0.
	SDCLK_1  SDCR = 0x02 << 10 //  Bit 1.
	RBURST   SDCR = 0x01 << 12 //+ Read burst.
	RPIPE    SDCR = 0x03 << 13 //+ Write protection.
	RPIPE_0  SDCR = 0x01 << 13 //  Bit 0.
	RPIPE_1  SDCR = 0x02 << 13 //  Bit 1.
)

const (
	NCn     = 0
	NRn     = 2
	SDMWIDn = 4
	NBn     = 6
	CASn    = 7
	WPn     = 9
	SDCLKn  = 10
	RBURSTn = 12
	RPIPEn  = 13
)

const (
	TMRD   SDTR = 0x0F << 0  //+ TMRD[3:0] bits (Load mode register to active).
	TMRD_0 SDTR = 0x01 << 0  //  Bit 0.
	TMRD_1 SDTR = 0x02 << 0  //  Bit 1.
	TMRD_2 SDTR = 0x04 << 0  //  Bit 2.
	TMRD_3 SDTR = 0x08 << 0  //  Bit 3.
	TXSR   SDTR = 0x0F << 4  //+ TXSR[3:0] bits (Exit self refresh).
	TXSR_0 SDTR = 0x01 << 4  //  Bit 0.
	TXSR_1 SDTR = 0x02 << 4  //  Bit 1.
	TXSR_2 SDTR = 0x04 << 4  //  Bit 2.
	TXSR_3 SDTR = 0x08 << 4  //  Bit 3.
	TRAS   SDTR = 0x0F << 8  //+ TRAS[3:0] bits (Self refresh time).
	TRAS_0 SDTR = 0x01 << 8  //  Bit 0.
	TRAS_1 SDTR = 0x02 << 8  //  Bit 1.
	TRAS_2 SDTR = 0x04 << 8  //  Bit 2.
	TRAS_3 SDTR = 0x08 << 8  //  Bit 3.
	TRC    SDTR = 0x0F << 12 //+ TRC[2:0] bits (Row cycle delay).
	TRC_0  SDTR = 0x01 << 12 //  Bit 0.
	TRC_1  SDTR = 0x02 << 12 //  Bit 1.
	TRC_2  SDTR = 0x04 << 12 //  Bit 2.
	TWR    SDTR = 0x0F << 16 //+ TRC[2:0] bits (Write recovery delay).
	TWR_0  SDTR = 0x01 << 16 //  Bit 0.
	TWR_1  SDTR = 0x02 << 16 //  Bit 1.
	TWR_2  SDTR = 0x04 << 16 //  Bit 2.
	TRP    SDTR = 0x0F << 20 //+ TRP[2:0] bits (Row precharge delay).
	TRP_0  SDTR = 0x01 << 20 //  Bit 0.
	TRP_1  SDTR = 0x02 << 20 //  Bit 1.
	TRP_2  SDTR = 0x04 << 20 //  Bit 2.
	TRCD   SDTR = 0x0F << 24 //+ TRP[2:0] bits (Row to column delay).
	TRCD_0 SDTR = 0x01 << 24 //  Bit 0.
	TRCD_1 SDTR = 0x02 << 24 //  Bit 1.
	TRCD_2 SDTR = 0x04 << 24 //  Bit 2.
)

const (
	TMRDn = 0
	TXSRn = 4
	TRASn = 8
	TRCn  = 12
	TWRn  = 16
	TRPn  = 20
	TRCDn = 24
)

const (
	MODE   SDCMR = 0x07 << 0   //+ MODE[2:0] bits (Command mode).
	MODE_0 SDCMR = 0x01 << 0   //  Bit 0.
	MODE_1 SDCMR = 0x02 << 0   //  Bit 1.
	MODE_2 SDCMR = 0x03 << 0   //  Bit 2.
	CTB2   SDCMR = 0x01 << 3   //+ Command target 2.
	CTB1   SDCMR = 0x01 << 4   //+ Command target 1.
	NRFS   SDCMR = 0x0F << 5   //+ NRFS[3:0] bits (Number of auto-refresh).
	NRFS_0 SDCMR = 0x01 << 5   //  Bit 0.
	NRFS_1 SDCMR = 0x02 << 5   //  Bit 1.
	NRFS_2 SDCMR = 0x04 << 5   //  Bit 2.
	NRFS_3 SDCMR = 0x08 << 5   //  Bit 3.
	MRD    SDCMR = 0x1FFF << 9 //+ MRD[12:0] bits (Mode register definition).
)

const (
	MODEn = 0
	CTB2n = 3
	CTB1n = 4
	NRFSn = 5
	MRDn  = 9
)

const (
	CRE   SDRTR = 0x01 << 0   //+ Clear refresh error flag.
	COUNT SDRTR = 0x1FFF << 1 //+ COUNT[12:0] bits (Refresh timer count).
	REIE  SDRTR = 0x01 << 14  //+ RES interupt enable.
)

const (
	CREn   = 0
	COUNTn = 1
	REIEn  = 14
)

const (
	RE       SDSR = 0x01 << 0 //+ Refresh error flag.
	MODES1   SDSR = 0x03 << 1 //+ MODES1[1:0]bits (Status mode for bank 1).
	MODES1_0 SDSR = 0x01 << 1 //  Bit 0.
	MODES1_1 SDSR = 0x02 << 1 //  Bit 1.
	MODES2   SDSR = 0x03 << 3 //+ MODES2[1:0]bits (Status mode for bank 2).
	MODES2_0 SDSR = 0x01 << 3 //  Bit 0.
	MODES2_1 SDSR = 0x02 << 3 //  Bit 1.
	BUSY     SDSR = 0x01 << 5 //+ Busy status.
)

const (
	REn     = 0
	MODES1n = 1
	MODES2n = 3
	BUSYn   = 5
)
