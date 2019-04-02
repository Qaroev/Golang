package main

import (
	"database/sql"
	"golang.org/x/exp/errors/fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}
func main() {
	db, err := sql.Open("mysql", "root:@/nodejs")
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT id, email,username FROM user")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.ID, &tag.Email,&tag.Username)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(tag.Email)
		fmt.Println(tag.ID)
		fmt.Println(tag.Username)
	}

	fmt.Println(results)
	fmt.Println(*db)
	fmt.Println("connect")
	defer db.Close()
}
