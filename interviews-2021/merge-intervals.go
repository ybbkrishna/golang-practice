package main

import (
	"math"
	"sort"
)

// start: Jun 18 2021 4:24 PM
// end: Jun 19 2021 4:39 PM

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	mergedIntervals := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := mergedIntervals[len(mergedIntervals)-1]
		if isIntersecting(last, intervals[i]) {
			mergedIntervals[len(mergedIntervals)-1] = mergeTwoIntervals(last, intervals[i])
			continue
		}
		mergedIntervals = append(mergedIntervals, intervals[i])
	}
	return mergedIntervals
}

func mergeTwoIntervals(int1, int2 []int) []int {
	return []int{int1[0], int(math.Max(float64(int1[1]), float64(int2[1])))}
}

func isIntersecting(int1, int2 []int) bool {
	return !(int2[0] > int1[1])
}
