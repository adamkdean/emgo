// +build f746xx

// Peripheral: ETH_Periph  Ethernet MAC.
// Instances:
//  ETH  mmap.ETH_BASE
// Registers:
//  0x00   32  MACCR
//  0x04   32  MACFFR
//  0x08   32  MACHTHR
//  0x0C   32  MACHTLR
//  0x10   32  MACMIIAR
//  0x14   32  MACMIIDR
//  0x18   32  MACFCR
//  0x1C   32  MACVLANTR   8.
//  0x28   32  MACRWUFFR   11.
//  0x2C   32  MACPMTCSR
//  0x38   32  MACSR       15.
//  0x3C   32  MACIMR
//  0x40   32  MACA0HR
//  0x44   32  MACA0LR
//  0x48   32  MACA1HR
//  0x4C   32  MACA1LR
//  0x50   32  MACA2HR
//  0x54   32  MACA2LR
//  0x58   32  MACA3HR
//  0x5C   32  MACA3LR     24.
//  0x100  32  MMCCR       65.
//  0x104  32  MMCRIR
//  0x108  32  MMCTIR
//  0x10C  32  MMCRIMR
//  0x110  32  MMCTIMR     69.
//  0x14C  32  MMCTGFSCCR  84.
//  0x150  32  MMCTGFMSCCR
//  0x168  32  MMCTGFCR
//  0x194  32  MMCRFCECR
//  0x198  32  MMCRFAECR
//  0x1C4  32  MMCRGUFCR
//  0x700  32  PTPTSCR
//  0x704  32  PTPSSIR
//  0x708  32  PTPTSHR
//  0x70C  32  PTPTSLR
//  0x710  32  PTPTSHUR
//  0x714  32  PTPTSLUR
//  0x718  32  PTPTSAR
//  0x71C  32  PTPTTHR
//  0x720  32  PTPTTLR
//  0x724  32  RESERVED8
//  0x728  32  PTPTSSR
//  0x1000 32  DMABMR
//  0x1004 32  DMATPDR
//  0x1008 32  DMARPDR
//  0x100C 32  DMARDLAR
//  0x1010 32  DMATDLAR
//  0x1014 32  DMASR
//  0x1018 32  DMAOMR
//  0x101C 32  DMAIER
//  0x1020 32  DMAMFBOCR
//  0x1024 32  DMARSWTR
//  0x1048 32  DMACHTDR
//  0x104C 32  DMACHRDR
//  0x1050 32  DMACHTBAR
//  0x1054 32  DMACHRBAR
// Import:
//  stm32/o/f746xx/mmap
package eth

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	WD        MACCR = 0x01 << 23 //+ Watchdog disable.
	JD        MACCR = 0x01 << 22 //+ Jabber disable.
	IFG       MACCR = 0x07 << 17 //+ Inter-frame gap.
	IFG_96Bit MACCR = 0x00 << 17 //  Minimum IFG between frames during transmission is 96Bit.
	IFG_88Bit MACCR = 0x01 << 17 //  Minimum IFG between frames during transmission is 88Bit.
	IFG_80Bit MACCR = 0x02 << 17 //  Minimum IFG between frames during transmission is 80Bit.
	IFG_72Bit MACCR = 0x03 << 17 //  Minimum IFG between frames during transmission is 72Bit.
	IFG_64Bit MACCR = 0x04 << 17 //  Minimum IFG between frames during transmission is 64Bit.
	IFG_56Bit MACCR = 0x05 << 17 //  Minimum IFG between frames during transmission is 56Bit.
	IFG_48Bit MACCR = 0x06 << 17 //  Minimum IFG between frames during transmission is 48Bit.
	IFG_40Bit MACCR = 0x07 << 17 //  Minimum IFG between frames during transmission is 40Bit.
	CSD       MACCR = 0x01 << 16 //+ Carrier sense disable (during transmission).
	FES       MACCR = 0x01 << 14 //+ Fast ethernet speed.
	ROD       MACCR = 0x01 << 13 //+ Receive own disable.
	LM        MACCR = 0x01 << 12 //+ loopback mode.
	DM        MACCR = 0x01 << 11 //+ Duplex mode.
	IPCO      MACCR = 0x01 << 10 //+ IP Checksum offload.
	RD        MACCR = 0x01 << 9  //+ Retry disable.
	APCS      MACCR = 0x01 << 7  //+ Automatic Pad/CRC stripping.
	BL        MACCR = 0x03 << 5  //+ Back-off limit: random integer number (r) of slot time delays before rescheduling.
	BL_10     MACCR = 0x00 << 5  //  k = min (n, 10).
	BL_8      MACCR = 0x01 << 5  //  k = min (n, 8).
	BL_4      MACCR = 0x02 << 5  //  k = min (n, 4).
	BL_1      MACCR = 0x03 << 5  //  k = min (n, 1).
	DC        MACCR = 0x01 << 4  //+ Defferal check.
	TE        MACCR = 0x01 << 3  //+ Transmitter enable.
	RE        MACCR = 0x01 << 2  //+ Receiver enable.
)

