package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function", num)
}

func main() {
	//var a [3] int
	//a[0] = 12
	//a[1] = 78
	//a[2] = 50
	//fmt.Println(a)
	//a := [3]int{12, 78, 50}
	//fmt.Println(a)
	//a := [...]string{"usa","china","india"}
	//b := a										//Go的数组是值类型而不是引用类型
	//b[0] = "Singapore"
	//fmt.Println("a is:", a)
	//fmt.Println("b is:", b)
	//num := [...]int{5, 6, 7, 8, 8}
	//fmt.Println("before passing to function ", num)
	//changeLocal(num)
	//fmt.Println("after passing to function ", num)
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)

	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)

}
