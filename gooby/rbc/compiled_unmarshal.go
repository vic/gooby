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
	s.number, err = u.readInt()
	return
}

func (s *compiled_rational) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *compiled_complex) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *compiled_string) unmarshal(u *unmarshaler) (err error) {
	p_encoding, err := u.unmarshal()
	s.encoding, _ = p_encoding.(*compiled_encoding)
	count, err := u.readInt()
	s.value = make([]byte, count)
	read_len, err := u.reader.Read(s.value)
	if read_len != count {
		failExit("Expected to find", count, "bytes but only got", read_len)
	}
	return
}

func (s *compiled_symbol) unmarshal(u *unmarshaler) (err error) {
	p_encoding, err := u.unmarshal()
	s.encoding, _ = p_encoding.(*compiled_encoding)
	s.name, err = u.readString()
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
	s.iseq = make([]int, count)
	for i := 0; err == nil && i < count; i++ {
		s.iseq[i], err = u.readInt()
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
