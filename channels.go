package main

import (
	"fmt"
	"time"
)

type Work struct {
	x, y, z int
}

func worker(i int, in <-chan *Work, out chan<- *Work) {
	fmt.Printf("Start[%d]: %v\n", i,time.Now())
	for w := range in {
		w.z = i
		time.Sleep(time.Duration(1) * time.Second)
		out <- w
	}
	fmt.Printf("End[%d]: %v\n",i,time.Now())
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < 20; i++ {
		go worker(i, in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
}

func sendLotsOfWork(in chan *Work) {
	for i:=0; i< 200; i++ {
		w1 := &Work{1*i,2*i,3*i}
		in <- w1
	}
	close(in)
	return
}

func receiveLotsOfResults(out chan *Work) {
	for {
		o, ok := <- out
		if ok == true {
			fmt.Printf("x=>%d y=>%d z=>%d\n", o.x,o.y,o.z)
		} else {
			fmt.Println("Channel broke")
		}
	}
}

func main() {
	fmt.Println("Hello, playground")
	Run()
}
