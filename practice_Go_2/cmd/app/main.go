package main

import (
	"fmt"
)

func main() {
	users := []string{"Pedro", "Juan", "Luis", "Matias", "Alon", "Fernando"}
	users2 := []string{"Matias", "Fernando"}

	fmt.Println(SearchUsers(users, users2))
}

func SearchUsers(users []string, users2 []string) bool {
	for _, u2 := range users2 {
		for j, u1 := range users {
			if u1 == u2 {
				break
			} else {
				if j == len(users) - 1 {
					return false
				}
			}
		}
	}

	return true
}
