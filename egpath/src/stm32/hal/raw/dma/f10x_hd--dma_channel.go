// +build f10x_hd

// Peripheral: DMA_Channel_Periph  DMA Controller.
// Instances:
//  DMA1_Channel1  mmap.DMA1_Channel1_BASE
//  DMA1_Channel2  mmap.DMA1_Channel2_BASE
//  DMA1_Channel3  mmap.DMA1_Channel3_BASE
//  DMA1_Channel4  mmap.DMA1_Channel4_BASE
//  DMA1_Channel5  mmap.DMA1_Channel5_BASE
//  DMA1_Channel6  mmap.DMA1_Channel6_BASE
//  DMA1_Channel7  mmap.DMA1_Channel7_BASE
//  DMA2_Channel1  mmap.DMA2_Channel1_BASE
//  DMA2_Channel2  mmap.DMA2_Channel2_BASE
//  DMA2_Channel3  mmap.DMA2_Channel3_BASE
//  DMA2_Channel4  mmap.DMA2_Channel4_BASE
//  DMA2_Channel5  mmap.DMA2_Channel5_BASE
// Registers:
//  0x00 32  CCR
//  0x04 32  CNDTR
//  0x08 32  CPAR
//  0x0C 32  CMAR
// Import:
//  stm32/o/f10x_hd/mmap
package dma

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	EN      CCR_Bits = 0x01 << 0  //+ Channel enable.
	TCIE    CCR_Bits = 0x01 << 1  //+ Transfer complete interrupt enable.
	HTIE    CCR_Bits = 0x01 << 2  //+ Half Transfer interrupt enable.
	TEIE    CCR_Bits = 0x01 << 3  //+ Transfer error interrupt enable.
	DIR     CCR_Bits = 0x01 << 4  //+ Data transfer direction.
	CIRC    CCR_Bits = 0x01 << 5  //+ Circular mode.
	PINC    CCR_Bits = 0x01 << 6  //+ Peripheral increment mode.
	MINC    CCR_Bits = 0x01 << 7  //+ Memory increment mode.
	PSIZE   CCR_Bits = 0x03 << 8  //+ PSIZE[1:0] bits (Peripheral size).
	PSIZE_0 CCR_Bits = 0x01 << 8  //  Bit 0.
	PSIZE_1 CCR_Bits = 0x02 << 8  //  Bit 1.
	MSIZE   CCR_Bits = 0x03 << 10 //+ MSIZE[1:0] bits (Memory size).
	MSIZE_0 CCR_Bits = 0x01 << 10 //  Bit 0.
	MSIZE_1 CCR_Bits = 0x02 << 10 //  Bit 1.
	PL      CCR_Bits = 0x03 << 12 //+ PL[1:0] bits(Channel Priority level).
	PL_0    CCR_Bits = 0x01 << 12 //  Bit 0.
	PL_1    CCR_Bits = 0x02 << 12 //  Bit 1.
	MEM2MEM CCR_Bits = 0x01 << 14 //+ Memory to memory mode.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR_Bits = 0xFFFF << 0 //+ Number of data to Transfer.
)

const (
	NDTn = 0
)

const (
	PA CPAR_Bits = 0xFFFFFFFF << 0 //+ Peripheral Address.
)

const (
	PAn = 0
)

const (
	MA CMAR_Bits = 0xFFFFFFFF << 0 //+ Memory Address.
)

const (
	MAn = 0
)
