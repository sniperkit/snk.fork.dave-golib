/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppc64

import (
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/compile/internal/gc"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/internal/obj"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/internal/obj/ppc64"
)

func (pstate *PackageState) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if cnt < int64(4*pstate.gc.Widthptr) {
		for i := int64(0); i < cnt; i += int64(pstate.gc.Widthptr) {
			p = pp.Appendpp(pstate.gc, p, ppc64.AMOVD, obj.TYPE_REG, ppc64.REGZERO, 0, obj.TYPE_MEM, ppc64.REGSP, pstate.gc.Ctxt.FixedFrameSize()+off+i)
		}
	} else if cnt <= int64(128*pstate.gc.Widthptr) {
		p = pp.Appendpp(pstate.gc, p, ppc64.AADD, obj.TYPE_CONST, 0, pstate.gc.Ctxt.FixedFrameSize()+off-8, obj.TYPE_REG, ppc64.REGRT1, 0)
		p.Reg = ppc64.REGSP
		p = pp.Appendpp(pstate.gc, p, obj.ADUFFZERO, obj.TYPE_NONE, 0, 0, obj.TYPE_MEM, 0, 0)
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = pstate.gc.Duffzero
		p.To.Offset = 4 * (128 - cnt/int64(pstate.gc.Widthptr))
	} else {
		p = pp.Appendpp(pstate.gc, p, ppc64.AMOVD, obj.TYPE_CONST, 0, pstate.gc.Ctxt.FixedFrameSize()+off-8, obj.TYPE_REG, ppc64.REGTMP, 0)
		p = pp.Appendpp(pstate.gc, p, ppc64.AADD, obj.TYPE_REG, ppc64.REGTMP, 0, obj.TYPE_REG, ppc64.REGRT1, 0)
		p.Reg = ppc64.REGSP
		p = pp.Appendpp(pstate.gc, p, ppc64.AMOVD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, ppc64.REGTMP, 0)
		p = pp.Appendpp(pstate.gc, p, ppc64.AADD, obj.TYPE_REG, ppc64.REGTMP, 0, obj.TYPE_REG, ppc64.REGRT2, 0)
		p.Reg = ppc64.REGRT1
		p = pp.Appendpp(pstate.gc, p, ppc64.AMOVDU, obj.TYPE_REG, ppc64.REGZERO, 0, obj.TYPE_MEM, ppc64.REGRT1, int64(pstate.gc.Widthptr))
		p1 := p
		p = pp.Appendpp(pstate.gc, p, ppc64.ACMP, obj.TYPE_REG, ppc64.REGRT1, 0, obj.TYPE_REG, ppc64.REGRT2, 0)
		p = pp.Appendpp(pstate.gc, p, ppc64.ABNE, obj.TYPE_NONE, 0, 0, obj.TYPE_BRANCH, 0, 0)
		pstate.gc.Patch(p, p1)
	}

	return p
}

func (pstate *PackageState) zeroAuto(pp *gc.Progs, n *gc.Node) {
	// Note: this code must not clobber any registers.
	sym := n.Sym.Linksym(pstate.types)
	size := n.Type.Size(pstate.types)
	for i := int64(0); i < size; i += 8 {
		p := pp.Prog(pstate.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = ppc64.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = ppc64.REGSP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (pstate *PackageState) ginsnop(pp *gc.Progs) {
	p := pp.Prog(pstate.gc, ppc64.AOR)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = ppc64.REG_R0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = ppc64.REG_R0
}

func (pstate *PackageState) ginsnop2(pp *gc.Progs) {
	// PPC64 is unusual because TWO nops are required
	// (see gc/cgen.go, gc/plive.go -- copy of comment below)
	//
	// On ppc64, when compiling Go into position
	// independent code on ppc64le we insert an
	// instruction to reload the TOC pointer from the
	// stack as well. See the long comment near
	// jmpdefer in runtime/asm_ppc64.s for why.
	// If the MOVD is not needed, insert a hardware NOP
	// so that the same number of instructions are used
	// on ppc64 in both shared and non-shared modes.

	pstate.ginsnop(pp)
	if pstate.gc.Ctxt.Flag_shared {
		p := pp.Prog(pstate.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_MEM
		p.From.Offset = 24
		p.From.Reg = ppc64.REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REG_R2
	} else {
		pstate.ginsnop(pp)
	}
}
