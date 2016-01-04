// +build f40_41xxx f411xe

package gpio

import (
	"mmio"
	"unsafe"

	"stm32/hal/raw/mmap"
	"stm32/hal/raw/rcc"
)

const (
	Low      Speed = 0 //   2 MHz (CL = 50 pF, VDD > 2.7 V)
	Medium   Speed = 1 //  25 MHz (CL = 50 pF, VDD > 2.7 V)
	High     Speed = 2 //  50 MHz (CL = 40 pF, VDD > 2.7 V)
	VeryHigh Speed = 3 // 100 MHz (CL = 30 pF, VDD > 2.7 V)
)

const (
	TIM1 = AF1
	TIM2 = AF1

	TIM3 = AF2
	TIM4 = AF2
	TIM5 = AF2

	TIM8  = AF3
	TIM9  = AF3
	TIM10 = AF3
	TIM11 = AF3

	I2C1 = AF4
	I2C2 = AF4
	I2C3 = AF4

	SPI1 = AF5
	SPI2 = AF5
	SPI4 = AF5
	SPI5 = AF5
	SPI6 = AF5

	SPI3 = AF6
	SAI1 = AF6

	USART1 = AF7
	USART2 = AF7
	USART3 = AF7

	UART4  = AF8
	UART5  = AF8
	USART6 = AF8
	UART7  = AF8
	UART8  = AF8

	CAN1  = AF9
	CAN2  = AF9
	TIM12 = AF9
	TIM13 = AF9
	TIM14 = AF9

	OTGFS = AF10
	OTGHS = AF10

	Eth = AF11

	FSMC    = AF12
	FMC     = AF12
	SDIO    = AF12
	OTGHSFS = AF12

	DCMI = AF13

	LTDC = AF14
)

var (
	A = (*Port)(unsafe.Pointer(mmap.GPIOA_BASE))
	B = (*Port)(unsafe.Pointer(mmap.GPIOB_BASE))
	C = (*Port)(unsafe.Pointer(mmap.GPIOC_BASE))
	D = (*Port)(unsafe.Pointer(mmap.GPIOD_BASE))
	E = (*Port)(unsafe.Pointer(mmap.GPIOE_BASE))
	F = (*Port)(unsafe.Pointer(mmap.GPIOF_BASE))
	G = (*Port)(unsafe.Pointer(mmap.GPIOG_BASE))
	H = (*Port)(unsafe.Pointer(mmap.GPIOH_BASE))
	I = (*Port)(unsafe.Pointer(mmap.GPIOI_BASE))
	J = (*Port)(unsafe.Pointer(mmap.GPIOJ_BASE))
	K = (*Port)(unsafe.Pointer(mmap.GPIOK_BASE))
)

func pnum(p *Port) int {
	return int(uintptr(unsafe.Pointer(p))-mmap.AHB1PERIPH_BASE) / 0x400
}

func enr() *mmio.U32   { return &rcc.RCC.AHB1ENR.U32 }
func lpenr() *mmio.U32 { return &rcc.RCC.AHB1LPENR.U32 }
func rstr() *mmio.U32  { return &rcc.RCC.AHB1RSTR.U32 }