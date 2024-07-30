package main

import (
	Collections "github.com/golangframework/collections"
	"github.com/golangframework/system"
)

type Iterator[TSource any] func() (item TSource, ok bool)

type IQueryable[TSource any] interface {
	Collections.IEnumerable[TSource]
	Aggregate(aggregate func(accumulator TSource, value TSource) TSource) (result TSource)
	AggregateFunc(accumulator TSource, aggregate func(accumulator TSource, value TSource) TSource) (result TSource)

	Any() bool
	AnyFunc(predicate system.Predicate[TSource]) bool
	Contains(value TSource) bool
	ContainsFunc(predicate system.Predicate[TSource]) bool
	Count() int
	CountFunc(predicate system.Predicate[TSource]) int
	Distinct() IQueryable[TSource]
	DistinctFunc(valueSelector ValueSelector[TSource, any]) IQueryable[TSource]
	First() (item TSource, err error)
	FirstFunc(predicate system.Predicate[TSource]) (item TSource, err error)
	FirstFuncOrDefault(predicate system.Predicate[TSource]) (item TSource)
	FirstFuncOrSpecified(predicate system.Predicate[TSource], value TSource) (item TSource)
	FirstOrDefault() (item TSource)
	FirstOrSpecified(value TSource) (item TSource)
	Select(valueSelector ValueSelector[TSource, any]) IQueryable[any]
	// Filters a sequence of values based on a predicate.
	//
	// Parameters:
	//
	// predicate - system.Predicate[TSource] - A function to test each element for a condition.
	//
	// Returns:
	//	IEnumerable[TSource] - An IEnumerable[TSource] that contains elements from the input sequence that satisfy the condition.
	Where(predicate system.Predicate[TSource]) IQueryable[TSource]
	ToQueue() Collections.Queue[TSource]
	ToStack() Collections.Stack[TSource]
	ToList() Collections.List[TSource]
	Zip(enumerable Collections.IEnumerable[TSource]) Collections.List[Tuple[TSource, TSource]]
}

type query[TSource any] struct {
	iterate func() Iterator[TSource]
}

type Tuple[T1 any, T2 any] struct {
	Item1 T1
	Item2 T2
}
