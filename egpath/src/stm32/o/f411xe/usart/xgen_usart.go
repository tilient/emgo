package usart

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type USART_Periph struct {
	SR   SR
	_    uint16
	DR   DR
	_    uint16
	BRR  BRR
	_    uint16
	CR1  CR1
	_    uint16
	CR2  CR2
	_    uint16
	CR3  CR3
	_    uint16
	GTPR GTPR
}

func (p *USART_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var USART2 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART2_BASE)))

//emgo:const
var USART3 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART3_BASE)))

//emgo:const
var UART4 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART4_BASE)))

//emgo:const
var UART5 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART5_BASE)))

//emgo:const
var UART7 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART7_BASE)))

//emgo:const
var UART8 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART8_BASE)))

//emgo:const
var USART1 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART1_BASE)))

//emgo:const
var USART6 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART6_BASE)))

type SR_Bits uint16

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U16 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U16.Bits(uint16(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U16.SetBits(uint16(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U16.ClearBits(uint16(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U16.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U16.Store(uint16(b)) }

type SR_Mask struct{ mmio.UM16 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM16.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) PE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(PE)}}
}

func (p *USART_Periph) FE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(FE)}}
}

func (p *USART_Periph) NE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(NE)}}
}

func (p *USART_Periph) ORE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(ORE)}}
}

func (p *USART_Periph) IDLE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(IDLE)}}
}

func (p *USART_Periph) RXNE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(RXNE)}}
}

func (p *USART_Periph) TC() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(TC)}}
}

func (p *USART_Periph) TXE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(TXE)}}
}

func (p *USART_Periph) LBD() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(LBD)}}
}

func (p *USART_Periph) CTS() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CTS)}}
}

type DR_Bits uint16

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U16 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U16.SetBits(uint16(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U16.ClearBits(uint16(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U16.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U16.Store(uint16(b)) }

type DR_Mask struct{ mmio.UM16 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM16.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM16.Store(uint16(b)) }

type BRR_Bits uint16

func (b BRR_Bits) Field(mask BRR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BRR_Bits) J(v int) BRR_Bits {
	return BRR_Bits(bits.Make32(v, uint32(mask)))
}

type BRR struct{ mmio.U16 }

func (r *BRR) Bits(mask BRR_Bits) BRR_Bits { return BRR_Bits(r.U16.Bits(uint16(mask))) }
func (r *BRR) StoreBits(mask, b BRR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *BRR) SetBits(mask BRR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *BRR) ClearBits(mask BRR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *BRR) Load() BRR_Bits              { return BRR_Bits(r.U16.Load()) }
func (r *BRR) Store(b BRR_Bits)            { r.U16.Store(uint16(b)) }

type BRR_Mask struct{ mmio.UM16 }

func (rm BRR_Mask) Load() BRR_Bits   { return BRR_Bits(rm.UM16.Load()) }
func (rm BRR_Mask) Store(b BRR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) DIV_Fraction() BRR_Mask {
	return BRR_Mask{mmio.UM16{&p.BRR.U16, uint16(DIV_Fraction)}}
}

func (p *USART_Periph) DIV_Mantissa() BRR_Mask {
	return BRR_Mask{mmio.UM16{&p.BRR.U16, uint16(DIV_Mantissa)}}
}

type CR1_Bits uint16

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U16 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U16.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U16.Store(uint16(b)) }

type CR1_Mask struct{ mmio.UM16 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM16.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) SBK() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(SBK)}}
}

func (p *USART_Periph) RWU() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(RWU)}}
}

func (p *USART_Periph) RE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(RE)}}
}

func (p *USART_Periph) TE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(TE)}}
}

func (p *USART_Periph) IDLEIE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(IDLEIE)}}
}

func (p *USART_Periph) RXNEIE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(RXNEIE)}}
}

func (p *USART_Periph) TCIE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(TCIE)}}
}

func (p *USART_Periph) TXEIE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(TXEIE)}}
}

func (p *USART_Periph) PEIE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(PEIE)}}
}

func (p *USART_Periph) PS() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(PS)}}
}

func (p *USART_Periph) PCE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(PCE)}}
}

