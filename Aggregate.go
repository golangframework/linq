package main

import Collections "github.com/golangframework/collections"

func (query query[TSource]) Aggregate(aggregate func(accumulator TSource, value TSource) TSource) (result TSource) {
	next := query.iterate()
	if aggregate != nil {
		for item, ok := next(); ok; item, ok = next() {
			result = aggregate(result, item)
		}
	}
	return result
}

func Aggregate[TSource any, TAccumulator any, TResult any](enumerable Collections.IEnumerable[TSource], aggregate func(accumulator TAccumulator, value TSource) TAccumulator) (result TAccumulator) {
	enumerator := enumerable.GetEnumerator()
	if aggregate != nil {
		for item, ok := enumerator.GetNext(); ok; item, ok = enumerator.GetNext() {
			result = aggregate(result, item)
		}
	}
	return result
}
