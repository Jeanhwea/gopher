package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id  int
	num int
}

type Result struct {
	job Job
	out int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func sumOfDigits(num int) int {
	sum := 0
	n := num
	for n != 0 {
		digit := n % 10
		sum += digit
		n /= 10
		time.Sleep(time.Duration(digit*10) * time.Millisecond)
	}
	return sum
}

func worker(wid int, wg *sync.WaitGroup) {
	fmt.Printf("Worker(#%d) start.\n", wid)
	for job := range jobs {
		results <- Result{
			job: job,
			out: sumOfDigits(job.num),
		}
	}
	fmt.Printf("Worker(#%d) finished.\n", wid)
	wg.Done()
}

func allocateJobs(numJobs int) {
	for i := 0; i < numJobs; i++ {
		jobs <- Job{
			id:  i,
			num: rand.Intn(999),
		}
	}
	fmt.Println("Finished allocate jobs.")
	close(jobs)
}

func result(done chan bool) {
	for res := range results {
		fmt.Printf(
			"Job id %d, input random %d, sum of digits = %d\n",
			res.job.id,
			res.job.num,
			res.out,
		)
	}
	fmt.Println("finish Print result")
	done <- true
}

func createWorkerPool(numWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go worker(i, &wg)
		fmt.Println("Create worker pool", i)
	}
	wg.Wait()
	fmt.Println("All worker pool done their jobs.")
	close(results)
}

func main() {
	startTime := time.Now()

	numJobs := 100
	go allocateJobs(numJobs)

	done := make(chan bool)
	go result(done)

	numWorker := 10
	createWorkerPool(numWorker)

	<-done
	endTime := time.Now()
	elapsed := endTime.Sub(startTime).Seconds()
	fmt.Printf("total time is %f second\n", elapsed)
}
