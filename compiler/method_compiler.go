package compiler

import (
	"github.com/vic/gooby/rbc"
	"go/ast"
	"go/token"
	"strconv"
)

type method_compiler struct {
	rbc.Method
	stack_top int
	body      *[]ast.Stmt
	iseq      *[]int
}

func new_method_compiler(method rbc.Method) (c *method_compiler) {
	iseq := method.ISeq()[:]
	c = &method_compiler{method, -1, &[]ast.Stmt{}, &iseq}
	return
}

func (self *method_compiler) name() (name string) {
	name = "gooby_" + self.Method.Name()
	return
}

func (self *method_compiler) compile() (f *ast.FuncDecl) {
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

	f = &ast.FuncDecl{
		Name: ast.NewIdent(self.name()),
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

func (self *method_compiler) local_var_decls() (decl ast.Stmt) {
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

func (self *method_compiler) shift_iseq() (val int) {
	val = (*self.iseq)[0]
	*self.iseq = (*self.iseq)[1:]
	return
}

func (self *method_compiler) compile_instructions() {
	for len(*self.iseq) > 0 {
		opcode := self.shift_iseq()
		compiler := opcode_compilers[opcode]
		compiler(self)
	}
}

func (self *method_compiler) rt_(name string, args ...ast.Expr) (expr ast.Expr) {
	expr = &ast.CallExpr{
		Fun:  ast.NewIdent("rt." + name),
		Args: args,
	}
	return
}

func (self *method_compiler) push(expr ast.Expr) {
	self.stack_top++
	self.set_top(expr)
}

func (self *method_compiler) rb_n(n int) ast.Expr {
	return ast.NewIdent("rb" + strconv.Itoa(n))
}

func (self *method_compiler) rb_many(n int) (ary []ast.Expr) {
	ary = make([]ast.Expr, n)
	for i := 0; i < n; i++ {
		ary[i] = self.rb_top()
	}
	return
}

func (self *method_compiler) rb_top() ast.Expr {
	return self.rb_n(self.stack_top)
}

func (self *method_compiler) assign(left ast.Expr, right ast.Expr) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{left},
		Rhs: []ast.Expr{right},
		Tok: token.ASSIGN,
	}
	self.append_stmt(stmt)

}

func (self *method_compiler) set_top(expr ast.Expr) {
	self.assign(self.rb_top(), expr)
}

func (self *method_compiler) literal(i int) (expr ast.Expr) {
	args := []ast.Expr{}
	literal := self.Method.Literal(i)
	method := ""

	var s rbc.String
	switch o := literal.(type) {
	case rbc.String:
		method = "String"
		s = o
	case rbc.Symbol:
		method = "Symbol"
		s = o
	}

	if method == "String" || method == "Symbol" {
		args = append(args, &ast.BasicLit{
			Value: `"` + s.HexBytes() + `"`,
			Kind:  token.STRING,
		}, &ast.BasicLit{
			Value: `"` + s.Encoding() + `"`,
			Kind:  token.STRING,
		})
	}

	expr = &ast.CallExpr{
		Fun:  ast.NewIdent("rt." + method + "Literal"),
		Args: args,
	}
	return
}

func (self *method_compiler) append_stmt(stmt ast.Stmt) {
	*self.body = append(*self.body, stmt)
}

func (self *method_compiler) append_expr(expr ast.Expr) {
	self.append_stmt(&ast.ExprStmt{X: expr})
}
