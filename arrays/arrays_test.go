package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	array := []int{12, 11, 13, 5, 6, 7}
	mergedArray := array

	MergeSort(&mergedArray, 0)
	assert.Equal(t, []int{5, 6, 7, 11, 12, 13}, mergedArray)
}

func TestQuickSort(t *testing.T) {
	array := []int{10, 7, 8, 9, 1, 5}
	QuickSort(array, 0, len(array)-1)
	assert.Equal(t, []int{1, 5, 7, 8, 9, 10}, array)
}
