package compiler

import (
	"fmt"
	"github.com/vic/gooby/rbc"
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func CompileRbc(filename string) {
	cf, _ := rbc.ReadFile(filename)
	compiler := &file_compiler{cf}
	file := compiler.compile(filename)

	fmt.Println("/* " + filename + " */")
	format.Node(os.Stdout, nil, file)
	fmt.Print("\n")
}

type file_compiler struct{ rbc.File }
type method_compiler struct{ rbc.Method }

func (self *file_compiler) compile(filename string) (f *ast.File) {
	gooby_import := &ast.ImportSpec{
		Name: ast.NewIdent("gooby"),
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"github.com/vic/gooby\"",
		},
	}

	body_compiler := &method_compiler{self.Method()}
	script_decl := body_compiler.compile()

	local_vm_decl := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{
			Name: ast.NewIdent("vm"),
			Type: &ast.StructType{
				Fields: &ast.FieldList{
					List: []*ast.Field{
						&ast.Field{Type: &ast.StarExpr{X: ast.NewIdent("gooby.VM")}},
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

func (self method_compiler) compile() (f *ast.FuncDecl) {
	f = &ast.FuncDecl{
		Name: ast.NewIdent(self.Name()),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{},
			},
		},
	}
	return f
}
