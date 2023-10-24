package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateImage90Degress(t *testing.T) {
	image := [][]Pixel{
		{Pixel{0, 0, 0, 1}, Pixel{0, 0, 1, 0}, Pixel{0, 1, 0, 0}},
		{Pixel{2, 0, 0, 1}, Pixel{0, 2, 1, 0}, Pixel{0, 2, 0, 0}},
		{Pixel{0, 0, 3, 1}, Pixel{0, 0, 3, 0}, Pixel{0, 1, 3, 0}},
	}
	expectedRotatedImage := [][]Pixel{
		{Pixel{0, 0, 3, 1}, Pixel{2, 0, 0, 1}, Pixel{0, 0, 0, 1}},
		{Pixel{0, 0, 3, 0}, Pixel{0, 2, 1, 0}, Pixel{0, 0, 1, 0}},
		{Pixel{0, 1, 3, 0}, Pixel{0, 2, 0, 0}, Pixel{0, 1, 0, 0}},
	}
	rotateImage90Degress(&image)
	for x := 0; x < len(image); x++ {
		for y := 0; y < len(image[x]); y++ {
			assert.Equal(t, expectedRotatedImage[y][x], image[y][x])
		}
	}
}

func TestRotateImage90DegressInPlace1(t *testing.T) {
	image := [][]Pixel{
		{Pixel{0, 0, 0, 1}, Pixel{0, 0, 1, 0}, Pixel{0, 1, 0, 0}},
		{Pixel{2, 0, 0, 1}, Pixel{0, 2, 1, 0}, Pixel{0, 2, 0, 0}},
		{Pixel{0, 0, 3, 1}, Pixel{0, 0, 3, 0}, Pixel{0, 1, 3, 0}},
	}
	expectedRotatedImage := [][]Pixel{
		{Pixel{0, 0, 3, 1}, Pixel{2, 0, 0, 1}, Pixel{0, 0, 0, 1}},
		{Pixel{0, 0, 3, 0}, Pixel{0, 2, 1, 0}, Pixel{0, 0, 1, 0}},
		{Pixel{0, 1, 3, 0}, Pixel{0, 2, 0, 0}, Pixel{0, 1, 0, 0}},
	}
	rotateImage90DegressInPlace1(&image)
	for x := 0; x < len(image); x++ {
		for y := 0; y < len(image[x]); y++ {
			assert.Equal(t, expectedRotatedImage[y][x], image[y][x])
		}
	}
}

func TestRotateImage90DegressInPlace2(t *testing.T) {
	image := [][]Pixel{
		{Pixel{0, 0, 0, 1}, Pixel{0, 0, 1, 0}, Pixel{0, 1, 0, 0}},
		{Pixel{2, 0, 0, 1}, Pixel{0, 2, 1, 0}, Pixel{0, 2, 0, 0}},
		{Pixel{0, 0, 3, 1}, Pixel{0, 0, 3, 0}, Pixel{0, 1, 3, 0}},
	}
	expectedRotatedImage := [][]Pixel{
		{Pixel{0, 0, 3, 1}, Pixel{2, 0, 0, 1}, Pixel{0, 0, 0, 1}},
		{Pixel{0, 0, 3, 0}, Pixel{0, 2, 1, 0}, Pixel{0, 0, 1, 0}},
		{Pixel{0, 1, 3, 0}, Pixel{0, 2, 0, 0}, Pixel{0, 1, 0, 0}},
	}
	rotateImage90DegressInPlace2(&image)
	for x := 0; x < len(image); x++ {
		for y := 0; y < len(image[x]); y++ {
			assert.Equal(t, expectedRotatedImage[y][x], image[y][x])
		}
	}

	image2 := [][]Pixel{
		{Pixel{0, 0, 0, 1}, Pixel{0, 0, 1, 0}, Pixel{0, 1, 0, 0}, Pixel{0, 9, 0, 0}},
		{Pixel{2, 0, 0, 1}, Pixel{0, 2, 1, 0}, Pixel{0, 2, 0, 0}, Pixel{0, 5, 0, 0}},
		{Pixel{0, 0, 3, 1}, Pixel{0, 0, 3, 0}, Pixel{0, 1, 3, 0}, Pixel{0, 2, 3, 0}},
		{Pixel{0, 91, 3, 1}, Pixel{0, 103, 3, 0}, Pixel{0, 21, 3, 0}, Pixel{0, 83, 3, 0}},
	}
	expectedRotatedImage2 := [][]Pixel{
		{Pixel{0, 91, 3, 1}, Pixel{0, 0, 3, 1}, Pixel{2, 0, 0, 1}, Pixel{0, 0, 0, 1}},
		{Pixel{0, 103, 3, 0}, Pixel{0, 0, 3, 0}, Pixel{0, 2, 1, 0}, Pixel{0, 0, 1, 0}},
		{Pixel{0, 21, 3, 0}, Pixel{0, 1, 3, 0}, Pixel{0, 2, 0, 0}, Pixel{0, 1, 0, 0}},
		{Pixel{0, 83, 3, 0}, Pixel{0, 2, 3, 0}, Pixel{0, 5, 0, 0}, Pixel{0, 9, 0, 0}},
	}
	rotateImage90DegressInPlace2(&image2)
	for x := 0; x < len(image2); x++ {
		for y := 0; y < len(image2[x]); y++ {
			assert.Equal(t, expectedRotatedImage2[y][x], image2[y][x])
		}
	}
}

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
