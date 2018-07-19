package ssa

import (
	"math"
)

// checkFunc checks invariants of f.
func (psess *PackageSession) checkFunc(f *Func) {
	blockMark := make([]bool, f.NumBlocks())
	valueMark := make([]bool, f.NumValues())

	for _, b := range f.Blocks {
		if blockMark[b.ID] {
			f.Fatalf("block %s appears twice in %s!", b, f.Name)
		}
		blockMark[b.ID] = true
		if b.Func != f {
			f.Fatalf("%s.Func=%s, want %s", b, b.Func.Name, f.Name)
		}

		for i, e := range b.Preds {
			if se := e.b.Succs[e.i]; se.b != b || se.i != i {
				f.Fatalf("block pred/succ not crosslinked correctly %d:%s %d:%s", i, b, se.i, se.b)
			}
		}
		for i, e := range b.Succs {
			if pe := e.b.Preds[e.i]; pe.b != b || pe.i != i {
				f.Fatalf("block succ/pred not crosslinked correctly %d:%s %d:%s", i, b, pe.i, pe.b)
			}
		}

		switch b.Kind {
		case BlockExit:
			if len(b.Succs) != 0 {
				f.Fatalf("exit block %s has successors", b)
			}
			if b.Control == nil {
				f.Fatalf("exit block %s has no control value", b)
			}
			if !b.Control.Type.IsMemory(psess.types) {
				f.Fatalf("exit block %s has non-memory control value %s", b, b.Control.LongString(psess))
			}
		case BlockRet:
			if len(b.Succs) != 0 {
				f.Fatalf("ret block %s has successors", b)
			}
			if b.Control == nil {
				f.Fatalf("ret block %s has nil control", b)
			}
			if !b.Control.Type.IsMemory(psess.types) {
				f.Fatalf("ret block %s has non-memory control value %s", b, b.Control.LongString(psess))
			}
		case BlockRetJmp:
			if len(b.Succs) != 0 {
				f.Fatalf("retjmp block %s len(Succs)==%d, want 0", b, len(b.Succs))
			}
			if b.Control == nil {
				f.Fatalf("retjmp block %s has nil control", b)
			}
			if !b.Control.Type.IsMemory(psess.types) {
				f.Fatalf("retjmp block %s has non-memory control value %s", b, b.Control.LongString(psess))
			}
			if b.Aux == nil {
				f.Fatalf("retjmp block %s has nil Aux field", b)
			}
		case BlockPlain:
			if len(b.Succs) != 1 {
				f.Fatalf("plain block %s len(Succs)==%d, want 1", b, len(b.Succs))
			}
			if b.Control != nil {
				f.Fatalf("plain block %s has non-nil control %s", b, b.Control.LongString(psess))
			}
		case BlockIf:
			if len(b.Succs) != 2 {
				f.Fatalf("if block %s len(Succs)==%d, want 2", b, len(b.Succs))
			}
			if b.Control == nil {
				f.Fatalf("if block %s has no control value", b)
			}
			if !b.Control.Type.IsBoolean() {
				f.Fatalf("if block %s has non-bool control value %s", b, b.Control.LongString(psess))
			}
		case BlockDefer:
			if len(b.Succs) != 2 {
				f.Fatalf("defer block %s len(Succs)==%d, want 2", b, len(b.Succs))
			}
			if b.Control == nil {
				f.Fatalf("defer block %s has no control value", b)
			}
			if !b.Control.Type.IsMemory(psess.types) {
				f.Fatalf("defer block %s has non-memory control value %s", b, b.Control.LongString(psess))
			}
		case BlockFirst:
			if len(b.Succs) != 2 {
				f.Fatalf("plain/dead block %s len(Succs)==%d, want 2", b, len(b.Succs))
			}
			if b.Control != nil {
				f.Fatalf("plain/dead block %s has a control value", b)
			}
		}
		if len(b.Succs) != 2 && b.Likely != BranchUnknown {
			f.Fatalf("likeliness prediction %d for block %s with %d successors", b.Likely, b, len(b.Succs))
		}

		for _, v := range b.Values {

			nArgs := psess.opcodeTable[v.Op].argLen
			if nArgs != -1 && int32(len(v.Args)) != nArgs {
				f.Fatalf("value %s has %d args, expected %d", v.LongString(psess),
					len(v.Args), nArgs)
			}

			canHaveAux := false
			canHaveAuxInt := false
			switch psess.opcodeTable[v.Op].auxType {
			case auxNone:
			case auxBool:
				if v.AuxInt < 0 || v.AuxInt > 1 {
					f.Fatalf("bad bool AuxInt value for %v", v)
				}
				canHaveAuxInt = true
			case auxInt8:
				if v.AuxInt != int64(int8(v.AuxInt)) {
					f.Fatalf("bad int8 AuxInt value for %v", v)
				}
				canHaveAuxInt = true
			case auxInt16:
				if v.AuxInt != int64(int16(v.AuxInt)) {
					f.Fatalf("bad int16 AuxInt value for %v", v)
				}
				canHaveAuxInt = true
			case auxInt32:
				if v.AuxInt != int64(int32(v.AuxInt)) {
					f.Fatalf("bad int32 AuxInt value for %v", v)
				}
				canHaveAuxInt = true
			case auxInt64, auxFloat64:
				canHaveAuxInt = true
			case auxInt128:

			case auxFloat32:
				canHaveAuxInt = true
				if !psess.isExactFloat32(v) {
					f.Fatalf("value %v has an AuxInt value that is not an exact float32", v)
				}
			case auxString, auxSym, auxTyp:
				canHaveAux = true
			case auxSymOff, auxSymValAndOff, auxTypSize:
				canHaveAuxInt = true
				canHaveAux = true
			case auxSymInt32:
				if v.AuxInt != int64(int32(v.AuxInt)) {
					f.Fatalf("bad int32 AuxInt value for %v", v)
				}
				canHaveAuxInt = true
				canHaveAux = true
			case auxCCop:
				if _, ok := v.Aux.(Op); !ok {
					f.Fatalf("bad type %T for CCop in %v", v.Aux, v)
				}
				canHaveAux = true
			default:
				f.Fatalf("unknown aux type for %s", v.Op)
			}
			if !canHaveAux && v.Aux != nil {
				f.Fatalf("value %s has an Aux value %v but shouldn't", v.LongString(psess), v.Aux)
			}
			if !canHaveAuxInt && v.AuxInt != 0 {
				f.Fatalf("value %s has an AuxInt value %d but shouldn't", v.LongString(psess), v.AuxInt)
			}

			for i, arg := range v.Args {
				if arg == nil {
					f.Fatalf("value %s has nil arg", v.LongString(psess))
				}
				if v.Op != OpPhi {

					if arg.Type.IsMemory(psess.types) && i != len(v.Args)-1 {
						f.Fatalf("value %s has non-final memory arg (%d < %d)", v.LongString(psess), i, len(v.Args)-1)
					}
				}
			}

			if valueMark[v.ID] {
				f.Fatalf("value %s appears twice!", v.LongString(psess))
			}
			valueMark[v.ID] = true

			if v.Block != b {
				f.Fatalf("%s.block != %s", v, b)
			}
			if v.Op == OpPhi && len(v.Args) != len(b.Preds) {
				f.Fatalf("phi length %s does not match pred length %d for block %s", v.LongString(psess), len(b.Preds), b)
			}

			if v.Op == OpAddr {
				if len(v.Args) == 0 {
					f.Fatalf("no args for OpAddr %s", v.LongString(psess))
				}
				if v.Args[0].Op != OpSP && v.Args[0].Op != OpSB {
					f.Fatalf("bad arg to OpAddr %v", v)
				}
			}

			if f.RegAlloc != nil && f.Config.SoftFloat && v.Type.IsFloat() {
				f.Fatalf("unexpected floating-point type %v", v.LongString(psess))
			}

			switch c := f.Config; v.Op {
			case OpSP, OpSB:
				if v.Type != c.Types.Uintptr {
					f.Fatalf("bad %s type: want uintptr, have %s",
						v.Op, v.Type.String(psess.types))
				}
			}

		}
	}

	if !blockMark[f.Entry.ID] {
		f.Fatalf("entry block %v is missing", f.Entry)
	}
	for _, b := range f.Blocks {
		for _, c := range b.Preds {
			if !blockMark[c.b.ID] {
				f.Fatalf("predecessor block %v for %v is missing", c, b)
			}
		}
		for _, c := range b.Succs {
			if !blockMark[c.b.ID] {
				f.Fatalf("successor block %v for %v is missing", c, b)
			}
		}
	}

	if len(f.Entry.Preds) > 0 {
		f.Fatalf("entry block %s of %s has predecessor(s) %v", f.Entry, f.Name, f.Entry.Preds)
	}

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			for i, a := range v.Args {
				if !valueMark[a.ID] {
					f.Fatalf("%v, arg %d of %s, is missing", a, i, v.LongString(psess))
				}
			}
		}
		if b.Control != nil && !valueMark[b.Control.ID] {
			f.Fatalf("control value for %s is missing: %v", b, b.Control)
		}
	}
	for b := f.freeBlocks; b != nil; b = b.succstorage[0].b {
		if blockMark[b.ID] {
			f.Fatalf("used block b%d in free list", b.ID)
		}
	}
	for v := f.freeValues; v != nil; v = v.argstorage[0] {
		if valueMark[v.ID] {
			f.Fatalf("used value v%d in free list", v.ID)
		}
	}

	if f.RegAlloc == nil {

		sdom := f.sdom()
		for _, b := range f.Blocks {
			for _, v := range b.Values {
				for i, arg := range v.Args {
					x := arg.Block
					y := b
					if v.Op == OpPhi {
						y = b.Preds[i].b
					}
					if !domCheck(f, sdom, x, y) {
						f.Fatalf("arg %d of value %s does not dominate, arg=%s", i, v.LongString(psess), arg.LongString(psess))
					}
				}
			}
			if b.Control != nil && !domCheck(f, sdom, b.Control.Block, b) {
				f.Fatalf("control value %s for %s doesn't dominate", b.Control, b)
			}
		}
	}

	if f.RegAlloc == nil && f.pass != nil {
		ln := f.loopnest(psess)
		if !ln.hasIrreducible {
			po := f.postorder()
			for _, b := range po {
				for _, s := range b.Succs {
					bb := s.Block()
					if ln.b2l[b.ID] == nil && ln.b2l[bb.ID] != nil && bb != ln.b2l[bb.ID].header {
						f.Fatalf("block %s not in loop branches to non-header block %s in loop", b.String(), bb.String())
					}
					if ln.b2l[b.ID] != nil && ln.b2l[bb.ID] != nil && bb != ln.b2l[bb.ID].header && !ln.b2l[b.ID].isWithinOrEq(ln.b2l[bb.ID]) {
						f.Fatalf("block %s in loop branches to non-header block %s in non-containing loop", b.String(), bb.String())
					}
				}
			}
		}
	}

	uses := make([]int32, f.NumValues())
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			for _, a := range v.Args {
				uses[a.ID]++
			}
		}
		if b.Control != nil {
			uses[b.Control.ID]++
		}
	}
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Uses != uses[v.ID] {
				f.Fatalf("%s has %d uses, but has Uses=%d", v, uses[v.ID], v.Uses)
			}
		}
	}
	psess.
		memCheck(f)
}

