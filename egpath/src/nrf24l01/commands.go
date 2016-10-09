package nrf24

func (d *Device) exec(args ...[]byte) {
	if d.Err != nil {
		return
	}
	_, d.Err = d.DCI.WriteRead(args...)
	d.Status = Status(args[1][0])
}

// Reg invokes R_REGISTER command.
func (d *Device) Reg(addr byte, val []byte) {
	cmd := []byte{addr}
	d.exec(cmd, cmd, nil, val)
}

// SetReg invokes W_REGISTER command.
func (d *Device) SetReg(addr byte, val ...byte) {
	cmd := []byte{addr | 0x20}
	d.exec(cmd, cmd, val)
}

func checkPayLen(pay []byte) {
	if len(pay) > 32 {
		panic("plen>32")
	}
}

// ReadRxP invokes R_RX_PAYLOAD command.
func (d *Device) ReadRx(pay []byte) {
	checkPayLen(pay)
	cmd := []byte{0x61}
	d.exec(cmd, cmd, nil, pay)
}

// WriteTxP invokes W_TX_PAYLOAD command.
func (d *Device) WriteTx(pay []byte) {
	checkPayLen(pay)
	cmd := []byte{0xa0}
	d.exec(cmd, cmd, pay)
}

// FlushTx invokes FLUSH_TX command.
func (d *Device) FlushTx() {
	cmd := []byte{0xe1}
	d.exec(cmd, cmd)
}

// FlushRx invokes FLUSH_RX command.
func (d *Device) FlushRx() {
	cmd := []byte{0xe2}
	d.exec(cmd, cmd)
}

// ReuseTx invokes REUSE_TX_PL command.
func (d *Device) ReuseTx() {
	cmd := []byte{0xe3}
	d.exec(cmd, cmd)
}

// Activate invokes nRF24L01 ACTIVATE command.
func (d *Device) Activate(b byte) {
	cmd := []byte{0x50}
	d.exec(cmd, cmd)
}

// RxLen invokes R_RX_PL_WID command.
func (d *Device) RxLen() int {
	cmd := []byte{0x60, 0xff}
	d.exec(cmd, cmd)
	return int(cmd[1])
}

// WriteAck invokes W_ACK_PAYLOAD command.
func (d *Device) WriteAck(pn int, pay []byte) {
	checkPayNum(pn)
	checkPayLen(pay)
	cmd := []byte{byte(0xa8 | pn)}
	d.exec(cmd, cmd, pay)
}

// WriteTxNoAck invokes W_TX_PAYLOAD_NOACK command.
func (d *Device) WriteTxNoAck(pay []byte) {
	checkPayLen(pay)
	cmd := []byte{0xa0}
	d.exec(cmd, cmd, pay)
}

// NOP invokes NOP command.
func (d *Device) NOP() {
	cmd := []byte{0xff}
	d.exec(cmd, cmd)
}
