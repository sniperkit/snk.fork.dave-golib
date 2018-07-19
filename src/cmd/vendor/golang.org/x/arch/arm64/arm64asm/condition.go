package arm64asm

func at_sys_cr_system_cond(instr uint32) bool {
	return sys_op_4((instr>>16)&0x7, 0x7, 0x8, (instr>>5)&0x7) == Sys_AT
}

func bfi_bfm_32m_bitfield_cond(instr uint32) bool {
	return (instr>>5)&0x1f != 0x1f && uint8((instr>>10)&0x3f) < uint8((instr>>16)&0x3f)
}

func bfi_bfm_64m_bitfield_cond(instr uint32) bool {
	return (instr>>5)&0x1f != 0x1f && uint8((instr>>10)&0x3f) < uint8((instr>>16)&0x3f)
}

func bfxil_bfm_32m_bitfield_cond(instr uint32) bool {
	return uint8((instr>>10)&0x3f) >= uint8((instr>>16)&0x3f)
}

func bfxil_bfm_64m_bitfield_cond(instr uint32) bool {
	return uint8((instr>>10)&0x3f) >= uint8((instr>>16)&0x3f)
}

func cinc_csinc_32_condsel_cond(instr uint32) bool {
	return instr&0x1f0000 != 0x1f0000 && instr&0xe000 != 0xe000 && instr&0x3e0 != 0x3e0 && (instr>>5)&0x1f == (instr>>16)&0x1f
}

func cinc_csinc_64_condsel_cond(instr uint32) bool {
	return instr&0x1f0000 != 0x1f0000 && instr&0xe000 != 0xe000 && instr&0x3e0 != 0x3e0 && (instr>>5)&0x1f == (instr>>16)&0x1f
}

func cinv_csinv_32_condsel_cond(instr uint32) bool {
	return instr&0x1f0000 != 0x1f0000 && instr&0xe000 != 0xe000 && instr&0x3e0 != 0x3e0 && (instr>>5)&0x1f == (instr>>16)&0x1f
}

func cinv_csinv_64_condsel_cond(instr uint32) bool {
	return instr&0x1f0000 != 0x1f0000 && instr&0xe000 != 0xe000 && instr&0x3e0 != 0x3e0 && (instr>>5)&0x1f == (instr>>16)&0x1f
}

func cneg_csneg_32_condsel_cond(instr uint32) bool {
	return instr&0xe000 != 0xe000 && (instr>>5)&0x1f == (instr>>16)&0x1f
}

func cneg_csneg_64_condsel_cond(instr uint32) bool {
	return instr&0xe000 != 0xe000 && (instr>>5)&0x1f == (instr>>16)&0x1f
}

func csinc_general_cond(instr uint32) bool {
	return instr&0xe000 != 0xe000
}
func csinv_general_cond(instr uint32) bool {
	return instr&0xe000 != 0xe000
}
func dc_sys_cr_system_cond(instr uint32) bool {
	return sys_op_4((instr>>16)&0x7, 0x7, (instr>>8)&0xf, (instr>>5)&0x7) == Sys_DC
}

func ic_sys_cr_system_cond(instr uint32) bool {
	return sys_op_4((instr>>16)&0x7, 0x7, (instr>>8)&0xf, (instr>>5)&0x7) == Sys_IC
}

func lsl_ubfm_32m_bitfield_cond(instr uint32) bool {
	return instr&0xfc00 != 0x7c00 && (instr>>10)&0x3f+1 == (instr>>16)&0x3f
}

func lsl_ubfm_64m_bitfield_cond(instr uint32) bool {
	return instr&0xfc00 != 0xfc00 && (instr>>10)&0x3f+1 == (instr>>16)&0x3f
}

func mov_orr_32_log_imm_cond(instr uint32) bool {
	return !move_wide_preferred_4((instr>>31)&0x1, (instr>>22)&0x1, (instr>>10)&0x3f, (instr>>16)&0x3f)
}

func mov_orr_64_log_imm_cond(instr uint32) bool {
	return !move_wide_preferred_4((instr>>31)&0x1, (instr>>22)&0x1, (instr>>10)&0x3f, (instr>>16)&0x3f)
}

func mov_movn_32_movewide_cond(instr uint32) bool {
	return !(is_zero((instr>>5)&0xffff) && (instr>>21)&0x3 != 0x0) && !is_ones_n16((instr>>5)&0xffff)
}

func mov_movn_64_movewide_cond(instr uint32) bool {
	return !(is_zero((instr>>5)&0xffff) && (instr>>21)&0x3 != 0x0)
}

func mov_add_32_addsub_imm_cond(instr uint32) bool {
	return instr&0x1f == 0x1f || (instr>>5)&0x1f == 0x1f
}

