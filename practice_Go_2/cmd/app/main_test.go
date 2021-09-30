package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchNumbersTrue(t *testing.T) {
	f_u1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f_u2 := []int{1, 2, 3}

	result := search(f_u1, f_u2)

	assert.True(t, result)
}

func TestSearchNumbersFalse(t *testing.T) {
	f_u1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	f_u2 := []int{}

	result := search(f_u1, f_u2)

	assert.False(t, result)
}
func TestSearchNumbersEmpty(t *testing.T) {
	f_u1 := []int{}
	f_u2 := []int{}

	result := search(f_u1, f_u2)

	assert.False(t, result)

}
