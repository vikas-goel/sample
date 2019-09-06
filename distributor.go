// Given a function Copy(src, dst int) that copies a file from node src to dst.
//
// A node has bandwidth to distribute the file to maximum three other nodes at
// a time. A node can distribute the file only when the file is available to it.
// Initially, the file is available only on node-0.
//
// Implement a synchronous function distribute(n int) that distributes the file to all the n nodes.

package main

import (
	"fmt"
	"sync"
	"time"
)

func Copy(src, dst int) {
	fmt.Printf("Copy from %3v to %3v in progress\n", src, dst)
	time.Sleep(time.Millisecond * time.Duration(1000))
	fmt.Printf("Copy from %3v to %3v is completed\n", src, dst)
}

type Distributor struct {
	maxcp, nodes int
	wg sync.WaitGroup
	lock sync.RWMutex
}

func (this *Distributor) Distribute(src int) {
	notify := make(chan bool, this.maxcp)

	check := false
	for i := 1; i <= this.maxcp && this.nodes > 0; i++ {
		check = true
		go this.copyAndFanOut(src, notify)
	}

	// Start file distribution as soon as bandwidth available and there
	// are nodes that need the file.
	for check {
		if check = <-notify; check {
			go this.copyAndFanOut(src, notify)
		}
	}

	if src == 0 {
		this.wg.Wait()
	}
}

func (this *Distributor) getNextNodeId() (nextNode int) {
	if this.nodes > 0 {
		this.lock.Lock()

		if this.nodes > 0 {
			nextNode = this.nodes
			this.nodes--
		}

		this.lock.Unlock()
	}

	return
}

func (this *Distributor) copyAndFanOut(src int, notify chan<- bool) {
	this.wg.Add(1)
	defer this.wg.Done()

	// Find the next destination node.
	dst := this.getNextNodeId()

	if dst > 0 {
		Copy(src, dst)

		// Once copy operation is complete on the destination,
		// make the destination node start distributing the file to
		// others.
		go this.Distribute(dst)

		// Notify the source about completion of Copy.
		notify <- true
	} else {
		// Notify the source about no more need of Copy.
		notify <- false
	}
}

func distribute(nodes int) {
	distributor := Distributor{maxcp: 3, nodes: nodes}
	distributor.Distribute(0)
}

func main() {
	distribute(100)
}
