package obj

import (
	"github.com/dave/golib/src/cmd/internal/src"
	"log"
)

const (
	PrologueEnd   = 2 + iota // overload "is_stmt" to include prologue_end
	EpilogueBegin            // overload "is_stmt" to include epilogue_end
)

func addvarint(d *Pcdata, v uint32) {
	for ; v >= 0x80; v >>= 7 {
		d.P = append(d.P, uint8(v|0x80))
	}
	d.P = append(d.P, uint8(v))
}

// funcpctab writes to dst a pc-value table mapping the code in func to the values
// returned by valfunc parameterized by arg. The invocation of valfunc to update the
// current value is, for each p,
//
//	val = valfunc(func, val, p, 0, arg);
//	record val as value at p->pc;
//	val = valfunc(func, val, p, 1, arg);
//
// where func is the function, val is the current value, p is the instruction being
// considered, and arg can be used to further parameterize valfunc.
func funcpctab(ctxt *Link, dst *Pcdata, func_ *LSym, desc string, valfunc func(*Link, *LSym, int32, *Prog, int32, interface{}) int32, arg interface{}) {
	dbg := desc == ctxt.Debugpcln

	dst.P = dst.P[:0]

	if dbg {
		ctxt.Logf("funcpctab %s [valfunc=%s]\n", func_.Name, desc)
	}

	val := int32(-1)
	oldval := val
	if func_.Func.Text == nil {
		return
	}

	pc := func_.Func.Text.Pc

	if dbg {
		ctxt.Logf("%6x %6d %v\n", uint64(pc), val, func_.Func.Text)
	}

	started := false
	var delta uint32
	for p := func_.Func.Text; p != nil; p = p.Link {

		val = valfunc(ctxt, func_, val, p, 0, arg)

		if val == oldval && started {
			val = valfunc(ctxt, func_, val, p, 1, arg)
			if dbg {
				ctxt.Logf("%6x %6s %v\n", uint64(p.Pc), "", p)
			}
			continue
		}

		if p.Link != nil && p.Link.Pc == p.Pc {
			val = valfunc(ctxt, func_, val, p, 1, arg)
			if dbg {
				ctxt.Logf("%6x %6s %v\n", uint64(p.Pc), "", p)
			}
			continue
		}

		if dbg {
			ctxt.Logf("%6x %6d %v\n", uint64(p.Pc), val, p)
		}

		if started {
			addvarint(dst, uint32((p.Pc-pc)/int64(ctxt.Arch.MinLC)))
			pc = p.Pc
		}

		delta = uint32(val) - uint32(oldval)
		if delta>>31 != 0 {
			delta = 1 | ^(delta << 1)
		} else {
			delta <<= 1
		}
		addvarint(dst, delta)
		oldval = val
		started = true
		val = valfunc(ctxt, func_, val, p, 1, arg)
	}

	if started {
		if dbg {
			ctxt.Logf("%6x done\n", uint64(func_.Func.Text.Pc+func_.Size))
		}
		addvarint(dst, uint32((func_.Size-pc)/int64(ctxt.Arch.MinLC)))
		addvarint(dst, 0)
	}

	if dbg {
		ctxt.Logf("wrote %d bytes to %p\n", len(dst.P), dst)
		for _, p := range dst.P {
			ctxt.Logf(" %02x", p)
		}
		ctxt.Logf("\n")
	}
}

// pctofileline computes either the file number (arg == 0)
// or the line number (arg == 1) to use at p.
// Because p.Pos applies to p, phase == 0 (before p)
// takes care of the update.
func (psess *PackageSession) pctofileline(ctxt *Link, sym *LSym, oldval int32, p *Prog, phase int32, arg interface{}) int32 {
	if p.As == ATEXT || p.As == ANOP || p.Pos.Line() == 0 || phase == 1 {
		return oldval
	}
	f, l := psess.linkgetlineFromPos(ctxt, p.Pos)
	if arg == nil {
		return l
	}
	pcln := arg.(*Pcln)

	if f == pcln.Lastfile {
		return int32(pcln.Lastindex)
	}

	for i, file := range pcln.File {
		if file == f {
			pcln.Lastfile = f
			pcln.Lastindex = i
			return int32(i)
		}
	}
	i := len(pcln.File)
	pcln.File = append(pcln.File, f)
	pcln.Lastfile = f
	pcln.Lastindex = i
	return int32(i)
}

