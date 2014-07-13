package runtime

import (
	"fmt"
	"github.com/vic/gooby/rbc"
	"github.com/vic/gooby/vm"
)

func InterpretRbc(filename string) {
	cf, _ := rbc.ReadFile(filename)
	vm := vm.NewVM()
	fmt.Print("Load compiled file into vm ", cf, vm)
}
