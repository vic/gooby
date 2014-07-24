package compiler

import (
	_ "fmt"
	"github.com/vic/gooby/rbc"
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func CompileRbc(filename string) (err error) {
	cf, err := rbc.ReadFile(filename)
	if err != nil {
		return
	}

	compiler := &file_compiler{cf}
	file := compiler.compile(filename)

	fset := token.NewFileSet()
	format.Node(os.Stdout, fset, file)
	return
}

type file_compiler struct{ rbc.File }

func (self *file_compiler) compile(filename string) (f *ast.File) {
	gooby_import := &ast.ImportSpec{
		Name: ast.NewIdent("gooby"),
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"github.com/vic/gooby/runtime\"",
		},
	}

	body_compiler := &method_compiler{self.Method(), 0, &[]ast.Stmt{}}
	script_decl := body_compiler.compile()

	local_vm_decl := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{
			Name: ast.NewIdent("runtime"),
			Type: &ast.StructType{
				Fields: &ast.FieldList{
					List: []*ast.Field{
						&ast.Field{Type: ast.NewIdent("gooby.Runtime")},
					},
				},
			},
		}},
	}

	f = &ast.File{
		Name: ast.NewIdent(filename),
		Decls: []ast.Decl{
			&ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{gooby_import}},
			local_vm_decl,
			script_decl,
		},
	}

	return
}
