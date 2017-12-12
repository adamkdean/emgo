package eve

import (
	"delay"
	"errors"
)

const chipidAddr = 0x0C0000

type mmap struct {
	ramdl       int
	ramcmd      int
	regdlswap   int
	regintflags int
}

//emgo:const
var eve1 = mmap{
	ramdl:       0x100000,
	ramcmd:      0x108000,
	regdlswap:   0x102450,
	regintflags: 0x102498,
}

//emgo:const
var eve2 = mmap{
	ramdl:       0x300000,
	ramcmd:      0x308000,
	regdlswap:   0x302054,
	regintflags: 0x3020a8,
}

// Register offsets relative to REG_DLSWAP.
const (
	ohcycle  = -40
	oswizzle = 16
	opclk    = 28
	ogpio    = 64
)

// Register offsets relative to REG_INT_FLAGS.
const (
	ointen   = 4
	ointmask = 8
	opwmduty = 44
)

type DisplayConfig struct {
	Hcycle  uint16 // Total number of clocks per line.
	Hsize   uint16 // Active width of LCD display.
	Hsync0  byte   // Start of horizontal sync pulse.
	Hsync1  byte   // End of horizontal sync pulse.
	Hoffset byte   // Start of active line.
	ClkPol  byte   // Define active edge of pixel clock.
	Vcycle  uint16 // Total number of lines per screen.
	Vsize   uint16 // Active height of LCD display.
	Vsync0  byte   // Start of vertical sync pulse.
	Vsync1  byte   // End of vertical sync pulse.
	Voffset byte   // Start of active screen.
	ClkPre  byte   // Pixel Clock prescaler.
	Swizzle byte   // Arrangement of output RGB pins.
	Spreed  byte   // Color signals spread, reduces EM noise
}

//emgo:const
var Default480x272 = DisplayConfig{
	Hcycle: 548, Hsize: 480, Hsync0: 0, Hsync1: 41, Hoffset: 43,
	Vcycle: 292, Vsize: 272, Vsync0: 0, Vsync1: 10, Voffset: 12,
	ClkPol: 1, ClkPre: 5,
}

// Host commands for initialisation.
const (
	cmdActive = 0
	cmdClkExt = 0x44
)

func (d *Driver) Init(cfg *DisplayConfig) error {
	d.dci.SetPDN(0)
	delay.Millisec(20)
	d.dci.SetPDN(1)
	delay.Millisec(20) // Wait 20 ms for internal oscilator and PLL.
	d.HostCmd(cmdActive, 0)

	/*
		// Simple triming algorithm if internal oscilator is used.
		for trim := uint32(0); trim <= 31; trim++ {
			d.W(ft80.REG_TRIM).W32(trim)
			if f := curFreq(d); f > 47040000 {
				lcd.W(ft80.REG_FREQUENCY).W32(f)
				break
			}
		}
	*/

	// Select external 12 MHz oscilator as clock source.
	d.HostCmd(cmdClkExt, 0)

	if err := d.Err(true); err != nil {
		return err
	}

	chipid := d.ReadUint32(chipidAddr)
	switch {
	case chipid == 0x10008:
		d.mmap = &eve1
	case 0x11008 <= chipid && chipid <= 0x111308:
		d.mmap = &eve2
	default:
		return errors.New("eve: unknown controller")
	}

	d.SetBacklight(0)
	d.SetIntMask(0)
	d.WriteByte(d.mmap.regintflags+ointen, 1)
	w := d.W(d.mmap.regdlswap + oswizzle)
	w.wr32(uint32(cfg.Swizzle))
	w.wr32(uint32(cfg.Spreed))
	w.wr32(uint32(cfg.ClkPol))
	w.wr32(0)
	w = d.W(d.mmap.regdlswap + ohcycle)
	w.wr32(uint32(cfg.Hcycle))
	w.wr32(uint32(cfg.Hoffset))
	w.wr32(uint32(cfg.Hsize))
	w.wr32(uint32(cfg.Hsync0))
	w.wr32(uint32(cfg.Hsync1))
	w.wr32(uint32(cfg.Vcycle))
	w.wr32(uint32(cfg.Voffset))
	w.wr32(uint32(cfg.Vsize))
	w.wr32(uint32(cfg.Vsync0))
	w.wr32(uint32(cfg.Vsync1))
	w = d.W(d.mmap.ramdl)
	w.wr32(CLEAR | CST)
	w.wr32(DISPLAY)
	d.SwapDL()
	b := d.ReadByte(d.mmap.regdlswap + ogpio)
	d.WriteByte(d.mmap.regdlswap+ogpio, b|0x80)     // Set DISP high.
	d.WriteByte(d.mmap.regdlswap+opclk, cfg.ClkPre) // Enable PCLK.

	delay.Millisec(20) // Wait for new main clock.

	return d.Err(true)
}

// SetBacklight sets backlight PWM duty cycle. Pwmduty range is from 0 to 128.
func (d *Driver) SetBacklight(pwmduty int) {
	d.WriteByte(d.mmap.regintflags+opwmduty, byte(pwmduty))
}

// IntFlags reads interrupt flags from EVE and accumulates them (using logical
// OR) in internal variable. It returns accumulated flags.
func (d *Driver) IntFlags() byte {
	d.intflags |= d.ReadByte(d.mmap.regintflags)
	return d.intflags
}

// ClearFlags clears interrupt flags specified by mask.
func (d *Driver) ClearFlags(mask byte) {
	d.intflags |= d.ReadByte(d.mmap.regintflags)
	d.intflags &^= mask
}

// IntMask returns current interrupt mask.
func (d *Driver) IntMask() byte {
	return d.ReadByte(d.mmap.regintflags + ointmask)
}

// SetIntMask sets interrupt mask.
func (d *Driver) SetIntMask(mask byte) {
	d.WriteByte(d.mmap.regintflags+ointmask, mask)
}

// SwapDL swaps display lists
func (d *Driver) SwapDL() {
	d.ClearFlags(INT_SWAP)
	d.WriteByte(d.mmap.regdlswap, DLSWAP_FRAME)
	mask := d.IntMask()
	d.SetIntMask(INT_SWAP)
	for {
		<-d.IRQ()
		if d.IntFlags()&INT_SWAP != 0 {
			d.SetIntMask(mask)
		}
	}
}