package i2c

import (
	"rtos"
	"sync"

	"arch/cortexm/nvic"

	"stm32/hal/raw/i2c"
)

// Driver implements polling and interrupt driven driver to I2C peripheral.
// Default mode is polling .
type Driver struct {
	Periph

	mutex  sync.Mutex
	evflag rtos.EventFlag
	evirq  nvic.IRQ
	errirq nvic.IRQ
}

func (d *Driver) SetIntMode(evirq, errirq nvic.IRQ) {
	d.evirq = evirq
	d.errirq = errirq
	d.Periph.raw.CR2.SetBits(i2c.ITBUFEN | i2c.ITEVTEN | i2c.ITERREN)
}

func (d *Driver) SetPollMode() {
	d.evirq = 0
	d.errirq = 0
	d.Periph.raw.CR2.ClearBits(i2c.ITBUFEN | i2c.ITEVTEN | i2c.ITERREN)
}

func (d *Driver) ISR() {
	d.evirq.Disable()
	d.errirq.Disable()
	d.evflag.Set()
}

type Error int16

const (
	BusErr   Error = 1 << 0
	ArbLost  Error = 1 << 1
	AckFail  Error = 1 << 2
	Overrun  Error = 1 << 3
	PECErr   Error = 1 << 4
	Timeout  Error = 1 << 6
	SMBAlert Error = 1 << 7

	SoftTimeout Error = 1 << 8
	BelatedStop Error = 1 << 9
)

func (e Error) Error() string {
	return "I2C error"
}

func (d *Driver) waitIRQ(ev i2c.SR1_Bits, deadline int64) Error {
	for {
		rtos.IRQ(d.evirq).Enable()
		rtos.IRQ(d.errirq).Enable()
		if !d.evflag.Wait(deadline) {
			return SoftTimeout
		}
		d.evflag.Clear()
		sr1 := d.Periph.raw.SR1.Load()
		if e := Error(sr1 >> 8); e != 0 {
			return e
		}
		if sr1&ev != 0 {
			return 0
		}
	}
}

func (d *Driver) pollEvent(ev i2c.SR1_Bits, deadline int64) Error {
	for {
		sr1 := d.raw.SR1.Load()
		if e := Error(sr1 >> 8); e != 0 {
			return e
		}
		if sr1&ev != 0 {
			return 0
		}
		if rtos.Nanosec() >= deadline {
			return SoftTimeout
		}
	}
}

func (d *Driver) waitEvent(ev i2c.SR1_Bits) Error {
	deadline := rtos.Nanosec() + 100e6 // 100 ms
	if d.evirq == 0 {
		return d.pollEvent(ev, deadline)
	}
	return d.waitIRQ(ev, deadline)
}

func (d *Driver) MasterConn(addr int16) MasterConn {
	return MasterConn{d: d, addr: uint16(addr << 1)}
	// TODO: Add support for 10-bit addr.
}