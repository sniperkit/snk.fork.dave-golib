package ld

import (
	"bytes"
	"debug/macho"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"reflect"
	"unsafe"
)

const (
	pageAlign = 12 // 4096 = 1 << 12
)

type loadCmd struct {
	Cmd macho.LoadCmd
	Len uint32
}

type dyldInfoCmd struct {
	Cmd                      macho.LoadCmd
	Len                      uint32
	RebaseOff, RebaseLen     uint32
	BindOff, BindLen         uint32
	WeakBindOff, WeakBindLen uint32
	LazyBindOff, LazyBindLen uint32
	ExportOff, ExportLen     uint32
}

type linkEditDataCmd struct {
	Cmd              macho.LoadCmd
	Len              uint32
	DataOff, DataLen uint32
}

type encryptionInfoCmd struct {
	Cmd                macho.LoadCmd
	Len                uint32
	CryptOff, CryptLen uint32
	CryptId            uint32
}

type loadCmdReader struct {
	offset, next int64
	f            *os.File
	order        binary.ByteOrder
}

func (r *loadCmdReader) Next() (cmd loadCmd, err error) {
	r.offset = r.next
	if _, err = r.f.Seek(r.offset, 0); err != nil {
		return
	}
	if err = binary.Read(r.f, r.order, &cmd); err != nil {
		return
	}
	r.next = r.offset + int64(cmd.Len)
	return
}

func (r loadCmdReader) ReadAt(offset int64, data interface{}) error {
	if _, err := r.f.Seek(r.offset+offset, 0); err != nil {
		return err
	}
	return binary.Read(r.f, r.order, data)
}

func (r loadCmdReader) WriteAt(offset int64, data interface{}) error {
	if _, err := r.f.Seek(r.offset+offset, 0); err != nil {
		return err
	}
	return binary.Write(r.f, r.order, data)
}

