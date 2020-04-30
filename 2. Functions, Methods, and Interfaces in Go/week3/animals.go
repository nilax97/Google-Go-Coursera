package main

import "fmt"
import "strings"

type Animals struct {
	food_eaten        string
	locomotion_method string
	spoken_sound      string
}

func (a *Animals) Eat() string {
	return a.food_eaten
}

func (a *Animals) Move() string {
	return a.locomotion_method
}

func (a *Animals) Speak() string {
	return a.spoken_sound
}

func main() {

	var animal_name string
	var info_request string

	cow := Animals{food_eaten: "grass", locomotion_method: "walk", spoken_sound: "moo"}
	bird := Animals{food_eaten: "worms", locomotion_method: "fly", spoken_sound: "peep"}
	snake := Animals{food_eaten: "mice", locomotion_method: "slither", spoken_sound: "hiss"}

	for {
		fmt.Printf(">")
		fmt.Scan(&animal_name)
		var query_animal Animals
		if strings.Compare(animal_name, "cow") == 0 {
			query_animal = cow
		} else if strings.Compare(animal_name, "bird") == 0 {
			query_animal = bird
		} else if strings.Compare(animal_name, "snake") == 0 {
			query_animal = snake
		} else {
			if strings.Compare(animal_name, "X") == 0 {
				break
			}
			fmt.Printf("Wrong Animal\n")
			continue
		}
		fmt.Scan(&info_request)
		if strings.Compare(info_request, "eat") == 0 {
			fmt.Printf(query_animal.Eat() + "\n")
		} else if strings.Compare(info_request, "move") == 0 {
			fmt.Printf(query_animal.Move() + "\n")
		} else if strings.Compare(info_request, "speak") == 0 {
			fmt.Printf(query_animal.Speak() + "\n")
		} else {
			fmt.Printf("Wrong information requested\n")
			continue
		}
	}
}
