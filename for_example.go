package main

import "fmt"

func main() {
	var arr [10]float64
	for i := 0; i < 10; i++ {
		arr[i] = float64(i * i)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%f ", arr[i])
	}
        fmt.Println()
	var sum float64
	for _, v := range arr {
		sum += v

	}

	fmt.Printf("%f\n", sum/float64(len(arr)))
}