// machoCombineDwarf merges dwarf info generated by dsymutil into a macho executable.
// machoCombineDwarf returns true and skips merging if the input executable is for iOS.
//
// With internal linking, DWARF is embedded into the executable, this lets us do the
// same for external linking.
// inexe is the path to the executable with no DWARF. It must have enough room in the macho
// header to add the DWARF sections. (Use ld's -headerpad option)
// dsym is the path to the macho file containing DWARF from dsymutil.
// outexe is the path where the combined executable should be saved.
func (psess *PackageSession) machoCombineDwarf(inexe, dsym, outexe string, buildmode BuildMode) (bool, error) {
	exef, err := os.Open(inexe)
	if err != nil {
		return false, err
	}
	exem, err := macho.NewFile(exef)
	if err != nil {
		return false, err
	}
	cmdOffset := unsafe.Sizeof(exem.FileHeader)
	is64bit := exem.Magic == macho.Magic64
	if is64bit {

		cmdOffset += unsafe.Sizeof(exem.Magic)
	}

	reader := loadCmdReader{next: int64(cmdOffset), f: exef, order: exem.ByteOrder}
	for i := uint32(0); i < exem.Ncmd; i++ {
		cmd, err := reader.Next()
		if err != nil {
			return false, err
		}
		if cmd.Cmd == LC_VERSION_MIN_IPHONEOS {

			return true, nil
		}
	}
	dwarff, err := os.Open(dsym)
	if err != nil {
		return false, err
	}
	outf, err := os.Create(outexe)
	if err != nil {
		return false, err
	}
	outf.Chmod(0755)

	dwarfm, err := macho.NewFile(dwarff)
	if err != nil {
		return false, err
	}
	psess.
		linkseg = exem.Segment("__LINKEDIT")
	if psess.linkseg == nil {
		return false, fmt.Errorf("missing __LINKEDIT segment")
	}

	if _, err = exef.Seek(0, 0); err != nil {
		return false, err
	}
	if _, err := io.CopyN(outf, exef, int64(psess.linkseg.Offset)); err != nil {
		return false, err
	}
	psess.
		realdwarf = dwarfm.Segment("__DWARF")
	if psess.realdwarf == nil {
		return false, fmt.Errorf("missing __DWARF segment")
	}
	psess.
		dwarfstart = machoCalcStart(psess.realdwarf.Offset, psess.linkseg.Offset, pageAlign)
	if _, err = outf.Seek(psess.dwarfstart, 0); err != nil {
		return false, err
	}
	psess.
		dwarfaddr = int64((psess.linkseg.Addr + psess.linkseg.Memsz + 1<<pageAlign - 1) &^ (1<<pageAlign - 1))

	if _, err = dwarff.Seek(int64(psess.realdwarf.Offset), 0); err != nil {
		return false, err
	}
	if _, err := io.CopyN(outf, dwarff, int64(psess.realdwarf.Filesz)); err != nil {
		return false, err
	}

	if _, err = exef.Seek(int64(psess.linkseg.Offset), 0); err != nil {
		return false, err
	}
	psess.
		linkstart = machoCalcStart(psess.linkseg.Offset, uint64(psess.dwarfstart)+psess.realdwarf.Filesz, pageAlign)
	psess.
		linkoffset = uint32(psess.linkstart - int64(psess.linkseg.Offset))
	if _, err = outf.Seek(psess.linkstart, 0); err != nil {
		return false, err
	}
	if _, err := io.Copy(outf, exef); err != nil {
		return false, err
	}

	textsect := exem.Section("__text")
	if psess.linkseg == nil {
		return false, fmt.Errorf("missing __text section")
	}

	dwarfCmdOffset := int64(cmdOffset) + int64(exem.FileHeader.Cmdsz)
	availablePadding := int64(textsect.Offset) - dwarfCmdOffset
	if availablePadding < int64(psess.realdwarf.Len) {
		return false, fmt.Errorf("No room to add dwarf info. Need at least %d padding bytes, found %d", psess.realdwarf.Len, availablePadding)
	}

	if _, err = outf.Seek(dwarfCmdOffset, 0); err != nil {
		return false, err
	}
	if _, err := io.CopyN(outf, bytes.NewReader(psess.realdwarf.Raw()), int64(psess.realdwarf.Len)); err != nil {
		return false, err
	}

	if _, err = outf.Seek(int64(unsafe.Offsetof(exem.FileHeader.Ncmd)), 0); err != nil {
		return false, err
	}
	if err = binary.Write(outf, exem.ByteOrder, exem.Ncmd+1); err != nil {
		return false, err
	}
	if err = binary.Write(outf, exem.ByteOrder, exem.Cmdsz+psess.realdwarf.Len); err != nil {
		return false, err
	}

	reader = loadCmdReader{next: int64(cmdOffset), f: outf, order: exem.ByteOrder}
	for i := uint32(0); i < exem.Ncmd; i++ {
		cmd, err := reader.Next()
		if err != nil {
			return false, err
		}
		switch cmd.Cmd {
		case macho.LoadCmdSegment64:
			err = psess.machoUpdateSegment(reader, &macho.Segment64{}, &macho.Section64{})
		case macho.LoadCmdSegment:
			err = psess.machoUpdateSegment(reader, &macho.Segment32{}, &macho.Section32{})
		case LC_DYLD_INFO, LC_DYLD_INFO_ONLY:
			err = psess.machoUpdateLoadCommand(reader, &dyldInfoCmd{}, "RebaseOff", "BindOff", "WeakBindOff", "LazyBindOff", "ExportOff")
		case macho.LoadCmdSymtab:
			err = psess.machoUpdateLoadCommand(reader, &macho.SymtabCmd{}, "Symoff", "Stroff")
		case macho.LoadCmdDysymtab:
			err = psess.machoUpdateLoadCommand(reader, &macho.DysymtabCmd{}, "Tocoffset", "Modtaboff", "Extrefsymoff", "Indirectsymoff", "Extreloff", "Locreloff")
		case LC_CODE_SIGNATURE, LC_SEGMENT_SPLIT_INFO, LC_FUNCTION_STARTS, LC_DATA_IN_CODE, LC_DYLIB_CODE_SIGN_DRS:
			err = psess.machoUpdateLoadCommand(reader, &linkEditDataCmd{}, "DataOff")
		case LC_ENCRYPTION_INFO, LC_ENCRYPTION_INFO_64:
			err = psess.machoUpdateLoadCommand(reader, &encryptionInfoCmd{}, "CryptOff")
		case macho.LoadCmdDylib, macho.LoadCmdThread, macho.LoadCmdUnixThread, LC_PREBOUND_DYLIB, LC_UUID, LC_VERSION_MIN_MACOSX, LC_VERSION_MIN_IPHONEOS, LC_SOURCE_VERSION, LC_MAIN, LC_LOAD_DYLINKER, LC_LOAD_WEAK_DYLIB, LC_REEXPORT_DYLIB, LC_RPATH, LC_ID_DYLIB, LC_SYMSEG, LC_LOADFVMLIB, LC_IDFVMLIB, LC_IDENT, LC_FVMFILE, LC_PREPAGE, LC_ID_DYLINKER, LC_ROUTINES, LC_SUB_FRAMEWORK, LC_SUB_UMBRELLA, LC_SUB_CLIENT, LC_SUB_LIBRARY, LC_TWOLEVEL_HINTS, LC_PREBIND_CKSUM, LC_ROUTINES_64, LC_LAZY_LOAD_DYLIB, LC_LOAD_UPWARD_DYLIB, LC_DYLD_ENVIRONMENT, LC_LINKER_OPTION, LC_LINKER_OPTIMIZATION_HINT, LC_VERSION_MIN_TVOS, LC_VERSION_MIN_WATCHOS, LC_VERSION_NOTE, LC_BUILD_VERSION:

		default:
			err = fmt.Errorf("Unknown load command 0x%x (%s)\n", int(cmd.Cmd), cmd.Cmd)
		}
		if err != nil {
			return false, err
		}
	}
	return false, psess.machoUpdateDwarfHeader(&reader, buildmode)
}

