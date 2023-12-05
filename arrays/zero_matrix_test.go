package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroOutAxisInMatrix(t *testing.T) {
	m1 := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	expected_m1_result := [][]int{
		{1, 0, 3},
		{0, 0, 0},
		{7, 0, 9},
	}

	zeroOutAxisInMatrix(m1)
	for r := 0; r < len(expected_m1_result); r++ {
		for c := 0; c < len(expected_m1_result[0]); c++ {
			assert.Equal(t, expected_m1_result[r][c], m1[r][c])
		}
	}

	m2 := [][]int{
		{0, 2, 0},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected_m2_result := [][]int{
		{0, 0, 0},
		{0, 5, 0},
		{0, 8, 0},
	}

	zeroOutAxisInMatrix(m2)
	for r := 0; r < len(expected_m2_result); r++ {
		for c := 0; c < len(expected_m2_result[0]); c++ {
			assert.Equal(t, expected_m2_result[r][c], m2[r][c])
		}
	}
}
