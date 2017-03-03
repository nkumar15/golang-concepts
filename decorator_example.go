package main

import (
	"fmt"
	"time"
        "log"
)

func doSomeTask1() {
	fmt.Println("Hey I have done a pretty time consuming task1")
}


func doSomeTask2() {
	fmt.Println("Hey I have done a pretty time consuming task2")
}
func logger_decorator(inner func()) {
	start := time.Now()
	inner()
	log.Printf("total time taken %s", time.Since(start))
}

func main() {
	logger_decorator(doSomeTask1)
	logger_decorator(doSomeTask2)
}
