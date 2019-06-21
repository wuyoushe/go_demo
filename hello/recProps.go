package main

import "fmt"

func rectProps1(length, width float64)(float64, float64) {
	var area = length * width;
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	area, _ := rectProps1(10.8, 5.6)
	fmt.Printf("Area %f", area)
}
