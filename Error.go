package main

import "errors"

var ErrEmptySequence error = errors.New("sequence is empty")
var ErrMoreThanOneElement error = errors.New("sequence contains more than 1 element")
var ErrNoSatisfyingElement error = errors.New("sequence contains no element satisfying condition")
var ErrNotComparableType error = errors.New("not comparable type")
