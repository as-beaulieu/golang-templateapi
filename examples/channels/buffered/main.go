package main

import (
	"fmt"
	"time"
)

func main() {
	//sends and receives to an unbuffered channel are blocking

	//It is possible to create a channel with a buffer.
	//Sends to a buffer channel are blocked only when the buffer is full
	//Receives from a buffered channel are blocked only when the buffer is empty

	//buffered channels can be created by passing an additional capacity parameter to the make function
	//ch := make(chan type, capacity)
	//capacity should be greater than 0 for a channel to have a buffer, 0 is the default

	ch := make(chan string, 2) //Capacity of 2, can write 2 strings in channel without being blocked
	ch <- "hello"
	ch <- "there"
	fmt.Println(<-ch)
	fmt.Println(<-ch) //can also read twice from the channel before blocking

	cha := make(chan int, 2)
	go func(ch chan int) {
		for i := 0; i < 5; i++ { //will write 0 and 1 and then block until read
			ch <- i
			fmt.Println("successfully wrote", i, "to channel")
		}
		close(ch)
	}(cha)
	time.Sleep(2 * time.Second)
	for v := range cha {
		fmt.Println("read value", v, "from channel") //after reading 0 and 1, will unblock to write more
		time.Sleep(2 * time.Second)
	}

	//Deadlock - when a buffered channel is blocked and no concurrency to unblock it
	//Will cause a panic at run time - "fatal error: all goroutines are asleep - deadlock!"

	//Length vs capacity of buffered channels
	//Length - number of elements currently queued in it
	//Capacity - number of values the channel can hold - specified when creating with make()
}
