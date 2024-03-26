package compiler

// ObjInfoVariable is the structure for the local variable.
type ObjInfoVariable struct {
	// Name is the name of the local variable.
	Name string
	// Index is the position of the variable in the current block.
	Index int
}

// ObjInfoExtendVariable is the structure for the extended variable.
type ObjInfoExtendVariable struct {
	// Name is the name of the extended variable.
	Name string
}
