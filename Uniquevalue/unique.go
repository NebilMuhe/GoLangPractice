package main

import "fmt"

func uniqueSlice(numbers []int) []int {
	val := []int{}
	valMaps := make(map[int]bool)

	for _, num := range numbers {
		if valMaps[num] {
			continue
		} else {
			valMaps[num] = true
			val = append(val, num)
		}
	}
	return val
}

func main() {
	numbers := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7}
	nums := uniqueSlice(numbers)
	fmt.Println(nums)
}
