// Peripheral: ADC_Periph  Analog to Digital Converter.
// Instances:
//  ADC1  mmap.ADC1_BASE
//  ADC2  mmap.ADC2_BASE
//  ADC3  mmap.ADC3_BASE
//  ADC4  mmap.ADC4_BASE
// Registers:
//  0x00 32  ISR     Interrupt and Status Register.
//  0x04 32  IER     Interrupt Enable Register.
//  0x08 32  CR      Control register.
//  0x0C 32  CFGR    Configuration register.
//  0x14 32  SMPR1   Sample time register 1.
//  0x18 32  SMPR2   Sample time register 2.
//  0x20 32  TR1     Watchdog threshold register 1.
//  0x24 32  TR2     Watchdog threshold register 2.
//  0x28 32  TR3     Watchdog threshold register 3.
//  0x30 32  SQR1    Regular sequence register 1.
//  0x34 32  SQR2    Regular sequence register 2.
//  0x38 32  SQR3    Regular sequence register 3.
//  0x3C 32  SQR4    Regular sequence register 4.
//  0x40 32  DR      Regular data register.
//  0x4C 32  JSQR    Injected sequence register.
//  0x60 32  OFR[4]  Offset registers.
//  0x80 32  JDR[4]  Injected data registers.
//  0xA0 32  AWD2CR  Analog Watchdog 2 Configuration Register.
//  0xA4 32  AWD3CR  Analog Watchdog 3 Configuration Register.
//  0xB0 32  DIFSEL  Differential Mode Selection Register.
//  0xB4 32  CALFACT Calibration Factors.
// Import:
//  stm32/o/f303xe/mmap
package adc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	ADRDY ISR_Bits = 0x01 << 0  //+ ADC ready flag.
	EOSMP ISR_Bits = 0x01 << 1  //+ ADC group regular end of sampling flag.
	EOC   ISR_Bits = 0x01 << 2  //+ ADC group regular end of unitary conversion flag.
	EOS   ISR_Bits = 0x01 << 3  //+ ADC group regular end of sequence conversions flag.
	OVR   ISR_Bits = 0x01 << 4  //+ ADC group regular overrun flag.
	JEOC  ISR_Bits = 0x01 << 5  //+ ADC group injected end of unitary conversion flag.
	JEOS  ISR_Bits = 0x01 << 6  //+ ADC group injected end of sequence conversions flag.
	AWD1  ISR_Bits = 0x01 << 7  //+ ADC analog watchdog 1 flag.
	AWD2  ISR_Bits = 0x01 << 8  //+ ADC analog watchdog 2 flag.
	AWD3  ISR_Bits = 0x01 << 9  //+ ADC analog watchdog 3 flag.
	JQOVF ISR_Bits = 0x01 << 10 //+ ADC group injected contexts queue overflow flag.
)

const (
	ADRDYn = 0
	EOSMPn = 1
	EOCn   = 2
	EOSn   = 3
	OVRn   = 4
	JEOCn  = 5
	JEOSn  = 6
	AWD1n  = 7
	AWD2n  = 8
	AWD3n  = 9
	JQOVFn = 10
)

const (
	ADRDYIEIE IER_Bits = 0x01 << 0  //+ ADC ready interrupt.
	EOSMPIEIE IER_Bits = 0x01 << 1  //+ ADC group regular end of sampling interrupt.
	EOCIEIE   IER_Bits = 0x01 << 2  //+ ADC group regular end of unitary conversion interrupt.
	EOSIEIE   IER_Bits = 0x01 << 3  //+ ADC group regular end of sequence conversions interrupt.
	OVRIEIE   IER_Bits = 0x01 << 4  //+ ADC group regular overrun interrupt.
	JEOCIEIE  IER_Bits = 0x01 << 5  //+ ADC group injected end of unitary conversion interrupt.
	JEOSIEIE  IER_Bits = 0x01 << 6  //+ ADC group injected end of sequence conversions interrupt.
	AWD1IEIE  IER_Bits = 0x01 << 7  //+ ADC analog watchdog 1 interrupt.
	AWD2IEIE  IER_Bits = 0x01 << 8  //+ ADC analog watchdog 2 interrupt.
	AWD3IEIE  IER_Bits = 0x01 << 9  //+ ADC analog watchdog 3 interrupt.
	JQOVFIEIE IER_Bits = 0x01 << 10 //+ ADC group injected contexts queue overflow interrupt.
)

