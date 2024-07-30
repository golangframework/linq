package main

import "github.com/golangframework/system"

// Returns the first element of the sequence that satisfies a condition or a default value if no such element is found.
//
// Parameters:
//
//	predicate system.Predicate[TSource] - A function to test each element for a condition.
//
// Returns:
//
//	item - The first element in the sequence that satisfies a specified condition, default value if there is no item that satisies a specified contition.
func (query query[TSource]) FirstFuncOrDefault(predicate system.Predicate[TSource]) (item TSource) {
	if predicate != nil {
		next := query.iterate()
		for item, ok := next(); ok; item, ok = next() {
			if predicate(item) {
				return item
			}
		}
	}
	return item
}
