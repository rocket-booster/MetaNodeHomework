package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	merged := [][]int{intervals[0]}
	
	for i := 1; i < len(intervals); i++ {		
		current := intervals[i]
		last := merged[len(merged)-1]
		
		if current[0] <= last[1] {			
			merged[len(merged)-1][1] = max(last[1], current[1])
		} else {			
			merged = append(merged, current)
		}
	}

	return merged
}



func main() {	
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 5}},
		{{1, 4}, {2, 3}},
	}
	
	for i, intervals := range testCases {
		result := merge(intervals)
		fmt.Printf("测试用例 %d: %v\n合并结果: %v\n\n", i+1, intervals, result)
	}
}