package compiler

import (
	"fmt"
	"github.com/vic/gooby/rbc"
)

func CompileRbc(filename string) {
	cf, _ := rbc.ReadFile(filename)
	fmt.Println(cf)
}
