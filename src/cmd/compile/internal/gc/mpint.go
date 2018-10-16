/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"math/big"
)

// implements integer arithmetic

// Mpint represents an integer constant.
type Mpint struct {
	Val  big.Int
	Ovf  bool // set if Val overflowed compiler limit (sticky)
	Rune bool // set if syntax indicates default type rune
}

func (a *Mpint) SetOverflow() {
	a.Val.SetUint64(1) // avoid spurious div-zero errors
	a.Ovf = true
}

func (a *Mpint) checkOverflow(extra int) bool {
	// We don't need to be precise here, any reasonable upper limit would do.
	// For now, use existing limit so we pass all the tests unchanged.
	if a.Val.BitLen()+extra > Mpprec {
		a.SetOverflow()
	}
	return a.Ovf
}

func (a *Mpint) Set(b *Mpint) {
	a.Val.Set(&b.Val)
}

func (a *Mpint) SetFloat(b *Mpflt) bool {
	// avoid converting huge floating-point numbers to integers
	// (2*Mpprec is large enough to permit all tests to pass)
	if b.Val.MantExp(nil) > 2*Mpprec {
		a.SetOverflow()
		return false
	}

	if _, acc := b.Val.Int(&a.Val); acc == big.Exact {
		return true
	}

	const delta = 16 // a reasonably small number of bits > 0
	var t big.Float
	t.SetPrec(Mpprec - delta)

	// try rounding down a little
	t.SetMode(big.ToZero)
	t.Set(&b.Val)
	if _, acc := t.Int(&a.Val); acc == big.Exact {
		return true
	}

	// try rounding up a little
	t.SetMode(big.AwayFromZero)
	t.Set(&b.Val)
	if _, acc := t.Int(&a.Val); acc == big.Exact {
		return true
	}

	a.Ovf = false
	return false
}

func (a *Mpint) Add(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Add")
		}
		a.SetOverflow()
		return
	}

	a.Val.Add(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		pstate.yyerror("constant addition overflow")
	}
}

func (a *Mpint) Sub(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Sub")
		}
		a.SetOverflow()
		return
	}

	a.Val.Sub(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		pstate.yyerror("constant subtraction overflow")
	}
}

func (a *Mpint) Mul(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Mul")
		}
		a.SetOverflow()
		return
	}

	a.Val.Mul(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		pstate.yyerror("constant multiplication overflow")
	}
}

func (a *Mpint) Quo(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Quo")
		}
		a.SetOverflow()
		return
	}

	a.Val.Quo(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		// can only happen for div-0 which should be checked elsewhere
		pstate.yyerror("constant division overflow")
	}
}

func (a *Mpint) Rem(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Rem")
		}
		a.SetOverflow()
		return
	}

	a.Val.Rem(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		// should never happen
		pstate.yyerror("constant modulo overflow")
	}
}

func (a *Mpint) Or(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Or")
		}
		a.SetOverflow()
		return
	}

	a.Val.Or(&a.Val, &b.Val)
}

func (a *Mpint) And(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint And")
		}
		a.SetOverflow()
		return
	}

	a.Val.And(&a.Val, &b.Val)
}

func (a *Mpint) AndNot(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint AndNot")
		}
		a.SetOverflow()
		return
	}

	a.Val.AndNot(&a.Val, &b.Val)
}

func (a *Mpint) Xor(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Xor")
		}
		a.SetOverflow()
		return
	}

	a.Val.Xor(&a.Val, &b.Val)
}

func (a *Mpint) Lsh(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Lsh")
		}
		a.SetOverflow()
		return
	}

	s := b.Int64(pstate)
	if s < 0 || s >= Mpprec {
		msg := "shift count too large"
		if s < 0 {
			msg = "invalid negative shift count"
		}
		pstate.yyerror("%s: %d", msg, s)
		a.SetInt64(0)
		return
	}

	if a.checkOverflow(int(s)) {
		pstate.yyerror("constant shift overflow")
		return
	}
	a.Val.Lsh(&a.Val, uint(s))
}

func (a *Mpint) Rsh(pstate *PackageState, b *Mpint) {
	if a.Ovf || b.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("ovf in Mpint Rsh")
		}
		a.SetOverflow()
		return
	}

	s := b.Int64(pstate)
	if s < 0 {
		pstate.yyerror("invalid negative shift count: %d", s)
		if a.Val.Sign() < 0 {
			a.SetInt64(-1)
		} else {
			a.SetInt64(0)
		}
		return
	}

	a.Val.Rsh(&a.Val, uint(s))
}

func (a *Mpint) Cmp(b *Mpint) int {
	return a.Val.Cmp(&b.Val)
}

func (a *Mpint) CmpInt64(c int64) int {
	if c == 0 {
		return a.Val.Sign() // common case shortcut
	}
	return a.Val.Cmp(big.NewInt(c))
}

func (a *Mpint) Neg() {
	a.Val.Neg(&a.Val)
}

func (a *Mpint) Int64(pstate *PackageState) int64 {
	if a.Ovf {
		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.Fatalf("constant overflow")
		}
		return 0
	}

	return a.Val.Int64()
}

func (a *Mpint) SetInt64(c int64) {
	a.Val.SetInt64(c)
}

func (a *Mpint) SetString(pstate *PackageState, as string) {
	_, ok := a.Val.SetString(as, 0)
	if !ok {
		// required syntax is [+-][0[x]]d*
		// At the moment we lose precise error cause;
		// the old code distinguished between:
		// - malformed hex constant
		// - malformed octal constant
		// - malformed decimal constant
		// TODO(gri) use different conversion function
		pstate.yyerror("malformed integer constant: %s", as)
		a.Val.SetUint64(0)
		return
	}
	if a.checkOverflow(0) {
		pstate.yyerror("constant too large: %s", as)
	}
}

func (x *Mpint) String() string {
	return bconv(x, 0)
}

func bconv(xval *Mpint, flag FmtFlag) string {
	if flag&FmtSharp != 0 {
		return fmt.Sprintf("%#x", &xval.Val)
	}
	return xval.Val.String()
}
