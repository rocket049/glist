package glist

import (
	"iter"
)

// generics element type of GList
type GElement[T comparable] struct {
	Data T
	Pre  *GElement[T]
	Next *GElement[T]
}

// generics list type
type GList[T comparable] struct {
	element *GElement[T]
	pool    *elementPool[T]
}

// use it like this: listNew := NewGList[TypeName]()
func NewGList[T comparable]() *GList[T] {
	return &GList[T]{pool: &elementPool[T]{}}
}

// get head of GList
func (p *GList[T]) Front() *GElement[T] {
	return p.element
}

// search one element of GList
func (p *GList[T]) SearchOne(item T) *GElement[T] {
	if p.element == nil {
		return nil
	}
	e := p.element
	for {
		if e.Data == item {
			return e
		}
		if e.Next == nil {
			break
		}
		e = e.Next
	}
	return nil
}

// search all elements of GList
func (p *GList[T]) SearchAll(item T) []*GElement[T] {
	if p.element == nil {
		return nil
	}
	ret := []*GElement[T]{}
	e := p.element
	for {
		if e.Data == item {
			ret = append(ret, e)
		}

		if e.Next == nil {
			break
		}
		e = e.Next
	}
	return ret
}

// Add: insert item at front, this is fastest
func (p *GList[T]) Add(item T) {
	if p.element == nil {
		p.element = p.pool.Get()
		p.element.Data = item
		return
	}
	e := p.pool.Get()
	e.Data = item
	e.Next = p.element
	p.element.Pre = e
	p.element = e
}

// AddUnique: insert item at front if it not exists.
func (p *GList[T]) AddUnique(item T) {
	if p.element == nil {
		p.element = p.pool.Get()
		p.element.Data = item
		return
	}

	//search
	p1 := p.element
	for {
		if p1.Data == item {
			return
		}
		if p1.Next == nil {
			break
		}
		p1 = p1.Next
	}
	//end search

	e := p.pool.Get()
	e.Data = item
	e.Next = p.element
	p.element.Pre = e
	p.element = e
}

// remove GLement e
func (p *GList[T]) Remove(e *GElement[T]) {

	defer p.pool.Put(e)

	if p.element == e {
		p.element = p.element.Next
		p.element.Pre = nil
		return
	}

	if e.Next != nil {
		e.Next.Pre = e.Pre
	}

	e.Pre.Next = e.Next

}

// remove all items of the list
func (p *GList[T]) Clear() {
	if p.element != nil {
		p.pool.PutList(p.element)
	}
	p.element = nil
}

// Append: append item at the end
func (p *GList[T]) Append(item T) {
	if p.element == nil {
		p.element = p.pool.Get()
		p.element.Data = item
		return
	}

	p1 := p.element

	for {
		if p1.Next == nil {
			break
		}
		p1 = p1.Next
	}

	p1.Next = p.pool.Get()
	p1.Next.Data = item
	p1.Next.Pre = p1

}

// AppendUnique: first search, append at the end if not found
func (p *GList[T]) AppendUnique(item T) {
	if p.element == nil {
		p.element = p.pool.Get()
		p.element.Data = item
		return
	}

	p1 := p.element

	//search

	for {
		if p1.Data == item {
			return
		}
		if p1.Next == nil {
			break
		}
		p1 = p1.Next
	}

	//end search

	p1.Next = p.pool.Get()
	p1.Next.Data = item
	p1.Next.Pre = p1
}

// insert item after element "pos"
func (p *GList[T]) Insert(item T, pos *GElement[T]) {
	e := p.pool.Get()
	e.Data = item

	if pos.Next == nil {
		pos.Next = e
		e.Pre = pos
		return
	}
	pos.Next.Pre = e
	e.Next = pos.Next

	e.Pre = pos
	pos.Next = e
}

// test is it empty
func (p *GList[T]) Empty() bool {
	if p.element == nil {
		return true
	}
	return false
}

// use it in for-range loop. usage:
//
//	for val := range list1.Range() {
//		// so some thing
//	}
func (p *GList[T]) Range() iter.Seq[T] {
	return func(yield func(T) bool) {
		e := p.element
		for {
			if e == nil {
				return
			}
			if !yield(e.Data) {
				return
			}
			e = e.Next
		}
	}
}

// return element value
func (p *GElement[T]) Value() T {
	return p.Data
}

// if current element is the last one, will return nil
func (p *GElement[T]) NextElement() *GElement[T] {
	return p.Next
}

// if current element is the last one, will return nil
func (p *GElement[T]) PreElement() *GElement[T] {
	return p.Pre
}
