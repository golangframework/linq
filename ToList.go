package main

import Collections "github.com/golangframework/collections"

// Creates a List[TSource] from an IEnumerable[TSource].
//
// Returns:
//
//	List[TSource] - A List[TSource] that contains elements from the input sequence.
func (query query[TSource]) ToList() (list Collections.List[TSource]) {
	list = make(Collections.List[TSource], 0)
	next := query.iterate()
	for item, ok := next(); ok; item, ok = next() {
		list = append(list, item)
	}
	return list
}

// Creates a Queue[TSource] from an IEnumerable[TSource].
//
// Returns:
//
//	Queue[TSource] - A Queue[TSource] that contains elements from the input sequence.
func (query query[TSource]) ToQueue() (list Collections.Queue[TSource]) {
	list = make(Collections.Queue[TSource], 0)
	next := query.iterate()
	for item, ok := next(); ok; item, ok = next() {
		list = append(list, item)
	}
	return list
}

// Creates a Stack[TSource] from an IEnumerable[TSource].
//
// Returns:
//
//	Stack[TSource] - A Stack[TSource] that contains elements from the input sequence.
func (query query[TSource]) ToStack() (list Collections.Stack[TSource]) {
	list = make(Collections.Stack[TSource], 0)
	next := query.iterate()
	for item, ok := next(); ok; item, ok = next() {
		list = append(list, item)
	}
	return list
}

func (q query[TSource]) Select(valueSelector ValueSelector[TSource, any]) IQueryable[any] {
	return query[any]{
		iterate: func() Iterator[any] {
			next := q.iterate()
			return func() (item any, ok bool) {
				if valueSelector != nil {
					if element, ok := next(); ok {
						return valueSelector(element), true
					}
				}
				return
			}
		},
	}
}

func Select[TSource any, TValue any](enumerable Collections.IEnumerable[TSource], valueSelector ValueSelector[TSource, TValue]) IQueryable[TValue] {
	return query[TValue]{
		iterate: func() Iterator[TValue] {
			enumerator := enumerable.GetEnumerator()
			return func() (item TValue, ok bool) {
				if valueSelector != nil {
					if element, ok := enumerator.GetNext(); ok {
						return valueSelector(element), true
					}
				}
				return
			}
		},
	}
}

type KeySelector[TSource any, TKey comparable] ValueSelector[TSource, TKey]

func ToDictionary[TSource any, TKey comparable, TValue any](enumerable Collections.IEnumerable[TSource], keySelector KeySelector[TSource, TKey], valueSelector ValueSelector[TSource, TValue]) (dictionary Collections.Dictionary[TKey, TValue]) {
	dictionary = make(Collections.Dictionary[TKey, TValue], 0)
	enumerator := enumerable.GetEnumerator()
	for item, ok := enumerator.GetNext(); ok; item, ok = enumerator.GetNext() {
		dictionary[keySelector(item)] = valueSelector(item)
	}
	return dictionary
}
