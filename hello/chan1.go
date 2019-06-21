package main

import "fmt"

func main() {
	//所有的信道都关联一个类型。信道只能运输这种类型的数据
	//chan T T表示类型的信道
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		//make来定义信道
		//简洁定义 a := make(chan int)
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
		//读取信道 data := <- a
		//写入信道 a <- data
	}
}