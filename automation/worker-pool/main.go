package main

import "fmt"

func main() {
	var jobChannel = make(chan int, 5)

	jobs := make([]int, 0)
	workers := 5

	for i := 0; i < 20; i++ {
		jobs = append(jobs, i)
	}

	for w := 0; w < workers; w++ {
		go func() {
			work := <-jobChannel

			fmt.Println(work)
		}()
	}

	for _, item := range jobs {
		jobChannel <- item
	}

	close(jobChannel)

}

///private/var/folders/lm/t_7591qx77n71k1jb3329qdc0000gn/T/___go_build_TemplateApi_automation_worker_pool
//0
//3
//1
//2
//4
//fatal error: all goroutines are asleep - deadlock!
