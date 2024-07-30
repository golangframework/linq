package main

import (
	"fmt"

	Collections "github.com/golangframework/collections"
)

func main() {
	list := make(Collections.List[int], 0)
	list.GetEnumerator()
	ok := From(list).Any()
	fmt.Println(ok)
	list.Add(4)
	ok = From(list).Any()
	fmt.Println(ok)
	ok = From(list).AnyFunc(func(item int) bool {
		return item > 4
	})
	fmt.Println(ok)
	list1 := Collections.List[int]{
		1, 2, 3,
	}
	list2 := Collections.List[int]{
		4, 5, 6,
	}
	list3 := From(list1).Zip(From(list2).Where(func(item int) bool { return item > 4 }))
	fmt.Println(list3)
	dict := make(Collections.Dictionary[int, string])
	_ = dict
	dict2 := make(map[int]string)
	_ = dict2

	ok = From(dict).AnyFunc(func(item Collections.KeyValue[int, string]) bool { return item.Key > 4 })
	ok2 := From(Collections.Dictionary[int, string](dict2)).AnyFunc(func(item Collections.KeyValue[int, string]) bool { return item.Key > 4 })

	fmt.Println(ok)
	fmt.Println(ok2)

	ala := From(dict).FirstOrSpecified(Collections.KeyValue[int, string]{Key: 1, Value: "Jeden"})
	fmt.Println(ala)

	i := From(list2).FirstOrDefault()
	fmt.Println(i)

	list4 := Collections.List[int]{
		1, 1, 2, 2, 3, 4, 5, 6, 7, 7,
	}
	fmt.Println(list4)
	list5 := From(list4).Distinct().ToList()
	fmt.Println(list5)
	list6 := From(list4).Where(func(item int) bool { return item > 4 }).ToList()
	fmt.Println(list6)

	a := Zip(list4, list5)
	fmt.Println(a)

	list7 := From(list4).Select(func(value int) any { return value + 1 }).ToList()
	fmt.Println(list7)

	b := Zip2(list4, list5).Where(func(item Tuple[int, int]) bool { return item.Item1 > item.Item2 }).ToList()
	fmt.Println(b)

}
