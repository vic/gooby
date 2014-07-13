package rbc

func (s *Nil) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *True) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *False) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *Int) unmarshal(u *unmarshaler) (err error) {
	s.Value, err = u.readInt()
	return
}

func (s *Rational) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *Complex) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *String) unmarshal(u *unmarshaler) (err error) {
	compiled, err := u.unmarshal()
	s.Encoding, _ = compiled.(*Encoding)
	count, err := u.readInt()
	s.Bytes = make([]byte, count)
	read_len, err := u.reader.Read(s.Bytes)
	if read_len != count {
		panicln("Expected to find", count, "bytes but only got", read_len)
	}
	return
}

func (s *Symbol) unmarshal(u *unmarshaler) (err error) {
	compiled, err := u.unmarshal()
	s.Encoding, _ = compiled.(*Encoding)
	count, err := u.readInt()
	s.Bytes = make([]byte, count)
	read_len, err := u.reader.Read(s.Bytes)
	if read_len != count {
		panicln("Expected to find", count, "bytes but only got", read_len)
	}
	return
}

func (s *Tuple) unmarshal(u *unmarshaler) (err error) {
	count, err := u.readInt()
	s.Items = make([]compiled, count)
	var comp compiled
	for i := 0; err == nil && i < count; i++ {
		comp, err = u.unmarshal()
		s.Items[i] = comp
	}
	return
}

func (s *Float) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *ISeq) unmarshal(u *unmarshaler) (err error) {
	count, err := u.readInt()
	s.Opcodes = make([]int, count)
	for i := 0; err == nil && i < count; i++ {
		s.Opcodes[i], err = u.readInt()
	}
	return
}

func (s *Constant) unmarshal(u *unmarshaler) (err error) {
	return
}

func (s *Encoding) unmarshal(u *unmarshaler) (err error) {
	s.Name, err = u.readString()
	return
}

func (cf *File) unmarshal(reader *unmarshaler) (err error) {
	reader.expectLine(MAGIC)

	cf.Signature, err = reader.readUint64()
	cf.Version, err = reader.readInt()
	body, err := reader.unmarshal()

	cf.Body, _ = body.(*Code)

	return
}

func (self *Code) unmarshal(reader *unmarshaler) (err error) {
	var comp compiled
	if self.Version, err = reader.readInt(); err != nil {
		return
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.Metadata = comp
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.Primitive, _ = comp.(*Symbol)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.Name, _ = comp.(*Symbol)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.ISeq, _ = comp.(*ISeq)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.StackSize, _ = comp.(*Int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.LocalCount, _ = comp.(*Int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.RequiredArgs, _ = comp.(*Int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.PostArgs, _ = comp.(*Int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.TotalArgs, _ = comp.(*Int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.Splat, _ = comp.(*Int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.Keywords, _ = comp.(*Tuple)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.Arity, _ = comp.(*Int)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.Literals, _ = comp.(*Tuple)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.Lines, _ = comp.(*Tuple)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.File, _ = comp.(*Symbol)
	}
	if comp, err = reader.unmarshal(); err == nil {
		self.LocalNames, _ = comp.(*Tuple)
	}

	return
}
