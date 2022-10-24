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

// go mod init github.com/Sree9446086944/buildapi
//go get -u github.com/gorilla/mux

// https://github.com/gorilla/mux

// model for course - this usually go into file
type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	// custome type, since Author is not created, we cannot reference(&) instead give as ponter type(*Author), since copy of object passed issue give pointer
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware / helper method - goes to file in real case
func (c *Course) IsEmpty() bool { // method part of struct, since Course struct is passed,get access to properties of Course struct
	// return c.CourseId == "" && c.CourseName == "" //return true if both empty
	return c.CourseName == "" //check only courseName since courseId we manually create
}

func main() {
	//controller - govern how we handle routes
	//seeding the data(means adding values initially to db for test) and routes that use the controlleres
	fmt.Println("Api")

	// mux give ability to have routers throught newRouter()
	r := mux.NewRouter() //create router

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Sree", Website: "lco.dev"}}) //Authod is a pointer type, so pass reference and that reference coming from Author
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Sree", Website: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET") //r- router, handleFunc(route,controller (not as fn passed instead reference, i.e without ())), and Method working on is Get
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET") //param "id"
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r)) // Fatal will stop if any error and logs error, use instead of if err!=null
	// :4000 means every url with 4000 port

	//go build .
	//go run main.go
	//localhost:4000/
	// http://localhost:4000/courses  - GET
	// http://localhost:4000/courses/2  - GET

	/*http://localhost:4000/course - POST
	{
		"coursename":"Java Spring",
		"price":199,
		"Author":{
		  "fullname":"Kunal",
		  "website":"communityclassroom.com"
		}
	  }*/

	/* http://localhost:4000/course/63 - PUT
	  {
		"coursename":"Nodejs",
		"price":199,
		"Author":{
		  "fullname":"Brad",
		  "website":"brad.com"
		}
	  }*/

	//   http://localhost:4000/course/2 - DELETE
}

//controllers - real case another file

// serve home route(/)
// governed by writer(where you write res for req) and reader(get the value from whoever sending the req- urls,params etc in r)
func serveHome(w http.ResponseWriter, r *http.Request) {
	//response
	w.Write([]byte("<h1>welcome to Api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses.")
	//setting header in w which is the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //json pkg have NewEncoder(writer), takes writer and Encode() writes to json whatever passed in w
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course.")
	w.Header().Set("Content-Type", "application/json")

	//grab id from req r
	params := mux.Vars(r) // grabs params from req

	//loop through courses and find matching id and return
	for _, course := range courses {
		if course.CourseId == params["id"] {
			//whatever given in Encode() will be converted to json
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with id.")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	//decode json value from req
	fmt.Println("Create one course.")
	w.Header().Set("Content-Type", "application/json")

	//what if body is empty, i.e if req body is empty
	if r.Body == nil { //direct access to body from request
		json.NewEncoder(w).Encode("Please send some data.")
	}
	//data like {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	//check only if title is duplicate
	//loop, title matches with course.coursename, JSON response - already exist

	//generate unique id, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())               //creates random value each time
	course.CourseId = strconv.Itoa(rand.Intn(100)) //convert int to string using strconv pkg Itoa() method

	courses = append(courses, course) // course added to courses
	json.NewEncoder(w).Encode(course) // returns added course
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	//by courseId, update values
	fmt.Println("Update one course.")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params := mux.Vars(r) // mux give access to Vars ,grab from req

	//loop through value, once match id, remove value, add value back in course with id
	for index, course := range courses {
		if course.CourseId == params["id"] { //id from route, path param
			courses = append(courses[:index], courses[index+1:]...) //removed the course from courses , veriadic operation so ... (varArgs - variable number of arguments)

			var course Course //new course struct
			//since course removed, again add course with id and new values in courses
			//grab data from req body(json)
			_ = json.NewDecoder(r.Body).Decode(&course) //decode req body and convert to course object/struct, &course reference since original course struct need to modify not copy
			course.CourseId = params["id"]
			courses = append(courses, course) // add new course with id
			json.NewEncoder(w).Encode(course) // send res as updated course
			return
		}
	}
	//send  response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course.")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, check by courseId, remove the reference of that course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //removed the course
			//send a confirm/deny response
			json.NewEncoder(w).Encode("Deleted the course")
			break //breaks the entire loop
		}
	}
}
