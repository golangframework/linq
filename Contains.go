package main

import "reflect"

// Determines whether a sequence contains a specified element.
//
// Returns:
//
//	bool - true if the source sequence contains an element that has the specified value, false otherwise.
func (query query[TSource]) Contains(value TSource) bool {
	next := query.iterate()
	for item, ok := next(); ok; item, ok = next() {
		if reflect.DeepEqual(item, value) {
			return true
		}
	}
	return false
}
