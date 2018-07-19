package mips

import (
	"debug/elf"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
)

func gentext(ctxt *ld.Link) {
	return
}

func adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	log.Fatalf("adddynrel not implemented")
	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write32(uint32(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		if r.Siz != 4 {
			return false
		}
		ctxt.Out.Write32(uint32(elf.R_MIPS_32) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_LO16) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPSU:
		ctxt.Out.Write32(uint32(elf.R_MIPS_HI16) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPSTLS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_TLS_TPREL_LO16) | uint32(elfsym)<<8)
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_26) | uint32(elfsym)<<8)
	}

	return true
}

func elfsetupplt(ctxt *ld.Link) {
	return
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	return false
}

func applyrel(arch *sys.Arch, r *sym.Reloc, s *sym.Symbol, val *int64, t int64) {
	o := arch.ByteOrder.Uint32(s.P[r.Off:])
	switch r.Type {
	case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSTLS:
		*val = int64(o&0xffff0000 | uint32(t)&0xffff)
	case objabi.R_ADDRMIPSU:
		*val = int64(o&0xffff0000 | uint32((t+(1<<15))>>16)&0xffff)
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		*val = int64(o&0xfc000000 | uint32(t>>2)&^0xfc000000)
	}
}

func (psess *PackageSession) archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val *int64) bool {
	if ctxt.LinkMode == ld.LinkExternal {
		switch r.Type {
		default:
			return false
		case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSU:
			r.Done = false

			rs := r.Sym
			r.Xadd = r.Add
			for rs.Outer != nil {
				r.Xadd += psess.ld.Symaddr(rs) - psess.ld.Symaddr(rs.Outer)
				rs = rs.Outer
			}

			if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Sect == nil {
				psess.ld.
					Errorf(s, "missing section for %s", rs.Name)
			}
			r.Xsym = rs
			applyrel(ctxt.Arch, r, s, val, r.Xadd)
			return true
		case objabi.R_ADDRMIPSTLS, objabi.R_CALLMIPS, objabi.R_JMPMIPS:
			r.Done = false
			r.Xsym = r.Sym
			r.Xadd = r.Add
			applyrel(ctxt.Arch, r, s, val, r.Add)
			return true
		}
	}

	switch r.Type {
	case objabi.R_CONST:
		*val = r.Add
		return true
	case objabi.R_GOTOFF:
		*val = psess.ld.Symaddr(r.Sym) + r.Add - psess.ld.Symaddr(ctxt.Syms.Lookup(".got", 0))
		return true
	case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSU:
		t := psess.ld.Symaddr(r.Sym) + r.Add
		applyrel(ctxt.Arch, r, s, val, t)
		return true
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		t := psess.ld.Symaddr(r.Sym) + r.Add

		if t&3 != 0 {
			psess.ld.
				Errorf(s, "direct call is not aligned: %s %x", r.Sym.Name, t)
		}

		if (s.Value+int64(r.Off)+4)&0xf0000000 != (t & 0xf0000000) {
			psess.ld.
				Errorf(s, "direct call too far: %s %x", r.Sym.Name, t)
		}

		applyrel(ctxt.Arch, r, s, val, t)
		return true
	case objabi.R_ADDRMIPSTLS:

		t := psess.ld.Symaddr(r.Sym) + r.Add - 0x7000
		if t < -32768 || t >= 32678 {
			psess.ld.
				Errorf(s, "TLS offset out of range %d", t)
		}
		applyrel(ctxt.Arch, r, s, val, t)
		return true
	}

	return false
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	return -1
}

func (psess *PackageSession) asmb(ctxt *ld.Link) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f asmb\n", psess.ld.Cputime())
	}

	if ctxt.IsELF {
		psess.ld.
			Asmbelfsetup()
	}

	sect := psess.ld.Segtext.Sections[0]
	ctxt.Out.SeekSet(psess.ld, int64(sect.Vaddr-psess.ld.Segtext.Vaddr+psess.ld.Segtext.Fileoff))
	psess.ld.
		Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	for _, sect = range psess.ld.Segtext.Sections[1:] {
		ctxt.Out.SeekSet(psess.ld, int64(sect.Vaddr-psess.ld.Segtext.Vaddr+psess.ld.Segtext.Fileoff))
		psess.ld.
			Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	}

	if psess.ld.Segrodata.Filelen > 0 {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f rodatblk\n", psess.ld.Cputime())
		}

		ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segrodata.Fileoff))
		psess.ld.
			Datblk(ctxt, int64(psess.ld.Segrodata.Vaddr), int64(psess.ld.Segrodata.Filelen))
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f datblk\n", psess.ld.Cputime())
	}

	ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segdata.Fileoff))
	psess.ld.
		Datblk(ctxt, int64(psess.ld.Segdata.Vaddr), int64(psess.ld.Segdata.Filelen))

	ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segdwarf.Fileoff))
	psess.ld.
		Dwarfblk(ctxt, int64(psess.ld.Segdwarf.Vaddr), int64(psess.ld.Segdwarf.Filelen))
	psess.ld.
		Symsize = 0
	psess.ld.
		Lcsize = 0
	symo := uint32(0)
	if !*psess.ld.FlagS {
		if !ctxt.IsELF {
			psess.ld.
				Errorf(nil, "unsupported executable format")
		}
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f sym\n", psess.ld.Cputime())
		}
		symo = uint32(psess.ld.Segdwarf.Fileoff + psess.ld.Segdwarf.Filelen)
		symo = uint32(ld.Rnd(int64(symo), int64(*psess.ld.FlagRound)))

		ctxt.Out.SeekSet(psess.ld, int64(symo))
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f elfsym\n", psess.ld.Cputime())
		}
		psess.ld.
			Asmelfsym(ctxt)
		ctxt.Out.Flush(psess.ld)
		ctxt.Out.Write(psess.ld.Elfstrdat)

		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f dwarf\n", psess.ld.Cputime())
		}

		if ctxt.LinkMode == ld.LinkExternal {
			psess.ld.
				Elfemitreloc(ctxt)
		}
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f header\n", psess.ld.Cputime())
	}

	ctxt.Out.SeekSet(psess.ld, 0)
	switch ctxt.HeadType {
	default:
		psess.ld.
			Errorf(nil, "unsupported operating system")
	case objabi.Hlinux:
		psess.ld.
			Asmbelf(ctxt, int64(symo))
	}

	ctxt.Out.Flush(psess.ld)
	if *psess.ld.FlagC {
		fmt.Printf("textsize=%d\n", psess.ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", psess.ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", psess.ld.Segdata.Length-psess.ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", psess.ld.Symsize)
		fmt.Printf("lcsize=%d\n", psess.ld.Lcsize)
		fmt.Printf("total=%d\n", psess.ld.Segtext.Filelen+psess.ld.Segdata.Length+uint64(psess.ld.Symsize)+uint64(psess.ld.Lcsize))
	}
}
