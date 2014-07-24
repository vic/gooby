package compiler

import (
	"github.com/vic/gooby/rbc"
	"go/ast"
	"go/token"
	"strconv"
)

type method_compiler struct{ rbc.Method }

type stack_usage struct {
	max  int
	used int
}

func (self method_compiler) compile() (f *ast.FuncDecl) {
	su := &stack_usage{max: self.StackSize()}

	params := []*ast.Field{}
	returns := []*ast.Field{}

	returns = append(returns,
		&ast.Field{
			Names: []*ast.Ident{ast.NewIdent("res")},
			Type:  ast.NewIdent("gooby.Object"),
		},
		&ast.Field{
			Names: []*ast.Ident{ast.NewIdent("err")},
			Type:  ast.NewIdent("gooby.Error"),
		},
	)

	body := []ast.Stmt{self.local_var_decls(su)}
	self.append_instructions(su, &body)
	body = append(body, &ast.ReturnStmt{})

	f = &ast.FuncDecl{
		Name: ast.NewIdent(self.Name()),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				&ast.Field{
					Names: []*ast.Ident{ast.NewIdent("rt")},
					Type:  ast.NewIdent("runtime"),
				},
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: params,
			},
			Results: &ast.FieldList{
				List: returns,
			},
		},
		Body: &ast.BlockStmt{
			List: body,
		},
	}
	return f
}

func (self method_compiler) local_var_decls(su *stack_usage) (decl ast.Stmt) {
	if su.max < 1 {
		return &ast.EmptyStmt{}
	}
	names := make([]*ast.Ident, su.max)
	for i := 0; i < su.max; i++ {
		names[i] = ast.NewIdent("rb" + strconv.Itoa(i))
	}
	names_spec := ast.ValueSpec{
		Names: names,
		Type:  ast.NewIdent("gooby.Object"),
	}
	vars := ast.GenDecl{
		Tok:   token.VAR,
		Specs: []ast.Spec{&names_spec},
	}
	decl = &ast.DeclStmt{
		Decl: &vars,
	}
	return
}

func (self method_compiler) append_instructions(su *stack_usage, body *[]ast.Stmt) {
	iseq := self.ISeq()
	for _, opcode := range iseq {
		compiler := opcode_compilers[opcode]
		compiler(self, su, body)
	}
}
