package rbc

import (
	"fmt"
)

type File interface {
	Version() int
	Method() Method
}

type Method interface {
	Name() string
	FileName() string
}

type String interface {
	Bytes() []byte
	Encoding() string
	String() string
}

func (self *compiled_file) Version() int {
	return self.version
}

func (self *compiled_file) Method() Method {
	return self.body
}

func (self *compiled_code) FileName() (filename string) {
	if self.file != nil {
		filename = self.file.String()
	}
	return
}

func (self *compiled_code) Name() string {
	return self.name.String()
}

func (self *compiled_symbol) String() string {
	return fmt.Sprintf("%s", self.bytes)
}

func (self *compiled_string) String() string {
	return fmt.Sprintf("%s", self.bytes)
}
