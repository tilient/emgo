// Peripheral: AUX_Periph
// Instances:
//  AUX  0xe000e008
// Registers:
//  0x00 32  ACTLR Auxiliary Control Register
package scb

const (
	DISMCYCINT ACTLR_Bits = 1 << 0 //+
	DISDEFWBUF ACTLR_Bits = 1 << 1 //+
	DISFOLD    ACTLR_Bits = 1 << 2 //+
	DISFPCA    ACTLR_Bits = 1 << 8 //+
	DISOOFP    ACTLR_Bits = 1 << 9 //+
)