const (
	WDn   = 23
	JDn   = 22
	IFGn  = 17
	CSDn  = 16
	FESn  = 14
	RODn  = 13
	LMn   = 12
	DMn   = 11
	IPCOn = 10
	RDn   = 9
	APCSn = 7
	BLn   = 5
	DCn   = 4
	TEn   = 3
	REn   = 2
)

const (
	RA                          MACFFR = 0x01 << 31 //+ Receive all.
	HPF                         MACFFR = 0x01 << 10 //+ Hash or perfect filter.
	SAF                         MACFFR = 0x01 << 9  //+ Source address filter enable.
	SAIF                        MACFFR = 0x01 << 8  //+ SA inverse filtering.
	PCF                         MACFFR = 0x03 << 6  //+ Pass control frames: 3 cases.
	PCF_BlockAll                MACFFR = 0x01 << 6  //  MAC filters all control frames from reaching the application.
	PCF_ForwardAll              MACFFR = 0x02 << 6  //  MAC forwards all control frames to application even if they fail the Address Filter.
	PCF_ForwardPassedAddrFilter MACFFR = 0x03 << 6  //  MAC forwards control frames that pass the Address Filter..
	BFD                         MACFFR = 0x01 << 5  //+ Broadcast frame disable.
	PAM                         MACFFR = 0x01 << 4  //+ Pass all mutlicast.
	DAIF                        MACFFR = 0x01 << 3  //+ DA Inverse filtering.
	HM                          MACFFR = 0x01 << 2  //+ Hash multicast.
	HU                          MACFFR = 0x01 << 1  //+ Hash unicast.
	PM                          MACFFR = 0x01 << 0  //+ Promiscuous mode.
)

const (
	RAn   = 31
	HPFn  = 10
	SAFn  = 9
	SAIFn = 8
	PCFn  = 6
	BFDn  = 5
	PAMn  = 4
	DAIFn = 3
	HMn   = 2
	HUn   = 1
	PMn   = 0
)

const (
	HTH MACHTHR = 0xFFFFFFFF << 0 //+ Hash table high.
)

const (
	HTHn = 0
)

const (
	HTL MACHTLR = 0xFFFFFFFF << 0 //+ Hash table low.
)

const (
	HTLn = 0
)

const (
	PA        MACMIIAR = 0x1F << 11 //+ Physical layer address.
	MR        MACMIIAR = 0x1F << 6  //+ MII register in the selected PHY.
	CR        MACMIIAR = 0x07 << 2  //+ CR clock range: 6 cases.
	CR_Div42  MACMIIAR = 0x00 << 2  //  HCLK:60-100 MHz; MDC clock= HCLK/42.
	CR_Div62  MACMIIAR = 0x01 << 2  //  HCLK:100-150 MHz; MDC clock= HCLK/62.
	CR_Div16  MACMIIAR = 0x02 << 2  //  HCLK:20-35 MHz; MDC clock= HCLK/16.
	CR_Div26  MACMIIAR = 0x03 << 2  //  HCLK:35-60 MHz; MDC clock= HCLK/26.
	CR_Div102 MACMIIAR = 0x04 << 2  //  HCLK:150-168 MHz; MDC clock= HCLK/102.
	MW        MACMIIAR = 0x01 << 1  //+ MII write.
	MB        MACMIIAR = 0x01 << 0  //+ MII busy.
)

const (
	PAn = 11
	MRn = 6
	CRn = 2
	MWn = 1
	MBn = 0
)

const (
	MD MACMIIDR = 0xFFFF << 0 //+ MII data: read/write data from/to PHY.
)

const (
	MDn = 0
)

const (
	PT           MACFCR = 0xFFFF << 16 //+ Pause time.
	ZQPD         MACFCR = 0x01 << 7    //+ Zero-quanta pause disable.
	PLT          MACFCR = 0x03 << 4    //+ Pause low threshold: 4 cases.
	PLT_Minus4   MACFCR = 0x00 << 4    //  Pause time minus 4 slot times.
	PLT_Minus28  MACFCR = 0x01 << 4    //  Pause time minus 28 slot times.
	PLT_Minus144 MACFCR = 0x02 << 4    //  Pause time minus 144 slot times.
	PLT_Minus256 MACFCR = 0x03 << 4    //  Pause time minus 256 slot times.
	UPFD         MACFCR = 0x01 << 3    //+ Unicast pause frame detect.
	RFCE         MACFCR = 0x01 << 2    //+ Receive flow control enable.
	TFCE         MACFCR = 0x01 << 1    //+ Transmit flow control enable.
	FCBBPA       MACFCR = 0x01 << 0    //+ Flow control busy/backpressure activate.
)

