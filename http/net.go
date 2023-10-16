package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var tasks = []Task{
	{ID: 1, Title: "Task 1", Content: "Description for Task 1"},
	{ID: 2, Title: "Task 2", Content: "Description for Task 2"},
}

var mu sync.Mutex

func main() {
	http.HandleFunc("/tasks", TasksHandler)
	http.HandleFunc("/tasks/", TaskHandler)
	http.ListenAndServe(":8080", nil)
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(tasks)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		mu.Lock()
		defer mu.Unlock()
		id, err := strconv.Atoi(r.URL.Path[len("/tasks/"):])
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		for _, task := range tasks {
			if task.ID == id {
				json.NewEncoder(w).Encode(task)
				return
			}
		}
		http.Error(w, "Task not found", http.StatusNotFound)
	} else if r.Method == http.MethodPost {
		mu.Lock()
		defer mu.Unlock()
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		task.ID = len(tasks) + 1
		tasks = append(tasks, task)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Task created successfully")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
