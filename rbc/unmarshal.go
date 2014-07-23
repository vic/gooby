package rbc

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const MAGIC = "!RBIX"

type unmarshaler struct {
	reader   *bufio.Reader
	filename string
}

type compiled interface {
	unmarshal(reader *unmarshaler) (err error)
}

func ReadFile(filename string) (cf File, err error) {
	if compiled, err := readRbc(filename); err == nil {
		cf = compiled.(File)
	}
	return
}

func readRbc(filename string) (cf compiled, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := &unmarshaler{
		filename: filename,
		reader:   bufio.NewReader(file),
	}

	cf = &compiled_file{}
	err = cf.unmarshal(reader)

	return
}

func (self *unmarshaler) unmarshal() (val compiled, err error) {
	code, err := self.readLine()
	if err != nil {
		return
	}
	switch code {
	case "n":
		val = &compiled_nil{}
	case "t":
		val = &compiled_true{}
	case "f":
		val = &compiled_false{}
	case "I":
		val = &compiled_int{}
	case "R":
		val = &compiled_rational{}
	case "C":
		val = &compiled_complex{}
	case "s":
		val = &compiled_string{}
	case "x":
		val = &compiled_symbol{}
	case "p":
		val = &compiled_tuple{}
	case "d":
		val = &compiled_float{}
	case "i":
		val = &compiled_iseq{}
	case "M":
		val = &compiled_method{}
	case "c":
		val = &compiled_constant{}
	case "E":
		val = &compiled_encoding{}
	default:
		err = errors.New("unknown marshal code: " + code)
	}
	if err == nil {
		err = val.unmarshal(self)
	}
	return
}

func (self *unmarshaler) expectLine(expected string) {
	line, err := self.readLine()
	if err != nil || line != expected {
		panic(fmt.Sprintf("%s: expected '%s' but got '%s'", self.filename, expected, line))
	}
}

func (self *unmarshaler) readLine() (line string, err error) {
	line, err = self.reader.ReadString('\n')
	if err == nil {
		line = line[:len(line)-1]
	}
	return
}

func (self *unmarshaler) readUint64() (val uint64, err error) {
	line, err := self.readLine()
	if err != nil {
		return
	}
	val, err = strconv.ParseUint(line, 10, 64)
	return
}

func (self *unmarshaler) readInt(base int) (val int, err error) {
	line, err := self.readLine()
	if err != nil {
		return
	}
	if i, err := strconv.ParseInt(line, base, 0); err == nil {
		val = int(i)
	}
	return
}

func (self *unmarshaler) readString() (val string, err error) {
	count, err := self.readInt(10)
	var bytes = make([]byte, count)
	read_len, err := self.reader.Read(bytes)
	if read_len != count {
		err = errors.New("Could not read all bytes")
		return
	}
	val = string(bytes)
	self.expectLine("")
	return
}
