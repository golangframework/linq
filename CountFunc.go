package main

import "github.com/golangframework/system"

// Returns a number that represents how many elements in the specified sequence satisfy a condition.
//
// Parameters:
//
//	predicate system.Predicate[TSource] - A function to test each element for a condition.
//
// Returns:
//
//	A number that represents how many elements in the sequence satisfy the condition in the predicate function.
func (query query[TSource]) CountFunc(predicate system.Predicate[TSource]) int {
	count := 0
	if predicate != nil {
		next := query.iterate()
		for _, ok := next(); ok; _, ok = next() {
			count++
		}
	}
	return count
}
