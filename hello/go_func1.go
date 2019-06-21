package main

import "fmt"

/*在Go语言中，函数声明的通用语法
func functionname(parametername type) returnType{}
函数中的参数列表和返回值并非是必须得所以下面这个函数的声明也是有效的
func functionname(){}
*/

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
}
