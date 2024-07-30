package main

import Collections "github.com/golangframework/collections"

func (query query[TSource]) AggregateFunc(accumulator TSource, aggregate func(accumulator TSource, value TSource) TSource) (result TSource) {
	next := query.iterate()
	result = accumulator
	if aggregate != nil {
		for item, ok := next(); ok; item, ok = next() {
			result = aggregate(result, item)
		}
	}
	return result
}

func AggregateFunc[TSource any, TResult any](enumerable Collections.IEnumerable[TSource], accumulator TResult, aggregate func(accumulator TResult, value TSource) TResult) (result TResult) {
	enumerator := enumerable.GetEnumerator()
	result = accumulator
	if aggregate != nil {
		for item, ok := enumerator.GetNext(); ok; item, ok = enumerator.GetNext() {
			result = aggregate(result, item)
		}
	}
	return result
}
