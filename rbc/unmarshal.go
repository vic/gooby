package rbc

import (
	"bufio"
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

func ReadFile(filename string) (cf *File, err error) {
	compiled, err := readRbc(filename)
	cf, _ = compiled.(*File)
	return
}

func readRbc(filename string) (cf compiled, err error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		panicln("Could not open", filename)
	}

	reader := &unmarshaler{
		filename: filename,
		reader:   bufio.NewReader(file),
	}

	cf = &File{}
	err = cf.unmarshal(reader)

	return
}

func panicln(args ...interface{}) {
	panic(fmt.Sprintln(args...))
}

func (self *unmarshaler) unmarshal() (val compiled, err error) {
	code, err := self.readLine()
	switch code {
	case "":
		return
	case "n":
		val = &Nil{}
	case "t":
		val = &True{}
	case "f":
		val = &False{}
	case "I":
		val = &Int{}
	case "R":
		val = &Rational{}
	case "C":
		val = &Complex{}
	case "s":
		val = &String{}
	case "x":
		val = &Symbol{}
	case "p":
		val = &Tuple{}
	case "d":
		val = &Float{}
	case "i":
		val = &ISeq{}
	case "M":
		val = &Code{}
	case "c":
		val = &Constant{}
	case "E":
		val = &Encoding{}
	default:
		panicln("unknown marshal code: ", code)
	}
	if err == nil {
		err = val.unmarshal(self)
	}
	return
}

func (self *unmarshaler) expectLine(expected string) {
	line, err := self.readLine()
	if err != nil || line != expected {
		panicln("Expected", expected, "in", self.filename)
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

func (self *unmarshaler) readInt() (val int, err error) {
	line, err := self.readLine()
	if err != nil {
		return
	}
	val, err = strconv.Atoi(line)
	return
}

func (self *unmarshaler) readString() (val string, err error) {
	count, err := self.readInt()
	var bytes = make([]byte, count)
	read_len, err := self.reader.Read(bytes)
	if read_len != count {
		panicln("Expected to find", count, "bytes but only got", read_len)
	}
	val = string(bytes)
	self.expectLine("")
	return
}
