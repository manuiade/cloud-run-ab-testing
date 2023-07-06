package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Data struct {
	Message string
	Color   string
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Print("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	magic_word, err := os.LookupEnv("FOOBAR")
	if err != false {
		m := Data{
			Message: "Hello " + magic_word,
			Color:   "green",
		}
		t.Execute(w, &m)
	} else {
		m := Data{
			Message: "Hello",
			Color:   "green",
		}
		t.Execute(w, &m)
	}
}
