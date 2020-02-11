/*
https://talks.golang.org/2012/waza.slide#39

Simple LoadBalancer between workers

Concurrency via channels and go routines

NOTES:
Channels are a typed conduit through which you can send and receive values with the channel operator
Channels are typed values that allow goroutines to synchronize and exchange information. ex: timerChan := make(chan time.Time)

Concurrency: Programming as the composition of independently executing processes.
Parallelism: Programming as the simultaneous execution of (possibly related) computations.

Concurrency is about dealing with lots of things at once.
Parallelism is about doing lots of things at once.

Not the same, but related.
Concurrency is about structure, parallelism is about execution.
Concurrency provides a way to structure a solution to solve a problem that may (but not necessarily) be parallelizable.

Concurrency plus communication(Channels)
Concurrency is a way to structure a program by breaking it into pieces that can be executed independently.
Communication is the means to coordinate the independent executions.
*/

package main

import (
	"fmt"
	"time"
)

//Work information
type Work struct {
	x, y, z int
}

//Perform some work looping over input channel
func worker(i int, in chan *Work, out chan *Work) {
	fmt.Printf("Start[%d]: %v\n", i,time.Now())
	for w := range in {
		w.z = i
		time.Sleep(time.Duration(1) * time.Second)
		out <- w
	}
	close(out)
	fmt.Printf("End[%d]: %v\n",i,time.Now())
}

//Create workers and launch go routines
func Run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < 2; i++ {
		go worker(i, in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
}

//Assign work to channel to keep workers busy
func sendLotsOfWork(in chan *Work) {
	for i:=0; i< 20; i++ {
		w1 := &Work{1*i,2*i,3*i}
		in <- w1
	}
	close(in)
	return
}

//Read the results
func receiveLotsOfResults(out chan *Work) {
	for o := range out {
		fmt.Printf("x=>%d y=>%d z=>%d\n", o.x,o.y,o.z)
	}
}

func main() {
	fmt.Println("Hello, playground")
	Run()
}
