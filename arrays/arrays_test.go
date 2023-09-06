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
