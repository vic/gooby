package compiler

import (
	"go/ast"
	"go/token"
	"strconv"
)

type opcode_compiler func(self *method_compiler)

func (self *method_compiler) noop() {
}

func (self *method_compiler) push_nil() {
	self.push(self.rt_("Nil"))
}

func (self *method_compiler) push_true() {
	self.push(self.rt_("True"))
}

func (self *method_compiler) push_false() {
	self.push(self.rt_("False"))
}

func (self *method_compiler) push_int() {
}

func (self *method_compiler) push_self() {
	self.push(self.rt_("Self"))
}

func (self *method_compiler) set_literal() {
}

func (self *method_compiler) push_literal() {
	self.push(self.literal(self.shift_iseq()))
}

func (self *method_compiler) goto_() {
}

func (self *method_compiler) goto_if_false() {
}

func (self *method_compiler) goto_if_true() {
}

func (self *method_compiler) ret() {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent("res")},
		Rhs: []ast.Expr{ast.NewIdent("rb" + strconv.Itoa(self.stack_top))},
		Tok: token.ASSIGN,
	}
	self.append_stmt(stmt)
	self.append_stmt(&ast.ReturnStmt{})
}

func (self *method_compiler) swap_stack() {
}

func (self *method_compiler) dup_top() {
}

func (self *method_compiler) dup_many() {
}

func (self *method_compiler) pop() {
	self.set_top(ast.NewIdent("nil"))
	self.stack_top--
}

func (self *method_compiler) pop_many() {
}

func (self *method_compiler) rotate() {
}

func (self *method_compiler) move_down() {
}

func (self *method_compiler) set_local() {
}

func (self *method_compiler) push_local() {
}

func (self *method_compiler) push_local_depth() {
}

func (self *method_compiler) set_local_depth() {
}

func (self *method_compiler) passed_arg() {
}

func (self *method_compiler) push_current_exception() {
}

func (self *method_compiler) clear_exception() {
}

func (self *method_compiler) push_exception_state() {
}

func (self *method_compiler) restore_exception_state() {
}

func (self *method_compiler) raise_exc() {
}

func (self *method_compiler) setup_unwind() {
}

func (self *method_compiler) pop_unwind() {
}

func (self *method_compiler) raise_return() {
}

func (self *method_compiler) ensure_return() {
}

func (self *method_compiler) raise_break() {
}

func (self *method_compiler) reraise() {
}

func (self *method_compiler) make_array() {
}

func (self *method_compiler) cast_array() {
}

func (self *method_compiler) shift_array() {
}

func (self *method_compiler) set_ivar() {
}

func (self *method_compiler) push_ivar() {
}

func (self *method_compiler) push_const() {
}

func (self *method_compiler) set_const() {
}

func (self *method_compiler) set_const_at() {
}

func (self *method_compiler) find_const() {
}

func (self *method_compiler) push_cpath_top() {
}

func (self *method_compiler) push_const_fast() {
}

func (self *method_compiler) find_const_fast() {
}

func (self *method_compiler) set_call_flags() {
}

func (self *method_compiler) allow_private() {
	self.append_expr(self.rt_("AllowPrivate"))
}

func (self *method_compiler) send_method() {
}

func (self *method_compiler) send_stack() {
}

func (self *method_compiler) send_stack_with_block() {
}

func (self *method_compiler) send_stack_with_splat() {
}

func (self *method_compiler) send_super_stack_with_block() {
}

func (self *method_compiler) send_super_stack_with_splat() {
}

func (self *method_compiler) push_block() {
}

func (self *method_compiler) passed_blockarg() {
}

func (self *method_compiler) create_block() {
}

func (self *method_compiler) cast_for_single_block_arg() {
}

func (self *method_compiler) cast_for_multi_block_arg() {
}

func (self *method_compiler) cast_for_splat_block_arg() {
}

func (self *method_compiler) yield_stack() {
}

func (self *method_compiler) yield_splat() {
}

func (self *method_compiler) string_append() {
}

func (self *method_compiler) string_build() {
}

func (self *method_compiler) string_dup() {
	self.append_expr(self.rt_("StringDup"))
}

func (self *method_compiler) push_scope() {
}

func (self *method_compiler) add_scope() {
}

func (self *method_compiler) push_variables() {
}

func (self *method_compiler) check_interrupts() {
}

func (self *method_compiler) yield_debugger() {
}

func (self *method_compiler) is_nil() {
}

func (self *method_compiler) check_serial() {
}

func (self *method_compiler) check_serial_private() {
}

func (self *method_compiler) push_my_field() {
}

func (self *method_compiler) store_my_field() {
}

func (self *method_compiler) kind_of() {
}

func (self *method_compiler) instance_of() {
}

func (self *method_compiler) meta_push_neg_1() {
}

func (self *method_compiler) meta_push_0() {
}

func (self *method_compiler) meta_push_1() {
}

func (self *method_compiler) meta_push_2() {
}

func (self *method_compiler) meta_send_op_plus() {
}

func (self *method_compiler) meta_send_op_minus() {
}

func (self *method_compiler) meta_send_op_equal() {
}

func (self *method_compiler) meta_send_op_lt() {
}

func (self *method_compiler) meta_send_op_gt() {
}

func (self *method_compiler) meta_send_op_tequal() {
}

func (self *method_compiler) meta_send_call() {
}

func (self *method_compiler) push_my_offset() {
}

func (self *method_compiler) zsuper() {
}

func (self *method_compiler) push_block_arg() {
}

func (self *method_compiler) push_undef() {
}

func (self *method_compiler) push_stack_local() {
}

func (self *method_compiler) set_stack_local() {
}

func (self *method_compiler) push_has_block() {
}

func (self *method_compiler) push_proc() {
}

func (self *method_compiler) check_frozen() {
}

func (self *method_compiler) cast_multi_value() {
}

func (self *method_compiler) invoke_primitive() {
}

func (self *method_compiler) push_rubinius() {
}

func (self *method_compiler) call_custom() {
}

func (self *method_compiler) meta_to_s() {
}

func (self *method_compiler) push_type() {
}

func (self *method_compiler) push_mirror() {
}
