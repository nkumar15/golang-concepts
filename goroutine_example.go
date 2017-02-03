package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
                // if sleep is not introduced
                // it shows the output in sequential order only
		amt := time.Duration(rand.Intn(25))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 1; i < 5; i++ {
		go PrintTable(i)
	}

	var input string
	fmt.Scanln(&input)
}
