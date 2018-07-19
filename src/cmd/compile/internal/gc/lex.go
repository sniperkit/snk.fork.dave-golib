package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/syntax"

	"github.com/dave/golib/src/cmd/internal/src"
	"strings"
)

// lineno is the source position at the start of the most recently lexed token.
// TODO(gri) rename and eventually remove

func (psess *PackageSession) makePos(base *src.PosBase, line, col uint) src.XPos {
	return psess.Ctxt.PosTable.XPos(src.MakePos(base, line, col))
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func isQuoted(s string) bool {
	return len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"'
}

const (
	// Func pragmas.
	Nointerface    syntax.Pragma = 1 << iota
	Noescape                     // func parameters don't escape
	Norace                       // func must not have race detector annotations
	Nosplit                      // func should not execute on separate stack
	Noinline                     // func should not be inlined
	CgoUnsafeArgs                // treat a pointer to one arg as a pointer to them all
	UintptrEscapes               // pointers converted to uintptr escape

	// Runtime-only func pragmas.
	// See ../../../../runtime/README.md for detailed descriptions.
	Systemstack        // func must run on system stack
	Nowritebarrier     // emit compiler error instead of write barrier
	Nowritebarrierrec  // error on write barrier in this or recursive callees
	Yeswritebarrierrec // cancels Nowritebarrierrec in this function and callees

	// Runtime-only type pragmas
	NotInHeap // values of this type must not be heap allocated
)

func (psess *PackageSession) pragmaValue(verb string) syntax.Pragma {
	switch verb {
	case "go:nointerface":
		if psess.objabi.Fieldtrack_enabled != 0 {
			return Nointerface
		}
	case "go:noescape":
		return Noescape
	case "go:norace":
		return Norace
	case "go:nosplit":
		return Nosplit
	case "go:noinline":
		return Noinline
	case "go:systemstack":
		return Systemstack
	case "go:nowritebarrier":
		return Nowritebarrier
	case "go:nowritebarrierrec":
		return Nowritebarrierrec | Nowritebarrier
	case "go:yeswritebarrierrec":
		return Yeswritebarrierrec
	case "go:cgo_unsafe_args":
		return CgoUnsafeArgs
	case "go:uintptrescapes":

		return UintptrEscapes
	case "go:notinheap":
		return NotInHeap
	}
	return 0
}

// pragcgo is called concurrently if files are parsed concurrently.
func (p *noder) pragcgo(pos syntax.Pos, text string) {
	f := pragmaFields(text)

	verb := strings.TrimPrefix(f[0], "go:")
	f[0] = verb

	switch verb {
	case "cgo_export_static", "cgo_export_dynamic":
		switch {
		case len(f) == 2 && !isQuoted(f[1]):
		case len(f) == 3 && !isQuoted(f[1]) && !isQuoted(f[2]):
		default:
			p.error(syntax.Error{Pos: pos, Msg: fmt.Sprintf("usage: //go:%s local [remote]", verb)})
			return
		}
	case "cgo_import_dynamic":
		switch {
		case len(f) == 2 && !isQuoted(f[1]):
		case len(f) == 3 && !isQuoted(f[1]) && !isQuoted(f[2]):
		case len(f) == 4 && !isQuoted(f[1]) && !isQuoted(f[2]) && isQuoted(f[3]):
			f[3] = strings.Trim(f[3], "\"")
		default:
			p.error(syntax.Error{Pos: pos, Msg: "usage: //go:cgo_import_dynamic local [remote [\"library\"]]"})
			return
		}
	case "cgo_import_static":
		switch {
		case len(f) == 2 && !isQuoted(f[1]):
		default:
			p.error(syntax.Error{Pos: pos, Msg: "usage: //go:cgo_import_static local"})
			return
		}
	case "cgo_dynamic_linker":
		switch {
		case len(f) == 2 && isQuoted(f[1]):
			f[1] = strings.Trim(f[1], "\"")
		default:
			p.error(syntax.Error{Pos: pos, Msg: "usage: //go:cgo_dynamic_linker \"path\""})
			return
		}
	case "cgo_ldflag":
		switch {
		case len(f) == 2 && isQuoted(f[1]):
			f[1] = strings.Trim(f[1], "\"")
		default:
			p.error(syntax.Error{Pos: pos, Msg: "usage: //go:cgo_ldflag \"arg\""})
			return
		}
	default:
		return
	}
	p.pragcgobuf = append(p.pragcgobuf, f)
}

// pragmaFields is similar to strings.FieldsFunc(s, isSpace)
// but does not split when inside double quoted regions and always
// splits before the start and after the end of a double quoted region.
// pragmaFields does not recognize escaped quotes. If a quote in s is not
// closed the part after the opening quote will not be returned as a field.
func pragmaFields(s string) []string {
	var a []string
	inQuote := false
	fieldStart := -1
	for i, c := range s {
		switch {
		case c == '"':
			if inQuote {
				inQuote = false
				a = append(a, s[fieldStart:i+1])
				fieldStart = -1
			} else {
				inQuote = true
				if fieldStart >= 0 {
					a = append(a, s[fieldStart:i])
				}
				fieldStart = i
			}
		case !inQuote && isSpace(c):
			if fieldStart >= 0 {
				a = append(a, s[fieldStart:i])
				fieldStart = -1
			}
		default:
			if fieldStart == -1 {
				fieldStart = i
			}
		}
	}
	if !inQuote && fieldStart >= 0 {
		a = append(a, s[fieldStart:])
	}
	return a
}
