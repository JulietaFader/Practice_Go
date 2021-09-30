package main

func BinarySearch(numbers []int, search int) bool {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		ihalf := (left + right) / 2
		ehalf := numbers[ihalf]

		if ehalf == search {
			return true
		}

		if search < ehalf {
			right = ihalf - 1
		} else {
			left = ihalf + 1
		}
	}

	return false
}

func BinarySearchString(users []string, user string) (i int) {
	return -1
}