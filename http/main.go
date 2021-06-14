package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.Id}} </h1>
<div> {{printf "User %d" .UserID}} </div>
<div> {{printf "%s (completed: %t)" .Title .Completed}} </div>`

func handler(w http.ResponseWriter, r *http.Request) {
	var url = "https://jsonplaceholder.typicode.com/"

	resp, err := http.Get(url + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	var todo Todo

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprint(resp.StatusCode), http.StatusBadGateway)
		return
	}

	json.NewDecoder(resp.Body).Decode(&todo)

	tem := template.New("mine")

	tem.Parse(form)
	tem.Execute(w, todo)
}

func main() {

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
