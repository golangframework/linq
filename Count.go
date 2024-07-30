package main

// Returns the number of elements in a sequence.
//
// Returns:
//
//	The number of elements in the input sequence.
func (query query[TSource]) Count() int {
	next := query.iterate()
	count := 0
	for _, ok := next(); ok; _, ok = next() {
		count++
	}
	return count
}
