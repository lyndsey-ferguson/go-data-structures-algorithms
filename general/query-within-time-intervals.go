package general

import (
	"sort"
)

type Interval struct {
	startTime int32
	endTime   int32
}

func CollapseIntervals(intervals []Interval) []Interval {
	collapsedIntervals := []Interval{intervals[0]}
	intervals = intervals[1:]
	for len(intervals) > 0 {
		interval := intervals[0]
		lastInterval := collapsedIntervals[len(collapsedIntervals)-1]
		collapsedIntervals = collapsedIntervals[:len(collapsedIntervals)-1]
		if interval.startTime >= lastInterval.startTime && interval.startTime <= lastInterval.endTime {
			// candidate for consolidation
			lastInterval.endTime = max(lastInterval.endTime, interval.endTime)
			collapsedIntervals = append(collapsedIntervals, lastInterval)
		} else {
			// new interval range
			collapsedIntervals = append(collapsedIntervals, lastInterval, interval)
		}
		intervals = intervals[1:]
	}
	return collapsedIntervals
}

func binarySearch(a []Interval, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case int(a[mid].endTime) < search:
		result, searchCount = binarySearch(a[mid+1:], search)
	case int(a[mid].startTime) > search:
		result, searchCount = binarySearch(a[:mid], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}

func IntervalWithinTimeIntervals(i int, intervals []Interval) bool {
	// 1. sort the list
	sort.SliceStable(
		intervals,
		func(i, j int) bool {
			return intervals[i].startTime < intervals[j].startTime
		},
	)
	// 2. collapse the list ranges that overlap
	collapsedIntervals := CollapseIntervals(intervals)

	// 3. binary search intervals
	result, _ := binarySearch(collapsedIntervals, i)
	return result != -1
}
