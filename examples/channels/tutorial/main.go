package main

import "fmt"

//Without specifying buffered, a channel defaults to unbuffered

func main() {
	//zero value of a channel is nil.
	var a chan int
	//nil channels must be defined with make like maps and slices
	if a == nil {
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

	//Sending and receiving from channel
	data := <-a //read from channel a
	a <- data   //write to channel a

	//Send and receives are blocking by default
	//When data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel
	//When data is read from a channel, the read is blocked until some Goroutine writes data to that channel
	//Allows effective communication without the use of explicit locks or conditional variables

	//Bidirectional channels by default - data can be both sent and received on them
	//It is possible to create unidirectional channels - can only send OR receive data
	sendch := make(chan<- int) //chan <- int denotes a send only channel as arrow is pointing to chan
	fmt.Println(sendch)

	//It is possible to convert a bidirectional channel to a send or a receive only, but not vice versa

	//senders have the ability to close the channel to notify receivers that no more data will be sent on the channel
	//receivers can use an additional variable whether the channel has been closed
	//v, ok := <- ch
	//ok == true: value was received by a successful send operation to a channel
	//ok == false: reading from a closed channel. Value read from closed channel will be zero value of channel's type

	ch := make(chan int)
	go func(chnl chan int) { //writes 0 to 9 to the channel and then closes it
		for i := 0; i < 10; i++ {
			chnl <- i
		}
		close(chnl)
	}(ch)
	for { //infinite for loop which checks whether the channel is closed using ok. If closed then loop is broken
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received", v, ok)
	}

	//The for range form of loop can be used to received values from a channel until it is closed
	rangeChannel := make(chan int)
	go func(chnl chan int) {
		for i := 0; i < 10; i++ {
			chnl <- i
		}
		close(chnl) //Once chnl is closed, the loop automatically exits
	}(rangeChannel)
	for v := range rangeChannel {
		fmt.Println("Received ", v)
	}
}
