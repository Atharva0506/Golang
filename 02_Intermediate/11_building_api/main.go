package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	Title      string  `json:"title"`
	Instructor string  `json:"instructor"`
	Price      float64 `json:"price"`
}

// Global "Database" to store courses
var courses []Course

// 1. GET Request Handler
func getCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// If the database is empty, let's load a default one!
	if len(courses) == 0 {
		courses = append(courses, Course{Title: "Golang Basics", Instructor: "Atharva", Price: 99.6})
	}

	json.NewEncoder(w).Encode(courses)
}

// 2. POST Request Handler (Reading incoming JSON from Client)
func createCourseHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON data from the Request Body!
	var newCourse Course
	err := json.NewDecoder(r.Body).Decode(&newCourse)

	if err != nil {
		// If they send bad JSON, send a 400 Bad Request error!
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Save it to our text database
	courses = append(courses, newCourse)

	// Return a 201 Created Status
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Course '%s' successfully created!", newCourse.Title)
}

// 3. Dynamic Path Variables ({id}) and Query Parameters (?key=value)
func getCourseByIDHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Path Variables: Use Go 1.22's r.PathValue to extract the {id} from the URL!
	id := r.PathValue("id")

	// 2. Query Parameters: Use r.URL.Query().Get() to read ?discount=true
	discount := r.URL.Query().Get("discount")

	w.Header().Set("Content-Type", "application/json")

	// Print a fake JSON string so we can see the variables at work!
	fmt.Fprintf(w, `{"message": "Fetching course ID: %s", "discount_requested": "%s"}`+"\n", id, discount)
}

// 4. DELETE Request using the Path Variable!
func deleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Course %s successfully deleted!\n", id)
}
func main() {
	// 3. Routing HTTP Methods (A massive new feature in Go 1.22+)
	// You just put the HTTP Method right before the path!
	http.HandleFunc("GET /courses", getCourseHandler)
	http.HandleFunc("POST /courses", createCourseHandler)

	// Handling Dynamic Paths with {id} syntax
	http.HandleFunc("GET /courses/{id}", getCourseByIDHandler)
	http.HandleFunc("DELETE /courses/{id}", deleteCourseHandler)

	fmt.Println("Server running on port 8080. Try:")
	fmt.Println("  [GET]    http://localhost:8080/courses")
	fmt.Println("  [POST]   http://localhost:8080/courses")
	fmt.Println("  [GET]    http://localhost:8080/courses/123?discount=true")
	fmt.Println("  [DELETE] http://localhost:8080/courses/123")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Crashed", err)
	}
}
