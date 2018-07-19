package arm64asm

const (
	_ Op = iota
	ABS
	ADC
	ADCS
	ADD
	ADDHN
	ADDHN2
	ADDP
	ADDS
	ADDV
	ADR
	ADRP
	AESD
	AESE
	AESIMC
	AESMC
	AND
	ANDS
	ASR
	ASRV
	AT
	B
	BFI
	BFM
	BFXIL
	BIC
	BICS
	BIF
	BIT
	BL
	BLR
	BR
	BRK
	BSL
	CBNZ
	CBZ
	CCMN
	CCMP
	CINC
	CINV
	CLREX
	CLS
	CLZ
	CMEQ
	CMGE
	CMGT
	CMHI
	CMHS
	CMLE
	CMLT
	CMN
	CMP
	CMTST
	CNEG
	CNT
	CRC32B
	CRC32CB
	CRC32CH
	CRC32CW
	CRC32CX
	CRC32H
	CRC32W
	CRC32X
	CSEL
	CSET
	CSETM
	CSINC
	CSINV
	CSNEG
	DC
	DCPS1
	DCPS2
	DCPS3
	DMB
	DRPS
	DSB
	DUP
	EON
	EOR
	ERET
	EXT
	EXTR
	FABD
	FABS
	FACGE
	FACGT
	FADD
	FADDP
	FCCMP
	FCCMPE
	FCMEQ
	FCMGE
	FCMGT
	FCMLE
	FCMLT
	FCMP
	FCMPE
	FCSEL
	FCVT
	FCVTAS
	FCVTAU
	FCVTL
	FCVTL2
	FCVTMS
	FCVTMU
	FCVTN
	FCVTN2
	FCVTNS
	FCVTNU
	FCVTPS
	FCVTPU
	FCVTXN
	FCVTXN2
	FCVTZS
	FCVTZU
	FDIV
	FMADD
	FMAX
	FMAXNM
	FMAXNMP
	FMAXNMV
	FMAXP
	FMAXV
	FMIN
	FMINNM
	FMINNMP
	FMINNMV
	FMINP
	FMINV
	FMLA
	FMLS
	FMOV
	FMSUB
	FMUL
	FMULX
	FNEG
	FNMADD
	FNMSUB
	FNMUL
	FRECPE
	FRECPS
	FRECPX
	FRINTA
	FRINTI
	FRINTM
	FRINTN
	FRINTP
	FRINTX
	FRINTZ
	FRSQRTE
	FRSQRTS
	FSQRT
	FSUB
	HINT
	HLT
	HVC
	IC
	INS
	ISB
	LD1
	LD1R
	LD2
	LD2R
	LD3
	LD3R
	LD4
	LD4R
	LDAR
	LDARB
	LDARH
	LDAXP
	LDAXR
	LDAXRB
	LDAXRH
	LDNP
	LDP
	LDPSW
	LDR
	LDRB
	LDRH
	LDRSB
	LDRSH
	LDRSW
	LDTR
	LDTRB
	LDTRH
	LDTRSB
	LDTRSH
	LDTRSW
	LDUR
	LDURB
	LDURH
	LDURSB
	LDURSH
	LDURSW
	LDXP
	LDXR
	LDXRB
	LDXRH
	LSL
	LSLV
	LSR
	LSRV
	MADD
	MLA
	MLS
	MNEG
	MOV
	MOVI
	MOVK
	MOVN
	MOVZ
	MRS
	MSR
	MSUB
	MUL
	MVN
	MVNI
	NEG
	NEGS
	NGC
	NGCS
	NOP
	NOT
	ORN
	ORR
	PMUL
	PMULL
	PMULL2
	PRFM
	PRFUM
	RADDHN
	RADDHN2
	RBIT
	RET
	REV
	REV16
	REV32
	REV64
	ROR
	RORV
	RSHRN
	RSHRN2
	RSUBHN
	RSUBHN2
	SABA
	SABAL
	SABAL2
	SABD
	SABDL
	SABDL2
	SADALP
	SADDL
	SADDL2
	SADDLP
	SADDLV
	SADDW
	SADDW2
	SBC
	SBCS
	SBFIZ
	SBFM
	SBFX
	SCVTF
	SDIV
	SEV
	SEVL
	SHA1C
	SHA1H
	SHA1M
	SHA1P
	SHA1SU0
	SHA1SU1
	SHA256H
	SHA256H2
	SHA256SU0
	SHA256SU1
	SHADD
	SHL
	SHLL
	SHLL2
	SHRN
	SHRN2
	SHSUB
	SLI
	SMADDL
	SMAX
	SMAXP
	SMAXV
	SMC
	SMIN
	SMINP
	SMINV
	SMLAL
	SMLAL2
	SMLSL
	SMLSL2
	SMNEGL
	SMOV
	SMSUBL
	SMULH
	SMULL
	SMULL2
	SQABS
	SQADD
	SQDMLAL
	SQDMLAL2
	SQDMLSL
	SQDMLSL2
	SQDMULH
	SQDMULL
	SQDMULL2
	SQNEG
	SQRDMULH
	SQRSHL
	SQRSHRN
	SQRSHRN2
	SQRSHRUN
	SQRSHRUN2
	SQSHL
	SQSHLU
	SQSHRN
	SQSHRN2
	SQSHRUN
	SQSHRUN2
	SQSUB
	SQXTN
	SQXTN2
	SQXTUN
	SQXTUN2
	SRHADD
	SRI
	SRSHL
	SRSHR
	SRSRA
	SSHL
	SSHLL
	SSHLL2
	SSHR
	SSRA
	SSUBL
	SSUBL2
	SSUBW
	SSUBW2
	ST1
	ST2
	ST3
	ST4
	STLR
	STLRB
	STLRH
	STLXP
	STLXR
	STLXRB
	STLXRH
	STNP
	STP
	STR
	STRB
	STRH
	STTR
	STTRB
	STTRH
	STUR
	STURB
	STURH
	STXP
	STXR
	STXRB
	STXRH
	SUB
	SUBHN
	SUBHN2
	SUBS
	SUQADD
	SVC
	SXTB
	SXTH
	SXTL
	SXTL2
	SXTW
	SYS
	SYSL
	TBL
	TBNZ
	TBX
	TBZ
	TLBI
	TRN1
	TRN2
	TST
	UABA
	UABAL
	UABAL2
	UABD
	UABDL
	UABDL2
	UADALP
	UADDL
	UADDL2
	UADDLP
	UADDLV
	UADDW
	UADDW2
	UBFIZ
	UBFM
	UBFX
	UCVTF
	UDIV
	UHADD
	UHSUB
	UMADDL
	UMAX
	UMAXP
	UMAXV
	UMIN
	UMINP
	UMINV
	UMLAL
	UMLAL2
	UMLSL
	UMLSL2
	UMNEGL
	UMOV
	UMSUBL
	UMULH
	UMULL
	UMULL2
	UQADD
	UQRSHL
	UQRSHRN
	UQRSHRN2
	UQSHL
	UQSHRN
	UQSHRN2
	UQSUB
	UQXTN
	UQXTN2
	URECPE
	URHADD
	URSHL
	URSHR
	URSQRTE
	URSRA
	USHL
	USHLL
	USHLL2
	USHR
	USQADD
	USRA
	USUBL
	USUBL2
	USUBW
	USUBW2
	UXTB
	UXTH
	UXTL
	UXTL2
	UZP1
	UZP2
	WFE
	WFI
	XTN
	XTN2
	YIELD
	ZIP1
	ZIP2
)

var opstr = [...]string{
	ABS:       "ABS",
	ADC:       "ADC",
	ADCS:      "ADCS",
	ADD:       "ADD",
	ADDHN:     "ADDHN",
	ADDHN2:    "ADDHN2",
	ADDP:      "ADDP",
	ADDS:      "ADDS",
	ADDV:      "ADDV",
	ADR:       "ADR",
	ADRP:      "ADRP",
	AESD:      "AESD",
	AESE:      "AESE",
	AESIMC:    "AESIMC",
	AESMC:     "AESMC",
	AND:       "AND",
	ANDS:      "ANDS",
	ASR:       "ASR",
	ASRV:      "ASRV",
	AT:        "AT",
	B:         "B",
	BFI:       "BFI",
	BFM:       "BFM",
	BFXIL:     "BFXIL",
	BIC:       "BIC",
	BICS:      "BICS",
	BIF:       "BIF",
	BIT:       "BIT",
	BL:        "BL",
	BLR:       "BLR",
	BR:        "BR",
	BRK:       "BRK",
	BSL:       "BSL",
	CBNZ:      "CBNZ",
	CBZ:       "CBZ",
	CCMN:      "CCMN",
	CCMP:      "CCMP",
	CINC:      "CINC",
	CINV:      "CINV",
	CLREX:     "CLREX",
	CLS:       "CLS",
	CLZ:       "CLZ",
	CMEQ:      "CMEQ",
	CMGE:      "CMGE",
	CMGT:      "CMGT",
	CMHI:      "CMHI",
	CMHS:      "CMHS",
	CMLE:      "CMLE",
	CMLT:      "CMLT",
	CMN:       "CMN",
	CMP:       "CMP",
	CMTST:     "CMTST",
	CNEG:      "CNEG",
	CNT:       "CNT",
	CRC32B:    "CRC32B",
	CRC32CB:   "CRC32CB",
	CRC32CH:   "CRC32CH",
	CRC32CW:   "CRC32CW",
	CRC32CX:   "CRC32CX",
	CRC32H:    "CRC32H",
	CRC32W:    "CRC32W",
	CRC32X:    "CRC32X",
	CSEL:      "CSEL",
	CSET:      "CSET",
	CSETM:     "CSETM",
	CSINC:     "CSINC",
	CSINV:     "CSINV",
	CSNEG:     "CSNEG",
	DC:        "DC",
	DCPS1:     "DCPS1",
	DCPS2:     "DCPS2",
	DCPS3:     "DCPS3",
	DMB:       "DMB",
	DRPS:      "DRPS",
	DSB:       "DSB",
	DUP:       "DUP",
	EON:       "EON",
	EOR:       "EOR",
	ERET:      "ERET",
	EXT:       "EXT",
	EXTR:      "EXTR",
	FABD:      "FABD",
	FABS:      "FABS",
	FACGE:     "FACGE",
	FACGT:     "FACGT",
	FADD:      "FADD",
	FADDP:     "FADDP",
	FCCMP:     "FCCMP",
	FCCMPE:    "FCCMPE",
	FCMEQ:     "FCMEQ",
	FCMGE:     "FCMGE",
	FCMGT:     "FCMGT",
	FCMLE:     "FCMLE",
	FCMLT:     "FCMLT",
	FCMP:      "FCMP",
	FCMPE:     "FCMPE",
	FCSEL:     "FCSEL",
	FCVT:      "FCVT",
	FCVTAS:    "FCVTAS",
	FCVTAU:    "FCVTAU",
	FCVTL:     "FCVTL",
	FCVTL2:    "FCVTL2",
	FCVTMS:    "FCVTMS",
	FCVTMU:    "FCVTMU",
	FCVTN:     "FCVTN",
	FCVTN2:    "FCVTN2",
	FCVTNS:    "FCVTNS",
	FCVTNU:    "FCVTNU",
	FCVTPS:    "FCVTPS",
	FCVTPU:    "FCVTPU",
	FCVTXN:    "FCVTXN",
	FCVTXN2:   "FCVTXN2",
	FCVTZS:    "FCVTZS",
	FCVTZU:    "FCVTZU",
	FDIV:      "FDIV",
	FMADD:     "FMADD",
	FMAX:      "FMAX",
	FMAXNM:    "FMAXNM",
	FMAXNMP:   "FMAXNMP",
	FMAXNMV:   "FMAXNMV",
	FMAXP:     "FMAXP",
	FMAXV:     "FMAXV",
	FMIN:      "FMIN",
	FMINNM:    "FMINNM",
	FMINNMP:   "FMINNMP",
	FMINNMV:   "FMINNMV",
	FMINP:     "FMINP",
	FMINV:     "FMINV",
	FMLA:      "FMLA",
	FMLS:      "FMLS",
	FMOV:      "FMOV",
	FMSUB:     "FMSUB",
	FMUL:      "FMUL",
	FMULX:     "FMULX",
	FNEG:      "FNEG",
	FNMADD:    "FNMADD",
	FNMSUB:    "FNMSUB",
	FNMUL:     "FNMUL",
	FRECPE:    "FRECPE",
	FRECPS:    "FRECPS",
	FRECPX:    "FRECPX",
	FRINTA:    "FRINTA",
	FRINTI:    "FRINTI",
	FRINTM:    "FRINTM",
	FRINTN:    "FRINTN",
	FRINTP:    "FRINTP",
	FRINTX:    "FRINTX",
	FRINTZ:    "FRINTZ",
	FRSQRTE:   "FRSQRTE",
	FRSQRTS:   "FRSQRTS",
	FSQRT:     "FSQRT",
	FSUB:      "FSUB",
	HINT:      "HINT",
	HLT:       "HLT",
	HVC:       "HVC",
	IC:        "IC",
	INS:       "INS",
	ISB:       "ISB",
	LD1:       "LD1",
	LD1R:      "LD1R",
	LD2:       "LD2",
	LD2R:      "LD2R",
	LD3:       "LD3",
	LD3R:      "LD3R",
	LD4:       "LD4",
	LD4R:      "LD4R",
	LDAR:      "LDAR",
	LDARB:     "LDARB",
	LDARH:     "LDARH",
	LDAXP:     "LDAXP",
	LDAXR:     "LDAXR",
	LDAXRB:    "LDAXRB",
	LDAXRH:    "LDAXRH",
	LDNP:      "LDNP",
	LDP:       "LDP",
	LDPSW:     "LDPSW",
	LDR:       "LDR",
	LDRB:      "LDRB",
	LDRH:      "LDRH",
	LDRSB:     "LDRSB",
	LDRSH:     "LDRSH",
	LDRSW:     "LDRSW",
	LDTR:      "LDTR",
	LDTRB:     "LDTRB",
	LDTRH:     "LDTRH",
	LDTRSB:    "LDTRSB",
	LDTRSH:    "LDTRSH",
	LDTRSW:    "LDTRSW",
	LDUR:      "LDUR",
	LDURB:     "LDURB",
	LDURH:     "LDURH",
	LDURSB:    "LDURSB",
	LDURSH:    "LDURSH",
	LDURSW:    "LDURSW",
	LDXP:      "LDXP",
	LDXR:      "LDXR",
	LDXRB:     "LDXRB",
	LDXRH:     "LDXRH",
	LSL:       "LSL",
	LSLV:      "LSLV",
	LSR:       "LSR",
	LSRV:      "LSRV",
	MADD:      "MADD",
	MLA:       "MLA",
	MLS:       "MLS",
	MNEG:      "MNEG",
	MOV:       "MOV",
	MOVI:      "MOVI",
	MOVK:      "MOVK",
	MOVN:      "MOVN",
	MOVZ:      "MOVZ",
	MRS:       "MRS",
	MSR:       "MSR",
	MSUB:      "MSUB",
	MUL:       "MUL",
	MVN:       "MVN",
	MVNI:      "MVNI",
	NEG:       "NEG",
	NEGS:      "NEGS",
	NGC:       "NGC",
	NGCS:      "NGCS",
	NOP:       "NOP",
	NOT:       "NOT",
	ORN:       "ORN",
	ORR:       "ORR",
	PMUL:      "PMUL",
	PMULL:     "PMULL",
	PMULL2:    "PMULL2",
	PRFM:      "PRFM",
	PRFUM:     "PRFUM",
	RADDHN:    "RADDHN",
	RADDHN2:   "RADDHN2",
	RBIT:      "RBIT",
	RET:       "RET",
	REV:       "REV",
	REV16:     "REV16",
	REV32:     "REV32",
	REV64:     "REV64",
	ROR:       "ROR",
	RORV:      "RORV",
	RSHRN:     "RSHRN",
	RSHRN2:    "RSHRN2",
	RSUBHN:    "RSUBHN",
	RSUBHN2:   "RSUBHN2",
	SABA:      "SABA",
	SABAL:     "SABAL",
	SABAL2:    "SABAL2",
	SABD:      "SABD",
	SABDL:     "SABDL",
	SABDL2:    "SABDL2",
	SADALP:    "SADALP",
	SADDL:     "SADDL",
	SADDL2:    "SADDL2",
	SADDLP:    "SADDLP",
	SADDLV:    "SADDLV",
	SADDW:     "SADDW",
	SADDW2:    "SADDW2",
	SBC:       "SBC",
	SBCS:      "SBCS",
	SBFIZ:     "SBFIZ",
	SBFM:      "SBFM",
	SBFX:      "SBFX",
	SCVTF:     "SCVTF",
	SDIV:      "SDIV",
	SEV:       "SEV",
	SEVL:      "SEVL",
	SHA1C:     "SHA1C",
	SHA1H:     "SHA1H",
	SHA1M:     "SHA1M",
	SHA1P:     "SHA1P",
	SHA1SU0:   "SHA1SU0",
	SHA1SU1:   "SHA1SU1",
	SHA256H:   "SHA256H",
	SHA256H2:  "SHA256H2",
	SHA256SU0: "SHA256SU0",
	SHA256SU1: "SHA256SU1",
	SHADD:     "SHADD",
	SHL:       "SHL",
	SHLL:      "SHLL",
	SHLL2:     "SHLL2",
	SHRN:      "SHRN",
	SHRN2:     "SHRN2",
	SHSUB:     "SHSUB",
	SLI:       "SLI",
	SMADDL:    "SMADDL",
	SMAX:      "SMAX",
	SMAXP:     "SMAXP",
	SMAXV:     "SMAXV",
	SMC:       "SMC",
	SMIN:      "SMIN",
	SMINP:     "SMINP",
	SMINV:     "SMINV",
	SMLAL:     "SMLAL",
	SMLAL2:    "SMLAL2",
	SMLSL:     "SMLSL",
	SMLSL2:    "SMLSL2",
	SMNEGL:    "SMNEGL",
	SMOV:      "SMOV",
	SMSUBL:    "SMSUBL",
	SMULH:     "SMULH",
	SMULL:     "SMULL",
	SMULL2:    "SMULL2",
	SQABS:     "SQABS",
	SQADD:     "SQADD",
	SQDMLAL:   "SQDMLAL",
	SQDMLAL2:  "SQDMLAL2",
	SQDMLSL:   "SQDMLSL",
	SQDMLSL2:  "SQDMLSL2",
	SQDMULH:   "SQDMULH",
	SQDMULL:   "SQDMULL",
	SQDMULL2:  "SQDMULL2",
	SQNEG:     "SQNEG",
	SQRDMULH:  "SQRDMULH",
	SQRSHL:    "SQRSHL",
	SQRSHRN:   "SQRSHRN",
	SQRSHRN2:  "SQRSHRN2",
	SQRSHRUN:  "SQRSHRUN",
	SQRSHRUN2: "SQRSHRUN2",
	SQSHL:     "SQSHL",
	SQSHLU:    "SQSHLU",
	SQSHRN:    "SQSHRN",
	SQSHRN2:   "SQSHRN2",
	SQSHRUN:   "SQSHRUN",
	SQSHRUN2:  "SQSHRUN2",
	SQSUB:     "SQSUB",
	SQXTN:     "SQXTN",
	SQXTN2:    "SQXTN2",
	SQXTUN:    "SQXTUN",
	SQXTUN2:   "SQXTUN2",
	SRHADD:    "SRHADD",
	SRI:       "SRI",
	SRSHL:     "SRSHL",
	SRSHR:     "SRSHR",
	SRSRA:     "SRSRA",
	SSHL:      "SSHL",
	SSHLL:     "SSHLL",
	SSHLL2:    "SSHLL2",
	SSHR:      "SSHR",
	SSRA:      "SSRA",
	SSUBL:     "SSUBL",
	SSUBL2:    "SSUBL2",
	SSUBW:     "SSUBW",
	SSUBW2:    "SSUBW2",
	ST1:       "ST1",
	ST2:       "ST2",
	ST3:       "ST3",
	ST4:       "ST4",
	STLR:      "STLR",
	STLRB:     "STLRB",
	STLRH:     "STLRH",
	STLXP:     "STLXP",
	STLXR:     "STLXR",
	STLXRB:    "STLXRB",
	STLXRH:    "STLXRH",
	STNP:      "STNP",
	STP:       "STP",
	STR:       "STR",
	STRB:      "STRB",
	STRH:      "STRH",
	STTR:      "STTR",
	STTRB:     "STTRB",
	STTRH:     "STTRH",
	STUR:      "STUR",
	STURB:     "STURB",
	STURH:     "STURH",
	STXP:      "STXP",
	STXR:      "STXR",
	STXRB:     "STXRB",
	STXRH:     "STXRH",
	SUB:       "SUB",
	SUBHN:     "SUBHN",
	SUBHN2:    "SUBHN2",
	SUBS:      "SUBS",
	SUQADD:    "SUQADD",
	SVC:       "SVC",
	SXTB:      "SXTB",
	SXTH:      "SXTH",
	SXTL:      "SXTL",
	SXTL2:     "SXTL2",
	SXTW:      "SXTW",
	SYS:       "SYS",
	SYSL:      "SYSL",
	TBL:       "TBL",
	TBNZ:      "TBNZ",
	TBX:       "TBX",
	TBZ:       "TBZ",
	TLBI:      "TLBI",
	TRN1:      "TRN1",
	TRN2:      "TRN2",
	TST:       "TST",
	UABA:      "UABA",
	UABAL:     "UABAL",
	UABAL2:    "UABAL2",
	UABD:      "UABD",
	UABDL:     "UABDL",
	UABDL2:    "UABDL2",
	UADALP:    "UADALP",
	UADDL:     "UADDL",
	UADDL2:    "UADDL2",
	UADDLP:    "UADDLP",
	UADDLV:    "UADDLV",
	UADDW:     "UADDW",
	UADDW2:    "UADDW2",
	UBFIZ:     "UBFIZ",
	UBFM:      "UBFM",
	UBFX:      "UBFX",
	UCVTF:     "UCVTF",
	UDIV:      "UDIV",
	UHADD:     "UHADD",
	UHSUB:     "UHSUB",
	UMADDL:    "UMADDL",
	UMAX:      "UMAX",
	UMAXP:     "UMAXP",
	UMAXV:     "UMAXV",
	UMIN:      "UMIN",
	UMINP:     "UMINP",
	UMINV:     "UMINV",
	UMLAL:     "UMLAL",
	UMLAL2:    "UMLAL2",
	UMLSL:     "UMLSL",
	UMLSL2:    "UMLSL2",
	UMNEGL:    "UMNEGL",
	UMOV:      "UMOV",
	UMSUBL:    "UMSUBL",
	UMULH:     "UMULH",
	UMULL:     "UMULL",
	UMULL2:    "UMULL2",
	UQADD:     "UQADD",
	UQRSHL:    "UQRSHL",
	UQRSHRN:   "UQRSHRN",
	UQRSHRN2:  "UQRSHRN2",
	UQSHL:     "UQSHL",
	UQSHRN:    "UQSHRN",
	UQSHRN2:   "UQSHRN2",
	UQSUB:     "UQSUB",
	UQXTN:     "UQXTN",
	UQXTN2:    "UQXTN2",
	URECPE:    "URECPE",
	URHADD:    "URHADD",
	URSHL:     "URSHL",
	URSHR:     "URSHR",
	URSQRTE:   "URSQRTE",
	URSRA:     "URSRA",
	USHL:      "USHL",
	USHLL:     "USHLL",
	USHLL2:    "USHLL2",
	USHR:      "USHR",
	USQADD:    "USQADD",
	USRA:      "USRA",
	USUBL:     "USUBL",
	USUBL2:    "USUBL2",
	USUBW:     "USUBW",
	USUBW2:    "USUBW2",
	UXTB:      "UXTB",
	UXTH:      "UXTH",
	UXTL:      "UXTL",
	UXTL2:     "UXTL2",
	UZP1:      "UZP1",
	UZP2:      "UZP2",
	WFE:       "WFE",
	WFI:       "WFI",
	XTN:       "XTN",
	XTN2:      "XTN2",
	YIELD:     "YIELD",
	ZIP1:      "ZIP1",
	ZIP2:      "ZIP2",
}

