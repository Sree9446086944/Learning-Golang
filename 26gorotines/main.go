package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signal []string

var wg sync.WaitGroup // usually this is pointer in real, we use regerence of wg

var mut sync.Mutex //usually pointer, since we need to pass on to multiple goroutines

// go mod init goroutines
// concurrency(task may be simultaneous) vs parallelism(multiple tasks done parallell)
func main() {
	//hello printed 6 times then world 6 times
	// greeter("Hello")
	// greeter("World")

	//creating goroutine by go keyword
	// go greeter("Hello") // fire up a thread that execute greeter("Hello"), but didnt tell when to come back, never waited to come back
	// greeter("World")
	//prints only "World"
	//main() exited as soon as all the methods fire up, we never waited that method to come back and print

	//method 1 - to prevent above issue
	//use time.sleep()

	// both thread starts, as soon as thread's job done, then come back and we recive it and prints

	//sleep.time() - not an ideal soln,
	//the main thread is not waiting for other thread to come back, as soon as main finish app stops even if other thread running

	//method 2 - using sync package
	// https://pkg.go.dev/sync

	// ----------------------------------------//
	// wait groups tutorial
	websitelist := []string{"https://lco.dev", "https://go.dev", "https://google.com", "https://fb.com", "https://github.com"}
	// for _, web := range websitelist {
	// 	getStatusCode(web)
	// }

	//took some time for each call, it will be fast if we fire each method call in different goroutine
	/*for _, web := range websitelist {
		go getStatusCode(web)
	}*/
	//this prints nothing since, the main method not waits for goroutines to report back, need to use time.sleep() or sync package

	//waitgroups
	//  need a variable that is waitgroup
	//  waitgroup is advanced version of time.sleep()
	/*while using waitgroup, there are 3 jobs to be done;
		   1- Add - adding goroutines counter,as soon as goroutine created, need to add that in to this waitgroup, so it will be responsible for management, to tell not end the method till this waitgroup end
	       2- Done - once the goroutine is done, we need to call Done()
		   3- Wait - optional wait for main thread to wait for goroutines to finish
	*/

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) //add 1 waitgroup in each iteration method call
	}
	wg.Wait() //Wait() always at end of main(), tell main to wait for other waitgroups to finish

	//this is much faster, since each call is done in different goroutines independently, not done one by one
	fmt.Println(signal)

	//mutex
	// goroutines are managed by go runtime , flexible stack -2kb
	// Thread is managed by os, fixed stack - 1mb

	//if goroutine managed by go runtime, who is managing lock on to memory?
	//there may be case when more than 1 goroutines simultaneously writes to single memory - not good
	//to avoid this there is mutux
	//mutex - provides lock over memory, i.e lock the memory if one goroutine is writing in a memory, until its done no other goroutine can us this memory
	// https://pkg.go.dev/sync#mutex

	// Read-Write mutex -  allows to read the memory, but if anybody comes to write that will lock that, whoever was reading it will be out till the write operation is done

	// create mutex variable - var mut sync.Mutex
	//whatever is the memory we are performing read/write operations, add mut.Lock() once done mut.Unlock()
	//while running , we cannot see any difference, but in big app, there might be difference
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	}
// }

// add time.sleep() to wait for execution
// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond) // each loop waits 3 millisecond
// 		fmt.Println(s)
// 	}
// }

//normal method
// func getStatusCode(endpoint string) {
// 	res, err := http.Get(endpoint)
// 	if err != nil {
// 		fmt.Println("error in endpoint")
// 	}
// 	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint) // %d for integer
// }

// func getStatusCode(endpoint string) {
// 	defer wg.Done() //pass signal done at end of this method, so that waitgroup knows this goroutine is done and contine other execution, waitgroup counter reduces
// 	res, err := http.Get(endpoint)
// 	if err != nil {
// 		fmt.Println("error in endpoint")
// 	} else {
// 		signal = append(signal, endpoint)
// 		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint) // %d for integer

// 	}
// }

// with mutex
func getStatusCode(endpoint string) {
	defer wg.Done() //pass signal done at end of this method, so that waitgroup knows this goroutine is done and contine other execution, waitgroup counter reduces
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("error in endpoint")
	} else {
		mut.Lock()
		signal = append(signal, endpoint) // read/write operation
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint) // %d for integer

	}
}
