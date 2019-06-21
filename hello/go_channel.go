/*信道可以想象成Go协程之间通信的管道 chan T表示T类型的管道
简洁定义 a := make(chan int)
通过信道进行发送和接收
data := <- a 	读取信道a
a <- data		写入信道a

*/
package main

import (
	"fmt"
)

func main(){
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Println("Type of a is %T", a)
	}
}