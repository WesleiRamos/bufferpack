package bufferpack

import "fmt"
import "testing"

var tbool bool = true
var tchar int32 = 'c'
var tbyte byte = 1
var tshort int = 2
var tushort uint = 2
var tint int = 4
var tuint uint = 4
var tlong int32 = 4
var tulong uint32 = 4
var tllong int64 = 8
var tullong uint64 = 8
var tfloat float32 = 4
var tdouble float64 = 8

var bufferLE *Buffer
var bufferBE *Buffer

func TestWrite(t *testing.T) {
	print("Writing values on Little Endian buffer... ")
	bufferLE = NewBuffer([]byte{})

	bufferLE.WriteBool(tbool)
	bufferLE.WriteChar(tchar)
	bufferLE.WriteByte(tbyte)

	bufferLE.WriteShortLE(tshort)
	bufferLE.WriteUShortLE(tushort)

	bufferLE.WriteIntLE(tint)
	bufferLE.WriteUIntLE(tuint)

	bufferLE.WriteInt32LE(tlong)
	bufferLE.WriteUInt32LE(tulong)

	bufferLE.WriteInt64LE(tllong)
	bufferLE.WriteUInt64LE(tullong)

	bufferLE.WriteFloatLE(tfloat)
	bufferLE.WriteDoubleLE(tdouble)
	println("OK")

	print("Writing values on Big Endian buffer... ")
	bufferBE = NewBuffer([]byte{})

	bufferBE.WriteBool(tbool)
	bufferBE.WriteChar(tchar)
	bufferBE.WriteByte(tbyte)

	bufferBE.WriteShortBE(tshort)
	bufferBE.WriteUShortBE(tushort)

	bufferBE.WriteIntBE(tint)
	bufferBE.WriteUIntBE(tuint)

	bufferBE.WriteInt32BE(tlong)
	bufferBE.WriteUInt32BE(tulong)

	bufferBE.WriteInt64BE(tllong)
	bufferBE.WriteUInt64BE(tullong)

	bufferBE.WriteFloatBE(tfloat)
	bufferBE.WriteDoubleBE(tdouble)
	println("OK\n\n")
}

func TestReadLE(t *testing.T) {
	println("Reading values from Little Endian buffer\n")

	Bool := bufferLE.ReadBool()
	Char := bufferLE.ReadChar()
	Byte := bufferLE.ReadByte()
	Short := bufferLE.ReadShortLE()
	UShort := bufferLE.ReadUShortLE()
	Int := bufferLE.ReadIntLE()
	UInt := bufferLE.ReadUIntLE()
	Long := bufferLE.ReadInt32LE()
	ULong := bufferLE.ReadUInt32LE()
	LLong := bufferLE.ReadInt64LE()
	ULLong := bufferLE.ReadUInt64LE()
	Float := bufferLE.ReadFloatLE()
	Double := bufferLE.ReadDoubleLE()

	fmt.Printf("%v == %v : %v\n", Bool, tbool, Bool == tbool)
	fmt.Printf("%v == %v : %v\n", string(Char), string(tchar), Char == tchar)
	fmt.Printf("%v == %v : %v\n", Byte, tbyte, Byte == tbyte)
	fmt.Printf("%v == %v : %v\n", Short, tshort, Short == tshort)
	fmt.Printf("%v == %v : %v\n", UShort, tushort, UShort == tushort)
	fmt.Printf("%v == %v : %v\n", Int, tint, Int == tint)
	fmt.Printf("%v == %v : %v\n", UInt, tuint, UInt == tuint)
	fmt.Printf("%v == %v : %v\n", Long, tlong, Long == tlong)
	fmt.Printf("%v == %v : %v\n", ULong, tulong, ULong == tulong)
	fmt.Printf("%v == %v : %v\n", LLong, tllong, LLong == tllong)
	fmt.Printf("%v == %v : %v\n", ULLong, tullong, ULLong == tullong)
	fmt.Printf("%.2f == %.2f : %v\n", Float, tfloat, Float == tfloat)
	fmt.Printf("%.2f == %.2f : %v\n", Double, tdouble, Double == tdouble)

	println("\nOK\n\n")
}

func TestReadBE(t *testing.T) {
	println("Reading values from Big Endian buffer\n")

	Bool := bufferBE.ReadBool()
	Char := bufferBE.ReadChar()
	Byte := bufferBE.ReadByte()
	Short := bufferBE.ReadShortBE()
	UShort := bufferBE.ReadUShortBE()
	Int := bufferBE.ReadIntBE()
	UInt := bufferBE.ReadUIntBE()
	Long := bufferBE.ReadInt32BE()
	ULong := bufferBE.ReadUInt32BE()
	LLong := bufferBE.ReadInt64BE()
	ULLong := bufferBE.ReadUInt64BE()
	Float := bufferBE.ReadFloatBE()
	Double := bufferBE.ReadDoubleBE()

	fmt.Printf("%v == %v : %v\n", Bool, tbool, Bool == tbool)
	fmt.Printf("%v == %v : %v\n", string(Char), string(tchar), Char == tchar)
	fmt.Printf("%v == %v : %v\n", Byte, tbyte, Byte == tbyte)
	fmt.Printf("%v == %v : %v\n", Short, tshort, Short == tshort)
	fmt.Printf("%v == %v : %v\n", UShort, tushort, UShort == tushort)
	fmt.Printf("%v == %v : %v\n", Int, tint, Int == tint)
	fmt.Printf("%v == %v : %v\n", UInt, tuint, UInt == tuint)
	fmt.Printf("%v == %v : %v\n", Long, tlong, Long == tlong)
	fmt.Printf("%v == %v : %v\n", ULong, tulong, ULong == tulong)
	fmt.Printf("%v == %v : %v\n", LLong, tllong, LLong == tllong)
	fmt.Printf("%v == %v : %v\n", ULLong, tullong, ULLong == tullong)
	fmt.Printf("%.2f == %.2f : %v\n", Float, tfloat, Float == tfloat)
	fmt.Printf("%.2f == %.2f : %v\n", Double, tdouble, Double == tdouble)

	println("\nOK\n\n")
}