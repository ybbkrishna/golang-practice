package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for end >= start {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func searchRepeatLowerBound(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for end >= start {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end + 1
}

func searchRepeatHigherBound(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for end >= start {
		mid := (start + end) / 2
		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start - 1
}

func main() {
	arr := []int{1, 4, 7}
	fmt.Println(search(arr, 2))
	fmt.Println(searchRepeatLowerBound(arr, 2))
	fmt.Println(searchRepeatHigherBound(arr, 2))
}
