package main

import (
	"fmt"
	"time"
)

//Write into channel
func pinger(c chan string) {
	for {
		c <- "ping"
	}
}

//Write into channel
func ponger(c chan string) {
	for {
		c <- "pong"
	}
}

//Read from channel
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
        // create a channel
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
