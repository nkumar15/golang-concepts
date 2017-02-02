package main

import "fmt"

func main() {

	var nums [10]int
	var search_num int
	var discard string

	fmt.Println("Enter 10 elements: ")
	for i := 0; i < 10; {

		fmt.Print("element ", i+1, ":")
		_, err := fmt.Scanln(&nums[i])

		if err != nil {
		        fmt.Scanln(&discard)
			fmt.Println("Bad input. Please enter again")
		} else {
			i++
		}
	}

	fmt.Print("Enter element to search ")
	_, err := fmt.Scanln(&search_num)

	if err != nil {
		fmt.Scanln(&discard)
		fmt.Println("Bad input. Exiting..")
		return
	}

	for i := 0; i < 10; i++ {
		if search_num == nums[i] {
			fmt.Println("Number found")
			return
		}
	}

	fmt.Println("Number not found")
	return
}
