// Original Program: How can we reduce the execution time of this program using concurrency?
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var c chan bool = make(chan bool)
var wg sync.WaitGroup

func task(i int) {
	// Task function simulates some work by printing a message and then sleeping for 100 milliseconds
	fmt.Println("Task", i)
	time.Sleep(100 * time.Millisecond)
	c <- true
	wg.Done()
}

func main() {
	start := time.Now()

	// Loop through 30 tasks sequentially
	wg.Add(30)
	for i := 1; i <= 30; i++ {
		go task(i)
	}

	for i := 1; i <= 30; i++ {
		<-c
	}

	elapsed := time.Since(start)
	log.Printf("Time taken %s", elapsed)
}
