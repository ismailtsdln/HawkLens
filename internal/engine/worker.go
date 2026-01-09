package engine

import (
	"context"
	"sync"

	"github.com/ismailtsdln/HawkLens/pkg/plugins"
)

// Job represents a scan task for a specific platform
type Job struct {
	Platform string
	Query    string
}

// ResultWrapper wraps plugin results with potential errors
type ResultWrapper struct {
	Platform string
	Results  []plugins.Result
	Error    error
}

// Dispatcher manages the worker pool
type Dispatcher struct {
	workerCount int
	jobQueue    chan Job
	results     chan ResultWrapper
	wg          sync.WaitGroup
}

// NewDispatcher creates a new Job Dispatcher
func NewDispatcher(workerCount int) *Dispatcher {
	return &Dispatcher{
		workerCount: workerCount,
		jobQueue:    make(chan Job, 100),
		results:     make(chan ResultWrapper, 100),
	}
}

// Run starts the worker pool
func (d *Dispatcher) Run(ctx context.Context) {
	for i := 0; i < d.workerCount; i++ {
		go d.worker(ctx)
	}
}

// worker handles jobs from the queue
func (d *Dispatcher) worker(ctx context.Context) {
	for job := range d.jobQueue {
		p, err := plugins.GetPlugin(job.Platform)
		if err != nil {
			d.results <- ResultWrapper{Platform: job.Platform, Error: err}
			d.wg.Done()
			continue
		}

		res, err := p.Fetch(ctx, job.Query)
		d.results <- ResultWrapper{Platform: job.Platform, Results: res, Error: err}
		d.wg.Done()
	}
}

// Submit adds jobs to the pool
func (d *Dispatcher) Submit(platform, query string) {
	d.wg.Add(1)
	d.jobQueue <- Job{Platform: platform, Query: query}
}

// Wait blocks until all jobs are finished
func (d *Dispatcher) Wait() {
	d.wg.Wait()
	close(d.jobQueue)
	close(d.results)
}

// Results returns the results channel
func (d *Dispatcher) Results() <-chan ResultWrapper {
	return d.results
}