const (
	ADRDYIEIEn = 0
	EOSMPIEIEn = 1
	EOCIEIEn   = 2
	EOSIEIEn   = 3
	OVRIEIEn   = 4
	JEOCIEIEn  = 5
	JEOSIEIEn  = 6
	AWD1IEIEn  = 7
	AWD2IEIEn  = 8
	AWD3IEIEn  = 9
	JQOVFIEIEn = 10
)

const (
	ADEN     CR_Bits = 0x01 << 0  //+ ADC enable.
	ADDIS    CR_Bits = 0x01 << 1  //+ ADC disable.
	ADSTART  CR_Bits = 0x01 << 2  //+ ADC group regular conversion start.
	JADSTART CR_Bits = 0x01 << 3  //+ ADC group injected conversion start.
	ADSTP    CR_Bits = 0x01 << 4  //+ ADC group regular conversion stop.
	JADSTP   CR_Bits = 0x01 << 5  //+ ADC group injected conversion stop.
	ADVREGEN CR_Bits = 0x03 << 28 //+ ADC voltage regulator enable.
	ADCALDIF CR_Bits = 0x01 << 30 //+ ADC differential mode for calibration.
	ADCAL    CR_Bits = 0x01 << 31 //+ ADC calibration.
)

const (
	ADENn     = 0
	ADDISn    = 1
	ADSTARTn  = 2
	JADSTARTn = 3
	ADSTPn    = 4
	JADSTPn   = 5
	ADVREGENn = 28
	ADCALDIFn = 30
	ADCALn    = 31
)

const (
	DMAEN   CFGR_Bits = 0x01 << 0  //+ ADC DMA enable.
	DMACFG  CFGR_Bits = 0x01 << 1  //+ ADC DMA configuration.
	RES     CFGR_Bits = 0x03 << 3  //+ ADC data resolution.
	ALIGN   CFGR_Bits = 0x01 << 5  //+ ADC data alignement.
	EXTSEL  CFGR_Bits = 0x0F << 6  //+ ADC group regular external trigger source.
	EXTEN   CFGR_Bits = 0x03 << 10 //+ ADC group regular external trigger polarity.
	OVRMOD  CFGR_Bits = 0x01 << 12 //+ ADC group regular overrun configuration.
	CONT    CFGR_Bits = 0x01 << 13 //+ ADC group regular continuous conversion mode.
	AUTDLY  CFGR_Bits = 0x01 << 14 //+ ADC low power auto wait.
	DISCEN  CFGR_Bits = 0x01 << 16 //+ ADC group regular sequencer discontinuous mode.
	DISCNUM CFGR_Bits = 0x07 << 17 //+ ADC Discontinuous mode channel count.
	JDISCEN CFGR_Bits = 0x01 << 20 //+ ADC Discontinuous mode on injected channels.
	JQM     CFGR_Bits = 0x01 << 21 //+ ADC group injected contexts queue mode.
	AWD1SGL CFGR_Bits = 0x01 << 22 //+ ADC analog watchdog 1 monitoring a single channel or all channels.
	AWD1EN  CFGR_Bits = 0x01 << 23 //+ ADC analog watchdog 1 enable on scope ADC group regular.
	JAWD1EN CFGR_Bits = 0x01 << 24 //+ ADC analog watchdog 1 enable on scope ADC group injected.
	JAUTO   CFGR_Bits = 0x01 << 25 //+ ADC group injected automatic trigger mode.
	AWD1CH  CFGR_Bits = 0x1F << 26 //+ ADC analog watchdog 1 monitored channel selection.
	AUTOFF  CFGR_Bits = 0x01 << 15 //+ ADC low power auto power off.
)

const (
	DMAENn   = 0
	DMACFGn  = 1
	RESn     = 3
	ALIGNn   = 5
	EXTSELn  = 6
	EXTENn   = 10
	OVRMODn  = 12
	CONTn    = 13
	AUTDLYn  = 14
	DISCENn  = 16
	DISCNUMn = 17
	JDISCENn = 20
	JQMn     = 21
	AWD1SGLn = 22
	AWD1ENn  = 23
	JAWD1ENn = 24
	JAUTOn   = 25
	AWD1CHn  = 26
	AUTOFFn  = 15
)

