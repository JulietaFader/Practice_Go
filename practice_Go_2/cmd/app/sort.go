package main

import (
	"sort"
)

func SortArray(numbers []int) []int {
	array := make([]int, len(numbers))
	for i, number := range numbers {
		array[i] = number
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	return array
}
