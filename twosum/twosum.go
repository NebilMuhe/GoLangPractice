package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var value []int

	for i := 0; i < len(nums); i++ {
		isCorrect := false
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				value = append(value, i, j)
				isCorrect = true
				break
			}
		}
		if isCorrect {
			break
		}
	}

	return value
}

func main() {
	numbers := []int{1, 2, 3, 4}
	target := 7
	val := twoSum(numbers, target)

	fmt.Println(val)
}
