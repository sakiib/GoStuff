package main

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums {
			if i != j && x + y == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.