package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"

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

// Get tasks All
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Get Tasks
func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Ivalid ID")
		return
	}

	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

// Create tasks
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// Delete Tasks
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Ivalid ID")
		return
	}

	for index, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:index], tasks[index + 1:]...)
			fmt.Fprintf(w, "The task with Id %v has been remove succesfully", taskId)
		}
	}
}

// Update Tasks
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	var updateTask task

	if err != nil {
		fmt.Fprintf(w, "Ivalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Date")
		return
	}

	json.Unmarshal(reqBody, &updateTask)

	for index, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:index], tasks[index + 1:]...)
			updateTask.ID = taskId
			tasks = append(tasks, updateTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskId)
		}
	}
}

// Home
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API 👌")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
