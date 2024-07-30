package main

// Determines whether any element of a sequence exists or satisfies a condition.
//
// Returns:
//
//	bool - true if the source sequence contains any elements, false otherwise.
func (query query[TSource]) Any() bool {
	next := query.iterate()
	_, ok := next()
	return ok
}
