package rbc

import (
	"encoding/hex"
	"fmt"
)

type File interface {
	Version() int
	Method() Method
}

type Method interface {
	Name() string
	FileName() string
	StackSize() int
	ISeq() []int
	Literal(int) compiled
	LiteralCount() int
}

type String interface {
	Bytes() []byte
	Encoding() string
	HexBytes() string
}

func (self *compiled_file) Version() int {
	return self.version
}

func (self *compiled_file) Method() (m Method) {
	m, _ = self.body.(*compiled_method)
	return
}

func (self *compiled_method) FileName() string {
	return _str(self.file)
}

func (self *compiled_method) Name() string {
	return _str(self.name)
}

func (self *compiled_method) StackSize() int {
	return _int(self.stackSize)
}

func (self *compiled_method) ISeq() []int {
	if is, ok := self.iseq.(*compiled_iseq); ok {
		return is.opcodes
	}
	return []int{}
}

func (self *compiled_method) LiteralCount() int {
	if tuple, ok := self.literals.(*compiled_tuple); ok {
		return len(tuple.items)
	}
	return 0
}

func (self *compiled_method) Literal(i int) (lit compiled) {
	if tuple, ok := self.literals.(*compiled_tuple); ok {
		lit = tuple.items[i]
	}
	return
}

func (self *compiled_symbol) Bytes() []byte {
	return self.bytes
}

func (self *compiled_symbol) Encoding() string {
	return self.encoding
}

func (self *compiled_symbol) String() string {
	return fmt.Sprintf("%s", self.bytes)
}

func (self *compiled_string) Bytes() []byte {
	return self.bytes
}

func (self *compiled_string) Encoding() string {
	return self.encoding
}

func (self *compiled_string) HexBytes() string {
	return hex.EncodeToString(self.bytes)
}

func (self *compiled_string) String() string {
	return fmt.Sprintf("%s", self.bytes)
}

func _int(c compiled) (r int) {
	if i, ok := c.(*compiled_int); ok {
		r = i.value
	}
	return
}

func _str(c compiled) (r string) {
	if s, ok := c.(*compiled_symbol); ok {
		r = s.String()
		return
	}
	if s, ok := c.(*compiled_string); ok {
		r = s.String()
		return
	}
	return
}
