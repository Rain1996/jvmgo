package rtda

const (
	DefaultMaxStackSize = 1024
)

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread(maxStackSize uint) *Thread {
	if maxStackSize < DefaultMaxStackSize {
		maxStackSize = DefaultMaxStackSize
	}
	return &Thread{
		stack: newStack(maxStackSize),
	}
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPc(pc int) {
	t.pc = pc
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}
