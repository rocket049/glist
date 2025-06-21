# GList: golang generics list type (go语言泛型列表)

功能齐全的泛型列表

```
type GElement
    func (p *GElement[T]) End() bool
    func (p *GElement[T]) Insert(item T)
    func (p *GElement[T]) NextElement() *GElement[T]
    func (p *GElement[T]) PreElement() *GElement[T]
    func (p *GElement[T]) Value() T
type GList
    func NewGList[T comparable]() *GList[T]
    func (p *GList[T]) Add(item T)
    func (p *GList[T]) AddUnique(item T)
    func (p *GList[T]) Append(item T)
    func (p *GList[T]) AppendUnique(item T)
    func (p *GList[T]) Clear()
    func (p *GList[T]) Empty() bool
    func (p *GList[T]) Front() *GElement[T]
    func (p *GList[T]) Remove(e *GElement[T])
    func (p *GList[T]) SearchAll(item T) []*GElement[T]
    func (p *GList[T]) SearchOne(item T) *GElement[T]
```