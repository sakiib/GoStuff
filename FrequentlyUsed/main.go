package main

import (
	"fmt"
	"sort"
)

type data struct {
	x int
	y int
}

func main() {
	fmt.Println("Frequently Used stuff!")
	// Slice sorting: sorting a slice of integers, need to import the sort library
	val := []int{3, 2, 1}
	sort.Ints(val)
	fmt.Println(val)

	// to sort an array []a, use it like a slice [l, r) -> r is exclusive, array indices 1 to n
	var a [100] int
	n := 10
	sort.Ints(a[1 : n+1]) // Simple array sorting

	var list []data // assume we have a list of data of the data format
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Scanf("%d %d\n", &x, &y)
		list = append(list, data{x: x, y: y})
	}
	// Struct sorting: SliceStable keep i < j stable, if val[i] is same as val[j]
	sort.SliceStable(list, func(i, j int) bool {
		if list[i].x == list[j].x {
			return list[i].y > list[j].y
		} else {
			return list[i].x < list[j].x
		}
	})
	// not stable sorting
	sort.Slice(list, func(i, j int) bool {
		if list[i].x == list[j].x {
			return list[i].y > list[j].y
		} else {
			return list[i].x < list[j].x
		}
	})
}