package main

import (
	"fmt"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func main() {
	//user := User{"Baby", 12}
	//user1 := user
	//user1.Name = "Val"
	//inc(&user)
	//fmt.Println(user)
	//fmt.Println(GetUser().Balance)
	//str := ` {"name":"Ali", "balance":12} `
	//user := User{}
	//json.Unmarshal([]byte(str), &user)
	//fmt.Printf("%v\n", user)
	user1 := User{"ab", "Ali", 12}
	user2 := User{"ba", "Ali", 12}

	m := map[string]User{user1.Id: user1, user2.Id: user2}

	u := m["ab"]
	fmt.Println(u)
}

func inc(user *User) {
	user.Balance++
	fmt.Println(user.Balance)
}

//func GetUser() *User {
//	return &User{"Baby", 12}
//}
