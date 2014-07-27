package runtime

import (
	"encoding/hex"
)

type _string struct {
	bytes    []byte
	encoding string
}

func (s *_string) String() string {
	return string(s.bytes)
}

func (rt *runtime) StringLiteral(hexbytes string, encoding string) Object {
	bytes, _ := hex.DecodeString(hexbytes)
	str := &_string{
		bytes:    bytes,
		encoding: encoding,
	}
	return str
}

func (rt *runtime) StringDup(str Object) Object {
	return str
}
