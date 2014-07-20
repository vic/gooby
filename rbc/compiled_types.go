package rbc

import (
	"fmt"
)

type File interface {
	Version() int
	Body() Method
}

type Method interface {
	String() string
}

func (self *compiled_file) Version() int {
	return self.version
}

func (self *compiled_file) Body() Method {
	return self.body
}

func (self *compiled_code) String() string {
	return fmt.Sprintf("%s", self.name.bytes)
}
