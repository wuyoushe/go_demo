package main

import (
	"fmt"
)

type Employee struct {
	name string
	salary int 
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee {
		name:	"Sam Adolf",
		salary: 5000,
		currency: "$",
	}
	emp1.displaySalary()
}