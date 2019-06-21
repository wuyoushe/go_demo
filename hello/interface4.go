package main

import "fmt"

type Describer interface {
	Describe()
}

type Person1 struct {
	name string
	age int
}

//值接受者
func(p Person1) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

//指针接受者

type Address1 struct {
	state string
	country string
}

//使用指针接受者实现
func(a *Address1) Describe() {
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person1{"Sam", 25}
	d1 = p1
	d1.Describe()

	p2 := Person1{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address1{"Washington", "USA"}

	d2 = &a
	d2.Describe()
}