const (
	PTn     = 16
	ZQPDn   = 7
	PLTn    = 4
	UPFDn   = 3
	RFCEn   = 2
	TFCEn   = 1
	FCBBPAn = 0
)

const (
	VLANTC MACVLANTR = 0x01 << 16  //+ 12-bit VLAN tag comparison.
	VLANTI MACVLANTR = 0xFFFF << 0 //+ VLAN tag identifier (for receive frames).
)

const (
	VLANTCn = 16
	VLANTIn = 0
)

const (
	D MACRWUFFR = 0xFFFFFFFF << 0 //+ Wake-up frame filter register data.
)

const (
	Dn = 0
)

const (
	WFFRPR MACPMTCSR = 0x01 << 31 //+ Wake-Up Frame Filter Register Pointer Reset.
	GU     MACPMTCSR = 0x01 << 9  //+ Global Unicast.
	WFR    MACPMTCSR = 0x01 << 6  //+ Wake-Up Frame Received.
	MPR    MACPMTCSR = 0x01 << 5  //+ Magic Packet Received.
	WFE    MACPMTCSR = 0x01 << 2  //+ Wake-Up Frame Enable.
	MPE    MACPMTCSR = 0x01 << 1  //+ Magic Packet Enable.
	PD     MACPMTCSR = 0x01 << 0  //+ Power Down.
)

const (
	WFFRPRn = 31
	GUn     = 9
	WFRn    = 6
	MPRn    = 5
	WFEn    = 2
	MPEn    = 1
	PDn     = 0
)

const (
	TSTS   MACSR = 0x01 << 9 //+ Time stamp trigger status.
	MMCTS  MACSR = 0x01 << 6 //+ MMC transmit status.
	MMMCRS MACSR = 0x01 << 5 //+ MMC receive status.
	MMCS   MACSR = 0x01 << 4 //+ MMC status.
	PMTS   MACSR = 0x01 << 3 //+ PMT status.
)

const (
	TSTSn   = 9
	MMCTSn  = 6
	MMMCRSn = 5
	MMCSn   = 4
	PMTSn   = 3
)

const (
	TSTIM MACIMR = 0x01 << 9 //+ Time stamp trigger interrupt mask.
	PMTIM MACIMR = 0x01 << 3 //+ PMT interrupt mask.
)

const (
	TSTIMn = 9
	PMTIMn = 3
)

const (
	MACA0H MACA0HR = 0xFFFF << 0 //+ MAC address0 high.
)

const (
	MACA0Hn = 0
)

const (
	MACA0L MACA0LR = 0xFFFFFFFF << 0 //+ MAC address0 low.
)

const (
	MACA0Ln = 0
)

const (
	AE             MACA1HR = 0x01 << 31  //+ Address enable.
	SA             MACA1HR = 0x01 << 30  //+ Source address.
	MBC            MACA1HR = 0x3F << 24  //+ Mask byte control: bits to mask for comparison of the MAC Address bytes.
	MBC_HBits15_8  MACA1HR = 0x20 << 24  //  Mask MAC Address high reg bits [15:8].
	MBC_HBits7_0   MACA1HR = 0x10 << 24  //  Mask MAC Address high reg bits [7:0].
	MBC_LBits31_24 MACA1HR = 0x08 << 24  //  Mask MAC Address low reg bits [31:24].
	MBC_LBits23_16 MACA1HR = 0x04 << 24  //  Mask MAC Address low reg bits [23:16].
	MBC_LBits15_8  MACA1HR = 0x02 << 24  //  Mask MAC Address low reg bits [15:8].
	MBC_LBits7_0   MACA1HR = 0x01 << 24  //  Mask MAC Address low reg bits [7:0].
	MACA1H         MACA1HR = 0xFFFF << 0 //+ MAC address1 high.
)

const (
	AEn     = 31
	SAn     = 30
	MBCn    = 24
	MACA1Hn = 0
)

const (
	MACA1L MACA1LR = 0xFFFFFFFF << 0 //+ MAC address1 low.
)

const (
	MACA1Ln = 0
)

const (
	AE             MACA2HR = 0x01 << 31  //+ Address enable.
	SA             MACA2HR = 0x01 << 30  //+ Source address.
	MBC            MACA2HR = 0x3F << 24  //+ Mask byte control.
	MBC_HBits15_8  MACA2HR = 0x20 << 24  //  Mask MAC Address high reg bits [15:8].
	MBC_HBits7_0   MACA2HR = 0x10 << 24  //  Mask MAC Address high reg bits [7:0].
	MBC_LBits31_24 MACA2HR = 0x08 << 24  //  Mask MAC Address low reg bits [31:24].
	MBC_LBits23_16 MACA2HR = 0x04 << 24  //  Mask MAC Address low reg bits [23:16].
	MBC_LBits15_8  MACA2HR = 0x02 << 24  //  Mask MAC Address low reg bits [15:8].
	MBC_LBits7_0   MACA2HR = 0x01 << 24  //  Mask MAC Address low reg bits [70].
	MACA2H         MACA2HR = 0xFFFF << 0 //+ MAC address1 high.
)

