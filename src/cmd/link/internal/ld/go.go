/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go-specific code shared across loaders (5l, 6l, 8l).

package ld

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/internal/bio"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/internal/objabi"
	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/link/internal/sym"
)

// go-specific code shared across loaders (5l, 6l, 8l).

// replace all "". with pkg.
func expandpkg(t0 string, pkg string) string {
	return strings.Replace(t0, "\"\".", pkg+".", -1)
}

// TODO:
//	generate debugging section in binary.
//	once the dust settles, try to move some code to
//		libmach, so that other linkers and ar can share.

func (pstate *PackageState) ldpkg(ctxt *Link, f *bio.Reader, lib *sym.Library, length int64, filename string) {
	if *pstate.flagG {
		return
	}

	if int64(int(length)) != length {
		fmt.Fprintf(os.Stderr, "%s: too much pkg data in %s\n", os.Args[0], filename)
		if *pstate.flagU {
			pstate.errorexit()
		}
		return
	}

	bdata := make([]byte, length)
	if _, err := io.ReadFull(f, bdata); err != nil {
		fmt.Fprintf(os.Stderr, "%s: short pkg read %s\n", os.Args[0], filename)
		if *pstate.flagU {
			pstate.errorexit()
		}
		return
	}
	data := string(bdata)

	// process header lines
	for data != "" {
		var line string
		if i := strings.Index(data, "\n"); i >= 0 {
			line, data = data[:i], data[i+1:]
		} else {
			line, data = data, ""
		}
		if line == "safe" {
			lib.Safe = true
		}
		if line == "main" {
			lib.Main = true
		}
		if line == "" {
			break
		}
	}

	// look for cgo section
	p0 := strings.Index(data, "\n$$  // cgo")
	var p1 int
	if p0 >= 0 {
		p0 += p1
		i := strings.IndexByte(data[p0+1:], '\n')
		if i < 0 {
			fmt.Fprintf(os.Stderr, "%s: found $$ // cgo but no newline in %s\n", os.Args[0], filename)
			if *pstate.flagU {
				pstate.errorexit()
			}
			return
		}
		p0 += 1 + i

		p1 = strings.Index(data[p0:], "\n$$")
		if p1 < 0 {
			p1 = strings.Index(data[p0:], "\n!\n")
		}
		if p1 < 0 {
			fmt.Fprintf(os.Stderr, "%s: cannot find end of // cgo section in %s\n", os.Args[0], filename)
			if *pstate.flagU {
				pstate.errorexit()
			}
			return
		}
		p1 += p0

		pstate.loadcgo(ctxt, filename, objabi.PathToPrefix(lib.Pkg), data[p0:p1])
	}
}

func (pstate *PackageState) loadcgo(ctxt *Link, file string, pkg string, p string) {
	var directives [][]string
	if err := json.NewDecoder(strings.NewReader(p)).Decode(&directives); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s: failed decoding cgo directives: %v\n", os.Args[0], file, err)
		pstate.nerrors++
		return
	}

	for _, f := range directives {
		switch f[0] {
		case "cgo_import_dynamic":
			if len(f) < 2 || len(f) > 4 {
				break
			}

			local := f[1]
			remote := local
			if len(f) > 2 {
				remote = f[2]
			}
			lib := ""
			if len(f) > 3 {
				lib = f[3]
			}

			if *pstate.FlagD {
				fmt.Fprintf(os.Stderr, "%s: %s: cannot use dynamic imports with -d flag\n", os.Args[0], file)
				pstate.nerrors++
				return
			}

			if local == "_" && remote == "_" {
				// allow #pragma dynimport _ _ "foo.so"
				// to force a link of foo.so.
				pstate.havedynamic = 1

				if ctxt.HeadType == objabi.Hdarwin {
					pstate.machoadddynlib(lib, ctxt.LinkMode)
				} else {
					pstate.dynlib = append(pstate.dynlib, lib)
				}
				continue
			}

			local = expandpkg(local, pkg)
			q := ""
			if i := strings.Index(remote, "#"); i >= 0 {
				remote, q = remote[:i], remote[i+1:]
			}
			s := ctxt.Syms.Lookup(local, 0)
			if s.Type == 0 || s.Type == sym.SXREF || s.Type == sym.SHOSTOBJ {
				s.Dynimplib = lib
				s.Extname = remote
				s.Dynimpvers = q
				if s.Type != sym.SHOSTOBJ {
					s.Type = sym.SDYNIMPORT
				}
				pstate.havedynamic = 1
			}
			continue

		case "cgo_import_static":
			if len(f) != 2 {
				break
			}
			local := f[1]

			s := ctxt.Syms.Lookup(local, 0)
			s.Type = sym.SHOSTOBJ
			s.Size = 0
			continue

		case "cgo_export_static", "cgo_export_dynamic":
			if len(f) < 2 || len(f) > 3 {
				break
			}
			local := f[1]
			remote := local
			if len(f) > 2 {
				remote = f[2]
			}
			local = expandpkg(local, pkg)

			s := ctxt.Syms.Lookup(local, 0)

			switch ctxt.BuildMode {
			case BuildModeCShared, BuildModeCArchive, BuildModePlugin:
				if s == ctxt.Syms.Lookup("main", 0) {
					continue
				}
			}

			// export overrides import, for openbsd/cgo.
			// see issue 4878.
			if s.Dynimplib != "" {
				s.Dynimplib = ""
				s.Extname = ""
				s.Dynimpvers = ""
				s.Type = 0
			}

			if !s.Attr.CgoExport() {
				s.Extname = remote
				pstate.dynexp = append(pstate.dynexp, s)
			} else if s.Extname != remote {
				fmt.Fprintf(os.Stderr, "%s: conflicting cgo_export directives: %s as %s and %s\n", os.Args[0], s.Name, s.Extname, remote)
				pstate.nerrors++
				return
			}

			if f[0] == "cgo_export_static" {
				s.Attr |= sym.AttrCgoExportStatic
			} else {
				s.Attr |= sym.AttrCgoExportDynamic
			}
			continue

		case "cgo_dynamic_linker":
			if len(f) != 2 {
				break
			}

			if *pstate.flagInterpreter == "" {
				if pstate.interpreter != "" && pstate.interpreter != f[1] {
					fmt.Fprintf(os.Stderr, "%s: conflict dynlinker: %s and %s\n", os.Args[0], pstate.interpreter, f[1])
					pstate.nerrors++
					return
				}

				pstate.interpreter = f[1]
			}
			continue

		case "cgo_ldflag":
			if len(f) != 2 {
				break
			}
			pstate.ldflag = append(pstate.ldflag, f[1])
			continue
		}

		fmt.Fprintf(os.Stderr, "%s: %s: invalid cgo directive: %q\n", os.Args[0], file, f)
		pstate.nerrors++
	}
}

