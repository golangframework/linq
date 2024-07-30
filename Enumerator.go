package main

import Collections "github.com/golangframework/collections"

func (query query[TSource]) GetEnumerator() Collections.IEnumerator[TSource] {
	return &enumerator[TSource]{
		object: query,
	}
}

type enumerator[T any] struct {
	object query[T]
}

func (enumerator *enumerator[T]) Reset() {
}

func (enumerator *enumerator[T]) GetNext() (object T, ok bool) {
	return enumerator.object.iterate()()
}
