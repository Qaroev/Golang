package main

import (
	"log"
	"net/http"
	"github.com/thedevsaddam/renderer"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "WEBCALC_GOLANG/*.html",
	}

	rnd = renderer.New(opts)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}

type Tag struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func home(w http.ResponseWriter, r *http.Request) {
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
		err = results.Scan(&tag.ID, &tag.Email, &tag.Username)
		if err != nil {
			panic(err.Error())
		}
		HomePageVars := Tag{ //store the date and time in a struct
			Email:    tag.Email,
			ID:       tag.ID,
			Username: tag.Username,
		}
		fmt.Println(tag.Email)
		fmt.Println(tag.ID)
		fmt.Println(tag.Username)
		rnd.HTML(w, http.StatusOK, "home", HomePageVars)
	}
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "about", nil)
}
