package main

import "fmt"
import "Strings"

func main() {
	var x string
	fmt.Scan(&x)
	var y string = strings.ToLower(x)
	if strings.Contains(y, "a") && strings.HasPrefix(y, "i") && strings.HasSuffix(y, "n") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
