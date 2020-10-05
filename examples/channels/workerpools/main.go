package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//WaitGroup
	//Used to wait for a collection of goroutines to finish executing
	//Control is blocked until all goroutines finish executing
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1) //wg's counter is incremented by the value passed to Add()
		go func(i int, wg *sync.WaitGroup) {
			fmt.Println("started goroutine ", i)
			time.Sleep(2 * time.Second)
			fmt.Println("Goroutine %d ended\n", i)
			wg.Done() //decrements the counter in wg
		}(i, &wg) //important to add address of wg.
		// if address of wg not passed, then each goroutine will have its own copy of waitgroup and main will not be notified when they finish
	}
	wg.Wait() //blocks the goroutine in which it's called until counter becomes zero
	fmt.Println("all go routines finished executing")

	//implementation of worker pool is important for buffered channels
	//Worker pool is a collection of threads which are waiting for tasks to be assigned to them.
	//Once they finish the task assigned, they make themselves available again for the next task
}
