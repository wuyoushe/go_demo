package main

import "fmt"

type Employee1 struct {
	name string
	salary int
	currency string
}

func (e Employee1) displaySalary1(){
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee1{
		name: "Sam Adolf",
		salary: 5000,
		currency:"$",
	}
	emp1.displaySalary1()
}
