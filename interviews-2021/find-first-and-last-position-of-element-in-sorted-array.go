package main

// start: Jun 18 2021 4:40 PM
// end: Jun 18 2021 4:59 PM

func searchRange(nums []int, target int) []int {
	l := searchLeft(nums, target)
	r := searchRight(nums, target)
	if l < 0 || r < 0 || l > len(nums)-1 || r > len(nums)-1 || nums[l] != target || nums[r] != target {
		return []int{-1, -1}
	}
	return []int{l, r}
}

func searchLeft(nums []int, target int) int {
	s := 0
	e := len(nums) - 1
	for s <= e {
		mid := (s + e) / 2
		if nums[mid] < target {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return e + 1
}

func searchRight(nums []int, target int) int {
	s := 0
	e := len(nums) - 1
	for s <= e {
		mid := (s + e) / 2
		if nums[mid] <= target {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return s - 1
}
