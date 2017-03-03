package main

import (
	"fmt"
)

func sum(nums ...int) {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	fmt.Println("Sum is: ", sum)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3, 4, 5, 6, 7)

}
