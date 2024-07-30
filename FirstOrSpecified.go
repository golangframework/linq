package main

// Returns the first element of a sequence, or a specified value if no element is found.
//
// Returns:
//
//	item - The first element in the sequence if sequence is not empty, specified value otherwise.
func (query query[TSource]) FirstOrSpecified(value TSource) (item TSource) {
	if item, ok := query.iterate()(); ok {
		return item
	}
	return value
}
