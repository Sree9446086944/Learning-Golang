package main

import (
	"fmt"
	"os"
	"strings"
)

// passing configuration into go app - using json file, commandline args, env
// use env for securing app, seet password without hardcoding
func main() {
	fmt.Println("OUTPUT: ", os.Getenv("OUTPUT"))
	// OUTPUT=/Users/sreek/output.txt go run main.go -> passing env and value at runtime, check in git bash

	// to see go envs -> go env

	fmt.Println("GO PATH: ", os.Getenv("GOPATH")) //already in system env, load already available in system path
	//go run main.go
	// o/p -> GO PATH:  C:\Users\sreek\go

	//get all env in system
	for _, env := range os.Environ() {
		fmt.Println(env) //prints all env as key=value
	}

	for _, env := range os.Environ() {
		fmt.Println(strings.Split(env, "=")) //prints all env as slice of [key value]
	}
}