// machoUpdateSegment updates the load command for a moved segment.
// Only the linkedit segment should move, and it should have 0 sections.
// seg should be a macho.Segment32 or macho.Segment64 as appropriate.
// sect should be a macho.Section32 or macho.Section64 as appropriate.
func (psess *PackageSession) machoUpdateSegment(r loadCmdReader, seg, sect interface{}) error {
	if err := r.ReadAt(0, seg); err != nil {
		return err
	}
	segValue := reflect.ValueOf(seg)
	offset := reflect.Indirect(segValue).FieldByName("Offset")

	if offset.Uint() < psess.linkseg.Offset {
		return nil
	}
	offset.SetUint(offset.Uint() + uint64(psess.linkoffset))
	if err := r.WriteAt(0, seg); err != nil {
		return err
	}

	return machoUpdateSections(r, segValue, reflect.ValueOf(sect), uint64(psess.linkoffset), 0)
}

func machoUpdateSections(r loadCmdReader, seg, sect reflect.Value, deltaOffset, deltaAddr uint64) error {
	iseg := reflect.Indirect(seg)
	nsect := iseg.FieldByName("Nsect").Uint()
	if nsect == 0 {
		return nil
	}
	sectOffset := int64(iseg.Type().Size())

	isect := reflect.Indirect(sect)
	offsetField := isect.FieldByName("Offset")
	reloffField := isect.FieldByName("Reloff")
	addrField := isect.FieldByName("Addr")
	sectSize := int64(isect.Type().Size())
	for i := uint64(0); i < nsect; i++ {
		if err := r.ReadAt(sectOffset, sect.Interface()); err != nil {
			return err
		}
		if offsetField.Uint() != 0 {
			offsetField.SetUint(offsetField.Uint() + deltaOffset)
		}
		if reloffField.Uint() != 0 {
			reloffField.SetUint(reloffField.Uint() + deltaOffset)
		}
		if addrField.Uint() != 0 {
			addrField.SetUint(addrField.Uint() + deltaAddr)
		}
		if err := r.WriteAt(sectOffset, sect.Interface()); err != nil {
			return err
		}
		sectOffset += sectSize
	}
	return nil
}

// machoUpdateDwarfHeader updates the DWARF segment load command.
func (psess *PackageSession) machoUpdateDwarfHeader(r *loadCmdReader, buildmode BuildMode) error {
	var seg, sect interface{}
	cmd, err := r.Next()
	if err != nil {
		return err
	}
	if cmd.Cmd == macho.LoadCmdSegment64 {
		seg = new(macho.Segment64)
		sect = new(macho.Section64)
	} else {
		seg = new(macho.Segment32)
		sect = new(macho.Section32)
	}
	if err := r.ReadAt(0, seg); err != nil {
		return err
	}
	segv := reflect.ValueOf(seg).Elem()

	segv.FieldByName("Offset").SetUint(uint64(psess.dwarfstart))
	segv.FieldByName("Addr").SetUint(uint64(psess.dwarfaddr))

	deltaOffset := uint64(psess.dwarfstart) - psess.realdwarf.Offset
	deltaAddr := uint64(psess.dwarfaddr) - psess.realdwarf.Addr

	if buildmode != BuildModeCShared {
		segv.FieldByName("Addr").SetUint(0)
		segv.FieldByName("Memsz").SetUint(0)
		deltaAddr = 0
	}

	if err := r.WriteAt(0, seg); err != nil {
		return err
	}
	return machoUpdateSections(*r, segv, reflect.ValueOf(sect), deltaOffset, deltaAddr)
}

func (psess *PackageSession) machoUpdateLoadCommand(r loadCmdReader, cmd interface{}, fields ...string) error {
	if err := r.ReadAt(0, cmd); err != nil {
		return err
	}
	value := reflect.Indirect(reflect.ValueOf(cmd))

	for _, name := range fields {
		field := value.FieldByName(name)
		fieldval := field.Uint()
		if fieldval >= psess.linkseg.Offset {
			field.SetUint(fieldval + uint64(psess.linkoffset))
		}
	}
	if err := r.WriteAt(0, cmd); err != nil {
		return err
	}
	return nil
}

func machoCalcStart(origAddr, newAddr uint64, alignExp uint32) int64 {
	align := uint64(1 << alignExp)
	if (origAddr % align) == (newAddr % align) {
		return int64(newAddr)
	}
	padding := (align - (newAddr % align))
	padding += origAddr % align
	return int64(padding + newAddr)
}
