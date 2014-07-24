package compiler

import (
	_ "fmt"
	"github.com/vic/gooby/rbc"
	"go/ast"
	"go/token"
	"strconv"
)

type method_compiler struct {
	rbc.Method
	stack_used int
	body       *[]ast.Stmt
}

func (self method_compiler) compile() (f *ast.FuncDecl) {

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

	self.append_stmt(self.local_var_decls())
	self.compile_instructions()
	self.append_stmt(&ast.ReturnStmt{})

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
			List: *self.body,
		},
	}
	return f
}

func (self method_compiler) local_var_decls() (decl ast.Stmt) {
	if self.StackSize() < 1 {
		return &ast.EmptyStmt{}
	}
	names := make([]*ast.Ident, self.StackSize())
	for i := 0; i < self.StackSize(); i++ {
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

func (self method_compiler) compile_instructions() {
	iseq := self.ISeq()
	for _, opcode := range iseq {
		compiler := opcode_compilers[opcode]
		compiler(self)
	}
}

func (self method_compiler) self() (expr ast.Expr) {
	expr = &ast.CallExpr{
		Fun:  ast.NewIdent("rt.self"),
		Args: []ast.Expr{},
	}
	return
}

func (self method_compiler) pop() {
	self.set_top(ast.NewIdent("nil"))
	self.stack_used--
}

func (self method_compiler) push(expr ast.Expr) {
	self.set_top(expr)
	self.stack_used++
}

func (self method_compiler) set_top(expr ast.Expr) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent("rb" + strconv.Itoa(self.stack_used))},
		Rhs: []ast.Expr{expr},
		Tok: token.ASSIGN,
	}
	self.append_stmt(stmt)
}

func (self method_compiler) append_stmt(stmt ast.Stmt) {
	*self.body = append(*self.body, stmt)
}
