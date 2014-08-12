package main

import (
	"stm32/f4/gpio"
	"stm32/f4/periph"
	"stm32/f4/setup"
	"stm32/f4/usart"
)

var udev = usart.USART2

func init() {
	setup.Performance168(8)

	periph.AHB1ClockEnable(periph.GPIOA)
	periph.AHB1Reset(periph.GPIOA)
	periph.APB1ClockEnable(periph.USART2)
	periph.APB1Reset(periph.USART2)

	io, tx, rx := gpio.A, 2, 3

	io.SetMode(tx, gpio.Alt)
	io.SetOutType(tx, gpio.PushPullOut)
	io.SetPull(tx, gpio.PullUp)
	io.SetOutSpeed(tx, gpio.Fast)
	io.SetAltFunc(tx, gpio.USART2)
	io.SetMode(rx, gpio.Alt)

	udev.SetBaudRate(115200)
	udev.SetWordLen(usart.Bits8)
	udev.SetParity(usart.None)
	udev.SetStopBits(usart.Stop1b)
	udev.Enable()
	udev.EnableTx()
	udev.EnableRx()
}

func writeByte(b byte) {
	udev.Store(b)
	for udev.Status()&usart.TxEmpty == 0 {
	}
}

func main() {
	for {
		writeByte('H')
		writeByte('i')
		writeByte('!')
		writeByte('\r')
		writeByte('\n')
	}
}