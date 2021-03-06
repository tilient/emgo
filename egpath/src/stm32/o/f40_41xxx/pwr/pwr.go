// Peripheral: PWR_Periph  Power Control.
// Instances:
//  PWR  mmap.PWR_BASE
// Registers:
//  0x00 32  CR  Power control register.
//  0x04 32  CSR Power control/status register.
// Import:
//  stm32/o/f40_41xxx/mmap
package pwr

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LPDS     CR_Bits = 0x01 << 0  //+ Low-Power Deepsleep.
	PDDS     CR_Bits = 0x01 << 1  //+ Power Down Deepsleep.
	CWUF     CR_Bits = 0x01 << 2  //+ Clear Wakeup Flag.
	CSBF     CR_Bits = 0x01 << 3  //+ Clear Standby Flag.
	PVDE     CR_Bits = 0x01 << 4  //+ Power Voltage Detector Enable.
	PLS      CR_Bits = 0x07 << 5  //+ PLS[2:0] bits (PVD Level Selection).
	PLS_0    CR_Bits = 0x01 << 5  //  Bit 0.
	PLS_1    CR_Bits = 0x02 << 5  //  Bit 1.
	PLS_2    CR_Bits = 0x04 << 5  //  Bit 2.
	PLS_LEV0 CR_Bits = 0x00 << 5  //  PVD level 0.
	PLS_LEV1 CR_Bits = 0x01 << 5  //  PVD level 1.
	PLS_LEV2 CR_Bits = 0x02 << 5  //  PVD level 2.
	PLS_LEV3 CR_Bits = 0x03 << 5  //  PVD level 3.
	PLS_LEV4 CR_Bits = 0x04 << 5  //  PVD level 4.
	PLS_LEV5 CR_Bits = 0x05 << 5  //  PVD level 5.
	PLS_LEV6 CR_Bits = 0x06 << 5  //  PVD level 6.
	PLS_LEV7 CR_Bits = 0x07 << 5  //  PVD level 7.
	DBP      CR_Bits = 0x01 << 8  //+ Disable Backup Domain write protection.
	FPDS     CR_Bits = 0x01 << 9  //+ Flash power down in Stop mode.
	LPUDS    CR_Bits = 0x01 << 10 //+ Low-Power Regulator in Stop under-drive mode.
	MRUDS    CR_Bits = 0x01 << 11 //+ Main regulator in Stop under-drive mode.
	LPLVDS   CR_Bits = 0x01 << 10 //  Low-power regulator Low Voltage in Deep Sleep mode.
	MRLVDS   CR_Bits = 0x01 << 11 //  Main regulator Low Voltage in Deep Sleep mode.
	ADCDC1   CR_Bits = 0x01 << 13 //+ Refer to AN4073 on how to use this bit.
	VOS      CR_Bits = 0x03 << 14 //+ VOS[1:0] bits (Regulator voltage scaling output selection).
	VOS_0    CR_Bits = 0x01 << 14 //  Bit 0.
	VOS_1    CR_Bits = 0x02 << 14 //  Bit 1.
	ODEN     CR_Bits = 0x01 << 16 //+ Over Drive enable.
	ODSWEN   CR_Bits = 0x01 << 17 //+ Over Drive switch enabled.
	UDEN     CR_Bits = 0x03 << 18 //+ Under Drive enable in stop mode.
	UDEN_0   CR_Bits = 0x01 << 18 //  Bit 0.
	UDEN_1   CR_Bits = 0x02 << 18 //  Bit 1.
	FMSSR    CR_Bits = 0x01 << 20 //+ Flash Memory Sleep System Run.
	FISSR    CR_Bits = 0x01 << 21 //+ Flash Interface Stop while System Run.
)

const (
	LPDSn   = 0
	PDDSn   = 1
	CWUFn   = 2
	CSBFn   = 3
	PVDEn   = 4
	PLSn    = 5
	DBPn    = 8
	FPDSn   = 9
	LPUDSn  = 10
	MRUDSn  = 11
	ADCDC1n = 13
	VOSn    = 14
	ODENn   = 16
	ODSWENn = 17
	UDENn   = 18
	FMSSRn  = 20
	FISSRn  = 21
)

const (
	WUF     CSR_Bits = 0x01 << 0  //+ Wakeup Flag.
	SBF     CSR_Bits = 0x01 << 1  //+ Standby Flag.
	PVDO    CSR_Bits = 0x01 << 2  //+ PVD Output.
	BRR     CSR_Bits = 0x01 << 3  //+ Backup regulator ready.
	WUPP    CSR_Bits = 0x01 << 7  //+ WKUP pin Polarity.
	EWUP    CSR_Bits = 0x01 << 8  //+ Enable WKUP pin.
	BRE     CSR_Bits = 0x01 << 9  //+ Backup regulator enable.
	VOSRDY  CSR_Bits = 0x01 << 14 //+ Regulator voltage scaling output selection ready.
	ODRDY   CSR_Bits = 0x01 << 16 //+ Over Drive generator ready.
	ODSWRDY CSR_Bits = 0x01 << 17 //+ Over Drive Switch ready.
	UDSWRDY CSR_Bits = 0x03 << 18 //+ Under Drive ready.
)

const (
	WUFn     = 0
	SBFn     = 1
	PVDOn    = 2
	BRRn     = 3
	WUPPn    = 7
	EWUPn    = 8
	BREn     = 9
	VOSRDYn  = 14
	ODRDYn   = 16
	ODSWRDYn = 17
	UDSWRDYn = 18
)
