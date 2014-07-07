package rbc

type compiled_nil struct{}
type compiled_true struct{}
type compiled_false struct{}

type compiled_int struct {
	number int
}

type compiled_rational struct{}

type compiled_complex struct{}

type compiled_string struct {
	value    []byte
	encoding *compiled_encoding
}

type compiled_symbol struct {
	name     string
	encoding *compiled_encoding
}

type compiled_tuple struct {
	items []compiled
}

type compiled_float struct{}

type compiled_iseq struct {
	iseq []int
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
	metadata      compiled
	primitive     compiled
	name          compiled
	iseq          compiled
	stack_size    compiled
	local_count   compiled
	required_args compiled
	post_args     compiled
	total_args    compiled
	splat         compiled
	keywords      compiled
	arity         compiled
	literals      compiled
	lines         compiled
	file          compiled
	local_names   compiled
}
