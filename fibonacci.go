package main

import (
	"fmt"
	"time"
)

func fib(n int, c chan int) {
	x,y := 0,1
	
	for i:=0; i< n; i++ {
		c <- x
		x,y = y,x+y

	}
	close(c)
}

func print(c chan int){
	for x := range c {
		fmt.Println("Reading " ,x)
	}
}

func main() {
	c := make(chan int, 10)
	go fib(cap(c), c)
	go print(c)
	time.Sleep(1)
}