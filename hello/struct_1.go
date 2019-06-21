package main

import "fmt"

type Employ struct{
	firstName, lastName string
	age, salary 		int
}

func main() {
	emp1 := Employ{
		firstName: "Sam",
		age:		25,
		salary:		500,
		lastName:	"Anderson",
	}

	emp2 := Employ{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee1", emp1)
	fmt.Println("Employee2", emp2)

	emp8 := &Employ{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name", (*emp8).firstName)
	fmt.Println("Age:",(*emp8).age)
}
