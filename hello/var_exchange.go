package main

import "fmt"

func main() {
	i := 55
	j := 67.8
	//sum := i + j		//invalid operation: i + j (mismatched types int and float64)
	sum := i + int(j)
	fmt.Println(sum)
}