// pcinlineState holds the state used to create a function's inlining
// tree and the PC-value table that maps PCs to nodes in that tree.
type pcinlineState struct {
	globalToLocal map[int]int
	localTree     InlTree
}

// addBranch adds a branch from the global inlining tree in ctxt to
// the function's local inlining tree, returning the index in the local tree.
func (s *pcinlineState) addBranch(ctxt *Link, globalIndex int) int {
	if globalIndex < 0 {
		return -1
	}

	localIndex, ok := s.globalToLocal[globalIndex]
	if ok {
		return localIndex
	}

	call := ctxt.InlTree.nodes[globalIndex]
	call.Parent = s.addBranch(ctxt, call.Parent)
	localIndex = len(s.localTree.nodes)
	s.localTree.nodes = append(s.localTree.nodes, call)
	s.globalToLocal[globalIndex] = localIndex
	return localIndex
}

// pctoinline computes the index into the local inlining tree to use at p.
// If p is not the result of inlining, pctoinline returns -1. Because p.Pos
// applies to p, phase == 0 (before p) takes care of the update.
func (s *pcinlineState) pctoinline(ctxt *Link, sym *LSym, oldval int32, p *Prog, phase int32, arg interface{}) int32 {
	if phase == 1 {
		return oldval
	}

	posBase := ctxt.PosTable.Pos(p.Pos).Base()
	if posBase == nil {
		return -1
	}

	globalIndex := posBase.InliningIndex()
	if globalIndex < 0 {
		return -1
	}

	if s.globalToLocal == nil {
		s.globalToLocal = make(map[int]int)
	}

	return int32(s.addBranch(ctxt, globalIndex))
}

// pctospadj computes the sp adjustment in effect.
// It is oldval plus any adjustment made by p itself.
// The adjustment by p takes effect only after p, so we
// apply the change during phase == 1.
func pctospadj(ctxt *Link, sym *LSym, oldval int32, p *Prog, phase int32, arg interface{}) int32 {
	if oldval == -1 {
		oldval = 0
	}
	if phase == 0 {
		return oldval
	}
	if oldval+p.Spadj < -10000 || oldval+p.Spadj > 1100000000 {
		ctxt.Diag("overflow in spadj: %d + %d = %d", oldval, p.Spadj, oldval+p.Spadj)
		ctxt.DiagFlush()
		log.Fatalf("bad code")
	}

	return oldval + p.Spadj
}

// pctostmt returns either,
// if phase==0, then whether the current instruction is a step-target (Dwarf is_stmt)
//     bit-or'd with whether the current statement is a prologue end or epilogue begin
// else (phase == 1), zero.
//
func pctostmt(ctxt *Link, sym *LSym, oldval int32, p *Prog, phase int32, arg interface{}) int32 {
	if phase == 1 {
		return 0
	}
	s := p.Pos.IsStmt()
	l := p.Pos.Xlogue()

	var is_stmt int32

	switch l {
	case src.PosPrologueEnd:
		is_stmt = PrologueEnd
	case src.PosEpilogueBegin:
		is_stmt = EpilogueBegin
	}

	if s != src.PosNotStmt {
		is_stmt |= 1
	}
	return is_stmt
}

