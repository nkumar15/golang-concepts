package main

import (
	"github.com/nkumar15/golang-design-patterns/singleton"
	"fmt"
)



func main() {

	first := singleton.GetInstance()
	sec := singleton.GetInstance()

	if first == sec {
		fmt.Println("Both are same")
	} else {
		fmt.Println("Both are not same")
	}

}
