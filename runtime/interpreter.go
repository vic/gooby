package runtime

import (
	"fmt"
	"github.com/vic/gooby/rbc"
)

func InterpretRbc(filename string) {
	cf, _ := rbc.ReadFile(filename)
	fmt.Print("Load compiled file into vm ", cf)
}