const (
	SMP0 SMPR1_Bits = 0x07 << 0  //+ ADC channel 0 sampling time selection.
	SMP1 SMPR1_Bits = 0x07 << 3  //+ ADC channel 1 sampling time selection.
	SMP2 SMPR1_Bits = 0x07 << 6  //+ ADC channel 2 sampling time selection.
	SMP3 SMPR1_Bits = 0x07 << 9  //+ ADC channel 3 sampling time selection.
	SMP4 SMPR1_Bits = 0x07 << 12 //+ ADC channel 4 sampling time selection.
	SMP5 SMPR1_Bits = 0x07 << 15 //+ ADC channel 5 sampling time selection.
	SMP6 SMPR1_Bits = 0x07 << 18 //+ ADC channel 6 sampling time selection.
	SMP7 SMPR1_Bits = 0x07 << 21 //+ ADC channel 7 sampling time selection.
	SMP8 SMPR1_Bits = 0x07 << 24 //+ ADC channel 8 sampling time selection.
	SMP9 SMPR1_Bits = 0x07 << 27 //+ ADC channel 9 sampling time selection.
)

const (
	SMP0n = 0
	SMP1n = 3
	SMP2n = 6
	SMP3n = 9
	SMP4n = 12
	SMP5n = 15
	SMP6n = 18
	SMP7n = 21
	SMP8n = 24
	SMP9n = 27
)

const (
	SMP10 SMPR2_Bits = 0x07 << 0  //+ ADC channel 10 sampling time selection.
	SMP11 SMPR2_Bits = 0x07 << 3  //+ ADC channel 11 sampling time selection.
	SMP12 SMPR2_Bits = 0x07 << 6  //+ ADC channel 12 sampling time selection.
	SMP13 SMPR2_Bits = 0x07 << 9  //+ ADC channel 13 sampling time selection.
	SMP14 SMPR2_Bits = 0x07 << 12 //+ ADC channel 14 sampling time selection.
	SMP15 SMPR2_Bits = 0x07 << 15 //+ ADC channel 15 sampling time selection.
	SMP16 SMPR2_Bits = 0x07 << 18 //+ ADC channel 16 sampling time selection.
	SMP17 SMPR2_Bits = 0x07 << 21 //+ ADC channel 17 sampling time selection.
	SMP18 SMPR2_Bits = 0x07 << 24 //+ ADC channel 18 sampling time selection.
)

const (
	SMP10n = 0
	SMP11n = 3
	SMP12n = 6
	SMP13n = 9
	SMP14n = 12
	SMP15n = 15
	SMP16n = 18
	SMP17n = 21
	SMP18n = 24
)

const (
	LT1 TR1_Bits = 0xFFF << 0  //+ ADC analog watchdog 1 threshold low.
	HT1 TR1_Bits = 0xFFF << 16 //+ ADC Analog watchdog 1 threshold high.
)

const (
	LT1n = 0
	HT1n = 16
)

const (
	LT2 TR2_Bits = 0xFF << 0  //+ ADC analog watchdog 2 threshold low.
	HT2 TR2_Bits = 0xFF << 16 //+ ADC analog watchdog 2 threshold high.
)

const (
	LT2n = 0
	HT2n = 16
)

const (
	LT3 TR3_Bits = 0xFF << 0  //+ ADC analog watchdog 3 threshold low.
	HT3 TR3_Bits = 0xFF << 16 //+ ADC analog watchdog 3 threshold high.
)

const (
	LT3n = 0
	HT3n = 16
)

const (
	L   SQR1_Bits = 0x0F << 0  //+ ADC group regular sequencer scan length.
	SQ1 SQR1_Bits = 0x1F << 6  //+ ADC group regular sequencer rank 1.
	SQ2 SQR1_Bits = 0x1F << 12 //+ ADC group regular sequencer rank 2.
	SQ3 SQR1_Bits = 0x1F << 18 //+ ADC group regular sequencer rank 3.
	SQ4 SQR1_Bits = 0x1F << 24 //+ ADC group regular sequencer rank 4.
)

