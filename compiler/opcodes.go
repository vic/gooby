package compiler

var opcode_compilers = []func(*method_compiler){
	(*method_compiler).noop,
	(*method_compiler).push_nil,
	(*method_compiler).push_true,
	(*method_compiler).push_false,
	(*method_compiler).push_int,
	(*method_compiler).push_self,
	(*method_compiler).set_literal,
	(*method_compiler).push_literal,
	(*method_compiler).goto_,
	(*method_compiler).goto_if_false,
	(*method_compiler).goto_if_true,
	(*method_compiler).ret,
	(*method_compiler).swap_stack,
	(*method_compiler).dup_top,
	(*method_compiler).dup_many,
	(*method_compiler).pop,
	(*method_compiler).pop_many,
	(*method_compiler).rotate,
	(*method_compiler).move_down,
	(*method_compiler).set_local,
	(*method_compiler).push_local,
	(*method_compiler).push_local_depth,
	(*method_compiler).set_local_depth,
	(*method_compiler).passed_arg,
	(*method_compiler).push_current_exception,
	(*method_compiler).clear_exception,
	(*method_compiler).push_exception_state,
	(*method_compiler).restore_exception_state,
	(*method_compiler).raise_exc,
	(*method_compiler).setup_unwind,
	(*method_compiler).pop_unwind,
	(*method_compiler).raise_return,
	(*method_compiler).ensure_return,
	(*method_compiler).raise_break,
	(*method_compiler).reraise,
	(*method_compiler).make_array,
	(*method_compiler).cast_array,
	(*method_compiler).shift_array,
	(*method_compiler).set_ivar,
	(*method_compiler).push_ivar,
	(*method_compiler).push_const,
	(*method_compiler).set_const,
	(*method_compiler).set_const_at,
	(*method_compiler).find_const,
	(*method_compiler).push_cpath_top,
	(*method_compiler).push_const_fast,
	(*method_compiler).find_const_fast,
	(*method_compiler).set_call_flags,
	(*method_compiler).allow_private,
	(*method_compiler).send_method,
	(*method_compiler).send_stack,
	(*method_compiler).send_stack_with_block,
	(*method_compiler).send_stack_with_splat,
	(*method_compiler).send_super_stack_with_block,
	(*method_compiler).send_super_stack_with_splat,
	(*method_compiler).push_block,
	(*method_compiler).passed_blockarg,
	(*method_compiler).create_block,
	(*method_compiler).cast_for_single_block_arg,
	(*method_compiler).cast_for_multi_block_arg,
	(*method_compiler).cast_for_splat_block_arg,
	(*method_compiler).yield_stack,
	(*method_compiler).yield_splat,
	(*method_compiler).string_append,
	(*method_compiler).string_build,
	(*method_compiler).string_dup,
	(*method_compiler).push_scope,
	(*method_compiler).add_scope,
	(*method_compiler).push_variables,
	(*method_compiler).check_interrupts,
	(*method_compiler).yield_debugger,
	(*method_compiler).is_nil,
	(*method_compiler).check_serial,
	(*method_compiler).check_serial_private,
	(*method_compiler).push_my_field,
	(*method_compiler).store_my_field,
	(*method_compiler).kind_of,
	(*method_compiler).instance_of,
	(*method_compiler).meta_push_neg_1,
	(*method_compiler).meta_push_0,
	(*method_compiler).meta_push_1,
	(*method_compiler).meta_push_2,
	(*method_compiler).meta_send_op_plus,
	(*method_compiler).meta_send_op_minus,
	(*method_compiler).meta_send_op_equal,
	(*method_compiler).meta_send_op_lt,
	(*method_compiler).meta_send_op_gt,
	(*method_compiler).meta_send_op_tequal,
	(*method_compiler).meta_send_call,
	(*method_compiler).push_my_offset,
	(*method_compiler).zsuper,
	(*method_compiler).push_block_arg,
	(*method_compiler).push_undef,
	(*method_compiler).push_stack_local,
	(*method_compiler).set_stack_local,
	(*method_compiler).push_has_block,
	(*method_compiler).push_proc,
	(*method_compiler).check_frozen,
	(*method_compiler).cast_multi_value,
	(*method_compiler).invoke_primitive,
	(*method_compiler).push_rubinius,
	(*method_compiler).call_custom,
	(*method_compiler).meta_to_s,
	(*method_compiler).push_type,
	(*method_compiler).push_mirror,
}
