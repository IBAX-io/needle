package vm

func (rt *Runtime) Stack() []any {
	return rt.stack
}

func (rt *Runtime) push(d any) {
	rt.stack = append(rt.stack, d)
}

func (rt *Runtime) pop() (ret any) {
	ret = rt.peek()
	rt.stack = rt.stack[:rt.len()-1]
	return
}

func (rt *Runtime) len() int {
	return len(rt.stack)
}

func (rt *Runtime) swap(n int) {
	rt.stack[rt.len()-n], rt.stack[rt.len()-1] = rt.peek(), rt.stack[rt.len()-n]
}

func (rt *Runtime) dup(n int) {
	rt.push(&rt.stack[rt.len()-n])
}

func (rt *Runtime) peek() any {
	if rt.len() == 0 {
		return nil
	}
	return rt.stack[rt.len()-1]
}

func (rt *Runtime) getStack(idx int) any {
	if idx >= 0 && rt.len() > 0 && rt.len() > idx {
		return rt.stack[idx]
	}
	return nil
}

func (rt *Runtime) resetByIdx(idx int) {
	rt.stack = rt.stack[:idx]
}

func (rt *Runtime) popBlock() (ret *blockStack) {
	ret = rt.blocks[len(rt.blocks)-1]
	rt.blocks = rt.blocks[:len(rt.blocks)-1]
	return
}
