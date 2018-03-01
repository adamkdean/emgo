// Peripheral: USB_OTG_HostChannel_Periph  USB_OTG_Host_Channel_Specific_Registers.
// Instances:
// Registers:
//  0x00 32  HCCHAR   Host Channel Characteristics Register    500h.
//  0x04 32  HCSPLT   Host Channel Split Control Register      504h.
//  0x08 32  HCINT    Host Channel Interrupt Register          508h.
//  0x0C 32  HCINTMSK Host Channel Interrupt Mask Register     50Ch.
//  0x10 32  HCTSIZ   Host Channel Transfer Size Register      510h.
//  0x14 32  HCDMA    Host Channel DMA Address Register        514h.
// Import:
//  stm32/o/f746xx/mmap
package usb

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MPSIZ   HCCHAR = 0x7FF << 0 //+ Maximum packet size.
	EPNUM   HCCHAR = 0x0F << 11 //+ Endpoint number.
	EPNUM_0 HCCHAR = 0x01 << 11 //  Bit 0.
	EPNUM_1 HCCHAR = 0x02 << 11 //  Bit 1.
	EPNUM_2 HCCHAR = 0x04 << 11 //  Bit 2.
	EPNUM_3 HCCHAR = 0x08 << 11 //  Bit 3.
	EPDIR   HCCHAR = 0x01 << 15 //+ Endpoint direction.
	LSDEV   HCCHAR = 0x01 << 17 //+ Low-speed device.
	EPTYP   HCCHAR = 0x03 << 18 //+ Endpoint type.
	EPTYP_0 HCCHAR = 0x01 << 18 //  Bit 0.
	EPTYP_1 HCCHAR = 0x02 << 18 //  Bit 1.
	MC      HCCHAR = 0x03 << 20 //+ Multi Count (MC) / Error Count (EC).
	MC_0    HCCHAR = 0x01 << 20 //  Bit 0.
	MC_1    HCCHAR = 0x02 << 20 //  Bit 1.
	DAD     HCCHAR = 0x7F << 22 //+ Device address.
	DAD_0   HCCHAR = 0x01 << 22 //  Bit 0.
	DAD_1   HCCHAR = 0x02 << 22 //  Bit 1.
	DAD_2   HCCHAR = 0x04 << 22 //  Bit 2.
	DAD_3   HCCHAR = 0x08 << 22 //  Bit 3.
	DAD_4   HCCHAR = 0x10 << 22 //  Bit 4.
	DAD_5   HCCHAR = 0x20 << 22 //  Bit 5.
	DAD_6   HCCHAR = 0x40 << 22 //  Bit 6.
	ODDFRM  HCCHAR = 0x01 << 29 //+ Odd frame.
	CHDIS   HCCHAR = 0x01 << 30 //+ Channel disable.
	CHENA   HCCHAR = 0x01 << 31 //+ Channel enable.
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCn     = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	PRTADDR   HCSPLT = 0x7F << 0  //+ Port address.
	PRTADDR_0 HCSPLT = 0x01 << 0  //  Bit 0.
	PRTADDR_1 HCSPLT = 0x02 << 0  //  Bit 1.
	PRTADDR_2 HCSPLT = 0x04 << 0  //  Bit 2.
	PRTADDR_3 HCSPLT = 0x08 << 0  //  Bit 3.
	PRTADDR_4 HCSPLT = 0x10 << 0  //  Bit 4.
	PRTADDR_5 HCSPLT = 0x20 << 0  //  Bit 5.
	PRTADDR_6 HCSPLT = 0x40 << 0  //  Bit 6.
	HUBADDR   HCSPLT = 0x7F << 7  //+ Hub address.
	HUBADDR_0 HCSPLT = 0x01 << 7  //  Bit 0.
	HUBADDR_1 HCSPLT = 0x02 << 7  //  Bit 1.
	HUBADDR_2 HCSPLT = 0x04 << 7  //  Bit 2.
	HUBADDR_3 HCSPLT = 0x08 << 7  //  Bit 3.
	HUBADDR_4 HCSPLT = 0x10 << 7  //  Bit 4.
	HUBADDR_5 HCSPLT = 0x20 << 7  //  Bit 5.
	HUBADDR_6 HCSPLT = 0x40 << 7  //  Bit 6.
	XACTPOS   HCSPLT = 0x03 << 14 //+ XACTPOS.
	XACTPOS_0 HCSPLT = 0x01 << 14 //  Bit 0.
	XACTPOS_1 HCSPLT = 0x02 << 14 //  Bit 1.
	COMPLSPLT HCSPLT = 0x01 << 16 //+ Do complete split.
	SPLITEN   HCSPLT = 0x01 << 31 //+ Split enable.
)

