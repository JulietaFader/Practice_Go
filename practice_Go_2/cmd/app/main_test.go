package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchUsersTrue(t *testing.T) {
	f_u1 := []string{"Pedro", "Juan", "Luis", "Matias", "Alon", "Fernando"}
	f_u2 := []string{"Matias", "Fernando"}

	result := SearchUsers(f_u1, f_u2)

	assert.True(t, result)
}

func TestSearchUsersFalse(t *testing.T) {
	f_u1 := []string{"Pedro", "Juan", "Luis", "Matias", "Alon", "Mariano"}
	f_u2 := []string{"Matias", "Fernando"}

	result := SearchUsers(f_u1, f_u2)

	assert.False(t, result)
}
