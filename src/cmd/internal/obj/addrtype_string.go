/*
Sniperkit-Bot
- Status: analyzed
*/

// Code generated by "stringer -type AddrType"; DO NOT EDIT.

package obj

import (
	"strconv"
)

const _AddrType_name = "TYPE_NONETYPE_BRANCHTYPE_TEXTSIZETYPE_MEMTYPE_CONSTTYPE_FCONSTTYPE_SCONSTTYPE_REGTYPE_ADDRTYPE_SHIFTTYPE_REGREGTYPE_REGREG2TYPE_INDIRTYPE_REGLIST"

func (i AddrType) String(pstate *PackageState) string {
	if i >= AddrType(len(pstate._AddrType_index)-1) {
		return "AddrType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AddrType_name[pstate._AddrType_index[i]:pstate._AddrType_index[i+1]]
}