const (
	PRTADDRn   = 0
	HUBADDRn   = 7
	XACTPOSn   = 14
	COMPLSPLTn = 16
	SPLITENn   = 31
)

const (
	XFRC   HCINT = 0x01 << 0  //+ Transfer completed.
	CHH    HCINT = 0x01 << 1  //+ Channel halted.
	AHBERR HCINT = 0x01 << 2  //+ AHB error.
	STALL  HCINT = 0x01 << 3  //+ STALL response received interrupt.
	NAK    HCINT = 0x01 << 4  //+ NAK response received interrupt.
	ACK    HCINT = 0x01 << 5  //+ ACK response received/transmitted interrupt.
	NYET   HCINT = 0x01 << 6  //+ Response received interrupt.
	TXERR  HCINT = 0x01 << 7  //+ Transaction error.
	BBERR  HCINT = 0x01 << 8  //+ Babble error.
	FRMOR  HCINT = 0x01 << 9  //+ Frame overrun.
	DTERR  HCINT = 0x01 << 10 //+ Data toggle error.
)

const (
	XFRCn   = 0
	CHHn    = 1
	AHBERRn = 2
	STALLn  = 3
	NAKn    = 4
	ACKn    = 5
	NYETn   = 6
	TXERRn  = 7
	BBERRn  = 8
	FRMORn  = 9
	DTERRn  = 10
)

const (
	XFRCM  HCINTMSK = 0x01 << 0  //+ Transfer completed mask.
	CHHM   HCINTMSK = 0x01 << 1  //+ Channel halted mask.
	AHBERR HCINTMSK = 0x01 << 2  //+ AHB error.
	STALLM HCINTMSK = 0x01 << 3  //+ STALL response received interrupt mask.
	NAKM   HCINTMSK = 0x01 << 4  //+ NAK response received interrupt mask.
	ACKM   HCINTMSK = 0x01 << 5  //+ ACK response received/transmitted interrupt mask.
	NYET   HCINTMSK = 0x01 << 6  //+ response received interrupt mask.
	TXERRM HCINTMSK = 0x01 << 7  //+ Transaction error mask.
	BBERRM HCINTMSK = 0x01 << 8  //+ Babble error mask.
	FRMORM HCINTMSK = 0x01 << 9  //+ Frame overrun mask.
	DTERRM HCINTMSK = 0x01 << 10 //+ Data toggle error mask.
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	AHBERRn = 2
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ HCTSIZ = 0x7FFFF << 0 //+ Transfer size.
	PKTCNT HCTSIZ = 0x3FF << 19  //+ Packet count.
	DOPING HCTSIZ = 0x01 << 31   //+ Do PING.
	DPID   HCTSIZ = 0x03 << 29   //+ Data PID.
	DPID_0 HCTSIZ = 0x01 << 29   //  Bit 0.
	DPID_1 HCTSIZ = 0x02 << 29   //  Bit 1.
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DOPINGn = 31
	DPIDn   = 29
)

const (
	DMAADDR HCDMA = 0xFFFFFFFF << 0 //+ DMA address.
)

const (
	DMAADDRn = 0
)
