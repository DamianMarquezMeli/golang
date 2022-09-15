package main

import (
	"fmt"
	"runtime"
	"time"

	cnt "github.com/devpablocristo/golang-examples/std-lib/concurrency/01-goroutines/05-add/count"
)

func main() {
	numbers := cnt.GenerateNumbers(1e7)

	// Comparsion Add Sequeential vs Add Concurrent

	numOfCores := runtime.NumCPU()
	fmt.Println("Number of CPUs: ", numOfCores)
	fmt.Println(runtime.GOARCH, runtime.GOOS)

	t := time.Now()
	sum := cnt.Add(numbers)
	fmt.Printf("Sequential Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))

	t = time.Now()
	sum = cnt.AddConcurrent(numbers)
	fmt.Printf("Concurrent Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
}
