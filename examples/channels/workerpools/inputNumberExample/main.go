package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Example: worker pool to find the sum of digits of the input number
//If 234 is passed, the output would be 9 (2 + 3 + 4)
//The input to the worker pool will be the list of pseudo random integers
//Core functionalities
//Creation of pool of goroutines which listen on an input buffered channel waiting for jobs
//Addition of jobs to the input buffered channel
//Writing results to an output buffered channel after job creation
//Read and print results from the output buffered channel

var (
	jobs    = make(chan Job, 10)
	results = make(chan Result, 10)
)

type Job struct { //worker goroutines listen for new tasks on the jobs buffered channel
	id       int
	randomNo int
}

type Result struct { //once task is done, result written to results buffered channel
	job         Job
	sumOfDigits int
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs { //waits for and reads from jobs channel
		output := Result{job, digits(job.randomNo)} //creates result from current job
		results <- output                           //writes to buffered channel
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)      //increments wg before creating goroutine
		go worker(&wg) //creates worker goroutine and passes address of wg
	}
	wg.Wait()      //waits for all goroutines to finish
	close(results) //when finished, closes results channel since no one else will be writing to results channel
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999) //generates pseudo random numbers with maximum of 998
		job := Job{i, randomNo}    //creates a job with loop counter i and random number
		jobs <- job                //then writes to the job channel
	}
	close(jobs) // closes jobs channel after writing all jobs
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf(
			"Job id %d, input random number %d, sum of digits %d\n",
			result.job.id,
			result.job.randomNo,
			result.sumOfDigits) //prints job id, number input, and sum output
	}
	done <- true //writes to it once has printed all results
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs) //add jobs to the jobs channel
	done := make(chan bool)
	go result(done) //done channel passed to result so it can print outputs and notify once everything is printed
	noOfWorkers := 20
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

//With 10 workers, time taken ~20 seconds (20.021324796)

//With 20 workers, time taken ~10 seconds (10.016307882)
//But used ~15 extra MB of ram (~237 vs ~255)
