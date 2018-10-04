// Code generated from gen/*Ops.go; DO NOT EDIT.

package ssa

import (
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/arm"
	"github.com/dave/golib/src/cmd/internal/obj/arm64"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
	"github.com/dave/golib/src/cmd/internal/obj/ppc64"
	"github.com/dave/golib/src/cmd/internal/obj/s390x"
	"github.com/dave/golib/src/cmd/internal/obj/wasm"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
)

const (
	BlockInvalid BlockKind = iota

	Block386EQ
	Block386NE
	Block386LT
	Block386LE
	Block386GT
	Block386GE
	Block386ULT
	Block386ULE
	Block386UGT
	Block386UGE
	Block386EQF
	Block386NEF
	Block386ORD
	Block386NAN

	BlockAMD64EQ
	BlockAMD64NE
	BlockAMD64LT
	BlockAMD64LE
	BlockAMD64GT
	BlockAMD64GE
	BlockAMD64ULT
	BlockAMD64ULE
	BlockAMD64UGT
	BlockAMD64UGE
	BlockAMD64EQF
	BlockAMD64NEF
	BlockAMD64ORD
	BlockAMD64NAN

	BlockARMEQ
	BlockARMNE
	BlockARMLT
	BlockARMLE
	BlockARMGT
	BlockARMGE
	BlockARMULT
	BlockARMULE
	BlockARMUGT
	BlockARMUGE

	BlockARM64EQ
	BlockARM64NE
	BlockARM64LT
	BlockARM64LE
	BlockARM64GT
	BlockARM64GE
	BlockARM64ULT
	BlockARM64ULE
	BlockARM64UGT
	BlockARM64UGE
	BlockARM64Z
	BlockARM64NZ
	BlockARM64ZW
	BlockARM64NZW
	BlockARM64TBZ
	BlockARM64TBNZ

	BlockMIPSEQ
	BlockMIPSNE
	BlockMIPSLTZ
	BlockMIPSLEZ
	BlockMIPSGTZ
	BlockMIPSGEZ
	BlockMIPSFPT
	BlockMIPSFPF

	BlockMIPS64EQ
	BlockMIPS64NE
	BlockMIPS64LTZ
	BlockMIPS64LEZ
	BlockMIPS64GTZ
	BlockMIPS64GEZ
	BlockMIPS64FPT
	BlockMIPS64FPF

	BlockPPC64EQ
	BlockPPC64NE
	BlockPPC64LT
	BlockPPC64LE
	BlockPPC64GT
	BlockPPC64GE
	BlockPPC64FLT
	BlockPPC64FLE
	BlockPPC64FGT
	BlockPPC64FGE

	BlockS390XEQ
	BlockS390XNE
	BlockS390XLT
	BlockS390XLE
	BlockS390XGT
	BlockS390XGE
	BlockS390XGTF
	BlockS390XGEF

	BlockPlain
	BlockIf
	BlockDefer
	BlockRet
	BlockRetJmp
	BlockExit
	BlockFirst
)

func (k BlockKind) String(pstate *PackageState) string { return pstate.blockString[k] }

