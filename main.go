package main

import (
	"fmt"
	"log"
	"sync"

	"worker-pool/config"
	"worker-pool/dispatcher"
	"worker-pool/job"
	"worker-pool/result"
)

func main() {
	cfg := config.Load()

	jobs := make(chan job.Job, cfg.NumJobs)
	results := make(chan result.Result, cfg.NumJobs)
	d := dispatcher.New(cfg.NumWorkers, jobs, results)
	d.Run()

	// creating and dispatching jobs
	go func() {
		for i := 1; i <= cfg.NumJobs; i++ {
			jobs <- job.New(i, fmt.Sprintf("Task %d", i))
		}
		close(jobs)
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for r := range results {
			log.Printf("Job %d completed. Result: %s\n", r.JobID, r.Output)
		}
	}()

	d.Wait()
	close(results)
	wg.Wait()
	log.Println("All jobs completed.")
}
