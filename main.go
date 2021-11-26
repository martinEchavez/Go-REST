package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Task Model
type task struct {
	ID          int    `json:ID`
	Name        string `json:Name`
	Description string `json:Description`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:          1,
		Name:        "Task One",
		Description: "Some Description",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)

	log.Fatal(http.ListenAndServe(":3000", router))
}
