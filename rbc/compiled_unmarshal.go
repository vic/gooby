package rbc

func (s *CompiledNil) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *CompiledTrue) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *CompiledFalse) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *CompiledInt) unmarshal(u *unmarshaler) (err error) {
	s.number, err = u.readInt()
	return
}

func (s *CompiledRational) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *CompiledComplex) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *CompiledString) unmarshal(u *unmarshaler) (err error) {
	p_encoding, err := u.unmarshal()
	s.encoding, _ = p_encoding.(*CompiledEncoding)
	count, err := u.readInt()
	s.value = make([]byte, count)
	read_len, err := u.reader.Read(s.value)
	if read_len != count {
		failExit("Expected to find", count, "bytes but only got", read_len)
	}
	return
}

func (s *CompiledSymbol) unmarshal(u *unmarshaler) (err error) {
	p_encoding, err := u.unmarshal()
	s.encoding, _ = p_encoding.(*CompiledEncoding)
	s.name, err = u.readString()
	return
}

func (s *CompiledTuple) unmarshal(u *unmarshaler) (err error) {
	count, err := u.readInt()
	s.items = make([]compiled, count)
	var comp compiled
	for i := 0; err == nil && i < count; i++ {
		comp, err = u.unmarshal()
		s.items[i] = comp
	}
	return
}

func (s *CompiledFloat) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *CompiledISeq) unmarshal(u *unmarshaler) (err error) {
	count, err := u.readInt()
	s.opcodes = make([]int, count)
	for i := 0; err == nil && i < count; i++ {
		s.opcodes[i], err = u.readInt()
	}
	return
}

func (s *CompiledConstant) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *CompiledEncoding) unmarshal(u *unmarshaler) (err error) {
	s.name, err = u.readString()
	return
}

func (cf *CompiledFile) unmarshal(reader *unmarshaler) (err error) {
	reader.expectLine(MAGIC)

	cf.signature, err = reader.readUint64()
	cf.version, err = reader.readInt()
	body, err := reader.unmarshal()

	cf.body, _ = body.(*CompiledCode)

	return
}

func (self *CompiledCode) unmarshal(reader *unmarshaler) (err error) {
	_, err = reader.readInt() // version ignored

	self.metadata, err = reader.unmarshal()
	self.primitive, err = reader.unmarshal()
	self.name, err = reader.unmarshal()
	self.iseq, err = reader.unmarshal()
	self.stack_size, err = reader.unmarshal()
	self.local_count, err = reader.unmarshal()
	self.required_args, err = reader.unmarshal()
	self.post_args, err = reader.unmarshal()
	self.total_args, err = reader.unmarshal()
	self.splat, err = reader.unmarshal()
	self.keywords, err = reader.unmarshal()
	self.arity, err = reader.unmarshal()
	self.literals, err = reader.unmarshal()
	self.lines, err = reader.unmarshal()
	self.file, err = reader.unmarshal()
	self.local_names, err = reader.unmarshal()

	return
}
