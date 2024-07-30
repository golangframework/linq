package main

import Collections "github.com/golangframework/collections"

func (query query[TSource]) Zip(enumerable Collections.IEnumerable[TSource]) Collections.List[Tuple[TSource, TSource]] {
	tuples := make(Collections.List[Tuple[TSource, TSource]], 0)
	next := query.iterate()
	enumerator := enumerable.GetEnumerator()
	item2, ok2 := enumerator.GetNext()
	for item1, ok1 := next(); ok1 && ok2; item1, ok1 = next() {
		tuples = append(tuples, Tuple[TSource, TSource]{
			Item1: item1,
			Item2: item2,
		})
		item2, ok2 = enumerator.GetNext()
	}
	return tuples
}

func (q query[TSource]) Zip2(enumerable Collections.IEnumerable[TSource]) IQueryable[any] {
	return query[any]{
		iterate: func() Iterator[any] {
			next := q.iterate()
			enumerator := enumerable.GetEnumerator()
			return func() (item any, ok bool) {
				item2, ok2 := enumerator.GetNext()
				for item1, ok1 := next(); ok1 && ok2; item1, ok1 = next() {
					item, ok = Tuple[TSource, TSource]{
						Item1: item1,
						Item2: item2,
					}, true
				}
				return
			}
		},
	}
}

// func (q query[TSource]) Zip2(enumerable Collections.IEnumerable[TSource]) IQueryable[Tuple[TSource, TSource]] {
// 	return query[Tuple[TSource,TSource]]{
// 		iterate: func() Iterator[Tuple[TSource,TSource]] {
// 			next := q.iterate()
// 			enumerator := enumerable.GetEnumerator()
// 			return func() (item Tuple[TSource,TSource], ok bool) {
// 					item2, ok2 := enumerator.GetNext()
// 					for item1, ok1 := next(); ok1 && ok2; item1, ok1 = next() {
// 						item, ok = Tuple[TSource, TSource]{
// 							Item1: item1,
// 							Item2: item2,
// 						}, true
// 					}
// 				return
// 			}
// 		},
// 	}
// }

// func (q query[TSource]) Zip2(enumerable Collections.IEnumerable[TSource]) query[Tuple[TSource, TSource]] {
// 	tuples := make(Collections.List[Tuple[TSource, TSource]], 0)
// 	next := q.iterate()
// 	enumerator := enumerable.GetEnumerator()
// 	enumerator.Reset()
// 	item2, ok2 := enumerator.GetNext()
// 	for item1, ok1 := next(); ok1 && ok2; item1, ok1 = next() {
// 		tuples = append(tuples, Tuple[TSource, TSource]{
// 			Item1: item1,
// 			Item2: item2,
// 		})
// 		item2, ok2 = enumerator.GetNext()
// 	}
// 	return tuples
// }

func Zip[TSource1 any, TSource2 any](enumerable1 Collections.IEnumerable[TSource1], enumerable2 Collections.IEnumerable[TSource2]) Collections.List[Tuple[TSource1, TSource2]] {
	tuples := make(Collections.List[Tuple[TSource1, TSource2]], 0)
	enumerator1 := enumerable1.GetEnumerator()
	enumerator2 := enumerable2.GetEnumerator()
	item2, ok2 := enumerator2.GetNext()
	for item1, ok1 := enumerator1.GetNext(); ok1 && ok2; item1, ok1 = enumerator1.GetNext() {
		tuples = append(tuples, Tuple[TSource1, TSource2]{
			Item1: item1,
			Item2: item2,
		})
		item2, ok2 = enumerator2.GetNext()
	}
	return tuples
}

func Zip2[TSource1 any, TSource2 any](enumerable1 Collections.IEnumerable[TSource1], enumerable2 Collections.IEnumerable[TSource2]) IQueryable[Tuple[TSource1, TSource2]] {
	return query[Tuple[TSource1, TSource2]]{
		iterate: func() Iterator[Tuple[TSource1, TSource2]] {
			enumerator1 := enumerable1.GetEnumerator()
			enumerator2 := enumerable2.GetEnumerator()
			return func() (item Tuple[TSource1, TSource2], ok bool) {
				item1, ok1 := enumerator1.GetNext()
				item2, ok2 := enumerator2.GetNext()
				ok = ok1 && ok2
				if ok {
					item = Tuple[TSource1, TSource2]{
						Item1: item1,
						Item2: item2,
					}
				}
				return
			}
		},
	}
}
