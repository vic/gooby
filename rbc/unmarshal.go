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

func ReadCompiledFile(filename string) (cf *CompiledFile, err error) {
	compiled, err := readRbc(filename)
	cf, _ = compiled.(*CompiledFile)
	return
}

func readRbc(filename string) (cf compiled, err error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		failExit("Could not open", filename)
	}

	reader := &unmarshaler{
		filename: filename,
		reader:   bufio.NewReader(file),
	}

	cf = &CompiledFile{}
	err = cf.unmarshal(reader)

	return
}

func failExit(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func (self *unmarshaler) unmarshal() (val compiled, err error) {
	code, err := self.readLine()
	switch code {
	case "":
		return
	case "n":
		val = &CompiledNil{}
	case "t":
		val = &CompiledTrue{}
	case "f":
		val = &CompiledFalse{}
	case "I":
		val = &CompiledInt{}
	case "R":
		val = &CompiledRational{}
	case "C":
		val = &CompiledComplex{}
	case "s":
		val = &CompiledString{}
	case "x":
		val = &CompiledSymbol{}
	case "p":
		val = &CompiledTuple{}
	case "d":
		val = &CompiledFloat{}
	case "i":
		val = &CompiledISeq{}
	case "M":
		val = &CompiledCode{}
	case "c":
		val = &CompiledConstant{}
	case "E":
		val = &CompiledEncoding{}
	default:
		failExit("unknown marshal code: ", code)
	}
	if err == nil {
		err = val.unmarshal(self)
	}
	return
}

func (self *unmarshaler) expectLine(expected string) {
	line, err := self.readLine()
	if err != nil || line != expected {
		failExit("Expected", expected, "in", self.filename)
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
		failExit("Expected to find", count, "bytes but only got", read_len)
	}
	val = string(bytes)
	self.expectLine("")
	return
}
