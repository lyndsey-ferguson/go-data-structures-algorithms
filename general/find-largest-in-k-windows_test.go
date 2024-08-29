package general

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxElementInKayWindows(t *testing.T) {
	arr := []int{1, 2, 3, 1, 4, 5, 2, 3, 6}
	k := 3

	res := findMaxElementInKayWindows(arr, k)
	assert.Equal(t, []int{3, 3, 4, 5, 5, 5, 6}, res)
}
