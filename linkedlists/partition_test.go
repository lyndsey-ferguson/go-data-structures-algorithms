package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionIntList(t *testing.T) {
	list := createList([]int32{3, 5, 8, 5, 10, 2, 1})
	expectedList := createList([]int32{3, 2, 1, 5, 10, 5, 8})

	partitionIntList(list, 5)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}

func TestPartitionIntListWithNoSmallerItems(t *testing.T) {
	list := createList([]int32{3, 5, 8, 5, 10, 7, 12})
	expectedList := createList([]int32{3, 5, 8, 5, 10, 7, 12})

	partitionIntList(list, 2)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}

func TestPartitionIntListWithNoEqualOrGreaterItems(t *testing.T) {
	list := createList([]int32{3, 5, 8, 5, 10, 7, 12})
	expectedList := createList([]int32{3, 5, 8, 5, 10, 7, 12})

	partitionIntList(list, 20)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}

func TestPartitionIntListMerging(t *testing.T) {
	list := createList([]int32{3, 5, 8, 5, 10, 2, 1})
	expectedList := createList([]int32{3, 2, 1, 5, 8, 5, 10})

	partitionIntListMerging(&list, 5)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}

func TestPartitionIntListMergingWithNoSmallerItems(t *testing.T) {
	list := createList([]int32{3, 5, 8, 5, 10, 7, 12})
	expectedList := createList([]int32{3, 5, 8, 5, 10, 7, 12})

	partitionIntListMerging(&list, 2)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}

func TestPartitionIntListMergingWithNoEqualOrGreaterItems(t *testing.T) {
	list := createList([]int32{3, 5, 8, 5, 10, 7, 12})
	expectedList := createList([]int32{3, 5, 8, 5, 10, 7, 12})

	partitionIntListMerging(&list, 20)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}
