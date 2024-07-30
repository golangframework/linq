package main

// Returns distinct elements from a sequence.
func (q query[TSource]) Distinct() IQueryable[TSource] {
	return query[TSource]{
		iterate: func() Iterator[TSource] {
			next := q.iterate()
			set := make(map[interface{}]struct{})
			return func() (item TSource, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if _, has := set[item]; !has {
						set[item] = struct{}{}
						return
					}
				}
				return
			}
		},
	}
}
