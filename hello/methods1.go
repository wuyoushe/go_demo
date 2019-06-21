package main

import "fmt"

//创建方法的语法
//func (t Type) methodName(parameter list) {}
type Employee struct {
	name string
	salary int
	currency string
}

/*displaySalary()方法将Employee作为接收器类型*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency,e.salary)
}

func main() {
	emp1 := Employee{
		name: "Sam Adlof",
		salary: 5000,
		currency: "$",
	}
	emp1.displaySalary()

}
