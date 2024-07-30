package main

type ValueSelector[TSource any, TValue any] func(value TSource) TValue

// Returns distinct elements from a sequence according to a specified value selector function.
//
// Parameters:
//
//	valueSelector ValueSelector[TSource, any] - A function to extract the value for each element.
func (q query[TSource]) DistinctFunc(valueSelector ValueSelector[TSource, any]) IQueryable[TSource] {
	return query[TSource]{
		iterate: func() Iterator[TSource] {
			next := q.iterate()
			set := make(map[interface{}]struct{})
			return func() (item TSource, ok bool) {
				if valueSelector != nil {
					for item, ok = next(); ok; item, ok = next() {
						if _, has := set[valueSelector(item)]; !has {
							set[item] = struct{}{}
							return
						}
					}
				}
				return
			}
		},
	}
}