func (pstate *PackageState) adddynlib(ctxt *Link, lib string) {
	if pstate.seenlib[lib] || ctxt.LinkMode == LinkExternal {
		return
	}
	pstate.seenlib[lib] = true

	if ctxt.IsELF {
		s := ctxt.Syms.Lookup(".dynstr", 0)
		if s.Size == 0 {
			pstate.Addstring(s, "")
		}
		pstate.Elfwritedynent(ctxt, ctxt.Syms.Lookup(".dynamic", 0), DT_NEEDED, uint64(pstate.Addstring(s, lib)))
	} else {
		pstate.Errorf(nil, "adddynlib: unsupported binary format")
	}
}

func (pstate *PackageState) Adddynsym(ctxt *Link, s *sym.Symbol) {
	if s.Dynid >= 0 || ctxt.LinkMode == LinkExternal {
		return
	}

	if ctxt.IsELF {
		pstate.elfadddynsym(ctxt, s)
	} else if ctxt.HeadType == objabi.Hdarwin {
		pstate.Errorf(s, "adddynsym: missed symbol (Extname=%s)", s.Extname)
	} else if ctxt.HeadType == objabi.Hwindows {
		// already taken care of
	} else {
		pstate.Errorf(s, "adddynsym: unsupported binary format")
	}
}

func (pstate *PackageState) fieldtrack(ctxt *Link) {
	// record field tracking references
	var buf bytes.Buffer
	for _, s := range ctxt.Syms.Allsym {
		if strings.HasPrefix(s.Name, "go.track.") {
			s.Attr |= sym.AttrSpecial // do not lay out in data segment
			s.Attr |= sym.AttrNotInSymbolTable
			if s.Attr.Reachable() {
				buf.WriteString(s.Name[9:])
				for p := s.Reachparent; p != nil; p = p.Reachparent {
					buf.WriteString("\t")
					buf.WriteString(p.Name)
				}
				buf.WriteString("\n")
			}

			s.Type = sym.SCONST
			s.Value = 0
		}
	}

	if *pstate.flagFieldTrack == "" {
		return
	}
	s := ctxt.Syms.ROLookup(*pstate.flagFieldTrack, 0)
	if s == nil || !s.Attr.Reachable() {
		return
	}
	s.Type = sym.SDATA
	pstate.addstrdata(ctxt, *pstate.flagFieldTrack, buf.String())
}

func (ctxt *Link) addexport(pstate *PackageState) {
	if ctxt.HeadType == objabi.Hdarwin {
		return
	}

	for _, exp := range pstate.dynexp {
		pstate.Adddynsym(ctxt, exp)
	}
	for _, lib := range pstate.dynlib {
		pstate.adddynlib(ctxt, lib)
	}
}

type Pkg struct {
	mark    bool
	checked bool
	path    string
	impby   []*Pkg
}

func (p *Pkg) cycle(pstate *PackageState) *Pkg {
	if p.checked {
		return nil
	}

	if p.mark {
		pstate.nerrors++
		fmt.Printf("import cycle:\n")
		fmt.Printf("\t%s\n", p.path)
		return p
	}

	p.mark = true
	for _, q := range p.impby {
		if bad := q.cycle(pstate); bad != nil {
			p.mark = false
			p.checked = true
			fmt.Printf("\timports %s\n", p.path)
			if bad == p {
				return nil
			}
			return bad
		}
	}

	p.checked = true
	p.mark = false
	return nil
}

func (pstate *PackageState) importcycles() {
	for _, p := range pstate.pkgall {
		p.cycle(pstate)
	}
}
