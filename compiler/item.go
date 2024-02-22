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
	Value any
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
