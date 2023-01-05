package main

import (
	"fmt"
	"strconv"
	"time"
)

// Recibes a int, as id, and 2 channels
// jobs is a only recibe channel
// results is a only send channel
func worker(id int, jobs <-chan int, results chan<- int, msg string) {

	fmt.Println(msg)
	// the for operator will end when the channel is closed.
	// jobs recibe 5 values (ints) and it it exits the loop, because the channel jobs is closed.
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// jobs channels
	const numJobs = 5

	// create buffered chan jobs, with 5 spaces
	jobs := make(chan int, numJobs)

	// create buffered chan results, with 5 spaces
	results := make(chan int, numJobs)

	// creates 3 goroutnes
	// sends 3 ints and 2 chans (jobs and result)
	msg := ""
	for w := 1; w <= 3; w++ {
		msg = strconv.Itoa(w) + " sent worker"
		go worker(w, jobs, results, msg)
	}

	// Sends 5 ints throw the channel jobs
	// jobs receive work on the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// Close jobs
	close(jobs)

	// Recibes 5 ints throw the channel result
	// results send the corresponding results on results.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
