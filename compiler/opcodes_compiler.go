package compiler

type opcode_compiler func(self method_compiler)

func compile_noop(self method_compiler) {
}

func compile_push_nil(self method_compiler) {
}

func compile_push_true(self method_compiler) {
}

func compile_push_false(self method_compiler) {
}

func compile_push_int(self method_compiler) {
}

func compile_push_self(self method_compiler) {
	self.push(self.self())
}

func compile_set_literal(self method_compiler) {
}

func compile_push_literal(self method_compiler) {
}

func compile_goto(self method_compiler) {
}

func compile_goto_if_false(self method_compiler) {
}

func compile_goto_if_true(self method_compiler) {
}

func compile_ret(self method_compiler) {
}

func compile_swap_stack(self method_compiler) {
}

func compile_dup_top(self method_compiler) {
}

func compile_dup_many(self method_compiler) {
}

func compile_pop(self method_compiler) {
	self.pop()
}

func compile_pop_many(self method_compiler) {
}

func compile_rotate(self method_compiler) {
}

func compile_move_down(self method_compiler) {
}

func compile_set_local(self method_compiler) {
}

func compile_push_local(self method_compiler) {
}

func compile_push_local_depth(self method_compiler) {
}

func compile_set_local_depth(self method_compiler) {
}

func compile_passed_arg(self method_compiler) {
}

func compile_push_current_exception(self method_compiler) {
}

func compile_clear_exception(self method_compiler) {
}

func compile_push_exception_state(self method_compiler) {
}

func compile_restore_exception_state(self method_compiler) {
}

func compile_raise_exc(self method_compiler) {
}

func compile_setup_unwind(self method_compiler) {
}

func compile_pop_unwind(self method_compiler) {
}

func compile_raise_return(self method_compiler) {
}

func compile_ensure_return(self method_compiler) {
}

func compile_raise_break(self method_compiler) {
}

func compile_reraise(self method_compiler) {
}

func compile_make_array(self method_compiler) {
}

func compile_cast_array(self method_compiler) {
}

func compile_shift_array(self method_compiler) {
}

func compile_set_ivar(self method_compiler) {
}

func compile_push_ivar(self method_compiler) {
}

func compile_push_const(self method_compiler) {
}

func compile_set_const(self method_compiler) {
}

func compile_set_const_at(self method_compiler) {
}

func compile_find_const(self method_compiler) {
}

func compile_push_cpath_top(self method_compiler) {
}

func compile_push_const_fast(self method_compiler) {
}

func compile_find_const_fast(self method_compiler) {
}

func compile_set_call_flags(self method_compiler) {
}

func compile_allow_private(self method_compiler) {
}

func compile_send_method(self method_compiler) {
}

func compile_send_stack(self method_compiler) {
}

func compile_send_stack_with_block(self method_compiler) {
}

func compile_send_stack_with_splat(self method_compiler) {
}

func compile_send_super_stack_with_block(self method_compiler) {
}

func compile_send_super_stack_with_splat(self method_compiler) {
}

func compile_push_block(self method_compiler) {
}

func compile_passed_blockarg(self method_compiler) {
}

func compile_create_block(self method_compiler) {
}

func compile_cast_for_single_block_arg(self method_compiler) {
}

func compile_cast_for_multi_block_arg(self method_compiler) {
}

func compile_cast_for_splat_block_arg(self method_compiler) {
}

func compile_yield_stack(self method_compiler) {
}

func compile_yield_splat(self method_compiler) {
}

func compile_string_append(self method_compiler) {
}

func compile_string_build(self method_compiler) {
}

func compile_string_dup(self method_compiler) {
}

func compile_push_scope(self method_compiler) {
}

func compile_add_scope(self method_compiler) {
}

func compile_push_variables(self method_compiler) {
}

func compile_check_interrupts(self method_compiler) {
}

func compile_yield_debugger(self method_compiler) {
}

func compile_is_nil(self method_compiler) {
}

func compile_check_serial(self method_compiler) {
}

func compile_check_serial_private(self method_compiler) {
}

func compile_push_my_field(self method_compiler) {
}

func compile_store_my_field(self method_compiler) {
}

func compile_kind_of(self method_compiler) {
}

func compile_instance_of(self method_compiler) {
}

func compile_meta_push_neg_1(self method_compiler) {
}

func compile_meta_push_0(self method_compiler) {
}

func compile_meta_push_1(self method_compiler) {
}

func compile_meta_push_2(self method_compiler) {
}

func compile_meta_send_op_plus(self method_compiler) {
}

func compile_meta_send_op_minus(self method_compiler) {
}

func compile_meta_send_op_equal(self method_compiler) {
}

func compile_meta_send_op_lt(self method_compiler) {
}

func compile_meta_send_op_gt(self method_compiler) {
}

func compile_meta_send_op_tequal(self method_compiler) {
}

func compile_meta_send_call(self method_compiler) {
}

func compile_push_my_offset(self method_compiler) {
}

func compile_zsuper(self method_compiler) {
}

func compile_push_block_arg(self method_compiler) {
}

func compile_push_undef(self method_compiler) {
}

func compile_push_stack_local(self method_compiler) {
}

func compile_set_stack_local(self method_compiler) {
}

func compile_push_has_block(self method_compiler) {
}

func compile_push_proc(self method_compiler) {
}

func compile_check_frozen(self method_compiler) {
}

func compile_cast_multi_value(self method_compiler) {
}

func compile_invoke_primitive(self method_compiler) {
}

func compile_push_rubinius(self method_compiler) {
}

func compile_call_custom(self method_compiler) {
}

func compile_meta_to_s(self method_compiler) {
}

func compile_push_type(self method_compiler) {
}

func compile_push_mirror(self method_compiler) {
}
