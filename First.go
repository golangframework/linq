package main

// Returns the first element of a sequence.
//
// Returns:
//
//	item - The first element in the sequence.
//	err - The error if the sequence is empty, nil otherwise.
func (query query[TSource]) First() (item TSource, err error) {
	next := query.iterate()
	item, ok := next()
	if !ok {
		err = ErrEmptySequence
	}
	return item, err
}
