package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// NOTE: Models

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float64 `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	AuthorName    string `json:"name"`
	AuthorWebsite string `json:"website"`
}

// Fake DB
var courses []Course

// middleware <--> helper function
func (c *Course) IsValidCourse() bool {
	return c.CourseName != ""
}

func main() {
	fmt.Println("Building API in Golang !!")
	router := mux.NewRouter()

	// seeding ( adding random data )
	courses = append(courses, Course{"123", "Golang Backend", 199, &Author{"Hitesh Chaudhary Sir", "hitesh@loc.dev"}})
	courses = append(courses, Course{"456", "Microservices", 299, &Author{"Hitesh Chaudhary Sir", "hitesh@loc.dev"}})

	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getCourseByID).Methods("GET")
	router.HandleFunc("/course/", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":3000", router))
}

// NOTE: Controllers

// Home Route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to APIs</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all the courses")

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(courses): This line is responsible for encoding the courses data into JSON format and sending it as the HTTP response. Here's what it does:

	// json.NewEncoder(w): This creates a new JSON encoder that writes data to the w (http.ResponseWriter) output stream. The w parameter is an interface provided by the Go HTTP package for writing HTTP responses.

	// .Encode(courses): This method takes the courses variable (presumably a slice or some data structure containing course information) and encodes it into JSON format. The resulting JSON data is then written to the HTTP response stream.
	json.NewEncoder(w).Encode(courses)
}

func getCourseByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting a single course")
	w.Header().Set("Content-Type", "application/json")

	// grabbing id from the r
	params := mux.Vars(r)

	// loop through courses, find matching id and return the responses
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the given id.")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a course")
	w.Header().Add("Content-Type", "application/json")

	// in case the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Data cannot be empty")
		return
	}

	// in case of post request like in form of {}
	var course Course
	// destructuring
	_ = json.NewDecoder(r.Body).Decode(&course)
	if !course.IsValidCourse() {
		json.NewEncoder(w).Encode("Data body cannot be empty")
		return
	}

	course_Name := course.CourseName

	// in case the title of the course already matches
	for _, course := range courses {
		if course.CourseName == course_Name {
			json.NewEncoder(w).Encode("Course already present")
			fmt.Println("Course already present")
			return
		}
	}

	// NOTE:
	// Generate a unique ID --> Convert it to string --> append the new course to the database/ slice (here)
	// Seed the random number generator with the current time.
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(100)
	courseIdStr := strconv.Itoa(randomInt)
	course.CourseId = courseIdStr

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating a course")
	w.Header().Set("Content-Type", "application/json")

	// getting the courseId to be updated
	params := mux.Vars(r)

	// loop through all the ids, find that id and remove it
	// add the updated value

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			// grabbing the course data from the req body
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a course")
	w.Header().Set("Content-Type", "application/json")

	// fetching the id from the req
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The ID is successful deleted.")
			break
		}
	}
}
