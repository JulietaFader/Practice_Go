package main

import "fmt"

func main() {
	users := []string{"Pedro", "Juan", "Luis", "Matias", "Fernando","Alon"}
	users2 := []string{"Matias", "Fernando"}

	fmt.Println(SearchUsers(users, users2))
}

func SearchUsers(users []string, users2 []string) []string {

	var users3 []string

	for _, value := range users {
		for _, value1 := range users2 {
			if value == value1 {
				users3 = append(users3, value)
			}
		}
	}
	return users3
}
