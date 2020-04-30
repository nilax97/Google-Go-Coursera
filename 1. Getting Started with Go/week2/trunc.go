package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	var y int = int(x)
	fmt.Printf("%d\n", y)
}
