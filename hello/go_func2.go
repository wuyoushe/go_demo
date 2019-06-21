package main

import "fmt"

/*
go语言支持多返回值
_在Go中被用作空白符，可以用做表示任何类型的任何值
例如我们只用到函数的一个返回值area可以 area_ := rectProps(10.8, 5.6)
*/

func rectProps(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Permeter %f", area, perimeter)
}
