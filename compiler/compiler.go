package compiler

import (
	"fmt"
	"github.com/vic/gooby/rbc"
)

func CompileRbc(filename string) {
	cf, _ := rbc.ReadCompiledFile(filename)
	fmt.Println(cf)
}
