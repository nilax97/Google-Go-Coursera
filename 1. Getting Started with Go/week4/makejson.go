package main

import "fmt"
import "encoding/json"

func main() {
	var name string
	var address string
	fmt.Printf("Enter the name : ")
	fmt.Scan(&name)
	fmt.Printf("Enter the address : ")
	fmt.Scan(&address)

	idmap := make(map[string]string)
	idmap["name"] = name
	idmap["address"] = address

	barr, _ := json.Marshal(idmap)
	fmt.Printf(string(barr) + "\n")
}
