package main

import "fmt"

//我们可以创建一个具有缓冲（Buffer)的信道。
//ch := make(chan type, capacity) 要让一个信道有缓冲，capacity应该大于0，无缓冲信道的容量默认为0.
func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}
