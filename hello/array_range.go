package main

import "fmt"

func main() {
	a:= [...]float64{67.7, 89.8, 21, 78}
	for i :=0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
	sum:= float64(0)
	//for _, v:= range a{}
	//忽略下标是通过用_空白标识符号替换索引来执行
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}