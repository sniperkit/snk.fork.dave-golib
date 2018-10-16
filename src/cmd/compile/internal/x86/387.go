/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x86

import (
	"math"

	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/compile/internal/gc"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/compile/internal/ssa"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/compile/internal/types"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/internal/obj"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/internal/obj/x86"
)

// Generates code for v using 387 instructions.
func (pstate *PackageState) ssaGenValue387(s *gc.SSAGenState, v *ssa.Value) {
	// The SSA compiler pretends that it has an SSE backend.
	// If we don't have one of those, we need to translate
	// all the SSE ops to equivalent 387 ops. That's what this
	// function does.

	switch v.Op {
	case ssa.Op386MOVSSconst, ssa.Op386MOVSDconst:
		p := s.Prog(pstate.gc, pstate.loadPush(v.Type))
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		pstate.popAndSave(s, v)

	case ssa.Op386MOVSSconst2, ssa.Op386MOVSDconst2:
		p := s.Prog(pstate.gc, pstate.loadPush(v.Type))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		pstate.popAndSave(s, v)

	case ssa.Op386MOVSSload, ssa.Op386MOVSDload, ssa.Op386MOVSSloadidx1, ssa.Op386MOVSDloadidx1, ssa.Op386MOVSSloadidx4, ssa.Op386MOVSDloadidx8:
		p := s.Prog(pstate.gc, pstate.loadPush(v.Type))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		switch v.Op {
		case ssa.Op386MOVSSloadidx1, ssa.Op386MOVSDloadidx1:
			p.From.Scale = 1
			p.From.Index = v.Args[1].Reg(pstate.ssa)
			if p.From.Index == x86.REG_SP {
				p.From.Reg, p.From.Index = p.From.Index, p.From.Reg
			}
		case ssa.Op386MOVSSloadidx4:
			p.From.Scale = 4
			p.From.Index = v.Args[1].Reg(pstate.ssa)
		case ssa.Op386MOVSDloadidx8:
			p.From.Scale = 8
			p.From.Index = v.Args[1].Reg(pstate.ssa)
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		pstate.popAndSave(s, v)

	case ssa.Op386MOVSSstore, ssa.Op386MOVSDstore:
		// Push to-be-stored value on top of stack.
		pstate.push(s, v.Args[1])

		// Pop and store value.
		var op obj.As
		switch v.Op {
		case ssa.Op386MOVSSstore:
			op = x86.AFMOVFP
		case ssa.Op386MOVSDstore:
			op = x86.AFMOVDP
		}
		p := s.Prog(pstate.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)

	case ssa.Op386MOVSSstoreidx1, ssa.Op386MOVSDstoreidx1, ssa.Op386MOVSSstoreidx4, ssa.Op386MOVSDstoreidx8:
		pstate.push(s, v.Args[2])
		var op obj.As
		switch v.Op {
		case ssa.Op386MOVSSstoreidx1, ssa.Op386MOVSSstoreidx4:
			op = x86.AFMOVFP
		case ssa.Op386MOVSDstoreidx1, ssa.Op386MOVSDstoreidx8:
			op = x86.AFMOVDP
		}
		p := s.Prog(pstate.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)
		switch v.Op {
		case ssa.Op386MOVSSstoreidx1, ssa.Op386MOVSDstoreidx1:
			p.To.Scale = 1
			p.To.Index = v.Args[1].Reg(pstate.ssa)
			if p.To.Index == x86.REG_SP {
				p.To.Reg, p.To.Index = p.To.Index, p.To.Reg
			}
		case ssa.Op386MOVSSstoreidx4:
			p.To.Scale = 4
			p.To.Index = v.Args[1].Reg(pstate.ssa)
		case ssa.Op386MOVSDstoreidx8:
			p.To.Scale = 8
			p.To.Index = v.Args[1].Reg(pstate.ssa)
		}

	case ssa.Op386ADDSS, ssa.Op386ADDSD, ssa.Op386SUBSS, ssa.Op386SUBSD,
		ssa.Op386MULSS, ssa.Op386MULSD, ssa.Op386DIVSS, ssa.Op386DIVSD:
		if v.Reg(pstate.ssa) != v.Args[0].Reg(pstate.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(pstate.ssa))
		}

		// Push arg1 on top of stack
		pstate.push(s, v.Args[1])

		// Set precision if needed.  64 bits is the default.
		switch v.Op {
		case ssa.Op386ADDSS, ssa.Op386SUBSS, ssa.Op386MULSS, ssa.Op386DIVSS:
			p := s.Prog(pstate.gc, x86.AFSTCW)
			s.AddrScratch(pstate.gc, &p.To)
			p = s.Prog(pstate.gc, x86.AFLDCW)
			p.From.Type = obj.TYPE_MEM
			p.From.Name = obj.NAME_EXTERN
			p.From.Sym = pstate.gc.ControlWord32
		}

		var op obj.As
		switch v.Op {
		case ssa.Op386ADDSS, ssa.Op386ADDSD:
			op = x86.AFADDDP
		case ssa.Op386SUBSS, ssa.Op386SUBSD:
			op = x86.AFSUBDP
		case ssa.Op386MULSS, ssa.Op386MULSD:
			op = x86.AFMULDP
		case ssa.Op386DIVSS, ssa.Op386DIVSD:
			op = x86.AFDIVDP
		}
		p := s.Prog(pstate.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = s.SSEto387[v.Reg(pstate.ssa)] + 1

		// Restore precision if needed.
		switch v.Op {
		case ssa.Op386ADDSS, ssa.Op386SUBSS, ssa.Op386MULSS, ssa.Op386DIVSS:
			p := s.Prog(pstate.gc, x86.AFLDCW)
			s.AddrScratch(pstate.gc, &p.From)
		}

	case ssa.Op386UCOMISS, ssa.Op386UCOMISD:
		pstate.push(s, v.Args[0])

		// Compare.
		p := s.Prog(pstate.gc, x86.AFUCOMP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = s.SSEto387[v.Args[1].Reg(pstate.ssa)] + 1

		// Save AX.
		p = s.Prog(pstate.gc, x86.AMOVL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_AX
		s.AddrScratch(pstate.gc, &p.To)

		// Move status word into AX.
		p = s.Prog(pstate.gc, x86.AFSTSW)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_AX

		// Then move the flags we need to the integer flags.
		s.Prog(pstate.gc, x86.ASAHF)

		// Restore AX.
		p = s.Prog(pstate.gc, x86.AMOVL)
		s.AddrScratch(pstate.gc, &p.From)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_AX

	case ssa.Op386SQRTSD:
		pstate.push(s, v.Args[0])
		s.Prog(pstate.gc, x86.AFSQRT)
		pstate.popAndSave(s, v)

	case ssa.Op386FCHS:
		pstate.push(s, v.Args[0])
		s.Prog(pstate.gc, x86.AFCHS)
		pstate.popAndSave(s, v)

	case ssa.Op386CVTSL2SS, ssa.Op386CVTSL2SD:
		p := s.Prog(pstate.gc, x86.AMOVL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		s.AddrScratch(pstate.gc, &p.To)
		p = s.Prog(pstate.gc, x86.AFMOVL)
		s.AddrScratch(pstate.gc, &p.From)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		pstate.popAndSave(s, v)

	case ssa.Op386CVTTSD2SL, ssa.Op386CVTTSS2SL:
		pstate.push(s, v.Args[0])

		// Save control word.
		p := s.Prog(pstate.gc, x86.AFSTCW)
		s.AddrScratch(pstate.gc, &p.To)
		p.To.Offset += 4

		// Load control word which truncates (rounds towards zero).
		p = s.Prog(pstate.gc, x86.AFLDCW)
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_EXTERN
		p.From.Sym = pstate.gc.ControlWord64trunc

		// Now do the conversion.
		p = s.Prog(pstate.gc, x86.AFMOVLP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		s.AddrScratch(pstate.gc, &p.To)
		p = s.Prog(pstate.gc, x86.AMOVL)
		s.AddrScratch(pstate.gc, &p.From)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

		// Restore control word.
		p = s.Prog(pstate.gc, x86.AFLDCW)
		s.AddrScratch(pstate.gc, &p.From)
		p.From.Offset += 4

	case ssa.Op386CVTSS2SD:
		// float32 -> float64 is a nop
		pstate.push(s, v.Args[0])
		pstate.popAndSave(s, v)

	case ssa.Op386CVTSD2SS:
		// Round to nearest float32.
		pstate.push(s, v.Args[0])
		p := s.Prog(pstate.gc, x86.AFMOVFP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		s.AddrScratch(pstate.gc, &p.To)
		p = s.Prog(pstate.gc, x86.AFMOVF)
		s.AddrScratch(pstate.gc, &p.From)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		pstate.popAndSave(s, v)

	case ssa.OpLoadReg:
		if !v.Type.IsFloat() {
			pstate.ssaGenValue(s, v)
			return
		}
		// Load+push the value we need.
		p := s.Prog(pstate.gc, pstate.loadPush(v.Type))
		pstate.gc.AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		// Move the value to its assigned register.
		pstate.popAndSave(s, v)

	case ssa.OpStoreReg:
		if !v.Type.IsFloat() {
			pstate.ssaGenValue(s, v)
			return
		}
		pstate.push(s, v.Args[0])
		var op obj.As
		switch v.Type.Size(pstate.types) {
		case 4:
			op = x86.AFMOVFP
		case 8:
			op = x86.AFMOVDP
		}
		p := s.Prog(pstate.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		pstate.gc.AddrAuto(&p.To, v)

	case ssa.OpCopy:
		if !v.Type.IsFloat() {
			pstate.ssaGenValue(s, v)
			return
		}
		pstate.push(s, v.Args[0])
		pstate.popAndSave(s, v)

	case ssa.Op386CALLstatic, ssa.Op386CALLclosure, ssa.Op386CALLinter:
		pstate.flush387(s) // Calls must empty the FP stack.
		fallthrough        // then issue the call as normal
	default:
		pstate.ssaGenValue(s, v)
	}
}

// push pushes v onto the floating-point stack.  v must be in a register.
func (pstate *PackageState) push(s *gc.SSAGenState, v *ssa.Value) {
	p := s.Prog(pstate.gc, x86.AFMOVD)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = s.SSEto387[v.Reg(pstate.ssa)]
	p.To.Type = obj.TYPE_REG
	p.To.Reg = x86.REG_F0
}

// popAndSave pops a value off of the floating-point stack and stores
// it in the reigster assigned to v.
func (pstate *PackageState) popAndSave(s *gc.SSAGenState, v *ssa.Value) {
	r := v.Reg(pstate.ssa)
	if _, ok := s.SSEto387[r]; ok {
		// Pop value, write to correct register.
		p := s.Prog(pstate.gc, x86.AFMOVDP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = s.SSEto387[v.Reg(pstate.ssa)] + 1
	} else {
		// Don't actually pop value. This 387 register is now the
		// new home for the not-yet-assigned-a-home SSE register.
		// Increase the register mapping of all other registers by one.
		for rSSE, r387 := range s.SSEto387 {
			s.SSEto387[rSSE] = r387 + 1
		}
		s.SSEto387[r] = x86.REG_F0
	}
}

// loadPush returns the opcode for load+push of the given type.
func (pstate *PackageState) loadPush(t *types.Type) obj.As {
	if t.Size(pstate.types) == 4 {
		return x86.AFMOVF
	}
	return x86.AFMOVD
}

// flush387 removes all entries from the 387 floating-point stack.
func (pstate *PackageState) flush387(s *gc.SSAGenState) {
	for k := range s.SSEto387 {
		p := s.Prog(pstate.gc, x86.AFMOVDP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		delete(s.SSEto387, k)
	}
}

func (pstate *PackageState) ssaGenBlock387(s *gc.SSAGenState, b, next *ssa.Block) {
	// Empty the 387's FP stack before the block ends.
	pstate.flush387(s)

	pstate.ssaGenBlock(s, b, next)
}
