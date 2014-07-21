package rbc

import "errors"

func (s *compiled_nil) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *compiled_true) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *compiled_false) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *compiled_int) unmarshal(u *unmarshaler) (err error) {
	s.value, err = u.readInt()
	return
}

func (s *compiled_rational) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *compiled_complex) unmarshal(u *unmarshaler) (err error) {
	return
}

func unmarshal_encoded(u *unmarshaler) (bytes []byte, encoding string, err error) {
	var compiled compiled
	var count int
	if compiled, err = u.unmarshal(); err != nil {
		return
	}
	if comp_encoding, ok := compiled.(*compiled_encoding); ok {
		encoding = comp_encoding.name
	} else {
		return
	}
	if count, err = u.readInt(); err == nil {
		bytes = make([]byte, count)
	} else {
		return
	}
	read_len, err := u.reader.Read(bytes)
	if read_len != count {
		err = errors.New("Could not read all bytes")
	}
	return
}

func (s *compiled_string) unmarshal(u *unmarshaler) (err error) {
	s.bytes, s.encoding, err = unmarshal_encoded(u)
	return
}

func (s *compiled_symbol) unmarshal(u *unmarshaler) (err error) {
	s.bytes, s.encoding, err = unmarshal_encoded(u)
	return
}

func (s *compiled_tuple) unmarshal(u *unmarshaler) (err error) {
	count, err := u.readInt()
	s.items = make([]compiled, count)
	var comp compiled
	for i := 0; err == nil && i < count; i++ {
		comp, err = u.unmarshal()
		s.items[i] = comp
	}
	return
}

func (s *compiled_float) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *compiled_iseq) unmarshal(u *unmarshaler) (err error) {
	count, err := u.readInt()
	s.opcodes = make([]int, count)
	for i := 0; err == nil && i < count; i++ {
		s.opcodes[i], err = u.readInt()
	}
	return
}

func (s *compiled_constant) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *compiled_encoding) unmarshal(u *unmarshaler) (err error) {
	s.name, err = u.readString()
	return
}

func (cf *compiled_file) unmarshal(reader *unmarshaler) (err error) {
	reader.expectLine(MAGIC)
	if cf.signature, err = reader.readUint64(); err != nil {
		return
	}
	if cf.version, err = reader.readInt(); err != nil {
		return
	}
	if cf.body, err = reader.unmarshal(); err != nil {
		return
	}
	return
}

func (self *compiled_method) unmarshal(reader *unmarshaler) (err error) {
	if self.version, err = reader.readInt(); err != nil {
		return
	}
	if self.metadata, err = reader.unmarshal(); err != nil {
		return
	}
	if self.primitive, err = reader.unmarshal(); err != nil {
		return
	}
	if self.name, err = reader.unmarshal(); err != nil {
		return
	}
	if self.iseq, err = reader.unmarshal(); err != nil {
		return
	}
	if self.stackSize, err = reader.unmarshal(); err != nil {
		return
	}
	if self.localCount, err = reader.unmarshal(); err != nil {
		return
	}
	if self.requiredArgs, err = reader.unmarshal(); err != nil {
		return
	}
	if self.postArgs, err = reader.unmarshal(); err != nil {
		return
	}
	if self.totalArgs, err = reader.unmarshal(); err != nil {
		return
	}
	if self.splat, err = reader.unmarshal(); err != nil {
		return
	}
	if self.keywords, err = reader.unmarshal(); err != nil {
		return
	}
	if self.arity, err = reader.unmarshal(); err != nil {
		return
	}
	if self.literals, err = reader.unmarshal(); err != nil {
		return
	}
	if self.lines, err = reader.unmarshal(); err != nil {
		return
	}
	if self.file, err = reader.unmarshal(); err != nil {
		return
	}
	if self.localNames, err = reader.unmarshal(); err != nil {
		return
	}
	return
}
