package gc

import "github.com/dave/golib/src/cmd/compile/internal/types"

func (psess *PackageSession) runtimeTypes() []*types.Type {
	var typs [113]*types.Type
	typs[0] = psess.types.Bytetype
	typs[1] = psess.types.NewPtr(typs[0])
	typs[2] = psess.types.Types[TANY]
	typs[3] = psess.types.NewPtr(typs[2])
	typs[4] = psess.functype(nil, []*Node{psess.anonfield(typs[1])}, []*Node{psess.anonfield(typs[3])})
	typs[5] = psess.functype(nil, nil, nil)
	typs[6] = psess.types.Types[TINTER]
	typs[7] = psess.functype(nil, []*Node{psess.anonfield(typs[6])}, nil)
	typs[8] = psess.types.Types[TINT32]
	typs[9] = psess.types.NewPtr(typs[8])
	typs[10] = psess.functype(nil, []*Node{psess.anonfield(typs[9])}, []*Node{psess.anonfield(typs[6])})
	typs[11] = psess.types.Types[TBOOL]
	typs[12] = psess.functype(nil, []*Node{psess.anonfield(typs[11])}, nil)
	typs[13] = psess.types.Types[TFLOAT64]
	typs[14] = psess.functype(nil, []*Node{psess.anonfield(typs[13])}, nil)
	typs[15] = psess.types.Types[TINT64]
	typs[16] = psess.functype(nil, []*Node{psess.anonfield(typs[15])}, nil)
	typs[17] = psess.types.Types[TUINT64]
	typs[18] = psess.functype(nil, []*Node{psess.anonfield(typs[17])}, nil)
	typs[19] = psess.types.Types[TCOMPLEX128]
	typs[20] = psess.functype(nil, []*Node{psess.anonfield(typs[19])}, nil)
	typs[21] = psess.types.Types[TSTRING]
	typs[22] = psess.functype(nil, []*Node{psess.anonfield(typs[21])}, nil)
	typs[23] = psess.functype(nil, []*Node{psess.anonfield(typs[2])}, nil)
	typs[24] = psess.types.NewArray(typs[0], 32)
	typs[25] = psess.types.NewPtr(typs[24])
	typs[26] = psess.functype(nil, []*Node{psess.anonfield(typs[25]), psess.anonfield(typs[21]), psess.anonfield(typs[21])}, []*Node{psess.anonfield(typs[21])})
	typs[27] = psess.functype(nil, []*Node{psess.anonfield(typs[25]), psess.anonfield(typs[21]), psess.anonfield(typs[21]), psess.anonfield(typs[21])}, []*Node{psess.anonfield(typs[21])})
	typs[28] = psess.functype(nil, []*Node{psess.anonfield(typs[25]), psess.anonfield(typs[21]), psess.anonfield(typs[21]), psess.anonfield(typs[21]), psess.anonfield(typs[21])}, []*Node{psess.anonfield(typs[21])})
	typs[29] = psess.functype(nil, []*Node{psess.anonfield(typs[25]), psess.anonfield(typs[21]), psess.anonfield(typs[21]), psess.anonfield(typs[21]), psess.anonfield(typs[21]), psess.anonfield(typs[21])}, []*Node{psess.anonfield(typs[21])})
	typs[30] = psess.types.NewSlice(typs[21])
	typs[31] = psess.functype(nil, []*Node{psess.anonfield(typs[25]), psess.anonfield(typs[30])}, []*Node{psess.anonfield(typs[21])})
	typs[32] = psess.types.Types[TINT]
	typs[33] = psess.functype(nil, []*Node{psess.anonfield(typs[21]), psess.anonfield(typs[21])}, []*Node{psess.anonfield(typs[32])})
	typs[34] = psess.types.NewArray(typs[0], 4)
	typs[35] = psess.types.NewPtr(typs[34])
	typs[36] = psess.functype(nil, []*Node{psess.anonfield(typs[35]), psess.anonfield(typs[15])}, []*Node{psess.anonfield(typs[21])})
	typs[37] = psess.types.NewSlice(typs[0])
	typs[38] = psess.functype(nil, []*Node{psess.anonfield(typs[25]), psess.anonfield(typs[37])}, []*Node{psess.anonfield(typs[21])})
	typs[39] = psess.functype(nil, []*Node{psess.anonfield(typs[37])}, []*Node{psess.anonfield(typs[21])})
	typs[40] = psess.types.Runetype
	typs[41] = psess.types.NewSlice(typs[40])
	typs[42] = psess.functype(nil, []*Node{psess.anonfield(typs[25]), psess.anonfield(typs[41])}, []*Node{psess.anonfield(typs[21])})
	typs[43] = psess.functype(nil, []*Node{psess.anonfield(typs[25]), psess.anonfield(typs[21])}, []*Node{psess.anonfield(typs[37])})
	typs[44] = psess.types.NewArray(typs[40], 32)
	typs[45] = psess.types.NewPtr(typs[44])
	typs[46] = psess.functype(nil, []*Node{psess.anonfield(typs[45]), psess.anonfield(typs[21])}, []*Node{psess.anonfield(typs[41])})
	typs[47] = psess.types.Types[TUINTPTR]
	typs[48] = psess.functype(nil, []*Node{psess.anonfield(typs[2]), psess.anonfield(typs[2]), psess.anonfield(typs[47])}, []*Node{psess.anonfield(typs[32])})
	typs[49] = psess.functype(nil, []*Node{psess.anonfield(typs[2]), psess.anonfield(typs[2])}, []*Node{psess.anonfield(typs[32])})
	typs[50] = psess.functype(nil, []*Node{psess.anonfield(typs[21]), psess.anonfield(typs[32])}, []*Node{psess.anonfield(typs[40]), psess.anonfield(typs[32])})
	typs[51] = psess.functype(nil, []*Node{psess.anonfield(typs[21])}, []*Node{psess.anonfield(typs[32])})
	typs[52] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[2])}, []*Node{psess.anonfield(typs[2])})
	typs[53] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[3])}, []*Node{psess.anonfield(typs[2])})
	typs[54] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[2])}, []*Node{psess.anonfield(typs[2]), psess.anonfield(typs[11])})
	typs[55] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[1]), psess.anonfield(typs[1])}, nil)
	typs[56] = psess.functype(nil, []*Node{psess.anonfield(typs[1])}, nil)
	typs[57] = psess.types.NewPtr(typs[47])
	typs[58] = psess.types.Types[TUNSAFEPTR]
	typs[59] = psess.functype(nil, []*Node{psess.anonfield(typs[57]), psess.anonfield(typs[58]), psess.anonfield(typs[58])}, []*Node{psess.anonfield(typs[11])})
	typs[60] = psess.types.Types[TUINT32]
	typs[61] = psess.functype(nil, nil, []*Node{psess.anonfield(typs[60])})
	typs[62] = psess.types.NewMap(typs[2], typs[2])
	typs[63] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[15]), psess.anonfield(typs[3])}, []*Node{psess.anonfield(typs[62])})
	typs[64] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[32]), psess.anonfield(typs[3])}, []*Node{psess.anonfield(typs[62])})
	typs[65] = psess.functype(nil, nil, []*Node{psess.anonfield(typs[62])})
	typs[66] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62]), psess.anonfield(typs[3])}, []*Node{psess.anonfield(typs[3])})
	typs[67] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62]), psess.anonfield(typs[2])}, []*Node{psess.anonfield(typs[3])})
	typs[68] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62]), psess.anonfield(typs[3]), psess.anonfield(typs[1])}, []*Node{psess.anonfield(typs[3])})
	typs[69] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62]), psess.anonfield(typs[3])}, []*Node{psess.anonfield(typs[3]), psess.anonfield(typs[11])})
	typs[70] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62]), psess.anonfield(typs[2])}, []*Node{psess.anonfield(typs[3]), psess.anonfield(typs[11])})
	typs[71] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62]), psess.anonfield(typs[3]), psess.anonfield(typs[1])}, []*Node{psess.anonfield(typs[3]), psess.anonfield(typs[11])})
	typs[72] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62]), psess.anonfield(typs[3])}, nil)
	typs[73] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62]), psess.anonfield(typs[2])}, nil)
	typs[74] = psess.functype(nil, []*Node{psess.anonfield(typs[3])}, nil)
	typs[75] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[62])}, nil)
	typs[76] = psess.types.NewChan(typs[2], types.Cboth)
	typs[77] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[15])}, []*Node{psess.anonfield(typs[76])})
	typs[78] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[32])}, []*Node{psess.anonfield(typs[76])})
	typs[79] = psess.types.NewChan(typs[2], types.Crecv)
	typs[80] = psess.functype(nil, []*Node{psess.anonfield(typs[79]), psess.anonfield(typs[3])}, nil)
	typs[81] = psess.functype(nil, []*Node{psess.anonfield(typs[79]), psess.anonfield(typs[3])}, []*Node{psess.anonfield(typs[11])})
	typs[82] = psess.types.NewChan(typs[2], types.Csend)
	typs[83] = psess.functype(nil, []*Node{psess.anonfield(typs[82]), psess.anonfield(typs[3])}, nil)
	typs[84] = psess.types.NewArray(typs[0], 3)
	typs[85] = psess.tostruct([]*Node{psess.namedfield("enabled", typs[11]), psess.namedfield("pad", typs[84]), psess.namedfield("needed", typs[11]), psess.namedfield("cgo", typs[11]), psess.namedfield("alignme", typs[17])})
	typs[86] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[3]), psess.anonfield(typs[3])}, nil)
	typs[87] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[3])}, nil)
	typs[88] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[2]), psess.anonfield(typs[2])}, []*Node{psess.anonfield(typs[32])})
	typs[89] = psess.functype(nil, []*Node{psess.anonfield(typs[82]), psess.anonfield(typs[3])}, []*Node{psess.anonfield(typs[11])})
	typs[90] = psess.functype(nil, []*Node{psess.anonfield(typs[3]), psess.anonfield(typs[79])}, []*Node{psess.anonfield(typs[11])})
	typs[91] = psess.types.NewPtr(typs[11])
	typs[92] = psess.functype(nil, []*Node{psess.anonfield(typs[3]), psess.anonfield(typs[91]), psess.anonfield(typs[79])}, []*Node{psess.anonfield(typs[11])})
	typs[93] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[1]), psess.anonfield(typs[32])}, []*Node{psess.anonfield(typs[32]), psess.anonfield(typs[11])})
	typs[94] = psess.types.NewSlice(typs[2])
	typs[95] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[32]), psess.anonfield(typs[32])}, []*Node{psess.anonfield(typs[94])})
	typs[96] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[15]), psess.anonfield(typs[15])}, []*Node{psess.anonfield(typs[94])})
	typs[97] = psess.functype(nil, []*Node{psess.anonfield(typs[1]), psess.anonfield(typs[94]), psess.anonfield(typs[32])}, []*Node{psess.anonfield(typs[94])})
	typs[98] = psess.functype(nil, []*Node{psess.anonfield(typs[3]), psess.anonfield(typs[3]), psess.anonfield(typs[47])}, nil)
	typs[99] = psess.functype(nil, []*Node{psess.anonfield(typs[58]), psess.anonfield(typs[47])}, nil)
	typs[100] = psess.functype(nil, []*Node{psess.anonfield(typs[3]), psess.anonfield(typs[3]), psess.anonfield(typs[47])}, []*Node{psess.anonfield(typs[11])})
	typs[101] = psess.functype(nil, []*Node{psess.anonfield(typs[3]), psess.anonfield(typs[3])}, []*Node{psess.anonfield(typs[11])})
	typs[102] = psess.functype(nil, []*Node{psess.anonfield(typs[15]), psess.anonfield(typs[15])}, []*Node{psess.anonfield(typs[15])})
	typs[103] = psess.functype(nil, []*Node{psess.anonfield(typs[17]), psess.anonfield(typs[17])}, []*Node{psess.anonfield(typs[17])})
	typs[104] = psess.functype(nil, []*Node{psess.anonfield(typs[13])}, []*Node{psess.anonfield(typs[15])})
	typs[105] = psess.functype(nil, []*Node{psess.anonfield(typs[13])}, []*Node{psess.anonfield(typs[17])})
	typs[106] = psess.functype(nil, []*Node{psess.anonfield(typs[13])}, []*Node{psess.anonfield(typs[60])})
	typs[107] = psess.functype(nil, []*Node{psess.anonfield(typs[15])}, []*Node{psess.anonfield(typs[13])})
	typs[108] = psess.functype(nil, []*Node{psess.anonfield(typs[17])}, []*Node{psess.anonfield(typs[13])})
	typs[109] = psess.functype(nil, []*Node{psess.anonfield(typs[60])}, []*Node{psess.anonfield(typs[13])})
	typs[110] = psess.functype(nil, []*Node{psess.anonfield(typs[19]), psess.anonfield(typs[19])}, []*Node{psess.anonfield(typs[19])})
	typs[111] = psess.functype(nil, []*Node{psess.anonfield(typs[47])}, nil)
	typs[112] = psess.functype(nil, []*Node{psess.anonfield(typs[47]), psess.anonfield(typs[47])}, nil)
	return typs[:]
}