func (psess *PackageSession) memCheck(f *Func) {

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Type.IsTuple() && v.Type.FieldType(psess.types, 0).IsMemory(psess.types) {
				f.Fatalf("memory is first in a tuple: %s\n", v.LongString(psess))
			}
		}
	}

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if (v.Op == OpCopy || v.Uses == 0) && v.Type.IsMemory(psess.types) {
				return
			}
		}
		if b != f.Entry && len(b.Preds) == 0 {
			return
		}
	}

	lastmem := make([]*Value, f.NumBlocks())
	ss := newSparseSet(f.NumValues())
	for _, b := range f.Blocks {

		ss.clear()
		for _, v := range b.Values {
			if v.Op == OpPhi || !v.Type.IsMemory(psess.types) {
				continue
			}
			if m := v.MemoryArg(psess); m != nil {
				ss.add(m.ID)
			}
		}

		for _, v := range b.Values {
			if !v.Type.IsMemory(psess.types) {
				continue
			}
			if ss.contains(v.ID) {
				continue
			}
			if lastmem[b.ID] != nil {
				f.Fatalf("two live memory values in %s: %s and %s", b, lastmem[b.ID], v)
			}
			lastmem[b.ID] = v
		}

		if lastmem[b.ID] == nil {
			for _, v := range b.Values {
				if v.Op == OpPhi {
					continue
				}
				m := v.MemoryArg(psess)
				if m == nil {
					continue
				}
				if lastmem[b.ID] != nil && lastmem[b.ID] != m {
					f.Fatalf("two live memory values in %s: %s and %s", b, lastmem[b.ID], m)
				}
				lastmem[b.ID] = m
			}
		}
	}

	for {
		changed := false
		for _, b := range f.Blocks {
			if lastmem[b.ID] != nil {
				continue
			}
			for _, e := range b.Preds {
				p := e.b
				if lastmem[p.ID] != nil {
					lastmem[b.ID] = lastmem[p.ID]
					changed = true
					break
				}
			}
		}
		if !changed {
			break
		}
	}

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op == OpPhi && v.Type.IsMemory(psess.types) {
				for i, a := range v.Args {
					if a != lastmem[b.Preds[i].b.ID] {
						f.Fatalf("inconsistent memory phi %s %d %s %s", v.LongString(psess), i, a, lastmem[b.Preds[i].b.ID])
					}
				}
			}
		}
	}

	if f.scheduled {
		for _, b := range f.Blocks {
			var mem *Value // the current live memory in the block
			for _, v := range b.Values {
				if v.Op == OpPhi {
					if v.Type.IsMemory(psess.types) {
						mem = v
					}
					continue
				}
				if mem == nil && len(b.Preds) > 0 {

					mem = lastmem[b.Preds[0].b.ID]
				}
				for _, a := range v.Args {
					if a.Type.IsMemory(psess.types) && a != mem {
						f.Fatalf("two live mems @ %s: %s and %s", v, mem, a)
					}
				}
				if v.Type.IsMemory(psess.types) {
					mem = v
				}
			}
		}
	}

	if f.scheduled {
		for _, b := range f.Blocks {
			seenNonPhi := false
			for _, v := range b.Values {
				switch v.Op {
				case OpPhi:
					if seenNonPhi {
						f.Fatalf("phi after non-phi @ %s: %s", b, v)
					}
				default:
					seenNonPhi = true
				}
			}
		}
	}
}

// domCheck reports whether x dominates y (including x==y).
func domCheck(f *Func, sdom SparseTree, x, y *Block) bool {
	if !sdom.isAncestorEq(f.Entry, y) {

		return true
	}
	return sdom.isAncestorEq(x, y)
}

// isExactFloat32 reports whether v has an AuxInt that can be exactly represented as a float32.
func (psess *PackageSession) isExactFloat32(v *Value) bool {
	x := v.AuxFloat(psess)
	return math.Float64bits(x) == math.Float64bits(float64(float32(x)))
}
