package main

import Collections "github.com/golangframework/collections"

type Iterator[TSource any] func() (item TSource, ok bool)

type Query[TSource any] struct {
	iterate func() Iterator[TSource]
}

func (query Query[TSource]) GetEnumerator() Collections.IEnumerator[TSource] {
	return &enumerator[TSource]{
		object: query,
	}
}

type enumerator[T any] struct {
	object Query[T]
}

func (enumerator *enumerator[T]) Reset() {
}

func (enumerator *enumerator[T]) GetNext() (object T, ok bool) {
	return enumerator.object.iterate()()
}

func From[TSource any](source Collections.IEnumerable[TSource]) *Query[TSource] {
	enumerator := source.GetEnumerator()
	enumerator.Reset()
	return &Query[TSource]{
		iterate: func() Iterator[TSource] {
			return func() (item TSource, ok bool) {
				item, ok = enumerator.GetNext()
				return item, ok
			}
		},
	}
}

func (query Query[TSource]) Any() bool {
	next := query.iterate()
	_, ok := next()
	return ok
}

type Predicate[T any] func(item T) bool

func (query Query[TSource]) AnyWith(predicate Predicate[TSource]) bool {
	next := query.iterate()
	if predicate != nil {
		for item, ok := next(); ok; item, ok = next() {
			if predicate(item) {
				return true
			}
		}
	}
	return false
}

func (query Query[TSource]) Where(predicate Predicate[TSource]) Query[TSource] {
	return Query[TSource]{
		iterate: func() Iterator[TSource] {
			next := query.iterate()
			return func() (item TSource, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(item) {
						return
					}
				}
				return
			}
		},
	}
}

type Tuple[T1 any, T2 any] struct {
	Item1 T1
	Item2 T2
}

func (query Query[TSource]) Zip(enumerable Collections.IEnumerable[TSource]) Collections.List[Tuple[TSource, TSource]] {
	tuples := make(Collections.List[Tuple[TSource, TSource]], 0)
	next := query.iterate()
	enumerator := enumerable.GetEnumerator()
	enumerator.Reset()
	item2, ok2 := enumerator.GetNext()
	for item1, ok1 := next(); ok1 && ok2; item1, ok1 = next() {
		tuples = append(tuples, Tuple[TSource, TSource]{
			Item1: item1,
			Item2: item2,
		})
		item2, ok2 = enumerator.GetNext()
	}
	return tuples
}
