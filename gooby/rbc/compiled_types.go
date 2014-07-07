package rbc

type CompiledNil struct{}
type CompiledTrue struct{}
type CompiledFalse struct{}

type CompiledInt struct {
	number int
}

type CompiledRational struct{}

type CompiledComplex struct{}

type CompiledString struct {
	value    []byte
	encoding *CompiledEncoding
}

type CompiledSymbol struct {
	name     string
	encoding *CompiledEncoding
}

type CompiledTuple struct {
	items []compiled
}

type CompiledFloat struct{}

type CompiledISeq struct {
	opcodes []int
}

type CompiledConstant struct{}

type CompiledEncoding struct {
	name string
}

type CompiledFile struct {
	signature uint64
	version   int
	body      *CompiledCode
}

type CompiledCode struct {
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
