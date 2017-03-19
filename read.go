package bufferpack

import "reflect"
import "github.com/WesleiRamos/pystruct"

func (self *Buffer) read(t string, end int) interface{} {
	i, err := pystruct.Unpack(t, self.buffer[self.position:self.position+end])
	if err != nil {
		panic(err)
	}
	self.position += end
	return i[0]
}

func (self *Buffer) ReadByte() byte {
	// byte
	b := self.read("b", 1)
	return byte(reflect.ValueOf(b).Uint())
}

func (self *Buffer) ReadChar() int32 {
	// char
	return int32(self.ReadByte())
}

func (self *Buffer) ReadBool() bool {
	// bool
	b := self.read("?", 1)
	return reflect.ValueOf(b).Bool()
}

// LITTLE ENDIAN

func (self *Buffer) ReadIntLE() int {
	// Int
	i := self.read("i", 4)
	return int(reflect.ValueOf(i).Uint())
}
func (self *Buffer) ReadUIntLE() uint {
	// unsigned int
	ui := self.read("I", 4)
	return uint(reflect.ValueOf(ui).Uint())
}

func (self *Buffer) ReadInt32LE() int32 {
	// long
	i := self.read("i", 4)
	return int32(reflect.ValueOf(i).Uint())
}
func (self *Buffer) ReadUInt32LE() uint32 {
	// unsigned long
	ui := self.read("I", 4)
	return uint32(reflect.ValueOf(ui).Uint())
}

func (self *Buffer) ReadInt64LE() int64 {
	// long long
	ll := self.read("q", 8)
	return reflect.ValueOf(ll).Int()
}
func (self *Buffer) ReadUInt64LE() uint64 {
	// unsigned long long
	ull := self.read("Q", 8)
	return reflect.ValueOf(ull).Uint()
}

func (self *Buffer) ReadFloatLE() float32 {
	// float
	f := self.read("f", 4)
	return float32(reflect.ValueOf(f).Float())
}

func (self *Buffer) ReadDoubleLE() float64 {
	// double
	d := self.read("d", 8)
	return reflect.ValueOf(d).Float()
}

func (self *Buffer) ReadShortLE() int {
	// short
	s := self.read("h", 2)
	return int(reflect.ValueOf(s).Int())
}

func (self *Buffer) ReadUShortLE() uint {
	// unsigned short
	s := self.read("H", 2)
	return uint(reflect.ValueOf(s).Uint())
}

func (self *Buffer) ReadStringLE() string {
	// string
	l := self.ReadShortLE()
	s := self.buffer[self.position : self.position+l]
	self.position += l
	return string(s)
}

// BIG ENDIAN

func (self *Buffer) ReadIntBE() int {
	// Int
	i := self.read("!i", 4)
	return int(reflect.ValueOf(i).Uint())
}
func (self *Buffer) ReadUIntBE() uint {
	// unsigned int
	ui := self.read("!I", 4)
	return uint(reflect.ValueOf(ui).Uint())
}

func (self *Buffer) ReadInt32BE() int32 {
	// long
	i := self.read("!i", 4)
	return int32(reflect.ValueOf(i).Uint())
}
func (self *Buffer) ReadUInt32BE() uint32 {
	// unsigned long
	ui := self.read("!I", 4)
	return uint32(reflect.ValueOf(ui).Uint())
}

func (self *Buffer) ReadInt64BE() int64 {
	// long long
	ll := self.read("!q", 8)
	return reflect.ValueOf(ll).Int()
}
func (self *Buffer) ReadUInt64BE() uint64 {
	// unsigned long long
	ull := self.read("!Q", 8)
	return reflect.ValueOf(ull).Uint()
}

func (self *Buffer) ReadFloatBE() float32 {
	// float
	f := self.read("!f", 4)
	return float32(reflect.ValueOf(f).Float())
}

func (self *Buffer) ReadDoubleBE() float64 {
	// double
	d := self.read("!d", 8)
	return reflect.ValueOf(d).Float()
}

func (self *Buffer) ReadShortBE() int {
	// short
	s := self.read("!h", 2)
	return int(reflect.ValueOf(s).Int())
}

func (self *Buffer) ReadUShortBE() uint {
	// unsigned short
	s := self.read("!H", 2)
	return uint(reflect.ValueOf(s).Uint())
}

func (self *Buffer) ReadStringBE() string {
	// string
	l := self.ReadShortBE()
	s := self.buffer[self.position : self.position+l]
	self.position += l
	return string(s)
}