const (
	AEn     = 31
	SAn     = 30
	MBCn    = 24
	MACA2Hn = 0
)

const (
	MACA2L MACA2LR = 0xFFFFFFFF << 0 //+ MAC address2 low.
)

const (
	MACA2Ln = 0
)

const (
	AE             MACA3HR = 0x01 << 31  //+ Address enable.
	SA             MACA3HR = 0x01 << 30  //+ Source address.
	MBC            MACA3HR = 0x3F << 24  //+ Mask byte control.
	MBC_HBits15_8  MACA3HR = 0x20 << 24  //  Mask MAC Address high reg bits [15:8].
	MBC_HBits7_0   MACA3HR = 0x10 << 24  //  Mask MAC Address high reg bits [7:0].
	MBC_LBits31_24 MACA3HR = 0x08 << 24  //  Mask MAC Address low reg bits [31:24].
	MBC_LBits23_16 MACA3HR = 0x04 << 24  //  Mask MAC Address low reg bits [23:16].
	MBC_LBits15_8  MACA3HR = 0x02 << 24  //  Mask MAC Address low reg bits [15:8].
	MBC_LBits7_0   MACA3HR = 0x01 << 24  //  Mask MAC Address low reg bits [70].
	MACA3H         MACA3HR = 0xFFFF << 0 //+ MAC address3 high.
)

const (
	AEn     = 31
	SAn     = 30
	MBCn    = 24
	MACA3Hn = 0
)

const (
	MACA3L MACA3LR = 0xFFFFFFFF << 0 //+ MAC address3 low.
)

const (
	MACA3Ln = 0
)

const (
	MCFHP MMCCR = 0x01 << 5 //+ MMC counter Full-Half preset.
	MCP   MMCCR = 0x01 << 4 //+ MMC counter preset.
	MCF   MMCCR = 0x01 << 3 //+ MMC Counter Freeze.
	ROR   MMCCR = 0x01 << 2 //+ Reset on Read.
	CSR   MMCCR = 0x01 << 1 //+ Counter Stop Rollover.
	CR    MMCCR = 0x01 << 0 //+ Counters Reset.
)

const (
	MCFHPn = 5
	MCPn   = 4
	MCFn   = 3
	RORn   = 2
	CSRn   = 1
	CRn    = 0
)

const (
	RGUFS MMCRIR = 0x01 << 17 //+ Set when Rx good unicast frames counter reaches half the maximum value.
	RFAES MMCRIR = 0x01 << 6  //+ Set when Rx alignment error counter reaches half the maximum value.
	RFCES MMCRIR = 0x01 << 5  //+ Set when Rx crc error counter reaches half the maximum value.
)

const (
	RGUFSn = 17
	RFAESn = 6
	RFCESn = 5
)

const (
	TGFS    MMCTIR = 0x01 << 21 //+ Set when Tx good frame count counter reaches half the maximum value.
	TGFMSCS MMCTIR = 0x01 << 15 //+ Set when Tx good multi col counter reaches half the maximum value.
	TGFSCS  MMCTIR = 0x01 << 14 //+ Set when Tx good single col counter reaches half the maximum value.
)

const (
	TGFSn    = 21
	TGFMSCSn = 15
	TGFSCSn  = 14
)

const (
	RGUFM MMCRIMR = 0x01 << 17 //+ Mask the interrupt when Rx good unicast frames counter reaches half the maximum value.
	RFAEM MMCRIMR = 0x01 << 6  //+ Mask the interrupt when when Rx alignment error counter reaches half the maximum value.
	RFCEM MMCRIMR = 0x01 << 5  //+ Mask the interrupt when Rx crc error counter reaches half the maximum value.
)

const (
	RGUFMn = 17
	RFAEMn = 6
	RFCEMn = 5
)

const (
	TGFM    MMCTIMR = 0x01 << 21 //+ Mask the interrupt when Tx good frame count counter reaches half the maximum value.
	TGFMSCM MMCTIMR = 0x01 << 15 //+ Mask the interrupt when Tx good multi col counter reaches half the maximum value.
	TGFSCM  MMCTIMR = 0x01 << 14 //+ Mask the interrupt when Tx good single col counter reaches half the maximum value.
)

const (
	TGFMn    = 21
	TGFMSCMn = 15
	TGFSCMn  = 14
)

