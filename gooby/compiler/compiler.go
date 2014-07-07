package compiler

import (
	"../rbc"
	"fmt"
)

func CompileRbc(filename string) {
	cf, _ := rbc.ReadCompiledFile(filename)
	fmt.Println(cf)
}
