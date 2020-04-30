package main

import "fmt"
import "strconv"
import "Strings"

func main() {
	sli := make([]int, 3)
	var x string
	for i := 0; true; i = i + 1 {
		fmt.Scan(&x)
		if s, err := strconv.Atoi(x); err == nil {
			if i < len(sli) {
				sli[i] = s
			} else {
				sli = append(sli, s)
			}
			for j := 0; j <= i; j = j + 1 {
				for k := j + 1; k <= i; k = k + 1 {
					if sli[j] > sli[k] {
						var temp int = sli[j]
						sli[j] = sli[k]
						sli[k] = temp
					}
				}
			}
			fmt.Printf("%v\n", sli)
		} else {
			if strings.Compare(x, "X") == 0 {
				break
			} else {
				fmt.Printf("Invalid input\n")
			}
		}
	}
	fmt.Printf("%v\n", sli)
}