const (
	Ln   = 0
	SQ1n = 6
	SQ2n = 12
	SQ3n = 18
	SQ4n = 24
)

const (
	SQ5 SQR2_Bits = 0x1F << 0  //+ ADC group regular sequencer rank 5.
	SQ6 SQR2_Bits = 0x1F << 6  //+ ADC group regular sequencer rank 6.
	SQ7 SQR2_Bits = 0x1F << 12 //+ ADC group regular sequencer rank 7.
	SQ8 SQR2_Bits = 0x1F << 18 //+ ADC group regular sequencer rank 8.
	SQ9 SQR2_Bits = 0x1F << 24 //+ ADC group regular sequencer rank 9.
)

const (
	SQ5n = 0
	SQ6n = 6
	SQ7n = 12
	SQ8n = 18
	SQ9n = 24
)

const (
	SQ10 SQR3_Bits = 0x1F << 0  //+ ADC group regular sequencer rank 10.
	SQ11 SQR3_Bits = 0x1F << 6  //+ ADC group regular sequencer rank 11.
	SQ12 SQR3_Bits = 0x1F << 12 //+ ADC group regular sequencer rank 12.
	SQ13 SQR3_Bits = 0x1F << 18 //+ ADC group regular sequencer rank 13.
	SQ14 SQR3_Bits = 0x1F << 24 //+ ADC group regular sequencer rank 14.
)

const (
	SQ10n = 0
	SQ11n = 6
	SQ12n = 12
	SQ13n = 18
	SQ14n = 24
)

const (
	SQ15 SQR4_Bits = 0x1F << 0 //+ ADC group regular sequencer rank 15.
	SQ16 SQR4_Bits = 0x1F << 6 //+ ADC group regular sequencer rank 16.
)

const (
	SQ15n = 0
	SQ16n = 6
)

const (
	RDATA DR_Bits = 0xFFFF << 0 //+ ADC group regular conversion data.
)

const (
	RDATAn = 0
)

const (
	JL      JSQR_Bits = 0x03 << 0  //+ ADC group injected sequencer scan length.
	JEXTSEL JSQR_Bits = 0x0F << 2  //+ ADC group injected external trigger source.
	JEXTEN  JSQR_Bits = 0x03 << 6  //+ ADC group injected external trigger polarity.
	JSQ1    JSQR_Bits = 0x1F << 8  //+ ADC group injected sequencer rank 1.
	JSQ2    JSQR_Bits = 0x1F << 14 //+ ADC group injected sequencer rank 2.
	JSQ3    JSQR_Bits = 0x1F << 20 //+ ADC group injected sequencer rank 3.
	JSQ4    JSQR_Bits = 0x1F << 26 //+ ADC group injected sequencer rank 4.
)

const (
	JLn      = 0
	JEXTSELn = 2
	JEXTENn  = 6
	JSQ1n    = 8
	JSQ2n    = 14
	JSQ3n    = 20
	JSQ4n    = 26
)

const (
	OFFSET1    OFR_Bits = 0xFFF << 0 //+ ADC offset number 1 offset level.
	OFFSET1_CH OFR_Bits = 0x1F << 26 //+ ADC offset number 1 channel selection.
	OFFSET1_EN OFR_Bits = 0x01 << 31 //+ ADC offset number 1 enable.
)

const (
	OFFSET1n    = 0
	OFFSET1_CHn = 26
	OFFSET1_ENn = 31
)

const (
	JDATA JDR_Bits = 0xFFFF << 0 //+ ADC group injected sequencer rank 1 conversion data.
)

const (
	JDATAn = 0
)

const (
	AWD2CH AWD2CR_Bits = 0x7FFFF << 0 //+ ADC analog watchdog 2 monitored channel selection.
)

const (
	AWD2CHn = 0
)

const (
	AWD3CH AWD3CR_Bits = 0x7FFFF << 0 //+ ADC analog watchdog 3 monitored channel selection.
)

const (
	AWD3CHn = 0
)

const (
	CALFACT_S CALFACT_Bits = 0x7F << 0  //+ ADC calibration factor in single-ended mode.
	CALFACT_D CALFACT_Bits = 0x7F << 16 //+ ADC calibration factor in differential mode.
)

const (
	CALFACT_Sn = 0
	CALFACT_Dn = 16
)
