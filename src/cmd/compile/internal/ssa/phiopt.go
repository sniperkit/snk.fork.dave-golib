package ssa

// phiopt eliminates boolean Phis based on the previous if.
//
// Main use case is to transform:
//   x := false
//   if b {
//     x = true
//   }
// into x = b.
//
// In SSA code this appears as
//
// b0
//   If b -> b1 b2
// b1
//   Plain -> b2
// b2
//   x = (OpPhi (ConstBool [true]) (ConstBool [false]))
//
// In this case we can replace x with a copy of b.
func (psess *PackageSession) phiopt(f *Func) {
	sdom := f.sdom()
	for _, b := range f.Blocks {
		if len(b.Preds) != 2 || len(b.Values) == 0 {

			continue
		}

		pb0, b0 := b, b.Preds[0].b
		for len(b0.Succs) == 1 && len(b0.Preds) == 1 {
			pb0, b0 = b0, b0.Preds[0].b
		}
		if b0.Kind != BlockIf {
			continue
		}
		pb1, b1 := b, b.Preds[1].b
		for len(b1.Succs) == 1 && len(b1.Preds) == 1 {
			pb1, b1 = b1, b1.Preds[0].b
		}
		if b1 != b0 {
			continue
		}

		// reverse is the predecessor from which the truth value comes.
		var reverse int
		if b0.Succs[0].b == pb0 && b0.Succs[1].b == pb1 {
			reverse = 0
		} else if b0.Succs[0].b == pb1 && b0.Succs[1].b == pb0 {
			reverse = 1
		} else {
			b.Fatalf("invalid predecessors\n")
		}

		for _, v := range b.Values {
			if v.Op != OpPhi {
				continue
			}

			if v.Type.IsInteger() {
				psess.
					phioptint(v, b0, reverse)
			}

			if !v.Type.IsBoolean() {
				continue
			}

			if v.Args[0].Op == OpConstBool && v.Args[1].Op == OpConstBool {
				if v.Args[reverse].AuxInt != v.Args[1-reverse].AuxInt {
					ops := [2]Op{OpNot, OpCopy}
					v.reset(ops[v.Args[reverse].AuxInt])
					v.AddArg(b0.Control)
					if f.pass.debug > 0 {
						f.Warnl(b.Pos, "converted OpPhi to %v", v.Op)
					}
					continue
				}
			}

			if v.Args[reverse].Op == OpConstBool && v.Args[reverse].AuxInt == 1 {
				if tmp := v.Args[1-reverse]; sdom.isAncestorEq(tmp.Block, b) {
					v.reset(OpOrB)
					v.SetArgs2(b0.Control, tmp)
					if f.pass.debug > 0 {
						f.Warnl(b.Pos, "converted OpPhi to %v", v.Op)
					}
					continue
				}
			}

			if v.Args[1-reverse].Op == OpConstBool && v.Args[1-reverse].AuxInt == 0 {
				if tmp := v.Args[reverse]; sdom.isAncestorEq(tmp.Block, b) {
					v.reset(OpAndB)
					v.SetArgs2(b0.Control, tmp)
					if f.pass.debug > 0 {
						f.Warnl(b.Pos, "converted OpPhi to %v", v.Op)
					}
					continue
				}
			}
		}
	}
}

func (psess *PackageSession) phioptint(v *Value, b0 *Block, reverse int) {
	a0 := v.Args[0]
	a1 := v.Args[1]
	if a0.Op != a1.Op {
		return
	}

	switch a0.Op {
	case OpConst8, OpConst16, OpConst32, OpConst64:
	default:
		return
	}

	negate := false
	switch {
	case a0.AuxInt == 0 && a1.AuxInt == 1:
		negate = true
	case a0.AuxInt == 1 && a1.AuxInt == 0:
	default:
		return
	}

	if reverse == 1 {
		negate = !negate
	}

	switch v.Type.Size(psess.types) {
	case 1:
		v.reset(OpCopy)
	case 2:
		v.reset(OpZeroExt8to16)
	case 4:
		v.reset(OpZeroExt8to32)
	case 8:
		v.reset(OpZeroExt8to64)
	default:
		v.Fatalf("bad int size %d", v.Type.Size(psess.types))
	}

	a := b0.Control
	if negate {
		a = v.Block.NewValue1(v.Pos, OpNot, a.Type, a)
	}
	v.AddArg(a)

	f := b0.Func
	if f.pass.debug > 0 {
		f.Warnl(v.Block.Pos, "converted OpPhi bool -> int%d", v.Type.Size(psess.types)*8)
	}
}