const (
	TGFSCC MMCTGFSCCR = 0xFFFFFFFF << 0 //+ Number of successfully transmitted frames after a single collision in Half-duplex mode..
)

const (
	TGFSCCn = 0
)

const (
	TGFMSCC MMCTGFMSCCR = 0xFFFFFFFF << 0 //+ Number of successfully transmitted frames after more than a single collision in Half-duplex mode..
)

const (
	TGFMSCCn = 0
)

const (
	TGFC MMCTGFCR = 0xFFFFFFFF << 0 //+ Number of good frames transmitted..
)

const (
	TGFCn = 0
)

const (
	RFCEC MMCRFCECR = 0xFFFFFFFF << 0 //+ Number of frames received with CRC error..
)

const (
	RFCECn = 0
)

const (
	RFAEC MMCRFAECR = 0xFFFFFFFF << 0 //+ Number of frames received with alignment (dribble) error.
)

const (
	RFAECn = 0
)

const (
	RGUFC MMCRGUFCR = 0xFFFFFFFF << 0 //+ Number of good unicast frames received..
)

const (
	RGUFCn = 0
)

const (
	TSCNT PTPTSCR = 0x03 << 16 //+ Time stamp clock node type.
	TSARU PTPTSCR = 0x01 << 5  //+ Addend register update.
	TSITE PTPTSCR = 0x01 << 4  //+ Time stamp interrupt trigger enable.
	TSSTU PTPTSCR = 0x01 << 3  //+ Time stamp update.
	TSSTI PTPTSCR = 0x01 << 2  //+ Time stamp initialize.
	TSFCU PTPTSCR = 0x01 << 1  //+ Time stamp fine or coarse update.
	TSE   PTPTSCR = 0x01 << 0  //+ Time stamp enable.
)

const (
	TSCNTn = 16
	TSARUn = 5
	TSITEn = 4
	TSSTUn = 3
	TSSTIn = 2
	TSFCUn = 1
	TSEn   = 0
)

const (
	STSSI PTPSSIR = 0xFF << 0 //+ System time Sub-second increment value.
)

const (
	STSSIn = 0
)

const (
	STS PTPTSHR = 0xFFFFFFFF << 0 //+ System Time second.
)

const (
	STSn = 0
)

const (
	STPNS PTPTSLR = 0x01 << 31      //+ System Time Positive or negative time.
	STSS  PTPTSLR = 0x7FFFFFFF << 0 //+ System Time sub-seconds.
)

const (
	STPNSn = 31
	STSSn  = 0
)

const (
	TSUS PTPTSHUR = 0xFFFFFFFF << 0 //+ Time stamp update seconds.
)

const (
	TSUSn = 0
)

const (
	TSUPNS PTPTSLUR = 0x01 << 31      //+ Time stamp update Positive or negative time.
	TSUSS  PTPTSLUR = 0x7FFFFFFF << 0 //+ Time stamp update sub-seconds.
)

const (
	TSUPNSn = 31
	TSUSSn  = 0
)

const (
	TSA PTPTSAR = 0xFFFFFFFF << 0 //+ Time stamp addend.
)

const (
	TSAn = 0
)

const (
	TTSH PTPTTHR = 0xFFFFFFFF << 0 //+ Target time stamp high.
)

const (
	TTSHn = 0
)

const (
	TTSL PTPTTLR = 0xFFFFFFFF << 0 //+ Target time stamp low.
)

const (
	TTSLn = 0
)

const (
	TSSMRME    PTPTSSR = 0x01 << 15 //+ Time stamp snapshot for message relevant to master enable.
	TSSEME     PTPTSSR = 0x01 << 14 //+ Time stamp snapshot for event message enable.
	TSSIPV4FE  PTPTSSR = 0x01 << 13 //+ Time stamp snapshot for IPv4 frames enable.
	TSSIPV6FE  PTPTSSR = 0x01 << 12 //+ Time stamp snapshot for IPv6 frames enable.
	TSSPTPOEFE PTPTSSR = 0x01 << 11 //+ Time stamp snapshot for PTP over ethernet frames enable.
	TSPTPPSV2E PTPTSSR = 0x01 << 10 //+ Time stamp PTP packet snooping for version2 format enable.
	TSSSR      PTPTSSR = 0x01 << 9  //+ Time stamp Sub-seconds rollover.
	TSSARFE    PTPTSSR = 0x01 << 8  //+ Time stamp snapshot for all received frames enable.
	TSTTR      PTPTSSR = 0x01 << 5  //+ Time stamp target time reached.
	TSSO       PTPTSSR = 0x01 << 4  //+ Time stamp seconds overflow.
)

const (
	TSSMRMEn    = 15
	TSSEMEn     = 14
	TSSIPV4FEn  = 13
	TSSIPV6FEn  = 12
	TSSPTPOEFEn = 11
	TSPTPPSV2En = 10
	TSSSRn      = 9
	TSSARFEn    = 8
	TSTTRn      = 5
	TSSOn       = 4
)

