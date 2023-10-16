package main

import (
	"fmt"
	"time"
)

func confibonacci(n int) {
	ch := make(chan int)
	go func() {
		fibonacci(n, ch)
	}()
	for i := range ch {
		fmt.Print(" ", i)
	}
}

func fibonacci(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {

		ch <- x
		x, y = y, x+y
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)

}
