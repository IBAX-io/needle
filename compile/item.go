package compile

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

// SliceItem is a slice item, Index[0] is low, Index[1] is high
type SliceItem struct {
	Index [2]int
}

func (i *SliceItem) SetLow(low int)   { i.Index[0] = low }
func (i *SliceItem) SetHigh(high int) { i.Index[1] = high }
