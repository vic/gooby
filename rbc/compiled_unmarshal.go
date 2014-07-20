package rbc

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

func (s *compiled_string) unmarshal(u *unmarshaler) (err error) {
	compiled, err := u.unmarshal()
	s.encoding, _ = compiled.(*compiled_encoding)
	count, err := u.readInt()
	s.bytes = make([]byte, count)
	read_len, err := u.reader.Read(s.bytes)
	if read_len != count {
		panicln("Expected to find", count, "bytes but only got", read_len)
	}
	return
}

func (s *compiled_symbol) unmarshal(u *unmarshaler) (err error) {
	compiled, err := u.unmarshal()
	s.encoding, _ = compiled.(*compiled_encoding)
	count, err := u.readInt()
	s.bytes = make([]byte, count)
	read_len, err := u.reader.Read(s.bytes)
	if read_len != count {
		panicln("Expected to find", count, "bytes but only got", read_len)
	}
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

	cf.signature, err = reader.readUint64()
	cf.version, err = reader.readInt()
	body, err := reader.unmarshal()

	cf.body, _ = body.(*compiled_code)

	return
}

func (self *compiled_code) unmarshal(reader *unmarshaler) (err error) {
	var comp compiled
	if self.version, err = reader.readInt(); err != nil {
		return
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.metadata = comp
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.primitive, _ = comp.(*compiled_symbol)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.name, _ = comp.(*compiled_symbol)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.iseq, _ = comp.(*compiled_iseq)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.stackSize, _ = comp.(*compiled_int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.localCount, _ = comp.(*compiled_int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.requiredArgs, _ = comp.(*compiled_int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.postArgs, _ = comp.(*compiled_int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.totalArgs, _ = comp.(*compiled_int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.splat, _ = comp.(*compiled_int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.keywords, _ = comp.(*compiled_tuple)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.arity, _ = comp.(*compiled_int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.literals, _ = comp.(*compiled_tuple)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.lines, _ = comp.(*compiled_tuple)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.file, _ = comp.(*compiled_symbol)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.localNames, _ = comp.(*compiled_tuple)
	}

	return
}
