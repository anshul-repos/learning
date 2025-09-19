package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	NumberOfWorkers int
	JobQueue        chan Job
	Result          chan int
	wg              sync.WaitGroup
}

type Job struct {
	JobID int
}

func NewWorkerPool(nJobs, nWorkers int) *WorkerPool {
	return &WorkerPool{
		NumberOfWorkers: nWorkers,
		JobQueue:        make(chan Job, nJobs),
		Result:          make(chan int, nJobs),
	}
}

func (wp *WorkerPool) AddJobs(nJobs int) {
	for i := 1; i <= nJobs; i++ {
		wp.JobQueue <- Job{JobID: i}
	}
}

func (wp *WorkerPool) StartWorker() {
	for i := 1; i <= wp.NumberOfWorkers; i++ {
		wp.wg.Add(1)
		go wp.Worker(i)
	}
}

func (wp *WorkerPool) Worker(wId int) {
	defer wp.wg.Done()

	for job := range wp.JobQueue {
		fmt.Println()
		fmt.Printf("\nWorker: `%d` started executing job: %d \n", wId, job.JobID)
		time.Sleep(time.Second * 2)
		fmt.Println()
		fmt.Printf("\nWorker: `%d` done executing job: %d \n", wId, job.JobID)
		wp.Result <- job.JobID
	}
}

func (wp *WorkerPool) CollectResult() {
	for job := range wp.Result {
		fmt.Printf("\nResults for job: %d\n", job)
	}
}

func main() {

	// initialize workerpool
	nWorkers, nJobs := 5, 20
	wp := NewWorkerPool(nJobs, nWorkers)

	// add jobs
	wp.AddJobs(nJobs)
	close(wp.JobQueue)

	// start worker

	wp.StartWorker()
	wp.wg.Wait()

	close(wp.Result)
	// collect result
	wp.CollectResult()

}
