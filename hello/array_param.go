package main

import (
	"fmt"
)

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func main() {
	num:=[...] int{5, 6, 7, 8, 9}
	fmt.Println("before passing to function ", num)
	//当数组作为参数传递给函数时，它们是按值传递的，而原始数组保持不变
	changeLocal(num)
	fmt.Println("after passing to function ", num)

}