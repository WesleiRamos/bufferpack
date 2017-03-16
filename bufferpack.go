package bufferpack

import "fmt"

type Buffer struct {
	buffer   []byte
	position int
}

func (self Buffer) Length() int {
	return len(self.buffer)
}

func (self Buffer) Bytes() []byte {
	return self.buffer
}

func (self *Buffer) ToString() string {
	return fmt.Sprintf("%q", self.buffer)
}

func (self Buffer) GetPosition() int {
	return self.position
}

func (self *Buffer) SetPosition(pos int) {
	self.position = pos
}

func (self *Buffer) AddPosition(add int) {
	self.position += add
}

func NewBuffer(buff []byte) *Buffer {
	return &Buffer{buffer: buff}
}