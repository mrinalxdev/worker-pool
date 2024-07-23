package dispatcher

import (
	"sync"

	"worker-pool/job"
	"worker-pool/result"
	"worker-pool/worker"
)

type Dispatcher struct {
	Workers []*worker.Worker
	Jobs    <-chan job.Job
	Results chan<- result.Result
	wg      sync.WaitGroup
}

func New(numWorkers int, jobs <-chan job.Job, results chan<- result.Result) *Dispatcher {
	d := &Dispatcher{
		Workers: make([]*worker.Worker, numWorkers),
		Jobs:    jobs,
		Results: results,
	}

	for i := 0; i < numWorkers; i++ {
		d.Workers[i] = worker.New(i+1, jobs, results)
	}

	return d
}

func (d *Dispatcher) Run() {
	for _, w := range d.Workers {
		d.wg.Add(1)
		go func(w *worker.Worker) {
			defer d.wg.Done()
			w.Start()
		}(w)
	}
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}
