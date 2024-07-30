package main

// Returns the first element of a sequence, or a default value if no element is found.
//
// Returns:
//
//	item - The first element in the sequence if sequence is not empty, default value otherwise.
func (query query[TSource]) FirstOrDefault() (item TSource) {
	item, _ = query.iterate()()
	return item
}
