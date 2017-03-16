package bufferpack

import "github.com/WesleiRamos/pystruct"

func (self *Buffer) WriteByte(value byte) {
	self.buffer = append(self.buffer, pystruct.Pack("b", value)...)
}
func (self *Buffer) WriteChar(value int32) {
	self.buffer = append(self.buffer, pystruct.Pack("c", value)...)
}

func (self *Buffer) WriteBool(value bool) {
	self.buffer = append(self.buffer, pystruct.Pack("?", value)...)
}
func (self *Buffer) WriteString(value string) {
	self.buffer = append(self.buffer, pystruct.Pack("s", value)...)
}

// LITTLE ENDIAN

func (self *Buffer) WriteIntLE(value int) {
	self.buffer = append(self.buffer, pystruct.Pack("i", value)...)
}
func (self *Buffer) WriteUIntLE(value uint) {
	self.buffer = append(self.buffer, pystruct.Pack("I", value)...)
}

func (self *Buffer) WriteInt32LE(value int32) {
	self.buffer = append(self.buffer, pystruct.Pack("l", value)...)
}
func (self *Buffer) WriteUInt32LE(value uint32) {
	self.buffer = append(self.buffer, pystruct.Pack("L", value)...)
}

func (self *Buffer) WriteInt64LE(value int64) {
	self.buffer = append(self.buffer, pystruct.Pack("q", value)...)
}
func (self *Buffer) WriteUInt64LE(value uint64) {
	self.buffer = append(self.buffer, pystruct.Pack("Q", value)...)
}

func (self *Buffer) WriteShortLE(value int) {
	self.buffer = append(self.buffer, pystruct.Pack("h", value)...)
}
func (self *Buffer) WriteUShortLE(value uint) {
	self.buffer = append(self.buffer, pystruct.Pack("H", value)...)
}

func (self *Buffer) WriteFloatLE(value float32) {
	self.buffer = append(self.buffer, pystruct.Pack("f", value)...)
}
func (self *Buffer) WriteDoubleLE(value float64) {
	self.buffer = append(self.buffer, pystruct.Pack("d", value)...)
}

// BIG ENDIAN

func (self *Buffer) WriteIntBE(value int) {
	self.buffer = append(self.buffer, pystruct.Pack("!i", value)...)
}
func (self *Buffer) WriteUIntBE(value uint) {
	self.buffer = append(self.buffer, pystruct.Pack("!I", value)...)
}

func (self *Buffer) WriteInt32BE(value int32) {
	self.buffer = append(self.buffer, pystruct.Pack("!l", value)...)
}
func (self *Buffer) WriteUInt32BE(value uint32) {
	self.buffer = append(self.buffer, pystruct.Pack("!L", value)...)
}

func (self *Buffer) WriteInt64BE(value int64) {
	self.buffer = append(self.buffer, pystruct.Pack("!q", value)...)
}
func (self *Buffer) WriteUInt64BE(value uint64) {
	self.buffer = append(self.buffer, pystruct.Pack("!Q", value)...)
}

func (self *Buffer) WriteShortBE(value int) {
	self.buffer = append(self.buffer, pystruct.Pack("!h", value)...)
}
func (self *Buffer) WriteUShortBE(value uint) {
	self.buffer = append(self.buffer, pystruct.Pack("!H", value)...)
}

func (self *Buffer) WriteFloatBE(value float32) {
	self.buffer = append(self.buffer, pystruct.Pack("!f", value)...)
}
func (self *Buffer) WriteDoubleBE(value float64) {
	self.buffer = append(self.buffer, pystruct.Pack("!d", value)...)
}