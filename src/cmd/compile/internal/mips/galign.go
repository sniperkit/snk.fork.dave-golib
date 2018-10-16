/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips

import (
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/compile/internal/gc"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/compile/internal/ssa"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/internal/obj/mips"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/internal/objabi"
)

func (pstate *PackageState) Init(arch *gc.Arch) {
	arch.LinkArch = &pstate.mips.Linkmips
	if pstate.objabi.GOARCH == "mipsle" {
		arch.LinkArch = &pstate.mips.Linkmipsle
	}
	arch.REGSP = mips.REGSP
	arch.MAXWIDTH = (1 << 31) - 1
	arch.SoftFloat = (pstate.objabi.GOMIPS == "softfloat")
	arch.ZeroRange = pstate.zerorange
	arch.ZeroAuto = pstate.zeroAuto
	arch.Ginsnop = pstate.ginsnop
	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = pstate.ssaGenValue
	arch.SSAGenBlock = pstate.ssaGenBlock
}