func (p *USART_Periph) WAKE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(WAKE)}}
}

func (p *USART_Periph) M() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(M)}}
}

func (p *USART_Periph) UE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(UE)}}
}

func (p *USART_Periph) OVER8() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(OVER8)}}
}

type CR2_Bits uint16

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U16 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U16.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U16.Store(uint16(b)) }

type CR2_Mask struct{ mmio.UM16 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM16.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) ADD() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(ADD)}}
}

func (p *USART_Periph) LBDL() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(LBDL)}}
}

func (p *USART_Periph) LBDIE() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(LBDIE)}}
}

func (p *USART_Periph) LBCL() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(LBCL)}}
}

func (p *USART_Periph) CPHA() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(CPHA)}}
}

func (p *USART_Periph) CPOL() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(CPOL)}}
}

func (p *USART_Periph) CLKEN() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(CLKEN)}}
}

func (p *USART_Periph) STOP() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(STOP)}}
}

func (p *USART_Periph) LINEN() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(LINEN)}}
}

type CR3_Bits uint16

func (b CR3_Bits) Field(mask CR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR3_Bits) J(v int) CR3_Bits {
	return CR3_Bits(bits.Make32(v, uint32(mask)))
}

type CR3 struct{ mmio.U16 }

func (r *CR3) Bits(mask CR3_Bits) CR3_Bits { return CR3_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR3) StoreBits(mask, b CR3_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR3) SetBits(mask CR3_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR3) ClearBits(mask CR3_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR3) Load() CR3_Bits              { return CR3_Bits(r.U16.Load()) }
func (r *CR3) Store(b CR3_Bits)            { r.U16.Store(uint16(b)) }

type CR3_Mask struct{ mmio.UM16 }

func (rm CR3_Mask) Load() CR3_Bits   { return CR3_Bits(rm.UM16.Load()) }
func (rm CR3_Mask) Store(b CR3_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) EIE() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(EIE)}}
}

func (p *USART_Periph) IREN() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(IREN)}}
}

func (p *USART_Periph) IRLP() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(IRLP)}}
}

func (p *USART_Periph) HDSEL() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(HDSEL)}}
}

func (p *USART_Periph) NACK() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(NACK)}}
}

func (p *USART_Periph) SCEN() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(SCEN)}}
}

func (p *USART_Periph) DMAR() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(DMAR)}}
}

func (p *USART_Periph) DMAT() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(DMAT)}}
}

func (p *USART_Periph) RTSE() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(RTSE)}}
}

func (p *USART_Periph) CTSE() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(CTSE)}}
}

func (p *USART_Periph) CTSIE() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(CTSIE)}}
}

func (p *USART_Periph) ONEBIT() CR3_Mask {
	return CR3_Mask{mmio.UM16{&p.CR3.U16, uint16(ONEBIT)}}
}

type GTPR_Bits uint16

func (b GTPR_Bits) Field(mask GTPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask GTPR_Bits) J(v int) GTPR_Bits {
	return GTPR_Bits(bits.Make32(v, uint32(mask)))
}

type GTPR struct{ mmio.U16 }

func (r *GTPR) Bits(mask GTPR_Bits) GTPR_Bits { return GTPR_Bits(r.U16.Bits(uint16(mask))) }
func (r *GTPR) StoreBits(mask, b GTPR_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *GTPR) SetBits(mask GTPR_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *GTPR) ClearBits(mask GTPR_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *GTPR) Load() GTPR_Bits               { return GTPR_Bits(r.U16.Load()) }
func (r *GTPR) Store(b GTPR_Bits)             { r.U16.Store(uint16(b)) }

type GTPR_Mask struct{ mmio.UM16 }

func (rm GTPR_Mask) Load() GTPR_Bits   { return GTPR_Bits(rm.UM16.Load()) }
func (rm GTPR_Mask) Store(b GTPR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) PSC() GTPR_Mask {
	return GTPR_Mask{mmio.UM16{&p.GTPR.U16, uint16(PSC)}}
}

func (p *USART_Periph) GT() GTPR_Mask {
	return GTPR_Mask{mmio.UM16{&p.GTPR.U16, uint16(GT)}}
}
