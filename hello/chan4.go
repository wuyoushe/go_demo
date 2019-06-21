package main

import "fmt"

func hello(done chan bool){
	fmt.Print("Hello world goroutine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<- done
	fmt.Println("main function")
}

