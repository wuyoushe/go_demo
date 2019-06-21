package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()								//	WaitGroup的Done()方法可以减少计数
}

func main() {
	no := 3
	var wg sync.WaitGroup 		//创建了WaitGroup起初始值卫零值
	for i :=0; i < no; i++ {
		//当我们调用WaitGroup的Add()方法并传入一个int时，WaitGroup的计数器会加上Add的传参。
		//要减少计数器，可以调用WaitGroup的Done()方法。
		wg.Add(1)				
		go process(i, &wg)
	}
	wg.Wait()					//wg.Wait()，确保Go主协程等待计数器变为0.
	fmt.Println("All go routines finished executing")
}




