var instFormats = [...]instFormat{

	{0xffe0fc00, 0x1a000000, ADC, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9a000000, ADC, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x3a000000, ADCS, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0xba000000, ADCS, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe00000, 0x0b200000, ADD, instArgs{arg_Wds, arg_Wns, arg_Wm_extend__UXTB_0__UXTH_1__LSL_UXTW_2__UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xffe00000, 0x8b200000, ADD, instArgs{arg_Xds, arg_Xns, arg_Rm_extend__UXTB_0__UXTH_1__UXTW_2__LSL_UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xfffffc00, 0x11000000, MOV, instArgs{arg_Wds, arg_Wns}, mov_add_32_addsub_imm_cond},

	{0xff000000, 0x11000000, ADD, instArgs{arg_Wds, arg_Wns, arg_IAddSub}, nil},

	{0xfffffc00, 0x91000000, MOV, instArgs{arg_Xds, arg_Xns}, mov_add_64_addsub_imm_cond},

	{0xff000000, 0x91000000, ADD, instArgs{arg_Xds, arg_Xns, arg_IAddSub}, nil},

	{0xff208000, 0x0b000000, ADD, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__0_31}, nil},

	{0xff200000, 0x8b000000, ADD, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__0_63}, nil},

	{0xffe0001f, 0x2b20001f, CMN, instArgs{arg_Wns, arg_Wm_extend__UXTB_0__UXTH_1__LSL_UXTW_2__UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xffe00000, 0x2b200000, ADDS, instArgs{arg_Wd, arg_Wns, arg_Wm_extend__UXTB_0__UXTH_1__LSL_UXTW_2__UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xffe0001f, 0xab20001f, CMN, instArgs{arg_Xns, arg_Rm_extend__UXTB_0__UXTH_1__UXTW_2__LSL_UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xffe00000, 0xab200000, ADDS, instArgs{arg_Xd, arg_Xns, arg_Rm_extend__UXTB_0__UXTH_1__UXTW_2__LSL_UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xff00001f, 0x3100001f, CMN, instArgs{arg_Wns, arg_IAddSub}, nil},

	{0xff000000, 0x31000000, ADDS, instArgs{arg_Wd, arg_Wns, arg_IAddSub}, nil},

	{0xff00001f, 0xb100001f, CMN, instArgs{arg_Xns, arg_IAddSub}, nil},

	{0xff000000, 0xb1000000, ADDS, instArgs{arg_Xd, arg_Xns, arg_IAddSub}, nil},

	{0xff20801f, 0x2b00001f, CMN, instArgs{arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__0_31}, nil},

	{0xff208000, 0x2b000000, ADDS, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__0_31}, nil},

	{0xff20001f, 0xab00001f, CMN, instArgs{arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__0_63}, nil},

	{0xff200000, 0xab000000, ADDS, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__0_63}, nil},

	{0x9f000000, 0x10000000, ADR, instArgs{arg_Xd, arg_slabel_immhi_immlo_0}, nil},

	{0x9f000000, 0x90000000, ADRP, instArgs{arg_Xd, arg_slabel_immhi_immlo_12}, nil},

	{0xffc00000, 0x12000000, AND, instArgs{arg_Wds, arg_Wn, arg_immediate_bitmask_32_imms_immr}, nil},

	{0xff800000, 0x92000000, AND, instArgs{arg_Xds, arg_Xn, arg_immediate_bitmask_64_N_imms_immr}, nil},

	{0xff208000, 0x0a000000, AND, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff200000, 0x8a000000, AND, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xffc0001f, 0x7200001f, TST, instArgs{arg_Wn, arg_immediate_bitmask_32_imms_immr}, nil},

	{0xffc00000, 0x72000000, ANDS, instArgs{arg_Wd, arg_Wn, arg_immediate_bitmask_32_imms_immr}, nil},

	{0xff80001f, 0xf200001f, TST, instArgs{arg_Xn, arg_immediate_bitmask_64_N_imms_immr}, nil},

	{0xff800000, 0xf2000000, ANDS, instArgs{arg_Xd, arg_Xn, arg_immediate_bitmask_64_N_imms_immr}, nil},

	{0xff20801f, 0x6a00001f, TST, instArgs{arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff208000, 0x6a000000, ANDS, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff20001f, 0xea00001f, TST, instArgs{arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xff200000, 0xea000000, ANDS, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xffc0fc00, 0x13007c00, ASR, instArgs{arg_Wd, arg_Wn, arg_immediate_ASR_SBFM_32M_bitfield_0_31_immr}, nil},

	{0xffc00000, 0x13000000, SBFIZ, instArgs{arg_Wd, arg_Wn, arg_immediate_SBFIZ_SBFM_32M_bitfield_lsb_32_immr, arg_immediate_SBFIZ_SBFM_32M_bitfield_width_32_imms}, sbfiz_sbfm_32m_bitfield_cond},

	{0xffc00000, 0x13000000, SBFX, instArgs{arg_Wd, arg_Wn, arg_immediate_SBFX_SBFM_32M_bitfield_lsb_32_immr, arg_immediate_SBFX_SBFM_32M_bitfield_width_32_imms}, sbfx_sbfm_32m_bitfield_cond},

	{0xfffffc00, 0x13001c00, SXTB, instArgs{arg_Wd, arg_Wn}, nil},

	{0xfffffc00, 0x13003c00, SXTH, instArgs{arg_Wd, arg_Wn}, nil},

	{0xffc00000, 0x13000000, SBFM, instArgs{arg_Wd, arg_Wn, arg_immediate_0_31_immr, arg_immediate_0_31_imms}, nil},

	{0xffc0fc00, 0x9340fc00, ASR, instArgs{arg_Xd, arg_Xn, arg_immediate_ASR_SBFM_64M_bitfield_0_63_immr}, nil},

	{0xffc00000, 0x93400000, SBFIZ, instArgs{arg_Xd, arg_Xn, arg_immediate_SBFIZ_SBFM_64M_bitfield_lsb_64_immr, arg_immediate_SBFIZ_SBFM_64M_bitfield_width_64_imms}, sbfiz_sbfm_64m_bitfield_cond},

	{0xffc00000, 0x93400000, SBFX, instArgs{arg_Xd, arg_Xn, arg_immediate_SBFX_SBFM_64M_bitfield_lsb_64_immr, arg_immediate_SBFX_SBFM_64M_bitfield_width_64_imms}, sbfx_sbfm_64m_bitfield_cond},

	{0xfffffc00, 0x93401c00, SXTB, instArgs{arg_Xd, arg_Wn}, nil},

	{0xfffffc00, 0x93403c00, SXTH, instArgs{arg_Xd, arg_Wn}, nil},

	{0xfffffc00, 0x93407c00, SXTW, instArgs{arg_Xd, arg_Wn}, nil},

	{0xffc00000, 0x93400000, SBFM, instArgs{arg_Xd, arg_Xn, arg_immediate_0_63_immr, arg_immediate_0_63_imms}, nil},

	{0xffe0fc00, 0x1ac02800, ASR, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x1ac02800, ASRV, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9ac02800, ASR, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x9ac02800, ASRV, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xfff8ff00, 0xd5087800, AT, instArgs{arg_sysop_AT_SYS_CR_system}, at_sys_cr_system_cond},

	{0xfff8f000, 0xd5087000, DC, instArgs{arg_sysop_DC_SYS_CR_system}, dc_sys_cr_system_cond},

	{0xfff8f000, 0xd5087000, IC, instArgs{arg_sysop_IC_SYS_CR_system}, ic_sys_cr_system_cond},

	{0xfff8f000, 0xd5088000, TLBI, instArgs{arg_sysop_TLBI_SYS_CR_system}, tlbi_sys_cr_system_cond},

	{0xfff80000, 0xd5080000, SYS, instArgs{arg_immediate_0_7_op1, arg_Cn, arg_Cm, arg_sysop_SYS_CR_system}, nil},

	{0xfc000000, 0x14000000, B, instArgs{arg_slabel_imm26_2}, nil},

	{0xff000010, 0x54000000, B, instArgs{arg_conditional, arg_slabel_imm19_2}, nil},

	{0xffc00000, 0x33000000, BFI, instArgs{arg_Wd, arg_Wn, arg_immediate_BFI_BFM_32M_bitfield_lsb_32_immr, arg_immediate_BFI_BFM_32M_bitfield_width_32_imms}, bfi_bfm_32m_bitfield_cond},

	{0xffc00000, 0x33000000, BFXIL, instArgs{arg_Wd, arg_Wn, arg_immediate_BFXIL_BFM_32M_bitfield_lsb_32_immr, arg_immediate_BFXIL_BFM_32M_bitfield_width_32_imms}, bfxil_bfm_32m_bitfield_cond},

	{0xffc00000, 0x33000000, BFM, instArgs{arg_Wd, arg_Wn, arg_immediate_0_31_immr, arg_immediate_0_31_imms}, nil},

	{0xffc00000, 0xb3400000, BFI, instArgs{arg_Xd, arg_Xn, arg_immediate_BFI_BFM_64M_bitfield_lsb_64_immr, arg_immediate_BFI_BFM_64M_bitfield_width_64_imms}, bfi_bfm_64m_bitfield_cond},

	{0xffc00000, 0xb3400000, BFXIL, instArgs{arg_Xd, arg_Xn, arg_immediate_BFXIL_BFM_64M_bitfield_lsb_64_immr, arg_immediate_BFXIL_BFM_64M_bitfield_width_64_imms}, bfxil_bfm_64m_bitfield_cond},

	{0xffc00000, 0xb3400000, BFM, instArgs{arg_Xd, arg_Xn, arg_immediate_0_63_immr, arg_immediate_0_63_imms}, nil},

	{0xff208000, 0x0a200000, BIC, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff200000, 0x8a200000, BIC, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xff208000, 0x6a200000, BICS, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff200000, 0xea200000, BICS, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xfc000000, 0x94000000, BL, instArgs{arg_slabel_imm26_2}, nil},

	{0xfffffc1f, 0xd63f0000, BLR, instArgs{arg_Xn}, nil},

	{0xfffffc1f, 0xd61f0000, BR, instArgs{arg_Xn}, nil},

	{0xffe0001f, 0xd4200000, BRK, instArgs{arg_immediate_0_65535_imm16}, nil},

	{0xff000000, 0x35000000, CBNZ, instArgs{arg_Wt, arg_slabel_imm19_2}, nil},

	{0xff000000, 0xb5000000, CBNZ, instArgs{arg_Xt, arg_slabel_imm19_2}, nil},

	{0xff000000, 0x34000000, CBZ, instArgs{arg_Wt, arg_slabel_imm19_2}, nil},

	{0xff000000, 0xb4000000, CBZ, instArgs{arg_Xt, arg_slabel_imm19_2}, nil},

	{0xffe00c10, 0x3a400800, CCMN, instArgs{arg_Wn, arg_immediate_0_31_imm5, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0xba400800, CCMN, instArgs{arg_Xn, arg_immediate_0_31_imm5, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0x3a400000, CCMN, instArgs{arg_Wn, arg_Wm, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0xba400000, CCMN, instArgs{arg_Xn, arg_Xm, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0x7a400800, CCMP, instArgs{arg_Wn, arg_immediate_0_31_imm5, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0xfa400800, CCMP, instArgs{arg_Xn, arg_immediate_0_31_imm5, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0x7a400000, CCMP, instArgs{arg_Wn, arg_Wm, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0xfa400000, CCMP, instArgs{arg_Xn, arg_Xm, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c00, 0x1a800400, CINC, instArgs{arg_Wd, arg_Wn, arg_cond_NotAllowALNV_Invert}, cinc_csinc_32_condsel_cond},

	{0xffff0fe0, 0x1a9f07e0, CSET, instArgs{arg_Wd, arg_cond_NotAllowALNV_Invert}, csinc_general_cond},

	{0xffe00c00, 0x1a800400, CSINC, instArgs{arg_Wd, arg_Wn, arg_Wm, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c00, 0x9a800400, CINC, instArgs{arg_Xd, arg_Xn, arg_cond_NotAllowALNV_Invert}, cinc_csinc_64_condsel_cond},

	{0xffff0fe0, 0x9a9f07e0, CSET, instArgs{arg_Xd, arg_cond_NotAllowALNV_Invert}, csinc_general_cond},

	{0xffe00c00, 0x9a800400, CSINC, instArgs{arg_Xd, arg_Xn, arg_Xm, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c00, 0x5a800000, CINV, instArgs{arg_Wd, arg_Wn, arg_cond_NotAllowALNV_Invert}, cinv_csinv_32_condsel_cond},

	{0xffff0fe0, 0x5a9f03e0, CSETM, instArgs{arg_Wd, arg_cond_NotAllowALNV_Invert}, csinv_general_cond},

	{0xffe00c00, 0x5a800000, CSINV, instArgs{arg_Wd, arg_Wn, arg_Wm, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c00, 0xda800000, CINV, instArgs{arg_Xd, arg_Xn, arg_cond_NotAllowALNV_Invert}, cinv_csinv_64_condsel_cond},

	{0xffff0fe0, 0xda9f03e0, CSETM, instArgs{arg_Xd, arg_cond_NotAllowALNV_Invert}, csinv_general_cond},

	{0xffe00c00, 0xda800000, CSINV, instArgs{arg_Xd, arg_Xn, arg_Xm, arg_cond_AllowALNV_Normal}, nil},

	{0xfffff0ff, 0xd503305f, CLREX, instArgs{arg_immediate_optional_0_15_CRm}, nil},

	{0xfffffc00, 0x5ac01400, CLS, instArgs{arg_Wd, arg_Wn}, nil},

	{0xfffffc00, 0xdac01400, CLS, instArgs{arg_Xd, arg_Xn}, nil},

	{0xfffffc00, 0x5ac01000, CLZ, instArgs{arg_Wd, arg_Wn}, nil},

	{0xfffffc00, 0xdac01000, CLZ, instArgs{arg_Xd, arg_Xn}, nil},

	{0xffe0001f, 0x6b20001f, CMP, instArgs{arg_Wns, arg_Wm_extend__UXTB_0__UXTH_1__LSL_UXTW_2__UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xffe00000, 0x6b200000, SUBS, instArgs{arg_Wd, arg_Wns, arg_Wm_extend__UXTB_0__UXTH_1__LSL_UXTW_2__UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xffe0001f, 0xeb20001f, CMP, instArgs{arg_Xns, arg_Rm_extend__UXTB_0__UXTH_1__UXTW_2__LSL_UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xffe00000, 0xeb200000, SUBS, instArgs{arg_Xd, arg_Xns, arg_Rm_extend__UXTB_0__UXTH_1__UXTW_2__LSL_UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xff00001f, 0x7100001f, CMP, instArgs{arg_Wns, arg_IAddSub}, nil},

	{0xff000000, 0x71000000, SUBS, instArgs{arg_Wd, arg_Wns, arg_IAddSub}, nil},

	{0xff00001f, 0xf100001f, CMP, instArgs{arg_Xns, arg_IAddSub}, nil},

	{0xff000000, 0xf1000000, SUBS, instArgs{arg_Xd, arg_Xns, arg_IAddSub}, nil},

	{0xff20801f, 0x6b00001f, CMP, instArgs{arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__0_31}, nil},

	{0xff2003e0, 0x6b0003e0, NEGS, instArgs{arg_Wd, arg_Wm_shift__LSL_0__LSR_1__ASR_2__0_31}, nil},

	{0xff208000, 0x6b000000, SUBS, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__0_31}, nil},

	{0xff20001f, 0xeb00001f, CMP, instArgs{arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__0_63}, nil},

	{0xff2003e0, 0xeb0003e0, NEGS, instArgs{arg_Xd, arg_Xm_shift__LSL_0__LSR_1__ASR_2__0_63}, nil},

	{0xff200000, 0xeb000000, SUBS, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__0_63}, nil},

	{0xffe00c00, 0x5a800400, CNEG, instArgs{arg_Wd, arg_Wn, arg_cond_NotAllowALNV_Invert}, cneg_csneg_32_condsel_cond},

	{0xffe00c00, 0x5a800400, CSNEG, instArgs{arg_Wd, arg_Wn, arg_Wm, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c00, 0xda800400, CNEG, instArgs{arg_Xd, arg_Xn, arg_cond_NotAllowALNV_Invert}, cneg_csneg_64_condsel_cond},

	{0xffe00c00, 0xda800400, CSNEG, instArgs{arg_Xd, arg_Xn, arg_Xm, arg_cond_AllowALNV_Normal}, nil},

	{0xffe0fc00, 0x1ac04000, CRC32B, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x1ac04400, CRC32H, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x1ac04800, CRC32W, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9ac04c00, CRC32X, instArgs{arg_Wd, arg_Wn, arg_Xm}, nil},

	{0xffe0fc00, 0x1ac05000, CRC32CB, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x1ac05400, CRC32CH, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x1ac05800, CRC32CW, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9ac05c00, CRC32CX, instArgs{arg_Wd, arg_Wn, arg_Xm}, nil},

	{0xffe00c00, 0x1a800000, CSEL, instArgs{arg_Wd, arg_Wn, arg_Wm, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c00, 0x9a800000, CSEL, instArgs{arg_Xd, arg_Xn, arg_Xm, arg_cond_AllowALNV_Normal}, nil},

	{0xffe0001f, 0xd4a00001, DCPS1, instArgs{arg_immediate_optional_0_65535_imm16}, nil},

	{0xffe0001f, 0xd4a00002, DCPS2, instArgs{arg_immediate_optional_0_65535_imm16}, nil},

	{0xffe0001f, 0xd4a00003, DCPS3, instArgs{arg_immediate_optional_0_65535_imm16}, nil},

	{0xfffff0ff, 0xd50330bf, DMB, instArgs{arg_option_DMB_BO_system_CRm}, nil},

	{0xffffffff, 0xd6bf03e0, DRPS, instArgs{}, nil},

	{0xfffff0ff, 0xd503309f, DSB, instArgs{arg_option_DSB_BO_system_CRm}, nil},

	{0xff208000, 0x4a200000, EON, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff200000, 0xca200000, EON, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xffc00000, 0x52000000, EOR, instArgs{arg_Wds, arg_Wn, arg_immediate_bitmask_32_imms_immr}, nil},

	{0xff800000, 0xd2000000, EOR, instArgs{arg_Xds, arg_Xn, arg_immediate_bitmask_64_N_imms_immr}, nil},

	{0xff208000, 0x4a000000, EOR, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff200000, 0xca000000, EOR, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xffffffff, 0xd69f03e0, ERET, instArgs{}, nil},

	{0xffe08000, 0x13800000, ROR, instArgs{arg_Wd, arg_Ws, arg_immediate_0_31_imms}, ror_extr_32_extract_cond},

	{0xffe08000, 0x13800000, EXTR, instArgs{arg_Wd, arg_Wn, arg_Wm, arg_immediate_0_31_imms}, nil},

	{0xffe00000, 0x93c00000, ROR, instArgs{arg_Xd, arg_Xs, arg_immediate_0_63_imms}, ror_extr_64_extract_cond},

	{0xffe00000, 0x93c00000, EXTR, instArgs{arg_Xd, arg_Xn, arg_Xm, arg_immediate_0_63_imms}, nil},

	{0xffffffff, 0xd503201f, NOP, instArgs{}, nil},

	{0xffffffff, 0xd503209f, SEV, instArgs{}, nil},

	{0xffffffff, 0xd50320bf, SEVL, instArgs{}, nil},

	{0xffffffff, 0xd503205f, WFE, instArgs{}, nil},

	{0xffffffff, 0xd503207f, WFI, instArgs{}, nil},

	{0xffffffff, 0xd503203f, YIELD, instArgs{}, nil},

	{0xfffff01f, 0xd503201f, HINT, instArgs{arg_immediate_0_127_CRm_op2}, nil},

	{0xffe0001f, 0xd4400000, HLT, instArgs{arg_immediate_0_65535_imm16}, nil},

	{0xfffff0ff, 0xd50330df, ISB, instArgs{arg_option_ISB_BI_system_CRm}, nil},

	{0xffe08000, 0x88c08000, LDAR, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8c08000, LDAR, instArgs{arg_Xt, arg_Xns_mem}, nil},

	{0xffe08000, 0x08c08000, LDARB, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0x48c08000, LDARH, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0x88608000, LDAXP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8608000, LDAXP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem}, nil},

	{0xffe08000, 0x88408000, LDAXR, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8408000, LDAXR, instArgs{arg_Xt, arg_Xns_mem}, nil},

	{0xffe08000, 0x08408000, LDAXRB, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0x48408000, LDAXRH, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffc00000, 0x28400000, LDNP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffc00000, 0xa8400000, LDNP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_optional_imm7_8_signed}, nil},

	{0xffc00000, 0x28c00000, LDP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem_post_imm7_4_signed}, nil},

	{0xffc00000, 0xa8c00000, LDP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_post_imm7_8_signed}, nil},

	{0xffc00000, 0x29c00000, LDP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem_wb_imm7_4_signed}, nil},

	{0xffc00000, 0xa9c00000, LDP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_wb_imm7_8_signed}, nil},

	{0xffc00000, 0x29400000, LDP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffc00000, 0xa9400000, LDP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_optional_imm7_8_signed}, nil},

	{0xffc00000, 0x68c00000, LDPSW, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_post_imm7_4_signed}, nil},

	{0xffc00000, 0x69c00000, LDPSW, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_wb_imm7_4_signed}, nil},

	{0xffc00000, 0x69400000, LDPSW, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffe00c00, 0xb8400400, LDR, instArgs{arg_Wt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xf8400400, LDR, instArgs{arg_Xt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xb8400c00, LDR, instArgs{arg_Wt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0xf8400c00, LDR, instArgs{arg_Xt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0xb9400000, LDR, instArgs{arg_Wt, arg_Xns_mem_optional_imm12_4_unsigned}, nil},

	{0xffc00000, 0xf9400000, LDR, instArgs{arg_Xt, arg_Xns_mem_optional_imm12_8_unsigned}, nil},

	{0xff000000, 0x18000000, LDR, instArgs{arg_Wt, arg_slabel_imm19_2}, nil},

	{0xff000000, 0x58000000, LDR, instArgs{arg_Xt, arg_slabel_imm19_2}, nil},

	{0xffe00c00, 0xb8600800, LDR, instArgs{arg_Wt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__2_1}, nil},

	{0xffe00c00, 0xf8600800, LDR, instArgs{arg_Xt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__3_1}, nil},

	{0xffe00c00, 0x38400400, LDRB, instArgs{arg_Wt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x38400c00, LDRB, instArgs{arg_Wt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0x39400000, LDRB, instArgs{arg_Wt, arg_Xns_mem_optional_imm12_1_unsigned}, nil},

	{0xffe00c00, 0x38600800, LDRB, instArgs{arg_Wt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__absent_0__0_1}, nil},

	{0xffe00c00, 0x78400400, LDRH, instArgs{arg_Wt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x78400c00, LDRH, instArgs{arg_Wt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0x79400000, LDRH, instArgs{arg_Wt, arg_Xns_mem_optional_imm12_2_unsigned}, nil},

	{0xffe00c00, 0x78600800, LDRH, instArgs{arg_Wt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__1_1}, nil},

	{0xffe00c00, 0x38c00400, LDRSB, instArgs{arg_Wt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x38800400, LDRSB, instArgs{arg_Xt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x38c00c00, LDRSB, instArgs{arg_Wt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0x38800c00, LDRSB, instArgs{arg_Xt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0x39c00000, LDRSB, instArgs{arg_Wt, arg_Xns_mem_optional_imm12_1_unsigned}, nil},

	{0xffc00000, 0x39800000, LDRSB, instArgs{arg_Xt, arg_Xns_mem_optional_imm12_1_unsigned}, nil},

	{0xffe00c00, 0x38e00800, LDRSB, instArgs{arg_Wt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__absent_0__0_1}, nil},

	{0xffe00c00, 0x38a00800, LDRSB, instArgs{arg_Xt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__absent_0__0_1}, nil},

	{0xffe00c00, 0x78c00400, LDRSH, instArgs{arg_Wt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x78800400, LDRSH, instArgs{arg_Xt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x78c00c00, LDRSH, instArgs{arg_Wt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0x78800c00, LDRSH, instArgs{arg_Xt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0x79c00000, LDRSH, instArgs{arg_Wt, arg_Xns_mem_optional_imm12_2_unsigned}, nil},

	{0xffc00000, 0x79800000, LDRSH, instArgs{arg_Xt, arg_Xns_mem_optional_imm12_2_unsigned}, nil},

	{0xffe00c00, 0x78e00800, LDRSH, instArgs{arg_Wt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__1_1}, nil},

	{0xffe00c00, 0x78a00800, LDRSH, instArgs{arg_Xt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__1_1}, nil},

	{0xffe00c00, 0xb8800400, LDRSW, instArgs{arg_Xt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xb8800c00, LDRSW, instArgs{arg_Xt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0xb9800000, LDRSW, instArgs{arg_Xt, arg_Xns_mem_optional_imm12_4_unsigned}, nil},

	{0xff000000, 0x98000000, LDRSW, instArgs{arg_Xt, arg_slabel_imm19_2}, nil},

	{0xffe00c00, 0xb8a00800, LDRSW, instArgs{arg_Xt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__2_1}, nil},

	{0xffe00c00, 0xb8400800, LDTR, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xf8400800, LDTR, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x38400800, LDTRB, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x78400800, LDTRH, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x38c00800, LDTRSB, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x38800800, LDTRSB, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x78c00800, LDTRSH, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x78800800, LDTRSH, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xb8800800, LDTRSW, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xb8400000, LDUR, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xf8400000, LDUR, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x38400000, LDURB, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x78400000, LDURH, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x38c00000, LDURSB, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x38800000, LDURSB, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x78c00000, LDURSH, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x78800000, LDURSH, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xb8800000, LDURSW, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe08000, 0x88600000, LDXP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8600000, LDXP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem}, nil},

	{0xffe08000, 0x88400000, LDXR, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8400000, LDXR, instArgs{arg_Xt, arg_Xns_mem}, nil},

	{0xffe08000, 0x08400000, LDXRB, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0x48400000, LDXRH, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffc08000, 0x53000000, LSL, instArgs{arg_Wd, arg_Wn, arg_immediate_LSL_UBFM_32M_bitfield_0_31_immr}, lsl_ubfm_32m_bitfield_cond},

	{0xffc0fc00, 0x53007c00, LSR, instArgs{arg_Wd, arg_Wn, arg_immediate_LSR_UBFM_32M_bitfield_0_31_immr}, nil},

	{0xffc00000, 0x53000000, UBFIZ, instArgs{arg_Wd, arg_Wn, arg_immediate_UBFIZ_UBFM_32M_bitfield_lsb_32_immr, arg_immediate_UBFIZ_UBFM_32M_bitfield_width_32_imms}, ubfiz_ubfm_32m_bitfield_cond},

	{0xffc00000, 0x53000000, UBFX, instArgs{arg_Wd, arg_Wn, arg_immediate_UBFX_UBFM_32M_bitfield_lsb_32_immr, arg_immediate_UBFX_UBFM_32M_bitfield_width_32_imms}, ubfx_ubfm_32m_bitfield_cond},

	{0xfffffc00, 0x53001c00, UXTB, instArgs{arg_Wd, arg_Wn}, nil},

	{0xfffffc00, 0x53003c00, UXTH, instArgs{arg_Wd, arg_Wn}, nil},

	{0xffc00000, 0x53000000, UBFM, instArgs{arg_Wd, arg_Wn, arg_immediate_0_31_immr, arg_immediate_0_31_imms}, nil},

	{0xffc00000, 0xd3400000, LSL, instArgs{arg_Xd, arg_Xn, arg_immediate_LSL_UBFM_64M_bitfield_0_63_immr}, lsl_ubfm_64m_bitfield_cond},

	{0xffc0fc00, 0xd340fc00, LSR, instArgs{arg_Xd, arg_Xn, arg_immediate_LSR_UBFM_64M_bitfield_0_63_immr}, nil},

	{0xffc00000, 0xd3400000, UBFIZ, instArgs{arg_Xd, arg_Xn, arg_immediate_UBFIZ_UBFM_64M_bitfield_lsb_64_immr, arg_immediate_UBFIZ_UBFM_64M_bitfield_width_64_imms}, ubfiz_ubfm_64m_bitfield_cond},

	{0xffc00000, 0xd3400000, UBFX, instArgs{arg_Xd, arg_Xn, arg_immediate_UBFX_UBFM_64M_bitfield_lsb_64_immr, arg_immediate_UBFX_UBFM_64M_bitfield_width_64_imms}, ubfx_ubfm_64m_bitfield_cond},

	{0xffc00000, 0xd3400000, UBFM, instArgs{arg_Xd, arg_Xn, arg_immediate_0_63_immr, arg_immediate_0_63_imms}, nil},

	{0xffe0fc00, 0x1ac02000, LSL, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x1ac02000, LSLV, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9ac02000, LSL, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x9ac02000, LSLV, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x1ac02400, LSR, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x1ac02400, LSRV, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9ac02400, LSR, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x9ac02400, LSRV, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x1b007c00, MUL, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe08000, 0x1b000000, MADD, instArgs{arg_Wd, arg_Wn, arg_Wm, arg_Wa}, nil},

	{0xffe0fc00, 0x9b007c00, MUL, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe08000, 0x9b000000, MADD, instArgs{arg_Xd, arg_Xn, arg_Xm, arg_Xa}, nil},

	{0xffe0fc00, 0x1b00fc00, MNEG, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe08000, 0x1b008000, MSUB, instArgs{arg_Wd, arg_Wn, arg_Wm, arg_Wa}, nil},

	{0xffe0fc00, 0x9b00fc00, MNEG, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe08000, 0x9b008000, MSUB, instArgs{arg_Xd, arg_Xn, arg_Xm, arg_Xa}, nil},

	{0xffc003e0, 0x320003e0, MOV, instArgs{arg_Wds, arg_immediate_bitmask_32_imms_immr}, mov_orr_32_log_imm_cond},

	{0xffc00000, 0x32000000, ORR, instArgs{arg_Wds, arg_Wn, arg_immediate_bitmask_32_imms_immr}, nil},

	{0xff8003e0, 0xb20003e0, MOV, instArgs{arg_Xds, arg_immediate_bitmask_64_N_imms_immr}, mov_orr_64_log_imm_cond},

	{0xff800000, 0xb2000000, ORR, instArgs{arg_Xds, arg_Xn, arg_immediate_bitmask_64_N_imms_immr}, nil},

	{0xff800000, 0x12800000, MOV, instArgs{arg_Wd, arg_immediate_shift_32_implicit_inverse_imm16_hw}, mov_movn_32_movewide_cond},

	{0xff800000, 0x12800000, MOVN, instArgs{arg_Wd, arg_immediate_OptLSL_amount_16_0_16}, nil},

	{0xff800000, 0x92800000, MOV, instArgs{arg_Xd, arg_immediate_shift_64_implicit_inverse_imm16_hw}, mov_movn_64_movewide_cond},

	{0xff800000, 0x92800000, MOVN, instArgs{arg_Xd, arg_immediate_OptLSL_amount_16_0_48}, nil},

	{0xffe0ffe0, 0x2a0003e0, MOV, instArgs{arg_Wd, arg_Wm}, nil},

	{0xff208000, 0x2a000000, ORR, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xffe0ffe0, 0xaa0003e0, MOV, instArgs{arg_Xd, arg_Xm}, nil},

	{0xff200000, 0xaa000000, ORR, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xff800000, 0x52800000, MOV, instArgs{arg_Wd, arg_immediate_shift_32_implicit_imm16_hw}, mov_movz_32_movewide_cond},

	{0xff800000, 0x52800000, MOVZ, instArgs{arg_Wd, arg_immediate_OptLSL_amount_16_0_16}, nil},

	{0xff800000, 0xd2800000, MOV, instArgs{arg_Xd, arg_immediate_shift_64_implicit_imm16_hw}, mov_movz_64_movewide_cond},

	{0xff800000, 0xd2800000, MOVZ, instArgs{arg_Xd, arg_immediate_OptLSL_amount_16_0_48}, nil},

	{0xff800000, 0x72800000, MOVK, instArgs{arg_Wd, arg_immediate_OptLSL_amount_16_0_16}, nil},

	{0xff800000, 0xf2800000, MOVK, instArgs{arg_Xd, arg_immediate_OptLSL_amount_16_0_48}, nil},

	{0xfff00000, 0xd5300000, MRS, instArgs{arg_Xt, arg_sysreg_o0_op1_CRn_CRm_op2}, nil},

	{0xfff8f01f, 0xd500401f, MSR, instArgs{arg_pstatefield_op1_op2__SPSel_05__DAIFSet_36__DAIFClr_37, arg_immediate_0_15_CRm}, nil},

	{0xfff00000, 0xd5100000, MSR, instArgs{arg_sysreg_o0_op1_CRn_CRm_op2, arg_Xt}, nil},

	{0xff2003e0, 0x2a2003e0, MVN, instArgs{arg_Wd, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff208000, 0x2a200000, ORN, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_31}, nil},

	{0xff2003e0, 0xaa2003e0, MVN, instArgs{arg_Xd, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xff200000, 0xaa200000, ORN, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__ROR_3__0_63}, nil},

	{0xff2003e0, 0x4b0003e0, NEG, instArgs{arg_Wd, arg_Wm_shift__LSL_0__LSR_1__ASR_2__0_31}, nil},

	{0xff208000, 0x4b000000, SUB, instArgs{arg_Wd, arg_Wn, arg_Wm_shift__LSL_0__LSR_1__ASR_2__0_31}, nil},

	{0xff2003e0, 0xcb0003e0, NEG, instArgs{arg_Xd, arg_Xm_shift__LSL_0__LSR_1__ASR_2__0_63}, nil},

	{0xff200000, 0xcb000000, SUB, instArgs{arg_Xd, arg_Xn, arg_Xm_shift__LSL_0__LSR_1__ASR_2__0_63}, nil},

	{0xffe0ffe0, 0x5a0003e0, NGC, instArgs{arg_Wd, arg_Wm}, nil},

	{0xffe0fc00, 0x5a000000, SBC, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0ffe0, 0xda0003e0, NGC, instArgs{arg_Xd, arg_Xm}, nil},

	{0xffe0fc00, 0xda000000, SBC, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0ffe0, 0x7a0003e0, NGCS, instArgs{arg_Wd, arg_Wm}, nil},

	{0xffe0fc00, 0x7a000000, SBCS, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0ffe0, 0xfa0003e0, NGCS, instArgs{arg_Xd, arg_Xm}, nil},

	{0xffe0fc00, 0xfa000000, SBCS, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffc00000, 0xf9800000, PRFM, instArgs{arg_prfop_Rt, arg_Xns_mem_optional_imm12_8_unsigned}, nil},

	{0xff000000, 0xd8000000, PRFM, instArgs{arg_prfop_Rt, arg_slabel_imm19_2}, nil},

	{0xffe00c00, 0xf8a00800, PRFM, instArgs{arg_prfop_Rt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__3_1}, nil},

	{0xffe00c00, 0xf8800000, PRFUM, instArgs{arg_prfop_Rt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xfffffc00, 0x5ac00000, RBIT, instArgs{arg_Wd, arg_Wn}, nil},

	{0xfffffc00, 0xdac00000, RBIT, instArgs{arg_Xd, arg_Xn}, nil},

	{0xfffffc1f, 0xd65f0000, RET, instArgs{arg_Xn}, nil},

	{0xfffffc00, 0x5ac00800, REV, instArgs{arg_Wd, arg_Wn}, nil},

	{0xfffffc00, 0xdac00c00, REV, instArgs{arg_Xd, arg_Xn}, nil},

	{0xfffffc00, 0x5ac00400, REV16, instArgs{arg_Wd, arg_Wn}, nil},

	{0xfffffc00, 0xdac00400, REV16, instArgs{arg_Xd, arg_Xn}, nil},

	{0xfffffc00, 0xdac00800, REV32, instArgs{arg_Xd, arg_Xn}, nil},

	{0xffe0fc00, 0x1ac02c00, ROR, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x1ac02c00, RORV, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9ac02c00, ROR, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x9ac02c00, RORV, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x1ac00c00, SDIV, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9ac00c00, SDIV, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x9b207c00, SMULL, instArgs{arg_Xd, arg_Wn, arg_Wm}, nil},

	{0xffe08000, 0x9b200000, SMADDL, instArgs{arg_Xd, arg_Wn, arg_Wm, arg_Xa}, nil},

	{0xffe0fc00, 0x9b20fc00, SMNEGL, instArgs{arg_Xd, arg_Wn, arg_Wm}, nil},

	{0xffe08000, 0x9b208000, SMSUBL, instArgs{arg_Xd, arg_Wn, arg_Wm, arg_Xa}, nil},

	{0xffe08000, 0x9b400000, SMULH, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe08000, 0x88808000, STLR, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8808000, STLR, instArgs{arg_Xt, arg_Xns_mem}, nil},

	{0xffe08000, 0x08808000, STLRB, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0x48808000, STLRH, instArgs{arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0x88208000, STLXP, instArgs{arg_Ws, arg_Wt, arg_Wt2, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8208000, STLXP, instArgs{arg_Ws, arg_Xt, arg_Xt2, arg_Xns_mem}, nil},

	{0xffe08000, 0x88008000, STLXR, instArgs{arg_Ws, arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8008000, STLXR, instArgs{arg_Ws, arg_Xt, arg_Xns_mem}, nil},

	{0xffe08000, 0x08008000, STLXRB, instArgs{arg_Ws, arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0x48008000, STLXRH, instArgs{arg_Ws, arg_Wt, arg_Xns_mem}, nil},

	{0xffc00000, 0x28000000, STNP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffc00000, 0xa8000000, STNP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_optional_imm7_8_signed}, nil},

	{0xffc00000, 0x28800000, STP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem_post_imm7_4_signed}, nil},

	{0xffc00000, 0xa8800000, STP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_post_imm7_8_signed}, nil},

	{0xffc00000, 0x29800000, STP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem_wb_imm7_4_signed}, nil},

	{0xffc00000, 0xa9800000, STP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_wb_imm7_8_signed}, nil},

	{0xffc00000, 0x29000000, STP, instArgs{arg_Wt, arg_Wt2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffc00000, 0xa9000000, STP, instArgs{arg_Xt, arg_Xt2, arg_Xns_mem_optional_imm7_8_signed}, nil},

	{0xffe00c00, 0xb8000400, STR, instArgs{arg_Wt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xf8000400, STR, instArgs{arg_Xt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xb8000c00, STR, instArgs{arg_Wt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0xf8000c00, STR, instArgs{arg_Xt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0xb9000000, STR, instArgs{arg_Wt, arg_Xns_mem_optional_imm12_4_unsigned}, nil},

	{0xffc00000, 0xf9000000, STR, instArgs{arg_Xt, arg_Xns_mem_optional_imm12_8_unsigned}, nil},

	{0xffe00c00, 0xb8200800, STR, instArgs{arg_Wt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__2_1}, nil},

	{0xffe00c00, 0xf8200800, STR, instArgs{arg_Xt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__3_1}, nil},

	{0xffe00c00, 0x38000400, STRB, instArgs{arg_Wt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x38000c00, STRB, instArgs{arg_Wt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0x39000000, STRB, instArgs{arg_Wt, arg_Xns_mem_optional_imm12_1_unsigned}, nil},

	{0xffe00c00, 0x38200800, STRB, instArgs{arg_Wt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__absent_0__0_1}, nil},

	{0xffe00c00, 0x78000400, STRH, instArgs{arg_Wt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x78000c00, STRH, instArgs{arg_Wt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0x79000000, STRH, instArgs{arg_Wt, arg_Xns_mem_optional_imm12_2_unsigned}, nil},

	{0xffe00c00, 0x78200800, STRH, instArgs{arg_Wt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__1_1}, nil},

	{0xffe00c00, 0xb8000800, STTR, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xf8000800, STTR, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x38000800, STTRB, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x78000800, STTRH, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xb8000000, STUR, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xf8000000, STUR, instArgs{arg_Xt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x38000000, STURB, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x78000000, STURH, instArgs{arg_Wt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe08000, 0x88200000, STXP, instArgs{arg_Ws, arg_Wt, arg_Wt2, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8200000, STXP, instArgs{arg_Ws, arg_Xt, arg_Xt2, arg_Xns_mem}, nil},

	{0xffe08000, 0x88000000, STXR, instArgs{arg_Ws, arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0xc8000000, STXR, instArgs{arg_Ws, arg_Xt, arg_Xns_mem}, nil},

	{0xffe08000, 0x08000000, STXRB, instArgs{arg_Ws, arg_Wt, arg_Xns_mem}, nil},

	{0xffe08000, 0x48000000, STXRH, instArgs{arg_Ws, arg_Wt, arg_Xns_mem}, nil},

	{0xffe00000, 0x4b200000, SUB, instArgs{arg_Wds, arg_Wns, arg_Wm_extend__UXTB_0__UXTH_1__LSL_UXTW_2__UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xffe00000, 0xcb200000, SUB, instArgs{arg_Xds, arg_Xns, arg_Rm_extend__UXTB_0__UXTH_1__UXTW_2__LSL_UXTX_3__SXTB_4__SXTH_5__SXTW_6__SXTX_7__0_4}, nil},

	{0xff000000, 0x51000000, SUB, instArgs{arg_Wds, arg_Wns, arg_IAddSub}, nil},

	{0xff000000, 0xd1000000, SUB, instArgs{arg_Xds, arg_Xns, arg_IAddSub}, nil},

	{0xffe0001f, 0xd4000001, SVC, instArgs{arg_immediate_0_65535_imm16}, nil},

	{0xfff80000, 0xd5280000, SYSL, instArgs{arg_Xt, arg_immediate_0_7_op1, arg_Cn, arg_Cm, arg_immediate_0_7_op2}, nil},

	{0x7f000000, 0x37000000, TBNZ, instArgs{arg_Rt_31_1__W_0__X_1, arg_immediate_0_63_b5_b40, arg_slabel_imm14_2}, nil},

	{0x7f000000, 0x36000000, TBZ, instArgs{arg_Rt_31_1__W_0__X_1, arg_immediate_0_63_b5_b40, arg_slabel_imm14_2}, nil},

	{0xffe0fc00, 0x1ac00800, UDIV, instArgs{arg_Wd, arg_Wn, arg_Wm}, nil},

	{0xffe0fc00, 0x9ac00800, UDIV, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xffe0fc00, 0x9ba07c00, UMULL, instArgs{arg_Xd, arg_Wn, arg_Wm}, nil},

	{0xffe08000, 0x9ba00000, UMADDL, instArgs{arg_Xd, arg_Wn, arg_Wm, arg_Xa}, nil},

	{0xffe0fc00, 0x9ba0fc00, UMNEGL, instArgs{arg_Xd, arg_Wn, arg_Wm}, nil},

	{0xffe08000, 0x9ba08000, UMSUBL, instArgs{arg_Xd, arg_Wn, arg_Wm, arg_Xa}, nil},

	{0xffe08000, 0x9bc00000, UMULH, instArgs{arg_Xd, arg_Xn, arg_Xm}, nil},

	{0xff3ffc00, 0x5e20b800, ABS, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3}, nil},

	{0xbf3ffc00, 0x0e20b800, ABS, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff20fc00, 0x5e208400, ADD, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x0e208400, ADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff20fc00, 0x0e204000, ADDHN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff20fc00, 0x4e204000, ADDHN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff3ffc00, 0x5e31b800, ADDP, instArgs{arg_Vd_22_2__D_3, arg_Vn_arrangement_size___2D_3}, nil},

	{0xbf20fc00, 0x0e20bc00, ADDP, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xbf3ffc00, 0x0e31b800, ADDV, instArgs{arg_Vd_22_2__B_0__H_1__S_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__4S_21}, nil},

	{0xfffffc00, 0x4e285800, AESD, instArgs{arg_Vd_arrangement_16B, arg_Vn_arrangement_16B}, nil},

	{0xfffffc00, 0x4e284800, AESE, instArgs{arg_Vd_arrangement_16B, arg_Vn_arrangement_16B}, nil},

	{0xfffffc00, 0x4e287800, AESIMC, instArgs{arg_Vd_arrangement_16B, arg_Vn_arrangement_16B}, nil},

	{0xfffffc00, 0x4e286800, AESMC, instArgs{arg_Vd_arrangement_16B, arg_Vn_arrangement_16B}, nil},

	{0xbfe0fc00, 0x0e201c00, AND, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbff8dc00, 0x2f009400, BIC, instArgs{arg_Vd_arrangement_Q___4H_0__8H_1, arg_immediate_OptLSL__a_b_c_d_e_f_g_h_cmode__0_0__8_1}, nil},

	{0xbff89c00, 0x2f001400, BIC, instArgs{arg_Vd_arrangement_Q___2S_0__4S_1, arg_immediate_OptLSL__a_b_c_d_e_f_g_h_cmode__0_0__8_1__16_2__24_3}, nil},

	{0xbfe0fc00, 0x0e601c00, BIC, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x2ee01c00, BIF, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x2ea01c00, BIT, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x2e601c00, BSL, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbf3ffc00, 0x0e204800, CLS, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x2e204800, CLZ, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x7e208c00, CMEQ, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x2e208c00, CMEQ, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff3ffc00, 0x5e209800, CMEQ, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_immediate_zero}, nil},

	{0xbf3ffc00, 0x0e209800, CMEQ, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_immediate_zero}, nil},

	{0xff20fc00, 0x5e203c00, CMGE, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x0e203c00, CMGE, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff3ffc00, 0x7e208800, CMGE, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_immediate_zero}, nil},

	{0xbf3ffc00, 0x2e208800, CMGE, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_immediate_zero}, nil},

	{0xff20fc00, 0x5e203400, CMGT, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x0e203400, CMGT, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff3ffc00, 0x5e208800, CMGT, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_immediate_zero}, nil},

	{0xbf3ffc00, 0x0e208800, CMGT, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_immediate_zero}, nil},

	{0xff20fc00, 0x7e203400, CMHI, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x2e203400, CMHI, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff20fc00, 0x7e203c00, CMHS, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x2e203c00, CMHS, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff3ffc00, 0x7e209800, CMLE, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_immediate_zero}, nil},

	{0xbf3ffc00, 0x2e209800, CMLE, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_immediate_zero}, nil},

	{0xff3ffc00, 0x5e20a800, CMLT, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_immediate_zero}, nil},

	{0xbf3ffc00, 0x0e20a800, CMLT, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_immediate_zero}, nil},

	{0xff20fc00, 0x5e208c00, CMTST, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x0e208c00, CMTST, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xbf3ffc00, 0x0e205800, CNT, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01, arg_Vn_arrangement_size_Q___8B_00__16B_01}, nil},

	{0xffe0fc00, 0x5e000400, MOV, instArgs{arg_Vd_16_5__B_1__H_2__S_4__D_8, arg_Vn_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4__imm5lt4gt_8_1}, nil},

	{0xffe0fc00, 0x5e000400, DUP, instArgs{arg_Vd_16_5__B_1__H_2__S_4__D_8, arg_Vn_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4__imm5lt4gt_8_1}, nil},

	{0xbfe0fc00, 0x0e000400, DUP, instArgs{arg_Vd_arrangement_imm5_Q___8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4__imm5lt4gt_8_1}, nil},

	{0xbfe0fc00, 0x0e000c00, DUP, instArgs{arg_Vd_arrangement_imm5_Q___8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Rn_16_5__W_1__W_2__W_4__X_8}, nil},

	{0xbfe0fc00, 0x2e201c00, EOR, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe08400, 0x2e000000, EXT, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1, arg_immediate_index_Q_imm4__imm4lt20gt_00__imm4_10}, nil},

	{0xffa0fc00, 0x7ea0d400, FABD, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x2ea0d400, FABD, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e20c000, FABS, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e60c000, FABS, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x0ea0f800, FABS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffa0fc00, 0x7e20ec00, FACGE, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x2e20ec00, FACGE, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffa0fc00, 0x7ea0ec00, FACGT, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x2ea0ec00, FACGT, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe0fc00, 0x1e202800, FADD, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e602800, FADD, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xbfa0fc00, 0x0e20d400, FADD, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x7e30d800, FADDP, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_arrangement_sz___2S_0__2D_1}, nil},

	{0xbfa0fc00, 0x2e20d400, FADDP, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe00c10, 0x1e200400, FCCMP, instArgs{arg_Sn, arg_Sm, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0x1e600400, FCCMP, instArgs{arg_Dn, arg_Dm, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0x1e200410, FCCMPE, instArgs{arg_Sn, arg_Sm, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c10, 0x1e600410, FCCMPE, instArgs{arg_Dn, arg_Dm, arg_immediate_0_15_nzcv, arg_cond_AllowALNV_Normal}, nil},

	{0xffa0fc00, 0x5e20e400, FCMEQ, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x0e20e400, FCMEQ, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x5ea0d800, FCMEQ, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_immediate_floatzero}, nil},

	{0xbfbffc00, 0x0ea0d800, FCMEQ, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_immediate_floatzero}, nil},

	{0xffa0fc00, 0x7e20e400, FCMGE, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x2e20e400, FCMGE, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x7ea0c800, FCMGE, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_immediate_floatzero}, nil},

	{0xbfbffc00, 0x2ea0c800, FCMGE, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_immediate_floatzero}, nil},

	{0xffa0fc00, 0x7ea0e400, FCMGT, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x2ea0e400, FCMGT, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x5ea0c800, FCMGT, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_immediate_floatzero}, nil},

	{0xbfbffc00, 0x0ea0c800, FCMGT, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_immediate_floatzero}, nil},

	{0xffbffc00, 0x7ea0d800, FCMLE, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_immediate_floatzero}, nil},

	{0xbfbffc00, 0x2ea0d800, FCMLE, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_immediate_floatzero}, nil},

	{0xffbffc00, 0x5ea0e800, FCMLT, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_immediate_floatzero}, nil},

	{0xbfbffc00, 0x0ea0e800, FCMLT, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_immediate_floatzero}, nil},

	{0xffe0fc1f, 0x1e202000, FCMP, instArgs{arg_Sn, arg_Sm}, nil},

	{0xffe0fc1f, 0x1e202008, FCMP, instArgs{arg_Sn, arg_immediate_floatzero}, nil},

	{0xffe0fc1f, 0x1e602000, FCMP, instArgs{arg_Dn, arg_Dm}, nil},

	{0xffe0fc1f, 0x1e602008, FCMP, instArgs{arg_Dn, arg_immediate_floatzero}, nil},

	{0xffe0fc1f, 0x1e202010, FCMPE, instArgs{arg_Sn, arg_Sm}, nil},

	{0xffe0fc1f, 0x1e202018, FCMPE, instArgs{arg_Sn, arg_immediate_floatzero}, nil},

	{0xffe0fc1f, 0x1e602010, FCMPE, instArgs{arg_Dn, arg_Dm}, nil},

	{0xffe0fc1f, 0x1e602018, FCMPE, instArgs{arg_Dn, arg_immediate_floatzero}, nil},

	{0xffe00c00, 0x1e200c00, FCSEL, instArgs{arg_Sd, arg_Sn, arg_Sm, arg_cond_AllowALNV_Normal}, nil},

	{0xffe00c00, 0x1e600c00, FCSEL, instArgs{arg_Dd, arg_Dn, arg_Dm, arg_cond_AllowALNV_Normal}, nil},

	{0xfffffc00, 0x1ee24000, FCVT, instArgs{arg_Sd, arg_Hn}, nil},

	{0xfffffc00, 0x1ee2c000, FCVT, instArgs{arg_Dd, arg_Hn}, nil},

	{0xfffffc00, 0x1e23c000, FCVT, instArgs{arg_Hd, arg_Sn}, nil},

	{0xfffffc00, 0x1e22c000, FCVT, instArgs{arg_Dd, arg_Sn}, nil},

	{0xfffffc00, 0x1e63c000, FCVT, instArgs{arg_Hd, arg_Dn}, nil},

	{0xfffffc00, 0x1e624000, FCVT, instArgs{arg_Sd, arg_Dn}, nil},

	{0xfffffc00, 0x1e240000, FCVTAS, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e240000, FCVTAS, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e640000, FCVTAS, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e640000, FCVTAS, instArgs{arg_Xd, arg_Dn}, nil},

	{0xffbffc00, 0x5e21c800, FCVTAS, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x0e21c800, FCVTAS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e250000, FCVTAU, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e250000, FCVTAU, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e650000, FCVTAU, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e650000, FCVTAU, instArgs{arg_Xd, arg_Dn}, nil},

	{0xffbffc00, 0x7e21c800, FCVTAU, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x2e21c800, FCVTAU, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x0e217800, FCVTL, instArgs{arg_Vd_arrangement_sz___4S_0__2D_1, arg_Vn_arrangement_sz_Q___4H_00__8H_01__2S_10__4S_11}, nil},

	{0xffbffc00, 0x4e217800, FCVTL2, instArgs{arg_Vd_arrangement_sz___4S_0__2D_1, arg_Vn_arrangement_sz_Q___4H_00__8H_01__2S_10__4S_11}, nil},

	{0xfffffc00, 0x1e300000, FCVTMS, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e300000, FCVTMS, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e700000, FCVTMS, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e700000, FCVTMS, instArgs{arg_Xd, arg_Dn}, nil},

	{0xffbffc00, 0x5e21b800, FCVTMS, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x0e21b800, FCVTMS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e310000, FCVTMU, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e310000, FCVTMU, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e710000, FCVTMU, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e710000, FCVTMU, instArgs{arg_Xd, arg_Dn}, nil},

	{0xffbffc00, 0x7e21b800, FCVTMU, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x2e21b800, FCVTMU, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x0e216800, FCVTN, instArgs{arg_Vd_arrangement_sz_Q___4H_00__8H_01__2S_10__4S_11, arg_Vn_arrangement_sz___4S_0__2D_1}, nil},

	{0xffbffc00, 0x4e216800, FCVTN2, instArgs{arg_Vd_arrangement_sz_Q___4H_00__8H_01__2S_10__4S_11, arg_Vn_arrangement_sz___4S_0__2D_1}, nil},

	{0xfffffc00, 0x1e200000, FCVTNS, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e200000, FCVTNS, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e600000, FCVTNS, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e600000, FCVTNS, instArgs{arg_Xd, arg_Dn}, nil},

	{0xffbffc00, 0x5e21a800, FCVTNS, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x0e21a800, FCVTNS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e210000, FCVTNU, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e210000, FCVTNU, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e610000, FCVTNU, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e610000, FCVTNU, instArgs{arg_Xd, arg_Dn}, nil},

	{0xffbffc00, 0x7e21a800, FCVTNU, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x2e21a800, FCVTNU, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e280000, FCVTPS, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e280000, FCVTPS, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e680000, FCVTPS, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e680000, FCVTPS, instArgs{arg_Xd, arg_Dn}, nil},

	{0xffbffc00, 0x5ea1a800, FCVTPS, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x0ea1a800, FCVTPS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e290000, FCVTPU, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e290000, FCVTPU, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e690000, FCVTPU, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e690000, FCVTPU, instArgs{arg_Xd, arg_Dn}, nil},

	{0xffbffc00, 0x7ea1a800, FCVTPU, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x2ea1a800, FCVTPU, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x7e216800, FCVTXN, instArgs{arg_Vd_22_1__S_1, arg_Vn_22_1__D_1}, nil},

	{0xffbffc00, 0x2e216800, FCVTXN, instArgs{arg_Vd_arrangement_sz_Q___2S_10__4S_11, arg_Vn_arrangement_sz___2D_1}, nil},

	{0xffbffc00, 0x6e216800, FCVTXN2, instArgs{arg_Vd_arrangement_sz_Q___2S_10__4S_11, arg_Vn_arrangement_sz___2D_1}, nil},

	{0xffff0000, 0x1e180000, FCVTZS, instArgs{arg_Wd, arg_Sn, arg_immediate_fbits_min_1_max_32_sub_64_scale}, nil},

	{0xffff0000, 0x9e180000, FCVTZS, instArgs{arg_Xd, arg_Sn, arg_immediate_fbits_min_1_max_64_sub_64_scale}, nil},

	{0xffff0000, 0x1e580000, FCVTZS, instArgs{arg_Wd, arg_Dn, arg_immediate_fbits_min_1_max_32_sub_64_scale}, nil},

	{0xffff0000, 0x9e580000, FCVTZS, instArgs{arg_Xd, arg_Dn, arg_immediate_fbits_min_1_max_64_sub_64_scale}, nil},

	{0xfffffc00, 0x1e380000, FCVTZS, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e380000, FCVTZS, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e780000, FCVTZS, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e780000, FCVTZS, instArgs{arg_Xd, arg_Dn}, nil},

	{0xff80fc00, 0x5f00fc00, FCVTZS, instArgs{arg_Vd_19_4__S_4__D_8, arg_Vn_19_4__S_4__D_8, arg_immediate_fbits_min_1_max_0_sub_0_immh_immb__64UIntimmhimmb_4__128UIntimmhimmb_8}, fcvtzs_asisdshf_c_cond},

	{0xbf80fc00, 0x0f00fc00, FCVTZS, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__2S_40__4S_41__2D_81, arg_immediate_fbits_min_1_max_0_sub_0_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__64UIntimmhimmb_4__128UIntimmhimmb_8}, fcvtzs_asimdshf_c_cond},

	{0xffbffc00, 0x5ea1b800, FCVTZS, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x0ea1b800, FCVTZS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffff0000, 0x1e190000, FCVTZU, instArgs{arg_Wd, arg_Sn, arg_immediate_fbits_min_1_max_32_sub_64_scale}, nil},

	{0xffff0000, 0x9e190000, FCVTZU, instArgs{arg_Xd, arg_Sn, arg_immediate_fbits_min_1_max_64_sub_64_scale}, nil},

	{0xffff0000, 0x1e590000, FCVTZU, instArgs{arg_Wd, arg_Dn, arg_immediate_fbits_min_1_max_32_sub_64_scale}, nil},

	{0xffff0000, 0x9e590000, FCVTZU, instArgs{arg_Xd, arg_Dn, arg_immediate_fbits_min_1_max_64_sub_64_scale}, nil},

	{0xfffffc00, 0x1e390000, FCVTZU, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e390000, FCVTZU, instArgs{arg_Xd, arg_Sn}, nil},

	{0xfffffc00, 0x1e790000, FCVTZU, instArgs{arg_Wd, arg_Dn}, nil},

	{0xfffffc00, 0x9e790000, FCVTZU, instArgs{arg_Xd, arg_Dn}, nil},

	{0xff80fc00, 0x7f00fc00, FCVTZU, instArgs{arg_Vd_19_4__S_4__D_8, arg_Vn_19_4__S_4__D_8, arg_immediate_fbits_min_1_max_0_sub_0_immh_immb__64UIntimmhimmb_4__128UIntimmhimmb_8}, fcvtzu_asisdshf_c_cond},

	{0xbf80fc00, 0x2f00fc00, FCVTZU, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__2S_40__4S_41__2D_81, arg_immediate_fbits_min_1_max_0_sub_0_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__64UIntimmhimmb_4__128UIntimmhimmb_8}, fcvtzu_asimdshf_c_cond},

	{0xffbffc00, 0x7ea1b800, FCVTZU, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x2ea1b800, FCVTZU, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe0fc00, 0x1e201800, FDIV, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e601800, FDIV, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xbfa0fc00, 0x2e20fc00, FDIV, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe08000, 0x1f000000, FMADD, instArgs{arg_Sd, arg_Sn, arg_Sm, arg_Sa}, nil},

	{0xffe08000, 0x1f400000, FMADD, instArgs{arg_Dd, arg_Dn, arg_Dm, arg_Da}, nil},

	{0xffe0fc00, 0x1e204800, FMAX, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e604800, FMAX, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xbfa0fc00, 0x0e20f400, FMAX, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe0fc00, 0x1e206800, FMAXNM, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e606800, FMAXNM, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xbfa0fc00, 0x0e20c400, FMAXNM, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x7e30c800, FMAXNMP, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_arrangement_sz___2S_0__2D_1}, nil},

	{0xbfa0fc00, 0x2e20c400, FMAXNMP, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xbfbffc00, 0x2e30c800, FMAXNMV, instArgs{arg_Vd_22_1__S_0, arg_Vn_arrangement_Q_sz___4S_10}, nil},

	{0xffbffc00, 0x7e30f800, FMAXP, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_arrangement_sz___2S_0__2D_1}, nil},

	{0xbfa0fc00, 0x2e20f400, FMAXP, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xbfbffc00, 0x2e30f800, FMAXV, instArgs{arg_Vd_22_1__S_0, arg_Vn_arrangement_Q_sz___4S_10}, nil},

	{0xffe0fc00, 0x1e205800, FMIN, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e605800, FMIN, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xbfa0fc00, 0x0ea0f400, FMIN, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe0fc00, 0x1e207800, FMINNM, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e607800, FMINNM, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xbfa0fc00, 0x0ea0c400, FMINNM, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x7eb0c800, FMINNMP, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_arrangement_sz___2S_0__2D_1}, nil},

	{0xbfa0fc00, 0x2ea0c400, FMINNMP, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xbfbffc00, 0x2eb0c800, FMINNMV, instArgs{arg_Vd_22_1__S_0, arg_Vn_arrangement_Q_sz___4S_10}, nil},

	{0xffbffc00, 0x7eb0f800, FMINP, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_arrangement_sz___2S_0__2D_1}, nil},

	{0xbfa0fc00, 0x2ea0f400, FMINP, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xbfbffc00, 0x2eb0f800, FMINV, instArgs{arg_Vd_22_1__S_0, arg_Vn_arrangement_Q_sz___4S_10}, nil},

	{0xff80f400, 0x5f801000, FMLA, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_arrangement_sz___S_0__D_1_index__sz_L_H__HL_00__H_10_1}, nil},

	{0xbf80f400, 0x0f801000, FMLA, instArgs{arg_Vd_arrangement_Q_sz___2S_00__4S_10__2D_11, arg_Vn_arrangement_Q_sz___2S_00__4S_10__2D_11, arg_Vm_arrangement_sz___S_0__D_1_index__sz_L_H__HL_00__H_10_1}, nil},

	{0xbfa0fc00, 0x0e20cc00, FMLA, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xff80f400, 0x5f805000, FMLS, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_arrangement_sz___S_0__D_1_index__sz_L_H__HL_00__H_10_1}, nil},

	{0xbf80f400, 0x0f805000, FMLS, instArgs{arg_Vd_arrangement_Q_sz___2S_00__4S_10__2D_11, arg_Vn_arrangement_Q_sz___2S_00__4S_10__2D_11, arg_Vm_arrangement_sz___S_0__D_1_index__sz_L_H__HL_00__H_10_1}, nil},

	{0xbfa0fc00, 0x0ea0cc00, FMLS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e270000, FMOV, instArgs{arg_Sd, arg_Wn}, nil},

	{0xfffffc00, 0x1e260000, FMOV, instArgs{arg_Wd, arg_Sn}, nil},

	{0xfffffc00, 0x9e670000, FMOV, instArgs{arg_Dd, arg_Xn}, nil},

	{0xfffffc00, 0x9eaf0000, FMOV, instArgs{arg_Vd_arrangement_D_index__1, arg_Xn}, nil},

	{0xfffffc00, 0x9e660000, FMOV, instArgs{arg_Xd, arg_Dn}, nil},

	{0xfffffc00, 0x9eae0000, FMOV, instArgs{arg_Xd, arg_Vn_arrangement_D_index__1}, nil},

	{0xfffffc00, 0x1e204000, FMOV, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e604000, FMOV, instArgs{arg_Dd, arg_Dn}, nil},

	{0xffe01fe0, 0x1e201000, FMOV, instArgs{arg_Sd, arg_immediate_exp_3_pre_4_imm8}, nil},

	{0xffe01fe0, 0x1e601000, FMOV, instArgs{arg_Dd, arg_immediate_exp_3_pre_4_imm8}, nil},

	{0xbff8fc00, 0x0f00f400, FMOV, instArgs{arg_Vd_arrangement_Q___2S_0__4S_1, arg_immediate_exp_3_pre_4_a_b_c_d_e_f_g_h}, nil},

	{0xfff8fc00, 0x6f00f400, FMOV, instArgs{arg_Vd_arrangement_2D, arg_immediate_exp_3_pre_4_a_b_c_d_e_f_g_h}, nil},

	{0xffe08000, 0x1f008000, FMSUB, instArgs{arg_Sd, arg_Sn, arg_Sm, arg_Sa}, nil},

	{0xffe08000, 0x1f408000, FMSUB, instArgs{arg_Dd, arg_Dn, arg_Dm, arg_Da}, nil},

	{0xff80f400, 0x5f809000, FMUL, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_arrangement_sz___S_0__D_1_index__sz_L_H__HL_00__H_10_1}, nil},

	{0xbf80f400, 0x0f809000, FMUL, instArgs{arg_Vd_arrangement_Q_sz___2S_00__4S_10__2D_11, arg_Vn_arrangement_Q_sz___2S_00__4S_10__2D_11, arg_Vm_arrangement_sz___S_0__D_1_index__sz_L_H__HL_00__H_10_1}, nil},

	{0xffe0fc00, 0x1e200800, FMUL, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e600800, FMUL, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xbfa0fc00, 0x2e20dc00, FMUL, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffa0fc00, 0x5e20dc00, FMULX, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x0e20dc00, FMULX, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xff80f400, 0x7f809000, FMULX, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_arrangement_sz___S_0__D_1_index__sz_L_H__HL_00__H_10_1}, nil},

	{0xbf80f400, 0x2f809000, FMULX, instArgs{arg_Vd_arrangement_Q_sz___2S_00__4S_10__2D_11, arg_Vn_arrangement_Q_sz___2S_00__4S_10__2D_11, arg_Vm_arrangement_sz___S_0__D_1_index__sz_L_H__HL_00__H_10_1}, nil},

	{0xfffffc00, 0x1e214000, FNEG, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e614000, FNEG, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x2ea0f800, FNEG, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe08000, 0x1f200000, FNMADD, instArgs{arg_Sd, arg_Sn, arg_Sm, arg_Sa}, nil},

	{0xffe08000, 0x1f600000, FNMADD, instArgs{arg_Dd, arg_Dn, arg_Dm, arg_Da}, nil},

	{0xffe08000, 0x1f208000, FNMSUB, instArgs{arg_Sd, arg_Sn, arg_Sm, arg_Sa}, nil},

	{0xffe08000, 0x1f608000, FNMSUB, instArgs{arg_Dd, arg_Dn, arg_Dm, arg_Da}, nil},

	{0xffe0fc00, 0x1e208800, FNMUL, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e608800, FNMUL, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xffbffc00, 0x5ea1d800, FRECPE, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x0ea1d800, FRECPE, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffa0fc00, 0x5e20fc00, FRECPS, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x0e20fc00, FRECPS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x5ea1f800, FRECPX, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xfffffc00, 0x1e264000, FRINTA, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e664000, FRINTA, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x2e218800, FRINTA, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e27c000, FRINTI, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e67c000, FRINTI, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x2ea19800, FRINTI, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e254000, FRINTM, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e654000, FRINTM, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x0e219800, FRINTM, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e244000, FRINTN, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e644000, FRINTN, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x0e218800, FRINTN, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e24c000, FRINTP, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e64c000, FRINTP, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x0ea18800, FRINTP, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e274000, FRINTX, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e674000, FRINTX, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x2e219800, FRINTX, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e25c000, FRINTZ, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e65c000, FRINTZ, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x0ea19800, FRINTZ, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffbffc00, 0x7ea1d800, FRSQRTE, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x2ea1d800, FRSQRTE, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffa0fc00, 0x5ea0fc00, FRSQRTS, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1, arg_Vm_22_1__S_0__D_1}, nil},

	{0xbfa0fc00, 0x0ea0fc00, FRSQRTS, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xfffffc00, 0x1e21c000, FSQRT, instArgs{arg_Sd, arg_Sn}, nil},

	{0xfffffc00, 0x1e61c000, FSQRT, instArgs{arg_Dd, arg_Dn}, nil},

	{0xbfbffc00, 0x2ea1f800, FSQRT, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe0fc00, 0x1e203800, FSUB, instArgs{arg_Sd, arg_Sn, arg_Sm}, nil},

	{0xffe0fc00, 0x1e603800, FSUB, instArgs{arg_Dd, arg_Dn, arg_Dm}, nil},

	{0xbfa0fc00, 0x0ea0d400, FSUB, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vm_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe08400, 0x6e000400, MOV, instArgs{arg_Vd_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4__imm5lt4gt_8_1, arg_Vn_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5_imm4__imm4lt30gt_1__imm4lt31gt_2__imm4lt32gt_4__imm4lt3gt_8_1}, nil},

	{0xffe08400, 0x6e000400, INS, instArgs{arg_Vd_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4__imm5lt4gt_8_1, arg_Vn_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5_imm4__imm4lt30gt_1__imm4lt31gt_2__imm4lt32gt_4__imm4lt3gt_8_1}, nil},

	{0xffe0fc00, 0x4e001c00, MOV, instArgs{arg_Vd_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4__imm5lt4gt_8_1, arg_Rn_16_5__W_1__W_2__W_4__X_8}, nil},

	{0xffe0fc00, 0x4e001c00, INS, instArgs{arg_Vd_arrangement_imm5___B_1__H_2__S_4__D_8_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4__imm5lt4gt_8_1, arg_Rn_16_5__W_1__W_2__W_4__X_8}, nil},

	{0xbffff000, 0x0c407000, LD1, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c40a000, LD1, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c406000, LD1, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c402000, LD1, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0cdf7000, LD1, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Q__8_0__16_1}, nil},

	{0xbfe0f000, 0x0cc07000, LD1, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0cdfa000, LD1, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Q__16_0__32_1}, nil},

	{0xbfe0f000, 0x0cc0a000, LD1, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0cdf6000, LD1, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Q__24_0__48_1}, nil},

	{0xbfe0f000, 0x0cc06000, LD1, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0cdf2000, LD1, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Q__32_0__64_1}, nil},

	{0xbfe0f000, 0x0cc02000, LD1, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe000, 0x0d400000, LD1, instArgs{arg_Vt_1_arrangement_B_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffe400, 0x0d404000, LD1, instArgs{arg_Vt_1_arrangement_H_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffec00, 0x0d408000, LD1, instArgs{arg_Vt_1_arrangement_S_index__Q_S_1, arg_Xns_mem_offset}, nil},

	{0xbffffc00, 0x0d408400, LD1, instArgs{arg_Vt_1_arrangement_D_index__Q_1, arg_Xns_mem_offset}, nil},

	{0xbfffe000, 0x0ddf0000, LD1, instArgs{arg_Vt_1_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_1}, nil},

	{0xbfe0e000, 0x0dc00000, LD1, instArgs{arg_Vt_1_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe400, 0x0ddf4000, LD1, instArgs{arg_Vt_1_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_2}, nil},

	{0xbfe0e400, 0x0dc04000, LD1, instArgs{arg_Vt_1_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffec00, 0x0ddf8000, LD1, instArgs{arg_Vt_1_arrangement_S_index__Q_S_1, arg_Xns_mem_post_fixedimm_4}, nil},

	{0xbfe0ec00, 0x0dc08000, LD1, instArgs{arg_Vt_1_arrangement_S_index__Q_S_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffffc00, 0x0ddf8400, LD1, instArgs{arg_Vt_1_arrangement_D_index__Q_1, arg_Xns_mem_post_fixedimm_8}, nil},

	{0xbfe0fc00, 0x0dc08400, LD1, instArgs{arg_Vt_1_arrangement_D_index__Q_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0d40c000, LD1R, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0ddfc000, LD1R, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_size__1_0__2_1__4_2__8_3}, nil},

	{0xbfe0f000, 0x0dc0c000, LD1R, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c408000, LD2, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0cdf8000, LD2, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Q__16_0__32_1}, nil},

	{0xbfe0f000, 0x0cc08000, LD2, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe000, 0x0d600000, LD2, instArgs{arg_Vt_2_arrangement_B_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffe400, 0x0d604000, LD2, instArgs{arg_Vt_2_arrangement_H_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffec00, 0x0d608000, LD2, instArgs{arg_Vt_2_arrangement_S_index__Q_S_1, arg_Xns_mem_offset}, nil},

	{0xbffffc00, 0x0d608400, LD2, instArgs{arg_Vt_2_arrangement_D_index__Q_1, arg_Xns_mem_offset}, nil},

	{0xbfffe000, 0x0dff0000, LD2, instArgs{arg_Vt_2_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_2}, nil},

	{0xbfe0e000, 0x0de00000, LD2, instArgs{arg_Vt_2_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe400, 0x0dff4000, LD2, instArgs{arg_Vt_2_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_4}, nil},

	{0xbfe0e400, 0x0de04000, LD2, instArgs{arg_Vt_2_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffec00, 0x0dff8000, LD2, instArgs{arg_Vt_2_arrangement_S_index__Q_S_1, arg_Xns_mem_post_fixedimm_8}, nil},

	{0xbfe0ec00, 0x0de08000, LD2, instArgs{arg_Vt_2_arrangement_S_index__Q_S_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffffc00, 0x0dff8400, LD2, instArgs{arg_Vt_2_arrangement_D_index__Q_1, arg_Xns_mem_post_fixedimm_16}, nil},

	{0xbfe0fc00, 0x0de08400, LD2, instArgs{arg_Vt_2_arrangement_D_index__Q_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0d60c000, LD2R, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0dffc000, LD2R, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_size__2_0__4_1__8_2__16_3}, nil},

	{0xbfe0f000, 0x0de0c000, LD2R, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c404000, LD3, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0cdf4000, LD3, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Q__24_0__48_1}, nil},

	{0xbfe0f000, 0x0cc04000, LD3, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe000, 0x0d402000, LD3, instArgs{arg_Vt_3_arrangement_B_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffe400, 0x0d406000, LD3, instArgs{arg_Vt_3_arrangement_H_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffec00, 0x0d40a000, LD3, instArgs{arg_Vt_3_arrangement_S_index__Q_S_1, arg_Xns_mem_offset}, nil},

	{0xbffffc00, 0x0d40a400, LD3, instArgs{arg_Vt_3_arrangement_D_index__Q_1, arg_Xns_mem_offset}, nil},

	{0xbfffe000, 0x0ddf2000, LD3, instArgs{arg_Vt_3_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_3}, nil},

	{0xbfe0e000, 0x0dc02000, LD3, instArgs{arg_Vt_3_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe400, 0x0ddf6000, LD3, instArgs{arg_Vt_3_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_6}, nil},

	{0xbfe0e400, 0x0dc06000, LD3, instArgs{arg_Vt_3_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffec00, 0x0ddfa000, LD3, instArgs{arg_Vt_3_arrangement_S_index__Q_S_1, arg_Xns_mem_post_fixedimm_12}, nil},

	{0xbfe0ec00, 0x0dc0a000, LD3, instArgs{arg_Vt_3_arrangement_S_index__Q_S_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffffc00, 0x0ddfa400, LD3, instArgs{arg_Vt_3_arrangement_D_index__Q_1, arg_Xns_mem_post_fixedimm_24}, nil},

	{0xbfe0fc00, 0x0dc0a400, LD3, instArgs{arg_Vt_3_arrangement_D_index__Q_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0d40e000, LD3R, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0ddfe000, LD3R, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_size__3_0__6_1__12_2__24_3}, nil},

	{0xbfe0f000, 0x0dc0e000, LD3R, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c400000, LD4, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0cdf0000, LD4, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Q__32_0__64_1}, nil},

	{0xbfe0f000, 0x0cc00000, LD4, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe000, 0x0d602000, LD4, instArgs{arg_Vt_4_arrangement_B_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffe400, 0x0d606000, LD4, instArgs{arg_Vt_4_arrangement_H_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffec00, 0x0d60a000, LD4, instArgs{arg_Vt_4_arrangement_S_index__Q_S_1, arg_Xns_mem_offset}, nil},

	{0xbffffc00, 0x0d60a400, LD4, instArgs{arg_Vt_4_arrangement_D_index__Q_1, arg_Xns_mem_offset}, nil},

	{0xbfffe000, 0x0dff2000, LD4, instArgs{arg_Vt_4_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_4}, nil},

	{0xbfe0e000, 0x0de02000, LD4, instArgs{arg_Vt_4_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe400, 0x0dff6000, LD4, instArgs{arg_Vt_4_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_8}, nil},

	{0xbfe0e400, 0x0de06000, LD4, instArgs{arg_Vt_4_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffec00, 0x0dffa000, LD4, instArgs{arg_Vt_4_arrangement_S_index__Q_S_1, arg_Xns_mem_post_fixedimm_16}, nil},

	{0xbfe0ec00, 0x0de0a000, LD4, instArgs{arg_Vt_4_arrangement_S_index__Q_S_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffffc00, 0x0dffa400, LD4, instArgs{arg_Vt_4_arrangement_D_index__Q_1, arg_Xns_mem_post_fixedimm_32}, nil},

	{0xbfe0fc00, 0x0de0a400, LD4, instArgs{arg_Vt_4_arrangement_D_index__Q_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0d60e000, LD4R, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0dffe000, LD4R, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_size__4_0__8_1__16_2__32_3}, nil},

	{0xbfe0f000, 0x0de0e000, LD4R, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xffc00000, 0x2c400000, LDNP, instArgs{arg_St, arg_St2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffc00000, 0x6c400000, LDNP, instArgs{arg_Dt, arg_Dt2, arg_Xns_mem_optional_imm7_8_signed}, nil},

	{0xffc00000, 0xac400000, LDNP, instArgs{arg_Qt, arg_Qt2, arg_Xns_mem_optional_imm7_16_signed}, nil},

	{0xffc00000, 0x2cc00000, LDP, instArgs{arg_St, arg_St2, arg_Xns_mem_post_imm7_4_signed}, nil},

	{0xffc00000, 0x6cc00000, LDP, instArgs{arg_Dt, arg_Dt2, arg_Xns_mem_post_imm7_8_signed}, nil},

	{0xffc00000, 0xacc00000, LDP, instArgs{arg_Qt, arg_Qt2, arg_Xns_mem_post_imm7_16_signed}, nil},

	{0xffc00000, 0x2dc00000, LDP, instArgs{arg_St, arg_St2, arg_Xns_mem_wb_imm7_4_signed}, nil},

	{0xffc00000, 0x6dc00000, LDP, instArgs{arg_Dt, arg_Dt2, arg_Xns_mem_wb_imm7_8_signed}, nil},

	{0xffc00000, 0xadc00000, LDP, instArgs{arg_Qt, arg_Qt2, arg_Xns_mem_wb_imm7_16_signed}, nil},

	{0xffc00000, 0x2d400000, LDP, instArgs{arg_St, arg_St2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffc00000, 0x6d400000, LDP, instArgs{arg_Dt, arg_Dt2, arg_Xns_mem_optional_imm7_8_signed}, nil},

	{0xffc00000, 0xad400000, LDP, instArgs{arg_Qt, arg_Qt2, arg_Xns_mem_optional_imm7_16_signed}, nil},

	{0xffe00c00, 0x3c400400, LDR, instArgs{arg_Bt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x7c400400, LDR, instArgs{arg_Ht, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xbc400400, LDR, instArgs{arg_St, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xfc400400, LDR, instArgs{arg_Dt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x3cc00400, LDR, instArgs{arg_Qt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x3c400c00, LDR, instArgs{arg_Bt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0x7c400c00, LDR, instArgs{arg_Ht, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0xbc400c00, LDR, instArgs{arg_St, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0xfc400c00, LDR, instArgs{arg_Dt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0x3cc00c00, LDR, instArgs{arg_Qt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0x3d400000, LDR, instArgs{arg_Bt, arg_Xns_mem_optional_imm12_1_unsigned}, nil},

	{0xffc00000, 0x7d400000, LDR, instArgs{arg_Ht, arg_Xns_mem_optional_imm12_2_unsigned}, nil},

	{0xffc00000, 0xbd400000, LDR, instArgs{arg_St, arg_Xns_mem_optional_imm12_4_unsigned}, nil},

	{0xffc00000, 0xfd400000, LDR, instArgs{arg_Dt, arg_Xns_mem_optional_imm12_8_unsigned}, nil},

	{0xffc00000, 0x3dc00000, LDR, instArgs{arg_Qt, arg_Xns_mem_optional_imm12_16_unsigned}, nil},

	{0xff000000, 0x1c000000, LDR, instArgs{arg_St, arg_slabel_imm19_2}, nil},

	{0xff000000, 0x5c000000, LDR, instArgs{arg_Dt, arg_slabel_imm19_2}, nil},

	{0xff000000, 0x9c000000, LDR, instArgs{arg_Qt, arg_slabel_imm19_2}, nil},

	{0xffe00c00, 0x3c600800, LDR, instArgs{arg_Bt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__absent_0__0_1}, nil},

	{0xffe00c00, 0x7c600800, LDR, instArgs{arg_Ht, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__1_1}, nil},

	{0xffe00c00, 0xbc600800, LDR, instArgs{arg_St, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__2_1}, nil},

	{0xffe00c00, 0xfc600800, LDR, instArgs{arg_Dt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__3_1}, nil},

	{0xffe00c00, 0x3ce00800, LDR, instArgs{arg_Qt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__4_1}, nil},

	{0xffe00c00, 0x3c400000, LDUR, instArgs{arg_Bt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x7c400000, LDUR, instArgs{arg_Ht, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xbc400000, LDUR, instArgs{arg_St, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xfc400000, LDUR, instArgs{arg_Dt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x3cc00000, LDUR, instArgs{arg_Qt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xbf00f400, 0x2f000000, MLA, instArgs{arg_Vd_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xbf20fc00, 0x0e209400, MLA, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf00f400, 0x2f004000, MLS, instArgs{arg_Vd_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xbf20fc00, 0x2e209400, MLS, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xffe0fc00, 0x0e003c00, MOV, instArgs{arg_Wd, arg_Vn_arrangement_S_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4_1}, mov_umov_asimdins_w_w_cond},

	{0xffe0fc00, 0x0e003c00, UMOV, instArgs{arg_Wd, arg_Vn_arrangement_imm5___B_1__H_2__S_4_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4_1}, nil},

	{0xffe0fc00, 0x4e003c00, MOV, instArgs{arg_Xd, arg_Vn_arrangement_D_index__imm5_1}, mov_umov_asimdins_x_x_cond},

	{0xffe0fc00, 0x4e003c00, UMOV, instArgs{arg_Xd, arg_Vn_arrangement_imm5___D_8_index__imm5_1}, nil},

	{0xbfe0fc00, 0x0ea01c00, MOV, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1}, mov_orr_asimdsame_only_cond},

	{0xbfe0fc00, 0x0ea01c00, ORR, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbff8fc00, 0x0f00e400, MOVI, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_immediate_OptLSLZero__a_b_c_d_e_f_g_h}, nil},

	{0xbff8dc00, 0x0f008400, MOVI, instArgs{arg_Vd_arrangement_Q___4H_0__8H_1, arg_immediate_OptLSL__a_b_c_d_e_f_g_h_cmode__0_0__8_1}, nil},

	{0xbff89c00, 0x0f000400, MOVI, instArgs{arg_Vd_arrangement_Q___2S_0__4S_1, arg_immediate_OptLSL__a_b_c_d_e_f_g_h_cmode__0_0__8_1__16_2__24_3}, nil},

	{0xbff8ec00, 0x0f00c400, MOVI, instArgs{arg_Vd_arrangement_Q___2S_0__4S_1, arg_immediate_MSL__a_b_c_d_e_f_g_h_cmode__8_0__16_1}, nil},

	{0xfff8fc00, 0x2f00e400, MOVI, instArgs{arg_Dd, arg_immediate_8x8_a_b_c_d_e_f_g_h}, nil},

	{0xfff8fc00, 0x6f00e400, MOVI, instArgs{arg_Vd_arrangement_2D, arg_immediate_8x8_a_b_c_d_e_f_g_h}, nil},

	{0xbf00f400, 0x0f008000, MUL, instArgs{arg_Vd_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xbf20fc00, 0x0e209c00, MUL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbffffc00, 0x2e205800, MVN, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1}, nil},

	{0xbffffc00, 0x2e205800, NOT, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1}, nil},

	{0xbff8dc00, 0x2f008400, MVNI, instArgs{arg_Vd_arrangement_Q___4H_0__8H_1, arg_immediate_OptLSL__a_b_c_d_e_f_g_h_cmode__0_0__8_1}, nil},

	{0xbff89c00, 0x2f000400, MVNI, instArgs{arg_Vd_arrangement_Q___2S_0__4S_1, arg_immediate_OptLSL__a_b_c_d_e_f_g_h_cmode__0_0__8_1__16_2__24_3}, nil},

	{0xbff8ec00, 0x2f00c400, MVNI, instArgs{arg_Vd_arrangement_Q___2S_0__4S_1, arg_immediate_MSL__a_b_c_d_e_f_g_h_cmode__8_0__16_1}, nil},

	{0xff3ffc00, 0x7e20b800, NEG, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3}, nil},

	{0xbf3ffc00, 0x2e20b800, NEG, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xbfe0fc00, 0x0ee01c00, ORN, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbff8dc00, 0x0f009400, ORR, instArgs{arg_Vd_arrangement_Q___4H_0__8H_1, arg_immediate_OptLSL__a_b_c_d_e_f_g_h_cmode__0_0__8_1}, nil},

	{0xbff89c00, 0x0f001400, ORR, instArgs{arg_Vd_arrangement_Q___2S_0__4S_1, arg_immediate_OptLSL__a_b_c_d_e_f_g_h_cmode__0_0__8_1__16_2__24_3}, nil},

	{0xbf20fc00, 0x2e209c00, PMUL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01, arg_Vn_arrangement_size_Q___8B_00__16B_01, arg_Vm_arrangement_size_Q___8B_00__16B_01}, nil},

	{0xff20fc00, 0x0e20e000, PMULL, instArgs{arg_Vd_arrangement_size___8H_0__1Q_3, arg_Vn_arrangement_size_Q___8B_00__16B_01__1D_30__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__1D_30__2D_31}, nil},

	{0xff20fc00, 0x4e20e000, PMULL2, instArgs{arg_Vd_arrangement_size___8H_0__1Q_3, arg_Vn_arrangement_size_Q___8B_00__16B_01__1D_30__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__1D_30__2D_31}, nil},

	{0xff20fc00, 0x2e204000, RADDHN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff20fc00, 0x6e204000, RADDHN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xbffffc00, 0x2e605800, RBIT, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_arrangement_Q___8B_0__16B_1}, nil},

	{0xbf3ffc00, 0x0e201800, REV16, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01, arg_Vn_arrangement_size_Q___8B_00__16B_01}, nil},

	{0xbf3ffc00, 0x2e200800, REV32, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11}, nil},

	{0xbf3ffc00, 0x0e200800, REV64, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff80fc00, 0x0f008c00, RSHRN, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, rshrn_asimdshf_n_cond},

	{0xff80fc00, 0x4f008c00, RSHRN2, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, rshrn_asimdshf_n_cond},

	{0xff20fc00, 0x2e206000, RSUBHN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff20fc00, 0x6e206000, RSUBHN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xbf20fc00, 0x0e207c00, SABA, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x0e205000, SABAL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e205000, SABAL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x0e207400, SABD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x0e207000, SABDL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e207000, SABDL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x0e206800, SADALP, instArgs{arg_Vd_arrangement_size_Q___4H_00__8H_01__2S_10__4S_11__1D_20__2D_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x0e200000, SADDL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e200000, SADDL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x0e202800, SADDLP, instArgs{arg_Vd_arrangement_size_Q___4H_00__8H_01__2S_10__4S_11__1D_20__2D_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x0e303800, SADDLV, instArgs{arg_Vd_22_2__H_0__S_1__D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__4S_21}, nil},

	{0xff20fc00, 0x0e201000, SADDW, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e201000, SADDW2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xffff0000, 0x1e020000, SCVTF, instArgs{arg_Sd, arg_Wn, arg_immediate_fbits_min_1_max_32_sub_64_scale}, nil},

	{0xffff0000, 0x1e420000, SCVTF, instArgs{arg_Dd, arg_Wn, arg_immediate_fbits_min_1_max_32_sub_64_scale}, nil},

	{0xffff0000, 0x9e020000, SCVTF, instArgs{arg_Sd, arg_Xn, arg_immediate_fbits_min_1_max_64_sub_64_scale}, nil},

	{0xffff0000, 0x9e420000, SCVTF, instArgs{arg_Dd, arg_Xn, arg_immediate_fbits_min_1_max_64_sub_64_scale}, nil},

	{0xfffffc00, 0x1e220000, SCVTF, instArgs{arg_Sd, arg_Wn}, nil},

	{0xfffffc00, 0x1e620000, SCVTF, instArgs{arg_Dd, arg_Wn}, nil},

	{0xfffffc00, 0x9e220000, SCVTF, instArgs{arg_Sd, arg_Xn}, nil},

	{0xfffffc00, 0x9e620000, SCVTF, instArgs{arg_Dd, arg_Xn}, nil},

	{0xff80fc00, 0x5f00e400, SCVTF, instArgs{arg_Vd_19_4__S_4__D_8, arg_Vn_19_4__S_4__D_8, arg_immediate_fbits_min_1_max_0_sub_0_immh_immb__64UIntimmhimmb_4__128UIntimmhimmb_8}, scvtf_asisdshf_c_cond},

	{0xbf80fc00, 0x0f00e400, SCVTF, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__2S_40__4S_41__2D_81, arg_immediate_fbits_min_1_max_0_sub_0_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__64UIntimmhimmb_4__128UIntimmhimmb_8}, scvtf_asimdshf_c_cond},

	{0xffbffc00, 0x5e21d800, SCVTF, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x0e21d800, SCVTF, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xffe0fc00, 0x5e000000, SHA1C, instArgs{arg_Qd, arg_Sn, arg_Vm_arrangement_4S}, nil},

	{0xfffffc00, 0x5e280800, SHA1H, instArgs{arg_Sd, arg_Sn}, nil},

	{0xffe0fc00, 0x5e002000, SHA1M, instArgs{arg_Qd, arg_Sn, arg_Vm_arrangement_4S}, nil},

	{0xffe0fc00, 0x5e001000, SHA1P, instArgs{arg_Qd, arg_Sn, arg_Vm_arrangement_4S}, nil},

	{0xffe0fc00, 0x5e003000, SHA1SU0, instArgs{arg_Vd_arrangement_4S, arg_Vn_arrangement_4S, arg_Vm_arrangement_4S}, nil},

	{0xfffffc00, 0x5e281800, SHA1SU1, instArgs{arg_Vd_arrangement_4S, arg_Vn_arrangement_4S}, nil},

	{0xffe0fc00, 0x5e004000, SHA256H, instArgs{arg_Qd, arg_Qn, arg_Vm_arrangement_4S}, nil},

	{0xffe0fc00, 0x5e005000, SHA256H2, instArgs{arg_Qd, arg_Qn, arg_Vm_arrangement_4S}, nil},

	{0xfffffc00, 0x5e282800, SHA256SU0, instArgs{arg_Vd_arrangement_4S, arg_Vn_arrangement_4S}, nil},

	{0xffe0fc00, 0x5e006000, SHA256SU1, instArgs{arg_Vd_arrangement_4S, arg_Vn_arrangement_4S, arg_Vm_arrangement_4S}, nil},

	{0xbf20fc00, 0x0e200400, SHADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff80fc00, 0x5f005400, SHL, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_0_63_immh_immb__UIntimmhimmb64_8}, shl_asisdshf_r_cond},

	{0xbf80fc00, 0x0f005400, SHL, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4__UIntimmhimmb64_8}, shl_asimdshf_r_cond},

	{0xff3ffc00, 0x2e213800, SHLL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_immediate_0_width_size__8_0__16_1__32_2}, nil},

	{0xff3ffc00, 0x6e213800, SHLL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_immediate_0_width_size__8_0__16_1__32_2}, nil},

	{0xff80fc00, 0x0f008400, SHRN, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, shrn_asimdshf_n_cond},

	{0xff80fc00, 0x4f008400, SHRN2, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, shrn_asimdshf_n_cond},

	{0xbf20fc00, 0x0e202400, SHSUB, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff80fc00, 0x7f005400, SLI, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_0_63_immh_immb__UIntimmhimmb64_8}, sli_asisdshf_r_cond},

	{0xbf80fc00, 0x2f005400, SLI, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4__UIntimmhimmb64_8}, sli_asimdshf_r_cond},

	{0xbf20fc00, 0x0e206400, SMAX, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x0e20a400, SMAXP, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x0e30a800, SMAXV, instArgs{arg_Vd_22_2__B_0__H_1__S_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__4S_21}, nil},

	{0xbf20fc00, 0x0e206c00, SMIN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x0e20ac00, SMINP, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x0e31a800, SMINV, instArgs{arg_Vd_22_2__B_0__H_1__S_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__4S_21}, nil},

	{0xff00f400, 0x0f002000, SMLAL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x4f002000, SMLAL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x0e208000, SMLAL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e208000, SMLAL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff00f400, 0x0f006000, SMLSL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x4f006000, SMLSL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x0e20a000, SMLSL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e20a000, SMLSL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xffe0fc00, 0x0e002c00, SMOV, instArgs{arg_Wd, arg_Vn_arrangement_imm5___B_1__H_2_index__imm5__imm5lt41gt_1__imm5lt42gt_2_1}, nil},

	{0xffe0fc00, 0x4e002c00, SMOV, instArgs{arg_Xd, arg_Vn_arrangement_imm5___B_1__H_2__S_4_index__imm5__imm5lt41gt_1__imm5lt42gt_2__imm5lt43gt_4_1}, nil},

	{0xff00f400, 0x0f00a000, SMULL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x4f00a000, SMULL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x0e20c000, SMULL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e20c000, SMULL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff3ffc00, 0x5e207800, SQABS, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf3ffc00, 0x0e207800, SQABS, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff20fc00, 0x5e200c00, SQADD, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3, arg_Vm_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf20fc00, 0x0e200c00, SQADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff00f400, 0x5f003000, SQDMLAL, instArgs{arg_Vd_22_2__S_1__D_2, arg_Vn_22_2__H_1__S_2, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x0f003000, SQDMLAL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x4f003000, SQDMLAL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x5e209000, SQDMLAL, instArgs{arg_Vd_22_2__S_1__D_2, arg_Vn_22_2__H_1__S_2, arg_Vm_22_2__H_1__S_2}, nil},

	{0xff20fc00, 0x0e209000, SQDMLAL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e209000, SQDMLAL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff00f400, 0x5f007000, SQDMLSL, instArgs{arg_Vd_22_2__S_1__D_2, arg_Vn_22_2__H_1__S_2, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x0f007000, SQDMLSL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x4f007000, SQDMLSL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x5e20b000, SQDMLSL, instArgs{arg_Vd_22_2__S_1__D_2, arg_Vn_22_2__H_1__S_2, arg_Vm_22_2__H_1__S_2}, nil},

	{0xff20fc00, 0x0e20b000, SQDMLSL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e20b000, SQDMLSL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff00f400, 0x5f00c000, SQDMULH, instArgs{arg_Vd_22_2__H_1__S_2, arg_Vn_22_2__H_1__S_2, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xbf00f400, 0x0f00c000, SQDMULH, instArgs{arg_Vd_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x5e20b400, SQDMULH, instArgs{arg_Vd_22_2__H_1__S_2, arg_Vn_22_2__H_1__S_2, arg_Vm_22_2__H_1__S_2}, nil},

	{0xbf20fc00, 0x0e20b400, SQDMULH, instArgs{arg_Vd_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff00f400, 0x5f00b000, SQDMULL, instArgs{arg_Vd_22_2__S_1__D_2, arg_Vn_22_2__H_1__S_2, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x0f00b000, SQDMULL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x4f00b000, SQDMULL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x5e20d000, SQDMULL, instArgs{arg_Vd_22_2__S_1__D_2, arg_Vn_22_2__H_1__S_2, arg_Vm_22_2__H_1__S_2}, nil},

	{0xff20fc00, 0x0e20d000, SQDMULL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e20d000, SQDMULL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff3ffc00, 0x7e207800, SQNEG, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf3ffc00, 0x2e207800, SQNEG, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff00f400, 0x5f00d000, SQRDMULH, instArgs{arg_Vd_22_2__H_1__S_2, arg_Vn_22_2__H_1__S_2, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xbf00f400, 0x0f00d000, SQRDMULH, instArgs{arg_Vd_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x7e20b400, SQRDMULH, instArgs{arg_Vd_22_2__H_1__S_2, arg_Vn_22_2__H_1__S_2, arg_Vm_22_2__H_1__S_2}, nil},

	{0xbf20fc00, 0x2e20b400, SQRDMULH, instArgs{arg_Vd_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x5e205c00, SQRSHL, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3, arg_Vm_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf20fc00, 0x0e205c00, SQRSHL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff80fc00, 0x5f009c00, SQRSHRN, instArgs{arg_Vd_19_4__B_1__H_2__S_4, arg_Vn_19_4__H_1__S_2__D_4, arg_immediate_1_width_immh_immb__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqrshrn_asisdshf_n_cond},

	{0xff80fc00, 0x0f009c00, SQRSHRN, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqrshrn_asimdshf_n_cond},

	{0xff80fc00, 0x4f009c00, SQRSHRN2, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqrshrn_asimdshf_n_cond},

	{0xff80fc00, 0x7f008c00, SQRSHRUN, instArgs{arg_Vd_19_4__B_1__H_2__S_4, arg_Vn_19_4__H_1__S_2__D_4, arg_immediate_1_width_immh_immb__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqrshrun_asisdshf_n_cond},

	{0xff80fc00, 0x2f008c00, SQRSHRUN, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqrshrun_asimdshf_n_cond},

	{0xff80fc00, 0x6f008c00, SQRSHRUN2, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqrshrun_asimdshf_n_cond},

	{0xff80fc00, 0x5f007400, SQSHL, instArgs{arg_Vd_19_4__B_1__H_2__S_4__D_8, arg_Vn_19_4__B_1__H_2__S_4__D_8, arg_immediate_0_width_m1_immh_immb__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4__UIntimmhimmb64_8}, sqshl_asisdshf_r_cond},

	{0xbf80fc00, 0x0f007400, SQSHL, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4__UIntimmhimmb64_8}, sqshl_asimdshf_r_cond},

	{0xff20fc00, 0x5e204c00, SQSHL, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3, arg_Vm_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf20fc00, 0x0e204c00, SQSHL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff80fc00, 0x7f006400, SQSHLU, instArgs{arg_Vd_19_4__B_1__H_2__S_4__D_8, arg_Vn_19_4__B_1__H_2__S_4__D_8, arg_immediate_0_width_m1_immh_immb__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4__UIntimmhimmb64_8}, sqshlu_asisdshf_r_cond},

	{0xbf80fc00, 0x2f006400, SQSHLU, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4__UIntimmhimmb64_8}, sqshlu_asimdshf_r_cond},

	{0xff80fc00, 0x5f009400, SQSHRN, instArgs{arg_Vd_19_4__B_1__H_2__S_4, arg_Vn_19_4__H_1__S_2__D_4, arg_immediate_1_width_immh_immb__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqshrn_asisdshf_n_cond},

	{0xff80fc00, 0x0f009400, SQSHRN, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqshrn_asimdshf_n_cond},

	{0xff80fc00, 0x4f009400, SQSHRN2, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqshrn_asimdshf_n_cond},

	{0xff80fc00, 0x7f008400, SQSHRUN, instArgs{arg_Vd_19_4__B_1__H_2__S_4, arg_Vn_19_4__H_1__S_2__D_4, arg_immediate_1_width_immh_immb__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqshrun_asisdshf_n_cond},

	{0xff80fc00, 0x2f008400, SQSHRUN, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqshrun_asimdshf_n_cond},

	{0xff80fc00, 0x6f008400, SQSHRUN2, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, sqshrun_asimdshf_n_cond},

	{0xff20fc00, 0x5e202c00, SQSUB, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3, arg_Vm_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf20fc00, 0x0e202c00, SQSUB, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff3ffc00, 0x5e214800, SQXTN, instArgs{arg_Vd_22_2__B_0__H_1__S_2, arg_Vn_22_2__H_0__S_1__D_2}, nil},

	{0xff3ffc00, 0x0e214800, SQXTN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff3ffc00, 0x4e214800, SQXTN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff3ffc00, 0x7e212800, SQXTUN, instArgs{arg_Vd_22_2__B_0__H_1__S_2, arg_Vn_22_2__H_0__S_1__D_2}, nil},

	{0xff3ffc00, 0x2e212800, SQXTUN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff3ffc00, 0x6e212800, SQXTUN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xbf20fc00, 0x0e201400, SRHADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff80fc00, 0x7f004400, SRI, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, sri_asisdshf_r_cond},

	{0xbf80fc00, 0x2f004400, SRI, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, sri_asimdshf_r_cond},

	{0xff20fc00, 0x5e205400, SRSHL, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x0e205400, SRSHL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff80fc00, 0x5f002400, SRSHR, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, srshr_asisdshf_r_cond},

	{0xbf80fc00, 0x0f002400, SRSHR, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, srshr_asimdshf_r_cond},

	{0xff80fc00, 0x5f003400, SRSRA, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, srsra_asisdshf_r_cond},

	{0xbf80fc00, 0x0f003400, SRSRA, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, srsra_asimdshf_r_cond},

	{0xff20fc00, 0x5e204400, SSHL, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x0e204400, SSHL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff87fc00, 0x0f00a400, SXTL, instArgs{arg_Vd_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41}, sxtl_sshll_asimdshf_l_cond},

	{0xff87fc00, 0x4f00a400, SXTL2, instArgs{arg_Vd_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41}, sxtl_sshll_asimdshf_l_cond},

	{0xff80fc00, 0x0f00a400, SSHLL, instArgs{arg_Vd_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4}, sshll_asimdshf_l_cond},

	{0xff80fc00, 0x4f00a400, SSHLL2, instArgs{arg_Vd_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4}, sshll_asimdshf_l_cond},

	{0xff80fc00, 0x5f000400, SSHR, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, sshr_asisdshf_r_cond},

	{0xbf80fc00, 0x0f000400, SSHR, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, sshr_asimdshf_r_cond},

	{0xff80fc00, 0x5f001400, SSRA, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, ssra_asisdshf_r_cond},

	{0xbf80fc00, 0x0f001400, SSRA, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, ssra_asimdshf_r_cond},

	{0xff20fc00, 0x0e202000, SSUBL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e202000, SSUBL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x0e203000, SSUBW, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x4e203000, SSUBW2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbffff000, 0x0c007000, ST1, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c00a000, ST1, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c006000, ST1, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c002000, ST1, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c9f7000, ST1, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Q__8_0__16_1}, nil},

	{0xbfe0f000, 0x0c807000, ST1, instArgs{arg_Vt_1_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c9fa000, ST1, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Q__16_0__32_1}, nil},

	{0xbfe0f000, 0x0c80a000, ST1, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c9f6000, ST1, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Q__24_0__48_1}, nil},

	{0xbfe0f000, 0x0c806000, ST1, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c9f2000, ST1, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Q__32_0__64_1}, nil},

	{0xbfe0f000, 0x0c802000, ST1, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__1D_30__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe000, 0x0d000000, ST1, instArgs{arg_Vt_1_arrangement_B_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffe400, 0x0d004000, ST1, instArgs{arg_Vt_1_arrangement_H_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffec00, 0x0d008000, ST1, instArgs{arg_Vt_1_arrangement_S_index__Q_S_1, arg_Xns_mem_offset}, nil},

	{0xbffffc00, 0x0d008400, ST1, instArgs{arg_Vt_1_arrangement_D_index__Q_1, arg_Xns_mem_offset}, nil},

	{0xbfffe000, 0x0d9f0000, ST1, instArgs{arg_Vt_1_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_1}, nil},

	{0xbfe0e000, 0x0d800000, ST1, instArgs{arg_Vt_1_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe400, 0x0d9f4000, ST1, instArgs{arg_Vt_1_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_2}, nil},

	{0xbfe0e400, 0x0d804000, ST1, instArgs{arg_Vt_1_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffec00, 0x0d9f8000, ST1, instArgs{arg_Vt_1_arrangement_S_index__Q_S_1, arg_Xns_mem_post_fixedimm_4}, nil},

	{0xbfe0ec00, 0x0d808000, ST1, instArgs{arg_Vt_1_arrangement_S_index__Q_S_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffffc00, 0x0d9f8400, ST1, instArgs{arg_Vt_1_arrangement_D_index__Q_1, arg_Xns_mem_post_fixedimm_8}, nil},

	{0xbfe0fc00, 0x0d808400, ST1, instArgs{arg_Vt_1_arrangement_D_index__Q_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c008000, ST2, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c9f8000, ST2, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Q__16_0__32_1}, nil},

	{0xbfe0f000, 0x0c808000, ST2, instArgs{arg_Vt_2_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe000, 0x0d200000, ST2, instArgs{arg_Vt_2_arrangement_B_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffe400, 0x0d204000, ST2, instArgs{arg_Vt_2_arrangement_H_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffec00, 0x0d208000, ST2, instArgs{arg_Vt_2_arrangement_S_index__Q_S_1, arg_Xns_mem_offset}, nil},

	{0xbffffc00, 0x0d208400, ST2, instArgs{arg_Vt_2_arrangement_D_index__Q_1, arg_Xns_mem_offset}, nil},

	{0xbfffe000, 0x0dbf0000, ST2, instArgs{arg_Vt_2_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_2}, nil},

	{0xbfe0e000, 0x0da00000, ST2, instArgs{arg_Vt_2_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe400, 0x0dbf4000, ST2, instArgs{arg_Vt_2_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_4}, nil},

	{0xbfe0e400, 0x0da04000, ST2, instArgs{arg_Vt_2_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffec00, 0x0dbf8000, ST2, instArgs{arg_Vt_2_arrangement_S_index__Q_S_1, arg_Xns_mem_post_fixedimm_8}, nil},

	{0xbfe0ec00, 0x0da08000, ST2, instArgs{arg_Vt_2_arrangement_S_index__Q_S_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffffc00, 0x0dbf8400, ST2, instArgs{arg_Vt_2_arrangement_D_index__Q_1, arg_Xns_mem_post_fixedimm_16}, nil},

	{0xbfe0fc00, 0x0da08400, ST2, instArgs{arg_Vt_2_arrangement_D_index__Q_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c004000, ST3, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c9f4000, ST3, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Q__24_0__48_1}, nil},

	{0xbfe0f000, 0x0c804000, ST3, instArgs{arg_Vt_3_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe000, 0x0d002000, ST3, instArgs{arg_Vt_3_arrangement_B_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffe400, 0x0d006000, ST3, instArgs{arg_Vt_3_arrangement_H_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffec00, 0x0d00a000, ST3, instArgs{arg_Vt_3_arrangement_S_index__Q_S_1, arg_Xns_mem_offset}, nil},

	{0xbffffc00, 0x0d00a400, ST3, instArgs{arg_Vt_3_arrangement_D_index__Q_1, arg_Xns_mem_offset}, nil},

	{0xbfffe000, 0x0d9f2000, ST3, instArgs{arg_Vt_3_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_3}, nil},

	{0xbfe0e000, 0x0d802000, ST3, instArgs{arg_Vt_3_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe400, 0x0d9f6000, ST3, instArgs{arg_Vt_3_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_6}, nil},

	{0xbfe0e400, 0x0d806000, ST3, instArgs{arg_Vt_3_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffec00, 0x0d9fa000, ST3, instArgs{arg_Vt_3_arrangement_S_index__Q_S_1, arg_Xns_mem_post_fixedimm_12}, nil},

	{0xbfe0ec00, 0x0d80a000, ST3, instArgs{arg_Vt_3_arrangement_S_index__Q_S_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffffc00, 0x0d9fa400, ST3, instArgs{arg_Vt_3_arrangement_D_index__Q_1, arg_Xns_mem_post_fixedimm_24}, nil},

	{0xbfe0fc00, 0x0d80a400, ST3, instArgs{arg_Vt_3_arrangement_D_index__Q_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffff000, 0x0c000000, ST4, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_offset}, nil},

	{0xbffff000, 0x0c9f0000, ST4, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Q__32_0__64_1}, nil},

	{0xbfe0f000, 0x0c800000, ST4, instArgs{arg_Vt_4_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe000, 0x0d202000, ST4, instArgs{arg_Vt_4_arrangement_B_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffe400, 0x0d206000, ST4, instArgs{arg_Vt_4_arrangement_H_index__Q_S_size_1, arg_Xns_mem_offset}, nil},

	{0xbfffec00, 0x0d20a000, ST4, instArgs{arg_Vt_4_arrangement_S_index__Q_S_1, arg_Xns_mem_offset}, nil},

	{0xbffffc00, 0x0d20a400, ST4, instArgs{arg_Vt_4_arrangement_D_index__Q_1, arg_Xns_mem_offset}, nil},

	{0xbfffe000, 0x0dbf2000, ST4, instArgs{arg_Vt_4_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_4}, nil},

	{0xbfe0e000, 0x0da02000, ST4, instArgs{arg_Vt_4_arrangement_B_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffe400, 0x0dbf6000, ST4, instArgs{arg_Vt_4_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_fixedimm_8}, nil},

	{0xbfe0e400, 0x0da06000, ST4, instArgs{arg_Vt_4_arrangement_H_index__Q_S_size_1, arg_Xns_mem_post_Xm}, nil},

	{0xbfffec00, 0x0dbfa000, ST4, instArgs{arg_Vt_4_arrangement_S_index__Q_S_1, arg_Xns_mem_post_fixedimm_16}, nil},

	{0xbfe0ec00, 0x0da0a000, ST4, instArgs{arg_Vt_4_arrangement_S_index__Q_S_1, arg_Xns_mem_post_Xm}, nil},

	{0xbffffc00, 0x0dbfa400, ST4, instArgs{arg_Vt_4_arrangement_D_index__Q_1, arg_Xns_mem_post_fixedimm_32}, nil},

	{0xbfe0fc00, 0x0da0a400, ST4, instArgs{arg_Vt_4_arrangement_D_index__Q_1, arg_Xns_mem_post_Xm}, nil},

	{0xffc00000, 0x2c000000, STNP, instArgs{arg_St, arg_St2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffc00000, 0x6c000000, STNP, instArgs{arg_Dt, arg_Dt2, arg_Xns_mem_optional_imm7_8_signed}, nil},

	{0xffc00000, 0xac000000, STNP, instArgs{arg_Qt, arg_Qt2, arg_Xns_mem_optional_imm7_16_signed}, nil},

	{0xffc00000, 0x2c800000, STP, instArgs{arg_St, arg_St2, arg_Xns_mem_post_imm7_4_signed}, nil},

	{0xffc00000, 0x6c800000, STP, instArgs{arg_Dt, arg_Dt2, arg_Xns_mem_post_imm7_8_signed}, nil},

	{0xffc00000, 0xac800000, STP, instArgs{arg_Qt, arg_Qt2, arg_Xns_mem_post_imm7_16_signed}, nil},

	{0xffc00000, 0x2d800000, STP, instArgs{arg_St, arg_St2, arg_Xns_mem_wb_imm7_4_signed}, nil},

	{0xffc00000, 0x6d800000, STP, instArgs{arg_Dt, arg_Dt2, arg_Xns_mem_wb_imm7_8_signed}, nil},

	{0xffc00000, 0xad800000, STP, instArgs{arg_Qt, arg_Qt2, arg_Xns_mem_wb_imm7_16_signed}, nil},

	{0xffc00000, 0x2d000000, STP, instArgs{arg_St, arg_St2, arg_Xns_mem_optional_imm7_4_signed}, nil},

	{0xffc00000, 0x6d000000, STP, instArgs{arg_Dt, arg_Dt2, arg_Xns_mem_optional_imm7_8_signed}, nil},

	{0xffc00000, 0xad000000, STP, instArgs{arg_Qt, arg_Qt2, arg_Xns_mem_optional_imm7_16_signed}, nil},

	{0xffe00c00, 0x3c000400, STR, instArgs{arg_Bt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x7c000400, STR, instArgs{arg_Ht, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xbc000400, STR, instArgs{arg_St, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0xfc000400, STR, instArgs{arg_Dt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x3c800400, STR, instArgs{arg_Qt, arg_Xns_mem_post_imm9_1_signed}, nil},

	{0xffe00c00, 0x3c000c00, STR, instArgs{arg_Bt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0x7c000c00, STR, instArgs{arg_Ht, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0xbc000c00, STR, instArgs{arg_St, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0xfc000c00, STR, instArgs{arg_Dt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffe00c00, 0x3c800c00, STR, instArgs{arg_Qt, arg_Xns_mem_wb_imm9_1_signed}, nil},

	{0xffc00000, 0x3d000000, STR, instArgs{arg_Bt, arg_Xns_mem_optional_imm12_1_unsigned}, nil},

	{0xffc00000, 0x7d000000, STR, instArgs{arg_Ht, arg_Xns_mem_optional_imm12_2_unsigned}, nil},

	{0xffc00000, 0xbd000000, STR, instArgs{arg_St, arg_Xns_mem_optional_imm12_4_unsigned}, nil},

	{0xffc00000, 0xfd000000, STR, instArgs{arg_Dt, arg_Xns_mem_optional_imm12_8_unsigned}, nil},

	{0xffc00000, 0x3d800000, STR, instArgs{arg_Qt, arg_Xns_mem_optional_imm12_16_unsigned}, nil},

	{0xffe00c00, 0x3c200800, STR, instArgs{arg_Bt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__absent_0__0_1}, nil},

	{0xffe00c00, 0x7c200800, STR, instArgs{arg_Ht, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__1_1}, nil},

	{0xffe00c00, 0xbc200800, STR, instArgs{arg_St, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__2_1}, nil},

	{0xffe00c00, 0xfc200800, STR, instArgs{arg_Dt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__3_1}, nil},

	{0xffe00c00, 0x3ca00800, STR, instArgs{arg_Qt, arg_Xns_mem_extend_m__UXTW_2__LSL_3__SXTW_6__SXTX_7__0_0__4_1}, nil},

	{0xffe00c00, 0x3c000000, STUR, instArgs{arg_Bt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x7c000000, STUR, instArgs{arg_Ht, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xbc000000, STUR, instArgs{arg_St, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0xfc000000, STUR, instArgs{arg_Dt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xffe00c00, 0x3c800000, STUR, instArgs{arg_Qt, arg_Xns_mem_optional_imm9_1_signed}, nil},

	{0xff20fc00, 0x7e208400, SUB, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x2e208400, SUB, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff20fc00, 0x0e206000, SUBHN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff20fc00, 0x4e206000, SUBHN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff3ffc00, 0x5e203800, SUQADD, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf3ffc00, 0x0e203800, SUQADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xbfe0fc00, 0x0e002000, TBL, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_2_arrangement_16B, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x0e004000, TBL, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_3_arrangement_16B, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x0e006000, TBL, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_4_arrangement_16B, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x0e000000, TBL, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_1_arrangement_16B, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x0e003000, TBX, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_2_arrangement_16B, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x0e005000, TBX, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_3_arrangement_16B, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x0e007000, TBX, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_4_arrangement_16B, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbfe0fc00, 0x0e001000, TBX, instArgs{arg_Vd_arrangement_Q___8B_0__16B_1, arg_Vn_1_arrangement_16B, arg_Vm_arrangement_Q___8B_0__16B_1}, nil},

	{0xbf20fc00, 0x0e002800, TRN1, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xbf20fc00, 0x0e006800, TRN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xbf20fc00, 0x2e207c00, UABA, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x2e205000, UABAL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e205000, UABAL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x2e207400, UABD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x2e207000, UABDL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e207000, UABDL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x2e206800, UADALP, instArgs{arg_Vd_arrangement_size_Q___4H_00__8H_01__2S_10__4S_11__1D_20__2D_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x2e200000, UADDL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e200000, UADDL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x2e202800, UADDLP, instArgs{arg_Vd_arrangement_size_Q___4H_00__8H_01__2S_10__4S_11__1D_20__2D_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x2e303800, UADDLV, instArgs{arg_Vd_22_2__H_0__S_1__D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__4S_21}, nil},

	{0xff20fc00, 0x2e201000, UADDW, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e201000, UADDW2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xffff0000, 0x1e030000, UCVTF, instArgs{arg_Sd, arg_Wn, arg_immediate_fbits_min_1_max_32_sub_64_scale}, nil},

	{0xffff0000, 0x1e430000, UCVTF, instArgs{arg_Dd, arg_Wn, arg_immediate_fbits_min_1_max_32_sub_64_scale}, nil},

	{0xffff0000, 0x9e030000, UCVTF, instArgs{arg_Sd, arg_Xn, arg_immediate_fbits_min_1_max_64_sub_64_scale}, nil},

	{0xffff0000, 0x9e430000, UCVTF, instArgs{arg_Dd, arg_Xn, arg_immediate_fbits_min_1_max_64_sub_64_scale}, nil},

	{0xfffffc00, 0x1e230000, UCVTF, instArgs{arg_Sd, arg_Wn}, nil},

	{0xfffffc00, 0x1e630000, UCVTF, instArgs{arg_Dd, arg_Wn}, nil},

	{0xfffffc00, 0x9e230000, UCVTF, instArgs{arg_Sd, arg_Xn}, nil},

	{0xfffffc00, 0x9e630000, UCVTF, instArgs{arg_Dd, arg_Xn}, nil},

	{0xff80fc00, 0x7f00e400, UCVTF, instArgs{arg_Vd_19_4__S_4__D_8, arg_Vn_19_4__S_4__D_8, arg_immediate_fbits_min_1_max_0_sub_0_immh_immb__64UIntimmhimmb_4__128UIntimmhimmb_8}, ucvtf_asisdshf_c_cond},

	{0xbf80fc00, 0x2f00e400, UCVTF, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__2S_40__4S_41__2D_81, arg_immediate_fbits_min_1_max_0_sub_0_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__64UIntimmhimmb_4__128UIntimmhimmb_8}, ucvtf_asimdshf_c_cond},

	{0xffbffc00, 0x7e21d800, UCVTF, instArgs{arg_Vd_22_1__S_0__D_1, arg_Vn_22_1__S_0__D_1}, nil},

	{0xbfbffc00, 0x2e21d800, UCVTF, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01__2D_11, arg_Vn_arrangement_sz_Q___2S_00__4S_01__2D_11}, nil},

	{0xbf20fc00, 0x2e200400, UHADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x2e202400, UHSUB, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x2e206400, UMAX, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x2e20a400, UMAXP, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x2e30a800, UMAXV, instArgs{arg_Vd_22_2__B_0__H_1__S_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__4S_21}, nil},

	{0xbf20fc00, 0x2e206c00, UMIN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x2e20ac00, UMINP, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf3ffc00, 0x2e31a800, UMINV, instArgs{arg_Vd_22_2__B_0__H_1__S_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__4S_21}, nil},

	{0xff00f400, 0x2f002000, UMLAL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x6f002000, UMLAL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x2e208000, UMLAL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e208000, UMLAL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff00f400, 0x2f006000, UMLSL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x6f006000, UMLSL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x2e20a000, UMLSL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e20a000, UMLSL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff00f400, 0x2f00a000, UMULL, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff00f400, 0x6f00a000, UMULL2, instArgs{arg_Vd_arrangement_size___4S_1__2D_2, arg_Vn_arrangement_size_Q___4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size___H_1__S_2_index__size_L_H_M__HLM_1__HL_2_1}, nil},

	{0xff20fc00, 0x2e20c000, UMULL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e20c000, UMULL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x7e200c00, UQADD, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3, arg_Vm_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf20fc00, 0x2e200c00, UQADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff20fc00, 0x7e205c00, UQRSHL, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3, arg_Vm_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf20fc00, 0x2e205c00, UQRSHL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff80fc00, 0x7f009c00, UQRSHRN, instArgs{arg_Vd_19_4__B_1__H_2__S_4, arg_Vn_19_4__H_1__S_2__D_4, arg_immediate_1_width_immh_immb__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, uqrshrn_asisdshf_n_cond},

	{0xff80fc00, 0x2f009c00, UQRSHRN, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, uqrshrn_asimdshf_n_cond},

	{0xff80fc00, 0x6f009c00, UQRSHRN2, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, uqrshrn_asimdshf_n_cond},

	{0xff80fc00, 0x7f007400, UQSHL, instArgs{arg_Vd_19_4__B_1__H_2__S_4__D_8, arg_Vn_19_4__B_1__H_2__S_4__D_8, arg_immediate_0_width_m1_immh_immb__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4__UIntimmhimmb64_8}, uqshl_asisdshf_r_cond},

	{0xbf80fc00, 0x2f007400, UQSHL, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4__UIntimmhimmb64_8}, uqshl_asimdshf_r_cond},

	{0xff20fc00, 0x7e204c00, UQSHL, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3, arg_Vm_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf20fc00, 0x2e204c00, UQSHL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff80fc00, 0x7f009400, UQSHRN, instArgs{arg_Vd_19_4__B_1__H_2__S_4, arg_Vn_19_4__H_1__S_2__D_4, arg_immediate_1_width_immh_immb__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, uqshrn_asisdshf_n_cond},

	{0xff80fc00, 0x2f009400, UQSHRN, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, uqshrn_asimdshf_n_cond},

	{0xff80fc00, 0x6f009400, UQSHRN2, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_Vn_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4}, uqshrn_asimdshf_n_cond},

	{0xff20fc00, 0x7e202c00, UQSUB, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3, arg_Vm_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf20fc00, 0x2e202c00, UQSUB, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff3ffc00, 0x7e214800, UQXTN, instArgs{arg_Vd_22_2__B_0__H_1__S_2, arg_Vn_22_2__H_0__S_1__D_2}, nil},

	{0xff3ffc00, 0x2e214800, UQXTN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff3ffc00, 0x6e214800, UQXTN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xbfbffc00, 0x0ea1c800, URECPE, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01, arg_Vn_arrangement_sz_Q___2S_00__4S_01}, nil},

	{0xbf20fc00, 0x2e201400, URHADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x7e205400, URSHL, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x2e205400, URSHL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff80fc00, 0x7f002400, URSHR, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, urshr_asisdshf_r_cond},

	{0xbf80fc00, 0x2f002400, URSHR, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, urshr_asimdshf_r_cond},

	{0xbfbffc00, 0x2ea1c800, URSQRTE, instArgs{arg_Vd_arrangement_sz_Q___2S_00__4S_01, arg_Vn_arrangement_sz_Q___2S_00__4S_01}, nil},

	{0xff80fc00, 0x7f003400, URSRA, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, ursra_asisdshf_r_cond},

	{0xbf80fc00, 0x2f003400, URSRA, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, ursra_asimdshf_r_cond},

	{0xff20fc00, 0x7e204400, USHL, instArgs{arg_Vd_22_2__D_3, arg_Vn_22_2__D_3, arg_Vm_22_2__D_3}, nil},

	{0xbf20fc00, 0x2e204400, USHL, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff87fc00, 0x2f00a400, UXTL, instArgs{arg_Vd_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41}, uxtl_ushll_asimdshf_l_cond},

	{0xff87fc00, 0x6f00a400, UXTL2, instArgs{arg_Vd_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41}, uxtl_ushll_asimdshf_l_cond},

	{0xff80fc00, 0x2f00a400, USHLL, instArgs{arg_Vd_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4}, ushll_asimdshf_l_cond},

	{0xff80fc00, 0x6f00a400, USHLL2, instArgs{arg_Vd_arrangement_immh___SEEAdvancedSIMDmodifiedimmediate_0__8H_1__4S_2__2D_4, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41, arg_immediate_0_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__UIntimmhimmb8_1__UIntimmhimmb16_2__UIntimmhimmb32_4}, ushll_asimdshf_l_cond},

	{0xff80fc00, 0x7f000400, USHR, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, ushr_asisdshf_r_cond},

	{0xbf80fc00, 0x2f000400, USHR, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, ushr_asimdshf_r_cond},

	{0xff3ffc00, 0x7e203800, USQADD, instArgs{arg_Vd_22_2__B_0__H_1__S_2__D_3, arg_Vn_22_2__B_0__H_1__S_2__D_3}, nil},

	{0xbf3ffc00, 0x2e203800, USQADD, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff80fc00, 0x7f001400, USRA, instArgs{arg_Vd_19_4__D_8, arg_Vn_19_4__D_8, arg_immediate_1_64_immh_immb__128UIntimmhimmb_8}, usra_asisdshf_r_cond},

	{0xbf80fc00, 0x2f001400, USRA, instArgs{arg_Vd_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_Vn_arrangement_immh_Q___SEEAdvancedSIMDmodifiedimmediate_00__8B_10__16B_11__4H_20__8H_21__2S_40__4S_41__2D_81, arg_immediate_1_width_immh_immb__SEEAdvancedSIMDmodifiedimmediate_0__16UIntimmhimmb_1__32UIntimmhimmb_2__64UIntimmhimmb_4__128UIntimmhimmb_8}, usra_asimdshf_r_cond},

	{0xff20fc00, 0x2e202000, USUBL, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e202000, USUBL2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x2e203000, USUBW, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xff20fc00, 0x6e203000, USUBW2, instArgs{arg_Vd_arrangement_size___8H_0__4S_1__2D_2, arg_Vn_arrangement_size___8H_0__4S_1__2D_2, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21}, nil},

	{0xbf20fc00, 0x0e001800, UZP1, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xbf20fc00, 0x0e005800, UZP2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xff3ffc00, 0x0e212800, XTN, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xff3ffc00, 0x4e212800, XTN2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21, arg_Vn_arrangement_size___8H_0__4S_1__2D_2}, nil},

	{0xbf20fc00, 0x0e003800, ZIP1, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},

	{0xbf20fc00, 0x0e007800, ZIP2, instArgs{arg_Vd_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vn_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31, arg_Vm_arrangement_size_Q___8B_00__16B_01__4H_10__8H_11__2S_20__4S_21__2D_31}, nil},
}
