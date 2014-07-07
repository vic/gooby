package runtime

import (
	"../rbc"
	"../vm"
	"fmt"
)

func InterpretRbc(filename string) {
	cf, _ := rbc.ReadCompiledFile(filename)
	vm := vm.NewVM()
	fmt.Print("Load compiled file into vm ", cf, vm)
}
