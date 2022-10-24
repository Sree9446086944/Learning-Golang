package main

import (
	"fmt"
	"sync"
)

// go mod init channels

//goroutine doesnt know about other goroutines execution, it just tells to waitgroup once operation finishes
//channels - way/pipeline in which multiple goroutines can communicate each other

func main() {
	fmt.Println("Channels in golang.")

	//create channel
	// myCh := make(chan int) //define what kind of values will be passing b/w each other, any datatype allowed
	myCh := make(chan int, 2) //1 is default value of channel
	wg := &sync.WaitGroup{}

	// myCh <- 5              // push values into channel using <- , myCh like box and <- value adding into box
	// fmt.Println(<-myCh)    //recieve values from channel using <-

	//if run -> gives deadlock - blocking goroutine/error
	//channel only allow to accept values from other if somebody is listening to this channel

	wg.Add(2)
	//this goroutine recieve value from channel
	//<- made this goroutine Recieve ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh // trying to listen value from channel myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// we can check if the channel is open by isChannelOpen, and tell the value 0 is signal value or coming because of closed channel

		fmt.Println(<-myCh) //recieve value from channel, listen to only one value
		//if to listen multiple values from channel
		fmt.Println(<-myCh) // if no seperate fetching <-, then deadlock

		//or can use bufferedChannel
		wg.Done()
	}(myCh, wg) //(channel type,waitgroup)fn call immediately

	//this goroutine will add values to channel
	//adding <- make this goroutine Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5 //added 5 to channel
		//pass another value
		myCh <- 6
		myCh <- 0   // still shows 0 value, we dont know whether 0 value recieving is signal value or msg that recieve that u r listening to a closed channel
		close(myCh) //close channel
		wg.Done()
	}(myCh, wg)

	//now run app -> no error

	wg.Wait()
}
