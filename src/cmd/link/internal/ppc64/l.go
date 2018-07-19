package ppc64

const (
	maxAlign  = 32 // max data alignment
	minAlign  = 1  // min data alignment
	funcAlign = 16
)

/* Used by ../internal/ld/dwarf.go */
const (
	dwarfRegSP = 1
	dwarfRegLR = 65
)
