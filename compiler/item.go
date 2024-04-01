package compiler

const (
	MapConst = iota
	MapVar
	MapMap
	MapExtend
	MapArray

	MustKey
	MustColon
	MustComma
	MustValue
)

// MapItem represents a map item with a type and value.
type MapItem struct {
	Type  int
	Value isMapItemValue
}

type MapItemList []*MapItem

type isMapItemValue interface {
	isMapItemValue()
}

func (*Map) isMapItemValue()         {}
func (*VarInfo) isMapItemValue()     {}
func (*MapItemList) isMapItemValue() {}
func (*Lexeme) isMapItemValue()      {}

func NewMapItem(v isMapItemValue) *MapItem {
	var t int
	switch v.(type) {
	case *Map:
		t = MapMap
	case *MapItemList:
		t = MapArray
	case *Lexeme:
		switch v.(*Lexeme).Type {
		case EXTEND:
			t = MapExtend
		case NUMBER, LITERAL:
			t = MapConst
		}
	case *VarInfo:
		t = MapVar
	}
	return &MapItem{Type: t, Value: v}
}

func (m *MapItem) GetMap() *Map {
	if x, ok := m.Value.(*Map); ok {
		return x
	}
	return nil
}

func (m *MapItem) GetVarInfo() *VarInfo {
	if x, ok := m.Value.(*VarInfo); ok {
		return x
	}
	return nil
}

func (m *MapItem) GetMapItemList() *MapItemList {
	if x, ok := m.Value.(*MapItemList); ok {
		return x
	}
	return nil
}

func (m *MapItem) GetLexeme() *Lexeme {
	if x, ok := m.Value.(*Lexeme); ok {
		return x
	}
	return nil
}

const (
	SliceLow = iota
	SliceLowNum
	SliceHigh
	SliceHighNum
)

// SliceItem is a slice item, Index[0] is low, Index[1] is high.
type SliceItem struct {
	Index [2]int
}

// SetLow sets the low index of the SliceItem.
func (i *SliceItem) SetLow(low int) { i.Index[0] = low }

// SetHigh sets the high index of the SliceItem.
func (i *SliceItem) SetHigh(high int) { i.Index[1] = high }
