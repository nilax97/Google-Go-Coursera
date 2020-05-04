package main

import "fmt"
import "time"

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	time.Sleep(1 * time.Nanosecond) // Comment this to see the race condition
	return i
}

/*
Race Condition
A data race happens when two goroutines access the same variable concurrently, and at least one of the access is a write instruction.
Detecting a race
Use the command
	go run -race race.go to generate the race report
	uncomment line 15
	go run -race race.go to generate the race report
Conclusion
•	Result one and two generates 2 different numbers after adding wait period
•	Adding the pause caused the result to be 5 (should be 0)
•	-race parameter generates a report to find data races.
•	The report describes when the write value happened
The example provided commenting/uncommenting line 15 shows the difference of context switching
*/
