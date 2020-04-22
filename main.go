package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

const (
	port    = ":8080"
	baseAPI = "https://swapi.dev/api/"
)

// PageData defines the basic shape of page data
type PageData struct {
	Title  string
	Header string
	Data   Characters
}

// Character defines the character response from the api
type Character struct {
	Name string `json:"name"`
}

// Characters define initial response from api
type Characters struct {
	Count      int         `json:"count"`
	Characters []Character `json:"results"`
}

func getCharacters(w http.ResponseWriter, r *http.Request) {
	data := Characters{}
	pagedata := PageData{
		Title:  "Starwars Characters",
		Header: "Popular Characters",
		Data:   data,
	}

	resp, err := http.Get(baseAPI + "people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Something went wrong accessing the swapi api: ", err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &pagedata.Data)

	t, err := template.ParseFiles("templates/characters.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("An error occured parsing template file: ", err)
	}
	t.Execute(w, pagedata)
}

func index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Something went wrong parsing template: ", err)
	}

	data := PageData{
		Title:  "Go Web Server",
		Header: "Awesome",
	}

	err = t.Execute(w, data)

	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/characters", getCharacters)
	fmt.Println("Listening at port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
