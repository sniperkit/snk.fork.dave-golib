/*
Sniperkit-Bot
- Status: analyzed
*/

package wasm

type PackageState struct {
	obj             *obj.PackageState
	objabi          *objabi.PackageState
	sys             *sys.PackageState
	Anames          []string
	Linkwasm        obj.LinkArch
	Register        map[string]int16
	deferreturn     *obj.LSym
	gcWriteBarrier  *obj.LSym
	jmpdefer        *obj.LSym
	morestack       *obj.LSym
	morestackNoCtxt *obj.LSym
	registerNames   []string
	sigpanic        *obj.LSym
	unaryDst        map[obj.As]bool
}

func NewPackageState(obj_pstate *obj.PackageState, objabi_pstate *objabi.PackageState, sys_pstate *sys.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.obj = obj_pstate
	pstate.objabi = objabi_pstate
	pstate.sys = sys_pstate
	pstate.Anames = []string{
		obj.A_ARCHSPECIFIC: "CallImport",
		"Get",
		"Set",
		"Tee",
		"Not",
		"Unreachable",
		"Nop",
		"Block",
		"Loop",
		"If",
		"Else",
		"End",
		"Br",
		"BrIf",
		"BrTable",
		"Return",
		"Call",
		"CallIndirect",
		"Drop",
		"Select",
		"I32Load",
		"I64Load",
		"F32Load",
		"F64Load",
		"I32Load8S",
		"I32Load8U",
		"I32Load16S",
		"I32Load16U",
		"I64Load8S",
		"I64Load8U",
		"I64Load16S",
		"I64Load16U",
		"I64Load32S",
		"I64Load32U",
		"I32Store",
		"I64Store",
		"F32Store",
		"F64Store",
		"I32Store8",
		"I32Store16",
		"I64Store8",
		"I64Store16",
		"I64Store32",
		"CurrentMemory",
		"GrowMemory",
		"I32Const",
		"I64Const",
		"F32Const",
		"F64Const",
		"I32Eqz",
		"I32Eq",
		"I32Ne",
		"I32LtS",
		"I32LtU",
		"I32GtS",
		"I32GtU",
		"I32LeS",
		"I32LeU",
		"I32GeS",
		"I32GeU",
		"I64Eqz",
		"I64Eq",
		"I64Ne",
		"I64LtS",
		"I64LtU",
		"I64GtS",
		"I64GtU",
		"I64LeS",
		"I64LeU",
		"I64GeS",
		"I64GeU",
		"F32Eq",
		"F32Ne",
		"F32Lt",
		"F32Gt",
		"F32Le",
		"F32Ge",
		"F64Eq",
		"F64Ne",
		"F64Lt",
		"F64Gt",
		"F64Le",
		"F64Ge",
		"I32Clz",
		"I32Ctz",
		"I32Popcnt",
		"I32Add",
		"I32Sub",
		"I32Mul",
		"I32DivS",
		"I32DivU",
		"I32RemS",
		"I32RemU",
		"I32And",
		"I32Or",
		"I32Xor",
		"I32Shl",
		"I32ShrS",
		"I32ShrU",
		"I32Rotl",
		"I32Rotr",
		"I64Clz",
		"I64Ctz",
		"I64Popcnt",
		"I64Add",
		"I64Sub",
		"I64Mul",
		"I64DivS",
		"I64DivU",
		"I64RemS",
		"I64RemU",
		"I64And",
		"I64Or",
		"I64Xor",
		"I64Shl",
		"I64ShrS",
		"I64ShrU",
		"I64Rotl",
		"I64Rotr",
		"F32Abs",
		"F32Neg",
		"F32Ceil",
		"F32Floor",
		"F32Trunc",
		"F32Nearest",
		"F32Sqrt",
		"F32Add",
		"F32Sub",
		"F32Mul",
		"F32Div",
		"F32Min",
		"F32Max",
		"F32Copysign",
		"F64Abs",
		"F64Neg",
		"F64Ceil",
		"F64Floor",
		"F64Trunc",
		"F64Nearest",
		"F64Sqrt",
		"F64Add",
		"F64Sub",
		"F64Mul",
		"F64Div",
		"F64Min",
		"F64Max",
		"F64Copysign",
		"I32WrapI64",
		"I32TruncSF32",
		"I32TruncUF32",
		"I32TruncSF64",
		"I32TruncUF64",
		"I64ExtendSI32",
		"I64ExtendUI32",
		"I64TruncSF32",
		"I64TruncUF32",
		"I64TruncSF64",
		"I64TruncUF64",
		"F32ConvertSI32",
		"F32ConvertUI32",
		"F32ConvertSI64",
		"F32ConvertUI64",
		"F32DemoteF64",
		"F64ConvertSI32",
		"F64ConvertUI32",
		"F64ConvertSI64",
		"F64ConvertUI64",
		"F64PromoteF32",
		"I32ReinterpretF32",
		"I64ReinterpretF64",
		"F32ReinterpretI32",
		"F64ReinterpretI64",
		"RESUMEPOINT",
		"CALLNORESUME",
		"RETUNWIND",
		"MOVB",
		"MOVH",
		"MOVW",
		"MOVD",
		"WORD",
		"LAST",
	}
	pstate.unaryDst = map[obj.As]bool{
		ASet:          true,
		ATee:          true,
		ACall:         true,
		ACallIndirect: true,
		ACallImport:   true,
		ABr:           true,
		ABrIf:         true,
		ABrTable:      true,
		AI32Store:     true,
		AI64Store:     true,
		AF32Store:     true,
		AF64Store:     true,
		AI32Store8:    true,
		AI32Store16:   true,
		AI64Store8:    true,
		AI64Store16:   true,
		AI64Store32:   true,
		ACALLNORESUME: true,
	}
	pstate.Register = map[string]int16{
		"PC_F": REG_PC_F,
		"PC_B": REG_PC_B,
		"SP":   REG_SP,
		"CTXT": REG_CTXT,
		"g":    REG_g,
		"RET0": REG_RET0,
		"RET1": REG_RET1,
		"RET2": REG_RET2,
		"RET3": REG_RET3,
		"RUN":  REG_RUN,

		"R0":  REG_R0,
		"R1":  REG_R1,
		"R2":  REG_R2,
		"R3":  REG_R3,
		"R4":  REG_R4,
		"R5":  REG_R5,
		"R6":  REG_R6,
		"R7":  REG_R7,
		"R8":  REG_R8,
		"R9":  REG_R9,
		"R10": REG_R10,
		"R11": REG_R11,
		"R12": REG_R12,
		"R13": REG_R13,
		"R14": REG_R14,
		"R15": REG_R15,

		"F0":  REG_F0,
		"F1":  REG_F1,
		"F2":  REG_F2,
		"F3":  REG_F3,
		"F4":  REG_F4,
		"F5":  REG_F5,
		"F6":  REG_F6,
		"F7":  REG_F7,
		"F8":  REG_F8,
		"F9":  REG_F9,
		"F10": REG_F10,
		"F11": REG_F11,
		"F12": REG_F12,
		"F13": REG_F13,
		"F14": REG_F14,
		"F15": REG_F15,
	}
	pstate.Linkwasm = obj.LinkArch{
		Arch:       pstate.sys.ArchWasm,
		Init:       pstate.instinit,
		Preprocess: pstate.preprocess,
		Assemble:   assemble,
		UnaryDst:   pstate.unaryDst,
	}
	return pstate
}