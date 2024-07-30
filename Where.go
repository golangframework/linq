package main

import "github.com/golangframework/system"

// Filters a sequence of values based on a predicate.
//
// Parameters:
//
// predicate - system.Predicate[TSource] - A function to test each element for a condition.
//
// Returns:
//	IEnumerable[TSource] - An IEnumerable[TSource] that contains elements from the input sequence that satisfy the condition.
func (q query[TSource]) Where(predicate system.Predicate[TSource]) IQueryable[TSource] {
	return query[TSource]{
		iterate: func() Iterator[TSource] {
			next := q.iterate()
			return func() (item TSource, ok bool) {
				if predicate != nil {
					for item, ok = next(); ok; item, ok = next() {
						if predicate(item) {
							return
						}
					}
				}
				return
			}
		},
	}
}
