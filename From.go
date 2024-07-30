package main

import Collections "github.com/golangframework/collections"

func From[TSource any](source Collections.IEnumerable[TSource]) IQueryable[TSource] {
	enumerator := source.GetEnumerator()
	return &query[TSource]{
		iterate: func() Iterator[TSource] {
			return enumerator.GetNext
		},
	}
}
