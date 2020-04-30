package main

import "fmt"
import "io/ioutil"
import "strings"

func main() {
	type Names struct {
		firstname string
		lastname  string
	}
	var fname string
	fmt.Printf("Enter the file name : ")
	fmt.Scan(&fname)

	sli := make([]Names, 0)

	dat, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(dat), "\n")
	for i := 0; i < len(lines); i = i + 1 {
		res := strings.Split(lines[i], " ")
		var n1 Names
		n1.firstname = res[0]
		n1.lastname = res[1]
		sli = append(sli, n1)
	}

	for i := 0; i < len(sli); i = i + 1 {
		first_name := sli[i].firstname
		last_name := sli[i].lastname
		fmt.Printf("Firstname : %s, Lastname : %s\n", first_name, last_name)
	}
}
