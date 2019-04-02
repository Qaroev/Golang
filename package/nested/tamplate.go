package main

import (
	"net/http"
	"log"
	"time"
	"html/template"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	HomePageVars := PageVariables{
		now.Format("02-01-2006"),
		now.Format("15:04:05"),
	}
	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		log.Print("Template parsing error", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error", err)
	}

}
