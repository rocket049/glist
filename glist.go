package glist

// generics element type of GList
type GElement[T comparable] struct {
	Data  T
	Empty bool
	Pre   *GElement[T]
	Next  *GElement[T]
}

// generics list type
type GList[T comparable] struct {
	element *GElement[T]
}

// use it like this: listNew := NewGList[TypeName]()
func NewGList[T comparable]() *GList[T] {
	return &GList[T]{element: &GElement[T]{Empty: true}}
}

// get head of GList
func (p *GList[T]) Front() *GElement[T] {
	if p.element.End() {
		return nil
	}
	return p.element
}

// search one element of GList
func (p *GList[T]) SearchOne(item T) *GElement[T] {
	e := p.element
	for {
		if e.End() {
			return nil
		}
		if e.Data == item {
			return e
		}
		e = e.Next
	}
	return nil
}

// search all elements of GList
func (p *GList[T]) SearchAll(item T) []*GElement[T] {
	ret := []*GElement[T]{}
	e := p.element
	for {
		if e.End() {
			break
		}
		if e.Data == item {
			ret = append(ret, e)
		}
		e = e.Next
	}
	return ret
}

// Add: insert item at front, this is fastest
func (p *GList[T]) Add(item T) {
	if p.element.Empty {
		p.element.Data = item
		p.element.Empty = false
		return
	}
	e := &GElement[T]{Data: item, Empty: false}
	e.Next = p.element
	p.element.Pre = e
	p.element = e
}

// AddUnique: insert item at front if it not exists.
func (p *GList[T]) AddUnique(item T) {
	if p.element.Empty {
		p.element.Data = item
		p.element.Empty = false
		return
	}

	//search
	p1 := p.element
	for {
		if p1.Data == item {
			return
		}
		if p1.Next.End() {
			break
		}
		p1 = p1.Next
	}
	//end search

	e := &GElement[T]{Data: item, Empty: false}
	e.Next = p.element
	p.element.Pre = e
	p.element = e
}

// test the list is empty
func (p *GList[T]) Empty() bool {
	return p.element.Empty
}

// remove GLement e
func (p *GList[T]) Remove(e *GElement[T]) {
	e.Empty = true

	if p.element == e {
		if e.Next.End() {
			return
		}
		p.element = e.Next
		return
	}

	if e.Next != nil {
		e.Next.Pre = e.Pre
	}

	e.Pre.Next = e.Next

}

// remove all items of the list
func (p *GList[T]) Clear() {
	p1 := p.element
	for {

		if p1.End() {
			break
		}
		p1.Empty = true
		p1 = p1.Next
	}
}

// Append: append item at the end
func (p *GList[T]) Append(item T) {
	e := p.element
	if e.Empty {
		e.Data = item
		e.Empty = false
		return
	}
	p1 := e
	for {
		if p1.Next.End() {
			break
		}
		p1 = p1.Next
	}

	if p1.Next == nil {
		p1.Next = &GElement[T]{Data: item, Empty: false}
		p1.Next.Pre = p1
	} else {
		p1.Next.Data = item
		p1.Next.Empty = false
	}
}

// AppendUnique: first search, append at the end if not found
func (p *GList[T]) AppendUnique(item T) {
	e := p.element
	if e.Empty {
		e.Data = item
		e.Empty = false
		return
	}
	//search
	p1 := e

	for {
		if p1.Data == item {
			return
		}
		if p1.Next.End() {
			break
		}
		p1 = p1.Next
	}

	//end search

	if p1.Next == nil {
		p1.Next = &GElement[T]{Data: item, Empty: false}
		p1.Next.Pre = p1
	} else {
		p1.Next.Data = item
		p1.Next.Empty = false
	}
}

// test is this last element
func (p *GElement[T]) End() bool {
	if p == nil {
		return true
	}
	if p.Empty {
		return true
	}
	return false
}

// return element value
func (p *GElement[T]) Value() T {
	return p.Data
}

// insert item after current element
func (p *GElement[T]) Insert(item T) {
	e := &GElement[T]{Data: item, Empty: false}
	if p.End() {
		p.Next = e
		e.Pre = p
		return
	}
	if p.Next == nil {
		p.Next = e
		e.Pre = p
		return
	}
	p.Next.Pre = e
	e.Next = p.Next

	e.Pre = p
	p.Next = e
}

// if current element is the last one, will return nil
func (p *GElement[T]) NextElement() *GElement[T] {
	if p.Next.End() {
		return nil
	}
	return p.Next
}

// if current element is the last one, will return nil
func (p *GElement[T]) PreElement() *GElement[T] {
	return p.Pre
}
