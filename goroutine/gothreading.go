package main

import (
	//"fmt"
	"sync"
	//"time"
)

var bd = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter = 0

func sst() {
	// msg := "main"
	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	m.RLock()
	// 	go func() {
	// 		fmt.Println("counter:", counter)
	// 		m.RUnlock()
	// 		wg.Done()
	// 	}()
	// 	//wg.Add(1)
	// 	m.Lock()
	// 	go func() {
	// 		//fmt.Println("message1")
	// 		counter++
	// 		m.Unlock()
	// 		wg.Done()

	// 	}()
	// 	// go func(msg string){
	// 	// 	fmt.Println("hello there!!")
	// 	// 	fmt.Println(msg)
	// 	// }(msg)
	// 	//msg="not main"
	// 	//time.Sleep(100 * time.Millisecond)
	// 	//runtime.GOMAXPROCS(100)//will give 100 threads
	// }
	//channel()
	bd.Wait()
}
