package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	numbers := []int{0, 1, 5, 23, 24, 51, 600, 700, 890, 999}
	result := BinarySearch(numbers, 24)
	assert.True(t, result)
}

func TestSearchInMinBound(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := BinarySearch(numbers, 0)
	assert.True(t, result)

}

func TestSearchInMaxBound(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := BinarySearch(numbers, 9)
	assert.True(t, result)
}

func TestSearchNotFound(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := BinarySearch(numbers, 21)
	assert.False(t, result)
}

func TestSearchNotFoundNegative(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := BinarySearch(numbers, -21)
	assert.False(t, result)

}
func TestSearchFail(t *testing.T) {

}
