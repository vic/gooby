package rbc

type Nil struct{}
type True struct{}
type False struct{}

type Int struct {
	Value int
}

type Rational struct{}

type Complex struct{}

type String struct {
	Bytes    []byte
	Encoding *Encoding
}

type Symbol struct {
	Bytes    []byte
	Encoding *Encoding
}

type Tuple struct {
	Items []compiled
}

type Float struct{}

type ISeq struct {
	Opcodes []int
}

type Constant struct{}

type Encoding struct {
	Name string
}

type File struct {
	Signature uint64
	Version   int
	Body      *Code
}

type Code struct {
	Version      int
	Metadata     compiled
	Primitive    *Symbol
	Name         *Symbol
	ISeq         *ISeq
	StackSize    *Int
	LocalCount   *Int
	RequiredArgs *Int
	PostArgs     *Int
	TotalArgs    *Int
	Splat        *Int
	Keywords     *Tuple
	Arity        *Int
	Literals     *Tuple
	Lines        *Tuple
	File         *Symbol
	LocalNames   *Tuple
}
