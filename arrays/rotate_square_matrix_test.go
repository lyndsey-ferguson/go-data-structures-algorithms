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
