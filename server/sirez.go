package main

import "fmt"

func main() {
	users := []string{"Tom", "Ali", "Kate"}

	/// add to srez

	users = append(users, "Bob")
	for _, value := range users {
		fmt.Println(value)
	}

	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:8]
	fmt.Println(users1)
	users2 := initialUsers[:4]
	fmt.Println(users2)

	// delate element

	idUsers := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	var n = 3
	idUsers = append(idUsers[:n], idUsers[n+1:]...)
	fmt.Println(idUsers)
}
