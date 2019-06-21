package main

import (
	"fmt"
	"time"
)

//func main() {
//	ch := make(chan string, 2)
//	ch <- "naveen"
//	ch <- "paul"
//	fmt.Println(<- ch)
//	fmt.Println(<- ch)
//}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}
