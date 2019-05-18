package main

import (
	"fmt"
	"sync"
	"time"
)

type Resource struct {
	id, count int
}

func (this *Resource) doSomething() {
	interval := 10 * time.Second
	time.Sleep(interval)
	this.count += 1
	fmt.Printf("Id:%v, Count:%v\n", this.id, this.count)
}

type Scheduler struct {
	in chan *Resource
	wg sync.WaitGroup
}

func (this *Scheduler) process() {
	for r := range this.in {
		r.doSomething()
		this.wg.Done()
	}

}

func (this *Scheduler) schedule(r *Resource) {
	if r == nil {
		return
	}

	this.wg.Add(1)
	go func() { this.in <- r }()
}

func (this *Scheduler) finish() {
	this.wg.Wait()
	close(this.in)
}

func newScheduler(numThreads int) *Scheduler {
	scheduler := Scheduler{ in: make(chan *Resource) }

	for i := 0; i < numThreads; i++ {
		go scheduler.process()
	}

	return &scheduler
}

func main() {
	const numThreadPool = 5
	sched := newScheduler(numThreadPool)

	for i := 0; i < 10; i++ {
		r := Resource{id: i}
		sched.schedule(&r)
		fmt.Println("Scheduled ", i)
	}

	fmt.Println("Scheduled all.")
	sched.finish()
	//time.Sleep(60*time.Second)
}
