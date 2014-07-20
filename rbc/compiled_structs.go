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
	encoding *compiled_encoding
}

type compiled_symbol struct {
	bytes    []byte
	encoding *compiled_encoding
}

type compiled_tuple struct {
	items []compiled
}

type compiled_float struct{}

type compiled_iseq struct {
	opcodes []int
}

type compiled_constant struct{}

type compiled_encoding struct {
	name string
}

type compiled_file struct {
	signature uint64
	version   int
	body      *compiled_code
}

type compiled_code struct {
	version      int
	metadata     compiled
	primitive    *compiled_symbol
	name         *compiled_symbol
	iseq         *compiled_iseq
	stackSize    *compiled_int
	localCount   *compiled_int
	requiredArgs *compiled_int
	postArgs     *compiled_int
	totalArgs    *compiled_int
	splat        *compiled_int
	keywords     *compiled_tuple
	arity        *compiled_int
	literals     *compiled_tuple
	lines        *compiled_tuple
	file         *compiled_symbol
	localNames   *compiled_tuple
}
