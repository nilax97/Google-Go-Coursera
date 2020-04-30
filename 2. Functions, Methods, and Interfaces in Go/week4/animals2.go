package main

import "fmt"
import "strings"

type Animals interface {
	Eat()
	Move()
	Speak()
	Name() string
}

type Cow struct {
	name string
}
func (c Cow) Eat() {
	fmt.Printf("grass\n")
}
func (c Cow) Move() {
	fmt.Printf("walk\n")
}
func (c Cow) Speak() {
	fmt.Printf("moo\n")
}
func (c Cow) Name() string {
	return c.name
}

type Bird struct {
	name string
}
func (c Bird) Eat() {
	fmt.Printf("worms\n")
}
func (c Bird) Move() {
	fmt.Printf("fly\n")
}
func (c Bird) Speak() {
	fmt.Printf("peep\n")
}
func (c Bird) Name() string {
	return c.name
}

type Snake struct {
	name string
}
func (c Snake) Eat() {
	fmt.Printf("mice\n")
}
func (c Snake) Move() {
	fmt.Printf("slither\n")
}
func (c Snake) Speak() {
	fmt.Printf("hsss\n")
}
func (c Snake) Name() string {
	return c.name
}

func main() {

	var request_type string
	sli := make([]Animals, 0)
	for {
		fmt.Printf(">")
		fmt.Scan(&request_type)
		if strings.Compare(request_type, "X") == 0 {
			break
		}
		if strings.Compare(request_type, "newanimal") == 0 {
			var animal_name string
			fmt.Scan(&animal_name)
			var animal_type string
			fmt.Scan(&animal_type)
			var ani1 Animals
			if strings.Compare(animal_type, "cow") == 0 {
				cow1 := Cow{animal_name}
				ani1 = cow1
				sli = append(sli,ani1)
			} else if strings.Compare(animal_type, "bird") == 0 {
				bird1 := Bird{animal_name}
				ani1 = bird1
				sli = append(sli,ani1)
			} else if strings.Compare(animal_type, "snake") == 0 {
				snake1 := Snake{animal_name}
				ani1 = snake1
				sli = append(sli,ani1)
			} else {
				fmt.Printf("Enter a valid animal type\n")
			}
		} else if strings.Compare(request_type, "query") == 0 {
			var animal_name string
			fmt.Scan(&animal_name)
			var info_type string
			fmt.Scan(&info_type)

			var flag int
			for _, ani1 := range sli {
				if strings.Compare(animal_name,ani1.Name()) == 0 {
					if strings.Compare(info_type,"eat") == 0 {
						ani1.Eat()
						flag = 1
						break
					} else if strings.Compare(info_type,"move") == 0 {
						ani1.Move()
						flag = 1
						break
					} else if strings.Compare(info_type,"speak") == 0 {
						ani1.Speak()
						flag = 1
						break
					} else {
						fmt.Printf("Enter a information request\n")
					}
				}
			}
			if flag == 0 {
				fmt.Printf("Animal name not found\n")
			}
		}
	}
}
