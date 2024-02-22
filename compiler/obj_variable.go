package compiler

// ObjInfoVariable object variable name name.
type ObjInfoVariable struct {
	Name  string
	Index int
}

func (*ObjInfoVariable) isObjInfoValue() {}

// ObjInfoExtendVariable object extend variable name.
type ObjInfoExtendVariable struct {
	Name string
}

func (*ObjInfoExtendVariable) isObjInfoValue() {}
