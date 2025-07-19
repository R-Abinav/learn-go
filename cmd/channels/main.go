package main

//Channels 
//Channels enable go routines to pass around information
//Main features of channel are -> 1. They hold data, 2. They are thread safe (We avoid data races when we're reading and writing from memory), 
//3.We can listen when data is added or removed from a channel and we can block code exec until one of these events happens
/*
A data race occurs when multiple goroutines access the same memory location concurrently, and at least one of those accesses is a write operation, without proper synchronization. This can lead to unpredictable behavior and corrupted data.
Go channels are thread-safe because they handle all the necessary synchronization internally. 
When you send or receive values through a channel, the Go runtime ensures that these operations are properly coordinated between goroutines,
preventing data races. The channel implementation uses mutexes and other synchronization primitives under the hood to guarantee that concurrent access 
to the channel's internal state is safe. 
*/

// --- CODE SNIPPET 1 ---
//defer keyword -> The defer keyword in Go schedules a function call to be executed when the surrounding function returns, regardless of how it returns (normal return, panic, etc.).
//The code snippet here -> Unbuffered channel 
// Unbuffered Channel -> 1. Cannot store any values, 2. Send operations block until another goroutine receives, and vice versa, 3. Synchronous Comms -> sender and receiver must be ready at the same time

// import (
// 	"fmt"
// )

// func main(){
// 	var c = make(chan int)
// 	go process(c)
// 	for i := range c{
// 		fmt.Println(i)
// 	}
// }

// func process(c chan int){
// 	defer close(c)
// 	for i:=0; i<5; i++{
// 		c <- i
// 	}
// }

// --- CODE SNIPPET 2 ---
//Buffer channels -> 1.Can hold multiple values up to its capacity, 2.Sends only block when buffer is full, 3.Sender can send without immediate receiver (up to buffer limit)
//The previous program, keeps the process function running till main does its job, here process function does its job and exits!!


import (
	"fmt"
	"time"
)

func main(){
	var c = make(chan int, 5)
	go process(c)
	for i := range c{
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int){
	defer close(c)
	for i:=0; i<5; i++{
		c <- i
	}
	fmt.Println("Exiting function process")
}