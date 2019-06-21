package main

import (
	"fmt"
	"geometry/rectangle"//导入自定义包，必须指定自定义包相对于工作区的src的相对路径
)

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	//在go中，任何以大写字母开头的变量或函数都是被导出的名字。其他包只能访问被导出的函数和变量
	//例如rectangle.Area，rectangle.Diagonal
	fmt.Println("area of rectangle %.2f\n",rectangle.Area(rectLen, rectWidth))

	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}