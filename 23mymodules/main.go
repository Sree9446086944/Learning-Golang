package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	//can run directly - go run main.go

	//correct way
	// go mod init <host(github)>
	//eg: go mod init github.com/Sree9446086944/mymodules

	// go versions
	//  go follows sem ver versioning (semantic versioning) - starts from 1.0.0
	// https://go.dev/ref/mod

	//go get
	//go to web and bring assets from web, repo given after get will be fetch directly

	//routing system - gorilla/mux
	// https://github.com/gorilla/mux
	//install - go get -u github.com/gorilla/mux
	//you can see this module dependency in go.mod file
	//from go.mod ->  require github.com/gorilla/mux v1.8.0 // indirect  -> indirect means its not using now, when we use this module, indirect comment will be gone

	// go.sum file also created
	//sum file will check for malicious attacks
	// since pulling from github repo, it checks on that, file shows the gitrepo with hash ,
	// any change in that repo in this special version, the hash will fail
	//verifying sum is done by go mod

	// to see the files mentioned in go.sum
	// go env -> GOPATH -> go to the path -> pkg -> mod -> cache
	//cache folder is where all the package and libraries from internet placed and not in working directory
	//next time we fetch this file, it wont fetch from web, it is fetched from local copy
	// cache -> download -> github.com -> gorilla -> mux etc to see these files

	greeter()

	//from mux github pade doc
	r := mux.NewRouter() //  new router r
	//this router can handleit on route(using HandleFunc()), define method inside method(just pass reference, not call)
	//dont need to pass serveHome(), its auto suggested, instead pass reference serveHome without ()
	r.HandleFunc("/", serveHome).Methods("GET") //i.e only serving Get req, if route to \, serveHome() is called

	//Running a server
	// err := http.ListenAndServe(":4000", r) //port and router , expect error, catch
	//or
	log.Fatal(http.ListenAndServe(":4000", r)) //if somthing wrong, Fatal() will log the value, log package

	// go build . -> build everything inside this directory
	// go run main.go
	// now can see in localhost:4000/

	//eventhough we are using gorilla/mux, it has indirect comment in go.mod,
	// use go mod tidy -. tidys all libraries dependent and removes package not using
	//go mod is expensive operations

	//go mod verify -> check go.mod and check how many modules required, also go to go.sum file and take those hashes which are unique
	//and try to verify its ok

	//go list ->  shows modules which app dependent on
	//go list all -> shows all packages installed

	//go list -m all -> shows modules dependent and also out base package (github.com/Sree9446086944/mymodules)

	//go list -m -versions github.com/gorilla/mux
	// - shows all versions available for mux - github.com/gorilla/mux v1.2.0 v1.3.0 v1.4.0 v1.5.0 v1.6.0 v1.6.1 v1.6.2 v1.7.0 v1.7.1 v1.7.2 v1.7.3 v1.7.4 v1.8.0
	//if latest version didnt work need to downgrade version

	//go mod why github.com/gorilla/mux -> shows module dependent in mux

	// go mod graph -> pull up graph of dependency , o/p - github.com/Sree9446086944/mymodules github.com/gorilla/mux@v1.8.0 -> our pkg dependent on mux

	//go mod edit -go 1.16  -> to edit go.mod file and here it changes the version of go in mod file to 1.16
	//go mod edit -module <name> -> changes module name in mod file

	//go mod vender -> creates vender folder similar to a node modules, now modules from web in cache folder, if case where we need from vender folder
	//router may not be doing all things, may be need to overwrite code then publsh to production, so use go mod vender
	// go run -mod=vender main.go -> first look into vender folder, its going to bring everything from this folder

}

func greeter() {
	fmt.Println("Hey there mod users.")
}

// created for gorilla/mux
func serveHome(w http.ResponseWriter, r *http.Request) {
	//we need to send req - to use urls,params etc - that all in 'r'
	// to send some response back - that is w

	//write response - sample
	//everything from web is []byte
	w.Write([]byte("<h1>Welcome to golang</h1>"))

}
