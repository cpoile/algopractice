package linkedlist

import "errors"

var ErrEmptyList = errors.New("empty list")

// node and member functions
type node struct {
	Val        interface{}
	next, prev *node
}

func (n *node) Next() *node {
	return n.next
}

func (n *node) Prev() *node {
	return n.prev
}

// List and member functions
type List struct {
	first, last *node
}

func (ll *List) Reverse() {
	for n := ll.first; n != nil; n = n.prev {
		n.prev, n.next = n.next, n.prev
	}
	ll.first, ll.last = ll.last, ll.first
}

func (ll *List) First() *node {
	return ll.first
}

func (ll *List) Last() *node {
	return ll.last
}

func (ll *List) PushFront(elem interface{}) {
	n := &node{elem, ll.first, nil}
	ll.first = n
	if n.next != nil {
		n.next.prev = n
	}
	if ll.last == nil {
		ll.last = n
	}
}

func (ll *List) PopFront() (interface{}, error) {
	if ll.first == nil {
		return 0, ErrEmptyList
	}
	n := ll.first
	ll.first = ll.first.next
	if ll.first != nil {
		ll.first.prev = nil
	}
	if ll.last == n {
		ll.last = nil
	}
	return n.Val, nil
}

func (ll *List) PushBack(elem interface{}) {
	n := &node{elem, nil, ll.last}
	ll.last = n
	if n.prev != nil {
		n.prev.next = n
	}
	if ll.first == nil {
		ll.first = n
	}
}

func (ll *List) PopBack() (interface{}, error) {
	if ll.last == nil {
		return 0, ErrEmptyList
	}
	n := ll.last
	ll.last = ll.last.prev
	if ll.last != nil {
		ll.last.next = nil
	}
	if ll.first == n {
		ll.first = nil
	}
	return n.Val, nil
}

func NewList(in ...interface{}) *List {
	ll := &List{}
	for _, v := range in {
		ll.PushBack(v)
	}
	return ll
}
