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
