/*发送和接收默认是阻塞的，当数据发送到信道时，程序会在发送数据的语句处发生阻塞，知道有Go协程从信道读取数据才会接触阻塞
与此类似，当读取信道的数据时，如果没有其他的协程把数据写入信道，那么读取过程就会一直阻塞着*/
package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <-i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}