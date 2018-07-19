package objabi

import "strconv"

const _SymKind_name = "SxxxSTEXTSRODATASNOPTRDATASDATASBSSSNOPTRBSSSTLSBSSSDWARFINFOSDWARFRANGESDWARFLOCSDWARFMISC"

func (i SymKind) String(psess *PackageSession,) string {
	if i >= SymKind(len(psess._SymKind_index)-1) {
		return "SymKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SymKind_name[psess._SymKind_index[i]:psess._SymKind_index[i+1]]
}
