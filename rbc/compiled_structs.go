package rbc

type compiled_nil struct{}
type compiled_true struct{}
type compiled_false struct{}

type compiled_int struct {
	value int
}

type compiled_rational struct{}

type compiled_complex struct{}

type compiled_string struct {
	bytes    []byte
	encoding string
}

type compiled_symbol struct {
	bytes    []byte
	encoding string
}

type compiled_encoding struct {
	name string
}

type compiled_tuple struct {
	items []compiled
}

type compiled_float struct{}

type compiled_iseq struct {
	opcodes []int
}

type compiled_constant struct{}

type compiled_file struct {
	signature uint64
	version   int
	body      compiled
}

type compiled_method struct {
	version      int
	metadata     compiled
	primitive    compiled
	name         compiled
	iseq         compiled
	stackSize    compiled
	localCount   compiled
	requiredArgs compiled
	postArgs     compiled
	totalArgs    compiled
	splat        compiled
	keywords     compiled
	arity        compiled
	literals     compiled
	lines        compiled
	file         compiled
	localNames   compiled
}
