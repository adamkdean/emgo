// Peripheral: FLASH_Periph  FLASH Registers.
// Instances:
//  FLASH  mmap.FLASH_R_BASE
// Registers:
//  0x00 32  ACR      Access control register.
//  0x04 32  KEYR     Key register.
//  0x08 32  OPTKEYR  Option key register.
//  0x0C 32  SR       Status register.
//  0x10 32  CR       Control register.
//  0x14 32  OPTCR[2] Option control registers.
// Import:
//  stm32/o/f411xe/mmap
package flash

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LATENCY       ACR = 0x0F << 0 //+
	LATENCY_0WS   ACR = 0x00 << 0
	LATENCY_1WS   ACR = 0x01 << 0
	LATENCY_2WS   ACR = 0x02 << 0
	LATENCY_3WS   ACR = 0x03 << 0
	LATENCY_4WS   ACR = 0x04 << 0
	LATENCY_5WS   ACR = 0x05 << 0
	LATENCY_6WS   ACR = 0x06 << 0
	LATENCY_7WS   ACR = 0x07 << 0
	LATENCY_8WS   ACR = 0x08 << 0
	LATENCY_9WS   ACR = 0x09 << 0
	LATENCY_10WS  ACR = 0x0A << 0
	LATENCY_11WS  ACR = 0x0B << 0
	LATENCY_12WS  ACR = 0x0C << 0
	LATENCY_13WS  ACR = 0x0D << 0
	LATENCY_14WS  ACR = 0x0E << 0
	LATENCY_15WS  ACR = 0x0F << 0
	PRFTEN        ACR = 0x01 << 8  //+
	ICEN          ACR = 0x01 << 9  //+
	DCEN          ACR = 0x01 << 10 //+
	ICRST         ACR = 0x01 << 11 //+
	DCRST         ACR = 0x01 << 12 //+
	BYTE0_ADDRESS ACR = 0x10008F << 10
	BYTE2_ADDRESS ACR = 0x40023C03 << 0
)

const (
	LATENCYn = 0
	PRFTENn  = 8
	ICENn    = 9
	DCENn    = 10
	ICRSTn   = 11
	DCRSTn   = 12
)

const (
	EOP    SR = 0x01 << 0  //+
	SOP    SR = 0x01 << 1  //+
	WRPERR SR = 0x01 << 4  //+
	PGAERR SR = 0x01 << 5  //+
	PGPERR SR = 0x01 << 6  //+
	PGSERR SR = 0x01 << 7  //+
	BSY    SR = 0x01 << 16 //+
)

const (
	EOPn    = 0
	SOPn    = 1
	WRPERRn = 4
	PGAERRn = 5
	PGPERRn = 6
	PGSERRn = 7
	BSYn    = 16
)

const (
	PG      CR = 0x01 << 0 //+
	SER     CR = 0x01 << 1 //+
	MER     CR = 0x01 << 2 //+
	SNB     CR = 0x1F << 3 //+
	SNB_0   CR = 0x01 << 3
	SNB_1   CR = 0x02 << 3
	SNB_2   CR = 0x04 << 3
	SNB_3   CR = 0x08 << 3
	SNB_4   CR = 0x08 << 3
	PSIZE   CR = 0x03 << 8 //+
	PSIZE_0 CR = 0x01 << 8
	PSIZE_1 CR = 0x02 << 8
	MER2    CR = 0x01 << 15 //+
	STRT    CR = 0x01 << 16 //+
	EOPIE   CR = 0x01 << 24 //+
	LOCK    CR = 0x01 << 31 //+
)

const (
	PGn    = 0
	SERn   = 1
	MERn   = 2
	SNBn   = 3
	PSIZEn = 8
	MER2n  = 15
	STRTn  = 16
	EOPIEn = 24
	LOCKn  = 31
)

const (
	OPTLOCK    OPTCR = 0x01 << 0 //+
	OPTSTRT    OPTCR = 0x01 << 1 //+
	BOR_LEV_0  OPTCR = 0x01 << 2 //+
	BOR_LEV_1  OPTCR = 0x01 << 3 //+
	BOR_LEV    OPTCR = 0x03 << 2
	BFB2       OPTCR = 0x01 << 4 //+
	WDG_SW     OPTCR = 0x01 << 5 //+
	nRST_STOP  OPTCR = 0x01 << 6 //+
	nRST_STDBY OPTCR = 0x01 << 7 //+
	RDP        OPTCR = 0xFF << 8 //+
	RDP_0      OPTCR = 0x01 << 8
	RDP_1      OPTCR = 0x02 << 8
	RDP_2      OPTCR = 0x04 << 8
	RDP_3      OPTCR = 0x08 << 8
	RDP_4      OPTCR = 0x10 << 8
	RDP_5      OPTCR = 0x20 << 8
	RDP_6      OPTCR = 0x40 << 8
	RDP_7      OPTCR = 0x80 << 8
	nWRP       OPTCR = 0xFFF << 16 //+
	nWRP_0     OPTCR = 0x01 << 16
	nWRP_1     OPTCR = 0x02 << 16
	nWRP_2     OPTCR = 0x04 << 16
	nWRP_3     OPTCR = 0x08 << 16
	nWRP_4     OPTCR = 0x10 << 16
	nWRP_5     OPTCR = 0x20 << 16
	nWRP_6     OPTCR = 0x40 << 16
	nWRP_7     OPTCR = 0x80 << 16
	nWRP_8     OPTCR = 0x100 << 16
	nWRP_9     OPTCR = 0x200 << 16
	nWRP_10    OPTCR = 0x400 << 16
	nWRP_11    OPTCR = 0x800 << 16
	DB1M       OPTCR = 0x01 << 30 //+
	SPRMOD     OPTCR = 0x01 << 31 //+
)

const (
	OPTLOCKn    = 0
	OPTSTRTn    = 1
	BOR_LEV_0n  = 2
	BOR_LEV_1n  = 3
	BFB2n       = 4
	WDG_SWn     = 5
	nRST_STOPn  = 6
	nRST_STDBYn = 7
	RDPn        = 8
	nWRPn       = 16
	DB1Mn       = 30
	SPRMODn     = 31
)
