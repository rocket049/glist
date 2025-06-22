package glist

type elementPool[T comparable] struct {
	element *GElement[T]
}

func (p *elementPool[T]) Get() *GElement[T] {
	if p.element == nil {
		return &GElement[T]{}
	}
	ret := p.element
	p.element = p.element.Next
	ret.Pre = nil
	ret.Next = nil
	return ret
}

// change e.Next
func (p *elementPool[T]) Put(e *GElement[T]) {
	e.Next = p.element
	p.element = e
}

func (p *elementPool[T]) PutList(e *GElement[T]) {
	t := e
	for {
		if t.Next == nil {
			break
		}
		t = t.Next
	}
	t.Next = p.element
	p.element = e
}
