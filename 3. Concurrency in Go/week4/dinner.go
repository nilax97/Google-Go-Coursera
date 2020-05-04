package main

import "fmt"
import "sync"

var lock sync.Mutex
var wg sync.WaitGroup

var chops [5]int

func Request(philo int, status int) {
	left := philo
	right := (philo + 1) % 5
	for {
		lock.Lock()
		if status == 1 {
			if chops[left] == 0 && chops[right] == 0 {
				chops[left] = 1
				chops[right] = 1
				lock.Unlock()
				break
			} else {
				lock.Unlock()
				continue
			}
		} else if status == 2 {
			chops[left] = 0
			chops[right] = 0
			lock.Unlock()
			break
		}
	}
}

func Eat(philo int) {
	Request(philo, 1)
	fmt.Printf("starting to eat %d\n", philo+1)
	// fmt.Printf("%v\n",chops)
	Request(philo, 2)
	fmt.Printf("finished eating %d\n", philo+1)
	wg.Done()
}

func main() {
	wg.Add(15)
	for i := 0; i < 15; i = i + 1 {
		go Eat(i % 5)
	}
	wg.Wait()
}
