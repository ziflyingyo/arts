package main

import "fmt"

func main() {
	nums := []int{2, 1, 7, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] != target {
				continue
			} else {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}

	}
	return nil
}
