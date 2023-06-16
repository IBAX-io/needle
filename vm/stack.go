package vm

type Stack struct {
	element []any
}

func newStack() *Stack {
	return &Stack{element: make([]any, 0, 1024)}
}

func (s *Stack) Stack() []any {
	return s.element
}

func (s *Stack) size() int {
	return len(s.element)
}

func (s *Stack) set(ind int, d any) {
	s.element[ind] = d
}

func (s *Stack) reset() {
	s.element = s.element[:0]
}

func (s *Stack) push(d any) {
	s.element = append(s.element, d)
}

func (s *Stack) pushN(d []any) {
	s.element = append(s.element, d...)
	return
}

func (s *Stack) get(idx int) any {
	if idx >= 0 && s.size() > 0 && s.size() > idx {
		return s.element[idx]
	}
	return nil
}

func (s *Stack) getAndDel(idx int) any {
	if idx >= 0 && s.size() > 0 && s.size() > idx {
		ret := s.element[idx]
		s.element = append(s.element[:idx], s.element[idx+1:]...)
		return ret
	}
	return nil
}

func (s *Stack) peek() any {
	if s.size() == 0 {
		return nil
	}
	return s.element[s.size()-1]
}

func (s *Stack) peekN(n int) []any {
	sLen := len(s.element)
	var el []any = nil
	if sLen >= n {
		el = s.element[sLen-n:]
	}
	return el
}

// peekFromTo return elements from index from to index to
func (s *Stack) peekFromTo(from, to int) []any {
	var el []any = nil
	if from >= 0 && to >= 0 && from <= to && s.size() >= to {
		el = s.element[from:to]
	}
	return el
}

func (s *Stack) pop() (ret any) {
	ret = s.peek()
	s.element = s.element[:s.size()-1]
	return
}

func (s *Stack) popN(n int) []any {
	sLen := len(s.element)
	var el []any
	el = s.element[sLen-n:]
	s.element = s.element[:sLen-n]

	//reverse to make sure the order
	for i, j := 0, len(el)-1; i < j; i, j = i+1, j-1 {
		el[i], el[j] = el[j], el[i]
	}
	return el
}

func (s *Stack) swap(n int) {
	s.element[s.size()-n], s.element[s.size()-1] = s.peek(), s.element[s.size()-n]
}

func (s *Stack) dup(n int) {
	s.push(&s.element[s.size()-n])
}

func (s *Stack) resetByIdx(idx int) {
	if idx < 0 || idx > s.size() {
		idx = 0
	}
	s.element = s.element[:idx]
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
