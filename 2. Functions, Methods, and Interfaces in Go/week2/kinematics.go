package main

import "fmt"

func GenDisplaceFn(acceleration float64, init_velocity float64, init_displacement float64) func (float64) float64 {
	return func (time float64) float64 {
		acc_displacement := (0.5 * acceleration * time * time)
		init_vel_displacement := (init_velocity * time)
		return (acc_displacement + init_vel_displacement + init_displacement)
	}
}

func main() {
	var acceleration float64
	fmt.Printf("Please enter the acceleration\n")
	fmt.Scan(&acceleration)

	var init_velocity float64
	fmt.Printf("Please enter the initial velocity\n")
	fmt.Scan(&init_velocity)

	var init_displacement float64
	fmt.Printf("Please enter the initial displacement\n")
	fmt.Scan(&init_displacement)

	fn := GenDisplaceFn(acceleration, init_velocity, init_displacement)

	var time float64
	fmt.Printf("Please enter the time\n")
	fmt.Scan(&time)

	fmt.Printf("The total displacement is : %v\n",fn(time))
}
