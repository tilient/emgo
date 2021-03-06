package wwdg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type WWDG_Periph struct {
	CR  CR
	CFR CFR
	SR  SR
}

func (p *WWDG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var WWDG = (*WWDG_Periph)(unsafe.Pointer(uintptr(mmap.WWDG_BASE)))

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

func (r *CR) AtomicStoreBits(mask, b CR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CR) AtomicSetBits(mask CR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CR) AtomicClearBits(mask CR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *WWDG_Periph) T() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(T)}}
}

func (p *WWDG_Periph) WDGA() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(WDGA)}}
}

type CFR_Bits uint32

func (b CFR_Bits) Field(mask CFR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFR_Bits) J(v int) CFR_Bits {
	return CFR_Bits(bits.Make32(v, uint32(mask)))
}

type CFR struct{ mmio.U32 }

func (r *CFR) Bits(mask CFR_Bits) CFR_Bits { return CFR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFR) StoreBits(mask, b CFR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFR) SetBits(mask CFR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CFR) ClearBits(mask CFR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CFR) Load() CFR_Bits              { return CFR_Bits(r.U32.Load()) }
func (r *CFR) Store(b CFR_Bits)            { r.U32.Store(uint32(b)) }

func (r *CFR) AtomicStoreBits(mask, b CFR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CFR) AtomicSetBits(mask CFR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CFR) AtomicClearBits(mask CFR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CFR_Mask struct{ mmio.UM32 }

func (rm CFR_Mask) Load() CFR_Bits   { return CFR_Bits(rm.UM32.Load()) }
func (rm CFR_Mask) Store(b CFR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *WWDG_Periph) W() CFR_Mask {
	return CFR_Mask{mmio.UM32{&p.CFR.U32, uint32(W)}}
}

func (p *WWDG_Periph) WDGTB() CFR_Mask {
	return CFR_Mask{mmio.UM32{&p.CFR.U32, uint32(WDGTB)}}
}

func (p *WWDG_Periph) EWI() CFR_Mask {
	return CFR_Mask{mmio.UM32{&p.CFR.U32, uint32(EWI)}}
}

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

func (r *SR) AtomicStoreBits(mask, b SR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *SR) AtomicSetBits(mask SR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SR) AtomicClearBits(mask SR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *WWDG_Periph) EWIF() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(EWIF)}}
}
