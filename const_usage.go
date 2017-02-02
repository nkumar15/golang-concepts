package main

import "fmt"

func main() {
	const name string = "neeraj"
	fmt.Println(name)
	// Below line causes compile error
	// name ="kumar"
}
