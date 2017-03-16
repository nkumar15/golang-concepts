package main

import "fmt"

func main() {
	// Initialize way 1
	// initialize slice like this when you know contents ahead
	names := []string{"Neeraj", "Oswell", "Slava", "Lorenzo"}
	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("****************")
	names = append(names, "Yaron")
	for _, name := range names {
		fmt.Println(name)
	}

	// Initialize way 3, empty slice
	// When you don't know contents. Need to add at runtime.

	var numberList []int
	fmt.Println("Enter numbers. Input -1 to terminate prompt")
	for {
		var num int
		_, err := fmt.Scanf("%d", &num)

		if err != nil || num == -1 {
			break
		}

		numberList = append(numberList, num)
	}

	for _, n := range numberList {
		fmt.Println(n)
	}

	// Initialize way 3
	// when you need to know the length before
	newNames := make([]string, len(names))
	fmt.Println("Enter names. Input -1 to terminate prompt")
	//idx := 0

	for {
		var name string
		_, err := fmt.Scanf("%s", &name)

		if err != nil || name == "-1" {
			break
		}

		//If index approach is used you cannot cross this length of slice. However appends make it grow
		//newNames[idx] = names
		//idx = idx + 1
		newNames = append(newNames, name)
	}

	for _, n := range newNames {
		fmt.Println(n)
	}

	// Initialize way 4
	// when you need to know the length and capacity before
	// Its a matter of choice
	newNames1 := make([]string, len(names), 10)
	fmt.Println("Enter names. Input -1 to terminate prompt")

	for {
		var name string
		_, err := fmt.Scanf("%s", &name)

		if err != nil || name == "-1" {
			break
		}

		newNames1 = append(newNames1, name)
	}

	for _, n := range newNames1 {
		fmt.Println(n)
	}

}
