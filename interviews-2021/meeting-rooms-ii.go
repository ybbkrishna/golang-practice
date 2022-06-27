package main

import (
	"fmt"
	"sort"
)

// start: Jun 18 2021 4:06 PM
// end: Jun 18 2021 4:20 PM

func minMeetingRooms(intervals [][]int) int {
	sortByStart := make([][]int, len(intervals))
	copy(sortByStart, intervals)
	sort.SliceStable(sortByStart, func(i, j int) bool {
		return sortByStart[i][0] < sortByStart[j][0]
	})
	sortByEnd := make([][]int, len(intervals))
	copy(sortByEnd, intervals)
	sort.SliceStable(sortByEnd, func(i, j int) bool {
		return sortByEnd[i][1] < sortByEnd[j][1]
	})
	maxRooms := 0
	currentRooms := 0
	for i, j := 0, 0; i < len(sortByStart) && j < len(sortByEnd); {
		if sortByStart[i][0] < sortByEnd[j][1] {
			i++
			currentRooms++
		} else if sortByStart[i][0] > sortByEnd[j][1] {
			j++
			currentRooms--
		} else {
			i++
			j++
		}
		if currentRooms > maxRooms {
			maxRooms = currentRooms
		}
	}
	return maxRooms
}

func main() {
	val := minMeetingRooms([][]int{
		[]int{0, 30},
		[]int{5, 10},
		[]int{15, 20},
	})
	fmt.Println(val)
}
