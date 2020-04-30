package main

import "fmt"

func Swap(nums []int, index int) {
	x := nums[index]
	nums[index] = nums[index+1]
	nums[index+1] = x
}

func BubbleSort(nums []int) {
	for i:=0; i<len(nums); i=i+1 {
		for j:=0; j<(len(nums)-i-1); j=j+1 {
			if nums[j] > nums[j+1] {
				Swap(nums,j)
			}
		}
	}
}

func main() {
	sli := make([]int, 0)
	fmt.Printf("Enter the space sperated numbers. Enter an 'X' when done\n")
	var inp int
	for i := 0; i< 10; i=i+1 {
		_, err := fmt.Scan(&inp)
		if err != nil {
			break
		}
		sli = append(sli,inp)
	}
	BubbleSort(sli)
	fmt.Printf("%v\n",sli)
}
