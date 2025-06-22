package glist

import (
	"sort"
	"strings"
	"testing"
)

func TestGList(t *testing.T) {
	list1 := NewGList[string]()
	if !list1.Empty() {
		t.Fatal("fail GList.Empty()")
	}
	list1.Add("1hello")
	list1.AddUnique("1hello")
	list1.Add("2pig")

	list1.AppendUnique("3cow")
	list1.AppendUnique("2pig")
	list1.AddUnique("1hello")
	list1.AddUnique("2pig")
	p := list1.Front()
	res1 := []string{}
	for {
		if p == nil {
			break
		}
		res1 = append(res1, p.Value())
		p = p.Next
	}
	sort.Strings(res1)
	result1 := strings.Join(res1, ",")
	if "1hello,2pig,3cow" != result1 {
		t.Fatal("fail GList.Append/GList.AppendUnique(e)/GList.Add(e)/GList.AddUnique(e)")
	}

	list1.Clear()
	p = list1.Front()
	for {
		if p == nil {
			break
		}
		println(p.Value())
		p = p.Next
	}
	if !list1.Empty() {
		t.Fatal("fail Clear()")
	}
	list1.AddUnique("hello1")
	list1.AddUnique("hello1")
	list1.AddUnique("cow")
	list1.Add("mouse")
	list1.Add("dog2")
	list1.Add("pig2")
	list1.AddUnique("cow")
	list1.AddUnique("mouse")

	p = list1.SearchOne("dog2")
	if p == nil {
		t.Fatal("GList.Search(e) return nil")
	}
	if p.Value() != "dog2" {
		t.Fatal("GList.Search(e) return error result")
	}

	p = list1.Front()
	for {
		e := p.Next
		if p.Value() == "dog2" {
			list1.Remove(p)
		} else if p.Value() == "hello1" {
			list1.Remove(p)
		} else if p.Value() == "pig2" {
			list1.Remove(p)
		}

		if e == nil {
			break
		}

		p = e
	}
	res2 := []string{}
	p = list1.Front()
	for {
		if p == nil {
			break
		}

		res2 = append(res2, p.Value())
		p = p.Next
	}
	sort.Strings(res2)
	result2 := strings.Join(res2, ",")
	if result2 != "cow,mouse" {
		t.Fatal("faile GList.Remove(e): cow,mouse != result:", result2)
	}

	list1.Add("cow")
	list1.Add("cow")
	res3 := list1.SearchAll("cow")
	if len(res3) != 3 {
		t.Fatal("fail GList.SearchAll(item)")
	}

}

func TestGElement(t *testing.T) {
	list1 := NewGList[string]()
	e := list1.Front()
	if e != nil {
		t.Fatal("error GList.Front() of empty list")
	}

	list1.Add("first")
	e = list1.Front()
	list1.Insert("two", e)
	list1.Insert("three", e)

	e = e.NextElement()
	if e.Value() != "three" {
		t.Fatal("fail GElement.NextElement()")
	}
	e = e.PreElement()
	if e.Value() != "first" {
		t.Fatal("fail GElement.NextElement()")
	}
	e = e.PreElement()
	if e != nil {
		t.Fatal("fail GElement.PreElement() at head not nil")
	}

	e = list1.Front().NextElement().NextElement().NextElement()
	if e != nil {
		t.Fatal("fail GElement.NextElement() at last not nil")
	}

	list1.Clear()
	e = list1.Front()
	if e != nil {
		t.Fatal("error GList.Front() of empty list")
	}
}

func TestRange(t *testing.T) {
	list1 := NewGList[string]()
	list1.Add("a")
	list1.Add("b")
	list1.Add("c")
	res := []string{}
	for val := range list1.Range() {
		res = append(res, val)
	}
	sort.Strings(res)
	if strings.Join(res, "") != "abc" {
		t.Fatal("fail TestRange.")
	}
}
