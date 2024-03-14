package main

import "log"

// Function to check if number is multiple of 3 and send true in isMul channel
func ismultipleOf3(ch1 chan int, isMul chan bool) {
	for {
		select {
		case x := <-ch1:
			if x%3 == 0 && x%15 != 0 {
				isMul <- true
			} else {
				isMul <- false
			}
		}
	}

}

// Function to check if number is multiple of 5 and send true in isMul channel
func ismultipleOf5(ch2 chan int, isMul chan bool) {
	for {
		select {
		case x := <-ch2:
			if x%5 == 0 && x%15 != 0 {
				isMul <- true
			} else {
				isMul <- false
			}
		}
	}

}

// Function to check if number is multiple of 3 and send true in isMul channel
func ismultipleOf15(ch3 chan int, isMul chan bool) {
	for {
		select {
		case x := <-ch3:
			if x%15 == 0 {
				isMul <- true
			} else {
				isMul <- false
			}
		}
	}

}

func main() {
	// Three channels for read Data to check if number is multiple of 3,5 or 3 and 5
	var ch1 chan int = make(chan int)
	var ch2 chan int = make(chan int)
	var ch3 chan int = make(chan int)

	// Channel to return the boolean value if divisible by certain number
	var isMul chan bool = make(chan bool)

	// Running all goroutines and waiting for values from channel
	go ismultipleOf3(ch1, isMul)
	go ismultipleOf5(ch2, isMul)
	go ismultipleOf15(ch3, isMul)

	// Running for loop and provide values for channel
	for i := 1; i <= 30; i++ {
		// Add To channel 1 and check if number is multiple of 3
		ch1 <- i
		// If isMul channel has true it means , we have to print fizz
		if <-isMul == true {
			log.Println("fizz")
			continue
		}
		// Add To channel 2 and check if number is multiple of 3
		ch2 <- i
		// If isMul channel has true it means , we have to print buzz
		if <-isMul == true {
			log.Println("buzz")
			continue
		}
		// Add To channel 3 and check if number is multiple of 15
		ch3 <- i
		// If isMul channel has true it means , we have to print fizzBuzz
		if <-isMul == true {
			log.Println("fizzBuzz")
			continue
		}
		// Print the number if not divisible by any
		log.Println(i)
	}
}