const (
	OpInvalid Op = iota

	Op386ADDSS
	Op386ADDSD
	Op386SUBSS
	Op386SUBSD
	Op386MULSS
	Op386MULSD
	Op386DIVSS
	Op386DIVSD
	Op386MOVSSload
	Op386MOVSDload
	Op386MOVSSconst
	Op386MOVSDconst
	Op386MOVSSloadidx1
	Op386MOVSSloadidx4
	Op386MOVSDloadidx1
	Op386MOVSDloadidx8
	Op386MOVSSstore
	Op386MOVSDstore
	Op386MOVSSstoreidx1
	Op386MOVSSstoreidx4
	Op386MOVSDstoreidx1
	Op386MOVSDstoreidx8
	Op386ADDSSload
	Op386ADDSDload
	Op386SUBSSload
	Op386SUBSDload
	Op386MULSSload
	Op386MULSDload
	Op386ADDL
	Op386ADDLconst
	Op386ADDLcarry
	Op386ADDLconstcarry
	Op386ADCL
	Op386ADCLconst
	Op386SUBL
	Op386SUBLconst
	Op386SUBLcarry
	Op386SUBLconstcarry
	Op386SBBL
	Op386SBBLconst
	Op386MULL
	Op386MULLconst
	Op386HMULL
	Op386HMULLU
	Op386MULLQU
	Op386AVGLU
	Op386DIVL
	Op386DIVW
	Op386DIVLU
	Op386DIVWU
	Op386MODL
	Op386MODW
	Op386MODLU
	Op386MODWU
	Op386ANDL
	Op386ANDLconst
	Op386ORL
	Op386ORLconst
	Op386XORL
	Op386XORLconst
	Op386CMPL
	Op386CMPW
	Op386CMPB
	Op386CMPLconst
	Op386CMPWconst
	Op386CMPBconst
	Op386UCOMISS
	Op386UCOMISD
	Op386TESTL
	Op386TESTW
	Op386TESTB
	Op386TESTLconst
	Op386TESTWconst
	Op386TESTBconst
	Op386SHLL
	Op386SHLLconst
	Op386SHRL
	Op386SHRW
	Op386SHRB
	Op386SHRLconst
	Op386SHRWconst
	Op386SHRBconst
	Op386SARL
	Op386SARW
	Op386SARB
	Op386SARLconst
	Op386SARWconst
	Op386SARBconst
	Op386ROLLconst
	Op386ROLWconst
	Op386ROLBconst
	Op386ADDLload
	Op386SUBLload
	Op386ANDLload
	Op386ORLload
	Op386XORLload
	Op386NEGL
	Op386NOTL
	Op386BSFL
	Op386BSFW
	Op386BSRL
	Op386BSRW
	Op386BSWAPL
	Op386SQRTSD
	Op386SBBLcarrymask
	Op386SETEQ
	Op386SETNE
	Op386SETL
	Op386SETLE
	Op386SETG
	Op386SETGE
	Op386SETB
	Op386SETBE
	Op386SETA
	Op386SETAE
	Op386SETEQF
	Op386SETNEF
	Op386SETORD
	Op386SETNAN
	Op386SETGF
	Op386SETGEF
	Op386MOVBLSX
	Op386MOVBLZX
	Op386MOVWLSX
	Op386MOVWLZX
	Op386MOVLconst
	Op386CVTTSD2SL
	Op386CVTTSS2SL
	Op386CVTSL2SS
	Op386CVTSL2SD
	Op386CVTSD2SS
	Op386CVTSS2SD
	Op386PXOR
	Op386LEAL
	Op386LEAL1
	Op386LEAL2
	Op386LEAL4
	Op386LEAL8
	Op386MOVBload
	Op386MOVBLSXload
	Op386MOVWload
	Op386MOVWLSXload
	Op386MOVLload
	Op386MOVBstore
	Op386MOVWstore
	Op386MOVLstore
	Op386ADDLmodify
	Op386SUBLmodify
	Op386ANDLmodify
	Op386ORLmodify
	Op386XORLmodify
	Op386MOVBloadidx1
	Op386MOVWloadidx1
	Op386MOVWloadidx2
	Op386MOVLloadidx1
	Op386MOVLloadidx4
	Op386MOVBstoreidx1
	Op386MOVWstoreidx1
	Op386MOVWstoreidx2
	Op386MOVLstoreidx1
	Op386MOVLstoreidx4
	Op386MOVBstoreconst
	Op386MOVWstoreconst
	Op386MOVLstoreconst
	Op386MOVBstoreconstidx1
	Op386MOVWstoreconstidx1
	Op386MOVWstoreconstidx2
	Op386MOVLstoreconstidx1
	Op386MOVLstoreconstidx4
	Op386DUFFZERO
	Op386REPSTOSL
	Op386CALLstatic
	Op386CALLclosure
	Op386CALLinter
	Op386DUFFCOPY
	Op386REPMOVSL
	Op386InvertFlags
	Op386LoweredGetG
	Op386LoweredGetClosurePtr
	Op386LoweredGetCallerPC
	Op386LoweredGetCallerSP
	Op386LoweredNilCheck
	Op386LoweredWB
	Op386FlagEQ
	Op386FlagLT_ULT
	Op386FlagLT_UGT
	Op386FlagGT_UGT
	Op386FlagGT_ULT
	Op386FCHS
	Op386MOVSSconst1
	Op386MOVSDconst1
	Op386MOVSSconst2
	Op386MOVSDconst2

	OpAMD64ADDSS
	OpAMD64ADDSD
	OpAMD64SUBSS
	OpAMD64SUBSD
	OpAMD64MULSS
	OpAMD64MULSD
	OpAMD64DIVSS
	OpAMD64DIVSD
	OpAMD64MOVSSload
	OpAMD64MOVSDload
	OpAMD64MOVSSconst
	OpAMD64MOVSDconst
	OpAMD64MOVSSloadidx1
	OpAMD64MOVSSloadidx4
	OpAMD64MOVSDloadidx1
	OpAMD64MOVSDloadidx8
	OpAMD64MOVSSstore
	OpAMD64MOVSDstore
	OpAMD64MOVSSstoreidx1
	OpAMD64MOVSSstoreidx4
	OpAMD64MOVSDstoreidx1
	OpAMD64MOVSDstoreidx8
	OpAMD64ADDSSload
	OpAMD64ADDSDload
	OpAMD64SUBSSload
	OpAMD64SUBSDload
	OpAMD64MULSSload
	OpAMD64MULSDload
	OpAMD64ADDQ
	OpAMD64ADDL
	OpAMD64ADDQconst
	OpAMD64ADDLconst
	OpAMD64ADDQconstmodify
	OpAMD64ADDLconstmodify
	OpAMD64SUBQ
	OpAMD64SUBL
	OpAMD64SUBQconst
	OpAMD64SUBLconst
	OpAMD64MULQ
	OpAMD64MULL
	OpAMD64MULQconst
	OpAMD64MULLconst
	OpAMD64HMULQ
	OpAMD64HMULL
	OpAMD64HMULQU
	OpAMD64HMULLU
	OpAMD64AVGQU
	OpAMD64DIVQ
	OpAMD64DIVL
	OpAMD64DIVW
	OpAMD64DIVQU
	OpAMD64DIVLU
	OpAMD64DIVWU
	OpAMD64MULQU2
	OpAMD64DIVQU2
	OpAMD64ANDQ
	OpAMD64ANDL
	OpAMD64ANDQconst
	OpAMD64ANDLconst
	OpAMD64ORQ
	OpAMD64ORL
	OpAMD64ORQconst
	OpAMD64ORLconst
	OpAMD64XORQ
	OpAMD64XORL
	OpAMD64XORQconst
	OpAMD64XORLconst
	OpAMD64CMPQ
	OpAMD64CMPL
	OpAMD64CMPW
	OpAMD64CMPB
	OpAMD64CMPQconst
	OpAMD64CMPLconst
	OpAMD64CMPWconst
	OpAMD64CMPBconst
	OpAMD64CMPQload
	OpAMD64CMPLload
	OpAMD64CMPWload
	OpAMD64CMPBload
	OpAMD64CMPQconstload
	OpAMD64CMPLconstload
	OpAMD64CMPWconstload
	OpAMD64CMPBconstload
	OpAMD64UCOMISS
	OpAMD64UCOMISD
	OpAMD64BTL
	OpAMD64BTQ
	OpAMD64BTCL
	OpAMD64BTCQ
	OpAMD64BTRL
	OpAMD64BTRQ
	OpAMD64BTSL
	OpAMD64BTSQ
	OpAMD64BTLconst
	OpAMD64BTQconst
	OpAMD64BTCLconst
	OpAMD64BTCQconst
	OpAMD64BTRLconst
	OpAMD64BTRQconst
	OpAMD64BTSLconst
	OpAMD64BTSQconst
	OpAMD64TESTQ
	OpAMD64TESTL
	OpAMD64TESTW
	OpAMD64TESTB
	OpAMD64TESTQconst
	OpAMD64TESTLconst
	OpAMD64TESTWconst
	OpAMD64TESTBconst
	OpAMD64SHLQ
	OpAMD64SHLL
	OpAMD64SHLQconst
	OpAMD64SHLLconst
	OpAMD64SHRQ
	OpAMD64SHRL
	OpAMD64SHRW
	OpAMD64SHRB
	OpAMD64SHRQconst
	OpAMD64SHRLconst
	OpAMD64SHRWconst
	OpAMD64SHRBconst
	OpAMD64SARQ
	OpAMD64SARL
	OpAMD64SARW
	OpAMD64SARB
	OpAMD64SARQconst
	OpAMD64SARLconst
	OpAMD64SARWconst
	OpAMD64SARBconst
	OpAMD64ROLQ
	OpAMD64ROLL
	OpAMD64ROLW
	OpAMD64ROLB
	OpAMD64RORQ
	OpAMD64RORL
	OpAMD64RORW
	OpAMD64RORB
	OpAMD64ROLQconst
	OpAMD64ROLLconst
	OpAMD64ROLWconst
	OpAMD64ROLBconst
	OpAMD64ADDLload
	OpAMD64ADDQload
	OpAMD64SUBQload
	OpAMD64SUBLload
	OpAMD64ANDLload
	OpAMD64ANDQload
	OpAMD64ORQload
	OpAMD64ORLload
	OpAMD64XORQload
	OpAMD64XORLload
	OpAMD64NEGQ
	OpAMD64NEGL
	OpAMD64NOTQ
	OpAMD64NOTL
	OpAMD64BSFQ
	OpAMD64BSFL
	OpAMD64BSRQ
	OpAMD64BSRL
	OpAMD64CMOVQEQ
	OpAMD64CMOVQNE
	OpAMD64CMOVQLT
	OpAMD64CMOVQGT
	OpAMD64CMOVQLE
	OpAMD64CMOVQGE
	OpAMD64CMOVQLS
	OpAMD64CMOVQHI
	OpAMD64CMOVQCC
	OpAMD64CMOVQCS
	OpAMD64CMOVLEQ
	OpAMD64CMOVLNE
	OpAMD64CMOVLLT
	OpAMD64CMOVLGT
	OpAMD64CMOVLLE
	OpAMD64CMOVLGE
	OpAMD64CMOVLLS
	OpAMD64CMOVLHI
	OpAMD64CMOVLCC
	OpAMD64CMOVLCS
	OpAMD64CMOVWEQ
	OpAMD64CMOVWNE
	OpAMD64CMOVWLT
	OpAMD64CMOVWGT
	OpAMD64CMOVWLE
	OpAMD64CMOVWGE
	OpAMD64CMOVWLS
	OpAMD64CMOVWHI
	OpAMD64CMOVWCC
	OpAMD64CMOVWCS
	OpAMD64CMOVQEQF
	OpAMD64CMOVQNEF
	OpAMD64CMOVQGTF
	OpAMD64CMOVQGEF
	OpAMD64CMOVLEQF
	OpAMD64CMOVLNEF
	OpAMD64CMOVLGTF
	OpAMD64CMOVLGEF
	OpAMD64CMOVWEQF
	OpAMD64CMOVWNEF
	OpAMD64CMOVWGTF
	OpAMD64CMOVWGEF
	OpAMD64BSWAPQ
	OpAMD64BSWAPL
	OpAMD64POPCNTQ
	OpAMD64POPCNTL
	OpAMD64SQRTSD
	OpAMD64ROUNDSD
	OpAMD64SBBQcarrymask
	OpAMD64SBBLcarrymask
	OpAMD64SETEQ
	OpAMD64SETNE
	OpAMD64SETL
	OpAMD64SETLE
	OpAMD64SETG
	OpAMD64SETGE
	OpAMD64SETB
	OpAMD64SETBE
	OpAMD64SETA
	OpAMD64SETAE
	OpAMD64SETEQstore
	OpAMD64SETNEstore
	OpAMD64SETLstore
	OpAMD64SETLEstore
	OpAMD64SETGstore
	OpAMD64SETGEstore
	OpAMD64SETBstore
	OpAMD64SETBEstore
	OpAMD64SETAstore
	OpAMD64SETAEstore
	OpAMD64SETEQF
	OpAMD64SETNEF
	OpAMD64SETORD
	OpAMD64SETNAN
	OpAMD64SETGF
	OpAMD64SETGEF
	OpAMD64MOVBQSX
	OpAMD64MOVBQZX
	OpAMD64MOVWQSX
	OpAMD64MOVWQZX
	OpAMD64MOVLQSX
	OpAMD64MOVLQZX
	OpAMD64MOVLconst
	OpAMD64MOVQconst
	OpAMD64CVTTSD2SL
	OpAMD64CVTTSD2SQ
	OpAMD64CVTTSS2SL
	OpAMD64CVTTSS2SQ
	OpAMD64CVTSL2SS
	OpAMD64CVTSL2SD
	OpAMD64CVTSQ2SS
	OpAMD64CVTSQ2SD
	OpAMD64CVTSD2SS
	OpAMD64CVTSS2SD
	OpAMD64MOVQi2f
	OpAMD64MOVQf2i
	OpAMD64MOVLi2f
	OpAMD64MOVLf2i
	OpAMD64PXOR
	OpAMD64LEAQ
	OpAMD64LEAL
	OpAMD64LEAW
	OpAMD64LEAQ1
	OpAMD64LEAL1
	OpAMD64LEAW1
	OpAMD64LEAQ2
	OpAMD64LEAL2
	OpAMD64LEAW2
	OpAMD64LEAQ4
	OpAMD64LEAL4
	OpAMD64LEAW4
	OpAMD64LEAQ8
	OpAMD64LEAL8
	OpAMD64LEAW8
	OpAMD64MOVBload
	OpAMD64MOVBQSXload
	OpAMD64MOVWload
	OpAMD64MOVWQSXload
	OpAMD64MOVLload
	OpAMD64MOVLQSXload
	OpAMD64MOVQload
	OpAMD64MOVBstore
	OpAMD64MOVWstore
	OpAMD64MOVLstore
	OpAMD64MOVQstore
	OpAMD64MOVOload
	OpAMD64MOVOstore
	OpAMD64MOVBloadidx1
	OpAMD64MOVWloadidx1
	OpAMD64MOVWloadidx2
	OpAMD64MOVLloadidx1
	OpAMD64MOVLloadidx4
	OpAMD64MOVLloadidx8
	OpAMD64MOVQloadidx1
	OpAMD64MOVQloadidx8
	OpAMD64MOVBstoreidx1
	OpAMD64MOVWstoreidx1
	OpAMD64MOVWstoreidx2
	OpAMD64MOVLstoreidx1
	OpAMD64MOVLstoreidx4
	OpAMD64MOVLstoreidx8
	OpAMD64MOVQstoreidx1
	OpAMD64MOVQstoreidx8
	OpAMD64MOVBstoreconst
	OpAMD64MOVWstoreconst
	OpAMD64MOVLstoreconst
	OpAMD64MOVQstoreconst
	OpAMD64MOVBstoreconstidx1
	OpAMD64MOVWstoreconstidx1
	OpAMD64MOVWstoreconstidx2
	OpAMD64MOVLstoreconstidx1
	OpAMD64MOVLstoreconstidx4
	OpAMD64MOVQstoreconstidx1
	OpAMD64MOVQstoreconstidx8
	OpAMD64DUFFZERO
	OpAMD64MOVOconst
	OpAMD64REPSTOSQ
	OpAMD64CALLstatic
	OpAMD64CALLclosure
	OpAMD64CALLinter
	OpAMD64DUFFCOPY
	OpAMD64REPMOVSQ
	OpAMD64InvertFlags
	OpAMD64LoweredGetG
	OpAMD64LoweredGetClosurePtr
	OpAMD64LoweredGetCallerPC
	OpAMD64LoweredGetCallerSP
	OpAMD64LoweredNilCheck
	OpAMD64LoweredWB
	OpAMD64FlagEQ
	OpAMD64FlagLT_ULT
	OpAMD64FlagLT_UGT
	OpAMD64FlagGT_UGT
	OpAMD64FlagGT_ULT
	OpAMD64MOVLatomicload
	OpAMD64MOVQatomicload
	OpAMD64XCHGL
	OpAMD64XCHGQ
	OpAMD64XADDLlock
	OpAMD64XADDQlock
	OpAMD64AddTupleFirst32
	OpAMD64AddTupleFirst64
	OpAMD64CMPXCHGLlock
	OpAMD64CMPXCHGQlock
	OpAMD64ANDBlock
	OpAMD64ORBlock

	OpARMADD
	OpARMADDconst
	OpARMSUB
	OpARMSUBconst
	OpARMRSB
	OpARMRSBconst
	OpARMMUL
	OpARMHMUL
	OpARMHMULU
	OpARMCALLudiv
	OpARMADDS
	OpARMADDSconst
	OpARMADC
	OpARMADCconst
	OpARMSUBS
	OpARMSUBSconst
	OpARMRSBSconst
	OpARMSBC
	OpARMSBCconst
	OpARMRSCconst
	OpARMMULLU
	OpARMMULA
	OpARMMULS
	OpARMADDF
	OpARMADDD
	OpARMSUBF
	OpARMSUBD
	OpARMMULF
	OpARMMULD
	OpARMNMULF
	OpARMNMULD
	OpARMDIVF
	OpARMDIVD
	OpARMMULAF
	OpARMMULAD
	OpARMMULSF
	OpARMMULSD
	OpARMAND
	OpARMANDconst
	OpARMOR
	OpARMORconst
	OpARMXOR
	OpARMXORconst
	OpARMBIC
	OpARMBICconst
	OpARMBFX
	OpARMBFXU
	OpARMMVN
	OpARMNEGF
	OpARMNEGD
	OpARMSQRTD
	OpARMCLZ
	OpARMREV
	OpARMRBIT
	OpARMSLL
	OpARMSLLconst
	OpARMSRL
	OpARMSRLconst
	OpARMSRA
	OpARMSRAconst
	OpARMSRRconst
	OpARMADDshiftLL
	OpARMADDshiftRL
	OpARMADDshiftRA
	OpARMSUBshiftLL
	OpARMSUBshiftRL
	OpARMSUBshiftRA
	OpARMRSBshiftLL
	OpARMRSBshiftRL
	OpARMRSBshiftRA
	OpARMANDshiftLL
	OpARMANDshiftRL
	OpARMANDshiftRA
	OpARMORshiftLL
	OpARMORshiftRL
	OpARMORshiftRA
	OpARMXORshiftLL
	OpARMXORshiftRL
	OpARMXORshiftRA
	OpARMXORshiftRR
	OpARMBICshiftLL
	OpARMBICshiftRL
	OpARMBICshiftRA
	OpARMMVNshiftLL
	OpARMMVNshiftRL
	OpARMMVNshiftRA
	OpARMADCshiftLL
	OpARMADCshiftRL
	OpARMADCshiftRA
	OpARMSBCshiftLL
	OpARMSBCshiftRL
	OpARMSBCshiftRA
	OpARMRSCshiftLL
	OpARMRSCshiftRL
	OpARMRSCshiftRA
	OpARMADDSshiftLL
	OpARMADDSshiftRL
	OpARMADDSshiftRA
	OpARMSUBSshiftLL
	OpARMSUBSshiftRL
	OpARMSUBSshiftRA
	OpARMRSBSshiftLL
	OpARMRSBSshiftRL
	OpARMRSBSshiftRA
	OpARMADDshiftLLreg
	OpARMADDshiftRLreg
	OpARMADDshiftRAreg
	OpARMSUBshiftLLreg
	OpARMSUBshiftRLreg
	OpARMSUBshiftRAreg
	OpARMRSBshiftLLreg
	OpARMRSBshiftRLreg
	OpARMRSBshiftRAreg
	OpARMANDshiftLLreg
	OpARMANDshiftRLreg
	OpARMANDshiftRAreg
	OpARMORshiftLLreg
	OpARMORshiftRLreg
	OpARMORshiftRAreg
	OpARMXORshiftLLreg
	OpARMXORshiftRLreg
	OpARMXORshiftRAreg
	OpARMBICshiftLLreg
	OpARMBICshiftRLreg
	OpARMBICshiftRAreg
	OpARMMVNshiftLLreg
	OpARMMVNshiftRLreg
	OpARMMVNshiftRAreg
	OpARMADCshiftLLreg
	OpARMADCshiftRLreg
	OpARMADCshiftRAreg
	OpARMSBCshiftLLreg
	OpARMSBCshiftRLreg
	OpARMSBCshiftRAreg
	OpARMRSCshiftLLreg
	OpARMRSCshiftRLreg
	OpARMRSCshiftRAreg
	OpARMADDSshiftLLreg
	OpARMADDSshiftRLreg
	OpARMADDSshiftRAreg
	OpARMSUBSshiftLLreg
	OpARMSUBSshiftRLreg
	OpARMSUBSshiftRAreg
	OpARMRSBSshiftLLreg
	OpARMRSBSshiftRLreg
	OpARMRSBSshiftRAreg
	OpARMCMP
	OpARMCMPconst
	OpARMCMN
	OpARMCMNconst
	OpARMTST
	OpARMTSTconst
	OpARMTEQ
	OpARMTEQconst
	OpARMCMPF
	OpARMCMPD
	OpARMCMPshiftLL
	OpARMCMPshiftRL
	OpARMCMPshiftRA
	OpARMCMNshiftLL
	OpARMCMNshiftRL
	OpARMCMNshiftRA
	OpARMTSTshiftLL
	OpARMTSTshiftRL
	OpARMTSTshiftRA
	OpARMTEQshiftLL
	OpARMTEQshiftRL
	OpARMTEQshiftRA
	OpARMCMPshiftLLreg
	OpARMCMPshiftRLreg
	OpARMCMPshiftRAreg
	OpARMCMNshiftLLreg
	OpARMCMNshiftRLreg
	OpARMCMNshiftRAreg
	OpARMTSTshiftLLreg
	OpARMTSTshiftRLreg
	OpARMTSTshiftRAreg
	OpARMTEQshiftLLreg
	OpARMTEQshiftRLreg
	OpARMTEQshiftRAreg
	OpARMCMPF0
	OpARMCMPD0
	OpARMMOVWconst
	OpARMMOVFconst
	OpARMMOVDconst
	OpARMMOVWaddr
	OpARMMOVBload
	OpARMMOVBUload
	OpARMMOVHload
	OpARMMOVHUload
	OpARMMOVWload
	OpARMMOVFload
	OpARMMOVDload
	OpARMMOVBstore
	OpARMMOVHstore
	OpARMMOVWstore
	OpARMMOVFstore
	OpARMMOVDstore
	OpARMMOVWloadidx
	OpARMMOVWloadshiftLL
	OpARMMOVWloadshiftRL
	OpARMMOVWloadshiftRA
	OpARMMOVBUloadidx
	OpARMMOVBloadidx
	OpARMMOVHUloadidx
	OpARMMOVHloadidx
	OpARMMOVWstoreidx
	OpARMMOVWstoreshiftLL
	OpARMMOVWstoreshiftRL
	OpARMMOVWstoreshiftRA
	OpARMMOVBstoreidx
	OpARMMOVHstoreidx
	OpARMMOVBreg
	OpARMMOVBUreg
	OpARMMOVHreg
	OpARMMOVHUreg
	OpARMMOVWreg
	OpARMMOVWnop
	OpARMMOVWF
	OpARMMOVWD
	OpARMMOVWUF
	OpARMMOVWUD
	OpARMMOVFW
	OpARMMOVDW
	OpARMMOVFWU
	OpARMMOVDWU
	OpARMMOVFD
	OpARMMOVDF
	OpARMCMOVWHSconst
	OpARMCMOVWLSconst
	OpARMSRAcond
	OpARMCALLstatic
	OpARMCALLclosure
	OpARMCALLinter
	OpARMLoweredNilCheck
	OpARMEqual
	OpARMNotEqual
	OpARMLessThan
	OpARMLessEqual
	OpARMGreaterThan
	OpARMGreaterEqual
	OpARMLessThanU
	OpARMLessEqualU
	OpARMGreaterThanU
	OpARMGreaterEqualU
	OpARMDUFFZERO
	OpARMDUFFCOPY
	OpARMLoweredZero
	OpARMLoweredMove
	OpARMLoweredGetClosurePtr
	OpARMLoweredGetCallerSP
	OpARMLoweredGetCallerPC
	OpARMFlagEQ
	OpARMFlagLT_ULT
	OpARMFlagLT_UGT
	OpARMFlagGT_UGT
	OpARMFlagGT_ULT
	OpARMInvertFlags
	OpARMLoweredWB

	OpARM64ADD
	OpARM64ADDconst
	OpARM64SUB
	OpARM64SUBconst
	OpARM64MUL
	OpARM64MULW
	OpARM64MNEG
	OpARM64MNEGW
	OpARM64MULH
	OpARM64UMULH
	OpARM64MULL
	OpARM64UMULL
	OpARM64DIV
	OpARM64UDIV
	OpARM64DIVW
	OpARM64UDIVW
	OpARM64MOD
	OpARM64UMOD
	OpARM64MODW
	OpARM64UMODW
	OpARM64FADDS
	OpARM64FADDD
	OpARM64FSUBS
	OpARM64FSUBD
	OpARM64FMULS
	OpARM64FMULD
	OpARM64FNMULS
	OpARM64FNMULD
	OpARM64FDIVS
	OpARM64FDIVD
	OpARM64AND
	OpARM64ANDconst
	OpARM64OR
	OpARM64ORconst
	OpARM64XOR
	OpARM64XORconst
	OpARM64BIC
	OpARM64EON
	OpARM64ORN
	OpARM64LoweredMuluhilo
	OpARM64MVN
	OpARM64NEG
	OpARM64FNEGS
	OpARM64FNEGD
	OpARM64FSQRTD
	OpARM64REV
	OpARM64REVW
	OpARM64REV16W
	OpARM64RBIT
	OpARM64RBITW
	OpARM64CLZ
	OpARM64CLZW
	OpARM64VCNT
	OpARM64VUADDLV
	OpARM64LoweredRound32F
	OpARM64LoweredRound64F
	OpARM64FMADDS
	OpARM64FMADDD
	OpARM64FNMADDS
	OpARM64FNMADDD
	OpARM64FMSUBS
	OpARM64FMSUBD
	OpARM64FNMSUBS
	OpARM64FNMSUBD
	OpARM64SLL
	OpARM64SLLconst
	OpARM64SRL
	OpARM64SRLconst
	OpARM64SRA
	OpARM64SRAconst
	OpARM64RORconst
	OpARM64RORWconst
	OpARM64EXTRconst
	OpARM64EXTRWconst
	OpARM64CMP
	OpARM64CMPconst
	OpARM64CMPW
	OpARM64CMPWconst
	OpARM64CMN
	OpARM64CMNconst
	OpARM64CMNW
	OpARM64CMNWconst
	OpARM64TST
	OpARM64TSTconst
	OpARM64TSTW
	OpARM64TSTWconst
	OpARM64FCMPS
	OpARM64FCMPD
	OpARM64ADDshiftLL
	OpARM64ADDshiftRL
	OpARM64ADDshiftRA
	OpARM64SUBshiftLL
	OpARM64SUBshiftRL
	OpARM64SUBshiftRA
	OpARM64ANDshiftLL
	OpARM64ANDshiftRL
	OpARM64ANDshiftRA
	OpARM64ORshiftLL
	OpARM64ORshiftRL
	OpARM64ORshiftRA
	OpARM64XORshiftLL
	OpARM64XORshiftRL
	OpARM64XORshiftRA
	OpARM64BICshiftLL
	OpARM64BICshiftRL
	OpARM64BICshiftRA
	OpARM64EONshiftLL
	OpARM64EONshiftRL
	OpARM64EONshiftRA
	OpARM64ORNshiftLL
	OpARM64ORNshiftRL
	OpARM64ORNshiftRA
	OpARM64CMPshiftLL
	OpARM64CMPshiftRL
	OpARM64CMPshiftRA
	OpARM64BFI
	OpARM64BFXIL
	OpARM64SBFIZ
	OpARM64SBFX
	OpARM64UBFIZ
	OpARM64UBFX
	OpARM64MOVDconst
	OpARM64FMOVSconst
	OpARM64FMOVDconst
	OpARM64MOVDaddr
	OpARM64MOVBload
	OpARM64MOVBUload
	OpARM64MOVHload
	OpARM64MOVHUload
	OpARM64MOVWload
	OpARM64MOVWUload
	OpARM64MOVDload
	OpARM64FMOVSload
	OpARM64FMOVDload
	OpARM64MOVDloadidx
	OpARM64MOVWloadidx
	OpARM64MOVWUloadidx
	OpARM64MOVHloadidx
	OpARM64MOVHUloadidx
	OpARM64MOVBloadidx
	OpARM64MOVBUloadidx
	OpARM64MOVHloadidx2
	OpARM64MOVHUloadidx2
	OpARM64MOVWloadidx4
	OpARM64MOVWUloadidx4
	OpARM64MOVDloadidx8
	OpARM64MOVBstore
	OpARM64MOVHstore
	OpARM64MOVWstore
	OpARM64MOVDstore
	OpARM64STP
	OpARM64FMOVSstore
	OpARM64FMOVDstore
	OpARM64MOVBstoreidx
	OpARM64MOVHstoreidx
	OpARM64MOVWstoreidx
	OpARM64MOVDstoreidx
	OpARM64MOVHstoreidx2
	OpARM64MOVWstoreidx4
	OpARM64MOVDstoreidx8
	OpARM64MOVBstorezero
	OpARM64MOVHstorezero
	OpARM64MOVWstorezero
	OpARM64MOVDstorezero
	OpARM64MOVQstorezero
	OpARM64MOVBstorezeroidx
	OpARM64MOVHstorezeroidx
	OpARM64MOVWstorezeroidx
	OpARM64MOVDstorezeroidx
	OpARM64MOVHstorezeroidx2
	OpARM64MOVWstorezeroidx4
	OpARM64MOVDstorezeroidx8
	OpARM64FMOVDgpfp
	OpARM64FMOVDfpgp
	OpARM64MOVBreg
	OpARM64MOVBUreg
	OpARM64MOVHreg
	OpARM64MOVHUreg
	OpARM64MOVWreg
	OpARM64MOVWUreg
	OpARM64MOVDreg
	OpARM64MOVDnop
	OpARM64SCVTFWS
	OpARM64SCVTFWD
	OpARM64UCVTFWS
	OpARM64UCVTFWD
	OpARM64SCVTFS
	OpARM64SCVTFD
	OpARM64UCVTFS
	OpARM64UCVTFD
	OpARM64FCVTZSSW
	OpARM64FCVTZSDW
	OpARM64FCVTZUSW
	OpARM64FCVTZUDW
	OpARM64FCVTZSS
	OpARM64FCVTZSD
	OpARM64FCVTZUS
	OpARM64FCVTZUD
	OpARM64FCVTSD
	OpARM64FCVTDS
	OpARM64FRINTAD
	OpARM64FRINTMD
	OpARM64FRINTPD
	OpARM64FRINTZD
	OpARM64CSEL
	OpARM64CSEL0
	OpARM64CALLstatic
	OpARM64CALLclosure
	OpARM64CALLinter
	OpARM64LoweredNilCheck
	OpARM64Equal
	OpARM64NotEqual
	OpARM64LessThan
	OpARM64LessEqual
	OpARM64GreaterThan
	OpARM64GreaterEqual
	OpARM64LessThanU
	OpARM64LessEqualU
	OpARM64GreaterThanU
	OpARM64GreaterEqualU
	OpARM64DUFFZERO
	OpARM64LoweredZero
	OpARM64DUFFCOPY
	OpARM64LoweredMove
	OpARM64LoweredGetClosurePtr
	OpARM64LoweredGetCallerSP
	OpARM64LoweredGetCallerPC
	OpARM64FlagEQ
	OpARM64FlagLT_ULT
	OpARM64FlagLT_UGT
	OpARM64FlagGT_UGT
	OpARM64FlagGT_ULT
	OpARM64InvertFlags
	OpARM64LDAR
	OpARM64LDARW
	OpARM64STLR
	OpARM64STLRW
	OpARM64LoweredAtomicExchange64
	OpARM64LoweredAtomicExchange32
	OpARM64LoweredAtomicAdd64
	OpARM64LoweredAtomicAdd32
	OpARM64LoweredAtomicAdd64Variant
	OpARM64LoweredAtomicAdd32Variant
	OpARM64LoweredAtomicCas64
	OpARM64LoweredAtomicCas32
	OpARM64LoweredAtomicAnd8
	OpARM64LoweredAtomicOr8
	OpARM64LoweredWB

	OpMIPSADD
	OpMIPSADDconst
	OpMIPSSUB
	OpMIPSSUBconst
	OpMIPSMUL
	OpMIPSMULT
	OpMIPSMULTU
	OpMIPSDIV
	OpMIPSDIVU
	OpMIPSADDF
	OpMIPSADDD
	OpMIPSSUBF
	OpMIPSSUBD
	OpMIPSMULF
	OpMIPSMULD
	OpMIPSDIVF
	OpMIPSDIVD
	OpMIPSAND
	OpMIPSANDconst
	OpMIPSOR
	OpMIPSORconst
	OpMIPSXOR
	OpMIPSXORconst
	OpMIPSNOR
	OpMIPSNORconst
	OpMIPSNEG
	OpMIPSNEGF
	OpMIPSNEGD
	OpMIPSSQRTD
	OpMIPSSLL
	OpMIPSSLLconst
	OpMIPSSRL
	OpMIPSSRLconst
	OpMIPSSRA
	OpMIPSSRAconst
	OpMIPSCLZ
	OpMIPSSGT
	OpMIPSSGTconst
	OpMIPSSGTzero
	OpMIPSSGTU
	OpMIPSSGTUconst
	OpMIPSSGTUzero
	OpMIPSCMPEQF
	OpMIPSCMPEQD
	OpMIPSCMPGEF
	OpMIPSCMPGED
	OpMIPSCMPGTF
	OpMIPSCMPGTD
	OpMIPSMOVWconst
	OpMIPSMOVFconst
	OpMIPSMOVDconst
	OpMIPSMOVWaddr
	OpMIPSMOVBload
	OpMIPSMOVBUload
	OpMIPSMOVHload
	OpMIPSMOVHUload
	OpMIPSMOVWload
	OpMIPSMOVFload
	OpMIPSMOVDload
	OpMIPSMOVBstore
	OpMIPSMOVHstore
	OpMIPSMOVWstore
	OpMIPSMOVFstore
	OpMIPSMOVDstore
	OpMIPSMOVBstorezero
	OpMIPSMOVHstorezero
	OpMIPSMOVWstorezero
	OpMIPSMOVBreg
	OpMIPSMOVBUreg
	OpMIPSMOVHreg
	OpMIPSMOVHUreg
	OpMIPSMOVWreg
	OpMIPSMOVWnop
	OpMIPSCMOVZ
	OpMIPSCMOVZzero
	OpMIPSMOVWF
	OpMIPSMOVWD
	OpMIPSTRUNCFW
	OpMIPSTRUNCDW
	OpMIPSMOVFD
	OpMIPSMOVDF
	OpMIPSCALLstatic
	OpMIPSCALLclosure
	OpMIPSCALLinter
	OpMIPSLoweredAtomicLoad
	OpMIPSLoweredAtomicStore
	OpMIPSLoweredAtomicStorezero
	OpMIPSLoweredAtomicExchange
	OpMIPSLoweredAtomicAdd
	OpMIPSLoweredAtomicAddconst
	OpMIPSLoweredAtomicCas
	OpMIPSLoweredAtomicAnd
	OpMIPSLoweredAtomicOr
	OpMIPSLoweredZero
	OpMIPSLoweredMove
	OpMIPSLoweredNilCheck
	OpMIPSFPFlagTrue
	OpMIPSFPFlagFalse
	OpMIPSLoweredGetClosurePtr
	OpMIPSLoweredGetCallerSP
	OpMIPSLoweredGetCallerPC
	OpMIPSLoweredWB

	OpMIPS64ADDV
	OpMIPS64ADDVconst
	OpMIPS64SUBV
	OpMIPS64SUBVconst
	OpMIPS64MULV
	OpMIPS64MULVU
	OpMIPS64DIVV
	OpMIPS64DIVVU
	OpMIPS64ADDF
	OpMIPS64ADDD
	OpMIPS64SUBF
	OpMIPS64SUBD
	OpMIPS64MULF
	OpMIPS64MULD
	OpMIPS64DIVF
	OpMIPS64DIVD
	OpMIPS64AND
	OpMIPS64ANDconst
	OpMIPS64OR
	OpMIPS64ORconst
	OpMIPS64XOR
	OpMIPS64XORconst
	OpMIPS64NOR
	OpMIPS64NORconst
	OpMIPS64NEGV
	OpMIPS64NEGF
	OpMIPS64NEGD
	OpMIPS64SQRTD
	OpMIPS64SLLV
	OpMIPS64SLLVconst
	OpMIPS64SRLV
	OpMIPS64SRLVconst
	OpMIPS64SRAV
	OpMIPS64SRAVconst
	OpMIPS64SGT
	OpMIPS64SGTconst
	OpMIPS64SGTU
	OpMIPS64SGTUconst
	OpMIPS64CMPEQF
	OpMIPS64CMPEQD
	OpMIPS64CMPGEF
	OpMIPS64CMPGED
	OpMIPS64CMPGTF
	OpMIPS64CMPGTD
	OpMIPS64MOVVconst
	OpMIPS64MOVFconst
	OpMIPS64MOVDconst
	OpMIPS64MOVVaddr
	OpMIPS64MOVBload
	OpMIPS64MOVBUload
	OpMIPS64MOVHload
	OpMIPS64MOVHUload
	OpMIPS64MOVWload
	OpMIPS64MOVWUload
	OpMIPS64MOVVload
	OpMIPS64MOVFload
	OpMIPS64MOVDload
	OpMIPS64MOVBstore
	OpMIPS64MOVHstore
	OpMIPS64MOVWstore
	OpMIPS64MOVVstore
	OpMIPS64MOVFstore
	OpMIPS64MOVDstore
	OpMIPS64MOVBstorezero
	OpMIPS64MOVHstorezero
	OpMIPS64MOVWstorezero
	OpMIPS64MOVVstorezero
	OpMIPS64MOVBreg
	OpMIPS64MOVBUreg
	OpMIPS64MOVHreg
	OpMIPS64MOVHUreg
	OpMIPS64MOVWreg
	OpMIPS64MOVWUreg
	OpMIPS64MOVVreg
	OpMIPS64MOVVnop
	OpMIPS64MOVWF
	OpMIPS64MOVWD
	OpMIPS64MOVVF
	OpMIPS64MOVVD
	OpMIPS64TRUNCFW
	OpMIPS64TRUNCDW
	OpMIPS64TRUNCFV
	OpMIPS64TRUNCDV
	OpMIPS64MOVFD
	OpMIPS64MOVDF
	OpMIPS64CALLstatic
	OpMIPS64CALLclosure
	OpMIPS64CALLinter
	OpMIPS64DUFFZERO
	OpMIPS64LoweredZero
	OpMIPS64LoweredMove
	OpMIPS64LoweredAtomicLoad32
	OpMIPS64LoweredAtomicLoad64
	OpMIPS64LoweredAtomicStore32
	OpMIPS64LoweredAtomicStore64
	OpMIPS64LoweredAtomicStorezero32
	OpMIPS64LoweredAtomicStorezero64
	OpMIPS64LoweredAtomicExchange32
	OpMIPS64LoweredAtomicExchange64
	OpMIPS64LoweredAtomicAdd32
	OpMIPS64LoweredAtomicAdd64
	OpMIPS64LoweredAtomicAddconst32
	OpMIPS64LoweredAtomicAddconst64
	OpMIPS64LoweredAtomicCas32
	OpMIPS64LoweredAtomicCas64
	OpMIPS64LoweredNilCheck
	OpMIPS64FPFlagTrue
	OpMIPS64FPFlagFalse
	OpMIPS64LoweredGetClosurePtr
	OpMIPS64LoweredGetCallerSP
	OpMIPS64LoweredGetCallerPC
	OpMIPS64LoweredWB

	OpPPC64ADD
	OpPPC64ADDconst
	OpPPC64FADD
	OpPPC64FADDS
	OpPPC64SUB
	OpPPC64FSUB
	OpPPC64FSUBS
	OpPPC64MULLD
	OpPPC64MULLW
	OpPPC64MULHD
	OpPPC64MULHW
	OpPPC64MULHDU
	OpPPC64MULHWU
	OpPPC64FMUL
	OpPPC64FMULS
	OpPPC64FMADD
	OpPPC64FMADDS
	OpPPC64FMSUB
	OpPPC64FMSUBS
	OpPPC64SRAD
	OpPPC64SRAW
	OpPPC64SRD
	OpPPC64SRW
	OpPPC64SLD
	OpPPC64SLW
	OpPPC64ROTL
	OpPPC64ROTLW
	OpPPC64ADDconstForCarry
	OpPPC64MaskIfNotCarry
	OpPPC64SRADconst
	OpPPC64SRAWconst
	OpPPC64SRDconst
	OpPPC64SRWconst
	OpPPC64SLDconst
	OpPPC64SLWconst
	OpPPC64ROTLconst
	OpPPC64ROTLWconst
	OpPPC64CNTLZD
	OpPPC64CNTLZW
	OpPPC64POPCNTD
	OpPPC64POPCNTW
	OpPPC64POPCNTB
	OpPPC64FDIV
	OpPPC64FDIVS
	OpPPC64DIVD
	OpPPC64DIVW
	OpPPC64DIVDU
	OpPPC64DIVWU
	OpPPC64FCTIDZ
	OpPPC64FCTIWZ
	OpPPC64FCFID
	OpPPC64FCFIDS
	OpPPC64FRSP
	OpPPC64MFVSRD
	OpPPC64MTVSRD
	OpPPC64AND
	OpPPC64ANDN
	OpPPC64OR
	OpPPC64ORN
	OpPPC64NOR
	OpPPC64XOR
	OpPPC64EQV
	OpPPC64NEG
	OpPPC64FNEG
	OpPPC64FSQRT
	OpPPC64FSQRTS
	OpPPC64FFLOOR
	OpPPC64FCEIL
	OpPPC64FTRUNC
	OpPPC64FROUND
	OpPPC64FABS
	OpPPC64FNABS
	OpPPC64FCPSGN
	OpPPC64ORconst
	OpPPC64XORconst
	OpPPC64ANDconst
	OpPPC64ANDCCconst
	OpPPC64MOVBreg
	OpPPC64MOVBZreg
	OpPPC64MOVHreg
	OpPPC64MOVHZreg
	OpPPC64MOVWreg
	OpPPC64MOVWZreg
	OpPPC64MOVBZload
	OpPPC64MOVHload
	OpPPC64MOVHZload
	OpPPC64MOVWload
	OpPPC64MOVWZload
	OpPPC64MOVDload
	OpPPC64MOVDBRload
	OpPPC64MOVWBRload
	OpPPC64MOVHBRload
	OpPPC64MOVDBRstore
	OpPPC64MOVWBRstore
	OpPPC64MOVHBRstore
	OpPPC64FMOVDload
	OpPPC64FMOVSload
	OpPPC64MOVBstore
	OpPPC64MOVHstore
	OpPPC64MOVWstore
	OpPPC64MOVDstore
	OpPPC64FMOVDstore
	OpPPC64FMOVSstore
	OpPPC64MOVBstorezero
	OpPPC64MOVHstorezero
	OpPPC64MOVWstorezero
	OpPPC64MOVDstorezero
	OpPPC64MOVDaddr
	OpPPC64MOVDconst
	OpPPC64FMOVDconst
	OpPPC64FMOVSconst
	OpPPC64FCMPU
	OpPPC64CMP
	OpPPC64CMPU
	OpPPC64CMPW
	OpPPC64CMPWU
	OpPPC64CMPconst
	OpPPC64CMPUconst
	OpPPC64CMPWconst
	OpPPC64CMPWUconst
	OpPPC64Equal
	OpPPC64NotEqual
	OpPPC64LessThan
	OpPPC64FLessThan
	OpPPC64LessEqual
	OpPPC64FLessEqual
	OpPPC64GreaterThan
	OpPPC64FGreaterThan
	OpPPC64GreaterEqual
	OpPPC64FGreaterEqual
	OpPPC64LoweredGetClosurePtr
	OpPPC64LoweredGetCallerSP
	OpPPC64LoweredGetCallerPC
	OpPPC64LoweredNilCheck
	OpPPC64LoweredRound32F
	OpPPC64LoweredRound64F
	OpPPC64CALLstatic
	OpPPC64CALLclosure
	OpPPC64CALLinter
	OpPPC64LoweredZero
	OpPPC64LoweredMove
	OpPPC64LoweredAtomicStore32
	OpPPC64LoweredAtomicStore64
	OpPPC64LoweredAtomicLoad32
	OpPPC64LoweredAtomicLoad64
	OpPPC64LoweredAtomicLoadPtr
	OpPPC64LoweredAtomicAdd32
	OpPPC64LoweredAtomicAdd64
	OpPPC64LoweredAtomicExchange32
	OpPPC64LoweredAtomicExchange64
	OpPPC64LoweredAtomicCas64
	OpPPC64LoweredAtomicCas32
	OpPPC64LoweredAtomicAnd8
	OpPPC64LoweredAtomicOr8
	OpPPC64LoweredWB
	OpPPC64InvertFlags
	OpPPC64FlagEQ
	OpPPC64FlagLT
	OpPPC64FlagGT

	OpS390XFADDS
	OpS390XFADD
	OpS390XFSUBS
	OpS390XFSUB
	OpS390XFMULS
	OpS390XFMUL
	OpS390XFDIVS
	OpS390XFDIV
	OpS390XFNEGS
	OpS390XFNEG
	OpS390XFMADDS
	OpS390XFMADD
	OpS390XFMSUBS
	OpS390XFMSUB
	OpS390XLPDFR
	OpS390XLNDFR
	OpS390XCPSDR
	OpS390XFIDBR
	OpS390XFMOVSload
	OpS390XFMOVDload
	OpS390XFMOVSconst
	OpS390XFMOVDconst
	OpS390XFMOVSloadidx
	OpS390XFMOVDloadidx
	OpS390XFMOVSstore
	OpS390XFMOVDstore
	OpS390XFMOVSstoreidx
	OpS390XFMOVDstoreidx
	OpS390XADD
	OpS390XADDW
	OpS390XADDconst
	OpS390XADDWconst
	OpS390XADDload
	OpS390XADDWload
	OpS390XSUB
	OpS390XSUBW
	OpS390XSUBconst
	OpS390XSUBWconst
	OpS390XSUBload
	OpS390XSUBWload
	OpS390XMULLD
	OpS390XMULLW
	OpS390XMULLDconst
	OpS390XMULLWconst
	OpS390XMULLDload
	OpS390XMULLWload
	OpS390XMULHD
	OpS390XMULHDU
	OpS390XDIVD
	OpS390XDIVW
	OpS390XDIVDU
	OpS390XDIVWU
	OpS390XMODD
	OpS390XMODW
	OpS390XMODDU
	OpS390XMODWU
	OpS390XAND
	OpS390XANDW
	OpS390XANDconst
	OpS390XANDWconst
	OpS390XANDload
	OpS390XANDWload
	OpS390XOR
	OpS390XORW
	OpS390XORconst
	OpS390XORWconst
	OpS390XORload
	OpS390XORWload
	OpS390XXOR
	OpS390XXORW
	OpS390XXORconst
	OpS390XXORWconst
	OpS390XXORload
	OpS390XXORWload
	OpS390XCMP
	OpS390XCMPW
	OpS390XCMPU
	OpS390XCMPWU
	OpS390XCMPconst
	OpS390XCMPWconst
	OpS390XCMPUconst
	OpS390XCMPWUconst
	OpS390XFCMPS
	OpS390XFCMP
	OpS390XSLD
	OpS390XSLW
	OpS390XSLDconst
	OpS390XSLWconst
	OpS390XSRD
	OpS390XSRW
	OpS390XSRDconst
	OpS390XSRWconst
	OpS390XSRAD
	OpS390XSRAW
	OpS390XSRADconst
	OpS390XSRAWconst
	OpS390XRLLGconst
	OpS390XRLLconst
	OpS390XNEG
	OpS390XNEGW
	OpS390XNOT
	OpS390XNOTW
	OpS390XFSQRT
	OpS390XMOVDEQ
	OpS390XMOVDNE
	OpS390XMOVDLT
	OpS390XMOVDLE
	OpS390XMOVDGT
	OpS390XMOVDGE
	OpS390XMOVDGTnoinv
	OpS390XMOVDGEnoinv
	OpS390XMOVBreg
	OpS390XMOVBZreg
	OpS390XMOVHreg
	OpS390XMOVHZreg
	OpS390XMOVWreg
	OpS390XMOVWZreg
	OpS390XMOVDreg
	OpS390XMOVDnop
	OpS390XMOVDconst
	OpS390XLDGR
	OpS390XLGDR
	OpS390XCFDBRA
	OpS390XCGDBRA
	OpS390XCFEBRA
	OpS390XCGEBRA
	OpS390XCEFBRA
	OpS390XCDFBRA
	OpS390XCEGBRA
	OpS390XCDGBRA
	OpS390XLEDBR
	OpS390XLDEBR
	OpS390XMOVDaddr
	OpS390XMOVDaddridx
	OpS390XMOVBZload
	OpS390XMOVBload
	OpS390XMOVHZload
	OpS390XMOVHload
	OpS390XMOVWZload
	OpS390XMOVWload
	OpS390XMOVDload
	OpS390XMOVWBR
	OpS390XMOVDBR
	OpS390XMOVHBRload
	OpS390XMOVWBRload
	OpS390XMOVDBRload
	OpS390XMOVBstore
	OpS390XMOVHstore
	OpS390XMOVWstore
	OpS390XMOVDstore
	OpS390XMOVHBRstore
	OpS390XMOVWBRstore
	OpS390XMOVDBRstore
	OpS390XMVC
	OpS390XMOVBZloadidx
	OpS390XMOVBloadidx
	OpS390XMOVHZloadidx
	OpS390XMOVHloadidx
	OpS390XMOVWZloadidx
	OpS390XMOVWloadidx
	OpS390XMOVDloadidx
	OpS390XMOVHBRloadidx
	OpS390XMOVWBRloadidx
	OpS390XMOVDBRloadidx
	OpS390XMOVBstoreidx
	OpS390XMOVHstoreidx
	OpS390XMOVWstoreidx
	OpS390XMOVDstoreidx
	OpS390XMOVHBRstoreidx
	OpS390XMOVWBRstoreidx
	OpS390XMOVDBRstoreidx
	OpS390XMOVBstoreconst
	OpS390XMOVHstoreconst
	OpS390XMOVWstoreconst
	OpS390XMOVDstoreconst
	OpS390XCLEAR
	OpS390XCALLstatic
	OpS390XCALLclosure
	OpS390XCALLinter
	OpS390XInvertFlags
	OpS390XLoweredGetG
	OpS390XLoweredGetClosurePtr
	OpS390XLoweredGetCallerSP
	OpS390XLoweredGetCallerPC
	OpS390XLoweredNilCheck
	OpS390XLoweredRound32F
	OpS390XLoweredRound64F
	OpS390XLoweredWB
	OpS390XFlagEQ
	OpS390XFlagLT
	OpS390XFlagGT
	OpS390XMOVWZatomicload
	OpS390XMOVDatomicload
	OpS390XMOVWatomicstore
	OpS390XMOVDatomicstore
	OpS390XLAA
	OpS390XLAAG
	OpS390XAddTupleFirst32
	OpS390XAddTupleFirst64
	OpS390XLoweredAtomicCas32
	OpS390XLoweredAtomicCas64
	OpS390XLoweredAtomicExchange32
	OpS390XLoweredAtomicExchange64
	OpS390XFLOGR
	OpS390XSTMG2
	OpS390XSTMG3
	OpS390XSTMG4
	OpS390XSTM2
	OpS390XSTM3
	OpS390XSTM4
	OpS390XLoweredMove
	OpS390XLoweredZero

	OpWasmLoweredStaticCall
	OpWasmLoweredClosureCall
	OpWasmLoweredInterCall
	OpWasmLoweredAddr
	OpWasmLoweredMove
	OpWasmLoweredZero
	OpWasmLoweredGetClosurePtr
	OpWasmLoweredGetCallerPC
	OpWasmLoweredGetCallerSP
	OpWasmLoweredNilCheck
	OpWasmLoweredWB
	OpWasmLoweredRound32F
	OpWasmLoweredConvert
	OpWasmSelect
	OpWasmI64Load8U
	OpWasmI64Load8S
	OpWasmI64Load16U
	OpWasmI64Load16S
	OpWasmI64Load32U
	OpWasmI64Load32S
	OpWasmI64Load
	OpWasmI64Store8
	OpWasmI64Store16
	OpWasmI64Store32
	OpWasmI64Store
	OpWasmF32Load
	OpWasmF64Load
	OpWasmF32Store
	OpWasmF64Store
	OpWasmI64Const
	OpWasmF64Const
	OpWasmI64Eqz
	OpWasmI64Eq
	OpWasmI64Ne
	OpWasmI64LtS
	OpWasmI64LtU
	OpWasmI64GtS
	OpWasmI64GtU
	OpWasmI64LeS
	OpWasmI64LeU
	OpWasmI64GeS
	OpWasmI64GeU
	OpWasmF64Eq
	OpWasmF64Ne
	OpWasmF64Lt
	OpWasmF64Gt
	OpWasmF64Le
	OpWasmF64Ge
	OpWasmI64Add
	OpWasmI64AddConst
	OpWasmI64Sub
	OpWasmI64Mul
	OpWasmI64DivS
	OpWasmI64DivU
	OpWasmI64RemS
	OpWasmI64RemU
	OpWasmI64And
	OpWasmI64Or
	OpWasmI64Xor
	OpWasmI64Shl
	OpWasmI64ShrS
	OpWasmI64ShrU
	OpWasmF64Neg
	OpWasmF64Add
	OpWasmF64Sub
	OpWasmF64Mul
	OpWasmF64Div
	OpWasmI64TruncSF64
	OpWasmI64TruncUF64
	OpWasmF64ConvertSI64
	OpWasmF64ConvertUI64

	OpAdd8
	OpAdd16
	OpAdd32
	OpAdd64
	OpAddPtr
	OpAdd32F
	OpAdd64F
	OpSub8
	OpSub16
	OpSub32
	OpSub64
	OpSubPtr
	OpSub32F
	OpSub64F
	OpMul8
	OpMul16
	OpMul32
	OpMul64
	OpMul32F
	OpMul64F
	OpDiv32F
	OpDiv64F
	OpHmul32
	OpHmul32u
	OpHmul64
	OpHmul64u
	OpMul32uhilo
	OpMul64uhilo
	OpAvg32u
	OpAvg64u
	OpDiv8
	OpDiv8u
	OpDiv16
	OpDiv16u
	OpDiv32
	OpDiv32u
	OpDiv64
	OpDiv64u
	OpDiv128u
	OpMod8
	OpMod8u
	OpMod16
	OpMod16u
	OpMod32
	OpMod32u
	OpMod64
	OpMod64u
	OpAnd8
	OpAnd16
	OpAnd32
	OpAnd64
	OpOr8
	OpOr16
	OpOr32
	OpOr64
	OpXor8
	OpXor16
	OpXor32
	OpXor64
	OpLsh8x8
	OpLsh8x16
	OpLsh8x32
	OpLsh8x64
	OpLsh16x8
	OpLsh16x16
	OpLsh16x32
	OpLsh16x64
	OpLsh32x8
	OpLsh32x16
	OpLsh32x32
	OpLsh32x64
	OpLsh64x8
	OpLsh64x16
	OpLsh64x32
	OpLsh64x64
	OpRsh8x8
	OpRsh8x16
	OpRsh8x32
	OpRsh8x64
	OpRsh16x8
	OpRsh16x16
	OpRsh16x32
	OpRsh16x64
	OpRsh32x8
	OpRsh32x16
	OpRsh32x32
	OpRsh32x64
	OpRsh64x8
	OpRsh64x16
	OpRsh64x32
	OpRsh64x64
	OpRsh8Ux8
	OpRsh8Ux16
	OpRsh8Ux32
	OpRsh8Ux64
	OpRsh16Ux8
	OpRsh16Ux16
	OpRsh16Ux32
	OpRsh16Ux64
	OpRsh32Ux8
	OpRsh32Ux16
	OpRsh32Ux32
	OpRsh32Ux64
	OpRsh64Ux8
	OpRsh64Ux16
	OpRsh64Ux32
	OpRsh64Ux64
	OpEq8
	OpEq16
	OpEq32
	OpEq64
	OpEqPtr
	OpEqInter
	OpEqSlice
	OpEq32F
	OpEq64F
	OpNeq8
	OpNeq16
	OpNeq32
	OpNeq64
	OpNeqPtr
	OpNeqInter
	OpNeqSlice
	OpNeq32F
	OpNeq64F
	OpLess8
	OpLess8U
	OpLess16
	OpLess16U
	OpLess32
	OpLess32U
	OpLess64
	OpLess64U
	OpLess32F
	OpLess64F
	OpLeq8
	OpLeq8U
	OpLeq16
	OpLeq16U
	OpLeq32
	OpLeq32U
	OpLeq64
	OpLeq64U
	OpLeq32F
	OpLeq64F
	OpGreater8
	OpGreater8U
	OpGreater16
	OpGreater16U
	OpGreater32
	OpGreater32U
	OpGreater64
	OpGreater64U
	OpGreater32F
	OpGreater64F
	OpGeq8
	OpGeq8U
	OpGeq16
	OpGeq16U
	OpGeq32
	OpGeq32U
	OpGeq64
	OpGeq64U
	OpGeq32F
	OpGeq64F
	OpCondSelect
	OpAndB
	OpOrB
	OpEqB
	OpNeqB
	OpNot
	OpNeg8
	OpNeg16
	OpNeg32
	OpNeg64
	OpNeg32F
	OpNeg64F
	OpCom8
	OpCom16
	OpCom32
	OpCom64
	OpCtz8
	OpCtz16
	OpCtz32
	OpCtz64
	OpCtz8NonZero
	OpCtz16NonZero
	OpCtz32NonZero
	OpCtz64NonZero
	OpBitLen8
	OpBitLen16
	OpBitLen32
	OpBitLen64
	OpBswap32
	OpBswap64
	OpBitRev8
	OpBitRev16
	OpBitRev32
	OpBitRev64
	OpPopCount8
	OpPopCount16
	OpPopCount32
	OpPopCount64
	OpSqrt
	OpFloor
	OpCeil
	OpTrunc
	OpRound
	OpRoundToEven
	OpAbs
	OpCopysign
	OpPhi
	OpCopy
	OpConvert
	OpConstBool
	OpConstString
	OpConstNil
	OpConst8
	OpConst16
	OpConst32
	OpConst64
	OpConst32F
	OpConst64F
	OpConstInterface
	OpConstSlice
	OpInitMem
	OpArg
	OpAddr
	OpSP
	OpSB
	OpLoad
	OpStore
	OpMove
	OpZero
	OpStoreWB
	OpMoveWB
	OpZeroWB
	OpWB
	OpClosureCall
	OpStaticCall
	OpInterCall
	OpSignExt8to16
	OpSignExt8to32
	OpSignExt8to64
	OpSignExt16to32
	OpSignExt16to64
	OpSignExt32to64
	OpZeroExt8to16
	OpZeroExt8to32
	OpZeroExt8to64
	OpZeroExt16to32
	OpZeroExt16to64
	OpZeroExt32to64
	OpTrunc16to8
	OpTrunc32to8
	OpTrunc32to16
	OpTrunc64to8
	OpTrunc64to16
	OpTrunc64to32
	OpCvt32to32F
	OpCvt32to64F
	OpCvt64to32F
	OpCvt64to64F
	OpCvt32Fto32
	OpCvt32Fto64
	OpCvt64Fto32
	OpCvt64Fto64
	OpCvt32Fto64F
	OpCvt64Fto32F
	OpRound32F
	OpRound64F
	OpIsNonNil
	OpIsInBounds
	OpIsSliceInBounds
	OpNilCheck
	OpGetG
	OpGetClosurePtr
	OpGetCallerPC
	OpGetCallerSP
	OpPtrIndex
	OpOffPtr
	OpSliceMake
	OpSlicePtr
	OpSliceLen
	OpSliceCap
	OpComplexMake
	OpComplexReal
	OpComplexImag
	OpStringMake
	OpStringPtr
	OpStringLen
	OpIMake
	OpITab
	OpIData
	OpStructMake0
	OpStructMake1
	OpStructMake2
	OpStructMake3
	OpStructMake4
	OpStructSelect
	OpArrayMake0
	OpArrayMake1
	OpArraySelect
	OpStoreReg
	OpLoadReg
	OpFwdRef
	OpUnknown
	OpVarDef
	OpVarKill
	OpVarLive
	OpKeepAlive
	OpInt64Make
	OpInt64Hi
	OpInt64Lo
	OpAdd32carry
	OpAdd32withcarry
	OpSub32carry
	OpSub32withcarry
	OpSignmask
	OpZeromask
	OpSlicemask
	OpCvt32Uto32F
	OpCvt32Uto64F
	OpCvt32Fto32U
	OpCvt64Fto32U
	OpCvt64Uto32F
	OpCvt64Uto64F
	OpCvt32Fto64U
	OpCvt64Fto64U
	OpSelect0
	OpSelect1
	OpAtomicLoad32
	OpAtomicLoad64
	OpAtomicLoadPtr
	OpAtomicStore32
	OpAtomicStore64
	OpAtomicStorePtrNoWB
	OpAtomicExchange32
	OpAtomicExchange64
	OpAtomicAdd32
	OpAtomicAdd64
	OpAtomicCompareAndSwap32
	OpAtomicCompareAndSwap64
	OpAtomicAnd8
	OpAtomicOr8
	OpAtomicAdd32Variant
	OpAtomicAdd64Variant
	OpClobber
)

func (o Op) Asm(pstate *PackageState) obj.As          { return pstate.opcodeTable[o].asm }
func (o Op) String(pstate *PackageState) string       { return pstate.opcodeTable[o].name }
func (o Op) UsesScratch(pstate *PackageState) bool    { return pstate.opcodeTable[o].usesScratch }
func (o Op) SymEffect(pstate *PackageState) SymEffect { return pstate.opcodeTable[o].symEffect }
func (o Op) IsCall(pstate *PackageState) bool         { return pstate.opcodeTable[o].call }
