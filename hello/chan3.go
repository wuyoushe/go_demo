package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
	//信道的声明 chan T
	//简短声明 a := make(chan int)
	//2 通过信道进行发送和接收
	//data := <- a 读取信道
	// a <- data 写入信道
}