package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	array := []int{10, 7, 8, 9, 1, 5}
	QuickSort(array, 0, len(array)-1)
	assert.Equal(t, []int{1, 5, 7, 8, 9, 10}, array)
}
