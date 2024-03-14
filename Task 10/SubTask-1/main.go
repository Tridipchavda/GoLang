package main

import (
	"log"
	"time"
)

// Task 1
// Made channel for sending messages between a and b
var ch chan int = make(chan int)

// Go routine to give input to channel first as 0 and print 1 after recieving data from channel
func a() {
	for {
		ch <- 0
		log.Println(<-ch)
	}
}

// Go routine to print 0 and send 1 to channel
func b() {
	for {
		log.Println(<-ch)
		ch <- 1
	}
}

func main() {
	// Run both goroutines and wait for two seconds for them to execute
	go a()
	go b()
	time.Sleep(time.Second * 2)

}
