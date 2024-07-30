package main

import "github.com/golangframework/system"

// Determines whether any element of a sequence satisfies a condition.
//
// Parameters:
//
//	predicate system.Predicate[TSource] - A function to test each element for a condition.
//
// Returns:
//
//	bool - true if the source sequence is not empty and at least one of its elements passes the test in the specified predicate, false otherwise.
func (query query[TSource]) AnyFunc(predicate system.Predicate[TSource]) bool {
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
