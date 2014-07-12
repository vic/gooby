package vm

type VM interface {
}

type vm struct {
}

func NewVM() VM {
	return &vm{}
}
