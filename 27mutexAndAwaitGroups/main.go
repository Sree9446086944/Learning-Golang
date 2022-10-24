package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")
	// A race condition in Go occurs when two or more goroutines have shared data and interact with it simultaneously.

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{} // pass in write operation, sometimes read operations also
	var score = []int{0} //default value 0

	//creating seperate goroutines
	//ifis / immediately invoked fns - like lambda/anonymous fn
	//no name and body and run that immediately using () at end
	//pass waitgroup pointer type since reference
	// wg.Add(1) //or
	wg.Add(3)                                    //give total number at goroutines at start or Add(1) at each fn
	go func(wg *sync.WaitGroup, m *sync.Mutex) { // making this fn goroutine by go keyword, using this goroutine to add values in score
		fmt.Println("One R")
		mut.Lock() // lock and unlock in write operation
		score = append(score, 1)
		mut.Unlock()
		wg.Done() //each goroutine done after execution, counter reduces
	}(wg, mut) // fn call, pass waitgroup wg, mut lock resource not copy is passed
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//all the 3 goroutine is trying to write score slice

	wg.Wait() // waits for all other goroutines to finish before main goroutine finish
	fmt.Println(score)
	//o/p  // no guarantee which goroutine will finish first, so value added order cannot be guaranteed
	/*	Three R
		One R
		Two R
		[0 3 1 2]*/

	//in real case we need to check cloud resources utilizing in better way
	//go run --race .  -> checks our program have race condition or not

	//race created will access memory
	//to avoid problems - use mutex -> lock and unlock in write operation
	//again check for race condition -> go run --race . -> now no issue shown -> great performance

	//can use mut.RWMutex(), can add RLock and RUnlock on read operation also
	//if use Mutex(not RWMutex) dont need RLock and RUnlock. Read is quick operation compared to write

	// i.e, if during write, some other thread come to write same resource its a problem, but for read its fine
}
