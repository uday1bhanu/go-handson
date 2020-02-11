/*
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
cap, len  2 0
Put ->-  0
cap, len  2 1
Put ->-  1
cap, len  2 2
Put ->-  2
Read -<-  0
Read -<-  1
Read -<-  2
cap, len  2 0
Put ->-  3
cap, len  2 0
Put ->-  4
cap, len  2 1
Put ->-  5
cap, len  2 2
Put ->-  6
Read -<-  3
Read -<-  4
Read -<-  5
Read -<-  6
cap, len  2 0
Put ->-  7
cap, len  2 0
Put ->-  8
cap, len  2 1
Put ->-  9
cap, len  2 2

Put  7,8,9 is not sent to channel as the buffer is not full

*/

package main

import (
		"fmt"
		)

func main() {
	ch := make(chan int, 2)
	fmt.Println("cap, len ", cap(ch), len(ch))
	go read(ch)
	
	for i:=0; i<10; i++ {
		fmt.Println("Put ->- ",i)
		ch <- i
		fmt.Println("cap, len ", cap(ch), len(ch))
	}
}

func read(c chan int) {
	for {
		fmt.Println("Read -<- ",<-c)
	}
}
