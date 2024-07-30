package main

import "github.com/golangframework/system"

// Returns the first element in a sequence that satisfies a specified condition.
//
// Parameters:
//
//	predicate system.Predicate[TSource] - A function to test each element for a condition.
//
// Returns:
//
//	item - The first element in the sequence that satisfies a specified condition.
//	err - The error if the sequence does not contain an element that satisfies a specified condition or is empty, nil otherwise.
func (query query[TSource]) FirstFunc(predicate system.Predicate[TSource]) (item TSource, err error) {
	next := query.iterate()
	item, ok := next()
	if !ok {
		return item, ErrEmptySequence
	}
	if predicate != nil {
		for ; ok; item, ok = next() {
			if predicate(item) {
				return item, nil
			}
		}
	}
	return item, ErrNoSatisfyingElement
}
