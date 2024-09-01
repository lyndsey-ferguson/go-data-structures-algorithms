package general

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollapseIntervals(t *testing.T) {
	intervals := []Interval{
		{1, 4},
		{1, 1},
		{2, 5},
		{4, 5},
	}
	expectedIntervals := []Interval{{1, 5}}
	actualIntervals := CollapseIntervals(intervals)
	assert.Equal(t, expectedIntervals, actualIntervals)

	intervals = []Interval{
		{1, 4},
		{1, 1},
		{2, 5},
		{10, 14},
		{11, 18},
	}

	expectedIntervals = []Interval{{1, 5}, {10, 18}}
	actualIntervals = CollapseIntervals(intervals)
	assert.Equal(t, expectedIntervals, actualIntervals)
}

func TestInstantWithinTimeIntervals(t *testing.T) {
	intervals := []Interval{
		{1, 4},
		{1, 1},
		{2, 5},
		{10, 14},
		{11, 18},
	}

	assert.True(t, IntervalWithinTimeIntervals(3, intervals))
	assert.False(t, IntervalWithinTimeIntervals(19, intervals))
	assert.False(t, IntervalWithinTimeIntervals(0, intervals))
}
