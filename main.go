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
	ok = From(list).AnyWith(func(item int) bool {
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
}
