package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup){
	fmt.Println("started Goroutine")
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()								//要减少计数器我们用Done()方法，wg.Done()可以让计数器递减
}

func main() {
	no := 3
	var wg sync.WaitGroup					//	创建WaitGroup其初始值为0
	for i := 0; i < no; i++ {
		wg.Add(1)						//WaitGroup使用计数器来工作
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