const (
	AAB               DMABMR = 0x01 << 25   //+ Address-Aligned beats.
	FPM               DMABMR = 0x01 << 24   //+ 4xPBL mode.
	USP               DMABMR = 0x01 << 23   //+ Use separate PBL.
	RDP               DMABMR = 0x3F << 17   //+ RxDMA PBL.
	RDP_1Beat         DMABMR = 0x01 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 1.
	RDP_2Beat         DMABMR = 0x02 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 2.
	RDP_4Beat         DMABMR = 0x04 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 4.
	RDP_8Beat         DMABMR = 0x08 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 8.
	RDP_16Beat        DMABMR = 0x10 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 16.
	RDP_32Beat        DMABMR = 0x20 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 32.
	RDP_4xPBL_4Beat   DMABMR = 0x81 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 4.
	RDP_4xPBL_8Beat   DMABMR = 0x82 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 8.
	RDP_4xPBL_16Beat  DMABMR = 0x84 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 16.
	RDP_4xPBL_32Beat  DMABMR = 0x88 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 32.
	RDP_4xPBL_64Beat  DMABMR = 0x90 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 64.
	RDP_4xPBL_128Beat DMABMR = 0xA0 << 17   //  maximum number of beats to be transferred in one RxDMA transaction is 128.
	FB                DMABMR = 0x01 << 16   //+ Fixed Burst.
	RTPR              DMABMR = 0x03 << 14   //+ Rx Tx priority ratio.
	RTPR_1_1          DMABMR = 0x00 << 14   //  Rx Tx priority ratio.
	RTPR_2_1          DMABMR = 0x01 << 14   //  Rx Tx priority ratio.
	RTPR_3_1          DMABMR = 0x02 << 14   //  Rx Tx priority ratio.
	RTPR_4_1          DMABMR = 0x03 << 14   //  Rx Tx priority ratio.
	PBL               DMABMR = 0x3F << 8    //+ Programmable burst length.
	PBL_1Beat         DMABMR = 0x01 << 8    //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 1.
	PBL_2Beat         DMABMR = 0x02 << 8    //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 2.
	PBL_4Beat         DMABMR = 0x04 << 8    //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 4.
	PBL_8Beat         DMABMR = 0x08 << 8    //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 8.
	PBL_16Beat        DMABMR = 0x10 << 8    //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 16.
	PBL_32Beat        DMABMR = 0x20 << 8    //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 32.
	PBL_4xPBL_4Beat   DMABMR = 0x10001 << 8 //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 4.
	PBL_4xPBL_8Beat   DMABMR = 0x10002 << 8 //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 8.
	PBL_4xPBL_16Beat  DMABMR = 0x10004 << 8 //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 16.
	PBL_4xPBL_32Beat  DMABMR = 0x10008 << 8 //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 32.
	PBL_4xPBL_64Beat  DMABMR = 0x10010 << 8 //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 64.
	PBL_4xPBL_128Beat DMABMR = 0x10020 << 8 //  maximum number of beats to be transferred in one TxDMA (or both) transaction is 128.
	EDE               DMABMR = 0x01 << 7    //+ Enhanced Descriptor Enable.
	DSL               DMABMR = 0x1F << 2    //+ Descriptor Skip Length.
	DA                DMABMR = 0x01 << 1    //+ DMA arbitration scheme.
	SR                DMABMR = 0x01 << 0    //+ Software reset.
)

const (
	AABn  = 25
	FPMn  = 24
	USPn  = 23
	RDPn  = 17
	FBn   = 16
	RTPRn = 14
	PBLn  = 8
	EDEn  = 7
	DSLn  = 2
	DAn   = 1
	SRn   = 0
)

const (
	TPD DMATPDR = 0xFFFFFFFF << 0 //+ Transmit poll demand.
)

const (
	TPDn = 0
)

const (
	RPD DMARPDR = 0xFFFFFFFF << 0 //+ Receive poll demand.
)

const (
	RPDn = 0
)

const (
	SRL DMARDLAR = 0xFFFFFFFF << 0 //+ Start of receive list.
)

const (
	SRLn = 0
)

const (
	STL DMATDLAR = 0xFFFFFFFF << 0 //+ Start of transmit list.
)

const (
	STLn = 0
)