// pctopcdata computes the pcdata value in effect at p.
// A PCDATA instruction sets the value in effect at future
// non-PCDATA instructions.
// Since PCDATA instructions have no width in the final code,
// it does not matter which phase we use for the update.
func pctopcdata(ctxt *Link, sym *LSym, oldval int32, p *Prog, phase int32, arg interface{}) int32 {
	if phase == 0 || p.As != APCDATA || p.From.Offset != int64(arg.(uint32)) {
		return oldval
	}
	if int64(int32(p.To.Offset)) != p.To.Offset {
		ctxt.Diag("overflow in PCDATA instruction: %v", p)
		ctxt.DiagFlush()
		log.Fatalf("bad code")
	}

	return int32(p.To.Offset)
}

// stmtData writes out pc-linked is_stmt data for eventual use in the DWARF line numbering table.
func stmtData(ctxt *Link, cursym *LSym) {
	var pctostmtData Pcdata
	funcpctab(ctxt, &pctostmtData, cursym, "pctostmt", pctostmt, nil)
	cursym.Func.dwarfIsStmtSym.P = pctostmtData.P
}

func (psess *PackageSession) linkpcln(ctxt *Link, cursym *LSym) {
	pcln := &cursym.Func.Pcln

	npcdata := 0
	nfuncdata := 0
	for p := cursym.Func.Text; p != nil; p = p.Link {

		if p.As == APCDATA && p.From.Offset >= int64(npcdata) && p.To.Offset != -1 {
			npcdata = int(p.From.Offset + 1)
		}

		if p.As == AFUNCDATA && p.From.Offset >= int64(nfuncdata) {
			nfuncdata = int(p.From.Offset + 1)
		}
	}

	pcln.Pcdata = make([]Pcdata, npcdata)
	pcln.Pcdata = pcln.Pcdata[:npcdata]
	pcln.Funcdata = make([]*LSym, nfuncdata)
	pcln.Funcdataoff = make([]int64, nfuncdata)
	pcln.Funcdataoff = pcln.Funcdataoff[:nfuncdata]

	funcpctab(ctxt, &pcln.Pcsp, cursym, "pctospadj", pctospadj, nil)
	funcpctab(ctxt, &pcln.Pcfile, cursym, "pctofile", psess.pctofileline, pcln)
	funcpctab(ctxt, &pcln.Pcline, cursym, "pctoline", psess.pctofileline, nil)

	pcinlineState := new(pcinlineState)
	funcpctab(ctxt, &pcln.Pcinline, cursym, "pctoinline", pcinlineState.pctoinline, nil)
	pcln.InlTree = pcinlineState.localTree
	if ctxt.Debugpcln == "pctoinline" && len(pcln.InlTree.nodes) > 0 {
		ctxt.Logf("-- inlining tree for %s:\n", cursym)
		dumpInlTree(ctxt, pcln.InlTree)
		ctxt.Logf("--\n")
	}

	havepc := make([]uint32, (npcdata+31)/32)
	havefunc := make([]uint32, (nfuncdata+31)/32)
	for p := cursym.Func.Text; p != nil; p = p.Link {
		if p.As == AFUNCDATA {
			if (havefunc[p.From.Offset/32]>>uint64(p.From.Offset%32))&1 != 0 {
				ctxt.Diag("multiple definitions for FUNCDATA $%d", p.From.Offset)
			}
			havefunc[p.From.Offset/32] |= 1 << uint64(p.From.Offset%32)
		}

		if p.As == APCDATA && p.To.Offset != -1 {
			havepc[p.From.Offset/32] |= 1 << uint64(p.From.Offset%32)
		}
	}

	for i := 0; i < npcdata; i++ {
		if (havepc[i/32]>>uint(i%32))&1 == 0 {
			continue
		}
		funcpctab(ctxt, &pcln.Pcdata[i], cursym, "pctopcdata", pctopcdata, interface{}(uint32(i)))
	}

	if nfuncdata > 0 {
		for p := cursym.Func.Text; p != nil; p = p.Link {
			if p.As != AFUNCDATA {
				continue
			}
			i := int(p.From.Offset)
			pcln.Funcdataoff[i] = p.To.Offset
			if p.To.Type != TYPE_CONST {

				pcln.Funcdata[i] = p.To.Sym
			}
		}
	}
}
