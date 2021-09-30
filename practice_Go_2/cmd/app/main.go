package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 22, 23, 24, 25, 26, 27, 6, 7, 9, 10, 11, 12, 13, 20, 21, 28, 29, 14, 15, 16, 17, 18, 19, 30}
	numbers2 := []int{5, 3}
	SearchNumbersBinary(numbers, numbers2)
	fmt.Println(numbers, numbers2)
}

func search(numbers []int, numbers2 []int) bool {
	//return SearchNumbersSequential(numbers, numbers2)
	return SearchNumbersBinary(numbers, numbers2)
}

func SearchNumbersSequential(numbers []int, numbers2 []int) bool {

	if len(numbers) <= 0 || len(numbers2) <= 0 {
		return false
	}

	for _, u2 := range numbers2 {
		for j, u1 := range numbers {
			if u1 == u2 {
				break
			} else if j == len(numbers)-1 {
				return false
			}
		}
	}

	return true
}

func SearchNumbersBinary(numbers []int, numbers2 []int) bool {
	if len(numbers) <= 0 || len(numbers2) <= 0 {
		return false
	}

	numbers = SortArray(numbers)

	for _, u2 := range numbers2 {
		if !BinarySearch(numbers, u2) {
			return false
		}
	}

	return true
}
