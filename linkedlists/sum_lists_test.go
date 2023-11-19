package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumListsInOnePass(t *testing.T) {
	list1 := createList([]int32{6, 1, 7})
	list2 := createList([]int32{2, 9, 5})

	expectedList := createList([]int32{9, 1, 2})

	assert.Equal(t, expectedList.ToString(), sumListsInOnePass(list1, list2).ToString())
}

func TestSumListsInOnePassWithCarry(t *testing.T) {
	list1 := createList([]int32{9, 2, 7})
	list2 := createList([]int32{9, 9, 4})

	expectedList := createList([]int32{1, 9, 2, 1})

	assert.Equal(t, expectedList.ToString(), sumListsInOnePass(list1, list2).ToString())
}
