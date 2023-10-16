package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	for j := 0; j < 4; j++ {
		wg.Add(2)
		go func(ch <-chan int) {
			//i := <-ch

			for {
				if i, ok := <-ch; ok {
					fmt.Println(" ", i)
				} else {
					break
				}
			}
			//ch <- 21
			defer wg.Done()
		}(ch)
		go func(ch chan<- int) {
			//k:=<-ch
			ch <- 42
			ch <- 32
			//ch <- 22
			//ch <- 21

			defer wg.Done()
		}(ch)

	}
	close(ch)
	wg.Wait()
	// confibonacci(10)
	// test(
	// 	[]string{
	// 		"hi friend",
	// 		"What's going on?",
	// 		"Welcome to the business",
	// 		"I'll pay you to be my friend",
	// 	},
	// 	[]string{
	// 		"Will you make your appointment?",
	// 		"Let's be friends",
	// 		"What are you doing?",
	// 		"I can't believe you've done this.",
	// 	},
	// )
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func main() {
// 	channel()
// }

// func channel() {
// 	ch := make(chan int, 50)
// 	for j := 0; j < 4; j++ {
// 		wg.Add(2)
// 		go func(ch <-chan int) {
// 			for {
// 				i, ok := <-ch
// 				if !ok {
// 					fmt.Println("Channel Closed")
// 					wg.Done() // Decrement the counter when the goroutine completes
// 					return
// 				}
// 				fmt.Println("Received:", i)
// 			}
// 		}(ch)
// 		go func(ch chan<- int) {
// 			ch <- 42
// 			ch <- 32
// 			wg.Done() // Decrement the counter when the goroutine completes
// 			close(ch) // Close the channel after sending values
// 		}(ch)
// 	}
// 	wg.Wait() // Wait for all goroutines to complete
// }
