package rbc

import "fmt"

func Compile(filename string) {
	var cf *compiled_file

	p_cf, _ := readRbc(filename)
	cf, _ = p_cf.(*compiled_file)

	fmt.Println(cf)

}
