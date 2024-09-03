package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var solution []int
	for i := 0 ; i < len(nums); i++ {
		for j := len(nums)-1; j > i ; j--{
			if nums[i]+nums[j] == target {
				solution = append(solution, i, j)
			}
		}
	}
	return solution
}

func main() {
	fmt.Println("Solution is: ", twoSum([]int{2, 5, 9, 11}, 13))
}