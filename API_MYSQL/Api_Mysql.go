package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

func routers() *chi.Mux {
	router.Get("/posts", AllPosts)
	//router.Get("/posts/{id}", DetailPost)
	router.Post("/posts", CreatePost)
	router.Put("/posts/{id}", UpdatePost)
	router.Delete("/posts/{id}", DeletePost)

	return router
}
func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	dbSource := "root:@/nodejs"

	var err error
	db, err = sql.Open("mysql", dbSource)

	catch(err)
}

type Post struct {
	ID       int    `json: "id"`
	Email    string `json: "email"`
	Username string `json: "Username"`
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	results, err := db.Query("SELECT id, email,username FROM user")
	fmt.Println(results)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var tag Post
		err = results.Scan(&tag.ID, &tag.Email, &tag.Username)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(tag.Email)
		fmt.Println(tag.ID)
		fmt.Println(tag.Username)
	}
	fmt.Println(results)
	defer results.Close()
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})

}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Insert user SET email =?,username =?")
	catch(err)
	fmt.Println(query)
	res, er := query.Exec(post.Email, post.Username)
	fmt.Println(res)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Update user set email =?, username=? where id=?")
	catch(err)
	_, er := query.Exec(post.Email, post.Username, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

}

// DeletePost remove a spesific post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("delete from user where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Logger return log message
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}