func mov_add_64_addsub_imm_cond(instr uint32) bool {
	return instr&0x1f == 0x1f || (instr>>5)&0x1f == 0x1f
}

func mov_movz_32_movewide_cond(instr uint32) bool {
	return !(is_zero((instr>>5)&0xffff) && (instr>>21)&0x3 != 0x0)
}

func mov_movz_64_movewide_cond(instr uint32) bool {
	return !(is_zero((instr>>5)&0xffff) && (instr>>21)&0x3 != 0x0)
}

func ror_extr_32_extract_cond(instr uint32) bool {
	return (instr>>5)&0x1f == (instr>>16)&0x1f
}

func ror_extr_64_extract_cond(instr uint32) bool {
	return (instr>>5)&0x1f == (instr>>16)&0x1f
}

func sbfiz_sbfm_32m_bitfield_cond(instr uint32) bool {
	return uint8((instr>>10)&0x3f) < uint8((instr>>16)&0x3f)
}

func sbfiz_sbfm_64m_bitfield_cond(instr uint32) bool {
	return uint8((instr>>10)&0x3f) < uint8((instr>>16)&0x3f)
}

func sbfx_sbfm_32m_bitfield_cond(instr uint32) bool {
	return bfxpreferred_4((instr>>31)&0x1, extract_bit((instr>>29)&0x3, 1), (instr>>10)&0x3f, (instr>>16)&0x3f)
}

func sbfx_sbfm_64m_bitfield_cond(instr uint32) bool {
	return bfxpreferred_4((instr>>31)&0x1, extract_bit((instr>>29)&0x3, 1), (instr>>10)&0x3f, (instr>>16)&0x3f)
}

func tlbi_sys_cr_system_cond(instr uint32) bool {
	return sys_op_4((instr>>16)&0x7, 0x8, (instr>>8)&0xf, (instr>>5)&0x7) == Sys_TLBI
}

func ubfiz_ubfm_32m_bitfield_cond(instr uint32) bool {
	return uint8((instr>>10)&0x3f) < uint8((instr>>16)&0x3f)
}

func ubfiz_ubfm_64m_bitfield_cond(instr uint32) bool {
	return uint8((instr>>10)&0x3f) < uint8((instr>>16)&0x3f)
}

func ubfx_ubfm_32m_bitfield_cond(instr uint32) bool {
	return bfxpreferred_4((instr>>31)&0x1, extract_bit((instr>>29)&0x3, 1), (instr>>10)&0x3f, (instr>>16)&0x3f)
}

func ubfx_ubfm_64m_bitfield_cond(instr uint32) bool {
	return bfxpreferred_4((instr>>31)&0x1, extract_bit((instr>>29)&0x3, 1), (instr>>10)&0x3f, (instr>>16)&0x3f)
}

func fcvtzs_asisdshf_c_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func fcvtzs_asimdshf_c_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func fcvtzu_asisdshf_c_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func fcvtzu_asimdshf_c_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func mov_umov_asimdins_w_w_cond(instr uint32) bool {
	return ((instr>>16)&0x1f)&0x7 == 0x4
}

func mov_umov_asimdins_x_x_cond(instr uint32) bool {
	return ((instr>>16)&0x1f)&0xf == 0x8
}

func mov_orr_asimdsame_only_cond(instr uint32) bool {
	return (instr>>16)&0x1f == (instr>>5)&0x1f
}

func rshrn_asimdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func scvtf_asisdshf_c_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func scvtf_asimdshf_c_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func shl_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func shl_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func shrn_asimdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sli_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sli_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqrshrn_asisdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqrshrn_asimdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqrshrun_asisdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqrshrun_asimdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqshl_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqshl_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqshlu_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqshlu_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqshrn_asisdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqshrn_asimdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqshrun_asisdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sqshrun_asimdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sri_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sri_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func srshr_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func srshr_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func srsra_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func srsra_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sshll_asimdshf_l_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sshr_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sshr_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func ssra_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func ssra_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func sxtl_sshll_asimdshf_l_cond(instr uint32) bool {
	return instr&0x780000 != 0x0 && bit_count((instr>>19)&0xf) == 1
}

func ucvtf_asisdshf_c_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func ucvtf_asimdshf_c_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func uqrshrn_asisdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func uqrshrn_asimdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func uqshl_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func uqshl_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func uqshrn_asisdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func uqshrn_asimdshf_n_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func urshr_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func urshr_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func ursra_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func ursra_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func ushll_asimdshf_l_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func ushr_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func ushr_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func usra_asisdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func usra_asimdshf_r_cond(instr uint32) bool {
	return instr&0x780000 != 0x0
}
func uxtl_ushll_asimdshf_l_cond(instr uint32) bool {
	return instr&0x780000 != 0x0 && bit_count((instr>>19)&0xf) == 1
}
