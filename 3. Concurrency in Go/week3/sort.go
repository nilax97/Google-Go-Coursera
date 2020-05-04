package main

import "fmt"
import "sync"

var wg sync.WaitGroup

const UintSize = 32 << (^uint(0) >> 32 & 1)
const MaxInt = 1<<(UintSize-1) - 1

func Swap(nums []int, index int) {
	x := nums[index]
	nums[index] = nums[index+1]
	nums[index+1] = x
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i = i + 1 {
		for j := 0; j < (len(nums) - i - 1); j = j + 1 {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func PartSort(nums []int) {
	BubbleSort(nums)
	fmt.Printf("Sorted part array : %v\n", nums)
	wg.Done()
}

func Calc_min_index(nums []int) int {
	min := 0
	for i := 1; i < len(nums); i = i + 1 {
		if nums[i] < nums[min] {
			min = i
		}
	}
	return min
}

func main() {
	sli := make([]int, 0)
	fmt.Printf("Enter the numbers. Enter an 'X' when done\n")
	var inp int
	for {
		_, err := fmt.Scan(&inp)
		if err != nil {
			break
		}
		sli = append(sli, inp)
	}
	n := len(sli)
	num_splits := 4
	wg.Add(4)
	for i := 0; i < num_splits; i = i + 1 {
		go PartSort(sli[n*i/4 : n*(i+1)/4])
	}
	wg.Wait()
	sort_sli := make([][]int, 4)
	for i := 0; i < num_splits; i = i + 1 {
		sort_sli[i] = sli[n*i/4 : n*(i+1)/4]
	}
	final_sli := make([]int, 0)
	index := make([]int, 4)
	for i := 0; i < n; i = i + 1 {
		min_sli := make([]int, 0)
		for j := 0; j < num_splits; j = j + 1 {
			if index[j] < len(sort_sli[j]) {
				min_sli = append(min_sli, sort_sli[j][index[j]])
			} else {
				min_sli = append(min_sli, MaxInt)
			}
		}
		min_index := Calc_min_index(min_sli)
		final_sli = append(final_sli, sort_sli[min_index][index[min_index]])
		index[min_index] = index[min_index] + 1
	}

	fmt.Printf("%v\n", final_sli)
}
