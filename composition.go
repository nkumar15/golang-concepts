package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) String() {
	fmt.Printf("Name  : %s\n", p.Name)
}

type Employee struct {
	*Person
	Salary int
}

func (e *Employee) String() {
	e.Person.String()
	fmt.Printf("Salary: %d\n", e.Salary)
}

func main() {
	emp := Employee{Person: &Person{Name: "Ravi"}, Salary: 800}
	emp.String()
}