const (
	TSTS             DMASR = 0x01 << 29 //+ Time-stamp trigger status.
	PMTS             DMASR = 0x01 << 28 //+ PMT status.
	MMCS             DMASR = 0x01 << 27 //+ MMC status.
	EBS              DMASR = 0x07 << 23 //+ Error bits status.
	EBS_DescAccess   DMASR = 0x04 << 23 //  Error bits 0-data buffer, 1-desc. access.
	EBS_ReadTransf   DMASR = 0x02 << 23 //  Error bits 0-write trnsf, 1-read transfr.
	EBS_DataTransfTx DMASR = 0x01 << 23 //  Error bits 0-Rx DMA, 1-Tx DMA.
	TPS              DMASR = 0x07 << 20 //+ Transmit process state.
	TPS_Stopped      DMASR = 0x00 << 20 //  Stopped - Reset or Stop Tx Command issued.
	TPS_Fetching     DMASR = 0x01 << 20 //  Running - fetching the Tx descriptor.
	TPS_Waiting      DMASR = 0x02 << 20 //  Running - waiting for status.
	TPS_Reading      DMASR = 0x03 << 20 //  Running - reading the data from host memory.
	TPS_Suspended    DMASR = 0x06 << 20 //  Suspended - Tx Descriptor unavailabe.
	TPS_Closing      DMASR = 0x07 << 20 //  Running - closing Rx descriptor.
	RPS              DMASR = 0x07 << 17 //+ Receive process state.
	RPS_Stopped      DMASR = 0x00 << 17 //  Stopped - Reset or Stop Rx Command issued.
	RPS_Fetching     DMASR = 0x01 << 17 //  Running - fetching the Rx descriptor.
	RPS_Waiting      DMASR = 0x03 << 17 //  Running - waiting for packet.
	RPS_Suspended    DMASR = 0x04 << 17 //  Suspended - Rx Descriptor unavailable.
	RPS_Closing      DMASR = 0x05 << 17 //  Running - closing descriptor.
	RPS_Queuing      DMASR = 0x07 << 17 //  Running - queuing the recieve frame into host memory.
	NIS              DMASR = 0x01 << 16 //+ Normal interrupt summary.
	AIS              DMASR = 0x01 << 15 //+ Abnormal interrupt summary.
	ERS              DMASR = 0x01 << 14 //+ Early receive status.
	FBES             DMASR = 0x01 << 13 //+ Fatal bus error status.
	ETS              DMASR = 0x01 << 10 //+ Early transmit status.
	RWTS             DMASR = 0x01 << 9  //+ Receive watchdog timeout status.
	RPSS             DMASR = 0x01 << 8  //+ Receive process stopped status.
	RBUS             DMASR = 0x01 << 7  //+ Receive buffer unavailable status.
	RS               DMASR = 0x01 << 6  //+ Receive status.
	TUS              DMASR = 0x01 << 5  //+ Transmit underflow status.
	ROS              DMASR = 0x01 << 4  //+ Receive overflow status.
	TJTS             DMASR = 0x01 << 3  //+ Transmit jabber timeout status.
	TBUS             DMASR = 0x01 << 2  //+ Transmit buffer unavailable status.
	TPSS             DMASR = 0x01 << 1  //+ Transmit process stopped status.
	TS               DMASR = 0x01 << 0  //+ Transmit status.
)

const (
	TSTSn = 29
	PMTSn = 28
	MMCSn = 27
	EBSn  = 23
	TPSn  = 20
	RPSn  = 17
	NISn  = 16
	AISn  = 15
	ERSn  = 14
	FBESn = 13
	ETSn  = 10
	RWTSn = 9
	RPSSn = 8
	RBUSn = 7
	RSn   = 6
	TUSn  = 5
	ROSn  = 4
	TJTSn = 3
	TBUSn = 2
	TPSSn = 1
	TSn   = 0
)

const (
	DTCEFD       DMAOMR = 0x01 << 26 //+ Disable Dropping of TCP/IP checksum error frames.
	RSF          DMAOMR = 0x01 << 25 //+ Receive store and forward.
	DFRF         DMAOMR = 0x01 << 24 //+ Disable flushing of received frames.
	TSF          DMAOMR = 0x01 << 21 //+ Transmit store and forward.
	FTF          DMAOMR = 0x01 << 20 //+ Flush transmit FIFO.
	TTC          DMAOMR = 0x07 << 14 //+ Transmit threshold control.
	TTC_64Bytes  DMAOMR = 0x00 << 14 //  threshold level of the MTL Transmit FIFO is 64 Bytes.
	TTC_128Bytes DMAOMR = 0x01 << 14 //  threshold level of the MTL Transmit FIFO is 128 Bytes.
	TTC_192Bytes DMAOMR = 0x02 << 14 //  threshold level of the MTL Transmit FIFO is 192 Bytes.
	TTC_256Bytes DMAOMR = 0x03 << 14 //  threshold level of the MTL Transmit FIFO is 256 Bytes.
	TTC_40Bytes  DMAOMR = 0x04 << 14 //  threshold level of the MTL Transmit FIFO is 40 Bytes.
	TTC_32Bytes  DMAOMR = 0x05 << 14 //  threshold level of the MTL Transmit FIFO is 32 Bytes.
	TTC_24Bytes  DMAOMR = 0x06 << 14 //  threshold level of the MTL Transmit FIFO is 24 Bytes.
	TTC_16Bytes  DMAOMR = 0x07 << 14 //  threshold level of the MTL Transmit FIFO is 16 Bytes.
	ST           DMAOMR = 0x01 << 13 //+ Start/stop transmission command.
	FEF          DMAOMR = 0x01 << 7  //+ Forward error frames.
	FUGF         DMAOMR = 0x01 << 6  //+ Forward undersized good frames.
	RTC          DMAOMR = 0x03 << 3  //+ receive threshold control.
	RTC_64Bytes  DMAOMR = 0x00 << 3  //  threshold level of the MTL Receive FIFO is 64 Bytes.
	RTC_32Bytes  DMAOMR = 0x01 << 3  //  threshold level of the MTL Receive FIFO is 32 Bytes.
	RTC_96Bytes  DMAOMR = 0x02 << 3  //  threshold level of the MTL Receive FIFO is 96 Bytes.
	RTC_128Bytes DMAOMR = 0x03 << 3  //  threshold level of the MTL Receive FIFO is 128 Bytes.
	OSF          DMAOMR = 0x01 << 2  //+ operate on second frame.
	SR           DMAOMR = 0x01 << 1  //+ Start/stop receive.
)

