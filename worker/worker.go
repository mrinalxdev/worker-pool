package worker

import (
	"fmt"
	"math/rand"
	"time"
	"worker-pool/job"
	"worker-pool/result"
)

type Worker struct {
	ID      int
	Jobs    <-chan job.Job
	Results chan<- result.Result
}

func New(id int, jobs <-chan job.Job, results chan<- result.Result) *Worker {
	return &Worker{
		ID:      id,
		Jobs:    jobs,
		Results: results,
	}
}

func (w *Worker) Start() {
	go func() {
		for j := range w.Jobs {
			fmt.Printf("Worker %d started job %d\n", w.ID, j.ID)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			output := fmt.Sprintf("Processed %s", j.Description)
			w.Results <- result.New(j.ID, output)
			fmt.Printf("Worker %d finished job %d\n", w.ID, j.ID)
		}
	}()
}
