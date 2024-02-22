package vm

const (
	// ShiftContractId is the offset of identifiers
	ShiftContractId = 5000

	// MaxArrayIndex is the maximum index of an array
	MaxArrayIndex = 1000000

	// MaxMapCount is the maximum length of map
	MaxMapCount = 100000

	// MaxCallDepth is the maximum call depth of functions
	MaxCallDepth = 1000

	// MemoryLimit is the maximum memory limit of VM
	MemoryLimit = 128 << 20 // 128 MB
)