const (
	DTCEFDn = 26
	RSFn    = 25
	DFRFn   = 24
	TSFn    = 21
	FTFn    = 20
	TTCn    = 14
	STn     = 13
	FEFn    = 7
	FUGFn   = 6
	RTCn    = 3
	OSFn    = 2
	SRn     = 1
)

const (
	NISE  DMAIER = 0x01 << 16 //+ Normal interrupt summary enable.
	AISE  DMAIER = 0x01 << 15 //+ Abnormal interrupt summary enable.
	ERIE  DMAIER = 0x01 << 14 //+ Early receive interrupt enable.
	FBEIE DMAIER = 0x01 << 13 //+ Fatal bus error interrupt enable.
	ETIE  DMAIER = 0x01 << 10 //+ Early transmit interrupt enable.
	RWTIE DMAIER = 0x01 << 9  //+ Receive watchdog timeout interrupt enable.
	RPSIE DMAIER = 0x01 << 8  //+ Receive process stopped interrupt enable.
	RBUIE DMAIER = 0x01 << 7  //+ Receive buffer unavailable interrupt enable.
	RIE   DMAIER = 0x01 << 6  //+ Receive interrupt enable.
	TUIE  DMAIER = 0x01 << 5  //+ Transmit Underflow interrupt enable.
	ROIE  DMAIER = 0x01 << 4  //+ Receive Overflow interrupt enable.
	TJTIE DMAIER = 0x01 << 3  //+ Transmit jabber timeout interrupt enable.
	TBUIE DMAIER = 0x01 << 2  //+ Transmit buffer unavailable interrupt enable.
	TPSIE DMAIER = 0x01 << 1  //+ Transmit process stopped interrupt enable.
	TIE   DMAIER = 0x01 << 0  //+ Transmit interrupt enable.
)

const (
	NISEn  = 16
	AISEn  = 15
	ERIEn  = 14
	FBEIEn = 13
	ETIEn  = 10
	RWTIEn = 9
	RPSIEn = 8
	RBUIEn = 7
	RIEn   = 6
	TUIEn  = 5
	ROIEn  = 4
	TJTIEn = 3
	TBUIEn = 2
	TPSIEn = 1
	TIEn   = 0
)

const (
	OFOC DMAMFBOCR = 0x01 << 28  //+ Overflow bit for FIFO overflow counter.
	MFA  DMAMFBOCR = 0x7FF << 17 //+ Number of frames missed by the application.
	OMFC DMAMFBOCR = 0x01 << 16  //+ Overflow bit for missed frame counter.
	MFC  DMAMFBOCR = 0xFFFF << 0 //+ Number of frames missed by the controller.
)

const (
	OFOCn = 28
	MFAn  = 17
	OMFCn = 16
	MFCn  = 0
)

const (
	HTDAP DMACHTDR = 0xFFFFFFFF << 0 //+ Host transmit descriptor address pointer.
)

const (
	HTDAPn = 0
)

const (
	HRDAP DMACHRDR = 0xFFFFFFFF << 0 //+ Host receive descriptor address pointer.
)

const (
	HRDAPn = 0
)

const (
	HTBAP DMACHTBAR = 0xFFFFFFFF << 0 //+ Host transmit buffer address pointer.
)

const (
	HTBAPn = 0
)

const (
	HRBAP DMACHRBAR = 0xFFFFFFFF << 0 //+ Host receive buffer address pointer.
)

const (
	HRBAPn = 0
)
