package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Resource struct {
	num int
}

func NewResource() *Resource {
	return &Resource{rand.Intn(10000)}
}

type ProducerConsumer struct {
	closing bool
	terminate, closed chan bool
	itemQueue []*Resource
	itemLength, itemLimit int
	numProducer, numConsumer int
	producer, consumer chan* Resource
}

func NewProducerConsumer(itemLimit int) *ProducerConsumer {
	if itemLimit < 1 {
		return nil
	}

	pc := new(ProducerConsumer)
	pc.itemLimit = itemLimit
	pc.itemQueue = make([]*Resource, 0, pc.itemLimit)

	pc.terminate, pc.closed = make(chan bool), make(chan bool)
	pc.producer, pc.consumer = make(chan *Resource), make(chan *Resource)

	go pc.ResourceManager()

	return pc
}

// ResourceManager coordinates with producers and consumers.
// It also imposes the maximum items that can be stored in pending queue.
// It waits for all producers and consumers to terminate before terminating self.
func (this *ProducerConsumer) ResourceManager() {
	for !this.closing {
		var firstItem *Resource = nil
		var fetch, send chan *Resource = nil, nil

		if this.itemLength < this.itemLimit {
			fetch = this.producer
		}

		if this.itemLength > 0 {
			send = this.consumer
			firstItem = this.itemQueue[0]
		}

		select {
			case <-this.terminate:
				this.closing = true
			case send<- firstItem:
				this.itemQueue = this.itemQueue[1:]
				this.itemLength -= 1
			case item := <-fetch:
				this.itemQueue = append(this.itemQueue, item)
				this.itemLength += 1
		}
		fmt.Println("Queue length", this.itemLength)
	}

	close(this.producer)
	close(this.consumer)

	for this.numProducer != 0 || this.numConsumer != 0 {
		time.Sleep(2 * time.Second)
	}

	this.closed<- true
}

// Creates a producer thread on-demand.
// The producer must accept a write-only channel.
// The producer must return a delay in second before next call.
func (this *ProducerConsumer) CreateProducer(produce func(int, chan<- *Resource)(int)) {
	go func(id int) {
		defer func() { this.numProducer -= 1 }()
		for this.closing == false {
			seconds := produce(id, this.producer)
			fmt.Printf("Producer %d sleeping for %d seconds.\n", id, seconds)
			time.Sleep(time.Duration(seconds) * time.Second)
		}
	}(this.numProducer)
	this.numProducer += 1
}

// Creates a consumer thread on-demand.
// The consumer must accept a read-only channel.
// The consumer must return a delay in second before next call.
func (this *ProducerConsumer) CreateConsumer(consume func(int, <-chan *Resource)(int)) {
	go func(id int) {
		defer func() { this.numConsumer -= 1 }()
		for this.closing == false {
			seconds := consume(id, this.consumer)
			fmt.Printf("Consumer %d sleeping for %d seconds.\n", id, seconds)
			time.Sleep(time.Duration(seconds) * time.Second)
		}
	}(this.numConsumer)
	this.numConsumer += 1
}

// Notifies the ResourceManager to prepare termination.
// Waits for the ResourceManager to terminate.
func (this *ProducerConsumer) Wait() {
	this.terminate<- true
	<-this.closed
}

// Producer.
func produce(id int, put chan<- *Resource) int {
	item := NewResource()
	fmt.Printf("Producer %d, item: %v.\n", id, *item)
	put<- item
	return rand.Intn(10)
}

// Consumer.
func consume(id int, get <-chan *Resource) int {
	select {
		case <-time.NewTimer(5*time.Second).C:
			fmt.Printf("Consumer %d, item: NA.\n", id)
		case item := <-get:
			fmt.Printf("Consumer %d, item: %v.\n", id, *item)
	}
	return rand.Intn(10)
}

func main() {
	// A ProducerConsumer object with maximum 1000 queued items.
	pc := NewProducerConsumer(1000)

	// Create 10 thread of producer and consumer each.
	for i := 0; i < 10; i++ {
		pc.CreateProducer(produce)
		pc.CreateConsumer(consume)
	}

	// Wait for user-input for terminating the program.
	fmt.Scanln()
	pc.Wait()
}
