package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	port = ":8080"
)

func index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Something went wrong parsing template: ", err)
	}

	err = t.Execute(w, nil)

}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Listening at port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
