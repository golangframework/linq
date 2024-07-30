package main

import "github.com/golangframework/system"

// Determines whether a sequence contains any element that satisfies a condition.
//
// Parameters:
//
//	predicate system.Predicate[TSource] - A function to test each element for a condition.
//
// Returns:
//
//	bool - true if the source sequence contains element that passes the test in the specified predicate, false otherwise.
func (query query[TSource]) ContainsFunc(predicate system.Predicate[TSource]) bool {
	next := query.iterate()
	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			return true
		}
	}
	return false
}
