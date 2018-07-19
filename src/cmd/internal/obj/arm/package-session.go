package arm

import (
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
)

type PackageSession struct {
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	sys    *sys.PackageSession

	ARMDWARFRegisters map[int16]int16
	Anames            []string

	Linkarm obj.LinkArch
	cnames5 []string

	deferreturn *obj.LSym
	mbOp        []struct {
		reg int16
		enc uint32
	}

	oprange              [150][]Optab
	optab                []Optab
	progedit_tlsfallback *obj.LSym

	symdiv  *obj.LSym
	symdivu *obj.LSym
	symmod  *obj.LSym
	symmodu *obj.LSym

	unaryDst map[obj.As]bool
	xcmp     [45][45]bool
}

func NewPackageSession(obj_psess *obj.PackageSession, objabi_psess *objabi.PackageSession, sys_psess *sys.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.obj = obj_psess
	psess.objabi = objabi_psess
	psess.sys = sys_psess
	psess.cnames5 = []string{
		"NONE",
		"REG",
		"REGREG",
		"REGREG2",
		"REGLIST",
		"SHIFT",
		"SHIFTADDR",
		"FREG",
		"PSR",
		"FCR",
		"SPR",
		"RCON",
		"NCON",
		"RCON2A",
		"RCON2S",
		"SCON",
		"LCON",
		"LCONADDR",
		"ZFCON",
		"SFCON",
		"LFCON",
		"RACON",
		"LACON",
		"SBRA",
		"LBRA",
		"HAUTO",
		"FAUTO",
		"HFAUTO",
		"SAUTO",
		"LAUTO",
		"HOREG",
		"FOREG",
		"HFOREG",
		"SOREG",
		"ROREG",
		"SROREG",
		"LOREG",
		"PC",
		"SP",
		"HREG",
		"ADDR",
		"C_TLS_LE",
		"C_TLS_IE",
		"TEXTSIZE",
		"GOK",
		"NCLASS",
		"SCOND = (1<<4)-1",
		"SBIT = 1<<4",
		"PBIT = 1<<5",
		"WBIT = 1<<6",
		"FBIT = 1<<7",
		"UBIT = 1<<7",
		"SCOND_XOR = 14",
		"SCOND_EQ = 0 ^ C_SCOND_XOR",
		"SCOND_NE = 1 ^ C_SCOND_XOR",
		"SCOND_HS = 2 ^ C_SCOND_XOR",
		"SCOND_LO = 3 ^ C_SCOND_XOR",
		"SCOND_MI = 4 ^ C_SCOND_XOR",
		"SCOND_PL = 5 ^ C_SCOND_XOR",
		"SCOND_VS = 6 ^ C_SCOND_XOR",
		"SCOND_VC = 7 ^ C_SCOND_XOR",
		"SCOND_HI = 8 ^ C_SCOND_XOR",
		"SCOND_LS = 9 ^ C_SCOND_XOR",
		"SCOND_GE = 10 ^ C_SCOND_XOR",
		"SCOND_LT = 11 ^ C_SCOND_XOR",
		"SCOND_GT = 12 ^ C_SCOND_XOR",
		"SCOND_LE = 13 ^ C_SCOND_XOR",
		"SCOND_NONE = 14 ^ C_SCOND_XOR",
		"SCOND_NV = 15 ^ C_SCOND_XOR",
	}
	psess.ARMDWARFRegisters = map[int16]int16{}
	psess.mbOp = []struct {
		reg int16
		enc uint32
	}{
		{REG_MB_SY, 15},
		{REG_MB_ST, 14},
		{REG_MB_ISH, 11},
		{REG_MB_ISHST, 10},
		{REG_MB_NSH, 7},
		{REG_MB_NSHST, 6},
		{REG_MB_OSH, 3},
		{REG_MB_OSHST, 2},
	}
	psess.unaryDst = map[obj.As]bool{
		ASWI:  true,
		AWORD: true,
	}
	psess.optab = []Optab{

		{obj.ATEXT, C_ADDR, C_NONE, C_TEXTSIZE, 0, 0, 0, 0, 0, 0},
		{AADD, C_REG, C_REG, C_REG, 1, 4, 0, 0, 0, C_SBIT},
		{AADD, C_REG, C_NONE, C_REG, 1, 4, 0, 0, 0, C_SBIT},
		{AAND, C_REG, C_REG, C_REG, 1, 4, 0, 0, 0, C_SBIT},
		{AAND, C_REG, C_NONE, C_REG, 1, 4, 0, 0, 0, C_SBIT},
		{AORR, C_REG, C_REG, C_REG, 1, 4, 0, 0, 0, C_SBIT},
		{AORR, C_REG, C_NONE, C_REG, 1, 4, 0, 0, 0, C_SBIT},
		{AMOVW, C_REG, C_NONE, C_REG, 1, 4, 0, 0, 0, C_SBIT},
		{AMVN, C_REG, C_NONE, C_REG, 1, 4, 0, 0, 0, C_SBIT},
		{ACMP, C_REG, C_REG, C_NONE, 1, 4, 0, 0, 0, 0},
		{AADD, C_RCON, C_REG, C_REG, 2, 4, 0, 0, 0, C_SBIT},
		{AADD, C_RCON, C_NONE, C_REG, 2, 4, 0, 0, 0, C_SBIT},
		{AAND, C_RCON, C_REG, C_REG, 2, 4, 0, 0, 0, C_SBIT},
		{AAND, C_RCON, C_NONE, C_REG, 2, 4, 0, 0, 0, C_SBIT},
		{AORR, C_RCON, C_REG, C_REG, 2, 4, 0, 0, 0, C_SBIT},
		{AORR, C_RCON, C_NONE, C_REG, 2, 4, 0, 0, 0, C_SBIT},
		{AMOVW, C_RCON, C_NONE, C_REG, 2, 4, 0, 0, 0, 0},
		{AMVN, C_RCON, C_NONE, C_REG, 2, 4, 0, 0, 0, 0},
		{ACMP, C_RCON, C_REG, C_NONE, 2, 4, 0, 0, 0, 0},
		{AADD, C_SHIFT, C_REG, C_REG, 3, 4, 0, 0, 0, C_SBIT},
		{AADD, C_SHIFT, C_NONE, C_REG, 3, 4, 0, 0, 0, C_SBIT},
		{AAND, C_SHIFT, C_REG, C_REG, 3, 4, 0, 0, 0, C_SBIT},
		{AAND, C_SHIFT, C_NONE, C_REG, 3, 4, 0, 0, 0, C_SBIT},
		{AORR, C_SHIFT, C_REG, C_REG, 3, 4, 0, 0, 0, C_SBIT},
		{AORR, C_SHIFT, C_NONE, C_REG, 3, 4, 0, 0, 0, C_SBIT},
		{AMVN, C_SHIFT, C_NONE, C_REG, 3, 4, 0, 0, 0, C_SBIT},
		{ACMP, C_SHIFT, C_REG, C_NONE, 3, 4, 0, 0, 0, 0},
		{AMOVW, C_RACON, C_NONE, C_REG, 4, 4, REGSP, 0, 0, C_SBIT},
		{AB, C_NONE, C_NONE, C_SBRA, 5, 4, 0, LPOOL, 0, 0},
		{ABL, C_NONE, C_NONE, C_SBRA, 5, 4, 0, 0, 0, 0},
		{ABX, C_NONE, C_NONE, C_SBRA, 74, 20, 0, 0, 0, 0},
		{ABEQ, C_NONE, C_NONE, C_SBRA, 5, 4, 0, 0, 0, 0},
		{ABEQ, C_RCON, C_NONE, C_SBRA, 5, 4, 0, 0, 0, 0},
		{AB, C_NONE, C_NONE, C_ROREG, 6, 4, 0, LPOOL, 0, 0},
		{ABL, C_NONE, C_NONE, C_ROREG, 7, 4, 0, 0, 0, 0},
		{ABL, C_REG, C_NONE, C_ROREG, 7, 4, 0, 0, 0, 0},
		{ABX, C_NONE, C_NONE, C_ROREG, 75, 12, 0, 0, 0, 0},
		{ABXRET, C_NONE, C_NONE, C_ROREG, 76, 4, 0, 0, 0, 0},
		{ASLL, C_RCON, C_REG, C_REG, 8, 4, 0, 0, 0, C_SBIT},
		{ASLL, C_RCON, C_NONE, C_REG, 8, 4, 0, 0, 0, C_SBIT},
		{ASLL, C_REG, C_NONE, C_REG, 9, 4, 0, 0, 0, C_SBIT},
		{ASLL, C_REG, C_REG, C_REG, 9, 4, 0, 0, 0, C_SBIT},
		{ASWI, C_NONE, C_NONE, C_NONE, 10, 4, 0, 0, 0, 0},
		{ASWI, C_NONE, C_NONE, C_LCON, 10, 4, 0, 0, 0, 0},
		{AWORD, C_NONE, C_NONE, C_LCON, 11, 4, 0, 0, 0, 0},
		{AWORD, C_NONE, C_NONE, C_LCONADDR, 11, 4, 0, 0, 0, 0},
		{AWORD, C_NONE, C_NONE, C_ADDR, 11, 4, 0, 0, 0, 0},
		{AWORD, C_NONE, C_NONE, C_TLS_LE, 103, 4, 0, 0, 0, 0},
		{AWORD, C_NONE, C_NONE, C_TLS_IE, 104, 4, 0, 0, 0, 0},
		{AMOVW, C_NCON, C_NONE, C_REG, 12, 4, 0, 0, 0, 0},
		{AMOVW, C_SCON, C_NONE, C_REG, 12, 4, 0, 0, 0, 0},
		{AMOVW, C_LCON, C_NONE, C_REG, 12, 4, 0, LFROM, 0, 0},
		{AMOVW, C_LCONADDR, C_NONE, C_REG, 12, 4, 0, LFROM | LPCREL, 4, 0},
		{AMVN, C_NCON, C_NONE, C_REG, 12, 4, 0, 0, 0, 0},
		{AADD, C_NCON, C_REG, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AADD, C_NCON, C_NONE, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AAND, C_NCON, C_REG, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AAND, C_NCON, C_NONE, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AORR, C_NCON, C_REG, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AORR, C_NCON, C_NONE, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{ACMP, C_NCON, C_REG, C_NONE, 13, 8, 0, 0, 0, 0},
		{AADD, C_SCON, C_REG, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AADD, C_SCON, C_NONE, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AAND, C_SCON, C_REG, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AAND, C_SCON, C_NONE, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AORR, C_SCON, C_REG, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AORR, C_SCON, C_NONE, C_REG, 13, 8, 0, 0, 0, C_SBIT},
		{AMVN, C_SCON, C_NONE, C_REG, 13, 8, 0, 0, 0, 0},
		{ACMP, C_SCON, C_REG, C_NONE, 13, 8, 0, 0, 0, 0},
		{AADD, C_RCON2A, C_REG, C_REG, 106, 8, 0, 0, 0, 0},
		{AADD, C_RCON2A, C_NONE, C_REG, 106, 8, 0, 0, 0, 0},
		{AORR, C_RCON2A, C_REG, C_REG, 106, 8, 0, 0, 0, 0},
		{AORR, C_RCON2A, C_NONE, C_REG, 106, 8, 0, 0, 0, 0},
		{AADD, C_RCON2S, C_REG, C_REG, 107, 8, 0, 0, 0, 0},
		{AADD, C_RCON2S, C_NONE, C_REG, 107, 8, 0, 0, 0, 0},
		{AADD, C_LCON, C_REG, C_REG, 13, 8, 0, LFROM, 0, C_SBIT},
		{AADD, C_LCON, C_NONE, C_REG, 13, 8, 0, LFROM, 0, C_SBIT},
		{AAND, C_LCON, C_REG, C_REG, 13, 8, 0, LFROM, 0, C_SBIT},
		{AAND, C_LCON, C_NONE, C_REG, 13, 8, 0, LFROM, 0, C_SBIT},
		{AORR, C_LCON, C_REG, C_REG, 13, 8, 0, LFROM, 0, C_SBIT},
		{AORR, C_LCON, C_NONE, C_REG, 13, 8, 0, LFROM, 0, C_SBIT},
		{AMVN, C_LCON, C_NONE, C_REG, 13, 8, 0, LFROM, 0, 0},
		{ACMP, C_LCON, C_REG, C_NONE, 13, 8, 0, LFROM, 0, 0},
		{AMOVB, C_REG, C_NONE, C_REG, 1, 4, 0, 0, 0, 0},
		{AMOVBS, C_REG, C_NONE, C_REG, 14, 8, 0, 0, 0, 0},
		{AMOVBU, C_REG, C_NONE, C_REG, 58, 4, 0, 0, 0, 0},
		{AMOVH, C_REG, C_NONE, C_REG, 1, 4, 0, 0, 0, 0},
		{AMOVHS, C_REG, C_NONE, C_REG, 14, 8, 0, 0, 0, 0},
		{AMOVHU, C_REG, C_NONE, C_REG, 14, 8, 0, 0, 0, 0},
		{AMUL, C_REG, C_REG, C_REG, 15, 4, 0, 0, 0, C_SBIT},
		{AMUL, C_REG, C_NONE, C_REG, 15, 4, 0, 0, 0, C_SBIT},
		{ADIV, C_REG, C_REG, C_REG, 16, 4, 0, 0, 0, 0},
		{ADIV, C_REG, C_NONE, C_REG, 16, 4, 0, 0, 0, 0},
		{ADIVHW, C_REG, C_REG, C_REG, 105, 4, 0, 0, 0, 0},
		{ADIVHW, C_REG, C_NONE, C_REG, 105, 4, 0, 0, 0, 0},
		{AMULL, C_REG, C_REG, C_REGREG, 17, 4, 0, 0, 0, C_SBIT},
		{ABFX, C_LCON, C_REG, C_REG, 18, 4, 0, 0, 0, 0},
		{ABFX, C_LCON, C_NONE, C_REG, 18, 4, 0, 0, 0, 0},
		{AMOVW, C_REG, C_NONE, C_SAUTO, 20, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_REG, C_NONE, C_SOREG, 20, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_REG, C_NONE, C_SAUTO, 20, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_REG, C_NONE, C_SOREG, 20, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_REG, C_NONE, C_SAUTO, 20, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_REG, C_NONE, C_SOREG, 20, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_REG, C_NONE, C_SAUTO, 20, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_REG, C_NONE, C_SOREG, 20, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_SAUTO, C_NONE, C_REG, 21, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_SOREG, C_NONE, C_REG, 21, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_SAUTO, C_NONE, C_REG, 21, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_SOREG, C_NONE, C_REG, 21, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AXTAB, C_SHIFT, C_REG, C_REG, 22, 4, 0, 0, 0, 0},
		{AXTAB, C_SHIFT, C_NONE, C_REG, 22, 4, 0, 0, 0, 0},
		{AMOVW, C_SHIFT, C_NONE, C_REG, 23, 4, 0, 0, 0, C_SBIT},
		{AMOVB, C_SHIFT, C_NONE, C_REG, 23, 4, 0, 0, 0, 0},
		{AMOVBS, C_SHIFT, C_NONE, C_REG, 23, 4, 0, 0, 0, 0},
		{AMOVBU, C_SHIFT, C_NONE, C_REG, 23, 4, 0, 0, 0, 0},
		{AMOVH, C_SHIFT, C_NONE, C_REG, 23, 4, 0, 0, 0, 0},
		{AMOVHS, C_SHIFT, C_NONE, C_REG, 23, 4, 0, 0, 0, 0},
		{AMOVHU, C_SHIFT, C_NONE, C_REG, 23, 4, 0, 0, 0, 0},
		{AMOVW, C_REG, C_NONE, C_LAUTO, 30, 8, REGSP, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_REG, C_NONE, C_LOREG, 30, 8, 0, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_REG, C_NONE, C_ADDR, 64, 8, 0, LTO | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_REG, C_NONE, C_LAUTO, 30, 8, REGSP, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_REG, C_NONE, C_LOREG, 30, 8, 0, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_REG, C_NONE, C_ADDR, 64, 8, 0, LTO | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_REG, C_NONE, C_LAUTO, 30, 8, REGSP, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_REG, C_NONE, C_LOREG, 30, 8, 0, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_REG, C_NONE, C_ADDR, 64, 8, 0, LTO | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_REG, C_NONE, C_LAUTO, 30, 8, REGSP, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_REG, C_NONE, C_LOREG, 30, 8, 0, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_REG, C_NONE, C_ADDR, 64, 8, 0, LTO | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_TLS_LE, C_NONE, C_REG, 101, 4, 0, LFROM, 0, 0},
		{AMOVW, C_TLS_IE, C_NONE, C_REG, 102, 8, 0, LFROM, 0, 0},
		{AMOVW, C_LAUTO, C_NONE, C_REG, 31, 8, REGSP, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_LOREG, C_NONE, C_REG, 31, 8, 0, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_ADDR, C_NONE, C_REG, 65, 8, 0, LFROM | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_LAUTO, C_NONE, C_REG, 31, 8, REGSP, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_LOREG, C_NONE, C_REG, 31, 8, 0, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_ADDR, C_NONE, C_REG, 65, 8, 0, LFROM | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_LACON, C_NONE, C_REG, 34, 8, REGSP, LFROM, 0, C_SBIT},
		{AMOVW, C_PSR, C_NONE, C_REG, 35, 4, 0, 0, 0, 0},
		{AMOVW, C_REG, C_NONE, C_PSR, 36, 4, 0, 0, 0, 0},
		{AMOVW, C_RCON, C_NONE, C_PSR, 37, 4, 0, 0, 0, 0},
		{AMOVM, C_REGLIST, C_NONE, C_SOREG, 38, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVM, C_SOREG, C_NONE, C_REGLIST, 39, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{ASWPW, C_SOREG, C_REG, C_REG, 40, 4, 0, 0, 0, 0},
		{ARFE, C_NONE, C_NONE, C_NONE, 41, 4, 0, 0, 0, 0},
		{AMOVF, C_FREG, C_NONE, C_FAUTO, 50, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_FREG, C_NONE, C_FOREG, 50, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_FAUTO, C_NONE, C_FREG, 51, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_FOREG, C_NONE, C_FREG, 51, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_FREG, C_NONE, C_LAUTO, 52, 12, REGSP, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_FREG, C_NONE, C_LOREG, 52, 12, 0, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_LAUTO, C_NONE, C_FREG, 53, 12, REGSP, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_LOREG, C_NONE, C_FREG, 53, 12, 0, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_FREG, C_NONE, C_ADDR, 68, 8, 0, LTO | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVF, C_ADDR, C_NONE, C_FREG, 69, 8, 0, LFROM | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AADDF, C_FREG, C_NONE, C_FREG, 54, 4, 0, 0, 0, 0},
		{AADDF, C_FREG, C_FREG, C_FREG, 54, 4, 0, 0, 0, 0},
		{AMOVF, C_FREG, C_NONE, C_FREG, 55, 4, 0, 0, 0, 0},
		{ANEGF, C_FREG, C_NONE, C_FREG, 55, 4, 0, 0, 0, 0},
		{AMOVW, C_REG, C_NONE, C_FCR, 56, 4, 0, 0, 0, 0},
		{AMOVW, C_FCR, C_NONE, C_REG, 57, 4, 0, 0, 0, 0},
		{AMOVW, C_SHIFTADDR, C_NONE, C_REG, 59, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_SHIFTADDR, C_NONE, C_REG, 59, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_SHIFTADDR, C_NONE, C_REG, 60, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_SHIFTADDR, C_NONE, C_REG, 60, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_SHIFTADDR, C_NONE, C_REG, 60, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_SHIFTADDR, C_NONE, C_REG, 60, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_SHIFTADDR, C_NONE, C_REG, 60, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVW, C_REG, C_NONE, C_SHIFTADDR, 61, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_REG, C_NONE, C_SHIFTADDR, 61, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_REG, C_NONE, C_SHIFTADDR, 61, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBU, C_REG, C_NONE, C_SHIFTADDR, 61, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_REG, C_NONE, C_SHIFTADDR, 62, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_REG, C_NONE, C_SHIFTADDR, 62, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_REG, C_NONE, C_SHIFTADDR, 62, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_REG, C_NONE, C_HAUTO, 70, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_REG, C_NONE, C_HOREG, 70, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_REG, C_NONE, C_HAUTO, 70, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_REG, C_NONE, C_HOREG, 70, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_REG, C_NONE, C_HAUTO, 70, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_REG, C_NONE, C_HOREG, 70, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_HAUTO, C_NONE, C_REG, 71, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_HOREG, C_NONE, C_REG, 71, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_HAUTO, C_NONE, C_REG, 71, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_HOREG, C_NONE, C_REG, 71, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_HAUTO, C_NONE, C_REG, 71, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_HOREG, C_NONE, C_REG, 71, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_HAUTO, C_NONE, C_REG, 71, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_HOREG, C_NONE, C_REG, 71, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_HAUTO, C_NONE, C_REG, 71, 4, REGSP, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_HOREG, C_NONE, C_REG, 71, 4, 0, 0, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_REG, C_NONE, C_LAUTO, 72, 8, REGSP, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_REG, C_NONE, C_LOREG, 72, 8, 0, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_REG, C_NONE, C_ADDR, 94, 8, 0, LTO | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_REG, C_NONE, C_LAUTO, 72, 8, REGSP, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_REG, C_NONE, C_LOREG, 72, 8, 0, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_REG, C_NONE, C_ADDR, 94, 8, 0, LTO | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_REG, C_NONE, C_LAUTO, 72, 8, REGSP, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_REG, C_NONE, C_LOREG, 72, 8, 0, LTO, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_REG, C_NONE, C_ADDR, 94, 8, 0, LTO | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_LAUTO, C_NONE, C_REG, 73, 8, REGSP, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_LOREG, C_NONE, C_REG, 73, 8, 0, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVB, C_ADDR, C_NONE, C_REG, 93, 8, 0, LFROM | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_LAUTO, C_NONE, C_REG, 73, 8, REGSP, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_LOREG, C_NONE, C_REG, 73, 8, 0, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVBS, C_ADDR, C_NONE, C_REG, 93, 8, 0, LFROM | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_LAUTO, C_NONE, C_REG, 73, 8, REGSP, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_LOREG, C_NONE, C_REG, 73, 8, 0, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVH, C_ADDR, C_NONE, C_REG, 93, 8, 0, LFROM | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_LAUTO, C_NONE, C_REG, 73, 8, REGSP, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_LOREG, C_NONE, C_REG, 73, 8, 0, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHS, C_ADDR, C_NONE, C_REG, 93, 8, 0, LFROM | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_LAUTO, C_NONE, C_REG, 73, 8, REGSP, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_LOREG, C_NONE, C_REG, 73, 8, 0, LFROM, 0, C_PBIT | C_WBIT | C_UBIT},
		{AMOVHU, C_ADDR, C_NONE, C_REG, 93, 8, 0, LFROM | LPCREL, 4, C_PBIT | C_WBIT | C_UBIT},
		{ALDREX, C_SOREG, C_NONE, C_REG, 77, 4, 0, 0, 0, 0},
		{ASTREX, C_SOREG, C_REG, C_REG, 78, 4, 0, 0, 0, 0},
		{ADMB, C_NONE, C_NONE, C_NONE, 110, 4, 0, 0, 0, 0},
		{ADMB, C_LCON, C_NONE, C_NONE, 110, 4, 0, 0, 0, 0},
		{ADMB, C_SPR, C_NONE, C_NONE, 110, 4, 0, 0, 0, 0},
		{AMOVF, C_ZFCON, C_NONE, C_FREG, 80, 8, 0, 0, 0, 0},
		{AMOVF, C_SFCON, C_NONE, C_FREG, 81, 4, 0, 0, 0, 0},
		{ACMPF, C_FREG, C_FREG, C_NONE, 82, 8, 0, 0, 0, 0},
		{ACMPF, C_FREG, C_NONE, C_NONE, 83, 8, 0, 0, 0, 0},
		{AMOVFW, C_FREG, C_NONE, C_FREG, 84, 4, 0, 0, 0, C_UBIT},
		{AMOVWF, C_FREG, C_NONE, C_FREG, 85, 4, 0, 0, 0, C_UBIT},
		{AMOVFW, C_FREG, C_NONE, C_REG, 86, 8, 0, 0, 0, C_UBIT},
		{AMOVWF, C_REG, C_NONE, C_FREG, 87, 8, 0, 0, 0, C_UBIT},
		{AMOVW, C_REG, C_NONE, C_FREG, 88, 4, 0, 0, 0, 0},
		{AMOVW, C_FREG, C_NONE, C_REG, 89, 4, 0, 0, 0, 0},
		{ALDREXD, C_SOREG, C_NONE, C_REG, 91, 4, 0, 0, 0, 0},
		{ASTREXD, C_SOREG, C_REG, C_REG, 92, 4, 0, 0, 0, 0},
		{APLD, C_SOREG, C_NONE, C_NONE, 95, 4, 0, 0, 0, 0},
		{obj.AUNDEF, C_NONE, C_NONE, C_NONE, 96, 4, 0, 0, 0, 0},
		{ACLZ, C_REG, C_NONE, C_REG, 97, 4, 0, 0, 0, 0},
		{AMULWT, C_REG, C_REG, C_REG, 98, 4, 0, 0, 0, 0},
		{AMULA, C_REG, C_REG, C_REGREG2, 99, 4, 0, 0, 0, C_SBIT},
		{AMULAWT, C_REG, C_REG, C_REGREG2, 99, 4, 0, 0, 0, 0},
		{obj.APCDATA, C_LCON, C_NONE, C_LCON, 0, 0, 0, 0, 0, 0},
		{obj.AFUNCDATA, C_LCON, C_NONE, C_ADDR, 0, 0, 0, 0, 0, 0},
		{obj.ANOP, C_NONE, C_NONE, C_NONE, 0, 0, 0, 0, 0, 0},
		{obj.ADUFFZERO, C_NONE, C_NONE, C_SBRA, 5, 4, 0, 0, 0, 0},
		{obj.ADUFFCOPY, C_NONE, C_NONE, C_SBRA, 5, 4, 0, 0, 0, 0},
		{ADATABUNDLE, C_NONE, C_NONE, C_NONE, 100, 4, 0, 0, 0, 0},
		{ADATABUNDLEEND, C_NONE, C_NONE, C_NONE, 100, 0, 0, 0, 0, 0},
		{obj.AXXX, C_NONE, C_NONE, C_NONE, 0, 4, 0, 0, 0, 0},
	}
	psess.Linkarm = obj.LinkArch{
		Arch:           psess.sys.ArchARM,
		Init:           psess.buildop,
		Preprocess:     psess.preprocess,
		Assemble:       psess.span5,
		Progedit:       psess.progedit,
		UnaryDst:       psess.unaryDst,
		DWARFRegisters: psess.ARMDWARFRegisters,
	}
	psess.Anames = []string{
		obj.A_ARCHSPECIFIC: "AND",
		"EOR",
		"SUB",
		"RSB",
		"ADD",
		"ADC",
		"SBC",
		"RSC",
		"TST",
		"TEQ",
		"CMP",
		"CMN",
		"ORR",
		"BIC",
		"MVN",
		"BEQ",
		"BNE",
		"BCS",
		"BHS",
		"BCC",
		"BLO",
		"BMI",
		"BPL",
		"BVS",
		"BVC",
		"BHI",
		"BLS",
		"BGE",
		"BLT",
		"BGT",
		"BLE",
		"MOVWD",
		"MOVWF",
		"MOVDW",
		"MOVFW",
		"MOVFD",
		"MOVDF",
		"MOVF",
		"MOVD",
		"CMPF",
		"CMPD",
		"ADDF",
		"ADDD",
		"SUBF",
		"SUBD",
		"MULF",
		"MULD",
		"NMULF",
		"NMULD",
		"MULAF",
		"MULAD",
		"NMULAF",
		"NMULAD",
		"MULSF",
		"MULSD",
		"NMULSF",
		"NMULSD",
		"FMULAF",
		"FMULAD",
		"FNMULAF",
		"FNMULAD",
		"FMULSF",
		"FMULSD",
		"FNMULSF",
		"FNMULSD",
		"DIVF",
		"DIVD",
		"SQRTF",
		"SQRTD",
		"ABSF",
		"ABSD",
		"NEGF",
		"NEGD",
		"SRL",
		"SRA",
		"SLL",
		"MULU",
		"DIVU",
		"MUL",
		"MMUL",
		"DIV",
		"MOD",
		"MODU",
		"DIVHW",
		"DIVUHW",
		"MOVB",
		"MOVBS",
		"MOVBU",
		"MOVH",
		"MOVHS",
		"MOVHU",
		"MOVW",
		"MOVM",
		"SWPBU",
		"SWPW",
		"RFE",
		"SWI",
		"MULA",
		"MULS",
		"MMULA",
		"MMULS",
		"WORD",
		"MULL",
		"MULAL",
		"MULLU",
		"MULALU",
		"BX",
		"BXRET",
		"DWORD",
		"LDREX",
		"STREX",
		"LDREXD",
		"STREXD",
		"DMB",
		"PLD",
		"CLZ",
		"REV",
		"REV16",
		"REVSH",
		"RBIT",
		"XTAB",
		"XTAH",
		"XTABU",
		"XTAHU",
		"BFX",
		"BFXU",
		"BFC",
		"BFI",
		"MULWT",
		"MULWB",
		"MULBB",
		"MULAWT",
		"MULAWB",
		"MULABB",
		"DATABUNDLE",
		"DATABUNDLEEND",
		"MRC",
		"LAST",
	}
	return psess
}
