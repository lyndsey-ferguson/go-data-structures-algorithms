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
