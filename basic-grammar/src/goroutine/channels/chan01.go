package main

import (
	"fmt"
	"time"
)

func main() {
	chanDemo()
}

func chanDemo() {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)

}
func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}
