package vm

type Stack struct {
	element []any
}

func NewStack() *Stack {
	return &Stack{element: make([]any, 0, 1024)}
}

func (rt *Stack) Stack() []any {
	return rt.element
}

func (rt *Stack) size() int {
	return len(rt.element)
}

func (rt *Stack) set(ind int, d any) {
	rt.element[ind] = d
}

func (rt *Stack) push(d any) {
	rt.element = append(rt.element, d)
}

func (rt *Stack) pushN(d []any) {
	rt.element = append(rt.element, d...)
	return
}

func (rt *Stack) get(idx int) any {
	if idx >= 0 && rt.size() > 0 && rt.size() > idx {
		return rt.element[idx]
	}
	return nil
}

func (rt *Stack) peek() any {
	if rt.size() == 0 {
		return nil
	}
	return rt.element[rt.size()-1]
}

func (rt *Stack) PeekN(n int) []any {
	sLen := len(rt.element)
	var el []any = nil
	if sLen >= n {
		el = rt.element[sLen-n:]
	}
	return el
}

// PeekFromTo return elements from index from to index to
func (rt *Stack) PeekFromTo(from, to int) []any {
	var el []any = nil
	if from >= 0 && to >= 0 && from <= to && rt.size() > to {
		el = rt.element[from:to]
	}
	return el
}

func (rt *Stack) pop() (ret any) {
	ret = rt.peek()
	rt.element = rt.element[:rt.size()-1]
	return
}

func (rt *Stack) PopN(n int) []any {
	sLen := len(rt.element)
	var el []any
	el = rt.element[sLen-n:]
	rt.element = rt.element[:sLen-n]

	//reverse to make sure the order
	for i, j := 0, len(el)-1; i < j; i, j = i+1, j-1 {
		el[i], el[j] = el[j], el[i]
	}
	return el
}

func (rt *Stack) swap(n int) {
	rt.element[rt.size()-n], rt.element[rt.size()-1] = rt.peek(), rt.element[rt.size()-n]
}

func (rt *Stack) dup(n int) {
	rt.push(&rt.element[rt.size()-n])
}

func (rt *Stack) resetByIdx(idx int) {
	rt.element = rt.element[:idx]
}

func (rt *Runtime) peekBlock() *blockStack {
	if len(rt.blocks) == 0 {
		return nil
	}
	return rt.blocks[len(rt.blocks)-1]
}

func (rt *Runtime) popBlock() (ret *blockStack) {
	ret = rt.blocks[len(rt.blocks)-1]
	rt.blocks = rt.blocks[:len(rt.blocks)-1]
	return
}

func (rt *Runtime) pushBlock(bs *blockStack) {
	rt.blocks = append(rt.blocks, bs)
}
