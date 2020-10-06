package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Guarantee of Delivery
	//				Unbuffered				Buffered
	//	Delivery	*Guaranteed*		*Not Guaranteed*

	//State
	//	The behavior of a channel is directly influenced by its current State.
	//	State of a channel can be: *nil, open or closed*
	//	State determines how the send and receive operations behave
	//	Signals are sent and received through a channel. Don't say read/write: channels aren't I/O

	//nil channel - nil state when it is declared to its zero value
	var nilChannel chan string
	nilChannel = nil //explicitly setting to nil
	//When a channel is in a nil state, any send/receive attempted will block.

	//open channel - when it's made using built-in function make
	openChannel := make(chan string)
	//When channel is in open state, signals can be sent and received

	//closed channel - closed state when using close()
	close(nilChannel)
	close(openChannel)
	//When placed in closed state, signals can no longer be sent, but still possible to receive signals

	//				nil			open		closed
	//	Send		Blocked		Allowed		Panic!
	//	Receive		Blocked		Allowed		Allowed

	//With and Without Data
	//You signal WITH data by performing a send on a channel
	sendChannel := make(chan string)
	sendChannel <- "hello"
	//When you signal with data, usually because
	//	A goroutine is being asked to start a new task
	//	A goroutine reports back a result

	//You signal WITHOUT data by closing a channel
	close(sendChannel)
	//When you signal without data, it's usually because (major use cases, not all)
	//	A goroutine is being told to stop what they are doing
	//	A goroutine reports back that they are done with no result or return
	//	A goroutine reports that it has completed processing and shut down

	//A single goroutine can signal many goroutines at once
	//Signaling with data is always a 1:1 exchange between goroutines

	//Signaling With Data
	//Three channel options: Unbuffered, Buffered > 1, or Buffered = 1

	//					Guarantee		No Guarantee		Delayed Guarantee
	//	Channel			Unbuffered		Buffered > 1		Buffered = 1

	//Guarantee
	//		Unbuffered gives you a Guarantee that a signal being sent has been received
	//			Because the Receive of the signal *happens before* the *send* of the signal completes
	//No Guarantee
	//		Buffered channel of size > 1 gives you *No Guarantee* that a signal being sent has been received
	//			Because Send of the signal Happens Before the Receive of the singal completes
	//Delayed Guarantee
	//		Buffered channel of size = 1 gives you a Delayed Guarantee. It can guarantee that the previous signal that was sent has been received
	//			Because the receive of the first signal, happens before the send of the second signal completes

	//Signaling Without Data
	//	Mainly reserved for cancellation
	//		Allows one goroutine to signal another to cancel what they are doing and move on
	//	Cancellation can be implemented using both Unbuffered and Buffered channels
	//		But using a buffered channel when no data will be sent is a code smell

	//				First Choice		Second Choice	Smell
	//	Channel		context.Context		Unbuffered		Buffered

	//close() is used to signal without data.
	//		Can still receive signals on a channel that is closed
	//			Any receive on a closed channel will not block and receive operation always returns

	//Most cases "context" standard library package is best to implement signaling without data
	//	context uses an Unbuffered channel underneath for the signaling and close() to signal without data
	//	If you choose to use your own channel for cancellation, channel should be of type *chan struct {}*

	//Signal With Data - Guarantee - Unbuffered channels
	//Wait for Task
	waitTaskChannel := make(chan string) //unbuffered channel with string data as the signal
	go func() {                          //worker is created
		p := <-waitTaskChannel //the channel receive, causing worker to block while waiting for a signal

		// does work here
		fmt.Println(p)

		// done and free to go
	}()

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	waitTaskChannel <- "start"
	//the worker has no idea how long it's going to take for a signal to be sent to start

	//Wait for Result
	waitResultChannel := make(chan string)
	go func() { //worker is created
		time.Sleep(2 * time.Second) //some work is done

		waitResultChannel <- "signal" //the channel receive

		//worker is done and free to go
	}()

	p := <-waitResultChannel //the channel send. Since unbuffered the receive happens before the send
	fmt.Println(p)           //guaranteed that you have received the result
	//in this example, you have no idea how long it is going to take the worker to finish the task and send the result

	//cost/benefit
	//	Unbuffered channel provides a guarantee that a signal being sent was received
	//	Cost of this guarantee is unknown latency
	//	Both scenarios, unknown latency is something you have to live with because guarantee is required

	//Signal With Data - No Guarantee - Buffered Channels > 1
	//When you dont need to know that a signal being sent has been received, two options: Fan Out, and Drop

	//How to decide how much capacity a channel needs? Three questions:
	//	Do I have a well defined amount of work to be completed? How much work is there?
	//	If my worker can't keep up, can I discard any new work? How much outstanding work meets capacity?
	//	What level of risk is acceptable if the program terminates unexpectedly?
	//		Anything in the buffer will be lost

	//Fan Out Buffered Channel > 1
	//Allows you to allocate a well defined number of workers to solve the problem concurrently.
	//1 worker for every task, you know exactly how many reports you will receive
	//Benefit of workers not needing to wait to submit their results
	//	They do however need to each take a turn placing the result in the channel if they finish at/near the same time

	workers := 20
	fanChannel := make(chan string, workers)
	for e := 0; e < workers; e++ { //workers are created and given their tasks
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) //perform some work
			fanChannel <- "signal"                                       //send their results
		}()
	}

	for workers > 0 { //waiting for all 20 employees to finish and send their reports
		p := <-fanChannel //blocked in a channel receive waiting for reports
		fmt.Println(p)    //results printed
		workers--         //decrement counter
	}

	//Drop Pattern Buffered Channel > 1
	//Allows you to throw work away when your workers are at capacity
	//Benefit of continuing to accept work from clients and never applying back pressure or latency
	//The key is knowing when workers are at capacity so you don't over/under commit
	//	Integration testing or metrics is what is needed to help identify the right number
	const dropCapacity = 5
	selectDropChannel := make(chan string, dropCapacity)

	go func() { //a single worker is created to handle the work
		for p := range selectDropChannel { //used for the channel receive
			fmt.Println("worker : received:", p) //every time a signal is received it is processed
		}
	}()

	const work = 20
	for w := 0; w < work; w++ { // attempt to send 20 signals to the worker
		select {
		case selectDropChannel <- "signal": //
			fmt.Println("manager : send ack")
		default: //if the send is going to block because there is no room in buffer, it is abandoned
			fmt.Println("manager : drop")
		}
	}

	close(selectDropChannel) //signal without data to the worker they are done and free to go once completed

	//Cost/Benefit
	//No guarantee that a signal being sent is ever received
	//Benefit of walking away from this guarantee is reduced/no latency in communication between two goroutines
	//Fan Out - there is a buffer space for each worker that will be sending a report
	//Drop - buffer is measured for capacity and if capacity is reached work is dropped so things can keep moving
	//In both options, have to live with lack of guarantee because reduction in latency is more important

	//Signal With Data - Delayed Guarantee - Buffered Channel = 1
	//Wait for Tasks
	//	Giving worker more than one task, feeding them tasks one after the other
	//		The worker must finish each task before they can start a new one
	//		There can be latency issues between the handoff of work
	//	Buffered channel of 1 has benefit. If everything is running at expected pace, neither sender and app have to wait for the other
	//If any time a send is blocked because channel is at capacity, you know the worker is having a problem and stop
	//	If buffer is empty and you perform the send, you have the guarantee that worker has taken your previous signal
	//	If perform the send and can't, you have the guarantee that the worker hasn't taken the last signal
	waitForTasksChannel := make(chan string, 1)
	go func() { //worker is created
		for p := range waitForTasksChannel { //used for channel receive
			fmt.Println("worker : working :", p) //work is processed as it is received
		}
	}()

	const bufferedWork = 10
	for x := 0; x < bufferedWork; x++ { //if your worker can run as fast as you can send, latency is reduced
		waitForTasksChannel <- "signal"
	}

	close(waitForTasksChannel) //last piece of work submitted will be received (flushed) before the range is terminated

	//Signal Without Data - Context
	//When you are on a deadline and if the single worker doesn't finish in time, you are not willing to wait

	duration := 50 * time.Millisecond //how long the worker has to finish task
	//creates a context goroutine that will close unbuffered channel once duration is met
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel() //you are responsible for calling cancel function regardless of how things turn out
	//Can call the cancel() more than once, this is deferred until this function terminates

	contextChannel := make(chan string, 1)
	go func() { //worker created
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) //some work is done
		contextChannel <- "signal"                                   //worker sends results to channel
	}()

	select { //wait for worker to send result
	case p := <-contextChannel: //receive result from worker - worker completed on time
		fmt.Println("work complete", p)
	case <-ctx.Done(): //context goroutine closes - worker didn't finish before context closed
		fmt.Println("moving on")
	}

	//Conclusion
	//Use Channels to orchestrate and coordinate goroutines
	//	Focus on signaling attributes and not the sharing of data
	//	Signaling with data or without data
	//	Synchronizing access to shared state can be done easier in other ways
	//Unbuffered Channels
	//	Receive happens before the send
	//	Benefit: 100% guarantee the signal has been received
	//	Cost: Unknown latency on when the signal will be received
	//Buffered Channels
	//	Send happens before the receive
	//	Benefit: Reduce blocking latency between signaling
	//	Cost: no guarantee when the signal has been received
	//		The larger the buffer, the lower the guarantee
	//		Buffer of 1 can give you one delayed send of guarantee
	//Closing Channels:
	//	Close happens before the received(like buffered)
	//	Signaling without data
	//	Perfect for Signaling cancellations and deadlines
	//nil channels
	//	Send and Receive block
	//	Turn off signaling
	//	For rate limiting or short term stoppages

	//If any given Send on a channel CAN cause the sending goroutine to block
	//	Do not use a buffered channel larger than 1
	//		Buffers larger than 1 must have reason/measurements
	//	Must know what happens when the sending goroutine blocks
	//If any given send on a channel WON'T cause the sending goroutine to block
	//	You have the exact number of buffers for each send? - Fan Out Pattern
	//	You have the buffer measured for max capacity? - Drop Pattern
	//Less is more with buffers
	//	Don't think about performance when thinking about buffers
	//	Buffers can help to reduce blocking latency between signaling
	//		Reducing blocking latency towards zero does not necessarily mean better throughput
	//		If a buffer of one is giving you good enough throughput then keep it
	//		Question buffers that are larger than one and measure for size
	//		Find the smallest buffer possible that provides good enough throughput
}
